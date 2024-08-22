package unit

import (
	"encoding/json"
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"github.com/TrashPony/veliri-lib/game_objects/effect"
	"github.com/TrashPony/veliri-lib/game_objects/obstacle_point"
	"github.com/TrashPony/veliri-lib/game_objects/target"
	"math"
	"strconv"
	"sync/atomic"
)

type StateMS struct {
	UseEnergy    int `json:"use_energy"`
	MaxEnergy    int `json:"max_energy"`
	CapacitySize int `json:"capacity_size"`

	// живучесть
	MaxHP                       int `json:"max_hp"`
	ProtectionToKinetics        int `json:"protection_to_kinetics"`
	ProtectionToThermo          int `json:"protection_to_thermo"`
	ProtectionToExplosion       int `json:"protection_to_explosion"`
	MaxShieldHP                 int `json:"max_shield_hp"`
	ShieldProtectionToKinetics  int `json:"shield_protection_to_kinetics"`
	ShieldProtectionToThermo    int `json:"shield_protection_to_thermo"`
	ShieldProtectionToExplosion int `json:"shield_protection_to_explosion"`

	// движение
	Speed         float64 `json:"speed"`          // -- макс скорость вперед
	ReverseSpeed  float64 `json:"reverse_speed"`  //-- макс скорость назад
	PowerFactor   float64 `json:"power_factor"`   // -- сила ускорения вперед
	ReverseFactor float64 `json:"reverse_factor"` //-- сила ускорения назад
	TurnSpeed     float64 `json:"turn_speed"`     //-- скорость поворота в радианах
	CurrentSpeed  float64 `json:"current_speed"`

	// навигация
	RangeView  int `json:"range_view"`
	RangeRadar int `json:"range_radar"`

	// генератор
	MaxPower           int `json:"max_power"`
	RecoveryPower      int `json:"recovery_power"`
	RecoveryPowerCycle int `json:"recovery_power_cycle"`

	// оружие
	WeaponType        string  `json:"wt"`
	WeaponSize        int     `json:"ws"`
	BodyType          string  `json:"bt"`
	BodySize          int     `json:"bs"`
	WeaponInstall     bool    `json:"install"`
	WeaponMaxDamage   int     `json:"weapon_max_damage"`
	WeaponMinDamage   int     `json:"weapon_min_damage"`
	CountFireBullet   int     `json:"count_fire_bullet"`
	WeaponMinRange    int     `json:"weapon_min_range"`
	WeaponMaxRange    int     `json:"weapon_max_range"`
	WeaponAccuracy    int     `json:"weapon_accuracy"`
	GunRotateSpeed    int     `json:"gun_rotate_speed"`
	BulletSpeed       float64 `json:"bullet_speed"`
	ReloadTime        int     `json:"reload_time"`
	ReloadAmmoTime    int     `json:"reload_ammo_time"`
	EfficiencyReactor float64 `json:"efficiency_reactor"`
	KineticsDamage    int     `json:"kinetics_damage"`
	ExplosionDamage   int     `json:"explosion_damage"`
	ThermoDamage      int     `json:"thermo_damage"`
}

