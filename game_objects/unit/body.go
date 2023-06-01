package unit

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
)

func (unit *Unit) SetBody(body *detail.Body) {
	unit.mx.Lock()
	defer unit.mx.Unlock()

	unit.body = body
}

func (unit *Unit) getBody() *detail.Body {
	return unit.body
}

func (unit *Unit) GetBody() *detail.Body {
	unit.mx.Lock()
	defer unit.mx.Unlock()

	return unit.body
}

func (unit *Unit) GetBodyItem() *detail.Body {
	//TODO использовать надо как можно реже
	return unit.body
}

func (unit *Unit) GetUseEnergy() int {
	unit.mx.RLock()
	defer unit.mx.RUnlock()

	return unit.getBody().GetUseEnergy()
}

func (unit *Unit) GetUseCapacitySize() int {
	return unit.getBody().GetUseCapacitySize()
}

func (unit *Unit) GetBodyID() int {
	return unit.getBody().ID
}

func (unit *Unit) GetBodyScale() int {
	return unit.getBody().Scale
}

func (unit *Unit) GetBodyStandardSize() int {
	return unit.getBody().StandardSize
}

func (unit *Unit) GetBodyWeaponSlotKeys() []int {
	keys := make([]int, 0)

	for key := range unit.getBody().Weapons {
		keys = append(keys, key)
	}

	return keys
}

func (unit *Unit) GetBodyWeaponSlot(numberSlot int) *detail.BodyWeaponSlot {
	return unit.getBody().Weapons[numberSlot]
}

func (unit *Unit) GetBodyEquipSlot(typeSlot, numberSlot int) *detail.BodyEquipSlot {
	return unit.getBody().GetEquipSlot(typeSlot, numberSlot)
}

func (unit *Unit) GetBodyThoriumSlotKeys() []int {
	keys := make([]int, 0)

	for key := range unit.getBody().ThoriumSlots {
		keys = append(keys, key)
	}

	return keys
}

func (unit *Unit) GetMaxRecoveryReactorPower() int {
	maxRecover := 0
	for _, slot := range unit.getBody().ThoriumSlots {
		maxRecover += slot.ProcessingThorium * unit.GetRecoveryPower()
	}

	return maxRecover
}

func (unit *Unit) GetBodyThoriumSlot(numberSlot int) *detail.ThoriumSlot {
	return unit.getBody().GetThoriumSlot(numberSlot)
}

func (unit *Unit) GetBodyLength() int {
	return unit.getBody().Length
}

func (unit *Unit) GetBodyHeight() int {
	return unit.getBody().Height
}

func (unit *Unit) GetBodyWidth() int {
	return unit.getBody().Width
}

func (unit *Unit) GetAllBodyEquips() []*detail.BodyEquipSlot {
	return unit.getBody().GetAllEquipSlots()
}

func (unit *Unit) GetBodyApplicableEquips(applicable string) []*detail.BodyEquipSlot {
	return unit.getBody().GetApplicableEquips(applicable)
}

func (unit *Unit) GetBodyJSON() string {
	unit.mx.Lock()
	defer unit.mx.Unlock()

	return unit.getBody().GetJSON()
}

func (unit *Unit) SetBodySpeed(speed float64) {
	// метод только для ботов
	unit.getBody().Speed = speed
	unit.GetPhysicalModel().Speed = speed / _const.ServerTickSecPart
}

func (unit *Unit) SetSpeed(speed float64) {
	unit.GetPhysicalModel().Speed = speed / _const.ServerTickSecPart
}

func (unit *Unit) SetTurnSpeed(speed float64) {
	unit.GetPhysicalModel().TurnSpeed = speed / _const.ServerTickSecPart
}

func (unit *Unit) SetBodyMaxHP(maxHP int) {
	unit.getBody().MaxHP = maxHP
}

func (unit *Unit) GetJSONBodyEquipSlots() []string {
	return unit.getBody().GetJSONBodyEquipSlots()
}

func (unit *Unit) GetJSONBodyWeaponSlots() []string {
	return unit.getBody().GetJSONWeaponSlots()
}

func (unit *Unit) GetJSONBodyThoriumSlots() []string {
	return unit.getBody().GetJSONThoriumSlots()

}
