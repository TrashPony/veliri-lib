package inventory

import (
	"errors"
	"sync"
)

// TODO рефакторинг

type Inventory struct {
	slots      map[int]*Slot
	parentType string
	parentID   int
	logger     func(parentType string, parentID int, event string, data map[string]interface{})
	updater    func(slots map[int]*Slot)
	mx         sync.RWMutex
}

func (inv *Inventory) GetParentType() string {
	return inv.parentType
}

func (inv *Inventory) GetParentID() int {
	return inv.parentID
}

type PlaceMayItems struct {
	Type         string `json:"type"`
	ID           int    `json:"id"`
	Slot         int    `json:"slot"`
	CountMayPut  int    `json:"count_may_put"`
	CountCurrent int    `json:"count_current"`
	SlotQuantity int    `json:"slot_quantity"`
	Count        int    `json:"count"`
	Error        string `json:"error"`
	AllPlace     bool   `json:"all_place"`
}

func (inv *Inventory) CheckAndPlaceItems(toInventory *Inventory, toInventoryCapSize int, slots map[int]int, noPlace bool, userAccessID int) ([]*PlaceMayItems, error) {
	placeMayItems, allPlace := inv.CheckPlaceItems(toInventory, toInventoryCapSize, slots, userAccessID)
	if noPlace {
		return placeMayItems, nil
	}

	if allPlace {
		for slot, count := range slots {
			_, _, err, _, _ := inv.InventoryToInventory(toInventory, toInventoryCapSize, slot, userAccessID, count, noPlace)
			if err != nil {
				return placeMayItems, err
			}
		}

		return placeMayItems, nil
	} else {
		return placeMayItems, errors.New("weight exceeded")
	}
}

func (inv *Inventory) InventoryToInventory(toInventory *Inventory, toInventoryCapSize int, inventorySlotNumber, accessUserID, count int, noPlace bool) (int, int, error, string, int) {
	if inv == nil {
		return 0, 0, errors.New("no inventory"), "", 0
	}

	inv.mx.Lock()
	defer func() {
		inv.mx.Unlock()
		inv.update()
	}()

	s, _ := inv.getSlot(inventorySlotNumber, accessUserID)

	if s == nil {
		return 0, 0, errors.New("no find slot"), "", 0
	}

	if s.Infinite {
		return 0, 0, errors.New("no allow"), s.Type, s.ItemID
	}

	if s.Quantity == 0 || s.GetItem() == nil {
		return 0, 0, nil, s.Type, s.ItemID
	}

	countPlace := s.Quantity
	startQuantity := s.Quantity
	if count > 0 {
		countPlace = count
	}

	if countPlace > s.Quantity || s.GetOneSize() == 0 {
		return 0, 0, errors.New("wrong count items"), s.Type, s.ItemID
	}

	FreeSize := toInventoryCapSize - toInventory.GetSize()

	countPut := FreeSize / s.GetOneSize()

	if countPut < countPlace && toInventoryCapSize != -1 { // -1 означает что инвентарь бесконечен
		return startQuantity, countPut, errors.New("weight exceeded"), s.Type, s.ItemID
	} else {
		if !noPlace {
			ok := toInventory.AddItem(s.GetItem(), s.Type, s.ItemID, countPlace,
				s.HP, s.MaxHP, false, accessUserID, 0)

			if !ok {
				return startQuantity, countPut, errors.New("no free slots"), s.Type, s.ItemID
			}
			inv.log("RemoveItem", map[string]interface{}{"item_type": s.Type, "item_id": s.ItemID, "quantity_remove": countPlace, "real_remove": s.removeItemBySlot(countPlace)})
		}

		return startQuantity, countPut, nil, s.Type, s.ItemID
	}
}

func (inv *Inventory) IsNil() bool {
	return inv == nil || inv.slots == nil
}

func (inv *Inventory) Init(parentType string, parentID int, logger func(parentType string, parentID int, event string, data map[string]interface{}), updater func(map[int]*Slot)) {
	inv.mx.Lock()
	defer inv.mx.Unlock()
	inv.slots = make(map[int]*Slot)
	inv.parentType = parentType
	inv.parentID = parentID
	inv.logger = logger
	inv.updater = updater
}

func (inv *Inventory) SetSlots(slots map[int]*Slot) {
	inv.mx.Lock()
	defer inv.mx.Unlock()
	inv.slots = slots
}

