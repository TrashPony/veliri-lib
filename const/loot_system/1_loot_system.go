package loot_system

import _const "github.com/TrashPony/veliri-lib/const"

type LootLot struct {
	ItemType  string `json:"item_type"`
	ItemID    int    `json:"item_id"`
	Fraction  string `json:"fraction"`
	BaseCount int    `json:"base_count"` // умножается на count из конкретной системы [RAND(BaseCount+1) * SYSTEM_COUNT]
}

// Детали для изготовления не требуются другие детали
var t0Details = []LootLot{
	{ItemType: "detail", ItemID: 2},
	{ItemType: "detail", ItemID: 3},
	{ItemType: "detail", ItemID: 5},
	{ItemType: "detail", ItemID: 6},
	{ItemType: "detail", ItemID: 9},
}

// Детали которые делаютяс из другиз деталей
var t1Details = []LootLot{
	{ItemType: "detail", ItemID: 1},
	{ItemType: "detail", ItemID: 4},
	{ItemType: "detail", ItemID: 7},
	{ItemType: "detail", ItemID: 8},
}

// Чертежи снаряжения, снарядов, деталей, некоторых видов оружи, корпусов и модулей. Составлялся ручками
var t0BluePrints = []LootLot{
	// detail
	{ItemType: "blueprints", ItemID: 230}, //electronics_500_bp
	{ItemType: "blueprints", ItemID: 231}, //steel_500_bp
	{ItemType: "blueprints", ItemID: 232}, //wire_500_bp
	{ItemType: "blueprints", ItemID: 233}, //wires_500_bp
	{ItemType: "blueprints", ItemID: 234}, //gear_500_bp
	{ItemType: "blueprints", ItemID: 235}, //titanium_plate_500_bp
	{ItemType: "blueprints", ItemID: 236}, //batteries_500_bp
	{ItemType: "blueprints", ItemID: 237}, //armor_items_500_bp
	{ItemType: "blueprints", ItemID: 238}, //carbon_plate_500_bp

	// ammo - тут только те которые нельзя купить на базе
	{ItemType: "blueprints", ItemID: 185}, // piu-piu_3
	{ItemType: "blueprints", ItemID: 186}, // piu-piu_4
	{ItemType: "blueprints", ItemID: 187}, // piu-piu_5
	{ItemType: "blueprints", ItemID: 188}, // piu-piu_6
	{ItemType: "blueprints", ItemID: 189}, // ballistics_artillery_bullet_3
	{ItemType: "blueprints", ItemID: 190}, // ballistics_artillery_bullet_4
	{ItemType: "blueprints", ItemID: 192}, // big_lens_1
	{ItemType: "blueprints", ItemID: 194}, // medium_lens_beam

	// equip
	{ItemType: "blueprints", ItemID: 198, Fraction: _const.Reverses}, // reverses_combine_bp
	{ItemType: "blueprints", ItemID: 207, Fraction: _const.Explores}, // dust_collector
	{ItemType: "blueprints", ItemID: 103, Fraction: _const.Replicas}, // replic_miner_extracted_bp
	{ItemType: "blueprints", ItemID: 104, Fraction: _const.Explores}, // explores_miner_extracted_bp
	{ItemType: "blueprints", ItemID: 105, Fraction: _const.Reverses}, // reverses_miner_extracted_bp
	{ItemType: "blueprints", ItemID: 221},                            // scientific_module_1_bp
	{ItemType: "blueprints", ItemID: 142},                            // pump_extracted_bp
	{ItemType: "blueprints", ItemID: 20},                             // geo_scanner_bp
	{ItemType: "blueprints", ItemID: 66},                             // distance_repair_bp

	{ItemType: "blueprints", ItemID: 127}, // mine_bomb_1_bp
	{ItemType: "blueprints", ItemID: 128}, // mini_turret_1_bp
	{ItemType: "blueprints", ItemID: 129}, // smoke_screen_local_bp
	{ItemType: "blueprints", ItemID: 130}, // smoke_screen_destination_bp
	{ItemType: "blueprints", ItemID: 131}, // drone_scout_1_bp
	{ItemType: "blueprints", ItemID: 132}, // drone_defender_1_bp
	{ItemType: "blueprints", ItemID: 133}, // rope_trap_1_bp
	{ItemType: "blueprints", ItemID: 134}, // grenade_1_bp
	{ItemType: "blueprints", ItemID: 135}, // gravity_grenade_1_bp
	{ItemType: "blueprints", ItemID: 136}, // jump_drive_1_bp
	{ItemType: "blueprints", ItemID: 137}, // invisibility_1_bp
	{ItemType: "blueprints", ItemID: 138}, // mini_turret_2_bp
	{ItemType: "blueprints", ItemID: 139}, // drone_defender_2_bp
	{ItemType: "blueprints", ItemID: 140}, // wall_1_bp
	{ItemType: "blueprints", ItemID: 141}, // energy_shield_mini_structure_1_bp
	{ItemType: "blueprints", ItemID: 63},  // jammer_bp
	{ItemType: "blueprints", ItemID: 64},  // missile_defense_bp
	{ItemType: "blueprints", ItemID: 13},  // repair_kit_bp
	{ItemType: "blueprints", ItemID: 14},  // energy_shield_bp

	{ItemType: "blueprints", ItemID: 17},  // armored_bp,
	{ItemType: "blueprints", ItemID: 65},  // repair_passive_bp
	{ItemType: "blueprints", ItemID: 144}, // radar_booster_1
	{ItemType: "blueprints", ItemID: 145}, // view_booster_1
	{ItemType: "blueprints", ItemID: 146}, // improved_inventory_capacity_1
	{ItemType: "blueprints", ItemID: 147}, // antigravity_speed_booster_1
	{ItemType: "blueprints", ItemID: 148}, // caterpillars_speed_booster_1
	{ItemType: "blueprints", ItemID: 149}, // antigravity_mobility_booster_1_bp
	{ItemType: "blueprints", ItemID: 150}, // wheels_mobility_booster_1_bp
	{ItemType: "blueprints", ItemID: 151}, // ballistic_rotate_booster_1
	{ItemType: "blueprints", ItemID: 152}, // ballistic_damage_booster_1
	{ItemType: "blueprints", ItemID: 153}, // missile_rotate_booster_1
	{ItemType: "blueprints", ItemID: 154}, // missile_damage_booster_1
	{ItemType: "blueprints", ItemID: 155}, // laser_rotate_booster_1
	{ItemType: "blueprints", ItemID: 156}, // laser_damage_booster_1
	{ItemType: "blueprints", ItemID: 157}, // energy_capacity_1
	{ItemType: "blueprints", ItemID: 158}, // energy_charging_speed_1
	{ItemType: "blueprints", ItemID: 159}, // power_booster_1
	{ItemType: "blueprints", ItemID: 160}, // body_structure_booster_1_bp

	// weapon
	{ItemType: "blueprints", ItemID: 223},                            //melee_1_bp
	{ItemType: "blueprints", ItemID: 162, Fraction: _const.Reverses}, //reverses_weapon_5_bp
	{ItemType: "blueprints", ItemID: 78, Fraction: _const.Replicas},  //replic_weapon_2_bp
	{ItemType: "blueprints", ItemID: 79, Fraction: _const.Replicas},  //replic_weapon_3_bp
	{ItemType: "blueprints", ItemID: 80, Fraction: _const.Replicas},  //replic_weapon_4_bp
	{ItemType: "blueprints", ItemID: 82, Fraction: _const.Explores},  //explores_weapon_2_bp
	{ItemType: "blueprints", ItemID: 83, Fraction: _const.Explores},  //explores_weapon_3_bp
	{ItemType: "blueprints", ItemID: 84, Fraction: _const.Explores},  //explores_weapon_4_bp
	{ItemType: "blueprints", ItemID: 86, Fraction: _const.Reverses},  //reverses_weapon_2_bp
	{ItemType: "blueprints", ItemID: 87, Fraction: _const.Reverses},  //reverses_weapon_3_bp

	// body
	{ItemType: "blueprints", ItemID: 100, Fraction: _const.Replicas}, //replic_miner_bp
	{ItemType: "blueprints", ItemID: 101, Fraction: _const.Explores}, //explores_miner_bp
	{ItemType: "blueprints", ItemID: 102, Fraction: _const.Reverses}, //reverses_miner_bp
	{ItemType: "blueprints", ItemID: 124, Fraction: _const.Reverses}, //reverses_start_body_bp
	{ItemType: "blueprints", ItemID: 110, Fraction: _const.Explores}, //explores_start_body_bp
	{ItemType: "blueprints", ItemID: 119, Fraction: _const.Replicas}, //replic_start_body_bp
	{ItemType: "blueprints", ItemID: 50, Fraction: _const.Replicas},  //replic_scout_bp
	{ItemType: "blueprints", ItemID: 55, Fraction: _const.Explores},  //explores_scout_bp
	{ItemType: "blueprints", ItemID: 60, Fraction: _const.Reverses},  //reverses_scout_bp,
}

