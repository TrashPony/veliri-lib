package descriptions

import _const "github.com/TrashPony/veliri-lib/const"

var FuelDescription = map[string]map[string]DescriptionItem{
	_const.EN: {
		"gas_fuel":       {Name: "Газовое топливо", Description: "малая энергоемкость, продается на базах"},
		"solid_fuel":     {Name: "Твердое топливо", Description: "чуть больше"},
		"liquid_fuel":    {Name: "Жидкое топливо", Description: "средняя"},
		"nuclear_fuel":   {Name: "Ядерное топливо", Description: "большая + малый баф на выработку энергии"},
		"synthetic_fuel": {Name: "Синтетическое топливо", Description: "средняя + средний баф на выработку энергии"},
		"jet_fuel":       {Name: "Реактивное топливо", Description: "малая емкость + большой баф на выработку энергии + баф на скорость (быстро сгорает но излишки заряжуют аккамулятор)"},
	},
	_const.RU: {
		"gas_fuel":       {Name: "Газовое топливо", Description: "малая энергоемкость, продается на базах"},
		"solid_fuel":     {Name: "Твердое топливо", Description: "чуть больше"},
		"liquid_fuel":    {Name: "Жидкое топливо", Description: "средняя"},
		"nuclear_fuel":   {Name: "Ядерное топливо", Description: "большая + малый баф на выработку энергии"},
		"synthetic_fuel": {Name: "Синтетическое топливо", Description: "средняя + средний баф на выработку энергии"},
		"jet_fuel":       {Name: "Реактивное топливо", Description: "малая емкость + большой баф на выработку энергии + баф на скорость (быстро сгорает но излишки заряжуют аккамулятор)"},
	},
}
