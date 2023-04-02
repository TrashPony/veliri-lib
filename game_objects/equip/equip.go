package equip

import (
	"github.com/TrashPony/veliri-lib/game_objects/coordinate"
	"github.com/TrashPony/veliri-lib/game_objects/effect"
	"github.com/TrashPony/veliri-lib/game_objects/reservoir"
)

type Equip struct {
	ID              int                      `json:"id"`
	Name            string                   `json:"name"`
	Active          bool                     `json:"active"`
	Applicable      string                   `json:"applicable"`
	Region          int                      `json:"region"`
	Radius          int                      `json:"radius"`
	TypeSlot        int                      `json:"type_slot"`
	Reload          int                      `json:"reload"`
	Power           int                      `json:"power"`
	UsePower        int                      `json:"use_power"`
	Recovery        int                      `json:"recovery"`
	RecoveryCycle   int                      `json:"recovery_cycle"`
	MaxHP           int                      `json:"max_hp"`
	Size            int                      `json:"size"`
	MiningChecker   bool                     `json:"move_checker"`
	MiningReservoir *reservoir.Reservoir     `json:"mining_reservoir"`
	Count           int                      `json:"count"`
	Count2          int                      `json:"count_2"`
	MaxDiffAngle    int                      `json:"max_diff_angle"`
	XAttach         int                      `json:"x_attach"`
	YAttach         int                      `json:"y_attach"`
	Unit            bool                     `json:"unit"`
	Structure       bool                     `json:"structure"`
	Effects         []*effect.Effect         `json:"effects"`
	FirePositions   []*coordinate.Coordinate `json:"fire_positions"`
	RotateSpeed     int                      `json:"rotate_speed"`
	PassAngle       int                      `json:"pass_angle"`

	MaxAmmoCount int            `json:"max_ammo_count"` // 0 - бесконечно
	LifeTime     int            `json:"life_time"`      // время жизни обьекта или дрона
	ObjectID     int            `json:"object_id"`      // ид обьекта которые вызывает снаряжения (турель/стена)
	WeaponID     int            `json:"weapon_id"`
	Attributes   map[string]int `json:"attributes"`

	miningExit chan bool
}

func (e *Equip) GetName() string {
	return e.Name
}

func (e *Equip) GetItemType() string {
	return ""
}

func (e *Equip) GetItemName() string {
	return ""
}

func (e *Equip) GetInMap() bool {
	return false
}

func (e *Equip) GetIcon() string {
	return ""
}

func (e *Equip) GetStandardSize() int {
	return 0
}

func (e *Equip) GetTypeSlot() int {
	return e.TypeSlot
}

func (e *Equip) CreateMining() {
	if e.miningExit != nil {
		close(e.miningExit)
	}
	e.miningExit = make(chan bool)
}

func (e *Equip) GetMining() chan bool {
	return e.miningExit
}
