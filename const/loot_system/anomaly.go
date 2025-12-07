package loot_system

import (
	"github.com/TrashPony/veliri-lib/game_math"
	"math/rand"
)

// LootDrop — то, что реально выпадает
type LootDrop struct {
	LootLot
	Count int // итоговое количество (уже с учётом BaseCount и системы)
}

// anomalyDrop — внутренняя структура для конфига
type anomalyDrop struct {
	Type   string // "trap", "reservoir_group", "treasure"
	Weight int
	Params map[string]interface{} // гибко: reservoirs=[4,2], minGrade=1, maxGrade=3 и т.д.
}

// TreasureConfig — параметры содержимого "сокровища" для конкретного типа сектора
type TreasureConfig struct {
	PartTier         int
	PartMin, PartMax int

	GoodsMinGrade, GoodsMaxGrade int
	GoodsMinCount, GoodsMaxCount int

	BlueprintTier int

	ScienceMin, ScienceMax int

	// Веса (в процентах) — можно менять без перекомпиляции!
	WeightParts     int // 40
	WeightGoods     int // 40
	WeightBlueprint int // 10
	WeightScience   int // 10
}

// anomalyConfig — конфиг для типа сектора
type anomalyConfig struct {
	MaxPoints int
	Drops     []anomalyDrop
	Treasure  *TreasureConfig
}

func getAnomalyConfig(sectorType sectorType) *anomalyConfig {
	switch sectorType {
	case sectorSafe:
		return &anomalyConfig{
			MaxPoints: 30,
			Drops: []anomalyDrop{
				{"trap", 20, map[string]interface{}{"type": "bomb"}},
				{"trap", 0, map[string]interface{}{"type": "respawn"}},
				{"reservoir_group", 5, map[string]interface{}{"resources": []int{resourceThoriumOre, resourceCopperOre}}},
				{"treasure", 75, nil},
			},
			Treasure: &TreasureConfig{
				PartTier:        0,
				PartMin:         50,
				PartMax:         100,
				GoodsMinGrade:   1,
				GoodsMaxGrade:   3,
				GoodsMinCount:   10,
				GoodsMaxCount:   50,
				BlueprintTier:   0,
				ScienceMin:      1,
				ScienceMax:      2,
				WeightParts:     40,
				WeightGoods:     40,
				WeightBlueprint: 10,
				WeightScience:   10,
			},
		}
	case sectorFactionDanger:
		return &anomalyConfig{
			MaxPoints: 25,
			Drops: []anomalyDrop{
				{"trap", 20, map[string]interface{}{"type": "bomb"}},
				{"trap", 10, map[string]interface{}{"type": "respawn"}},
				{"reservoir_group", 5, map[string]interface{}{"resources": []int{resourceIronOre, resourceSiliconOre}}},
				{"treasure", 65, nil},
			},
			Treasure: &TreasureConfig{
				PartTier:        0,
				PartMin:         150,
				PartMax:         200,
				GoodsMinGrade:   2,
				GoodsMaxGrade:   5,
				GoodsMinCount:   25,
				GoodsMaxCount:   50,
				BlueprintTier:   0,
				ScienceMin:      2,
				ScienceMax:      3,
				WeightParts:     30,
				WeightGoods:     30,
				WeightBlueprint: 20,
				WeightScience:   20,
			},
		}
	case sectorPirate:
		return &anomalyConfig{
			MaxPoints: 20,
			Drops: []anomalyDrop{
				{"trap", 25, map[string]interface{}{"type": "bomb"}},
				{"trap", 10, map[string]interface{}{"type": "respawn"}},
				{"reservoir_group", 5, map[string]interface{}{"resources": []int{resourceTitaniumOre, resourceSiliconOre}}},
				{"treasure", 60, nil},
			},
			Treasure: &TreasureConfig{
				PartTier:        1,
				PartMin:         50,
				PartMax:         100,
				GoodsMinGrade:   4,
				GoodsMaxGrade:   5,
				GoodsMinCount:   25,
				GoodsMaxCount:   50,
				BlueprintTier:   1,
				ScienceMin:      3,
				ScienceMax:      4,
				WeightParts:     30,
				WeightGoods:     30,
				WeightBlueprint: 20,
				WeightScience:   20,
			},
		}
	case sectorWasteland: // одинаковые по луту
		return &anomalyConfig{
			MaxPoints: 15,
			Drops: []anomalyDrop{
				{"trap", 25, map[string]interface{}{"type": "bomb"}},
				{"trap", 10, map[string]interface{}{"type": "respawn"}},
				{"reservoir_group", 5, map[string]interface{}{"resources": []int{resourceTitaniumOre, resourceSiliconOre}}},
				{"treasure", 60, nil},
			},
			Treasure: &TreasureConfig{
				PartTier:        1,
				PartMin:         100,
				PartMax:         200,
				GoodsMinGrade:   5,
				GoodsMaxGrade:   9,
				GoodsMinCount:   25,
				GoodsMaxCount:   100,
				BlueprintTier:   1,
				ScienceMin:      3,
				ScienceMax:      4,
				WeightParts:     30,
				WeightGoods:     30,
				WeightBlueprint: 20,
				WeightScience:   20,
			},
		}
	case sectorCombat:
		return &anomalyConfig{
			MaxPoints: 15,
			Drops: []anomalyDrop{
				{"trap", 25, map[string]interface{}{"type": "bomb"}},
				{"trap", 10, map[string]interface{}{"type": "respawn"}},
				{"reservoir_group", 5, map[string]interface{}{"resources": []int{resourceOil, resourceTitaniumOre, resourceSiliconOre}}},
				{"treasure", 60, nil},
			},
			Treasure: &TreasureConfig{
				PartTier:        1,
				PartMin:         100,
				PartMax:         200,
				GoodsMinGrade:   5,
				GoodsMaxGrade:   9,
				GoodsMinCount:   25,
				GoodsMaxCount:   100,
				BlueprintTier:   1,
				ScienceMin:      3,
				ScienceMax:      4,
				WeightParts:     30,
				WeightGoods:     30,
				WeightBlueprint: 20,
				WeightScience:   20,
			},
		}
	default:
		return getAnomalyConfig(sectorSafe)
	}
}