func (inv *Inventory) DeleteEmptySlot(number int) {
	inv.mx.Lock()
	defer inv.mx.Unlock()

	slot := inv.slots[number]
	if slot != nil && slot.IsEmpty() {
		delete(inv.slots, number)
	}
}

func (inv *Inventory) DeleteSlot(number int) {
	inv.mx.Lock()
	defer func() {
		inv.mx.Unlock()
		inv.update()
	}()

	s, ok := inv.slots[number]
	if !ok {
		return
	}

	inv.log("DeleteSlot", map[string]interface{}{"item_type": s.Type, "item_id": s.ItemID, "quantity": s.Quantity})
	delete(inv.slots, number)
}

func (inv *Inventory) UnsafeDeleteSlot(number int) {
	_, ok := inv.slots[number]
	if !ok {
		return
	}

	delete(inv.slots, number)
}

func (inv *Inventory) AddItemFromSlot(slot *Slot, userID, accessUserID int) bool {

	if slot.GetQuantity() <= 0 || slot.Item == nil { // slot.Size/float32(slot.Quantity) деление на ноль все сломает
		return false
	}

	return inv.AddItem(slot.GetItem(), slot.Type, slot.ItemID, slot.GetQuantity(),
		slot.HP, slot.MaxHP, false, userID, accessUserID)
}

func (inv *Inventory) AddItemFromSlotByQuantity(slot *Slot, userID, accessUserID, quantity int) bool {

	if slot.GetQuantity() <= 0 { // slot.Size/float32(slot.Quantity) деление на ноль все сломает
		return false
	}

	return inv.AddItem(slot.GetItem(), slot.Type, slot.ItemID, quantity,
		slot.HP, slot.MaxHP, false, userID, accessUserID)
}

func (inv *Inventory) AddItem(item ItemInformer, itemType string, itemID, quantity, hp int, maxHP int, newSlot bool, userID, accessUserID int) bool {
	inv.mx.Lock()
	defer func() {
		inv.mx.Unlock()
		inv.update()
	}()

	if quantity == 0 {
		return true
	}

	//newSlot говорит о том что этот айтем надо положить в строго новый слот
	if !newSlot {
		for _, s := range inv.slots { // ищем стопку с такими же элементами
			if s.AccessUserID == accessUserID {
				if s.ItemID == itemID && s.Type == itemType && s.HP == hp && s.GetItem() != nil {

					s.SetQuantity(s.GetQuantity() + quantity)
					s.PlaceUserID = userID

					inv.log("AddItem", map[string]interface{}{"up_slot": true, "item_type": s.Type, "item_id": s.ItemID, "quantity": quantity})
					return true
				}
			}
		}
	}

	newNumberSlot := inv.GetEmptySlot()
	if newNumberSlot > 0 {

		newItem := Slot{
			Item:         GetInfoByItem(item),
			Type:         itemType,
			ItemID:       itemID,
			Quantity:     quantity,
			HP:           hp,
			MaxHP:        maxHP,
			PlaceUserID:  userID,
			Number:       newNumberSlot,
			AccessUserID: accessUserID,
		}

		inv.slots[newNumberSlot] = &newItem
		inv.log("AddItem", map[string]interface{}{"new_slot": true, "item_type": newItem.Type, "item_id": newItem.ItemID, "quantity": quantity})
		return true
	}

	return false
}

func (inv *Inventory) RemoveItem(itemID int, itemType string, quantityRemove int) error {
	// надо убедиться что необходимое количество есть и его можно удалить
	if inv.ViewItems(itemID, itemType, quantityRemove) {

		inv.mx.Lock()
		defer func() {
			inv.mx.Unlock()
			inv.update()
		}()

		for _, s := range inv.slots {
			if s.ItemID == itemID && s.Type == itemType {
				if s.GetQuantity() >= quantityRemove {
					inv.log("RemoveItem", map[string]interface{}{"item_type": s.Type, "item_id": s.ItemID, "quantity_remove": quantityRemove, "real_remove": s.removeItemBySlot(quantityRemove)})
					return nil
				} else {
					// если в слоте не чего либо для полного удаления,
					// то удаляем все из слота, и уменьшаем количество итемов которые еще надо удалить
					quantityRemove -= s.GetQuantity()
					inv.log("RemoveItem", map[string]interface{}{"item_type": s.Type, "item_id": s.ItemID, "quantity_remove": quantityRemove, "real_remove": s.removeItemBySlot(s.GetQuantity())})
				}
			}
		}
		return nil

	} else {
		return errors.New("few items")
	}
}

