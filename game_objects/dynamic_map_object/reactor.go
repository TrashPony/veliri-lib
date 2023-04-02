package dynamic_map_object

func (o *Object) SetThorium(count int) {
	o.EnergyCell.SetCount(o.EnergyCell.GetCount() + count)
}

func (o *Object) ReactorWork() {
	// у обьекта упрощеная модель работы реактора, если есть куда втолкнуть то вталкиваем

	const powerInWork = 100 // 1 заряд равен 100 энергии // TODO впихнуть это в обьект тория

	if o.EnergyCell != nil && o.EnergyCell.GetCount() > 0 && o.GetPower()+powerInWork < o.MaxEnergy {
		// идем от обратного что бы не парится о обьявление, когда работа стала 100 то -1 торий
		o.EnergyCell.WorkedOut++

		o.SetPower(o.GetPower() + powerInWork)

		// если ячейка отработала свой ресурс то удаляем ее
		if o.EnergyCell.WorkedOut >= 100 {
			o.EnergyCell.WorkedOut = 0
			o.EnergyCell.SetCount(o.EnergyCell.GetCount() - 1)
		}
	}
}
