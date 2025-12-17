package inventory

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/getlantern/deepcopy"
)

func (inv *Inventory) GetSlots() map[int]*Slot {

	if inv == nil {
		return nil
	}

	inv.mx.Lock()
	defer inv.mx.Unlock()

	slots := make(map[int]*Slot)
	for _, slot := range inv.slots {
		slots[slot.GetNumber()] = slot
	}

	return slots
}

func (inv *Inventory) GetJsonSlots() map[int]string {

	if inv == nil {
		return nil
	}

	inv.mx.Lock()
	defer inv.mx.Unlock()

	slots := make(map[int]string)

	for _, slot := range inv.slots {
		slots[slot.GetNumber()] = slot.GetJson()
	}

	return slots
}

func (inv *Inventory) GetSlotsChan() <-chan *Slot {
	inv.mx.Lock()
	chanSlots := make(chan *Slot, len(inv.slots))

	go func() {
		defer func() {
			close(chanSlots)
			inv.mx.Unlock()
		}()

		for number, slot := range inv.slots {
			slot.setNumber(number)
			chanSlots <- slot
		}
	}()

	return chanSlots
}

func (inv *Inventory) GetSlot(number, userID int) (*Slot, bool) {
	inv.mx.Lock()
	defer inv.mx.Unlock()
	return inv.getSlot(number, userID)
}

func (inv *Inventory) getSlot(number, userID int) (*Slot, bool) {
	slot, ok := inv.slots[number]
	if slot != nil && (slot.AccessUserID == 0 || slot.AccessUserID == userID || userID == -1) {
		return slot, ok
	} else {
		return nil, false
	}
}

func (inv *Inventory) GetSize() int {

	if inv == nil {
		return 0
	}

	inv.mx.Lock()
	defer inv.mx.Unlock()

	var inventorySquadSize int
	for _, slot := range inv.slots {
		if slot.GetItem() != nil {
			inventorySquadSize = inventorySquadSize + slot.GetSize()
		}
	}

	return inventorySquadSize
}

func (inv *Inventory) GetEmptySlot() int {
	for i := 1; i <= 9999; i++ { // ищем пустой слот
		_, ok := inv.slots[i]
		if !ok {
			return i
		}
	}
	return 0
}

// ViewItemsBySlots метод делает сравнение инвентарей слот к слоту
func (inv *Inventory) ViewItemsBySlots(slots map[int]*Slot) bool {
	checkItems := true
	for number, slot := range slots {
		realSlot, findSlot := inv.slots[number]
		if !findSlot || slot == nil || slot.GetQuantity() > realSlot.GetQuantity() {
			checkItems = false
		}
	}
	return checkItems
}

// SearchItemsByOtherInventory метод смотрит все предметы inv2 что бы они были в inv на наличие
func (inv *Inventory) SearchItemsByOtherInventory(inv2 *Inventory) bool {
	for _, slot := range inv2.slots {
		if !inv.ViewItems(slot.ItemID, slot.Type, slot.GetQuantity(), ItemFilter{}) {
			return false
		}
	}
	return true
}

// ViewQuantityItems метод считает все итемы в инвентаре
func (inv *Inventory) ViewQuantityItems(itemID int, itemType string, filter ItemFilter) int {
	inv.mx.Lock()
	defer inv.mx.Unlock()
	return inv.viewQuantityItems(itemID, itemType, filter)
}

func (inv *Inventory) viewQuantityItems(itemID int, itemType string, filter ItemFilter) int {
	countRealItems := 0
	for _, slot := range inv.slots {
		if slot.ItemID == itemID && slot.Type == itemType {

			if filter.OnlyPerfect && (slot.Durability < _const.StartDurability || slot.HP < slot.MaxHP) {
				continue
			}

			countRealItems += slot.GetQuantity()
		}
	}

	return countRealItems
}

// ViewItems метод смотрим естли необходимое количество предметов в инвентаре
func (inv *Inventory) ViewItems(itemID int, itemType string, quantityFind int, filter ItemFilter) bool {
	if inv.ViewQuantityItems(itemID, itemType, filter) >= quantityFind {
		return true
	} else {
		return false
	}
}