// Чертежи снаряжения, снарядов, деталей, некоторых видов оружи, корпусов и модулей. Составлялся ручками
var t1BluePrints = []LootLot{
	// detail
	{ItemType: "blueprints", ItemID: 230, BaseCount: 3}, //electronics_500_bp
	{ItemType: "blueprints", ItemID: 231, BaseCount: 3}, //steel_500_bp
	{ItemType: "blueprints", ItemID: 232, BaseCount: 3}, //wire_500_bp
	{ItemType: "blueprints", ItemID: 233, BaseCount: 3}, //wires_500_bp
	{ItemType: "blueprints", ItemID: 234, BaseCount: 3}, //gear_500_bp
	{ItemType: "blueprints", ItemID: 235, BaseCount: 3}, //titanium_plate_500_bp
	{ItemType: "blueprints", ItemID: 236, BaseCount: 3}, //batteries_500_bp
	{ItemType: "blueprints", ItemID: 237, BaseCount: 3}, //armor_items_500_bp
	{ItemType: "blueprints", ItemID: 238, BaseCount: 3}, //carbon_plate_500_bp

	// ammo - тут только те которые нельзя купить на базе
	{ItemType: "blueprints", ItemID: 185, BaseCount: 3}, // piu-piu_3
	{ItemType: "blueprints", ItemID: 186, BaseCount: 3}, // piu-piu_4
	{ItemType: "blueprints", ItemID: 187, BaseCount: 3}, // piu-piu_5
	{ItemType: "blueprints", ItemID: 188, BaseCount: 3}, // piu-piu_6
	{ItemType: "blueprints", ItemID: 189, BaseCount: 3}, // ballistics_artillery_bullet_3
	{ItemType: "blueprints", ItemID: 190, BaseCount: 3}, // ballistics_artillery_bullet_4
	{ItemType: "blueprints", ItemID: 192, BaseCount: 3}, // big_lens_1
	{ItemType: "blueprints", ItemID: 194, BaseCount: 3}, // medium_lens_beam

	// equip
	{ItemType: "blueprints", ItemID: 165}, // rig_armored
	{ItemType: "blueprints", ItemID: 166}, // rig_armored_explosion
	{ItemType: "blueprints", ItemID: 167}, // rig_armored_kinetics
	{ItemType: "blueprints", ItemID: 168}, // rig_armored_thermo
	{ItemType: "blueprints", ItemID: 169}, // rig_body_structure_booster
	{ItemType: "blueprints", ItemID: 170}, // rig_caterpillars_speed_booster_1
	{ItemType: "blueprints", ItemID: 171}, // rig_antigravity_speed_booster_1
	{ItemType: "blueprints", ItemID: 172}, // rig_wheels_speed_booster_1
	{ItemType: "blueprints", ItemID: 173}, // rig_radar_booster_1
	{ItemType: "blueprints", ItemID: 174}, // rig_view_booster_1
	{ItemType: "blueprints", ItemID: 175}, // rig_energy_capacity_1
	{ItemType: "blueprints", ItemID: 176}, // rig_energy_charging_speed_1
	{ItemType: "blueprints", ItemID: 177}, // rig_improved_inventory_capacity_1
	{ItemType: "blueprints", ItemID: 178}, // rig_power_booster_1
	{ItemType: "blueprints", ItemID: 179}, // rig_missile_damage_booster_1
	{ItemType: "blueprints", ItemID: 180}, // rig_laser_damage_booster_1
	{ItemType: "blueprints", ItemID: 181}, // rig_ballistic_damage_booster_1
	{ItemType: "blueprints", ItemID: 182}, // rig_ballistic_rotate_booster_1
	{ItemType: "blueprints", ItemID: 183}, // rig_missile_rotate_booster_1
	{ItemType: "blueprints", ItemID: 184}, // rig_laser_rotate_booster_1

	{ItemType: "blueprints", ItemID: 69, Fraction: _const.Replicas}, // replic_builder_bp
	{ItemType: "blueprints", ItemID: 71, Fraction: _const.Explores}, // explores_builder_bp
	{ItemType: "blueprints", ItemID: 73, Fraction: _const.Reverses}, // reverses_builder_bp
	{ItemType: "blueprints", ItemID: 143},                           // inventory_scanner_1_bp

	// weapon
	{ItemType: "blueprints", ItemID: 222},                            //melee_2_bp
	{ItemType: "blueprints", ItemID: 77, Fraction: _const.Replicas},  //replic_weapon_1_bp
	{ItemType: "blueprints", ItemID: 195, Fraction: _const.Explores}, //explores_weapon_5
	{ItemType: "blueprints", ItemID: 196, Fraction: _const.Explores}, //explores_weapon_6
	{ItemType: "blueprints", ItemID: 81, Fraction: _const.Explores},  //explores_weapon_1_bp
	{ItemType: "blueprints", ItemID: 85, Fraction: _const.Reverses},  //reverses_weapon_1_bp
	{ItemType: "blueprints", ItemID: 88, Fraction: _const.Reverses},  //reverses_weapon_4_bp

	// body
	{ItemType: "blueprints", ItemID: 48, Fraction: _const.Replicas},  //replic_builder_body_bp
	{ItemType: "blueprints", ItemID: 51, Fraction: _const.Replicas},  //replic_transport_bp
	{ItemType: "blueprints", ItemID: 52, Fraction: _const.Replicas},  //replic_warrior_bp
	{ItemType: "blueprints", ItemID: 53, Fraction: _const.Explores},  //explores_builder_body_bp
	{ItemType: "blueprints", ItemID: 56, Fraction: _const.Explores},  //explores_transport_bp
	{ItemType: "blueprints", ItemID: 57, Fraction: _const.Explores},  //explores_warrior_bp
	{ItemType: "blueprints", ItemID: 58, Fraction: _const.Reverses},  //reverses_builder_body_bp
	{ItemType: "blueprints", ItemID: 61, Fraction: _const.Reverses},  //reverses_transport_bp
	{ItemType: "blueprints", ItemID: 62, Fraction: _const.Reverses},  //reverses_warrior_bp
	{ItemType: "blueprints", ItemID: 120, Fraction: _const.Replicas}, //replic_pirat_bp
	{ItemType: "blueprints", ItemID: 121, Fraction: _const.Replicas}, //replic_transport_2_bp
	{ItemType: "blueprints", ItemID: 122, Fraction: _const.Explores}, //explores_pirat_bp
	{ItemType: "blueprints", ItemID: 123, Fraction: _const.Explores}, //explores_transport_2_bp
	{ItemType: "blueprints", ItemID: 125, Fraction: _const.Reverses}, //reverses_pirat_bp
	{ItemType: "blueprints", ItemID: 126, Fraction: _const.Reverses}, //reverses_transport_2_bp
	{ItemType: "blueprints", ItemID: 197, Fraction: _const.Replicas}, //replic_oil_body
}

