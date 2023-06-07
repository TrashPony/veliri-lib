package inventory

import (
	"errors"
	"github.com/getlantern/deepcopy"
	"sync"
)

type Inventory struct {
	slots map[int]*Slot
	size  int
	mx    sync.RWMutex
}

type PlaceMayItems struct {
	Slot         int    `json:"slot"`
	CountMayPut  int    `json:"count_may_put"`
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
			_, _, err := inv.InventoryToInventory(toInventory, toInventoryCapSize, slot, userAccessID, count, noPlace)
			if err != nil {
				return nil, err
			}
		}

		return nil, nil
	} else {
		return placeMayItems, errors.New("weight exceeded")
	}
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

			slotQuantity, countMayPut, err := srcInvCopy.InventoryToInventory(toInvCopy, toInventoryCapSize, slot, userAccessID, count, noPlace)
			if !fillCount && placeMayItems[slot] != nil {
				placeMayItems[slot].CountMayPut = placeMayItems[slot].Count + countMayPut
			} else {

				var errString string
				if err != nil {
					errString = err.Error()
				}

				placeMayItems[slot] = &PlaceMayItems{
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
		_, _, err := srcInvCopy2.InventoryToInventory(toInvCopy2, toInventoryCapSize, slot, userAccessID, count, false)

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

func (inv *Inventory) InventoryToInventory(toInventory *Inventory, toInventoryCapSize int, inventorySlotNumber, accessUserID, count int, noPlace bool) (int, int, error) {

	if inv == nil {
		return 0, 0, errors.New("no inventory")
	}

	slot, _ := inv.GetSlot(inventorySlotNumber, accessUserID)

	if slot == nil {
		return 0, 0, errors.New("no find slot")
	}

	if slot.Infinite {
		return 0, 0, errors.New("no allow")
	}

	if slot.Quantity == 0 || slot.GetItem() == nil {
		return 0, 0, nil
	}

	countPlace := slot.Quantity
	startQuantity := slot.Quantity
	if count > 0 {
		countPlace = count
	}

	if countPlace > slot.Quantity || slot.GetOneSize() == 0 {
		return 0, 0, errors.New("wrong count items")
	}

	FreeSize := toInventoryCapSize - toInventory.GetSize()

	countPut := FreeSize / slot.GetOneSize()

	if countPut < countPlace && toInventoryCapSize != -1 { // -1 означает что инвентарь бесконечен
		return startQuantity, countPut, errors.New("weight exceeded")
	} else {
		if !noPlace {
			ok := toInventory.AddItem(slot.GetItem(), slot.Type, slot.ItemID, countPlace,
				slot.HP, slot.GetOneSize(), slot.MaxHP, false, accessUserID, 0)

			if !ok {
				return startQuantity, countPut, errors.New("no free slots")
			}

			slot.RemoveItemBySlot(countPlace)
		}

		return startQuantity, countPut, nil
	}
}

func (inv *Inventory) GetCopyInventory() *Inventory {

	if inv == nil {
		return nil
	}

	inv.mx.RLock()
	defer inv.mx.RUnlock()

	copyInventory := Inventory{size: inv.size}
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

func (inv *Inventory) IsNil() bool {
	return inv == nil || inv.slots == nil || inv.size == 0
}

func (inv *Inventory) SetSlotsSize(size int) {

	inv.mx.Lock()
	defer inv.mx.Unlock()

	inv.slots = make(map[int]*Slot)
	inv.size = size
}

func (inv *Inventory) SetSlots(slots map[int]*Slot) {
	inv.mx.Lock()
	defer inv.mx.Unlock()
	inv.slots = slots
}

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

func (inv *Inventory) GetArraySlots() []*Slot {

	if inv == nil {
		return nil
	}

	inv.mx.Lock()
	defer inv.mx.Unlock()

	slots := make([]*Slot, 0, len(inv.slots))
	for _, slot := range inv.slots {
		slots = append(slots, slot)
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
			slot.SetNumber(number)
			chanSlots <- slot
		}
	}()

	return chanSlots
}

func (inv *Inventory) DeleteEmptySlot(number int) {
	inv.mx.Lock()
	defer inv.mx.Unlock()

	slot := inv.slots[number]
	if slot != nil && slot.Quantity == 0 {
		delete(inv.slots, number)
	}
}

func (inv *Inventory) DeleteSlot(number int) {
	inv.mx.Lock()
	defer inv.mx.Unlock()

	delete(inv.slots, number)
}

func (inv *Inventory) GetSlot(number, userID int) (*Slot, bool) {
	inv.mx.Lock()
	defer inv.mx.Unlock()

	slot, ok := inv.slots[number]
	if slot != nil && (slot.AccessUserID == 0 || slot.AccessUserID == userID || userID == -1) {
		return slot, ok
	} else {
		return nil, false
	}
}

func (inv *Inventory) AddItemFromSlot(slot *Slot, userID, accessUserID int) bool {

	if slot.GetQuantity() <= 0 { // slot.Size/float32(slot.Quantity) деление на ноль все сломает
		return false
	}

	return inv.AddItem(slot.GetItem(), slot.Type, slot.ItemID, slot.GetQuantity(),
		slot.HP, slot.GetSize()/slot.GetQuantity(), slot.MaxHP, false, userID, accessUserID)
}

func (inv *Inventory) AddItemFromSlotByQuantity(slot *Slot, userID, accessUserID, quantity int) bool {

	if slot.GetQuantity() <= 0 { // slot.Size/float32(slot.Quantity) деление на ноль все сломает
		return false
	}

	return inv.AddItem(slot.GetItem(), slot.Type, slot.ItemID, quantity,
		slot.HP, slot.GetSize()/slot.GetQuantity(), slot.MaxHP, false, userID, accessUserID)
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

func (inv *Inventory) AddItem(item ItemInformer, itemType string, itemID, quantity, hp int, itemSize int, maxHP int, newSlot bool, userID, accessUserID int) bool {

	inv.mx.Lock()
	defer inv.mx.Unlock()

	if quantity == 0 {
		return true
	}

	//newSlot говорит о том что этот айтем надо положить в строго новый слот
	if !newSlot {
		for _, slot := range inv.slots { // ищем стопку с такими же элементами
			if slot.AccessUserID == accessUserID {
				if slot.ItemID == itemID && slot.Type == itemType && slot.HP == hp && slot.GetItem() != nil {
					slot.SetQuantity(slot.GetQuantity() + quantity)
					slot.SetSize(slot.GetSize() + (itemSize * quantity))
					slot.PlaceUserID = userID
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
			InsertToDB:   true,
			Quantity:     quantity,
			HP:           hp,
			MaxHP:        maxHP,
			Size:         itemSize * quantity,
			PlaceUserID:  userID,
			Number:       newNumberSlot,
			AccessUserID: accessUserID,
		}

		inv.slots[newNumberSlot] = &newItem
		return true
	}

	return false
}

func (inv *Inventory) GetEmptySlot() int {
	for i := 1; i <= inv.size; i++ { // ищем пустой слот
		_, ok := inv.slots[i]
		if !ok {
			return i
		}
	}
	return 0
}

func (inv *Inventory) RemoveItem(itemID int, itemType string, quantityRemove int) error {
	// надо убедиться что необходимое количество есть и его можно удалить
	if inv.ViewItems(itemID, itemType, quantityRemove) {

		inv.mx.Lock()
		defer inv.mx.Unlock()

		for _, slot := range inv.slots {
			if slot.ItemID == itemID && slot.Type == itemType {
				if slot.GetQuantity() >= quantityRemove {
					slot.RemoveItemBySlot(quantityRemove)
					return nil
				} else {
					// если в слоте не чего либо для полного удаления,
					// то удаляем все из слота, и уменьшаем количество итемов которые еще надо удалить
					quantityRemove -= slot.GetQuantity()
					slot.RemoveItemBySlot(slot.GetQuantity())
				}
			}
		}
		return nil

	} else {
		return errors.New("few items")
	}
}

// метод делает сравнение инвентарей слот к слоту
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

// метод смотрит все предметы inv2 что бы они были в inv на наличие
func (inv *Inventory) SearchItemsByOtherInventory(inv2 *Inventory) bool {
	for _, slot := range inv2.slots {
		if !inv.ViewItems(slot.ItemID, slot.Type, slot.GetQuantity()) {
			return false
		}
	}
	return true
}

// метод удаляем все итемы из inv которые есть в inv2 если они все в наличие
func (inv *Inventory) RemoveItemsByOtherInventory(inv2 *Inventory, force bool) bool {
	for _, slot := range inv2.slots {
		if !inv.ViewItems(slot.ItemID, slot.Type, slot.GetQuantity()) {
			if !force {
				return false
			}
		}
	}

	for _, removeSlot := range inv2.slots {
		quantityRemove := removeSlot.GetQuantity()
		for slot := range inv.GetSlotsChan() {
			if slot.ItemID == removeSlot.ItemID && slot.Type == removeSlot.Type {
				if slot.GetQuantity() >= quantityRemove {
					slot.RemoveItemBySlot(quantityRemove)
				} else {
					// если в слоте не чего либо для полного удаления,
					// то удаляем все из слота, и уменьшаем количество итемов которые еще надо удалить
					quantityRemove -= slot.GetQuantity()
					slot.RemoveItemBySlot(slot.GetQuantity())
				}
			}
		}
	}

	return true
}

// метод считает все итемы в инвентаре
func (inv *Inventory) ViewQuantityItems(itemID int, itemType string) int {
	inv.mx.Lock()
	defer inv.mx.Unlock()

	countRealItems := 0
	for _, slot := range inv.slots {
		if slot.ItemID == itemID && slot.Type == itemType {
			countRealItems += slot.GetQuantity()
		}
	}

	return countRealItems
}

// метод смотрим естли необходимое количество предметов в инвентаре
func (inv *Inventory) ViewItems(itemID int, itemType string, quantityFind int) bool {
	if inv.ViewQuantityItems(itemID, itemType) >= quantityFind {
		return true
	} else {
		return false
	}
}

func (inv *Inventory) AddSlot(slot int, inventorySlot *Slot) {
	inv.mx.Lock()
	inventorySlot.SetNumber(slot)
	inv.slots[slot] = inventorySlot
	inv.mx.Unlock()
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
		if slot != nil && slot.Type == itemType && slot.ItemID == itemID {
			slots = append(slots, slot)
		}
	}

	return slots
}
