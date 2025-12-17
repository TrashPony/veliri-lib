package dynamic_map_object

import (
	_const "github.com/TrashPony/veliri-lib/const"
)

func (o *Object) ReactorWork() {
	// TODO эта механика не работает
	// у обьекта упрощеная модель работы реактора, если есть куда втолкнуть то вталкиваем
	if o.EnergyCell != nil && o.GetPower()+_const.PowerInWork < o.MaxEnergy {
		// идем от обратного что бы не парится о обьявление, когда работа стала 100 то -1 торий
		o.EnergyCell.Worked--
		o.SetPower(o.GetPower() + _const.PowerInWork)

		// если ячейка отработала свой ресурс то удаляем ее
		if o.EnergyCell.Worked <= 0 {
			o.EnergyCell.Worked = 0
			if o.EnergyCell.NextFuel.ID != 0 {
				o.EnergyCell.Worked = o.EnergyCell.NextFuel.EnergyCap
			}
		}
	}
}
