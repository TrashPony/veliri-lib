package move_path

import (
	"github.com/TrashPony/veliri-lib/game_objects/coordinate"
	"github.com/TrashPony/veliri-lib/game_objects/target"
	"time"
)

type MovePath struct {
	typeFind     string
	angle        float64
	path         *[]*coordinate.Coordinate
	followTarget *target.Target
	currentPoint int
	needFindPath bool
	time         int64
	playerID     int
	unitID       int
	stop         int
	stopTimeOut  int64
}

func (m *MovePath) GetMovePathTime() int64 {
	return m.time
}

func (m *MovePath) GetNeedCalc() bool {
	return m.needFindPath
}

func (m *MovePath) GetMovePathState() (bool, string, float64, *target.Target, *[]*coordinate.Coordinate, int, bool, int64, bool) {
	stop := m.stop > 0
	if m.stopTimeOut < time.Now().Unix() {
		if m.stop >= 2 {
			m.needFindPath = true
		}

		stop = false
	}

	return true, m.typeFind, m.angle, m.followTarget, m.path, m.currentPoint, m.needFindPath, m.time, stop
}

func (m *MovePath) NextMovePoint() {
	m.currentPoint++
}

func (m *MovePath) SetFindMovePath() {
	m.needFindPath = true
}

func (m *MovePath) SetMovePath(path *[]*coordinate.Coordinate) {
	m.needFindPath = false
	m.path = path
	m.currentPoint = 0
	m.stop = 0
	m.stopTimeOut = 0
	m.time = time.Now().Unix()
}

func (m *MovePath) GetFollowTarget() *target.Target {
	return m.followTarget
}

func (m *MovePath) SetMovePathTarget(t *target.Target, playerID, unitID int) {
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

func (m *MovePath) Stop() {
	m.stop++
	m.stopTimeOut = time.Now().Unix() + 5
}
