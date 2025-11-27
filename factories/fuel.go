package factories

import (
	"github.com/TrashPony/veliri-lib/game_objects/effect"
	"github.com/TrashPony/veliri-lib/game_objects/fuel"
	"math/rand"
)

const sizeFuel = 1000
const capK = 2

var fuelTypes = map[int]fuel.Fuel{
	1: {
		ID:        1,
		Name:      "gas_fuel",
		Size:      sizeFuel,
		EnergyCap: 25000 * capK, // 25000
		Bonuses:   []effect.Effect{},
	},
	2: {
		ID:        2,
		Name:      "solid_fuel",
		Size:      sizeFuel,
		EnergyCap: 37500 * capK,
		Bonuses:   []effect.Effect{},
	},
	3: {
		ID:        3,
		Name:      "liquid_fuel",
		Size:      sizeFuel,
		EnergyCap: 62500 * capK,
		Bonuses:   []effect.Effect{},
	},
	4: {
		ID:        4,
		Name:      "nuclear_fuel",
		Size:      sizeFuel,
		EnergyCap: 100000 * capK,
		Bonuses:   []effect.Effect{},
		Tech:      1,
	},
	5: {
		ID:        5,
		Name:      "synthetic_fuel",
		Size:      sizeFuel,
		EnergyCap: 62500 * capK,
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
		Tech: 1,
	},
	6: {
		ID:        6,
		Name:      "jet_fuel",
		Size:      sizeFuel,
		EnergyCap: 37500 * capK,
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
		Tech: 1,
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
