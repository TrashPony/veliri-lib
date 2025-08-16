package damage_object

import (
	"github.com/TrashPony/veliri-lib/game_objects/unit"
	"github.com/TrashPony/veliri-lib/game_objects/web_socket_response"
)

type Object struct {
	TypeTarget        string                          `json:"type_target"`
	IdTarget          int                             `json:"id_target"`
	TargetOwnerID     int                             `json:"target_owner_id"`
	Damage            int                             `json:"damage"`
	DamageK           int                             `json:"damage_k"`
	KineticsDamage    int                             `json:"kinetics_damage"`
	ExplosionDamage   int                             `json:"explosion_damage"`
	ThermoDamage      int                             `json:"thermo_damage"`
	WeaponType        string                          `json:"weapon_type"`
	Dead              bool                            `json:"dead"`
	Obj               interface{}                     `json:"obj"`
	X                 int                             `json:"x"`
	Y                 int                             `json:"y"`
	PushPower         int                             `json:"push_power"`
	DealerType        string                          `json:"dealer_type"`
	DealerID          int                             `json:"dealer_id"`
	Explosion         bool                            `json:"explosion"`
	PointResponses    []*web_socket_response.Response `json:"-"`
	KillResponses     []*web_socket_response.Response `json:"-"`
	Decal             []*unit.Decal                   `json:"-"`
	Area              int                             `json:"-"`
	UnitShield        bool                            `json:"-"`
	DestroyUnitShield bool                            `json:"-"`
	Type              byte                            `json:"-"`
}
