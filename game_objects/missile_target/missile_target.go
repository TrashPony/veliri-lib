package missile_target

import (
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/generate_ids"
	"sync"
)

const (
	LockDist     = 45
	LockTime     = 3000 // TODO в параметры пусковой установки, если пустовых несколько брать лучшие показатели
	ProgressTime = 1500 // TODO в параметры пусковой установки
)

type Positioner interface {
	GetX() int
	GetY() int
}

type MissileTarget struct {
	UUID       int
	ID         int
	Type       string
	Dist       int
	Progress   int
	IsActive   bool
	Lock       int
	X          int
	Y          int
	TTL        int
	IsIncoming bool
	Parrent    Positioner
}

type MissileTargetList struct {
	targets []*MissileTarget
	mx      sync.Mutex
}

func (m *MissileTargetList) Add(target *MissileTarget, time int) {
	m.mx.Lock()
	defer m.mx.Unlock()

	if m.targets == nil {
		m.targets = make([]*MissileTarget, 0)
	}

	found := false
	for _, t := range m.targets {
		if target != nil && t.ID == target.ID && t.Type == target.Type {

			t.TTL = target.TTL
			t.X = target.X
			t.Y = target.Y

			if !target.IsIncoming {
				t.IsActive = true
				t.Progress += time
			} else if t.Progress < target.Progress {
				t.Progress = target.Progress
			}

			if t.Progress > ProgressTime {
				t.Progress = ProgressTime
				t.Lock = LockTime
			}

			found = true
		} else if target == nil {
			if t.IsActive && (t.Lock <= 0 || t.ID == -1) {
				t.IsActive = false
			}
		} else {
			t.IsActive = false
		}
	}

	if found {
		target = nil
	}

	if target != nil {
		target.UUID = generate_ids.GetMissileIDGenerate()

		if !target.IsIncoming {
			target.IsActive = true
			target.Progress = time
		}

		m.targets = append(m.targets, target)
	}
}

func (m *MissileTargetList) UpdatePos(target *MissileTarget) {
	for _, t := range m.targets {
		if target != nil && t.ID == target.ID && t.Type == target.Type {
			t.X = target.X
			t.Y = target.Y
		}
	}
}

func (m *MissileTargetList) UpdateTime(time int) {
	m.mx.Lock()
	defer m.mx.Unlock()

	if m.targets == nil {
		m.targets = make([]*MissileTarget, 0)
	}

	for _, t := range m.targets {

		if t.Parrent != nil {
			t.X = t.Parrent.GetX()
			t.Y = t.Parrent.GetY()
		}

		if t.IsIncoming {
			t.TTL -= time
			continue
		}

		if t.Lock >= 0 && t.ID != -1 {
			t.Lock -= time
		} else {
			if !t.IsActive {
				t.Progress -= time / 3
			}
		}
	}

	// remove targets with Progress <= 0
	alive := make([]*MissileTarget, 0, len(m.targets))
	for _, t := range m.targets {
		if (!t.IsIncoming && t.Progress >= 0) || (t.IsIncoming && t.TTL > 0) {
			alive = append(alive, t)
		}
	}

	m.targets = alive
}

func (m *MissileTargetList) GetActive() *MissileTarget {
	for _, t := range m.targets {
		if t.IsActive && t.Progress >= ProgressTime {
			return t
		}
	}

	return nil
}

func (m *MissileTargetList) GetBinDataLockTargets() []byte {
	m.mx.Lock()
	defer m.mx.Unlock()

	data := make([]byte, 0)

	for _, t := range m.targets {
		data = append(data, game_math.GetIntBytes(t.UUID)...)
		data = append(data, game_math.BoolToByte(t.IsActive))
		data = append(data, byte((float64(t.Progress)/float64(ProgressTime))*100))
		data = append(data, game_math.GetIntBytes(t.X)...)
		data = append(data, game_math.GetIntBytes(t.Y)...)
		data = append(data, game_math.BoolToByte(t.IsIncoming))

		if t.IsIncoming {
			t.Progress = 0
		}
	}

	return data
}
