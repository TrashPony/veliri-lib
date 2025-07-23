package unit

import (
	"sync/atomic"
)

func (u *Unit) SetPower(power int) {
	u.Power = power
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
