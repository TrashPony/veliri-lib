package unit

import (
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"github.com/TrashPony/veliri-lib/game_objects/obstacle_point"
	"math"
)

func (u *Unit) GetMeleeWeaponSlot(slotNumber int) *detail.BodyWeaponSlot { // по диз доку оружие в юните может быть только одно

	if u == nil || u.getBody() == nil {
		return nil
	}

	return u.getBody().MeleeWeapons[slotNumber]
}

func (u *Unit) RangeMeleeWeaponSlots() map[int]*detail.BodyWeaponSlot {
	// мы никогда не пишут в карту слотов оружия поэтому этот метод безопасен (по крайне мере пока)
	return u.getBody().MeleeWeapons
}

func (u *Unit) GetMeleeWeaponFirePos(slotNumber int) []*game_math.Positions {

	weaponSlot := u.GetMeleeWeaponSlot(slotNumber)

	return game_math.GetWeaponFirePositions(
		u.GetX(), u.GetY(), u.GetScale(), u.GetRotate(), weaponSlot.GetGunRotate(),
		weaponSlot.Weapon.XAttach, weaponSlot.Weapon.YAttach,
		weaponSlot.Weapon.FirePositions,
		float64(weaponSlot.XAttach),
		float64(weaponSlot.YAttach),
	)
}

func (u *Unit) getMeleeGunAccuracy(weaponSlotNumber int) int {
	weaponSlot := u.GetMeleeWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return 0
	}
	return int(math.Ceil(u.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Weapon.Accuracy), "accuracy", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (u *Unit) getMeleeGunRotateSpeed(weaponSlotNumber int) int {
	weaponSlot := u.GetMeleeWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return 0
	}
	return int(math.Ceil(u.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Weapon.RotateSpeed), "gun_speed_rotate", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (u *Unit) getMeleeWeaponReloadTime(weaponSlotNumber int) int {
	weaponSlot := u.GetMeleeWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return 0
	}
	return int(math.Ceil(u.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Weapon.ReloadTime), "reload", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (u *Unit) getMeleeWeaponAmmoReloadTime(weaponSlotNumber int) int {
	weaponSlot := u.GetMeleeWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return 0
	}
	return int(math.Ceil(u.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Weapon.ReloadAmmoTime), "reload_ammo", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (u *Unit) getMeleeMaxDamage(weaponSlotNumber int) int {
	weaponSlot := u.GetMeleeWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil || weaponSlot.Ammo == nil {
		return 0
	}
	return int(math.Ceil(u.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Ammo.MaxDamage), "damage", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (u *Unit) getMeleeMinDamage(weaponSlotNumber int) int {
	weaponSlot := u.GetMeleeWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil || weaponSlot.Ammo == nil {
		return 0
	}
	return int(math.Ceil(u.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Ammo.MinDamage), "damage", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (u *Unit) GetMeleeWeaponData() []*obstacle_point.ObstaclePoint {
	if u.GetPhysicalModel().MeleeWeaponData.Time == u.CurrentMapTime {
		return u.GetPhysicalModel().MeleeWeaponData.GeoData
	}

	meleePositions := make([]*obstacle_point.ObstaclePoint, 0) // TODO кеширование массива что бы каждый раз заного не создавать

	for k, slot := range u.RangeMeleeWeaponSlots() {
		if slot.On && slot.Weapon != nil && slot.Weapon.Power > 0 {
			for _, firePosition := range u.meller.GetWeaponFirePos(k) {
				meleePositions = append(meleePositions, &obstacle_point.ObstaclePoint{
					X:          firePosition.X,
					Y:          firePosition.Y,
					Radius:     firePosition.Radius,
					ParentType: "melee",
					K:          slot.Weapon.Power,
				})
			}
		}
	}

	u.GetPhysicalModel().MeleeWeaponData.Time = u.CurrentMapTime
	u.GetPhysicalModel().MeleeWeaponData.GeoData = meleePositions
	return meleePositions
}
