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
	_const.ZhCN: {
		"gas_fuel":       {Name: "气体燃料单元", Description: "<p>基于多种可燃和不可燃气体的混合物的能源。价格低廉，操作简单，但能量密度较低。</p>"},
		"solid_fuel":     {Name: "固体燃料单元", Description: "<p>基于压缩可燃物质的能源。价格低廉，操作简单，能量密度略高。</p>"},
		"liquid_fuel":    {Name: "碳氢燃料单元", Description: "<p>基于多种易溶碳氢化合物的能源。较为常见，具有中等能量密度。</p>"},
		"nuclear_fuel":   {Name: "中微子燃料单元", Description: "<p>基于受控核反应的能源。具有高能量密度。</p>"},
		"synthetic_fuel": {Name: "合成燃料单元", Description: "<p>基于合成和加工辛烷材料的能源。具有中等能量密度，能够自主产生小功率并提高整体机动性。</p>"},
		"jet_fuel":       {Name: "喷射燃料单元", Description: "<p>基于重混合物的化学反应的能源。能量密度较低，但能够产生中等功率并提高整体机动性。</p>"},
	},
}
