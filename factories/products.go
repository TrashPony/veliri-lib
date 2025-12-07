package factories

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_objects/product"
	"math/rand"
)

const priceK = 10
const sizeKBase = 5
const sizeKContraband = 4
const sizeKHigh = 2
const MaxBaseProducts = 5000.0

var productsTypes = map[int]product.Product{
	4: {
		ID:           4,
		Name:         "industrial_materials",
		Size:         500 / sizeKBase,
		DefaultPrice: 150 * priceK,
	},
	2: {
		ID:           2,
		Name:         "general_purpose_equipment",
		Size:         500 / sizeKBase,
		DefaultPrice: 350 * priceK,
	},
	7: {
		ID:           7,
		Name:         "civilian_weapon",
		Size:         500 / sizeKContraband,
		DefaultPrice: 420 * priceK,
		FairTrade:    []string{_const.Reverses, _const.Explores},
	},
	1: {
		ID:           1,
		Name:         "energy_storage",
		Size:         500 / sizeKBase,
		DefaultPrice: 500 * priceK,
	},
	9: {
		ID:           9,
		Name:         "fox_crystals",
		Size:         500 / sizeKContraband,
		DefaultPrice: 600 * priceK,
		FairTrade:    []string{_const.Explores, _const.Replicas},
	},
	3: {
		ID:           3,
		Name:         "weapon_parts",
		Size:         500 / sizeKHigh,
		DefaultPrice: 1000 * priceK,
	},
	8: {
		ID:           8,
		Name:         "neutrino_processors",
		Size:         500 / sizeKContraband,
		DefaultPrice: 1050 * priceK,
		FairTrade:    []string{_const.Reverses, _const.Replicas},
	},
	6: {
		ID:           6,
		Name:         "technological_maps",
		Size:         500 / sizeKHigh,
		DefaultPrice: 750 * priceK,
	},
	5: {
		ID:           5,
		Name:         "subatomic_composites",
		Size:         500 / sizeKHigh,
		DefaultPrice: 1250 * priceK,
	},
}

func GetProduct(id int) *product.Product {
	newProduct := productsTypes[id]
	return &newProduct
}

func GetProductByName(name string) *product.Product {
	var newProduct product.Product

	for _, p := range productsTypes {
		if p.Name == name {
			newProduct = p
			return &newProduct
		}
	}

	return nil
}

func GetAllProducts() map[int]product.Product {
	return productsTypes
}

func GetRandomProducts() product.Product {
	ps := make([]product.Product, 0)

	for _, p := range productsTypes {
		ps = append(ps, p)
	}

	return ps[rand.Intn(len(ps))]
}
