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
				{resourceThoriumOre, 32, 3},
				{resourceCopperOre, 32, 3},
				{resourceIronOre, 32, 3},
				{resourceSiliconOre, 2, 1},
				{resourceTitaniumOre, 2, 1},
			},
		}
	case sectorFactionDanger:
		return &sectorResourceConfig{
			drops: []resourceDrop{
				{resourceThoriumOre, 30, 3},
				{resourceCopperOre, 29, 3},
				{resourceIronOre, 29, 3},
				{resourceSiliconOre, 5, 1},
				{resourceTitaniumOre, 5, 1},
				{resourceOil, 2, 1},
			},
		}
	case sectorPirate:
		return &sectorResourceConfig{
			drops: []resourceDrop{
				{resourceThoriumOre, 20, 2},
				{resourceCopperOre, 20, 2},
				{resourceIronOre, 20, 2},
				{resourceSiliconOre, 10, 2},
				{resourceTitaniumOre, 10, 2},
				{resourceOil, 20, 3},
			},
		}
	case sectorWasteland:
		return &sectorResourceConfig{
			drops: []resourceDrop{
				{resourceThoriumOre, 10, 1},
				{resourceCopperOre, 10, 1},
				{resourceIronOre, 10, 1},
				{resourceSiliconOre, 25, 3},
				{resourceTitaniumOre, 25, 3},
				{resourceOil, 20, 3},
			},
		}
	case sectorCombat:
		return &sectorResourceConfig{
			drops: []resourceDrop{
				{resourceThoriumOre, 15, 1},
				{resourceCopperOre, 20, 1},
				{resourceIronOre, 20, 1},
				{resourceSiliconOre, 10, 2},
				{resourceTitaniumOre, 10, 2},
				{resourceOil, 35, 3},
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
