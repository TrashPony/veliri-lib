package loot_system

type LootLot struct {
	ItemType string `json:"item_type"`
	ItemID   int    `json:"item_id"`
	Fraction string `json:"fraction"`
}

// Детали для изготовления не требуются другие детали
var t0Details = []LootLot{
	{},
}

// Детали которые делаютяс из другиз деталей
var t1Details = []LootLot{
	{},
}

// Чертежи снаряжения, снарядов, деталей, некоторых видов оружи, корпусов и модулей. Подбирается ручками
var t0BluePrints = []LootLot{
	{},
}

// Чертежи снаряжения, снарядов, деталей, некоторых видов оружи, корпусов и модулей. Подбирается ручками
var t1BluePrints = []LootLot{
	{},
}

// Научные данные — типы
var frr = []LootLot{
	{},
}

// Товары — по грейдам
var products = map[int][]LootLot{
	1: {},
}
