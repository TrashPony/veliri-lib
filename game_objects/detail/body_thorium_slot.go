package detail

import (
	"encoding/json"
	"github.com/TrashPony/veliri-lib/game_objects/fuel"
)

type ThoriumSlot struct {
	ID                int64     `json:"id"`
	Number            int       `json:"number_slot"`
	Count             int       `json:"count"`
	MaxCount          int       `json:"max_count"`
	WorkedOut         int       `json:"worked_out"`
	Inversion         bool      `json:"inversion"`
	ProcessingThorium int       `json:"processing_thorium"`
	Fuel              fuel.Fuel `json:"fuel"`
}

func (t *ThoriumSlot) GetCount() int {
	return t.Count
}

func (t *ThoriumSlot) SetCount(count int, fuel fuel.Fuel) {
	t.Count = count
	t.Fuel = fuel
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