func GetAnomalyMaxPoints(id int, fraction string, freeLand, battle, secure bool) int {
	cfg := getAnomalyConfig(getTypeSector(id, fraction, freeLand, battle, secure))
	return cfg.MaxPoints
}

func GenerateAnomaly(id int, fraction string, freeLand, battle, secure bool, rng *rand.Rand) (LootDrop, int) {
	cfg := getAnomalyConfig(getTypeSector(id, fraction, freeLand, battle, secure))

	// Шаг 1: выбрать тип аномалии (ловушка / ресурс / сокровище)
	mainDrop := weightedPick(cfg.Drops, rng)
	if mainDrop == nil {
		return LootDrop{}, cfg.MaxPoints
	}

	switch mainDrop.Type {
	case "trap":
		trapType := mainDrop.Params["type"].(string)
		var trapID int
		switch trapType {
		case "bomb":
			trapID = 1 // или твой ID бомбы
		case "respawn":
			trapID = 2 // респавн-ловушка
		}

		return LootDrop{
			LootLot: LootLot{ItemType: "trap", ItemID: trapID},
		}, cfg.MaxPoints

	case "reservoir_group":
		resources := mainDrop.Params["resources"].([]int)
		if len(resources) == 0 {
			return LootDrop{}, cfg.MaxPoints
		}

		return LootDrop{
			LootLot: LootLot{ItemType: "reservoir"},
		}, cfg.MaxPoints

	case "treasure":
		return LootDrop{
			LootLot: LootLot{ItemType: "treasure"},
		}, cfg.MaxPoints
	}

	return LootDrop{}, cfg.MaxPoints
}

