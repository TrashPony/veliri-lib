package _const

type AssortmentPoint struct {
	ID       int    `json:"id"`
	ItemID   int    `json:"item_id"`
	ItemType string `json:"item_type"`
	Points   int    `json:"points"`
	Credits  int    `json:"credits"`
	Count    int    `json:"count"`
	MinRank  int    `json:"min_rank"`
	Priority int    `json:"priority"`
}

var ReplicAssortment = []AssortmentPoint{
	// ammo bp +
	{ID: 3, ItemID: 74, ItemType: "blueprints", Count: 5, Points: 1800, MinRank: 1, Priority: 1},
	{ID: 1, ItemID: 21, ItemType: "blueprints", Count: 5, Points: 5000, MinRank: 3, Priority: 1},
	{ID: 2, ItemID: 22, ItemType: "blueprints", Count: 5, Points: 9000, MinRank: 6, Priority: 1},

	// equip bp +
	// rang 1
	{ID: 8, ItemID: 103, ItemType: "blueprints", Count: 1, Points: 100000, MinRank: 1, Priority: 1},
	{ID: 49, ItemID: 63, ItemType: "blueprints", Count: 1, Points: 130000, MinRank: 1, Priority: 2},
	{ID: 50, ItemID: 64, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},

	// rang 3
	{ID: 4, ItemID: 13, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 3, Priority: 1},
	{ID: 51, ItemID: 19, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 3, Priority: 1},
	{ID: 52, ItemID: 20, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 3, Priority: 1},
	{ID: 54, ItemID: 130, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 3, Priority: 2},
	{ID: 55, ItemID: 131, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 3, Priority: 2},
	{ID: 56, ItemID: 133, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 3, Priority: 2},
	{ID: 58, ItemID: 142, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 3, Priority: 2},

	{ID: 60, ItemID: 17, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 3, Priority: 2},
	{ID: 61, ItemID: 148, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 3, Priority: 2},
	{ID: 62, ItemID: 150, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 3, Priority: 2},
	{ID: 63, ItemID: 151, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 3, Priority: 2},
	{ID: 64, ItemID: 152, ItemType: "blueprints", Count: 1, Points: 250000, MinRank: 3, Priority: 2},
	{ID: 65, ItemID: 157, ItemType: "blueprints", Count: 1, Points: 250000, MinRank: 3, Priority: 2},
	{ID: 66, ItemID: 159, ItemType: "blueprints", Count: 1, Points: 500000, MinRank: 3, Priority: 2},
	{ID: 67, ItemID: 160, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 3, Priority: 2},

	// rang 5
	{ID: 70, ItemID: 140, ItemType: "blueprints", Count: 1, Points: 650000, MinRank: 5, Priority: 2},
	{ID: 71, ItemID: 143, ItemType: "blueprints", Count: 1, Points: 400000, MinRank: 5, Priority: 2},
	{ID: 93, ItemID: 69, ItemType: "blueprints", Count: 1, Points: 450000, MinRank: 5, Priority: 2},

	// weapon bp +
	{ID: 10, ItemID: 80, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 1, Priority: 1},
	{ID: 9, ItemID: 78, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 2, Priority: 1},
	{ID: 11, ItemID: 79, ItemType: "blueprints", Count: 1, Points: 500000, MinRank: 5, Priority: 2},
	{ID: 47, ItemID: 77, ItemType: "blueprints", Count: 1, Points: 650000, MinRank: 10, Priority: 3},
	// body bp +
	{ID: 40, ItemID: 119, ItemType: "blueprints", Count: 1, Points: 250000, MinRank: 1, Priority: 1},
	{ID: 41, ItemID: 50, ItemType: "blueprints", Count: 1, Points: 500000, MinRank: 5, Priority: 1},
	{ID: 42, ItemID: 120, ItemType: "blueprints", Count: 1, Points: 850000, MinRank: 10, Priority: 2},
	{ID: 43, ItemID: 52, ItemType: "blueprints", Count: 1, Points: 1150000, MinRank: 15, Priority: 3},

	{ID: 35, ItemID: 100, ItemType: "blueprints", Count: 1, Points: 600000, MinRank: 5, Priority: 1},
	{ID: 44, ItemID: 51, ItemType: "blueprints", Count: 1, Points: 600000, MinRank: 5, Priority: 1},
	{ID: 45, ItemID: 48, ItemType: "blueprints", Count: 1, Points: 1150000, MinRank: 10, Priority: 3},
	{ID: 46, ItemID: 121, ItemType: "blueprints", Count: 1, Points: 1450000, MinRank: 10, Priority: 3},

	// rigs bp
	{ID: 2000, ItemID: 165, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 2001, ItemID: 166, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 2002, ItemID: 170, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 2003, ItemID: 173, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 2004, ItemID: 178, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 2005, ItemID: 181, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 2006, ItemID: 182, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},

	// ammo +
	{ID: 72, ItemID: 13, ItemType: "ammo", Count: 5 * 20, Points: 1800 * 3, MinRank: 1},
	{ID: 73, ItemID: 1, ItemType: "ammo", Count: 5 * 4, Points: 5000 * 3, MinRank: 3},
	{ID: 74, ItemID: 2, ItemType: "ammo", Count: 5 * 2, Points: 9000 * 3, MinRank: 6},
	// equip +
	// weapon +
	{ID: 39, ItemID: 10000011, ItemType: "weapon", Count: 1, Credits: 10000, MinRank: 1},
	// body +
	{ID: 38, ItemID: 10000009, ItemType: "body", Count: 1, Credits: 15000, MinRank: 1},

	{ID: 1000, ItemID: 39, ItemType: "trash", Count: 1, Points: 2000000, Credits: 1000000, MinRank: 1},
}

