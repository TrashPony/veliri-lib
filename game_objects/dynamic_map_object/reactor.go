package dynamic_map_object

import _const "github.com/TrashPony/veliri-lib/const"

func (o *Object) SetThorium(count int) {
	o.EnergyCell.SetCount(o.EnergyCell.GetCount() + count)
}

func (o *Object) ReactorWork() {
	// у обьекта упрощеная модель работы реактора, если есть куда втолкнуть то вталкиваем
	if o.EnergyCell != nil && o.EnergyCell.GetCount() > 0 && o.GetPower()+_const.PowerInWork < o.MaxEnergy {
		// идем от обратного что бы не парится о обьявление, когда работа стала 100 то -1 торий
		o.EnergyCell.WorkedOut++

		o.SetPower(o.GetPower() + _const.PowerInWork)

		// если ячейка отработала свой ресурс то удаляем ее
		if o.EnergyCell.WorkedOut >= 100 {
			o.EnergyCell.WorkedOut = 0
			o.EnergyCell.SetCount(o.EnergyCell.GetCount() - 1)
		}
	}
}
