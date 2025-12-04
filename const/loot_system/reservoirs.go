package loot_system

import (
	_const "github.com/TrashPony/veliri-lib/const"
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
	resourceID int
	weight     int // проценты × 100 для целочисленных вычислений (но можно и дроби — я покажу оба)
}

// sectorResourceConfig — полная конфигурация для сектора
type sectorResourceConfig struct {
	maxCount int            // макс ресурсов на сектор (на одну "точку" или на сектор целиком — как у тебя)
	drops    []resourceDrop // пул ресурсов с весами
}

func getResourceDropConfig(sectorType sectorType) *sectorResourceConfig {
	switch sectorType {
	case sectorSafe:
		return &sectorResourceConfig{
			maxCount: 3,
			drops: []resourceDrop{
				{resourceThoriumOre, 32},
				{resourceCopperOre, 32},
				{resourceIronOre, 32},
				{resourceSiliconOre, 2},
				{resourceTitaniumOre, 2},
			},
		}
	case sectorFactionDanger:
		return &sectorResourceConfig{
			maxCount: 5,
			drops: []resourceDrop{
				{resourceThoriumOre, 30},
				{resourceCopperOre, 29},
				{resourceIronOre, 29},
				{resourceSiliconOre, 5},
				{resourceTitaniumOre, 5},
				{resourceOil, 2},
			},
		}
	case sectorPirate:
		return &sectorResourceConfig{
			maxCount: 5,
			drops: []resourceDrop{
				{resourceThoriumOre, 20},
				{resourceCopperOre, 20},
				{resourceIronOre, 20},
				{resourceSiliconOre, 10},
				{resourceTitaniumOre, 10},
				{resourceOil, 20},
			},
		}
	case sectorWasteland:
		return &sectorResourceConfig{
			maxCount: 7,
			drops: []resourceDrop{
				{resourceThoriumOre, 10},
				{resourceCopperOre, 10},
				{resourceIronOre, 10},
				{resourceSiliconOre, 25},
				{resourceTitaniumOre, 25},
				{resourceOil, 20},
			},
		}
	case sectorCombat:
		return &sectorResourceConfig{
			maxCount: 7,
			drops: []resourceDrop{
				{resourceThoriumOre, 15},
				{resourceCopperOre, 20},
				{resourceIronOre, 20},
				{resourceSiliconOre, 10},
				{resourceTitaniumOre, 10},
				{resourceOil, 35},
			},
		}
	default:
		// fallback — безопасный сектор
		return getResourceDropConfig(sectorSafe)
	}
}

// sampleResource — возвращает ID ресурса по весам из конфига
func sampleResource(config *sectorResourceConfig, rng *rand.Rand) int {
	if len(config.drops) == 0 {
		return 0 // или паникуй, или логируй — как удобнее
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
			return d.resourceID
		}
	}

	// На всякий — последний
	return config.drops[len(config.drops)-1].resourceID
}

func getTypeSector(id int, fraction string, freeLand, battle, secure bool) sectorType {
	if freeLand || id < 0 {
		return sectorWasteland
	}

	if battle {
		return sectorCombat
	}

	if fraction == _const.RustbucketCartel {
		return sectorPirate
	}

	if !secure {
		return sectorFactionDanger
	}

	return sectorSafe
}

func GetSectorReservoir(id int, fraction string, freeLand, battle, secure bool, rng *rand.Rand) (int, int) {
	config := getResourceDropConfig(getTypeSector(id, fraction, freeLand, battle, secure))
	return sampleResource(config, rng), config.maxCount
}
