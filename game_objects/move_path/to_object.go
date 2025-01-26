package move_path

import (
	"github.com/TrashPony/veliri-lib/game_math"
	"time"
)

type To struct {
	Find    bool  `json:"find"`
	Path    bool  `json:"path"`
	Time    int64 `json:"time"`
	Source  game_math.Positions
	FindPos game_math.Positions
}

func (t *To) Check(sourceX, sourceY int) (bool, bool) {
	if t.Source.X != sourceX || t.Source.Y != sourceY {
		return false, true
	}

	if t.Time > 0 && (time.Now().Unix()-t.Time > 5) {
		return t.Path && t.Find, true
	}

	return t.Path && t.Find, !t.Find
}

func (t *To) SetCheck(path bool, sourceX, sourceY int, findX, findY int) {
	t.Find = true
	t.Source.X = sourceX
	t.Source.Y = sourceY
	t.FindPos.X = findX
	t.FindPos.Y = findY
	t.Path = path
	t.Time = time.Now().Unix()
}
