package rpc

import (
	"encoding/gob"
	"github.com/TrashPony/veliri-lib/game_objects/behavior_rule"
	"github.com/TrashPony/veliri-lib/game_objects/coordinate"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"github.com/TrashPony/veliri-lib/game_objects/dynamic_map_object"
	"github.com/TrashPony/veliri-lib/game_objects/effects_store"
	"github.com/TrashPony/veliri-lib/game_objects/lvl_map"
	"github.com/TrashPony/veliri-lib/game_objects/notify"
	"github.com/TrashPony/veliri-lib/game_objects/reservoir"
	"github.com/TrashPony/veliri-lib/game_objects/rpc_request"
	"github.com/TrashPony/veliri-lib/game_objects/select_mouse_equip"
	"github.com/TrashPony/veliri-lib/game_objects/skin"
	"github.com/TrashPony/veliri-lib/game_objects/spawn"
	"github.com/TrashPony/veliri-lib/game_objects/target"
	"github.com/TrashPony/veliri-lib/game_objects/unit"
	"github.com/valyala/gorpc"
)

type RPC struct {
	Server           *gorpc.Server
	VeliriMainClient *gorpc.Client
	Error            bool
}

func GobRegister() {
	gob.Register(rpc_request.Request{})
	gob.Register(map[int]string{})
	gob.Register(map[string]int{})
	gob.Register(map[int][]byte{})
	gob.Register(map[string]interface{}{})

	gob.Register(map[int][]*coordinate.Coordinate{})
	gob.Register([]*coordinate.Coordinate{})
	gob.Register(coordinate.Coordinate{})

	gob.Register(map[int]*skin.Skin{})
	gob.Register(skin.Skin{})

	gob.Register(map[int]map[int]lvl_map.LvlMap{})
	gob.Register(dynamic_map_object.Object{})
	gob.Register(reservoir.Reservoir{})
	gob.Register(notify.Notify{})
	gob.Register([]*spawn.Spawn{})
	gob.Register([]unit.Damage{})
	gob.Register(behavior_rule.BehaviorRules{})
	gob.Register(select_mouse_equip.SelectMouseEquip{})
	gob.Register(target.Target{})
	gob.Register(effects_store.EffectsStore{})
	gob.Register(detail.Body{})
	gob.Register(map[int]map[int]*dynamic_map_object.Flore{})
}
