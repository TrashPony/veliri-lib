package dynamic_map_object

import "github.com/TrashPony/veliri-lib/game_objects/coordinate"

type MinObject struct {
	ID               int                      `json:"id"`
	TypeID           int                      `json:"type_id"`
	X                int                      `json:"x"`
	Y                int                      `json:"y"`
	MapID            int                      `json:"map_id"`
	BoxID            int                      `json:"box_id"`
	Immortal         bool                     `json:"immortal"`
	Complete         float64                  `json:"complete"`
	CorporationID    int                      `json:"corporation_id"`
	HP               int                      `json:"hp"`
	MaxHP            int                      `json:"max_hp"`         // o.GetMaxHP()
	BaseData         []*coordinate.Coordinate `json:"type_base_data"` // o.GetBaseData()
	BodyID           int                      `json:"body_id"`
	Fraction         string                   `json:"fraction"`
	SpecialCell      bool                     `json:"special_cell"`
	Interactive      bool                     `json:"interactive"`
	Inventory        bool                     `json:"inventory"`
	Disable          bool                     `json:"disable"`
	AutoBuild        bool                     `json:"auto_build"`
	Texture          string                   `json:"texture"`
	Build            bool                     `json:"build"`
	Scale            int                      `json:"scale"`
	GrowTime         int                      `json:"grow_time"`
	GrowLeftTime     int                      `json:"grow_left_time"`
	EquipID          int                      `json:"equip_id"`
	CurrentEnergy    int                      `json:"current_energy"`
	OwnerID          int                      `json:"owner_id"`
	OwnerPlayerID    int                      `json:"owner_player_id"`
	OwnerType        string                   `json:"owner_type"`
	OwnerBaseID      int                      `json:"owner_base_id"`
	MaxScale         int                      `json:"max_scale"`
	LifeTime         int                      `json:"life_time"`
	DestroyLeftTimer bool                     `json:"destroy_left_timer"`
	DestroyTimer     int                      `json:"destroy_timer"`
}

func (o *Object) GetMinData() MinObject {
	d := MinObject{
		ID:               o.ID,
		TypeID:           o.TypeID,
		X:                o.GetX(),
		Y:                o.GetY(),
		MapID:            o.MapID,
		BoxID:            o.BoxID,
		Immortal:         o.Immortal,
		Complete:         o.Complete,
		CorporationID:    o.CorporationID,
		HP:               o.HP,
		MaxHP:            o.GetMaxHP(),
		BaseData:         o.GetBaseData(),
		BodyID:           o.BodyID,
		Fraction:         o.Fraction,
		SpecialCell:      o.SpecialCell,
		Interactive:      o.Interactive,
		Inventory:        o.Inventory,
		Disable:          o.Disable,
		AutoBuild:        o.AutoBuild,
		Texture:          o.Texture,
		Build:            o.Build,
		Scale:            o.Scale,
		GrowTime:         o.GrowTime,
		GrowLeftTime:     o.GrowLeftTime,
		CurrentEnergy:    o.CurrentEnergy,
		OwnerID:          o.OwnerID,
		OwnerPlayerID:    o.OwnerPlayerID,
		OwnerType:        o.OwnerType,
		OwnerBaseID:      o.OwnerBaseID,
		MaxScale:         o.MaxScale,
		LifeTime:         o.LifeTime,
		DestroyLeftTimer: o.DestroyLeftTimer,
		DestroyTimer:     o.DestroyTimer,
	}

	for _, e := range o.Equips {
		if e != nil && e.Equip != nil {
			d.EquipID = e.Equip.ID
			break
		}
	}

	return d
}