var ExploresAssortment = []AssortmentPoint{
	// ammo bp +
	{ID: 13, ItemID: 25, ItemType: "blueprints", Count: 5, Points: 1800, MinRank: 1, Priority: 1},
	{ID: 88, ItemID: 161, ItemType: "blueprints", Count: 5, Points: 2500, MinRank: 3, Priority: 1},
	{ID: 14, ItemID: 26, ItemType: "blueprints", Count: 5, Points: 7000, MinRank: 6, Priority: 1},

	// equip bp +
	// rang 1
	{ID: 94, ItemID: 104, ItemType: "blueprints", Count: 1, Points: 100000, MinRank: 1, Priority: 1},
	{ID: 95, ItemID: 64, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 102, ItemID: 129, ItemType: "blueprints", Count: 1, Points: 100000, MinRank: 1, Priority: 2},

	// rang 3
	{ID: 96, ItemID: 13, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 3, Priority: 1},
	{ID: 97, ItemID: 14, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 3, Priority: 2},
	{ID: 98, ItemID: 19, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 3, Priority: 1},
	{ID: 99, ItemID: 20, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 3, Priority: 1},
	{ID: 106, ItemID: 142, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 3, Priority: 2},

	{ID: 107, ItemID: 144, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 3, Priority: 2},
	{ID: 109, ItemID: 145, ItemType: "blueprints", Count: 1, Points: 250000, MinRank: 3, Priority: 2},
	{ID: 110, ItemID: 147, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 3, Priority: 2},
	{ID: 111, ItemID: 149, ItemType: "blueprints", Count: 1, Points: 250000, MinRank: 3, Priority: 2},
	{ID: 112, ItemID: 155, ItemType: "blueprints", Count: 1, Points: 250000, MinRank: 3, Priority: 2},
	{ID: 113, ItemID: 156, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 3, Priority: 2},
	{ID: 114, ItemID: 158, ItemType: "blueprints", Count: 1, Points: 400000, MinRank: 3, Priority: 2},
	{ID: 115, ItemID: 159, ItemType: "blueprints", Count: 1, Points: 500000, MinRank: 3, Priority: 2},
	{ID: 116, ItemID: 160, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 3, Priority: 2},

	// rang 5
	{ID: 100, ItemID: 71, ItemType: "blueprints", Count: 1, Points: 450000, MinRank: 5, Priority: 2},
	{ID: 101, ItemID: 128, ItemType: "blueprints", Count: 1, Points: 500000, MinRank: 5, Priority: 2},
	{ID: 102, ItemID: 132, ItemType: "blueprints", Count: 1, Points: 400000, MinRank: 5, Priority: 2},
	{ID: 103, ItemID: 135, ItemType: "blueprints", Count: 1, Points: 400000, MinRank: 5, Priority: 2},
	{ID: 104, ItemID: 137, ItemType: "blueprints", Count: 1, Points: 500000, MinRank: 5, Priority: 2},
	{ID: 105, ItemID: 141, ItemType: "blueprints", Count: 1, Points: 800000, MinRank: 5, Priority: 2},

	// weapon bp +
	{ID: 79, ItemID: 81, ItemType: "blueprints", Count: 1, Points: 650000, MinRank: 10, Priority: 3},
	{ID: 20, ItemID: 82, ItemType: "blueprints", Count: 1, Points: 500000, MinRank: 5, Priority: 2},
	{ID: 21, ItemID: 83, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 1, Priority: 1},
	{ID: 22, ItemID: 84, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 2, Priority: 1},

	// body bp +
	{ID: 80, ItemID: 110, ItemType: "blueprints", Count: 1, Points: 250000, MinRank: 1, Priority: 1},
	{ID: 81, ItemID: 55, ItemType: "blueprints", Count: 1, Points: 500000, MinRank: 5, Priority: 1},
	{ID: 82, ItemID: 122, ItemType: "blueprints", Count: 1, Points: 850000, MinRank: 10, Priority: 2},
	{ID: 83, ItemID: 57, ItemType: "blueprints", Count: 1, Points: 1150000, MinRank: 15, Priority: 3},

	{ID: 84, ItemID: 101, ItemType: "blueprints", Count: 1, Points: 600000, MinRank: 5, Priority: 1},
	{ID: 85, ItemID: 56, ItemType: "blueprints", Count: 1, Points: 600000, MinRank: 5, Priority: 1},
	{ID: 86, ItemID: 53, ItemType: "blueprints", Count: 1, Points: 1150000, MinRank: 10, Priority: 3},
	{ID: 87, ItemID: 123, ItemType: "blueprints", Count: 1, Points: 1450000, MinRank: 10, Priority: 3},

	// rigs bp
	{ID: 2007, ItemID: 167, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 2008, ItemID: 169, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 2009, ItemID: 172, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 2010, ItemID: 174, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 2011, ItemID: 177, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 2012, ItemID: 179, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 2013, ItemID: 183, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},

	// ammo +
	{ID: 89, ItemID: 5, ItemType: "ammo", Count: 5 * 8, Points: 1800 * 3, MinRank: 1},
	{ID: 90, ItemID: 18, ItemType: "ammo", Count: 5 * 100, Points: 2500 * 3, MinRank: 3},
	{ID: 91, ItemID: 6, ItemType: "ammo", Count: 5 * 4, Points: 7000 * 3, MinRank: 6},

	// equip +
	// weapon +
	{ID: 75, ItemID: 10000016, ItemType: "weapon", Count: 1, Credits: 10000, MinRank: 1},
	// body +
	{ID: 76, ItemID: 10000014, ItemType: "body", Count: 1, Credits: 15000, MinRank: 1},

	{ID: 1001, ItemID: 39, ItemType: "trash", Count: 1, Points: 2000000, Credits: 1000000, MinRank: 1},
}

