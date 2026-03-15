package unit

import (
	"github.com/TrashPony/veliri-lib/game_objects/aim_filter"
)

func (u *Unit) GetAimFilter() *aim_filter.AimFilter {
	if u.aimFilter == nil {
		u.aimFilter = aim_filter.NewAimFilter()
	}

	return u.aimFilter
}
