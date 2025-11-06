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
	ID              int          `json:"id"`
	ItemID          int          `json:"item_id"`
	ItemType        string       `json:"item_type"`
	MapID           int          `json:"map_id"`
	CacheCreateData CacheData    `json:"-"`
	CacheUpdateData CacheData    `json:"-"`
	LifeTime        int64        `json:"-"`
	ToPath          move_path.To `json:"-"`
	LastAction      string       `json:"-"`
	Destroy         bool         `json:"-"`
	ownerType       string
	ownerID         int
	ownerInv        int
	inv             *inventory.Inventory
	physicalModel   *physical_model.PhysicalModel
}

type CacheData struct {
	Data []byte `json:"-"`
	Time int64  `json:"-"`
}

func (i *MapItem) SetOwner(ownerType string, ownerID, ownerInv int, updater func(string, int, int)) {
	i.ownerType = ownerType
	i.ownerID = ownerID
	i.ownerInv = ownerInv

	if updater != nil {
		go updater(i.ownerType, i.ownerID, i.ownerInv)
	}
}

func (i *MapItem) GetOwnerType() string {
	return i.ownerType
}

func (i *MapItem) GetOwnerID() int {
	return i.ownerID
}

func (i *MapItem) GetOwnerInv() int {
	return i.ownerInv
}

func (i *MapItem) GetGunner() *gunner.Gunner {
	return nil
}

func (i *MapItem) GetMeleer() *gunner.Meleer {
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
	inv.Init("", 0, nil, nil)
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

func (i *MapItem) GetQuantity() int {
	return i.GetSlot().GetQuantity()
}

func (i *MapItem) SetQuantity(q int, updater func(int)) {
	i.GetSlot().SetQuantity(q)
	if updater != nil {
		go updater(q)
	}
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

	if i.CacheCreateData.Time == mapTime && len(i.CacheCreateData.Data) > 0 {
		return i.CacheCreateData.Data
	}

	if i.inv == nil {
		return []byte{}
	}

	s, _ := i.inv.GetSlot(1, -1)
	if s == nil {
		return []byte{}
	}

	if i.CacheCreateData.Data == nil {
		i.CacheCreateData.Data = []byte{}
	}
	i.CacheCreateData.Data = i.CacheCreateData.Data[:0]

	i.CacheCreateData.Data = append(i.CacheCreateData.Data, game_math.GetIntBytes(i.ID)...)
	i.CacheCreateData.Data = append(i.CacheCreateData.Data, game_math.GetIntBytes(i.GetX())...)
	i.CacheCreateData.Data = append(i.CacheCreateData.Data, game_math.GetIntBytes(i.GetY())...)
	i.CacheCreateData.Data = append(i.CacheCreateData.Data, game_math.GetIntBytes(int(i.GetPhysicalModel().GetRotate()))...)
	i.CacheCreateData.Data = append(i.CacheCreateData.Data, game_math.GetIntBytes(i.MapID)...)
	i.CacheCreateData.Data = append(i.CacheCreateData.Data, game_math.GetIntBytes(s.Quantity)...)
	i.CacheCreateData.Data = append(i.CacheCreateData.Data, game_math.GetIntBytes(s.ItemID)...)
	i.CacheCreateData.Data = append(i.CacheCreateData.Data, byte(_const.ItemBinTypes[s.Type]))
	i.CacheCreateData.Data = append(i.CacheCreateData.Data, byte(i.GetPhysicalModel().GetRadius()))

	i.CacheCreateData.Time = mapTime

	return i.CacheCreateData.Data
}

func (i *MapItem) GetUpdateData(mapTime int64) []byte {
	if i.CacheUpdateData.Time == mapTime && len(i.CacheUpdateData.Data) > 0 {
		return i.CacheUpdateData.Data
	}

	if i.inv == nil {
		return []byte{}
	}

	s, _ := i.inv.GetSlot(1, -1)
	if s == nil {
		return []byte{}
	}

	if i.CacheUpdateData.Data == nil {
		i.CacheUpdateData.Data = []byte{}
	}

	i.CacheUpdateData.Data = i.CacheUpdateData.Data[:0]

	i.CacheUpdateData.Data = append(i.CacheUpdateData.Data, game_math.GetIntBytes(s.GetQuantity())...)

	i.CacheUpdateData.Time = mapTime

	return i.CacheUpdateData.Data
}
