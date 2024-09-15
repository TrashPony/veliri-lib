package fuel

import "github.com/TrashPony/veliri-lib/game_objects/effect"

type Fuel struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Size      int             `json:"size"`
	EnergyCap int             `json:"energy_cap"`
	Bonuses   []effect.Effect `json:"bonuses"`
}

func (f *Fuel) GetName() string {
	return f.Name
}

func (f *Fuel) GetSize() int {
	return f.Size
}

func (f *Fuel) GetItemType() string {
	return ""
}

func (f *Fuel) GetItemName() string {
	return ""
}

func (f *Fuel) GetInMap() bool {
	return false
}

func (f *Fuel) GetIcon() string {
	return ""
}

func (f *Fuel) GetStandardSize() int {
	return 0
}

func (f *Fuel) GetTypeSlot() int {
	return 0
}