// viewItems метод смотрим естли необходимое количество предметов в инвентаре
func (inv *Inventory) viewItems(itemID int, itemType string, quantityFind int, filter ItemFilter) bool {
	if inv.viewQuantityItems(itemID, itemType, filter) >= quantityFind {
		return true
	} else {
		return false
	}
}

func (inv *Inventory) GetCellsByType(itemType string) []*Slot {

	inv.mx.Lock()
	defer inv.mx.Unlock()

	slots := make([]*Slot, 0)

	for _, slot := range inv.slots {
		if slot != nil && slot.Type == itemType || itemType == "all" {
			slots = append(slots, slot)
		}
	}

	return slots
}

func (inv *Inventory) GetSlotByTypeAndIDItem(itemID int, itemType string) []*Slot {
	inv.mx.Lock()
	defer inv.mx.Unlock()

	slots := make([]*Slot, 0)

	for _, slot := range inv.slots {
		if slot != nil && slot.Quantity > 0 && slot.Type == itemType && slot.ItemID == itemID {
			slots = append(slots, slot)
		}
	}

	return slots
}

func (inv *Inventory) GetCopyInventory() *Inventory {

	if inv == nil {
		return nil
	}

	inv.mx.RLock()
	defer inv.mx.RUnlock()

	copyInventory := Inventory{}
	if inv.slots == nil {
		return &copyInventory
	} else {
		copyInventory.slots = make(map[int]*Slot)
	}

	for number, slot := range inv.slots {

		var copySlot Slot

		err := deepcopy.Copy(&copySlot, &slot)
		if err != nil {
			println("GetCopyInventory", err.Error())
		}

		copyInventory.slots[number] = &copySlot
	}

	return &copyInventory
}

func (inv *Inventory) CheckPlaceItems(toInventory *Inventory, toInventoryCapSize int, slots map[int]int, userAccessID int) ([]*PlaceMayItems, bool) {

	srcInvCopy := inv.GetCopyInventory()
	toInvCopy := toInventory.GetCopyInventory()

	placeMayItems := make(map[int]*PlaceMayItems)

	put := func(fillCount bool) {
		for slot, count := range slots {

			if count == 0 && fillCount {
				continue
			}

			if !fillCount {
				count = 0
			}

			noPlace := !fillCount

			slotQuantity, countMayPut, err, itemType, itemID := srcInvCopy.InventoryToInventory(toInvCopy, toInventoryCapSize, slot, userAccessID, count, noPlace)
			if !fillCount && placeMayItems[slot] != nil {
				placeMayItems[slot].CountMayPut = placeMayItems[slot].Count + countMayPut
			} else {

				var errString string
				if err != nil {
					errString = err.Error()
				}

				placeMayItems[slot] = &PlaceMayItems{
					Type:         itemType,
					ID:           itemID,
					Slot:         slot,
					CountMayPut:  countMayPut,
					SlotQuantity: slotQuantity,
					Count:        count,
					Error:        errString,
					AllPlace:     count != 0 && count <= countMayPut,
				}

				if count > countMayPut {
					placeMayItems[slot].Count = countMayPut
				}
			}

			if placeMayItems[slot] != nil && placeMayItems[slot].AllPlace {
				continue
			}
		}
	}

	put(true)
	put(false)

	// todo мега еба костыль что бы посмотреть рально ли все итемы поместяться
	srcInvCopy2 := inv.GetCopyInventory()
	toInvCopy2 := toInventory.GetCopyInventory()
	allPlace := true
	for slot, count := range slots {
		_, _, err, _, _ := srcInvCopy2.InventoryToInventory(toInvCopy2, toInventoryCapSize, slot, userAccessID, count, false)

		if err != nil {
			allPlace = false
			if placeMayItems[slot] != nil && placeMayItems[slot].Error == "" {
				placeMayItems[slot].Error = ""
			}
		}
	}

	placeMayItemsArray := make([]*PlaceMayItems, 0)

	for _, item := range placeMayItems {
		placeMayItemsArray = append(placeMayItemsArray, item)
	}

	return placeMayItemsArray, allPlace
}
