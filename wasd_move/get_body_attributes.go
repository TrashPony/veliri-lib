package wasd_move

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"github.com/TrashPony/veliri-lib/game_objects/physical_model"
	"github.com/TrashPony/veliri-lib/game_objects/unit"
)

func GetBodyMaxSpeedAndTimeRacing(bodies map[int]detail.Body) map[int]map[string]int {
	bodyAttributes := make(map[int]map[string]int)
	for _, b := range bodies {

		u := &unit.Unit{}
		u.SetBody(&b)

		bodyAttributes[b.ID] = GetBodyAttributes(u)
	}

	return bodyAttributes
}

func GetBodyAttributes(u *unit.Unit) map[string]int {

	copypm := &physical_model.PhysicalModel{
		Speed:         u.GetMoveMaxPower() / _const.ServerTickSecPart,
		ReverseSpeed:  u.GetMaxReverse() / _const.ServerTickSecPart,
		PowerFactor:   u.GetPowerFactor() / _const.ServerTickSecPart,
		ReverseFactor: u.GetReverseFactor() / _const.ServerTickSecPart,
		TurnSpeed:     u.GetTurnSpeed() / _const.ServerTickSecPart,
		MoveDrag:      u.GetBody().MoveDrag,
		AngularDrag:   u.GetBody().AngularDrag,
		Weight:        float64(u.GetBody().CapacitySize),
		Type:          "unit",
		ID:            u.GetID(),
		ChassisType:   u.GetBody().ChassisType,
	}

	oldSpeed := 0
	speed := 0
	ticks := 0

	for speed != oldSpeed || speed == 0 {
		oldSpeed = speed
		ticks++
		copypm.SetWASD(true, false, false, false, false, false, false, false, false)
		WasdMove(copypm, nil)
		speed = int(copypm.GetCurrentSpeed() * 1000)
		copypm.MultiplyVelocity(copypm.GetMoveDrag(), copypm.GetMoveDrag())
	}

	for i := 0; i < 900; i++ {
		copypm.SetAngularVelocity(copypm.GetAngularVelocity() + copypm.GetTurnSpeed())
		copypm.SetAngularVelocity(copypm.GetAngularVelocity() * copypm.GetAngularDrag())
	}

	return map[string]int{
		"max_speed":   speed,
		"time_racing": ticks * _const.ServerTick,
		"turn_speed":  int(game_math.RadianToDeg(copypm.GetAngularVelocity()) * _const.ServerTickSecPart),
	}
}
