package unit

import "github.com/TrashPony/veliri-lib/game_objects/target"

func (u *Unit) GetWeaponTarget(slot int) *target.Target {
	u.mx.RLock()
	defer u.mx.RUnlock()

	if slot != 0 && u.GetWeaponSlot(slot) != nil && u.GetWeaponSlot(slot).Weapon != nil {
		t := u.GetWeaponSlot(slot).GetWeaponTarget()
		if t != nil {
			return t
		}
	}

	return u.weaponTarget
}

func (u *Unit) SetWeaponTarget(slot int, target *target.Target) {
	u.mx.Lock()
	defer u.mx.Unlock()

	if slot != 0 && u.GetWeaponSlot(slot) != nil && u.GetWeaponSlot(slot).Weapon != nil {
		u.GetWeaponSlot(slot).SetWeaponTarget(target)
		return
	}

	if target == nil {
		for _, weaponSlot := range u.getBody().Weapons {
			weaponSlot.SetWeaponTarget(target)
		}
	}

	u.weaponTarget = target
}

func (u *Unit) GetReservoirTarget() *target.Target {
	u.mx.RLock()
	defer u.mx.RUnlock()
	return u.reservoirTarget
}

func (u *Unit) SetReservoirTarget(target *target.Target) {
	u.mx.Lock()
	defer u.mx.Unlock()
	u.reservoirTarget = target
}
