package unit

import "github.com/TrashPony/veliri-lib/game_objects/target"

func (unit *Unit) GetWeaponTarget() *target.Target {
	unit.mx.RLock()
	defer unit.mx.RUnlock()
	return unit.weaponTarget
}

func (unit *Unit) SetWeaponTarget(target *target.Target) {
	unit.mx.Lock()
	defer unit.mx.Unlock()
	unit.weaponTarget = target
}

func (unit *Unit) GetReservoirTarget() *target.Target {
	unit.mx.RLock()
	defer unit.mx.RUnlock()
	return unit.reservoirTarget
}

func (unit *Unit) SetReservoirTarget(target *target.Target) {
	unit.mx.Lock()
	defer unit.mx.Unlock()
	unit.reservoirTarget = target
}
