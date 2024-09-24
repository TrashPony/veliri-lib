package descriptions

import _const "github.com/TrashPony/veliri-lib/const"

var FuelDescription = map[string]map[string]DescriptionItem{
	_const.EN: {
		"gas_fuel":       {Name: "Gas fuel cell", Description: "<p>An energy source based on a multicomponent mixture of combustible and non-combustible gases. It is quite cheap, easy to operate, and has low energy intensity.</p>"},
		"solid_fuel":     {Name: "Solid fuel cell", Description: "<p>Energy source based on compressed types of combustibles. Cheap, easy to use, slightly higher energy density.</p>"},
		"liquid_fuel":    {Name: "Hydrocarbon fuel cell", Description: "<p>An energy source based on various substances whose main composition is readily soluble hydrocarbons. Quite common, has an average energy capacity.</p>"},
		"nuclear_fuel":   {Name: "Neutrino fuel cell", Description: "<p>An energy source based on a controlled nuclear reaction. Has high energy intensity.</p>"},
		"synthetic_fuel": {Name: "Synthetic fuel cell", Description: "<p>Energy source based on the processes of synthesis and processing of octane materials. It has an average energy capacity, allows you to autonomously generate low power and increases overall mobility.</p>"},
		"jet_fuel":       {Name: "Jet fuel cell", Description: "<p>A source of energy based on chemical reactions of a heavy mixed fraction. It has low energy intensity, but allows you to generate medium power and increases overall mobility.</p>"},
	},
	_const.RU: {
		"gas_fuel":       {Name: "Ячейка газового элемента", Description: "<p>Источник энергии, основанный на многокомпонентной смеси горючих и негорючих газов. Достаточно дёшев, прост в эксплуатации, обладает малой энергоёмкостью.</p>"},
		"solid_fuel":     {Name: "Ячейка твердотопливного элемента", Description: "<p>Источник энергии, основанный на скомпрессованных видах горючих веществ. Дёшев, прост в эксплуатации, обладает слегка повышенной энергоёмкостью.</p>"},
		"liquid_fuel":    {Name: "Ячейка углеводородного элемента ", Description: "<p>Источник энергии, основанный на различных веществах, чей главный состав - легкорастворимые углеводороды. Достаточно распространён, имеет среднюю энергоёмкость.</p>"},
		"nuclear_fuel":   {Name: "Ячейка нейтринного элемента", Description: "<p>Источник энергии, основанный на контролируемой ядерной реакции. Имеет большую энергоёмкость.</p>"},
		"synthetic_fuel": {Name: "Ячейка синтетического элемента", Description: "<p>Источник энергии, основанный на процессах синтеза и переработки октановых материалов. Имеет среднюю энергоёмкость, позволяет автономно генерировать малые мощности и повышает общую подвижность.</p>"},
		"jet_fuel":       {Name: "Ячейка реактивного элемента", Description: "<p>Источник энергии, основанный на химических реакциях тяжёлой смесевой фракции. Имеет малую энергоёмкость, однако позволяет генерировать средние мощности и повышает общую подвижность.</p>"},
	},
}