// RemoveItemsByOtherInventory метод удаляем все итемы из inv которые есть в inv2 если они все в наличие
func (inv *Inventory) RemoveItemsByOtherInventory(inv2 *Inventory, force bool) bool {

	for _, slot := range inv2.slots {
		if !inv.ViewItems(slot.ItemID, slot.Type, slot.GetQuantity()) {
			if !force {
				return false
			}
		}
	}

	inv.mx.Lock()
	defer func() {
		inv.mx.Unlock()
		inv.update()
	}()

	for _, removeSlot := range inv2.slots {
		quantityRemove := removeSlot.GetQuantity()
		for _, s := range inv.slots {
			if s.ItemID == removeSlot.ItemID && s.Type == removeSlot.Type {
				if s.GetQuantity() >= quantityRemove {
					inv.log("RemoveItem", map[string]interface{}{"item_type": s.Type, "item_id": s.ItemID, "quantity_remove": quantityRemove, "real_remove": s.removeItemBySlot(quantityRemove)})
				} else {
					// если в слоте не чего либо для полного удаления,
					// то удаляем все из слота, и уменьшаем количество итемов которые еще надо удалить
					quantityRemove -= s.GetQuantity()
					inv.log("RemoveItem", map[string]interface{}{"item_type": s.Type, "item_id": s.ItemID, "quantity_remove": quantityRemove, "real_remove": s.removeItemBySlot(s.GetQuantity())})
				}
			}
		}
	}

	return true
}

func (inv *Inventory) AddItemBySlot(slot, quantity, userID int) bool {
	inv.mx.Lock()
	defer func() {
		inv.mx.Unlock()
		inv.update()
	}()

	s, ok := inv.getSlot(slot, userID)
	if !ok {
		return false
	}

	s.addItemBySlot(quantity, userID)
	inv.log("AddItem", map[string]interface{}{"item_type": s.Type, "item_id": s.ItemID, "quantity": quantity})

	return true
}

func (inv *Inventory) RemoveItemBySlot(slot, quantityRemove, userID int) int {
	inv.mx.Lock()
	defer func() {
		inv.mx.Unlock()
		inv.update()
	}()

	s, ok := inv.getSlot(slot, userID)
	if !ok {
		return -1
	}

	realRemove := s.removeItemBySlot(quantityRemove)
	inv.log("RemoveItem", map[string]interface{}{"item_type": s.Type, "item_id": s.ItemID, "quantity_remove": quantityRemove, "real_remove": realRemove})

	return realRemove
}

func (inv *Inventory) AddSlot(slot int, inventorySlot *Slot) {
	inv.mx.Lock()
	defer inv.mx.Unlock()

	inventorySlot.setNumber(slot)
	inv.slots[slot] = inventorySlot
}

func (inv *Inventory) InitSlot(slot int, item *ItemInfo, maxHP int) {
	inv.mx.Lock()
	defer inv.mx.Unlock()
	s, ok := inv.slots[slot]
	if !ok {
		return
	}

	s.setItem(item)
	s.MaxHP = maxHP
}

func (inv *Inventory) log(event string, data map[string]interface{}) {
	if inv.logger == nil || inv.parentType == "" || inv.parentID == 0 {
		return
	}

	inv.logger(inv.parentType, inv.parentID, event, data)
}

func (inv *Inventory) update() {
	inv.mx.Lock()
	defer inv.mx.Unlock()

	if inv.updater != nil {
		inv.updater(inv.slots)
	}
}

// ItemToSlot костыль метод используется только для конфигурации, использование в других участах может привести к потери итемов
func (inv *Inventory) ItemToSlot(itemID int, itemType string, toSLot int) {
	inv.mx.Lock()
	defer inv.mx.Unlock()

	slotNumber := -1
	for n, slot := range inv.slots {
		if slot.ItemID == itemID && slot.Type == itemType && toSLot > n {
			slotNumber = n
		}
	}

	if slotNumber == -1 || slotNumber == toSLot {
		return
	}

	buffer := inv.slots[toSLot]

	inv.slots[toSLot] = inv.slots[slotNumber]
	inv.slots[toSLot].Number = toSLot

	if buffer != nil {
		inv.slots[slotNumber] = buffer
		inv.slots[slotNumber].Number = slotNumber
	} else {
		delete(inv.slots, slotNumber)
	}
}
