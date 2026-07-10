package unit

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"github.com/TrashPony/veliri-lib/game_objects/effect"
	"strconv"
	"sync"
)

var moduleToByte = map[string]byte{
	"chassis":          1,
	"weapon":           2,
	"reactor":          3,
	"engine":           4,
	"ammo_storage":     5,
	"sensors":          6,
	"shield_generator": 7,
}

type damageModuleManager struct {
	damageModules []*detail.DamageModule
	mx            sync.Mutex
}

const TimeRepairDamageModule = 3000

func (u *Unit) RepairDamageModules() {
	u.damageModules.mx.Lock()
	defer u.damageModules.mx.Unlock()
	u.initDamageModules(true)
}

// инициируем нулевую сборку из корпуса на случай если таковой нет
func (u *Unit) initDamageModules(force bool) {
	if u.damageModules.damageModules == nil || len(u.damageModules.damageModules) == 0 || force {
		u.damageModules.damageModules = make([]*detail.DamageModule, 0, len(u.body.DamageModules))

		for i, m := range u.body.DamageModules {
			copyModule := m
			m.RepairTime = 0
			m.Damage = false
			m.ID = byte(i) + 1
			copyModule.CurrentHP = copyModule.MaxHP
			u.damageModules.damageModules = append(u.damageModules.damageModules, &copyModule)
		}
	} else {
		for i, m := range u.damageModules.damageModules {
			m.ID = byte(i) + 1
		}
	}
}

// SetDamageModules загружаем например из бд или при перебросе на другую ноду
func (u *Unit) SetDamageModules(modules []*detail.DamageModule) {
	u.damageModules.mx.Lock()
	defer u.damageModules.mx.Unlock()
	u.damageModules.damageModules = modules
	u.appendModuleDebuff()
}

// GetDamageModules забираем для проброса по апи или в бд
func (u *Unit) GetDamageModules() []*detail.DamageModule {
	u.damageModules.mx.Lock()
	defer u.damageModules.mx.Unlock()
	u.initDamageModules(false)

	copyArr := make([]*detail.DamageModule, 0)
	for _, m := range u.damageModules.damageModules {
		copyArr = append(copyArr, m)
	}

	return copyArr
}

// DamageModule наносим урон зоне, мы должны вернуть тип зоны и ее позицию (угол между стартов и концом) в случае ее уничтожения что бы можно было отыграть анимации и наложить соотвествующие эффекты
func (u *Unit) DamageModule(bulletX, bulletY, damage int) (string, int, byte) {
	u.damageModules.mx.Lock()
	defer u.damageModules.mx.Unlock()

	u.initDamageModules(false)

	angle := game_math.GetBetweenAngle(float64(bulletX), float64(bulletY), float64(u.GetPhysicalModel().GetX()), float64(u.GetPhysicalModel().GetY()))
	game_math.PrepareAngle(&angle)

	angle -= u.GetPhysicalModel().GetRotate()
	game_math.PrepareAngle(&angle)

	z := u.getDamageModuleByAngle(int(angle))
	if z != nil {
		// уже уничтожен
		if z.RepairTime > 0 {
			return "", 0, 0
		}

		if z.Type == "weapon" || z.Type == "ammo_storage" {
			weaponFind := false
			ammoFind := false
			for _, ws := range u.RangeWeaponSlots() {
				if ws.Weapon != nil {
					weaponFind = true

					if ws.Ammo != nil && ws.AmmoQuantity > 0 {
						ammoFind = true
					}
				}
			}

			if z.Type == "weapon" && !weaponFind {
				return "", 0, 0
			}

			if z.Type == "ammo_storage" && !ammoFind {
				return "", 0, 0
			}
		}

		z.CurrentHP -= damage
		if z.CurrentHP <= 0 {

			z.Damage = true
			z.RepairTime = TimeRepairDamageModule
			z.CurrentHP = z.MaxHP / 2

			u.appendModuleDebuff()
			return z.Type, getModuleCenter(z.StartAngle, z.EndAngle), z.ID
		}
	}

	return "", 0, 0
}

func getModuleCenter(start, end int) int {
	if start <= end {
		// Обычный случай: 30-60
		return (start + end) / 2
	} else {
		// Пересекает 0°: 350-10
		// Сдвигаем end на 360
		end += 360
		center := (start + end) / 2
		// Нормализуем обратно в 0-359
		return center % 360
	}
}

func (u *Unit) getDamageModuleByAngle(angle int) *detail.DamageModule {
	for _, z := range u.damageModules.damageModules {
		if z.StartAngle <= z.EndAngle {
			// Обычный случай: модуль не пересекает 0°
			// Проверяем: start <= angle <= end
			if z.StartAngle <= angle && z.EndAngle >= angle {
				return z
			}
		} else {
			// Модуль пересекает 0° (например, 350-10)
			// Проверяем: angle >= start ИЛИ angle <= end
			if angle >= z.StartAngle || angle <= z.EndAngle {
				return z
			}
		}
	}

	return nil
}

func (u *Unit) AppendModuleDebuff() {
	u.damageModules.mx.Lock()
	defer u.damageModules.mx.Unlock()
	u.appendModuleDebuff()
}

func (u *Unit) appendModuleDebuff() {
	u.initDamageModules(false)

	angularVelocityTilt := 0.0
	for _, z := range u.damageModules.damageModules {
		if z.Damage {
			for _, p := range getParameterByType(z.Type) {
				u.AddEffect(&effect.Effect{
					UUID:        "damage_module_" + p + z.Type + strconv.Itoa(z.StartAngle) + strconv.Itoa(z.EndAngle),
					Parameter:   p,
					Quantity:    z.Count,
					Percentages: true,
					Subtract:    true,
				})
			}

			if z.Type == "chassis" {
				angularVelocityTilt += u.GetPhysicalModel().TurnSpeed * float64(game_math.GetRangeRand(0, 50, nil)-25) * 0.01
			}
		}
	}

	u.GetPhysicalModel().AngularVelocityTilt = angularVelocityTilt
}

func getParameterByType(moduleType string) []string {
	switch moduleType {
	case "chassis":
		return []string{"turn_speed", "reverse_factor", "power_factor", "reverse_speed", "speed"}
	case "weapon":
		return []string{"damage", "gun_speed_rotate"}
	case "reactor":
		return []string{"charging_speed"}
	case "engine":
		return []string{"reverse_factor", "power_factor", "reverse_speed", "speed"}
	case "ammo_storage":
		return []string{"reload_ammo", "reload"}
	case "sensors":
		return []string{"view"}
	case "shield_generator":
		return []string{"off_body_shield"}
	default:
		return []string{}
	}
}

func (u *Unit) UpdateDamageModules() {
	u.damageModules.mx.Lock()
	defer u.damageModules.mx.Unlock()

	for _, m := range u.damageModules.damageModules {
		if m.RepairTime > 0 {
			m.RepairTime -= _const.ServerTick
		}
	}
}

func (u *Unit) GetBinDataModules() []byte {
	u.damageModules.mx.Lock()
	defer u.damageModules.mx.Unlock()

	bytes := make([]byte, 0)

	bytes = append(bytes, byte(len(u.damageModules.damageModules)))
	for _, m := range u.damageModules.damageModules {
		bytes = append(bytes, moduleToByte[m.Type])
		bytes = append(bytes, byte((float64(m.CurrentHP)/float64(m.MaxHP))*100))
	}

	return bytes
}
