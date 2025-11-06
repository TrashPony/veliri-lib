package gunner

import (
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"github.com/TrashPony/veliri-lib/game_objects/target"
	"math/rand"
)

type MeleeUser interface {
	GetX() int
	GetY() int
	GetRotate() float64
	GetMapHeight() float64
	GetScale() int
	GetWeaponTarget(int) *target.Target
	RangeMeleeWeaponSlots() map[int]*detail.BodyWeaponSlot
	GetMeleeWeaponSlot(slotNumber int) *detail.BodyWeaponSlot
}

type Meleer struct {
	MeleeUser        MeleeUser
	WeaponSlotsState []*WeaponSlotState
	r                *rand.Rand
}

func (g *Meleer) GetX() int {
	return g.MeleeUser.GetX()
}

func (g *Meleer) GetY() int {
	return g.MeleeUser.GetY()
}

func (g *Meleer) GetRotate() float64 {
	return g.MeleeUser.GetRotate()
}

func (g *Meleer) GetMapHeight() float64 {
	return g.MeleeUser.GetMapHeight()
}

func (g *Meleer) GetWeaponTarget(slot int) *target.Target {
	if g == nil {
		return nil
	}

	return g.MeleeUser.GetWeaponTarget(0)
}

func (g *Meleer) GetGunRotate(slotNumber int) float64 {
	weaponSlot := g.MeleeUser.GetMeleeWeaponSlot(slotNumber)
	return weaponSlot.GetGunRotate()
}

func (g *Meleer) GetGunRotateSpeed(slotNumber int) int {
	weaponSlot := g.MeleeUser.GetMeleeWeaponSlot(slotNumber)
	slotState := g.getSlotState(slotNumber)
	if weaponSlot == nil || slotState == nil {
		return 0
	}

	return slotState.RotateSpeed
}

func (g *Meleer) SetGunRotate(angle float64, slotNumber int) {
	weaponSlot := g.MeleeUser.GetMeleeWeaponSlot(slotNumber)
	if weaponSlot == nil {
		return
	}

	weaponSlot.SetGunRotate(angle)
}

func (g *Meleer) GetFirePos(slotNumber int) *game_math.Positions {
	// отдае координату откуда стрелять
	weaponSlot := g.MeleeUser.GetMeleeWeaponSlot(slotNumber)

	if len(weaponSlot.Weapon.FirePositions) <= weaponSlot.GetLastFirePosition() {
		weaponSlot.NextLastFirePosition()
	}

	return g.GetWeaponFirePosOne(slotNumber, weaponSlot.GetLastFirePosition())
}

func (g *Meleer) GetWeaponFirePosOne(slotNumber, position int) *game_math.Positions {
	weaponSlot := g.MeleeUser.GetMeleeWeaponSlot(slotNumber)

	return game_math.GetWeaponFirePosition(
		g.MeleeUser.GetX(), g.MeleeUser.GetY(), g.MeleeUser.GetScale(), g.MeleeUser.GetRotate(), weaponSlot.GetGunRotate(),
		weaponSlot.Weapon.XAttach, weaponSlot.Weapon.YAttach,
		weaponSlot.Weapon.FirePositions,
		float64(weaponSlot.XAttach),
		float64(weaponSlot.YAttach),
		position,
	)
}

func (g *Meleer) RangeWeaponSlots() map[int]*detail.BodyWeaponSlot {
	return g.MeleeUser.RangeMeleeWeaponSlots()
}

func (g *Meleer) getSlotState(number int) *WeaponSlotState {
	for _, s := range g.WeaponSlotsState {
		if s != nil && s.Number == number {
			return s
		}
	}

	return nil
}

func (g *Meleer) IsValid() bool {
	return g != nil
}
