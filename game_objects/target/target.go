package target

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_objects/inventory"
	"time"
)

type Target struct {
	Type         string            `json:"type"` // box, unit, map
	UUID         string            `json:"uuid"`
	ID           int               `json:"id"`
	X            int               `json:"x"`
	ParrentX     int               `json:"parrent_x"`
	Y            int               `json:"y"`
	ParrentY     int               `json:"parrent_y"`
	Z            float64           `json:"z"`
	Follow       bool              `json:"follow"` // преследовать цель используется для цели атак
	Attack       bool              `json:"attack"`
	Force        bool              `json:"force"`
	WeaponTarget bool              `json:"weapon_target"`
	Radius       int               `json:"radius"`      // радиус на котором держатся от цели
	OpenBox      bool              `json:"open_box"`    // если true и тип бокс, то он попытается открыть его как достигнет дистации меньше радиус
	OpenObject   bool              `json:"open_object"` // взаимодействовать с обьектом при достижение Radius
	MiningOre    bool              `json:"mining_ore"`  // начинать майнить по достижение
	ToBase       bool              `json:"to_base"`     // информационное поле о том что он идет в базу
	PlaceObject  bool              `json:"place_object"`
	Rotate       float64           `json:"rotate"`
	BaseID       int               `json:"base_id"`
	Ignore       bool              `json:"ignore"`
	Update       int64             `json:"-"`
	MouseTap     bool              `json:"-"`
	StartTap     bool              `json:"-"`
	LastTap      bool              `json:"-"`
	Points       int               `json:"-"`
	PassAngle    int               `json:"-"`
	Key          string            `json:"-"`
	Slots        []*inventory.Slot `json:"-"`
}

func (t *Target) GetX() int {
	return t.X
}

func (t *Target) SetX(x int) {
	t.X = x
}

func (t *Target) GetY() int {
	return t.Y
}

func (t *Target) SetY(y int) {
	t.Y = y
}

func (t *Target) GetZ() float64 {
	return t.Z
}

func (t *Target) SetZ(z float64) {
	t.Z = z
}

func (t *Target) GetFollow() bool {
	return t.Follow
}

func (t *Target) SetFollow(follow bool) {
	t.Follow = follow
}

func (t *Target) GetUpdate() int64 {
	return t.Update
}

func (t *Target) SetUpdate() {
	t.Update = time.Now().UnixNano()
}

func (t *Target) GetCopy() *Target {
	copyTarget := *t
	return &copyTarget
}

// todo костыль
func (t *Target) GetNumberType() int {
	if t.Type == "unit" {
		return _const.ObjTypeUnit
	}

	if t.Type == "box" {
		return _const.ObjTypeBox
	}

	if t.Type == "object" {
		return _const.ObjTypeDynamicObject
	}

	return 0
}