func (u *Unit) GetState() *StateMS {

	var weapon *detail.Weapon
	if u.GetWeaponSlot(1) != nil {
		weapon = u.GetWeaponSlot(1).Weapon
	}

	state := StateMS{
		UseEnergy:                   u.getBody().GetUseEnergy(),
		MaxEnergy:                   u.GetMaxEnergy(),
		MaxHP:                       u.GetMaxHP(),
		ProtectionToKinetics:        u.GetProtection("kinetics"),
		ProtectionToThermo:          u.GetProtection("thermo"),
		ProtectionToExplosion:       u.GetProtection("explosion"),
		MaxShieldHP:                 u.GetMaxShieldHP(),
		ShieldProtectionToKinetics:  u.GetProtection("shield_kinetics"),
		ShieldProtectionToThermo:    u.GetProtection("shield_thermo"),
		ShieldProtectionToExplosion: u.GetProtection("shield_explosion"),
		Speed:                       u.GetMoveMaxPower(),
		ReverseSpeed:                u.GetMaxReverse(),
		PowerFactor:                 u.GetPowerFactor(),
		ReverseFactor:               u.GetReverseFactor(),
		TurnSpeed:                   u.GetTurnSpeed(),
		RangeView:                   u.GetRangeView(),
		RangeRadar:                  u.GetRadarRange(),
		WeaponInstall: u.GetWeaponSlot(1) != nil &&
			weapon != nil &&
			u.GetWeaponSlot(1).GetAmmo() != nil &&
			u.GetWeaponSlot(1).GetAmmoQuantity() > 0,
		EfficiencyReactor:  u.EfficiencyReactor(),
		CurrentSpeed:       u.GetPhysicalModel().GetCurrentSpeed(),
		CapacitySize:       u.GetCapSize(),
		MaxPower:           u.GetMaxPower(),
		RecoveryPower:      u.GetRecoveryPower(),
		RecoveryPowerCycle: u.body.RecoveryPowerCycle,
		BodyType:           u.body.ChassisType,
		BodySize:           u.body.StandardSize,
	}

	if state.WeaponInstall {
		state.WeaponType = weapon.Type
		state.WeaponSize = weapon.StandardSize
		state.WeaponMaxDamage = u.GetGunner().GetMaxDamage(u.GetWeaponSlot(1).Number)
		state.WeaponMinDamage = u.GetGunner().GetMinDamage(u.GetWeaponSlot(1).Number)
		state.CountFireBullet = u.GetGunner().GetCountFireBullet(u.GetWeaponSlot(1).Number)
		state.WeaponMinRange, _ = u.GetGunner().GetWeaponMinRange(0.0, u.GetWeaponSlot(1).Number)
		state.WeaponMaxRange, _ = u.GetGunner().GetWeaponMaxRange(0.0, u.GetWeaponSlot(1).Number, false)
		state.WeaponAccuracy = u.GetGunner().GetWeaponAccuracy(u.GetWeaponSlot(1).Number)
		state.GunRotateSpeed = u.GetGunner().GetGunRotateSpeed(u.GetWeaponSlot(1).Number)
		state.BulletSpeed = u.GetGunner().GetBulletSpeed(u.GetWeaponSlot(1).Number)
		state.ReloadTime = u.GetGunner().GetWeaponReloadTime(u.GetWeaponSlot(1).Number)
		state.ReloadAmmoTime = u.GetGunner().GetWeaponReloadAmmoTime(u.GetWeaponSlot(1).Number)

		k, t, e := u.GetGunner().GetDamageType(u.GetWeaponSlot(1).Number)
		state.KineticsDamage = k
		state.ThermoDamage = t
		state.ExplosionDamage = e
	}

	return &state
}

// GetAllState метод возвращает полное состоиня машинки, влияние навыков, влияние снаряжения штрафы
func (u *Unit) GetAllState() (all *StateMS, effects []*effect.Effect) {

	all = u.GetState()
	effects = u.GetEffects().GetAllEffects()

	return all, effects
}

func (u *Unit) GetProtection(typeVulnerabilities string) int {

	var startValue float64

	if typeVulnerabilities == "thermo" {
		startValue = float64(u.body.ProtectionToThermo)
	}

	if typeVulnerabilities == "kinetics" {
		startValue = float64(u.body.ProtectionToKinetics)
	}

	if typeVulnerabilities == "explosion" {
		startValue = float64(u.body.ProtectionToExplosion)
	}

	if typeVulnerabilities == "shield_thermo" {
		startValue = float64(u.body.ShieldProtectionToThermo)
	}

	if typeVulnerabilities == "shield_kinetics" {
		startValue = float64(u.body.ShieldProtectionToKinetics)
	}

	if typeVulnerabilities == "shield_explosion" {
		startValue = float64(u.body.ShieldProtectionToExplosion)
	}

	for _, e := range u.GetEffects().GetAllEffects() {
		if e.Parameter == "protect_"+typeVulnerabilities {
			free := 100 - startValue

			if !e.Subtract {
				startValue += free * (float64(e.Quantity) / 100)
			} else {
				startValue -= free * (float64(e.Quantity) / 100)
			}
		}
	}

	if int(startValue) < 0 {
		return 0
	}

	return int(startValue)
}

func (u *Unit) GetOwnerID() int {
	return u.OwnerID
}

func (u *Unit) GetOwnerPlayerID() int {
	return u.OwnerID
}

func (u *Unit) GetPower() int {
	return u.Power
}

func (u *Unit) GetID() int {
	return u.ID
}

func (u *Unit) GetType() string {
	return "unit"
}

func (u *Unit) GetRotate() float64 {
	return u.GetPhysicalModel().Rotate
}

func (u *Unit) GetMapHeight() float64 {
	return _const.UnitHeight
}

func (u *Unit) GetStateSenderPos() bool {
	return false
}

func (u *Unit) GetWidth() float64 {
	return float64(u.body.Width) * 2
}

func (u *Unit) GetLength() float64 {
	return float64(u.body.Length) * 2
}

func (u *Unit) GetHeight() float64 {
	return float64(u.body.Height) * 2
}

