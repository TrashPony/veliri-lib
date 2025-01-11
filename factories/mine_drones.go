package factories

import (
	"github.com/TrashPony/veliri-lib/game_objects/mine_drone"
)

var mineDroneTypes = map[int]mine_drone.MineDrone{
	1: {
		ID:     1,
		Name:   "mine_drone",
		Size:   5000,
		BodyID: 51,
	},
	2: {
		ID:     2,
		Name:   "builder_drone",
		Size:   5000,
		BodyID: 53,
	},
	3: {
		ID:     3,
		Name:   "transport_drone",
		Size:   7500,
		BodyID: 52,
	},
}

func GetMineDrone(id int) *mine_drone.MineDrone {
	newFuel := mineDroneTypes[id]
	return &newFuel
}

func GetAllMineDrones() map[int]mine_drone.MineDrone {
	return mineDroneTypes
}
