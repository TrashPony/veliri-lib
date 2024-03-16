package special_hostiles

import (
	"encoding/json"
	"sync"
	"time"
)

const (
	max                = 200
	min                = -200
	shortMemoryTime    = 30
	fractionUpdateTime = 900
)

type SpecialHostile struct {
	UUID       string                     `json:"uuid"`
	ID         int                        `json:"id"`
	Type       string                     `json:"type"`
	Points     int                        `json:"points"`         // коротковременная память
	Moderate   map[string]*PointsModerate `json:"pointsModerate"` // долговременная память
	LastUpdate int64                      `json:"last_update"`
	mx         sync.RWMutex
}

func (s *SpecialHostile) SetPoints(hatePoint int) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.Points = hatePoint
	s.LastUpdate = time.Now().UnixNano()
}

func (s *SpecialHostile) AddPoints(hatePoint int, mod string) {
	s.mx.Lock()
	defer s.mx.Unlock()

	if s.Moderate == nil {
		s.Moderate = make(map[string]*PointsModerate)
	}

	if mod == "" || mod == "battle" {
		// кратковременная память живет всего 30 секунд
		s.Points += hatePoint

		if s.Points > max {
			s.Points = max
		}

		if s.Points < min {
			s.Points = min
		}

		s.LastUpdate = time.Now().UnixNano()
		return
	}

	if mod == "fraction" {
		m, ok := s.Moderate[mod]
		if ok && time.Now().UnixNano()-m.LastUpdate < int64(time.Second*fractionUpdateTime) {
			// фракционный мод мы считывает только когда видим цель первый раз
			return
		}

		s.Moderate[mod] = getModerate(mod)
		s.Moderate[mod].SetPoints(hatePoint)
		return
	}

	_, ok := s.Moderate[mod]
	if !ok {
		s.Moderate[mod] = getModerate(mod)
	}

	if s.Moderate[mod] != nil {
		s.Moderate[mod].AddPoints(hatePoint)
	}
}

func (s *SpecialHostile) GetPoints() int {
	s.mx.Lock()
	defer s.mx.Unlock()

	if time.Now().UnixNano()-s.LastUpdate > int64(time.Second*shortMemoryTime) {
		s.Points = 0
	}

	allPoints := s.Points

	for _, m := range s.Moderate {
		allPoints += m.GetPoints()
	}

	if allPoints > max {
		allPoints = max
	}

	if allPoints < min {
		allPoints = min
	}

	return allPoints
}

func (s *SpecialHostile) GetJSON() []byte {
	s.mx.Lock()
	defer s.mx.Unlock()

	jsonData, _ := json.Marshal(s)
	return jsonData
}
