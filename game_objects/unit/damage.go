package unit

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/damage_manager"
)

func (u *Unit) SetDamage(damage, k, t, e int, ignoreShield bool, callback func(playerID, startDamage, kDamage, tDamage, eDamage, finalDamage, startSHP, finalSHP, startHP, finalHP, k, t, e int)) (bool, bool, int, int) {

	var prefix string
	if u.ShieldHP > 0 && !ignoreShield {
		prefix = "shield_"
	}

	startDamage := damage
	startSHP := u.ShieldHP
	startHP := u.GetHP()

	kDamage := float64(damage) * (float64(k) / 100.0)
	tDamage := float64(damage) * (float64(t) / 100.0)
	eDamage := float64(damage) * (float64(e) / 100.0)

	// влияние типа атаки на урон, за счет защиты корпуса
	kDamage -= kDamage * float64(u.GetProtection(prefix+"kinetics")) / 100
	if kDamage < 0 {
		kDamage = 0
	}

	tDamage -= tDamage * float64(u.GetProtection(prefix+"thermo")) / 100
	if tDamage < 0 {
		tDamage = 0
	}

	eDamage -= eDamage * float64(u.GetProtection(prefix+"explosion")) / 100
	if eDamage < 0 {
		eDamage = 0
	}

	u.ShieldTimeOut = _const.UnitShieldRestoreTime

	damage = int(kDamage + tDamage + eDamage)
	if damage <= 0 {
		chance := game_math.GetRangeRand(0, 5, nil)
		if chance == 0 {
			damage = 1
		} else {
			damage = 0
		}
	}

	// щит
	if u.ShieldHP > 0 && !ignoreShield {
		if damage > u.ShieldHP {
			damage = u.ShieldHP
		}

		u.ShieldHP -= damage
		if u.ShieldHP < 0 {
			u.ShieldHP = 0
		}

		if callback != nil { // TODO только для дебага или улучшить
			callback(u.GetOwnerPlayerID(), startDamage, int(kDamage), int(tDamage), int(eDamage), damage, startSHP, u.ShieldHP, startHP, u.GetHP(), k, t, e)
		}

		return true, u.ShieldHP == 0, damage, 0
	}

	// корпус
	if damage > u.GetHP() {
		damage = u.GetHP()
	}

	u.SetHP(u.GetHP() - damage)
	if u.Immortal && u.GetHP() <= 0 {
		u.SetHP(1)
	}

	if callback != nil { // TODO только для дебага, потом удалить или улучшить
		callback(u.GetOwnerPlayerID(), startDamage, int(kDamage), int(tDamage), int(eDamage), damage, startSHP, u.ShieldHP, startHP, u.GetHP(), k, t, e)
	}

	return false, false, damage, u.AddUrepairableDamage(damage)
}

func (u *Unit) FillLastDamage(d []damage_manager.Damage) {
	u.damageManager.FillLastDamage(d)
}

func (u *Unit) GetArrayLastDamage() []damage_manager.Damage {
	return u.damageManager.GetArrayLastDamage()
}

func (u *Unit) SetLastDamage(playerID, damage int) *damage_manager.Damage {
	if playerID == u.OwnerID {
		return nil
	}

	return u.damageManager.SetLastDamage(playerID, damage)
}

func (u *Unit) GetLastDamage(sec int64) int {
	return u.damageManager.GetLastDamage(sec)
}

func (u *Unit) GetLastDamageByPlayerID(playerID, sec int) int {
	return u.damageManager.GetLastDamageByPlayerID(playerID, sec)
}

func (u *Unit) GetAllDamage(t int64) map[int]int {
	return u.damageManager.GetAllDamage(t)
}
