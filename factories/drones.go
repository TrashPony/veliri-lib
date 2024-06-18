package factories

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_objects/anchor"
	"github.com/TrashPony/veliri-lib/game_objects/coordinate"
	"github.com/TrashPony/veliri-lib/game_objects/drone"
	"strings"
)

func Drones() *factory {
	return drones
}

type factory struct {
	drones map[string]drone.Drone
	// read-only sync.Mutex
}

func (f *factory) GetDroneByID(name string) drone.Drone {
	return f.drones[name]
}

func (f *factory) GetAllType() map[string]drone.Drone {

	for k, d := range f.drones {
		d.SetEngagePositions()
		f.drones[k] = d
	}

	return f.drones
}

var drones = &factory{
	drones: map[string]drone.Drone{
		"transport_" + strings.ToLower(_const.Replicas): {
			Sprite:   "transport_" + strings.ToLower(_const.Replicas),
			Fraction: _const.Replicas,
			EngagePosNoScale: map[string]coordinate.Coordinate{
				"antigravity_place": {X: 64, Y: 64},
				"1_engage":          {X: 15, Y: 64},
			},
			EngageAnchors: map[string]anchor.Anchor{
				"antigravity_place": {Type: "1", Scale: 1},
				"1_engage":          {Type: "2", Scale: 1},
			},
		},
		"transport_" + strings.ToLower(_const.Explores): {
			Sprite:   "transport_" + strings.ToLower(_const.Explores),
			Fraction: _const.Explores,
			EngagePosNoScale: map[string]coordinate.Coordinate{
				"antigravity_place": {X: 64, Y: 64},
			},
			EngageAnchors: map[string]anchor.Anchor{
				"antigravity_place": {Type: "1", Scale: 1},
			},
		},
		"transport_" + strings.ToLower(_const.Reverses): {
			Sprite:   "transport_" + strings.ToLower(_const.Reverses),
			Fraction: _const.Reverses,
			EngagePosNoScale: map[string]coordinate.Coordinate{
				"antigravity_place": {X: 64, Y: 64},
			},
			EngageAnchors: map[string]anchor.Anchor{
				"antigravity_place": {Type: "1", Scale: 1},
			},
		},
		"drone_scout_1": {
			EngagePosNoScale: map[string]coordinate.Coordinate{
				"antigravity_place": {X: 47, Y: 64},
				"1_engage":          {X: 61, Y: 52},
				"2_engage":          {X: 61, Y: 76},
			},
			EngageAnchors: map[string]anchor.Anchor{
				"antigravity_place": {Type: "1", Scale: 0.35},
				"1_engage":          {Type: "3", Height: 3, Rotate: 45, Scale: 0.5},
				"2_engage":          {Type: "3", Height: 3, Rotate: -45, Scale: 0.5},
			},
		},
		"digger": {
			EngagePosNoScale: map[string]coordinate.Coordinate{
				"antigravity_place": {X: 64, Y: 64},
				"1_engage":          {X: 43, Y: 64},
			},
			EngageAnchors: map[string]anchor.Anchor{
				"antigravity_place": {Type: "1", Scale: 0.35},
				"1_engage":          {Type: "2", Scale: 0.5},
			},
		},
		"drone_defender_1": {
			EngagePosNoScale: map[string]coordinate.Coordinate{
				"antigravity_place": {X: 64, Y: 64},
				"1_engage":          {X: 37, Y: 64},
			},
			EngageAnchors: map[string]anchor.Anchor{
				"antigravity_place": {Type: "1", Scale: 0.35},
				"1_engage":          {Type: "2", Scale: 0.75},
			},
		},
		"drone_defender_2": {
			EngagePosNoScale: map[string]coordinate.Coordinate{
				"antigravity_place": {X: 64, Y: 64},
				"1_engage":          {X: 40, Y: 64},
			},
			EngageAnchors: map[string]anchor.Anchor{
				"antigravity_place": {Type: "1", Scale: 0.35},
				"1_engage":          {Type: "2", Scale: 0.75},
			},
		},
		"fly_fauna_drone_1": {
			EngagePosNoScale: map[string]coordinate.Coordinate{
				"antigravity_place": {X: 64, Y: 64},
				"1_engage":          {X: 30, Y: 22},
				"2_engage":          {X: 30, Y: 105},
				"3_engage":          {X: 13, Y: 46},
				"4_engage":          {X: 13, Y: 82},
			},
			EngageAnchors: map[string]anchor.Anchor{
				"antigravity_place": {Type: "1", Scale: 0.5},
				"1_engage":          {Type: "2", Height: 3, Rotate: 1, Scale: 0.5},
				"2_engage":          {Type: "2", Height: 3, Rotate: -1, Scale: 0.5},
				"3_engage":          {Type: "2", Rotate: 1, Scale: 0.25},
				"4_engage":          {Type: "2", Rotate: -1, Scale: 0.25},
			},
		},
		"fly_fauna_drone_2": {
			EngagePosNoScale: map[string]coordinate.Coordinate{
				"antigravity_place": {X: 64, Y: 64},
				"1_engage":          {X: 10, Y: 37},
				"2_engage":          {X: 10, Y: 90},
				"3_engage":          {X: 5, Y: 64},
			},
			EngageAnchors: map[string]anchor.Anchor{
				"antigravity_place": {Type: "1", Scale: 0.5},
				"1_engage":          {Type: "2", Height: 3, Rotate: 1, Scale: 0.5},
				"2_engage":          {Type: "2", Height: 3, Rotate: -1, Scale: 0.5},
				"3_engage":          {Type: "2", Height: 3, Scale: 1.5},
			},
		},
		"fly_fauna_drone_3": {
			EngagePosNoScale: map[string]coordinate.Coordinate{
				"antigravity_place": {X: 64, Y: 64},
				"1_engage":          {X: 40, Y: 14},
				"2_engage":          {X: 40, Y: 114},
				"3_engage":          {X: 6, Y: 64},
			},
			EngageAnchors: map[string]anchor.Anchor{
				"antigravity_place": {Type: "1", Scale: 0.5},
				"1_engage":          {Type: "2", Height: 3, Rotate: 1, Scale: 0.5},
				"2_engage":          {Type: "2", Height: 3, Rotate: -1, Scale: 0.5},
				"3_engage":          {Type: "2", Height: 3, Scale: 1.5},
			},
		},
		"builder_1": {
			EngagePosNoScale: map[string]coordinate.Coordinate{
				"antigravity_place": {X: 64, Y: 64},
				"1_engage":          {X: 30, Y: 22},
				"2_engage":          {X: 30, Y: 105},
				"3_engage":          {X: 13, Y: 46},
				"4_engage":          {X: 13, Y: 82},
			},
			EngageAnchors: map[string]anchor.Anchor{
				"antigravity_place": {Type: "1", Scale: 0.15},
				"1_engage":          {Type: "2", Height: 3, Rotate: 1, Scale: 0.5},
				"2_engage":          {Type: "2", Height: 3, Rotate: -1, Scale: 0.5},
				"3_engage":          {Type: "2", Rotate: 1, Scale: 0.25},
				"4_engage":          {Type: "2", Rotate: -1, Scale: 0.25},
			},
		},
	},
}
