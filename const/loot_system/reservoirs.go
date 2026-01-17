package loot_system

import (
	"math/rand"
)

const (
	resourceCopperOre   = 1
	resourceIronOre     = 2
	resourceSiliconOre  = 3
	resourceThoriumOre  = 4
	resourceTitaniumOre = 5
	resourceOil         = 6
)

type sectorType int

const (
	sectorSafe          sectorType = iota // 0
	sectorFactionDanger                   // 1
	sectorPirate                          // 2
	sectorWasteland                       // 3
	sectorCombat                          // 4
)

// resourceDrop — один элемент в пуле (ресурс + вес)
type resourceDrop struct {
	maxCount   int // макс ресурсов на сектор (на одну "точку" или на сектор целиком — как у тебя)
	resourceID int
	weight     int // проценты × 100 для целочисленных вычислений (но можно и дроби — я покажу оба)
}

// sectorResourceConfig — полная конфигурация для сектора
type sectorResourceConfig struct {
	drops []resourceDrop // пул ресурсов с весами
}

func getResourceDropConfig(sectorType sectorType) *sectorResourceConfig {
	switch sectorType {
	case sectorSafe:
		return &sectorResourceConfig{
			drops: []resourceDrop{
				{resourceID: resourceThoriumOre, weight: 32, maxCount: 3},
				{resourceID: resourceCopperOre, weight: 32, maxCount: 3},
				{resourceID: resourceIronOre, weight: 32, maxCount: 3},
				{resourceID: resourceSiliconOre, weight: 2, maxCount: 1},
				{resourceID: resourceTitaniumOre, weight: 2, maxCount: 1},
			},
		}
	case sectorFactionDanger:
		return &sectorResourceConfig{
			drops: []resourceDrop{
				{resourceID: resourceThoriumOre, weight: 30, maxCount: 3},
				{resourceID: resourceCopperOre, weight: 29, maxCount: 3},
				{resourceID: resourceIronOre, weight: 29, maxCount: 3},
				{resourceID: resourceSiliconOre, weight: 5, maxCount: 1},
				{resourceID: resourceTitaniumOre, weight: 5, maxCount: 1},
				{resourceID: resourceOil, weight: 2, maxCount: 1},
			},
		}
	case sectorPirate:
		return &sectorResourceConfig{
			drops: []resourceDrop{
				{resourceID: resourceThoriumOre, weight: 20, maxCount: 2},
				{resourceID: resourceCopperOre, weight: 20, maxCount: 2},
				{resourceID: resourceIronOre, weight: 20, maxCount: 2},
				{resourceID: resourceSiliconOre, weight: 10, maxCount: 2},
				{resourceID: resourceTitaniumOre, weight: 10, maxCount: 2},
				{resourceID: resourceOil, weight: 20, maxCount: 3},
			},
		}
	case sectorWasteland:
		return &sectorResourceConfig{
			drops: []resourceDrop{
				{resourceID: resourceThoriumOre, weight: 10, maxCount: 1},
				{resourceID: resourceCopperOre, weight: 10, maxCount: 1},
				{resourceID: resourceIronOre, weight: 10, maxCount: 1},
				{resourceID: resourceSiliconOre, weight: 25, maxCount: 3},
				{resourceID: resourceTitaniumOre, weight: 25, maxCount: 3},
				{resourceID: resourceOil, weight: 20, maxCount: 3},
			},
		}
	case sectorCombat:
		return &sectorResourceConfig{
			drops: []resourceDrop{
				{resourceID: resourceThoriumOre, weight: 15, maxCount: 1},
				{resourceID: resourceCopperOre, weight: 20, maxCount: 1},
				{resourceID: resourceIronOre, weight: 20, maxCount: 1},
				{resourceID: resourceSiliconOre, weight: 10, maxCount: 2},
				{resourceID: resourceTitaniumOre, weight: 10, maxCount: 2},
				{resourceID: resourceOil, weight: 35, maxCount: 3},
			},
		}
	default:
		// fallback — безопасный сектор
		return getResourceDropConfig(sectorSafe)
	}
}

// sampleResource — возвращает ID ресурса по весам из конфига
func sampleResource(config *sectorResourceConfig, rng *rand.Rand) (int, int) {
	if len(config.drops) == 0 {
		return 0, 0
	}

	// Считаем общий вес
	totalWeight := 0
	for _, d := range config.drops {
		totalWeight += d.weight
	}

	// Рандом от 1 до totalWeight (включительно)
	r := rng.Intn(totalWeight) + 1

	// Идём по дропам — кто "поглотил" r — тот и выпал
	accum := 0
	for _, d := range config.drops {
		accum += d.weight
		if r <= accum {
			return d.resourceID, d.maxCount
		}
	}

	// На всякий — последний
	res := config.drops[len(config.drops)-1]
	return res.resourceID, res.maxCount
}

func GetSectorReservoir(id int, fraction string, freeLand, battle, secure bool, rng *rand.Rand) (int, int) {
	config := getResourceDropConfig(getTypeSector(id, fraction, freeLand, battle, secure))
	return sampleResource(config, rng)
}
