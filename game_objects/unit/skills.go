package unit

import "strings"

var Size = map[int]string{1: "small", 2: "medium", 3: "big"}
var WType = map[string]string{"laser": "laser", "missile": "rocket", "firearms": "ballistic"}

func (unit *Unit) GetAmmoSkillName() string {
	weaponSlot := unit.GetWeaponSlot(1)
	if weaponSlot == nil || weaponSlot.Ammo == nil {
		return ""
	}

	return Size[weaponSlot.Ammo.StandardSize] + "_" + WType[weaponSlot.Ammo.Type] + "_ammo"
}

func (unit *Unit) GetWeaponSkillName() string {
	weaponSlot := unit.GetWeaponSlot(1)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return ""
	}

	return Size[weaponSlot.Weapon.StandardSize] + "_" + WType[weaponSlot.Weapon.Type] + "_weapon"
}

func (unit *Unit) GetMoveBodySkillName() string {
	return "move_" + Size[unit.getBody().StandardSize] + "_" + strings.ToLower(unit.getBody().Fraction) + "_body"
}

func (unit *Unit) GetArmorBodySkillName() string {
	return "armor_" + Size[unit.getBody().StandardSize] + "_" + strings.ToLower(unit.getBody().Fraction) + "_body"
}

type parameterVolume struct {
	All        float64 `json:"all"`
	EquipBonus float64 `json:"equip_bonus"`
	SkillBonus float64 `json:"skill_bonus"`
	BodyBonus  float64 `json:"body_bonus"`
	// TODO EffectBonus
}
