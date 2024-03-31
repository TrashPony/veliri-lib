package unit

import "github.com/TrashPony/veliri-lib/game_objects/target"

func (u *Unit) GetWeaponTarget() *target.Target {
	u.mx.RLock()
	defer u.mx.RUnlock()
	return u.weaponTarget
}

func (u *Unit) SetWeaponTarget(target *target.Target) {
	u.mx.Lock()
	defer u.mx.Unlock()
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
