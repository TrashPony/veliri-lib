package squad

import (
	"encoding/json"
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_math"
	"sort"
)

// ячейки которые отображены на панеле быстрого доступа игрока
type EquipSell struct {
	Number int `json:"number"`

	Source   string `json:"source"`    // squadInventory, Constructor, "empty" - убрать все выделения
	TypeSlot int    `json:"type_slot"` // 0 - weapon. 1,2,3 - body
	Slot     int    `json:"slot"`

	// эти параметры берутся каждый тик по состоянию снаряжения дл фронтенда
	StartReload  int64 `json:"sr"`
	EndReload    int64 `json:"er"`
	AmmoQuantity int   `json:"aq"`
	AmmoReload   bool  `json:"ar"`
	On           bool  `json:"on"`
	Mode         int   `json:"mode"`
}

func (s *Squad) GetEquipPanel() map[int]*EquipSell {

	if s.equipPanel == nil {
		s.mx.Lock()
		s.equipPanel = make(map[int]*EquipSell)
		s.equipPanel[0] = &EquipSell{TypeSlot: -1, Source: "empty"}
		s.mx.Unlock()
	}

	s.mx.RLock()
	defer s.mx.RUnlock()

	return s.equipPanel
}

func (s *Squad) GetCopyEquipPanel() map[int]*EquipSell {
	s.GetEquipPanel()

	s.mx.RLock()
	defer s.mx.RUnlock()

	equipPanel := make(map[int]*EquipSell)
	for k, v := range s.equipPanel {
		equipPanel[k] = v
	}

	return equipPanel
}

func (s *Squad) GetBinEquipPanel() []byte {

	s.fillStateEquip()

	s.mx.RLock()
	defer s.mx.RUnlock()

	keys := make([]int, 0, len(s.equipPanel))
	for k := range s.equipPanel {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	sizeBytes := 30
	binEquip := make([]byte, 1+(len(keys)*sizeBytes))
	binEquip[0] = byte(len(s.equipPanel))

	for i, slotKey := range keys {
		slot := s.equipPanel[slotKey]
		if slot == nil {
			continue
		}

		binEquip[((i+1)*sizeBytes)-29] = byte(_const.SourceItemBin[slot.Source])
		binEquip[((i+1)*sizeBytes)-28] = byte(slot.TypeSlot)
		game_math.ReuseByteSlice(&binEquip, ((i+1)*sizeBytes)-27, game_math.GetIntBytes(slot.Slot))
		game_math.ReuseByteSlice(&binEquip, ((i+1)*sizeBytes)-23, game_math.GetInt64Bytes(slot.StartReload))
		game_math.ReuseByteSlice(&binEquip, ((i+1)*sizeBytes)-15, game_math.GetInt64Bytes(slot.EndReload))
		game_math.ReuseByteSlice(&binEquip, ((i+1)*sizeBytes)-7, game_math.GetIntBytes(slot.AmmoQuantity))
		binEquip[((i+1)*sizeBytes)-3] = game_math.BoolToByte(slot.On)
		binEquip[((i+1)*sizeBytes)-2] = byte(0) // TODO Mode
		binEquip[((i+1)*sizeBytes)-1] = game_math.BoolToByte(slot.AmmoReload)
		binEquip[((i + 1) * sizeBytes)] = byte(slotKey)
	}

	return binEquip
}

func (s *Squad) fillStateEquip() {

	s.mx.RLock()
	defer s.mx.RUnlock()

	for k, slot := range s.equipPanel {

		if slot == nil || slot.Source == "" {
			delete(s.equipPanel, k)
			continue
		}

		if k == 0 {
			continue
		}

		if slot.Source == "Constructor" {
			if slot.TypeSlot == 4 {
				weaponSlot := s.GetMS().GetWeaponSlot(slot.Slot)
				if weaponSlot != nil && weaponSlot.Weapon != nil {
					slot.StartReload = weaponSlot.StartReloadTime
					slot.EndReload = weaponSlot.EndReloadTime
					slot.AmmoQuantity = weaponSlot.GetAmmoQuantity()
					slot.AmmoReload = weaponSlot.AmmoReload
					slot.On = weaponSlot.On
				} else {
					//slot.Source = ""
					//slot.TypeSlot = 0
					//slot.Slot = 0
				}
			}

			if slot.TypeSlot == 6 {
				weaponSlot := s.GetMS().GetMeleeWeaponSlot(slot.Slot)
				if weaponSlot != nil && weaponSlot.Weapon != nil {
					slot.AmmoReload = false
					slot.On = weaponSlot.On
				}
			}

			if slot.TypeSlot == 1 || slot.TypeSlot == 2 || slot.TypeSlot == 3 || slot.TypeSlot == 5 {
				equipSlot := s.GetMS().GetBodyEquipSlot(slot.TypeSlot, slot.Slot)
				if equipSlot != nil && equipSlot.Equip != nil {
					slot.StartReload = equipSlot.StartReloadTime
					slot.EndReload = equipSlot.EndReloadTime
					slot.AmmoQuantity = equipSlot.CurrentAmmoCount
					slot.On = equipSlot.On || equipSlot.Drone
					slot.Mode = 0 // TODO equipSlot.Mode
				} else {
					//slot.Source = ""
					//slot.TypeSlot = 0
					//slot.Slot = 0
				}
			}
		}
	}
}

func (s *Squad) SetEquipPanelCell(sell int, source string, typeSlot, slot int) {
	s.mx.Lock()
	defer s.mx.Unlock()

	if sell > 12 || sell < 1 {
		return
	}
	s.equipPanelRerender = true
	s.equipPanel[sell] = &EquipSell{
		Source:   source,
		TypeSlot: typeSlot,
		Slot:     slot,
	}
}

func (s *Squad) SwitchEquipPanelCell(src, dst int) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.equipPanelRerender = true

	srcSlot := s.equipPanel[src]
	dstSlot := s.equipPanel[dst]

	s.equipPanel[src] = dstSlot
	s.equipPanel[dst] = srcSlot
}

func (s *Squad) SetEquipPanelFromJSON(jsonData []byte) {
	s.mx.Lock()
	defer func() {
		s.mx.Unlock()
		s.fillStateEquip()
	}()

	s.equipPanel = make(map[int]*EquipSell)
	_ = json.Unmarshal(jsonData, &s.equipPanel)
	s.equipPanel[0] = &EquipSell{TypeSlot: -1, Source: "empty"}
	s.equipPanelRerender = true
}

func (s *Squad) GetJSONEquipPanel() []byte {

	s.fillStateEquip()

	s.mx.RLock()
	defer s.mx.RUnlock()

	jsonPanel, err := json.Marshal(&s.equipPanel)
	if err != nil {
		println("equip panel to json: ", err.Error())
	}

	return jsonPanel
}

func (s *Squad) GetRerenderEquipPanel() bool {
	defer func() {
		s.equipPanelRerender = false
	}()
	return s.equipPanelRerender
}
