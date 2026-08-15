package loot_system

// ---- СЕКТОРА ----

type AnomalyView struct {
	MaxPoints int               `json:"max_points"`
	Resources []ResourceViewDTO `json:"resources"`
	Treasure  TreasureConfig    `json:"treasure"`
}

type ResourceViewDTO struct {
	ResourceID int `json:"resource_id"`
	Weight     int `json:"weight"`
	MaxCount   int `json:"max_count"`
}

// GetAllAnomalies возвращает конфиги для всех типов секторов
func GetAllAnomalies() map[string]AnomalyView {
	types := map[string]sectorType{
		"fraction_safe":   sectorSafe,
		"fraction_danger": sectorFactionDanger,
		"pirate":          sectorPirate,
		"wasteland":       sectorWasteland,
		"combat":          sectorCombat,
	}

	result := make(map[string]AnomalyView)
	for name, st := range types {
		cfg := getAnomalyConfig(st)
		resCfg := getResourceDropConfig(st)

		var resources []ResourceViewDTO
		for _, r := range resCfg.drops {
			resources = append(resources, ResourceViewDTO{
				ResourceID: r.resourceID,
				Weight:     r.weight,
				MaxCount:   r.maxCount,
			})
		}

		var treasure TreasureConfig
		if cfg.Treasure != nil {
			treasure = *cfg.Treasure
		}

		result[name] = AnomalyView{
			MaxPoints: cfg.MaxPoints,
			Resources: resources,
			Treasure:  treasure,
		}
	}
	return result
}

// ---- ФОРПОСТЫ ----

func GetAllOutposts() map[string]OutpostLootConfig {
	types := map[string]OutpostType{
		"structure_apd":      OutpostStructure,
		"core_apd":           OutpostCore,
		"structure_fraction": OutpostStructureNPC,
		"core_fraction":      OutpostCoreNPC,
	}

	result := make(map[string]OutpostLootConfig)
	for name, ot := range types {
		cfg := getOutpostLootConfig(ot)
		if cfg != nil {
			result[name] = *cfg
		}
	}
	return result
}

// ---- БОТЫ ----

func GetAllBots() map[string]BotLootConfig {
	types := map[string]botType{
		"trader":     botTrader,
		"miner":      botMiner,
		"pirate":     botPirate,
		"police":     botPolice,
		"warrior":    botWarrior,
		"agent":      botAgent,
		"builder":    botBuilder,
		"mob_small":  botMobSmall,
		"mob_mid":    botMobMid,
		"mob_heavy":  botMobHeavy,
		"mob_gigant": botMobGigant,
	}

	result := make(map[string]BotLootConfig)
	for name, bt := range types {
		cfg := getBotLootConfig(bt)
		if cfg != nil {
			result[name] = *cfg
		}
	}
	return result
}
