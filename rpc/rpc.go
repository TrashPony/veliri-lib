package rpc

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/ammo"
	"github.com/TrashPony/veliri-lib/game_objects/anomaly"
	"github.com/TrashPony/veliri-lib/game_objects/behavior_rule"
	"github.com/TrashPony/veliri-lib/game_objects/blueprints"
	"github.com/TrashPony/veliri-lib/game_objects/box"
	"github.com/TrashPony/veliri-lib/game_objects/box_trap"
	"github.com/TrashPony/veliri-lib/game_objects/build_option"
	"github.com/TrashPony/veliri-lib/game_objects/builder_game_state"
	"github.com/TrashPony/veliri-lib/game_objects/coordinate"
	"github.com/TrashPony/veliri-lib/game_objects/damage_manager"
	"github.com/TrashPony/veliri-lib/game_objects/db"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"github.com/TrashPony/veliri-lib/game_objects/drone"
	"github.com/TrashPony/veliri-lib/game_objects/dynamic_map_object"
	"github.com/TrashPony/veliri-lib/game_objects/effect"
	"github.com/TrashPony/veliri-lib/game_objects/effects_store"
	"github.com/TrashPony/veliri-lib/game_objects/equip"
	"github.com/TrashPony/veliri-lib/game_objects/generate_map_config"
	"github.com/TrashPony/veliri-lib/game_objects/info_map"
	"github.com/TrashPony/veliri-lib/game_objects/inventory"
	"github.com/TrashPony/veliri-lib/game_objects/lvl_map"
	"github.com/TrashPony/veliri-lib/game_objects/map_item"
	"github.com/TrashPony/veliri-lib/game_objects/notify"
	"github.com/TrashPony/veliri-lib/game_objects/npc_request"
	"github.com/TrashPony/veliri-lib/game_objects/obstacle_point"
	"github.com/TrashPony/veliri-lib/game_objects/physical_model"
	"github.com/TrashPony/veliri-lib/game_objects/product"
	"github.com/TrashPony/veliri-lib/game_objects/relation_bonuses"
	"github.com/TrashPony/veliri-lib/game_objects/reservoir"
	"github.com/TrashPony/veliri-lib/game_objects/resource"
	"github.com/TrashPony/veliri-lib/game_objects/rpc_request"
	"github.com/TrashPony/veliri-lib/game_objects/select_mouse_equip"
	"github.com/TrashPony/veliri-lib/game_objects/skill"
	"github.com/TrashPony/veliri-lib/game_objects/skin"
	"github.com/TrashPony/veliri-lib/game_objects/spawn"
	"github.com/TrashPony/veliri-lib/game_objects/target"
	"github.com/TrashPony/veliri-lib/game_objects/trash_item"
	"github.com/TrashPony/veliri-lib/game_objects/unit"
	"github.com/TrashPony/veliri-lib/game_objects/violator"
	"github.com/TrashPony/veliri-lib/game_objects/visible_objects"
	"github.com/valyala/gorpc"
	"log"
	"runtime/debug"
)

type RPC struct {
	Server           *gorpc.Server
	VeliriMainClient *gorpc.Client
	Error            bool
}

func CheckMsg(request interface{}) {
	return

	defer func() {
		err := recover()
		if err != nil {
			debug.PrintStack()
			fmt.Println("CheckMsg: ", request, err)
			data, e := json.Marshal(request)
			if e != nil {
				fmt.Println("json_data: ", data)
			}
		}
	}()

	if request == nil {
		return
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(request); err != nil {
		log.Fatal(err)
	}

	return
}

func GobRegister() {
	gob.Register(map[int]string{})
	gob.Register(map[string]int{})
	gob.Register(map[int][]byte{})
	gob.Register(map[string]interface{}{})
	gob.Register(map[int]map[string]interface{}{})
	gob.Register([]map[string]interface{}{})
	gob.Register(map[int]interface{}{})
	gob.Register(map[string]string{})
	gob.Register(map[string]bool{})

	gob.Register(map[int]*info_map.InfoMap{})
	gob.Register(map[int]map[int]lvl_map.LvlMap{})

	gob.Register(reservoir.Reservoir{})
	gob.Register(notify.Notify{})
	gob.Register([]*spawn.Spawn{})
	gob.Register([]damage_manager.Damage{})
	gob.Register(behavior_rule.BehaviorRules{})
	gob.Register(select_mouse_equip.SelectMouseEquip{})
	gob.Register(target.Target{})
	gob.Register(effects_store.EffectsStore{})
	gob.Register(map[int]map[int]*dynamic_map_object.Flore{})
	gob.Register(map_item.MapItem{})
	gob.Register([]drone.Drone{})
	gob.Register(rpc_request.Request{})
	gob.Register(violator.Violator{})
	gob.Register(obstacle_point.ObstaclePoint{})
	gob.Register(box_trap.BoxTrap{})
	gob.Register(behavior_rule.BehaviorRule{})
	gob.Register(npc_request.DialogRequest{})
	gob.Register(behavior_rule.Meta{})
	gob.Register(map[int]int{})

	gob.Register([]*build_option.BuildOption{})
	gob.Register([]build_option.BaseConfig{})

	gob.Register(info_map.InfoMap{})
	gob.Register(behavior_rule.SubGroup{})
	gob.Register(map[string]*skill.Skill{})
	gob.Register([]*reservoir.Reservoir{})
	gob.Register(drone.Drone{})
	gob.Register(physical_model.PhysicalModel{})
	gob.Register(map[string]relation_bonuses.RelationBonuses{})

	gob.Register(map[int][]*coordinate.Coordinate{})
	gob.Register([]*coordinate.Coordinate{})
	gob.Register([]coordinate.Coordinate{})
	gob.Register(coordinate.Coordinate{})
	gob.Register(game_math.Positions{})

	gob.Register(unit.EvacuationState{})

	gob.Register(map[int]*skin.Skin{})
	gob.Register(skin.Skin{})

	gob.Register([]db.DynObject{})
	gob.Register([]db.ReservoirOption{})

	gob.Register(map[string]map[int]*inventory.Slot{})
	gob.Register(map[int]*inventory.Slot{})
	gob.Register([]*inventory.Slot{})
	gob.Register(map[int]*detail.ThoriumSlot{})
	gob.Register([]*inventory.PlaceMayItems{})

	gob.Register(anomaly.Anomaly{})
	gob.Register(resource.Resource{})
	gob.Register(resource.RecycledResource{})
	gob.Register(resource.CraftDetail{})
	gob.Register(ammo.Ammo{})
	gob.Register(equip.Equip{})
	gob.Register([]*detail.BodyEquipSlot{})
	gob.Register(detail.Body{})
	gob.Register(box.Box{})
	gob.Register(blueprints.Blueprint{})
	gob.Register(trash_item.TrashItem{})
	gob.Register(product.Product{})
	gob.Register(inventory.ItemInfo{})
	gob.Register(inventory.Slot{})

	gob.Register([]unit.Decal{})
	gob.Register(unit.Decal{})

	gob.Register(generate_map_config.GenerateMapConfig{})
	gob.Register(builder_game_state.BuilderGameState{})
	gob.Register([]builder_game_state.BuilderGameState{})
	gob.Register(builder_game_state.UnitState{})

	gob.Register(dynamic_map_object.MinObject{})
	gob.Register([]drone.Job{})
	gob.Register(unit.BossConfig{})
	gob.Register([]effect.Effect{})
	gob.Register(visible_objects.VisibleObject{})

	gob.Register(map[string][]int{})
	gob.Register(map[int]bool{})
}
