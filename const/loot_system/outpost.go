package loot_system

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_math"
	"math/rand"
)

type OutpostType int

const (
	OutpostStructure    OutpostType = iota // 0 — турели/щиты/радары
	OutpostCore                            // 1 — центральное строение
	OutpostStructureNPC                    // 2 — NPC структуры
	OutpostCoreNPC                         // 3 — NPC центр
)

type OutpostLootConfig struct {
	PartMin       int  `json:"part_min"`
	PartMax       int  `json:"part_max"` // детали Т1
	BlueprintMin  int  `json:"blueprint_min"`
	BlueprintMax  int  `json:"blueprint_max"` // чертежи Т1 (шт)
	CurrencyMin   int  `json:"currency_min"`
	CurrencyMax   int  `json:"currency_max"` // валюта
	ScienceMin    int  `json:"science_min"`
	ScienceMax    int  `json:"science_max"` // научные данные
	GoodsMinGrade int  `json:"goods_min_grade"`
	GoodsMaxGrade int  `json:"goods_max_grade"`
	GoodsMinCount int  `json:"goods_min_count"`
	GoodsMaxCount int  `json:"goods_max_count"`
	FractionAny   bool `json:"fraction_any"` // true = любые чертежи, false = только фракционные
}

func getOutpostLootConfig(outpostType OutpostType) *OutpostLootConfig {
	switch outpostType {
	case OutpostStructure:
		return &OutpostLootConfig{
			PartMin: 25, PartMax: 100,
			CurrencyMin: 5000, CurrencyMax: 10000,
		}
	case OutpostCore:
		return &OutpostLootConfig{
			PartMin: 100, PartMax: 200,
			BlueprintMin: 5, BlueprintMax: 10,
			CurrencyMin: 20000, CurrencyMax: 100000,
			ScienceMin: 50, ScienceMax: 100,
			GoodsMinGrade: 5, GoodsMaxGrade: 9,
			GoodsMinCount: 100, GoodsMaxCount: 200,
		}
	case OutpostStructureNPC:
		return &OutpostLootConfig{
			PartMin: 50, PartMax: 100,
			CurrencyMin: 5000, CurrencyMax: 10000,
		}
	case OutpostCoreNPC:
		return &OutpostLootConfig{
			PartMin: 150, PartMax: 300,
			BlueprintMin: 5, BlueprintMax: 10,
			CurrencyMin: 50000, CurrencyMax: 100000,
			ScienceMin: 75, ScienceMax: 150,
			FractionAny: true, // "любые t1"
		}
	default:
		return &OutpostLootConfig{}
	}
}

func GenerateOutpostLoot(structureType, fraction string, core bool, rng *rand.Rand) []LootDrop {

	var drops []LootDrop
	outpostType := getTypeOutpostStructure(structureType, fraction, core)
	if outpostType == -1 {
		return drops
	}

	cfg := getOutpostLootConfig(outpostType)

	// Детали Т1 (всегда)
	count := game_math.GetRangeRand(cfg.PartMin, cfg.PartMax, rng)
	if count > 0 && len(T1Details) > 0 {
		lot := T1Details[rng.Intn(len(T1Details))]
		total := count * game_math.GetRangeRand(1, lot.BaseCount+1, rng)
		drops = append(drops, LootDrop{LootLot: lot, Count: total})
	}

	// Чертежи Т1 (если есть)
	if cfg.BlueprintMin > 0 {
		blueprintCount := game_math.GetRangeRand(cfg.BlueprintMin, cfg.BlueprintMax, rng)
		for i := 0; i < blueprintCount; i++ {
			pool := filterByFraction(T1BluePrints, fraction)
			if cfg.FractionAny {
				pool = T1BluePrints // любые
			}
			if len(pool) == 0 {
				pool = T1BluePrints // fallback
			}
			if len(pool) > 0 {
				lot := pool[rng.Intn(len(pool))]
				drops = append(drops, LootDrop{
					LootLot: lot,
					Count:   game_math.GetRangeRand(1, lot.BaseCount+1, rng),
				})
			}
		}
	}

	// Валюта
	if cfg.CurrencyMin > 0 {
		amount := game_math.GetRangeRand(cfg.CurrencyMin, cfg.CurrencyMax, rng)
		drops = append(drops, LootDrop{
			LootLot: LootLot{ItemType: "trash", ItemID: 41}, // 1k_credits
			Count:   amount / 1000,
		})
	}

	// Научные данные
	if cfg.ScienceMin > 0 && len(Frr) > 0 {
		count := game_math.GetRangeRand(cfg.ScienceMin, cfg.ScienceMax, rng)
		lot := Frr[rng.Intn(len(Frr))]
		drops = append(drops, LootDrop{LootLot: lot, Count: count})
	}

	// Товары (только у OutpostCore)
	if cfg.GoodsMinGrade > 0 {
		grade := game_math.GetRangeRand(cfg.GoodsMinGrade, cfg.GoodsMaxGrade, rng)
		pool, ok := Products[grade]
		if !ok || len(pool) == 0 {
			for g := cfg.GoodsMinGrade; g <= cfg.GoodsMaxGrade; g++ {
				if p, exists := Products[g]; exists && len(p) > 0 {
					pool = p
					break
				}
			}
		}
		if len(pool) > 0 {
			count := game_math.GetRangeRand(cfg.GoodsMinCount, cfg.GoodsMaxCount, rng)
			lot := pool[rng.Intn(len(pool))]
			total := count * game_math.GetRangeRand(1, lot.BaseCount+1, rng)
			drops = append(drops, LootDrop{LootLot: lot, Count: total})
		}
	}

	return drops
}

func getTypeOutpostStructure(structureType, fraction string, core bool) OutpostType {
	if structureType == "storage" ||
		structureType == "shield" ||
		structureType == "radar" ||
		structureType == "turret" ||
		structureType == "jammer" ||
		structureType == "repair_station" ||
		structureType == "extractor" ||
		structureType == "generator" {

		if fraction == _const.Replicas || fraction == _const.Explores || fraction == _const.Reverses {
			if core {
				return OutpostCore
			} else {
				return OutpostStructure
			}
		} else {
			if core {
				return OutpostCoreNPC
			} else {
				return OutpostStructureNPC
			}
		}
	}

	return -1
}
