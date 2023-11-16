package detail

import (
	"encoding/json"
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/ammo"
	"time"
)

type BodyWeaponSlot struct {
	ID                  int64      `json:"id"`
	AmmoSlotID          int64      `json:"ammo_slot_id"`
	Type                int        `json:"type_slot"` // по диз доку он может быть только 3
	Number              int        `json:"number_slot"`
	Weapon              *Weapon    `json:"weapon"`
	WeaponType          string     `json:"weapon_type"`
	Ammo                *ammo.Ammo `json:"ammo"`
	AmmoQuantity        int        `json:"ammo_quantity"`
	XAttach             int        `json:"x_attach"`
	YAttach             int        `json:"y_attach"`
	RealXAttach         int        `json:"real_x_attach"`
	RealYAttach         int        `json:"real_y_attach"`
	XAnchor             float64    `json:"x_anchor"`
	YAnchor             float64    `json:"y_anchor"`
	Reload              bool       `json:"reload"`
	AmmoReload          bool       `json:"-"`
	GunRotate           float64    `json:"gun_rotate"`
	CurrentReload       int        `json:"current_reload"`
	WeaponTexture       string     `json:"weapon_texture"`
	WeaponColor1        string     `json:"weapon_color_1"`
	WeaponColor2        string     `json:"weapon_color_2"`
	LockSlot            int        `json:"lock_slot"`
	lastFirePosition    int
	AccumulationCurrent float64 `json:"-"` // 0-100, от этого зависит урон
	AccumulationTimeOut int     `json:"-"`
	StartReloadTime     int64   `json:"-"`
	EndReloadTime       int64   `json:"-"`
}

func (s *BodyWeaponSlot) StartReload(reloadTime int, ammoReload bool) {
	s.SetReload(true)
	s.SetCurrentReload(reloadTime)
	s.AmmoReload = ammoReload

	s.StartReloadTime = time.Now().UnixNano() / int64(time.Millisecond)
	s.EndReloadTime = s.StartReloadTime + int64(reloadTime)
}

//func (s *BodyWeaponSlot) StartAmmoReload() {
//	s.SetReload(true)
//	s.setTimeReload(s.Weapon.ReloadAmmoTime)
//	s.SetCurrentReload(s.Weapon.ReloadAmmoTime)
//}

func (s *BodyWeaponSlot) GetReload() bool {
	return s.Reload
}

func (s *BodyWeaponSlot) SetReload(reload bool) {
	s.Reload = reload
	if !reload {
		s.AmmoReload = false
		s.StartReloadTime = 0
		s.EndReloadTime = 0
	}
}

func (s *BodyWeaponSlot) GetAmmoQuantity() int {
	return s.AmmoQuantity
}

func (s *BodyWeaponSlot) SetAmmoQuantity(quantity int) {
	s.AmmoQuantity = quantity
}

func (s *BodyWeaponSlot) SetAnchor() {
	if s == nil || s.Weapon == nil {
		return
	}

	s.XAnchor, s.YAnchor, s.RealXAttach, s.RealYAttach = game_math.GetAnchorWeapon(s.Weapon.XAttach, s.Weapon.YAttach, s.XAttach, s.YAttach)
}

func (s *BodyWeaponSlot) GetAmmo() *ammo.Ammo {
	return s.Ammo
}

func (s *BodyWeaponSlot) SetAmmo(ammo *ammo.Ammo) {
	s.Ammo = ammo
}

func (s *BodyWeaponSlot) GetGunRotate() float64 {
	return s.GunRotate
}

func (s *BodyWeaponSlot) SetGunRotate(rotate float64) {
	s.GunRotate = rotate
}

func (s *BodyWeaponSlot) GetCurrentReload() int {
	return s.CurrentReload
}

func (s *BodyWeaponSlot) SetCurrentReload(time int) {
	if time <= 0 {
		s.SetReload(false)
	}
	s.CurrentReload = time
}

func (s *BodyWeaponSlot) GetXAnchor() float64 {
	return s.XAnchor
}

func (s *BodyWeaponSlot) GetYAnchor() float64 {
	return s.YAnchor
}

func (s *BodyWeaponSlot) GetRealXAttach() int {
	return s.RealXAttach
}

func (s *BodyWeaponSlot) GetRealYAttach() int {
	return s.RealYAttach
}

func (s *BodyWeaponSlot) GetJSON() string {
	jsonSlot, err := json.Marshal(s)
	if err != nil {
		println("weapon Slot to json: ", err.Error())
	}

	return string(jsonSlot)
}

func (s *BodyWeaponSlot) GetCopy() *BodyWeaponSlot {
	copySlot := *s
	return &copySlot
}

func (s *BodyWeaponSlot) NextLastFirePosition() {
	if s.Weapon == nil {
		return
	}

	s.lastFirePosition++
	if len(s.Weapon.FirePositions) == s.lastFirePosition {
		s.lastFirePosition = 0
	}
}

func (s *BodyWeaponSlot) GetLastFirePosition() int {
	return s.lastFirePosition
}

func (w *Weapon) GetWeaponMaxRange(ammo *ammo.Ammo, lvlMap, mapHeight float64, realBallistic bool) (int, float64) {
	if (w.Type == "laser" || w.Type == "missile") || (w.MaxRange > 0 && !realBallistic) {
		return w.MaxRange, float64(w.MaxAngle)
	}

	bulletSpeed := ammo.BulletSpeed + w.BulletSpeed

	angle := 0.0

	if w.Artillery {
		angle = float64(w.MinAngle)
	} else {
		angle = float64(w.MaxAngle)
	}

	maxRange := game_math.GetMaxDistBallisticWeapon(
		float64(bulletSpeed),
		lvlMap,
		lvlMap+mapHeight,
		game_math.DegToRadian(angle),
		w.Type,
		ammo.Gravity,
	)

	return int(maxRange), angle
}
