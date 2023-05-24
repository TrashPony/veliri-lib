package special_hostiles

import (
	"encoding/json"
	"strconv"
	"sync"
)

// отношение бота к другим ботам и игрокм
// 0   - нейтральное отношение, но сохраняется фракционная вражба во фри и батл секторах
// < 0 - дружаня, не атакует по причине фракционной вражды во фри, но атакуте в батл секторах, чаще помогает при запросах
// > 0 - враг, атакует его даже если тот в той же фракции, активно принимает участие в запросах против цели и не помогует цели
type SpecialHostiles struct {
	ignoreHate      map[string]bool
	specialHostiles map[string]*SpecialHostile
	mx              sync.RWMutex
}

func (s *SpecialHostiles) RangeHostiles() <-chan *SpecialHostile {
	s.mx.RLock()
	ch := make(chan *SpecialHostile, len(s.specialHostiles))

	go func() {
		defer func() {
			close(ch)
			s.mx.RUnlock()
		}()

		for _, s := range s.specialHostiles {
			ch <- s
		}
	}()

	return ch
}

func (s *SpecialHostiles) getHostile(typeHostile string, id int) *SpecialHostile {

	if typeHostile == "meteorite" {
		return &SpecialHostile{UUID: typeHostile} // отдаем фейк что бы нечего не сломать
	}

	if s.specialHostiles == nil {
		s.mx.Lock()
		s.specialHostiles = make(map[string]*SpecialHostile)
		s.mx.Unlock()
	}

	uuid := typeHostile + strconv.Itoa(id)

	s.mx.RLock()
	hostile := s.specialHostiles[uuid]
	_, ignore := s.ignoreHate[uuid]
	s.mx.RUnlock()

	if ignore {
		return &SpecialHostile{UUID: uuid} // отдаем фейк что бы нечего не сломать
	}

	if hostile == nil {
		hostile = &SpecialHostile{UUID: uuid, ID: id, Type: typeHostile}
		s.mx.Lock()
		s.specialHostiles[uuid] = hostile
		s.mx.Unlock()
	}

	return hostile
}

func (s *SpecialHostiles) SetPoints(typeHostile string, id, hatePoint int) {
	if s == nil || (typeHostile == "" && id == 0) {
		return
	}

	s.getHostile(typeHostile, id).SetPoints(hatePoint)
}

func (s *SpecialHostiles) AddPoints(typeHostile string, id, hatePoint int) {
	if s == nil || (typeHostile == "" && id == 0) {
		return
	}

	s.getHostile(typeHostile, id).AddPoints(hatePoint)
}

func (s *SpecialHostiles) CheckHostile(typeHostile string, id int) (bool, int) {
	hostile := s.getHostile(typeHostile, id)
	hatePoints := hostile.GetPoints()

	if hatePoints == 0 {
		s.mx.Lock()
		delete(s.specialHostiles, hostile.UUID)
		s.mx.Unlock()
	}

	return hatePoints > 100, hatePoints
}

func (s *SpecialHostiles) AddIgnore(typeHostile string, id int) {
	uuid := typeHostile + strconv.Itoa(id)

	s.mx.Lock()
	defer s.mx.Unlock()

	if s.ignoreHate == nil {
		s.ignoreHate = make(map[string]bool)
	}

	s.ignoreHate[uuid] = true
}

func (s *SpecialHostiles) GetJsonData() ([]byte, []byte) {
	s.mx.Lock()
	defer s.mx.Unlock()

	ignoreHate, _ := json.Marshal(s.ignoreHate)
	specialHostiles, _ := json.Marshal(s.specialHostiles)

	return ignoreHate, specialHostiles
}

func (s *SpecialHostiles) LoadFromJson(ignoreHate, specialHostiles []byte) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.ignoreHate = make(map[string]bool)
	s.specialHostiles = make(map[string]*SpecialHostile)

	json.Unmarshal(ignoreHate, &s.ignoreHate)
	json.Unmarshal(specialHostiles, &s.specialHostiles)
}
