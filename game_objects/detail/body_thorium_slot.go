package detail

import (
	"encoding/json"
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_objects/fuel"
)

type ThoriumSlot struct {
	ID          int64     `json:"id"`
	Number      int       `json:"number_slot"`
	Worked      int       `json:"worked_out"`
	CurrentFuel fuel.Fuel `json:"fuel"`
	NextFuel    fuel.Fuel `json:"next_fuel"`
	Durability  int       `json:"durability"`
	MaxCap      int       `json:"max_cap"`
	Reload      bool      `json:"-"`
	SendRequest bool      `json:"-"`
}

func (t *ThoriumSlot) SetFuel(fuel fuel.Fuel) {
	t.CurrentFuel = fuel
	t.Worked = fuel.EnergyCap
	t.Durability = _const.StartDurability
}

func (t *ThoriumSlot) GetJSON() string {
	jsonSlot, err := json.Marshal(t)
	if err != nil {
		println("equip Slot to json: ", err.Error())
	}

	return string(jsonSlot)
}

func (t *ThoriumSlot) GetMaxCap() int {
	x := float64(t.Durability) / _const.StartDurability // 0.0-1.0

	// Квадратичная кривая
	effectivePercent := 0.25*x*x + 0.75 // От 75% до 100%

	t.MaxCap = int(float64(t.CurrentFuel.EnergyCap) * effectivePercent)
	return t.MaxCap
}

func (t *ThoriumSlot) GetCopy() *ThoriumSlot {
	copy := *t
	return &copy
}
