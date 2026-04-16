package dynamic_map_object

import (
	_const "github.com/TrashPony/veliri-lib/const"
)

func (o *Object) ReactorWork() {
	if o.EnergyCell != nil && o.GetPower()+_const.PowerInWork < o.MaxEnergy {
		o.SetPower(o.GetPower() + _const.PowerInWork)
	}
}
