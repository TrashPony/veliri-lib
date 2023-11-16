package detail

import (
	"encoding/json"
)

type ThoriumSlot struct {
	ID                int64 `json:"id"`
	Number            int   `json:"number_slot"`
	Count             int   `json:"count"`
	MaxCount          int   `json:"max_count"`
	WorkedOut         int   `json:"worked_out"` /* параметр показывает что топливо вырабатано на сколь-ко то процентов */
	Inversion         bool  `json:"inversion"`
	ProcessingThorium int   `json:"processing_thorium"`
}

func (t *ThoriumSlot) GetCount() int {
	return t.Count
}

func (t *ThoriumSlot) SetCount(count int) {
	t.Count = count
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
