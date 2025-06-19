package move_path

import (
	"math"
	"sync"
	"time"

	"github.com/TrashPony/veliri-lib/game_objects/coordinate"
	"github.com/TrashPony/veliri-lib/game_objects/target"
)

type MovePath struct {
	typeFind     string
	angle        float64
	path         *[]*coordinate.Coordinate
	followTarget *target.Target
	currentPoint int
	distToPoint  int
	needFindPath bool
	time         int64
	playerID     int
	unitID       int
	stop         bool
	typePath     int
	mx           sync.RWMutex
}

func (m *MovePath) GetMovePathTime() int64 {
	return m.time
}

func (m *MovePath) GetNeedCalc() bool {
	return m.needFindPath
}

func (m *MovePath) GetMovePathState() (bool, string, float64, *target.Target, *[]*coordinate.Coordinate, int, bool, int64, bool, int, int) {
	m.mx.RLock()
	defer m.mx.RUnlock()

	return true, m.typeFind, m.angle, m.followTarget, m.path, m.currentPoint, m.needFindPath, m.time, m.stop, m.distToPoint, m.typePath
}

func (m *MovePath) NextMovePoint() {
	m.currentPoint++
	m.distToPoint = math.MaxInt
}

func (m *MovePath) SetDistToTaget(d int) {
	m.distToPoint = d
}

func (m *MovePath) SetFindMovePath() {
	m.needFindPath = true
}

func (m *MovePath) SetMovePath(path *[]*coordinate.Coordinate, typePath int) {
	m.mx.Lock()
	defer m.mx.Unlock()

	m.needFindPath = false
	m.path = path
	m.currentPoint = 0
	m.stop = false
	m.time = time.Now().Unix()
	m.typePath = typePath
	m.distToPoint = math.MaxInt
}

func (m *MovePath) GetMovePath() []*coordinate.Coordinate {
	m.mx.RLock()
	defer m.mx.RUnlock()

	if m.path == nil {
		return nil
	}

	path := make([]*coordinate.Coordinate, 0, len(*m.path))

	for i, p := range *m.path {
		if i >= m.currentPoint {
			path = append(path, p)
		}
	}

	return path
}

func (m *MovePath) GetFollowTarget() *target.Target {
	return m.followTarget
}

func (m *MovePath) SetMovePathTarget(t *target.Target, playerID, unitID int) {
	m.mx.Lock()
	defer m.mx.Unlock()

	m.needFindPath = true
	m.path = nil // &[]*coordinate.Coordinate{{X: t.X, Y: t.Y}}
	m.followTarget = t
	m.playerID = playerID
	m.unitID = unitID
}

func (m *MovePath) SetMovePathAngle(angle float64, playerID, unitID int) {
	m.typeFind = "angle"
	m.needFindPath = true
	m.angle = angle
	m.playerID = playerID
	m.unitID = unitID
	m.followTarget = &target.Target{Ignore: true}
}

func (m *MovePath) SetMovePathRotate(angle float64, playerID, unitID int) {
	m.typeFind = "rotate"
	m.needFindPath = true
	m.angle = angle
	m.playerID = playerID
	m.unitID = unitID
	m.followTarget = &target.Target{Ignore: true}
}

func (m *MovePath) Stop() {
	m.stop = true
}
