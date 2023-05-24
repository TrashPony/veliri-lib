package special_hostiles

import (
	"encoding/json"
	"sync"
	"time"
)

type SpecialHostile struct {
	UUID       string `json:"uuid"`
	ID         int    `json:"id"`
	Type       string `json:"type"`
	Points     int    `json:"points"`
	LastUpdate int64  `json:"last_update"`
	mx         sync.RWMutex
}

func (s *SpecialHostile) SetPoints(hatePoint int) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.Points = hatePoint
}

func (s *SpecialHostile) AddPoints(hatePoint int) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.Points += hatePoint

	if s.Points > 200 {
		s.Points = 200
	}

	if s.Points < -200 {
		s.Points = -200
	}

	s.LastUpdate = time.Now().UnixNano()
}

func (s *SpecialHostile) GetPoints() int {
	s.mx.Lock()
	defer s.mx.Unlock()

	if time.Now().UnixNano()-s.LastUpdate > int64(time.Second*30) {
		s.Points = 0
	}

	return s.Points
}

func (s *SpecialHostile) GetJSON() []byte {
	s.mx.Lock()
	defer s.mx.Unlock()

	jsonData, _ := json.Marshal(s)
	return jsonData
}
