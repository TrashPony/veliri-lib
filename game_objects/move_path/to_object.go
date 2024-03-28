package move_path

import "github.com/TrashPony/veliri-lib/game_math"

type To struct {
	Find   bool `json:"find"`
	Path   bool `json:"path"`
	Source game_math.Positions
}

func (t *To) Check(sourceX, sourceY int) (bool, bool) {
	if t.Source.X != sourceX || t.Source.Y != sourceY {
		return false, true
	}

	return t.Path && t.Find, !t.Find
}

func (t *To) SetCheck(path bool, sourceX, sourceY int) {
	t.Find = true
	t.Source.X = sourceX
	t.Source.Y = sourceY
	t.Path = path
}
