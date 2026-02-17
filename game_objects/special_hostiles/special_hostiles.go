package special_hostiles

import (
	"encoding/json"
	"strconv"
	"sync"
)

// SpecialHostiles отношение бота к другим ботам и игрокм
// 0 - нейтральное отношение, но сохраняется фракционная вражба во фри и батл секторах
// < 0 - дружаня, не атакует по причине фракционной вражды во фри, но атакуте в батл секторах, чаще помогает при запросах
// > 0 - враг, атакует его даже если тот в той же фракции, активно принимает участие в запросах против цели и не помогует цели
type SpecialHostiles struct {
	IgnoreHate      map[string]bool            `json:"ignore_hate"`
	SpecialHostiles map[string]*SpecialHostile `json:"special_hostiles"`
	mx              sync.RWMutex
}

func (s *SpecialHostiles) RangeHostiles() <-chan *SpecialHostile {
	s.mx.RLock()
	ch := make(chan *SpecialHostile, len(s.SpecialHostiles))

	go func() {
		defer func() {
			close(ch)
			s.mx.RUnlock()
		}()

		for _, s := range s.SpecialHostiles {
			ch <- s
		}
	}()

	return ch
}

func (s *SpecialHostiles) UnsafeRangeHostiles() (map[string]*SpecialHostile, *sync.RWMutex) {
	s.mx.RLock()
	return s.SpecialHostiles, &s.mx
}

func (s *SpecialHostiles) getHostile(typeHostile string, id int) *SpecialHostile {

	if typeHostile == "meteorite" {
		return &SpecialHostile{UUID: typeHostile} // отдаем фейк что бы нечего не сломать
	}

	if s.SpecialHostiles == nil {
		s.mx.Lock()
		s.SpecialHostiles = make(map[string]*SpecialHostile)
		s.mx.Unlock()
	}

	uuid := typeHostile + strconv.Itoa(id)

	s.mx.RLock()
	hostile := s.SpecialHostiles[uuid]
	_, ignore := s.IgnoreHate[uuid]
	s.mx.RUnlock()

	if ignore {
		return &SpecialHostile{UUID: uuid} // отдаем фейк что бы нечего не сломать
	}

	if hostile == nil {
		hostile = &SpecialHostile{UUID: uuid, ID: id, Type: typeHostile}
		s.mx.Lock()
		s.SpecialHostiles[uuid] = hostile
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

func (s *SpecialHostiles) AddPoints(typeHostile string, id, hatePoint int, mod string) {
	if s == nil || (typeHostile == "" && id == 0) {
		return
	}

	s.getHostile(typeHostile, id).AddPoints(hatePoint, mod)
}

func (s *SpecialHostiles) CheckHostile(typeHostile string, id int) (bool, int) {
	hostile := s.getHostile(typeHostile, id)
	hatePoints := hostile.GetPoints()

	return hatePoints > 100, hatePoints
}

func (s *SpecialHostiles) CheckHostileByMod(typeHostile string, id int, mod string) int {
	hostile := s.getHostile(typeHostile, id)
	hatePoints := hostile.GetPointsByMod(mod)

	return hatePoints
}

func (s *SpecialHostiles) AddIgnore(typeHostile string, id int) {
	uuid := typeHostile + strconv.Itoa(id)

	s.mx.Lock()
	defer s.mx.Unlock()

	if s.IgnoreHate == nil {
		s.IgnoreHate = make(map[string]bool)
	}

	s.IgnoreHate[uuid] = true
}

func (s *SpecialHostiles) GetJsonData() ([]byte, []byte) {
	s.mx.Lock()
	defer s.mx.Unlock()

	ignoreHate, _ := json.Marshal(s.IgnoreHate)
	specialHostiles, _ := json.Marshal(s.SpecialHostiles)

	return ignoreHate, specialHostiles
}

func (s *SpecialHostiles) LoadFromJson(ignoreHate, specialHostiles []byte) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.IgnoreHate = make(map[string]bool)
	s.SpecialHostiles = make(map[string]*SpecialHostile)

	json.Unmarshal(ignoreHate, &s.IgnoreHate)
	json.Unmarshal(specialHostiles, &s.SpecialHostiles)
}

func (s *SpecialHostiles) LoadData(ignoreHate map[string]bool, specialHostiles map[string]*SpecialHostile) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.IgnoreHate = ignoreHate
	s.SpecialHostiles = specialHostiles
}

func (s *SpecialHostiles) GetData() (map[string]bool, map[string]*SpecialHostile) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	ignoreCopy := make(map[string]bool, len(s.IgnoreHate))
	for k, v := range s.IgnoreHate {
		ignoreCopy[k] = v
	}

	hostilesCopy := make(map[string]*SpecialHostile, len(s.SpecialHostiles))
	for k, v := range s.SpecialHostiles {
		hostilesCopy[k] = v.Copy()
	}

	return ignoreCopy, hostilesCopy
}

func (s *SpecialHostiles) Clear() {
	s.mx.Lock()
	defer s.mx.Unlock()

	for k, h := range s.SpecialHostiles {
		if h.GetPoints() == 0 {
			delete(s.SpecialHostiles, k)
		}
	}
}
