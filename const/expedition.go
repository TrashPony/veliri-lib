package _const

import "math/rand"

var FractionMapExpedition = map[string][]int{
	Reverses: {
		55, 56, 57, 58, 49, 50, 51, 52,
	},
	Replicas: {
		41, 42, 43, 44, 47, 48, 49, 50,
	},
	Explores: {
		43, 44, 45, 46, 53, 54, 55, 58,
	},
}

func GetRandomExpeditionSector(fraction string) int {
	sectors, ok := FractionMapExpedition[fraction]
	if !ok {
		return 0
	}

	return sectors[rand.Intn(len(sectors))]
}

func GetFractionBySector(id int, excludeFraction string) string {
	for fraction, ids := range FractionMapExpedition {
		if fraction == excludeFraction {
			continue
		}

		for _, fid := range ids {
			if fid == id {
				return fraction
			}
		}
	}

	return ""
}
