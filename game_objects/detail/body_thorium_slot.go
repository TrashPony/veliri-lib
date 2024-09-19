package detail

import (
	"encoding/json"
	"github.com/TrashPony/veliri-lib/game_objects/fuel"
)

type ThoriumSlot struct {
	ID          int64     `json:"id"`
	Number      int       `json:"number_slot"`
	Worked      int       `json:"worked_out"`
	CurrentFuel fuel.Fuel `json:"fuel"`
	NextFuel    fuel.Fuel `json:"next_fuel"`
	Reload      bool      `json:"-"`
	SendRequest bool      `json:"-"`
}

func (t *ThoriumSlot) SetFuel(fuel fuel.Fuel) {
	t.CurrentFuel = fuel
	t.Worked = fuel.EnergyCap
}

func (t *ThoriumSlot) GetJSON() string {
	jsonSlot, err := json.Marshal(t)
	if err != nil {
		println("equip Slot to json: ", err.Error())
	}

	return string(jsonSlot)
}

func (t *ThoriumSlot) GetCopy() *ThoriumSlot {
	copy := *t
	return &copy
}
