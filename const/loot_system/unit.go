package loot_system

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_math"
	"math/rand"
)

type botType string

const (
	botTrader    botType = "trader"
	botMiner     botType = "miner"
	botPirate    botType = "pirate"
	botPolice    botType = "police"
	botWarrior   botType = "warrior"
	botAgent     botType = "agent"
	botBuilder   botType = "builder"
	botMobSmall  botType = "mob_small"
	botMobMid    botType = "mob_mid"
	botMobHeavy  botType = "mob_heavy"
	botMobGigant botType = "mob_gigant"

	minDetailCount = 33
	maxDetailCount = 66
)

func GenerateBotLoot(botRole string, fraction string, rng *rand.Rand, k float64, size int) []LootDrop {
	botType := getTypeBot(botRole, fraction, size)

	cfg := getBotLootConfig(botType)
	if k == 0 {
		k = 1
	}

	cfg.CurrencyChance = int(float64(cfg.CurrencyChance) * (k * 3))
	cfg.BlueprintChance = int(float64(cfg.BlueprintChance) * (k * 3))
	cfg.PartChance = int(float64(cfg.PartChance) * (k * 3))

	var drops []LootDrop

	// Валюта: 1к–5к
	if rng.Intn(100) < cfg.CurrencyChance {
		amount := 1000 + int(float64(rng.Intn(3001))*k)
		drops = append(drops, LootDrop{
			LootLot: LootLot{ItemType: "trash", ItemID: 41}, // 1k_credits
			Count:   amount / 1000,
		})
	}

	// Чертежи: фракционные, 1 шт (с BaseCount)
	if rng.Intn(100) < cfg.BlueprintChance {

		blueprints := filterByFraction(t0BluePrints, fraction)
		if cfg.PartTier == 1 { // Т1 чертежи — у агентов, строителей, гигантов
			blueprints = filterByFraction(t1BluePrints, fraction)
		}

		if len(blueprints) == 0 {
			// fallback
			if cfg.PartTier == 1 {
				blueprints = t1BluePrints
			} else {
				blueprints = t0BluePrints
			}
		}

		if len(blueprints) > 0 {
			lot := blueprints[rng.Intn(len(blueprints))]
			count := game_math.GetRangeRand(1, lot.BaseCount+1, rng)
			drops = append(drops, LootDrop{LootLot: lot, Count: int(float64(count) * k)})
		}
	}

	// Детали: пачки 50–100 (Т0) или 50–100/100–200 (Т1?) — но в таблице везде "50-100"
	// Уточним: в таблице — "пачки 50-100 шт" для всех, но у пиратов/полиции — Т0, у агентов — Т1
	// Значит: min=50, max=100 всегда, tier — по cfg.PartTier
	if rng.Intn(100) < cfg.PartChance {

		pool := t0Details
		if cfg.PartTier == 1 {
			pool = t1Details
		}

		if len(pool) > 0 {
			lot := pool[rng.Intn(len(pool))]
			count := game_math.GetRangeRand(minDetailCount, maxDetailCount, rng) * game_math.GetRangeRand(1, lot.BaseCount+1, rng)
			drops = append(drops, LootDrop{LootLot: lot, Count: int(float64(count) * k)})
		}
	}

	return drops
}

type BotLootConfig struct {
	CurrencyChance  int // 5 = 5%
	BlueprintChance int
	PartChance      int
	PartTier        int // 0 или 1
}

func getBotLootConfig(botType botType) *BotLootConfig {
	switch botType {
	case botTrader:
		return &BotLootConfig{3, 6, 10, 1}
	case botMiner:
		return &BotLootConfig{3, 6, 18, 1}
	case botPirate:
		return &BotLootConfig{10, 15, 26, 0}
	case botPolice:
		return &BotLootConfig{5, 6, 10, 0}
	case botWarrior:
		return &BotLootConfig{10, 10, 17, 0}
	case botAgent:
		return &BotLootConfig{10, 15, 17, 1}
	case botBuilder:
		return &BotLootConfig{3, 15, 35, 1}
	case botMobSmall:
		return &BotLootConfig{3, 6, 10, 0}
	case botMobMid:
		return &BotLootConfig{3, 7, 17, 0}
	case botMobHeavy:
		return &BotLootConfig{5, 8, 25, 0}
	case botMobGigant:
		return &BotLootConfig{7, 11, 17, 1}
	default:
		return &BotLootConfig{3, 6, 10, 0}
	}
}

func getTypeBot(botRole string, fraction string, size int) botType {
	// дикие
	if fraction == _const.APD {
		if botRole == "in_scout" || botRole == "exp-in_scout" || botRole == "expansion-in_scout" {
			return botMobSmall
		}

		if botRole == "guard" || botRole == "exp-guard" || botRole == "expansion-guard" {
			return botMobMid
		}

		if botRole == "warrior" || botRole == "exp-warrior" || botRole == "expansion-warrior" {
			return botMobGigant
		}
	}

	if fraction == _const.FarGod {
		return botMobSmall // на текущий момент бывают только в квестах
	}

	if fraction == _const.RustbucketCartel {
		if size <= 1 {
			return botMobSmall
		}

		if size == 2 {
			return botMobMid
		}

		if size >= 3 {
			return botMobHeavy
		}
	}

	// фракционные
	if fraction == _const.Replicas || fraction == _const.Explores || fraction == _const.Reverses {
		if botRole == "miner" || botRole == "exp-miner" {
			return botMiner
		}

		if botRole == "guard" || botRole == "guard_trader" || botRole == "exp-guard" || botRole == "expansion-guard" {
			return botPolice
		}

		if botRole == "warrior" || botRole == "exp-warrior" || botRole == "war-warrior" || botRole == "guard-warrior" || botRole == "expansion-warrior" {
			return botWarrior
		}

		if botRole == "out_scout" || botRole == "exp-out_scout" {
			return botAgent
		}

		if botRole == "in_scout" || botRole == "exp-in_scout" || botRole == "expansion-in_scout" {
			return botPirate
		}

		if botRole == "transport" || botRole == "trader" || botRole == "exp-transport" {
			return botTrader
		}

		if botRole == "builder" || botRole == "exp-builder" {
			return botBuilder
		}
	}

	return botMobSmall
}
