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
	4: { // industrial_materials - сырье, часто торгуется
		ID:                4,
		Name:              "industrial_materials",
		Size:              500 / sizeKBase,
		DefaultPrice:      150 * priceK,
		DiffPercent:       14, // Низкий спред - часто торгуется
		ChangePriceTicSec: 30,
	},
	2: { // general_purpose_equipment - базовый товар
		ID:                2,
		Name:              "general_purpose_equipment",
		Size:              500 / sizeKBase,
		DefaultPrice:      350 * priceK,
		DiffPercent:       12, // Средний спред
		ChangePriceTicSec: 30,
	},
	7: { // civilian_weapon - контрабанда
		ID:                7,
		Name:              "civilian_weapon",
		Size:              500 / sizeKContraband,
		DefaultPrice:      420 * priceK,
		DiffPercent:       14, // Высокий спред - риск контрабанды
		FairTrade:         []string{_const.Reverses, _const.Explores},
		ChangePriceTicSec: 10,
	},
	1: { // energy_storage - базовый, но важный
		ID:                1,
		Name:              "energy_storage",
		Size:              500 / sizeKBase,
		DefaultPrice:      500 * priceK,
		DiffPercent:       12, // Чуть ниже среднего
		ChangePriceTicSec: 30,
	},
	9: { // fox_crystals - контрабанда
		ID:                9,
		Name:              "fox_crystals",
		Size:              500 / sizeKContraband,
		DefaultPrice:      600 * priceK,
		DiffPercent:       14, // Очень высокий - редкая контрабанда
		FairTrade:         []string{_const.Explores, _const.Replicas},
		ChangePriceTicSec: 10,
	},
	3: { // weapon_parts - премиум, военный
		ID:                3,
		Name:              "weapon_parts",
		Size:              500 / sizeKHigh,
		DefaultPrice:      1000 * priceK,
		DiffPercent:       13, // Максимальный - дефицитный военный товар
		ChangePriceTicSec: 50,
	},
	8: { // neutrino_processors - контрабанда + высокотех
		ID:                8,
		Name:              "neutrino_processors",
		Size:              500 / sizeKContraband,
		DefaultPrice:      1050 * priceK,
		DiffPercent:       14, // Высокий спред
		FairTrade:         []string{_const.Reverses, _const.Replicas},
		ChangePriceTicSec: 10,
	},
	6: { // technological_maps - премиум, информация
		ID:                6,
		Name:              "technological_maps",
		Size:              500 / sizeKHigh,
		DefaultPrice:      750 * priceK,
		DiffPercent:       12, // Средне-высокий
		ChangePriceTicSec: 50,
	},
	5: { // subatomic_composites - самый премиум
		ID:                5,
		Name:              "subatomic_composites",
		Size:              500 / sizeKHigh,
		DefaultPrice:      1250 * priceK,
		DiffPercent:       14, // Максимальный - самый дорогой товар
		ChangePriceTicSec: 50,
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
