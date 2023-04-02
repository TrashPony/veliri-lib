package special_hostiles

import (
	"strconv"
	"sync"
	"time"
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
	return hatePoints > 100, hatePoints
}

type SpecialHostile struct {
	UUID                string `json:"uuid"`
	ID                  int    `json:"id"`
	Type                string `json:"type"`
	Points              int    `json:"points"`
	takeAwayPointsEvent map[int64]int
	mx                  sync.RWMutex
}

func (s *SpecialHostile) SetPoints(hatePoint int) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.Points = hatePoint
	s.takeAwayPointsEvent = make(map[int64]int)
}

func (s *SpecialHostile) AddPoints(hatePoint int) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.Points += hatePoint

	if s.Points > 200 {
		hatePoint = 200
	}

	if s.Points < -200 {
		hatePoint = -200
	}

	if s.takeAwayPointsEvent == nil {
		s.takeAwayPointsEvent = make(map[int64]int)
	}

	s.takeAwayPointsEvent[time.Now().UnixNano()] = hatePoint
}

func (s *SpecialHostile) GetPoints() int {
	s.mx.Lock()
	defer s.mx.Unlock()

	if len(s.takeAwayPointsEvent) > 0 {
		// отнимаем все поинты у которых вышло время
		for unixNanoTime, points := range s.takeAwayPointsEvent {
			if time.Now().UnixNano()-unixNanoTime > int64(time.Second*30) {
				s.Points -= points
				delete(s.takeAwayPointsEvent, unixNanoTime)
			}
		}
	}

	return s.Points
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
