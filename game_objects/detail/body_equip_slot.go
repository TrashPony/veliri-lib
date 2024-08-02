package detail

import (
	"encoding/json"
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/equip"
	"github.com/TrashPony/veliri-lib/game_objects/target"
	"time"
)

type BodyEquipSlot struct {
	ID               int64          `json:"id"`
	Type             int            `json:"type_slot"`
	UUID             string         `json:"uuid"`
	Number           int            `json:"number_slot"`
	Equip            *equip.Equip   `json:"equip"`
	Target           *target.Target `json:"-"`
	StandardSize     int            `json:"standard_size"` /* определяет тип вмещаемого юнита если это ангар */
	XAttach          int            `json:"x_attach"`
	YAttach          int            `json:"y_attach"`
	RealXAttach      int            `json:"real_x_attach"`
	RealYAttach      int            `json:"real_y_attach"`
	XAnchor          float64        `json:"x_anchor"`
	YAnchor          float64        `json:"y_anchor"`
	Reload           bool           `json:"-"`
	Mode             string         `json:"-"`
	CurrentReload    int            `json:"-"`
	CurrentAmmoCount int            `json:"-"`
	On               bool           `json:"on"`
	Drone            bool           `json:"-"` // поле существует только что бы рисовать на фронте
	LockSlot         int            `json:"lock_slot"`
	Rotate           float64        `json:"rotate"`

	Mining          bool `json:"mining"` //todo переименовать в "наружний"
	CatchBulletUUID string
	StartReloadTime int64 `json:"-"`
	EndReloadTime   int64 `json:"-"`

	TypeSlot int `json:"-"`
	Slot     int `json:"-"`

	Direction int                    `json:"direction"`
	PassAngle int                    `json:"pass_angle"`
	Attr      map[string]interface{} `json:"attr"`
}

func (s *BodyEquipSlot) GetReload() bool {
	return s.Reload
}

func (s *BodyEquipSlot) SetReload(reload bool) {
	s.Reload = reload
	if !reload {
		s.StartReloadTime = 0
		s.EndReloadTime = 0
	}
}

func (s *BodyEquipSlot) StartReload() {
	s.SetReload(true)
	s.SetCurrentReload(s.Equip.Reload)
	s.StartReloadTime = time.Now().UnixNano() / int64(time.Millisecond)
	s.EndReloadTime = s.StartReloadTime + int64(s.Equip.Reload)
}

func (s *BodyEquipSlot) SetAnchor() {
	if s == nil || s.Equip == nil {
		return
	}

	s.XAnchor, s.YAnchor, s.RealXAttach, s.RealYAttach = game_math.GetAnchorWeapon(s.Equip.XAttach, s.Equip.YAttach, s.XAttach, s.YAttach)
}

func (s *BodyEquipSlot) GetCurrentReload() int {
	return s.CurrentReload
}

func (s *BodyEquipSlot) SetCurrentReload(time int) {

	if time <= 0 {
		s.SetReload(false)
	}

	s.CurrentReload = time
}

func (s *BodyEquipSlot) GetXAnchor() float64 {
	return s.XAnchor
}

func (s *BodyEquipSlot) GetYAnchor() float64 {
	return s.YAnchor
}

func (s *BodyEquipSlot) GetRealXAttach() int {
	return s.RealXAttach
}

func (s *BodyEquipSlot) GetRealYAttach() int {
	return s.RealYAttach
}

func (s *BodyEquipSlot) GetJSON() string {
	jsonSlot, err := json.Marshal(s)
	if err != nil {
		println("equip Slot to json: ", err.Error())
	}

	return string(jsonSlot)
}

func (s *BodyEquipSlot) GetCopy() *BodyEquipSlot {
	copySlot := *s
	return &copySlot
}
