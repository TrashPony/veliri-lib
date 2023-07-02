package _const

type AssortmentPoint struct {
	ID       int    `json:"id"`
	ItemID   int    `json:"item_id"`
	ItemType string `json:"item_type"`
	Points   int    `json:"points"`
	Count    int    `json:"count"`
	MinRank  int    `json:"min_rank"`
}

// TODO баланс
var ReplicAssortment = []AssortmentPoint{
	// ammo bp
	{ID: 1, ItemID: 21, ItemType: "blueprints", Count: 1, Points: 10000, MinRank: 2},
	{ID: 2, ItemID: 22, ItemType: "blueprints", Count: 1, Points: 15000, MinRank: 2},
	{ID: 3, ItemID: 74, ItemType: "blueprints", Count: 1, Points: 5000, MinRank: 2},
	// equip bp
	{ID: 4, ItemID: 13, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 2},
	{ID: 5, ItemID: 19, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 2},
	{ID: 6, ItemID: 20, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 2},
	{ID: 7, ItemID: 69, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 2},
	{ID: 8, ItemID: 103, ItemType: "blueprints", Count: 1, Points: 100000, MinRank: 2},
	// weapon bp
	{ID: 9, ItemID: 78, ItemType: "blueprints", Count: 1, Points: 100000, MinRank: 2},
	{ID: 10, ItemID: 80, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 10},
	{ID: 11, ItemID: 79, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 20},
	// body bp
	{ID: 35, ItemID: 100, ItemType: "blueprints", Count: 1, Points: 250000, MinRank: 10},
	{ID: 12, ItemID: 50, ItemType: "blueprints", Count: 1, Points: 500000, MinRank: 20},
	// ammo // TODO
	// equip // TODO
	// weapon // TODO
	// body // TODO
}

var ExploresAssortment = []AssortmentPoint{
	// ammo bp
	{ID: 13, ItemID: 25, ItemType: "blueprints", Count: 1, Points: 5000, MinRank: 2},
	{ID: 14, ItemID: 26, ItemType: "blueprints", Count: 1, Points: 10000, MinRank: 2},
	// equip bp
	{ID: 15, ItemID: 13, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 2},
	{ID: 16, ItemID: 19, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 2},
	{ID: 17, ItemID: 20, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 2},
	{ID: 18, ItemID: 71, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 2},
	{ID: 19, ItemID: 104, ItemType: "blueprints", Count: 1, Points: 100000, MinRank: 2},
	// weapon bp
	{ID: 20, ItemID: 82, ItemType: "blueprints", Count: 1, Points: 100000, MinRank: 2},
	{ID: 21, ItemID: 83, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 10},
	{ID: 22, ItemID: 84, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 20},
	// body bp
	{ID: 36, ItemID: 101, ItemType: "blueprints", Count: 1, Points: 250000, MinRank: 10},
	{ID: 23, ItemID: 55, ItemType: "blueprints", Count: 1, Points: 500000, MinRank: 20},
	// ammo // TODO
	// equip // TODO
	// weapon // TODO
	// body // TODO
}

var ReversesAssortment = []AssortmentPoint{
	// ammo bp
	{ID: 24, ItemID: 23, ItemType: "blueprints", Count: 1, Points: 10000, MinRank: 2},
	{ID: 25, ItemID: 24, ItemType: "blueprints", Count: 1, Points: 15000, MinRank: 2},
	// equip bp
	{ID: 26, ItemID: 13, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 2},
	{ID: 27, ItemID: 19, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 2},
	{ID: 28, ItemID: 20, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 2},
	{ID: 29, ItemID: 73, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 2},
	{ID: 30, ItemID: 105, ItemType: "blueprints", Count: 1, Points: 100000, MinRank: 2},
	// weapon bp
	{ID: 31, ItemID: 85, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 2},
	{ID: 32, ItemID: 86, ItemType: "blueprints", Count: 1, Points: 100000, MinRank: 10},
	{ID: 33, ItemID: 87, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 20},
	// body bp
	{ID: 37, ItemID: 102, ItemType: "blueprints", Count: 1, Points: 250000, MinRank: 10},
	{ID: 34, ItemID: 60, ItemType: "blueprints", Count: 1, Points: 500000, MinRank: 20},
	// ammo // TODO
	// equip // TODO
	// weapon // TODO
	// body // TODO
}
