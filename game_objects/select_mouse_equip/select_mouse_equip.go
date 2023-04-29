package select_mouse_equip

type SelectMouseEquip struct {
	Type      string `json:"type"`
	TypeEquip int    `json:"type_equip"`
	SlotEquip int    `json:"slot_equip"`
	PanelSlot int    `json:"panel_slot"`
}
