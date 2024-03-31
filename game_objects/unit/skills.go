package unit

import "strings"

var Size = map[int]string{1: "small", 2: "medium", 3: "big"}
var WType = map[string]string{"laser": "laser", "missile": "rocket", "firearms": "ballistic"}

func (u *Unit) GetAmmoSkillName() string {
	weaponSlot := u.GetWeaponSlot(1)
	if weaponSlot == nil || weaponSlot.Ammo == nil {
		return ""
	}

	return Size[weaponSlot.Ammo.StandardSize] + "_" + WType[weaponSlot.Ammo.Type] + "_ammo"
}

func (u *Unit) GetRealoadWeaponSkillName() string {
	weaponSlot := u.GetWeaponSlot(1)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return ""
	}

	return "" + Size[weaponSlot.Weapon.StandardSize] + "_" + WType[weaponSlot.Weapon.Type] + "_weapon"
}

func (u *Unit) GetRotateWeaponSkillName() string {
	weaponSlot := u.GetWeaponSlot(1)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return ""
	}

	return "rotate_" + Size[weaponSlot.Weapon.StandardSize] + "_" + WType[weaponSlot.Weapon.Type] + "_weapon"
}

func (u *Unit) GetAccuracyWeaponSkillName() string {
	weaponSlot := u.GetWeaponSlot(1)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return ""
	}

	return "accuracy_" + Size[weaponSlot.Weapon.StandardSize] + "_" + WType[weaponSlot.Weapon.Type] + "_weapon"
}

func (u *Unit) GetRateWeaponSkillName() string {
	weaponSlot := u.GetWeaponSlot(1)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return ""
	}

	return "rate_" + Size[weaponSlot.Weapon.StandardSize] + "_" + WType[weaponSlot.Weapon.Type] + "_weapon"
}

func (u *Unit) GetMoveBodySkillName() string {
	return "move_" + Size[u.getBody().StandardSize] + "_" + strings.ToLower(u.getBody().Fraction) + "_body"
}

func (u *Unit) GetArmorBodySkillName() string {
	return "armor_" + Size[u.getBody().StandardSize] + "_" + strings.ToLower(u.getBody().Fraction) + "_body"
}

func (u *Unit) GetViewBodySkillName() string {
	return "view_" + Size[u.getBody().StandardSize] + "_" + strings.ToLower(u.getBody().Fraction) + "_body"
}

func (u *Unit) GetPowerCapBodySkillName() string {
	return "powercap_" + Size[u.getBody().StandardSize] + "_" + strings.ToLower(u.getBody().Fraction) + "_body"
}

func (u *Unit) GetRadarBodySkillName() string {
	return "radar_" + Size[u.getBody().StandardSize] + "_" + strings.ToLower(u.getBody().Fraction) + "_body"
}

func (u *Unit) GetChargingBodySkillName() string {
	return "charging_" + Size[u.getBody().StandardSize] + "_" + strings.ToLower(u.getBody().Fraction) + "_body"
}

func (u *Unit) GetSpeedBodySkillName() string {
	return "speed_" + Size[u.getBody().StandardSize] + "_" + strings.ToLower(u.getBody().Fraction) + "_body"
}

type parameterVolume struct {
	All        float64 `json:"all"`
	EquipBonus float64 `json:"equip_bonus"`
	SkillBonus float64 `json:"skill_bonus"`
	BodyBonus  float64 `json:"body_bonus"`
	// TODO EffectBonus
}
