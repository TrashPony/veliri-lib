package base

import (
	"github.com/TrashPony/veliri-lib/game_objects/coordinate"
)

type Base struct {
	ID            int                      `json:"id"`
	Name          string                   `json:"name"`
	Type          string                   `json:"type"`
	X             int                      `json:"x"`
	Y             int                      `json:"y"`
	MapID         int                      `json:"map_id"`
	GravityRadius int                      `json:"gravity_radius"`
	Respawns      []*coordinate.Coordinate `json:"respawns"`
	Fraction      string                   `json:"fraction"`
	Capital       bool                     `json:"capital"`
}