// Научные данные — типы
var frr = []LootLot{
	{ItemType: "frr", ItemID: 1},
	{ItemType: "frr", ItemID: 2},
	{ItemType: "frr", ItemID: 3},
	{ItemType: "frr", ItemID: 4},
	{ItemType: "frr", ItemID: 5},
	{ItemType: "frr", ItemID: 6},
	{ItemType: "frr", ItemID: 7},
}

// Товары — по грейдам
var products = map[int][]LootLot{
	1: {
		{ItemType: "product", ItemID: 4},
	},
	2: {
		{ItemType: "product", ItemID: 2},
	},
	3: {
		{ItemType: "product", ItemID: 7},
	},
	4: {
		{ItemType: "product", ItemID: 1},
	},
	5: {
		{ItemType: "product", ItemID: 9},
	},
	6: {
		{ItemType: "product", ItemID: 6},
	},
	7: {
		{ItemType: "product", ItemID: 3},
	},
	8: {
		{ItemType: "product", ItemID: 8},
	},
	9: {
		{ItemType: "product", ItemID: 5},
	},
}

func filterByFraction(pool []LootLot, fraction string) []LootLot {
	var result []LootLot
	for _, lot := range pool {
		if lot.Fraction == fraction || lot.Fraction == "" {
			result = append(result, lot)
		}
	}
	return result
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
