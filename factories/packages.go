package factories

import (
	_package "github.com/TrashPony/veliri-lib/game_objects/package"
)

const countRecycleItems = 20

var packages = map[int]_package.Package{
	1: {
		ID:        1,
		ItemId:    1,
		Name:      "package_enriched_thorium",
		ItemName:  "enriched_thorium",
		ItemType:  "recycle",
		ItemCount: countRecycleItems,
		Size:      25 * (countRecycleItems / 2),
	},
	2: {
		ID:        2,
		ItemId:    2,
		Name:      "package_copper",
		ItemName:  "copper",
		ItemType:  "recycle",
		ItemCount: countRecycleItems,
		Size:      30 * (countRecycleItems / 2),
	},
	3: {
		ID:        3,
		ItemId:    3,
		Name:      "package_iron",
		ItemName:  "iron",
		ItemType:  "recycle",
		ItemCount: countRecycleItems,
		Size:      35 * (countRecycleItems / 2),
	},
	4: {
		ID:        4,
		ItemId:    4,
		Name:      "package_plastic",
		ItemName:  "plastic",
		ItemType:  "recycle",
		ItemCount: countRecycleItems,
		Size:      50 * (countRecycleItems / 2),
	},
	5: {
		ID:        5,
		ItemId:    5,
		Name:      "package_silicon",
		ItemName:  "silicon",
		ItemType:  "recycle",
		ItemCount: countRecycleItems,
		Size:      40 * (countRecycleItems / 2),
	},
	6: {
		ID:        6,
		ItemId:    6,
		Name:      "package_titanium",
		ItemName:  "titanium",
		ItemType:  "recycle",
		ItemCount: countRecycleItems,
		Size:      50 * (countRecycleItems / 2),
	},
	7: {
		ID:        7,
		ItemId:    7,
		Name:      "package_carbon",
		ItemName:  "carbon",
		ItemType:  "recycle",
		ItemCount: countRecycleItems,
		Size:      25 * (countRecycleItems / 2),
	},
}

func GetPackage(id int) *_package.Package {
	newFuel := packages[id]
	return &newFuel
}

func GetAllPackages() map[int]_package.Package {
	return packages
}