// weightedPick — универсальный picker для []anomalyDrop
func weightedPick(drops []anomalyDrop, rng *rand.Rand) *anomalyDrop {
	if len(drops) == 0 {
		return nil
	}
	total := 0
	for _, d := range drops {
		total += d.Weight
	}
	r := rng.Intn(total)
	accum := 0
	for i, d := range drops {
		accum += d.Weight
		if r < accum {
			return &drops[i]
		}
	}
	return &drops[len(drops)-1]
}

func GenerateTreasureContent(id int, fraction string, freeLand, battle, secure bool, rng *rand.Rand) []LootDrop {
	sectorCfg := getAnomalyConfig(getTypeSector(id, fraction, freeLand, battle, secure))
	if sectorCfg == nil {
		return nil
	}

	cfg := sectorCfg.Treasure
	if cfg == nil {
		return nil
	}

	// Взвешенный выбор по явным весам
	totalWeight := cfg.WeightParts + cfg.WeightGoods + cfg.WeightBlueprint + cfg.WeightScience
	r := rng.Intn(totalWeight)

	switch {
	case r < cfg.WeightParts:
		return generateParts(cfg.PartTier, cfg.PartMin, cfg.PartMax, rng)

	case r < cfg.WeightParts+cfg.WeightGoods:
		return generateGoods(cfg.GoodsMinGrade, cfg.GoodsMaxGrade, cfg.GoodsMinCount, cfg.GoodsMaxCount, rng)

	case r < cfg.WeightParts+cfg.WeightGoods+cfg.WeightBlueprint:
		return generateBlueprint(cfg.BlueprintTier, fraction, rng)

	default: // science
		return generateScience(cfg.ScienceMin, cfg.ScienceMax, rng)
	}
}

// Вспомогательные функции — чистые, переиспользуемые
func generateParts(tier, minCount, maxCount int, rng *rand.Rand) []LootDrop {
	pool := t0Details
	if tier == 1 {
		pool = t1Details
	}

	if len(pool) == 0 {
		return nil
	}

	lot := pool[rng.Intn(len(pool))]

	count := game_math.GetRangeRand(minCount, maxCount, rng) * (game_math.GetRangeRand(1, lot.BaseCount+1, rng))
	return []LootDrop{{LootLot: lot, Count: count}}
}

func generateGoods(minGrade, maxGrade, minCount, maxCount int, rng *rand.Rand) []LootDrop {
	grade := minGrade + rng.Intn(maxGrade-minGrade+1)
	pool, ok := products[grade]
	if !ok || len(pool) == 0 {
		// fallback на первый непустой
		for g := minGrade; g <= maxGrade; g++ {
			if p, exists := products[g]; exists && len(p) > 0 {
				pool = p
				break
			}
		}
	}

	if len(pool) == 0 {
		return nil
	}

	lot := pool[rng.Intn(len(pool))]

	count := game_math.GetRangeRand(minCount, maxCount, rng) * (game_math.GetRangeRand(1, lot.BaseCount+1, rng))
	return []LootDrop{{LootLot: lot, Count: count}}
}

func generateBlueprint(tier int, fraction string, rng *rand.Rand) []LootDrop {
	pool := filterByFraction(t0BluePrints, fraction)
	if tier == 1 {
		pool = filterByFraction(t1BluePrints, fraction)
	}

	if len(pool) == 0 {
		// fallback: любая фракция
		if tier == 0 {
			pool = t0BluePrints
		} else {
			pool = t1BluePrints
		}
	}

	if len(pool) == 0 {
		return nil
	}

	lot := pool[rng.Intn(len(pool))]
	return []LootDrop{{LootLot: lot, Count: game_math.GetRangeRand(1, lot.BaseCount+1, rng)}}
}

func generateScience(minCount, maxCount int, rng *rand.Rand) []LootDrop {
	if len(frr) == 0 {
		return nil
	}

	lot := frr[rng.Intn(len(frr))]

	count := game_math.GetRangeRand(minCount, maxCount, rng) * (game_math.GetRangeRand(1, lot.BaseCount+1, rng))
	return []LootDrop{{LootLot: lot, Count: count}}
}
