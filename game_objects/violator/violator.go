package violator

import (
	"sync"
	"time"
)

type Violator struct {
	PlayerID      int    `json:"player_id"`
	Type          string `json:"type"`  // (pvp/pk)
	Price         int    `json:"price"` // награда за голову, пока она всегда на 1 убийство
	EndTime       int64  `json:"end_time"`
	Strikes       int    `json:"strikes"`
	TakeStrike    int64  `json:"take_strike"`
	TimeOutStrike int64  `json:"time_out_strike"`
	Fraction      string `json:"fraction"`
}

func (v *Violator) GetTime() int {
	if v.EndTime == -1 {
		return 99999
	}

	return int(v.EndTime - time.Now().Unix())
}

func (v *Violator) TakeStrikeTime() int {
	return int(v.TakeStrike - time.Now().Unix())
}

type DistressSignals struct {
	UUID     string `json:"-"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	MapID    int    `json:"map_id"`
	Fraction string `json:"-"`
	// содержит ид ботов которые посетили аномалию
	Bots        map[int]bool `json:"-"`
	DestroyTime int64        `json:"-"`
	mx          sync.Mutex   `json:"-"`
}

func (s *DistressSignals) CheckPassBot(botID int) bool {
	s.mx.Lock()
	defer s.mx.Unlock()
	return s.Bots[botID]
}

func (s *DistressSignals) AddPassBot(botID int) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.Bots[botID] = true
}
