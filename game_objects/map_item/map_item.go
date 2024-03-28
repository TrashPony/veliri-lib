package map_item

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/gunner"
	"github.com/TrashPony/veliri-lib/game_objects/inventory"
	"github.com/TrashPony/veliri-lib/game_objects/move_path"
	"github.com/TrashPony/veliri-lib/game_objects/physical_model"
)

type MapItem struct {
	ID             int          `json:"id"`
	ItemID         int          `json:"item_id"`
	ItemType       string       `json:"item_type"`
	MapID          int          `json:"map_id"`
	CacheJson      []byte       `json:"-"`
	CreateJsonTime int64        `json:"-"`
	LifeTime       int64        `json:"-"`
	ToPath         move_path.To `json:"-"`
	inv            *inventory.Inventory
	physicalModel  *physical_model.PhysicalModel
}

func (i *MapItem) GetGunner() *gunner.Gunner {
	return nil
}

func (i *MapItem) GetX() int {
	return i.GetPhysicalModel().GetX()
}

func (i *MapItem) GetY() int {
	return i.GetPhysicalModel().GetY()
}

func (i *MapItem) GetID() int {
	return i.ID
}

func (i *MapItem) SetSlot(slot *inventory.Slot) {
	inv := &inventory.Inventory{}
	inv.Init("", 0, nil)
	slot.Number = 1
	inv.AddSlot(slot.Number, slot)
	i.inv = inv
}

func (i *MapItem) GetInventory() *inventory.Inventory {
	return i.inv
}

func (i *MapItem) GetSlot() *inventory.Slot {
	if i.inv == nil {
		return nil
	}

	s, _ := i.inv.GetSlot(1, -1)
	return s
}

func (i *MapItem) GetPhysicalModel() *physical_model.PhysicalModel {

	if i.physicalModel == nil {
		i.initPhysicalModel()
	}

	return i.physicalModel
}

func (i *MapItem) LockedControl() bool {
	return false
}

func (i *MapItem) Ghost() bool {
	return false
}

func (i *MapItem) initPhysicalModel() {
	i.physicalModel = &physical_model.PhysicalModel{
		MoveDrag:    0.60,
		AngularDrag: 0.60,
		Weight:      3000.0,
		Type:        "map_item",
		ID:          i.ID,
		Height:      _const.ItemSizeRadius,
		Radius:      _const.ItemSizeRadius,
	}
}

func (i *MapItem) SetAllGunRotate(float64) {

}

func (i *MapItem) GetJSON(mapTime int64) []byte {

	if i.CreateJsonTime == mapTime && len(i.CacheJson) > 0 {
		return i.CacheJson
	}

	if i.inv == nil {
		return []byte{}
	}

	s, _ := i.inv.GetSlot(1, -1)
	if s == nil {
		return []byte{}
	}

	if i.CacheJson == nil {
		i.CacheJson = []byte{}
	}
	i.CacheJson = i.CacheJson[:0]

	i.CacheJson = append(i.CacheJson, game_math.GetIntBytes(i.ID)...)
	i.CacheJson = append(i.CacheJson, game_math.GetIntBytes(i.GetX())...)
	i.CacheJson = append(i.CacheJson, game_math.GetIntBytes(i.GetY())...)
	i.CacheJson = append(i.CacheJson, game_math.GetIntBytes(int(i.GetPhysicalModel().GetRotate()))...)
	i.CacheJson = append(i.CacheJson, game_math.GetIntBytes(i.MapID)...)
	i.CacheJson = append(i.CacheJson, game_math.GetIntBytes(s.Quantity)...)
	i.CacheJson = append(i.CacheJson, game_math.GetIntBytes(s.ItemID)...)
	i.CacheJson = append(i.CacheJson, byte(_const.ItemBinTypes[s.Type]))
	i.CacheJson = append(i.CacheJson, byte(i.GetPhysicalModel().GetRadius()))

	i.CreateJsonTime = mapTime

	return i.CacheJson
}

func (i *MapItem) GetUpdateData(mapTime int64) []byte {
	command := make([]byte, 4)

	if i.inv == nil {
		return []byte{}
	}

	s, _ := i.inv.GetSlot(1, -1)
	if s == nil {
		return []byte{}
	}

	game_math.ReuseByteSlice(&command, 0, game_math.GetIntBytes(s.GetQuantity()))
	return command
}
