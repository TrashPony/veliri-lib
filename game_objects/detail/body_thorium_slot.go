package detail

import (
	"encoding/json"
	"sync"
)

type ThoriumSlot struct {
	Number            int  `json:"number_slot"`
	Count             int  `json:"count"`
	MaxCount          int  `json:"max_count"`
	WorkedOut         int  `json:"worked_out"` /* параметр показывает что топливо вырабатано на сколь-ко то процентов */
	Inversion         bool `json:"inversion"`
	ProcessingThorium int  `json:"processing_thorium"`
	mx                sync.RWMutex
}

func (t *ThoriumSlot) GetCount() int {
	t.mx.RLock()
	defer t.mx.RUnlock()

	return t.Count
}

func (t *ThoriumSlot) SetCount(count int) {
	t.mx.Lock()
	defer t.mx.Unlock()

	t.Count = count
}

func (t *ThoriumSlot) GetJSON() string {
	t.mx.Lock()
	defer t.mx.Unlock()

	jsonSlot, err := json.Marshal(t)
	if err != nil {
		println("equip Slot to json: ", err.Error())
	}

	return string(jsonSlot)
}

func (t *ThoriumSlot) GetCopy() *ThoriumSlot {
	t.mx.Lock()
	copy := *t
	t.mx.Unlock()
	copy.mx = sync.RWMutex{}
	return &copy
}
