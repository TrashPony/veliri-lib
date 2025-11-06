package unit

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
)

func (u *Unit) SetBody(body *detail.Body) {
	u.body = body
}

func (u *Unit) getBody() *detail.Body {
	return u.body
}

func (u *Unit) GetBody() *detail.Body {
	return u.body
}

func (u *Unit) GetBodyItem() *detail.Body {
	//TODO использовать надо как можно реже
	return u.body
}

func (u *Unit) GetUseEnergy() int {
	return u.getBody().GetUseEnergy()
}

func (u *Unit) GetUseCapacitySize() int {
	return u.getBody().GetUseCapacitySize()
}

func (u *Unit) GetBodyID() int {
	return u.getBody().ID
}

func (u *Unit) GetBodyScale() int {
	return u.getBody().Scale
}

func (u *Unit) GetBodyStandardSize() int {
	return u.getBody().StandardSize
}

func (u *Unit) GetBodyWeaponSlotKeys() []int {
	keys := make([]int, 0)

	for key := range u.getBody().Weapons {
		keys = append(keys, key)
	}

	return keys
}

func (u *Unit) GetBodyMeleeWeaponSlotKeys() []int {
	keys := make([]int, 0)

	for key := range u.getBody().MeleeWeapons {
		keys = append(keys, key)
	}

	return keys
}

func (u *Unit) GetBodyWeaponSlot(numberSlot int) *detail.BodyWeaponSlot {
	return u.getBody().Weapons[numberSlot]
}

func (u *Unit) GetBodyEquipSlot(typeSlot, numberSlot int) *detail.BodyEquipSlot {
	return u.getBody().GetEquipSlot(typeSlot, numberSlot)
}

func (u *Unit) GetBodyThoriumSlotKeys() []int {
	keys := make([]int, 0)

	for key := range u.getBody().ThoriumSlots {
		keys = append(keys, key)
	}

	return keys
}

func (u *Unit) GetBodyThoriumSlot(numberSlot int) *detail.ThoriumSlot {
	return u.getBody().GetThoriumSlot(numberSlot)
}

func (u *Unit) GetBodyLength() int {
	return u.getBody().Length
}

func (u *Unit) GetBodyHeight() int {
	return u.getBody().Height
}

func (u *Unit) GetBodyWidth() int {
	return u.getBody().Width
}

func (u *Unit) GetAllBodyEquips() []*detail.BodyEquipSlot {
	return u.getBody().GetAllEquipSlots()
}

func (u *Unit) GetBodyApplicableEquips(applicable string) []*detail.BodyEquipSlot {
	return u.getBody().GetApplicableEquips(applicable)
}

func (u *Unit) GetBodyJSON() string {
	return u.getBody().GetJSON()
}

func (u *Unit) SetBodySpeed(speed float64) {
	// метод только для ботов
	u.getBody().Speed = speed
	u.GetPhysicalModel().Speed = speed / _const.ServerTickSecPart
}

func (u *Unit) SetSpeed(speed float64) {
	u.GetPhysicalModel().Speed = speed / _const.ServerTickSecPart
}

func (u *Unit) SetTurnSpeed(speed float64) {
	u.GetPhysicalModel().TurnSpeed = speed / _const.ServerTickSecPart
}

func (u *Unit) SetBodyMaxHP(maxHP int) {
	u.getBody().MaxHP = maxHP
}

func (u *Unit) GetJSONBodyEquipSlots() []string {
	return u.getBody().GetJSONBodyEquipSlots()
}

func (u *Unit) GetJSONBodyWeaponSlots() []string {
	return u.getBody().GetJSONWeaponSlots()
}

func (u *Unit) GetJSONBodyThoriumSlots() []string {
	return u.getBody().GetJSONThoriumSlots()
}