func (u *Unit) GetGeoData() []*obstacle_point.ObstaclePoint {
	return nil
}

func (u *Unit) GetWeight() float64 {
	return u.body.Weight
}

func (u *Unit) GetMoveMaxPower() float64 {
	return u.GetEffects().GetAllBodyBonus(u.body.Speed, "speed", u.getBody().ChassisType, u.getBody().StandardSize)
}

func (u *Unit) GetMaxReverse() float64 {
	return u.GetEffects().GetAllBodyBonus(u.body.ReverseSpeed, "reverse_speed", u.getBody().ChassisType, u.getBody().StandardSize)
}

func (u *Unit) GetPowerFactor() float64 {
	return u.GetEffects().GetAllBodyBonus(u.body.PowerFactor, "power_factor", u.getBody().ChassisType, u.getBody().StandardSize)
}

func (u *Unit) GetReverseFactor() float64 {
	return u.GetEffects().GetAllBodyBonus(u.body.ReverseFactor, "reverse_factor", u.getBody().ChassisType, u.getBody().StandardSize)
}

func (u *Unit) GetTurnSpeed() float64 {
	return u.GetEffects().GetAllBodyBonus(u.body.TurnSpeed, "turn_speed", u.getBody().ChassisType, u.getBody().StandardSize)
}

func (u *Unit) GetCapSize() int {
	if u == nil {
		return 0
	}

	return int(math.Ceil(u.GetEffects().GetAllBodyBonus(float64(u.body.CapacitySize), "cap_size", u.getBody().ChassisType, u.getBody().StandardSize)))
}

func (u *Unit) GetAdditionalCapSize(key string) int {
	if u == nil {
		return 0
	}

	inv := u.body.AdditionalInventory[key]
	return int(math.Ceil(u.GetEffects().GetAllBodyBonus(float64(inv.GetCapSize()), key+"_cap_size", u.getBody().ChassisType, u.getBody().StandardSize)))
}

func (u *Unit) GetFullFreeCapSize(itemType string, itemID int) int {
	freeSize := u.GetCapSize() - u.Inventory.GetSize()

	for key, inv := range u.AdditionalInventory {
		filter := u.body.AdditionalInventory[key].Filter
		for fType, ids := range filter {
			if fType == itemType {
				for _, id := range ids {
					if id == itemID {
						freeSize += u.GetAdditionalCapSize(key) - inv.GetSize()
					}
				}
			}
		}
	}

	return freeSize
}

func (u *Unit) GetRangeView() int {
	return u.viewRange
}

func (u *Unit) getRangeView() int {
	return int(math.Ceil(u.GetEffects().GetAllBodyBonus(float64(u.body.RangeView), "view", u.getBody().ChassisType, u.getBody().StandardSize)))
}

func (u *Unit) GetRadarRange() int {
	return u.radarRange
}

func (u *Unit) getRadarRange() int {
	return int(math.Ceil(u.GetEffects().GetAllBodyBonus(float64(u.body.RangeRadar), "radar", u.getBody().ChassisType, u.getBody().StandardSize)))
}

func (u *Unit) GetMaxHP() int {
	return int(math.Ceil(u.GetEffects().GetAllBodyBonus(float64(u.body.MaxHP), "max_hp", u.getBody().ChassisType, u.getBody().StandardSize)))
}

func (u *Unit) GetMaxShieldHP() int {
	return int(math.Ceil(u.GetEffects().GetAllBodyBonus(float64(u.body.MaxShieldHP), "max_shield_hp", u.getBody().ChassisType, u.getBody().StandardSize)))
}

func (u *Unit) GetMaxPower() int {
	return int(math.Ceil(u.GetEffects().GetAllBodyBonus(float64(u.body.MaxPower), "max_power", u.getBody().ChassisType, u.getBody().StandardSize)))
}

func (u *Unit) GetMaxEnergy() int {
	return int(math.Ceil(u.GetEffects().GetAllBodyBonus(float64(u.body.MaxEnergy), "max_energy", u.getBody().ChassisType, u.getBody().StandardSize)))
}

