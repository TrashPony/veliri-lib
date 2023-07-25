package unit

import (
	"github.com/TrashPony/veliri-lib/game_math"
	"sync/atomic"
)

func (unit *Unit) SetPower(power int) {
	unit.Power = power
}

func (unit *Unit) SetDamage(damage, k, t, e int) int {

	kDamage := float64(damage) * (float64(k) / 100.0)
	tDamage := float64(damage) * (float64(t) / 100.0)
	eDamage := float64(damage) * (float64(e) / 100.0)

	// влияние типа атаки на урон, за счет защиты корпуса
	kDamage -= kDamage * float64(unit.GetProtection("kinetics")) / 100
	if kDamage < 0 {
		kDamage = 0
	}

	tDamage -= tDamage * float64(unit.GetProtection("thermo")) / 100
	if tDamage < 0 {
		tDamage = 0
	}

	eDamage -= eDamage * float64(unit.GetProtection("explosion")) / 100
	if eDamage < 0 {
		eDamage = 0
	}

	damage = int(kDamage + tDamage + eDamage)
	if damage > unit.GetHP() {
		damage = unit.GetHP()
	}

	if damage <= 0 {
		chance := game_math.GetRangeRand(0, 5, nil)
		if chance == 0 {
			damage = 1
		} else {
			damage = 0
		}
	}

	unit.SetHP(unit.GetHP() - damage)

	return damage
}

func (unit *Unit) SetGunRotate(angle float64, slotNumber int) {
	weaponSlot := unit.GetWeaponSlot(slotNumber)
	if weaponSlot != nil {
		weaponSlot.SetGunRotate(angle)
	}
}

func (unit *Unit) SetStateSenderPos(bool) {
}

func (unit *Unit) SetFollowExit(exit bool) {
	unit.followExit = exit
}

func (unit *Unit) SetEvacuation(evacuation bool) {
	unit.evacuation = evacuation
	if unit.evacuation {
		unit.GetPhysicalModel().BlockControl = true
	} else {
		unit.GetPhysicalModel().BlockControl = false
	}
}

func (unit *Unit) SetForceEvacuation(evacuation bool) {
	unit.forceEvacuation = evacuation
}

func (unit *Unit) SetHP(hp int) {
	unit.HP = hp
}

func (unit *Unit) SetMapID(id int) {
	atomic.StoreInt32(&unit.MapID, int32(id))
}

func (unit *Unit) SetInSky(inSky bool) {
	unit.inSky = inSky
}

func (unit *Unit) SetOwnerID(id int) {
	unit.OwnerID = id
}

func (unit *Unit) SetID(id int) {
	unit.ID = id
}
