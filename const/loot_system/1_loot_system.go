package loot_system

type LootLot struct {
	ItemType string `json:"item_type"`
	ItemID   int    `json:"item_id"`
	Fraction string `json:"fraction"`
}

// Детали для изготовления не требуются другие детали
var t0Details = []LootLot{
	{ItemType: "detail", ItemID: 2},
	{ItemType: "detail", ItemID: 3},
	{ItemType: "detail", ItemID: 5},
	{ItemType: "detail", ItemID: 6},
	{ItemType: "detail", ItemID: 9},
}

// Детали которые делаютяс из другиз деталей
var t1Details = []LootLot{
	{ItemType: "detail", ItemID: 1},
	{ItemType: "detail", ItemID: 4},
	{ItemType: "detail", ItemID: 7},
	{ItemType: "detail", ItemID: 8},
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
	{ItemType: "frr", ItemID: 1},
	{ItemType: "frr", ItemID: 2},
	{ItemType: "frr", ItemID: 3},
	{ItemType: "frr", ItemID: 4},
	{ItemType: "frr", ItemID: 5},
	{ItemType: "frr", ItemID: 6},
	{ItemType: "frr", ItemID: 7},
}

// Товары — по грейдам
var products = map[int][]LootLot{
	1: {
		{ItemType: "product", ItemID: 4},
	},
	2: {
		{ItemType: "product", ItemID: 2},
	},
	3: {
		{ItemType: "product", ItemID: 7},
	},
	4: {
		{ItemType: "product", ItemID: 1},
	},
	5: {
		{ItemType: "product", ItemID: 9},
	},
	6: {
		{ItemType: "product", ItemID: 6},
	},
	7: {
		{ItemType: "product", ItemID: 3},
	},
	8: {
		{ItemType: "product", ItemID: 8},
	},
	9: {
		{ItemType: "product", ItemID: 5},
	},
}