// weapon
func (u *Unit) getGunAccuracy(weaponSlotNumber int) int {
	weaponSlot := u.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return 0
	}
	return int(math.Ceil(u.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Weapon.Accuracy), "accuracy", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (u *Unit) getGunRotateSpeed(weaponSlotNumber int) int {
	weaponSlot := u.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return 0
	}
	return int(math.Ceil(u.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Weapon.RotateSpeed), "gun_speed_rotate", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (u *Unit) getWeaponReloadTime(weaponSlotNumber int) int {
	weaponSlot := u.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return 0
	}
	return int(math.Ceil(u.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Weapon.ReloadTime), "reload", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (u *Unit) getWeaponAmmoReloadTime(weaponSlotNumber int) int {
	weaponSlot := u.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return 0
	}
	return int(math.Ceil(u.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Weapon.ReloadAmmoTime), "reload_ammo", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (u *Unit) getMaxDamage(weaponSlotNumber int) int {
	weaponSlot := u.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil || weaponSlot.Ammo == nil {
		return 0
	}
	return int(math.Ceil(u.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Ammo.MaxDamage), "damage", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (u *Unit) getMinDamage(weaponSlotNumber int) int {
	weaponSlot := u.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil || weaponSlot.Ammo == nil {
		return 0
	}
	return int(math.Ceil(u.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Ammo.MinDamage), "damage", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (u *Unit) AppendWeaponDamageModifier(weaponSlotNumber int) {
	u.RemoveEffect("weapon_damage_modifier")
	weaponSlot := u.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot != nil && weaponSlot.Weapon != nil && weaponSlot.Weapon.DamageModifier != 1 {
		damageModifier := 100 - (weaponSlot.Weapon.DamageModifier * 100)
		subtract := damageModifier > 0
		if damageModifier < 0 {
			damageModifier *= -1
		}

		u.AddEffect(&effect.Effect{
			UUID:        "weapon_damage_modifier",
			Parameter:   "damage",
			Quantity:    int(damageModifier),
			Percentages: true,
			Subtract:    subtract,
		})
	}
}

const stackDebuf = 15

func (u *Unit) AppendPassiveEquipModifier() {

	for _, slot := range u.body.GetAllEquipSlots() {
		u.RemoveBySlot(slot.Type, slot.Number)
	}

	for _, slot := range u.body.GetAllEquips() {
		if slot.Equip != nil && !slot.Equip.Active && len(slot.Equip.Effects) > 0 {
			for _, e := range slot.Equip.Effects {

				quantity := e.Quantity
				count := u.body.FindEquipByID(slot.Equip.ID)
				if len(count) > 1 {
					debuf := 100 - float64(stackDebuf*(len(count)-1))
					quantity = int(math.Floor(float64(quantity) * (debuf / 100)))
				}

				u.AddEffect(&effect.Effect{
					UUID:         "equip_passive_" + strconv.Itoa(slot.Type) + ":" + strconv.Itoa(slot.Number) + e.Name,
					Parameter:    e.Parameter,
					Quantity:     quantity,
					Percentages:  e.Percentages,
					Subtract:     e.Subtract,
					SlotType:     slot.Type,
					SlotNumber:   slot.Number,
					WeaponType:   e.WeaponType,
					BodyType:     e.BodyType,
					StandardSize: e.StandardSize,
					Fraction:     e.Fraction,
				})
			}
		}
	}
}

func (u *Unit) GetEvacuation() bool {
	return u.evacuation
}

func (u *Unit) GetForceEvacuation() bool {
	return u.forceEvacuation
}

func (u *Unit) GetHP() int {
	return u.HP
}

func (u *Unit) GetX() int {
	return u.GetPhysicalModel().GetX()
}

func (u *Unit) GetY() int {
	return u.GetPhysicalModel().GetY()
}

func (u *Unit) GetMapID() int {
	return int(atomic.LoadInt32(&u.MapID))
}

func (u *Unit) InSky() bool {
	return u.inSky
}

func (u *Unit) GetFullJSON() string {
	jsonUnit, err := json.Marshal(u)
	if err != nil {
		println("unit to json: ", err.Error())
	}

	return string(jsonUnit)
}

func (u *Unit) GetFollowTarget() *target.Target {
	mp := u.movePath
	if mp != nil {
		return mp.GetFollowTarget()
	}

	return nil
}

func (u *Unit) GetRecoveryPower() int {
	return int(math.Ceil(u.GetEffects().GetAllBodyBonus(float64(u.getBody().RecoveryPower), "charging_speed", u.getBody().ChassisType, u.getBody().StandardSize)))
}

func (u *Unit) GetUnrepairableDamage() int {
	maxUD := 2 * (u.GetMaxHP() / 3)
	ud := int(25 * (float64(u.UnrepairableDamage) / 100))

	if maxUD < ud {
		return maxUD
	}

	return ud
}

func (u *Unit) AddUrepairableDamage(d int) {
	percentUD := 90 - ((float64(u.GetUnrepairableDamage()) / float64(u.GetMaxHP())) * 100)
	if percentUD <= 0 {
		return
	}

	add := int(float64(d) * percentUD / 100)
	if add < 0 {
		add = 1
	}
	u.UnrepairableDamage += add
}
