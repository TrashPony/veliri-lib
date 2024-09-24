package factories

import (
	"github.com/TrashPony/veliri-lib/game_objects/effect"
	"github.com/TrashPony/veliri-lib/game_objects/fuel"
	"math/rand"
)

const size = 100

var fuelTypes = map[int]fuel.Fuel{
	1: {
		ID:        1,
		Name:      "gas_fuel",
		Size:      size,
		EnergyCap: 25000, // 25000
		Bonuses:   []effect.Effect{},
	},
	2: {
		ID:        2,
		Name:      "solid_fuel",
		Size:      size,
		EnergyCap: 37500,
		Bonuses:   []effect.Effect{},
	},
	3: {
		ID:        3,
		Name:      "liquid_fuel",
		Size:      size,
		EnergyCap: 62500,
		Bonuses:   []effect.Effect{},
	},
	4: {
		ID:        4,
		Name:      "nuclear_fuel",
		Size:      size,
		EnergyCap: 100000,
		Bonuses:   []effect.Effect{},
	},
	5: {
		ID:        5,
		Name:      "synthetic_fuel",
		Size:      size,
		EnergyCap: 62500,
		Bonuses: []effect.Effect{
			{
				Name:        "effect_synthetic_fuel",
				Parameter:   "charging_speed",
				Quantity:    1,
				Percentages: true,
			},
			{
				Parameter:   "speed",
				Quantity:    1,
				Percentages: true,
			},
			{
				Parameter:   "turn_speed",
				Quantity:    1,
				Percentages: true,
			},
			{
				Parameter:   "reverse_speed",
				Quantity:    1,
				Percentages: true,
			},
			{
				Parameter:   "power_factor",
				Quantity:    1,
				Percentages: true,
			},
			{
				Parameter:   "reverse_factor",
				Quantity:    1,
				Percentages: true,
			},
		},
	},
	6: {
		ID:        6,
		Name:      "jet_fuel",
		Size:      size,
		EnergyCap: 37500,
		Bonuses: []effect.Effect{
			{
				Name:        "effect_jet_fuel",
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
