package inventory

import (
	"encoding/json"
	"github.com/getlantern/deepcopy"
	"sync"
)

type ItemInfo struct {
	Name string `json:"name"`
	// для блюпринтов
	ItemType string `json:"item_type,omitempty"`
	ItemName string `json:"item_name,omitempty"`
	InMap    bool   `json:"in_map,omitempty"`
	Icon     string `json:"icon,omitempty"`
	// для магазина
	StandardSize int `json:"standard_size,omitempty"`
	TypeSlot     int `json:"type_slot,omitempty"`
}

type ItemInformer interface {
	GetName() string
	GetItemType() string
	GetItemName() string
	GetInMap() bool
	GetIcon() string
	GetStandardSize() int
	GetTypeSlot() int
}

func (i *ItemInfo) GetName() string {
	return i.Name
}

func (i *ItemInfo) GetItemType() string {
	return i.ItemType
}

func (i *ItemInfo) GetItemName() string {
	return i.ItemName
}

func (i *ItemInfo) GetInMap() bool {
	return i.InMap
}

func (i *ItemInfo) GetIcon() string {
	return i.Icon
}

func (i *ItemInfo) GetStandardSize() int {
	return i.StandardSize
}

func (i *ItemInfo) GetTypeSlot() int {
	return i.TypeSlot
}

func GetInfoByItem(informer ItemInformer) *ItemInfo {
	infoItem := &ItemInfo{
		Name:         informer.GetName(),
		ItemType:     informer.GetItemType(),
		ItemName:     informer.GetItemName(),
		InMap:        informer.GetInMap(),
		Icon:         informer.GetIcon(),
		StandardSize: informer.GetStandardSize(),
		TypeSlot:     informer.GetTypeSlot(),
	}

	if infoItem.Name == "" {
		return nil
	}
	return infoItem
}

type Slot struct {
	ID           int       `json:"id"`
	Item         *ItemInfo `json:"item"`
	Quantity     int       `json:"quantity"`
	Type         string    `json:"type"`
	ItemID       int       `json:"item_id"`
	HP           int       `json:"hp"`
	MaxHP        int       `json:"max_hp"`
	Size         int       `json:"size"`
	Number       int       `json:"number"`
	AccessUserID int       `json:"access_user_id"`
	Infinite     bool      `json:"infinite"` // если слот бесконечный то его нельзя разделить или перенести в трюм

	PlaceUserID int  `json:"place_user_id"`
	Tax         int  `json:"tax"`  // поле для налогов
	Find        bool `json:"find"` // поле для верстака, обозначающие естли такое количество итемов на складе или нет
	FindCount   int  `json:"find_count"`

	mx sync.RWMutex
}

func (slot *Slot) IsEmpty() bool {
	slot.mx.RLock()
	defer slot.mx.RUnlock()

	return slot.Item == nil || slot.GetQuantity() == 0
}

func (slot *Slot) GetQuantity() int {
	slot.mx.RLock()
	defer slot.mx.RUnlock()

	return slot.Quantity
}

func (slot *Slot) GetSize() int {
	slot.mx.RLock()
	defer slot.mx.RUnlock()

	return slot.Size
}

func (slot *Slot) GetOneSize() int {
	if slot.GetQuantity() == 0 {
		return 0
	}
	return slot.GetSize() / slot.GetQuantity()
}

func (slot *Slot) GetItem() *ItemInfo {
	slot.mx.RLock()
	defer slot.mx.RUnlock()

	return slot.Item
}

func (slot *Slot) GetNumber() int {
	slot.mx.RLock()
	defer slot.mx.RUnlock()

	return slot.Number
}

func (slot *Slot) GetCopy() *Slot {
	slot.mx.Lock()
	defer slot.mx.Unlock()

	var newSlot Slot
	err := deepcopy.Copy(&newSlot, &slot)
	if err != nil {
		println("copy slot: ", err.Error())
	}

	return &newSlot
}

func (slot *Slot) GetJson() string {
	slot.mx.Lock()
	defer slot.mx.Unlock()

	jsonSlot, err := json.Marshal(&slot)
	if err != nil {
		println("slot to json: ", err.Error())
	}

	return string(jsonSlot)
}

func (slot *Slot) setNumber(number int) {
	slot.mx.Lock()
	defer slot.mx.Unlock()

	slot.Number = number
}

func (slot *Slot) setItem(item *ItemInfo) {
	slot.mx.Lock()
	defer slot.mx.Unlock()

	if item == nil && slot.Infinite {
		return
	}

	slot.Item = item
}

// TODO сделать этот метод приватным и вызывать его через инвнвентарь
func (slot *Slot) SetSize(size int) {
	slot.mx.Lock()
	defer slot.mx.Unlock()

	if slot.Infinite {
		size = 0
	}

	slot.Size = size
}

// TODO сделать этот метод приватным и вызывать его через инвнвентарь
func (slot *Slot) SetQuantity(quantity int) {
	slot.mx.Lock()
	defer slot.mx.Unlock()

	if quantity <= 0 && slot.Infinite {
		return
	}

	slot.Quantity = quantity
}

func (slot *Slot) addItemBySlot(quantity, userID int) {
	// определяем вес 1 вещи
	sizeOneItem := slot.GetSize() / slot.GetQuantity()
	slot.SetQuantity(slot.GetQuantity() + quantity)
	slot.PlaceUserID = userID
	// находим новый вес для всей стопки
	slot.SetSize(sizeOneItem * slot.GetQuantity())
}

// RemoveItemBySlot когда slot.Item = nil он удалиться из бд при обновление данных
func (slot *Slot) RemoveItemBySlot(quantityRemove int) (CountRemove int) {

	if slot.GetQuantity() == 0 {
		slot.setItem(nil)
		countRemove := slot.GetQuantity()
		slot.SetQuantity(0)
		slot.SetSize(0)
		return countRemove
	}

	if quantityRemove < slot.GetQuantity() {
		// определяем вес 1 вещи
		itemSize := slot.GetSize() / slot.GetQuantity()
		// отнимает вес по количеству предметов
		slot.SetSize(slot.GetSize() - (itemSize * quantityRemove))
		// отнимаем количество итемов
		slot.SetQuantity(slot.GetQuantity() - quantityRemove)
		return quantityRemove
	} else {
		slot.setItem(nil)
		countRemove := slot.GetQuantity()
		slot.SetQuantity(0)
		slot.SetSize(0)
		return countRemove
	}
}
