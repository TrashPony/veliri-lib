package unit

type AvailableAmmoMsg struct {
	WeaponSlot int         `json:"weapon_slot"`
	Ammo       map[int]int `json:"ammo"` // [id]count
}
