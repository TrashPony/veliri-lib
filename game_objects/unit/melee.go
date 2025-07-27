package unit

import (
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
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
