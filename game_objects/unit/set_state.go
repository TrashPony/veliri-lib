package unit

import (
	"github.com/TrashPony/veliri-lib/game_math"
	"sync/atomic"
)

func (u *Unit) SetPower(power int) {
	u.Power = power
}

func (u *Unit) SetDamage(damage, k, t, e int) int {

	kDamage := float64(damage) * (float64(k) / 100.0)
	tDamage := float64(damage) * (float64(t) / 100.0)
	eDamage := float64(damage) * (float64(e) / 100.0)

	// влияние типа атаки на урон, за счет защиты корпуса
	kDamage -= kDamage * float64(u.GetProtection("kinetics")) / 100
	if kDamage < 0 {
		kDamage = 0
	}

	tDamage -= tDamage * float64(u.GetProtection("thermo")) / 100
	if tDamage < 0 {
		tDamage = 0
	}

	eDamage -= eDamage * float64(u.GetProtection("explosion")) / 100
	if eDamage < 0 {
		eDamage = 0
	}

	damage = int(kDamage + tDamage + eDamage)
	if damage > u.GetHP() {
		damage = u.GetHP()
	}

	if damage <= 0 {
		chance := game_math.GetRangeRand(0, 5, nil)
		if chance == 0 {
			damage = 1
		} else {
			damage = 0
		}
	}

	if !u.Immortal {
		u.SetHP(u.GetHP() - damage)
	}

	return damage
}

func (u *Unit) SetGunRotate(angle float64, slotNumber int) {
	weaponSlot := u.GetWeaponSlot(slotNumber)
	if weaponSlot != nil {
		weaponSlot.SetGunRotate(angle)
	}
}

func (u *Unit) SetStateSenderPos(bool) {
}

func (u *Unit) SetFollowExit(exit bool) {
	u.followExit = exit
}

func (u *Unit) SetEvacuation(evacuation bool) {
	u.evacuation = evacuation
	if u.evacuation {
		u.GetPhysicalModel().BlockControl = true
	} else {
		u.GetPhysicalModel().BlockControl = false
	}
}

func (u *Unit) SetForceEvacuation(evacuation bool) {
	u.forceEvacuation = evacuation
}

func (u *Unit) SetHP(hp int) {
	u.HP = hp
}

func (u *Unit) SetMapID(id int) {
	atomic.StoreInt32(&u.MapID, int32(id))
}

func (u *Unit) SetInSky(inSky bool) {
	u.inSky = inSky
}

func (u *Unit) SetOwnerID(id int) {
	u.OwnerID = id
}

func (u *Unit) SetID(id int) {
	u.ID = id
}
