package factories

import (
	"github.com/TrashPony/veliri-lib/game_objects/effect"
	"github.com/TrashPony/veliri-lib/game_objects/fuel"
	"math/rand"
)

const sizeFuel = 1000
const capK = 1

var fuelTypes = map[int]fuel.Fuel{
	1: {
		ID:             1,
		Name:           "gas_fuel",
		Size:           sizeFuel,
		EnergyCap:      25000 * capK, // 25000
		LostDurability: 250,
		Bonuses:        []effect.Effect{},
	},
	2: {
		ID:             2,
		Name:           "solid_fuel",
		Size:           sizeFuel,
		EnergyCap:      37500 * capK,
		LostDurability: 250,
		Bonuses: []effect.Effect{
			{
				Name:        "effect_solid_fuel",
				Parameter:   "charging_speed",
				Quantity:    2,
				Percentages: true,
			},
			{
				Parameter:   "speed",
				Quantity:    2,
				Percentages: true,
			},
			{
				Parameter:   "turn_speed",
				Quantity:    2,
				Percentages: true,
			},
			{
				Parameter:   "reverse_speed",
				Quantity:    2,
				Percentages: true,
			},
			{
				Parameter:   "power_factor",
				Quantity:    2,
				Percentages: true,
			},
			{
				Parameter:   "reverse_factor",
				Quantity:    2,
				Percentages: true,
			},
			{
				Parameter:   "view",
				Quantity:    2,
				Percentages: true,
			},
			{
				Parameter:   "radar",
				Quantity:    2,
				Percentages: true,
			},
		},
	},
	3: {
		ID:             3,
		Name:           "liquid_fuel",
		Size:           sizeFuel,
		EnergyCap:      62500 * capK,
		LostDurability: 250,
		Bonuses: []effect.Effect{
			{
				Name:        "effect_liquid_fuel",
				Parameter:   "view",
				Quantity:    5,
				Percentages: true,
			},
			{
				Parameter:   "radar",
				Quantity:    5,
				Percentages: true,
			},
		},
	},
	4: {
		ID:             4,
		Name:           "nuclear_fuel",
		Size:           sizeFuel,
		EnergyCap:      100000 * capK,
		LostDurability: 250,
		Bonuses: []effect.Effect{
			{
				Name:        "effect_nuclear_fuel",
				Parameter:   "charging_speed",
				Quantity:    10,
				Percentages: true,
			},
		},
		Tech: 1,
	},
	5: {
		ID:             5,
		Name:           "synthetic_fuel",
		Size:           sizeFuel,
		EnergyCap:      62500 * capK,
		LostDurability: 250,
		Bonuses: []effect.Effect{
			{
				Name:        "effect_synthetic_fuel",
				Parameter:   "view",
				Quantity:    10,
				Percentages: true,
			},
			{
				Parameter:   "radar",
				Quantity:    10,
				Percentages: true,
			},
		},
		Tech: 1,
	},
	6: {
		ID:             6,
		Name:           "jet_fuel",
		Size:           sizeFuel,
		EnergyCap:      37500 * capK,
		LostDurability: 250,
		Bonuses: []effect.Effect{
			{
				Name:        "effect_jet_fuel",
				Parameter:   "speed",
				Quantity:    5,
				Percentages: true,
			},
			{
				Parameter:   "turn_speed",
				Quantity:    5,
				Percentages: true,
			},
			{
				Parameter:   "reverse_speed",
				Quantity:    5,
				Percentages: true,
			},
			{
				Parameter:   "power_factor",
				Quantity:    5,
				Percentages: true,
			},
			{
				Parameter:   "reverse_factor",
				Quantity:    5,
				Percentages: true,
			},
		},
		Tech: 1,
	},
	7: {
		ID:               7,
		Name:             "emergency_fuel_t0",
		Size:             sizeFuel / 10,
		EnergyCap:        1000 * capK,
		DestroyOnExtract: true,
		NonRefuelable:    true,
		LostDurability:   25,
		Bonuses: []effect.Effect{
			{
				Name:        "effect_emergency_fuel_t0",
				Parameter:   "charging_speed",
				Quantity:    10,
				Percentages: true,
				Subtract:    true,
			},
			{
				Parameter:   "radar",
				Quantity:    10,
				Percentages: true,
				Subtract:    true,
			},
			{
				Parameter:   "view",
				Quantity:    10,
				Percentages: true,
				Subtract:    true,
			},
		},
	},
	8: {
		ID:               8,
		Name:             "emergency_fuel_t1",
		Size:             sizeFuel / 10,
		EnergyCap:        2500 * capK,
		Tech:             1,
		DestroyOnExtract: true,
		NonRefuelable:    true,
		LostDurability:   25,
		Bonuses: []effect.Effect{
			{
				Name:        "effect_emergency_fuel_t1",
				Parameter:   "charging_speed",
				Quantity:    7,
				Percentages: true,
				Subtract:    true,
			},
			{
				Parameter:   "radar",
				Quantity:    7,
				Percentages: true,
				Subtract:    true,
			},
			{
				Parameter:   "view",
				Quantity:    7,
				Percentages: true,
				Subtract:    true,
			},
		},
	},
	9: {
		ID:               9,
		Name:             "emergency_fuel_t2",
		Size:             sizeFuel / 10,
		EnergyCap:        5000 * capK,
		Tech:             1,
		DestroyOnExtract: true,
		NonRefuelable:    true,
		LostDurability:   25,
		Bonuses: []effect.Effect{
			{
				Name:        "effect_emergency_fuel_t2",
				Parameter:   "charging_speed",
				Quantity:    5,
				Percentages: true,
				Subtract:    true,
			},
			{
				Parameter:   "radar",
				Quantity:    5,
				Percentages: true,
				Subtract:    true,
			},
			{
				Parameter:   "view",
				Quantity:    5,
				Percentages: true,
				Subtract:    true,
			},
		},
	},
}

func GetFuel(id int) *fuel.Fuel {
	newFuel := fuelTypes[id]
	return &newFuel
}

func GetAllFuel() map[int]fuel.Fuel {
	return fuelTypes
}

func GetRandomFuel() fuel.Fuel {
	fs := make([]fuel.Fuel, 0)

	for _, p := range fuelTypes {
		fs = append(fs, p)
	}

	return fs[rand.Intn(len(fs))]
}