var ReversesAssortment = []AssortmentPoint{
	// ammo bp +
	{ID: 130, ItemID: 163, ItemType: "blueprints", Count: 5, Points: 2500, MinRank: 1, Priority: 1},
	{ID: 131, ItemID: 23, ItemType: "blueprints", Count: 5, Points: 8000, MinRank: 3, Priority: 1},
	{ID: 132, ItemID: 75, ItemType: "blueprints", Count: 5, Points: 14500, MinRank: 3, Priority: 1},
	{ID: 133, ItemID: 24, ItemType: "blueprints", Count: 5, Points: 20000, MinRank: 6, Priority: 1},
	{ID: 134, ItemID: 76, ItemType: "blueprints", Count: 5, Points: 30000, MinRank: 6, Priority: 1},

	// equip bp
	// rang 1
	{ID: 140, ItemID: 105, ItemType: "blueprints", Count: 1, Points: 100000, MinRank: 1, Priority: 1},
	{ID: 141, ItemID: 129, ItemType: "blueprints", Count: 1, Points: 100000, MinRank: 1, Priority: 2},
	// rang 3
	{ID: 142, ItemID: 13, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 3, Priority: 1},
	{ID: 143, ItemID: 65, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 3, Priority: 1},
	{ID: 144, ItemID: 19, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 3, Priority: 1},
	{ID: 145, ItemID: 20, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 3, Priority: 1},
	{ID: 146, ItemID: 127, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 3, Priority: 2},
	{ID: 147, ItemID: 134, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 3, Priority: 2},
	{ID: 148, ItemID: 142, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 3, Priority: 2},

	{ID: 149, ItemID: 148, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 3, Priority: 2},
	{ID: 150, ItemID: 150, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 3, Priority: 2},
	{ID: 151, ItemID: 146, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 3, Priority: 2},
	{ID: 152, ItemID: 153, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 3, Priority: 2},
	{ID: 153, ItemID: 154, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 3, Priority: 2},
	{ID: 154, ItemID: 158, ItemType: "blueprints", Count: 1, Points: 400000, MinRank: 3, Priority: 2},
	{ID: 155, ItemID: 159, ItemType: "blueprints", Count: 1, Points: 500000, MinRank: 3, Priority: 2},
	{ID: 156, ItemID: 160, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 3, Priority: 2},
	// rang 5
	{ID: 157, ItemID: 73, ItemType: "blueprints", Count: 1, Points: 450000, MinRank: 5, Priority: 2},
	{ID: 158, ItemID: 136, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 5, Priority: 2},
	{ID: 159, ItemID: 138, ItemType: "blueprints", Count: 1, Points: 700000, MinRank: 5, Priority: 2},
	{ID: 160, ItemID: 139, ItemType: "blueprints", Count: 1, Points: 600000, MinRank: 5, Priority: 2},
	{ID: 161, ItemID: 143, ItemType: "blueprints", Count: 1, Points: 400000, MinRank: 5, Priority: 2},

	// weapon bp +
	{ID: 125, ItemID: 162, ItemType: "blueprints", Count: 1, Points: 200000, MinRank: 1, Priority: 1},
	{ID: 126, ItemID: 86, ItemType: "blueprints", Count: 1, Points: 300000, MinRank: 2, Priority: 1},
	{ID: 127, ItemID: 87, ItemType: "blueprints", Count: 1, Points: 500000, MinRank: 5, Priority: 2},
	{ID: 128, ItemID: 88, ItemType: "blueprints", Count: 1, Points: 650000, MinRank: 10, Priority: 3},
	{ID: 129, ItemID: 85, ItemType: "blueprints", Count: 1, Points: 650000, MinRank: 10, Priority: 3},

	// body bp +
	{ID: 117, ItemID: 124, ItemType: "blueprints", Count: 1, Points: 250000, MinRank: 1, Priority: 1},
	{ID: 118, ItemID: 60, ItemType: "blueprints", Count: 1, Points: 500000, MinRank: 5, Priority: 1},
	{ID: 119, ItemID: 125, ItemType: "blueprints", Count: 1, Points: 850000, MinRank: 10, Priority: 2},
	{ID: 120, ItemID: 62, ItemType: "blueprints", Count: 1, Points: 1150000, MinRank: 15, Priority: 3},

	{ID: 121, ItemID: 102, ItemType: "blueprints", Count: 1, Points: 600000, MinRank: 5, Priority: 1},
	{ID: 122, ItemID: 61, ItemType: "blueprints", Count: 1, Points: 600000, MinRank: 5, Priority: 1},
	{ID: 123, ItemID: 58, ItemType: "blueprints", Count: 1, Points: 1150000, MinRank: 10, Priority: 3},
	{ID: 124, ItemID: 126, ItemType: "blueprints", Count: 1, Points: 1450000, MinRank: 10, Priority: 3},

	// rigs bp
	{ID: 2014, ItemID: 168, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 2015, ItemID: 171, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 2016, ItemID: 175, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 2017, ItemID: 176, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 2018, ItemID: 180, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},
	{ID: 2018, ItemID: 184, ItemType: "blueprints", Count: 1, Points: 150000, MinRank: 1, Priority: 2},

	// ammo +
	{ID: 135, ItemID: 33, ItemType: "ammo", Count: 5 * 10, Points: 2500 * 3, MinRank: 1},
	{ID: 136, ItemID: 3, ItemType: "ammo", Count: 5 * 8, Points: 8000 * 3, MinRank: 3},
	{ID: 137, ItemID: 14, ItemType: "ammo", Count: 5 * 4, Points: 14500 * 3, MinRank: 3},
	{ID: 138, ItemID: 4, ItemType: "ammo", Count: 5 * 2, Points: 20000 * 3, MinRank: 6},
	{ID: 139, ItemID: 15, ItemType: "ammo", Count: 5 * 2, Points: 30000 * 3, MinRank: 6},

	// equip +
	// weapon +
	{ID: 77, ItemID: 10000042, ItemType: "weapon", Count: 1, Credits: 10000, MinRank: 1},
	// body +
	{ID: 78, ItemID: 10000019, ItemType: "body", Count: 1, Credits: 15000, MinRank: 1},

	{ID: 1002, ItemID: 39, ItemType: "trash", Count: 1, Points: 2000000, Credits: 1000000, MinRank: 1},
}
