package wasd_move

import (
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/gunner"
	"math"
)

func antigravity(obj MoveObject, g *gunner.Gunner) {

	if obj.CheckGrowthPower() {
		obj.SetPowerMove(obj.GetPowerMove() + obj.GetPowerFactor())
	} else {
		obj.SetPowerMove(0)
	}

	if obj.CheckGrowthRevers() {
		obj.SetReverse(obj.GetReverse() + obj.GetReverseFactor())
	} else {
		obj.SetReverse(0)
	}

	if obj.CheckLeftRotate() {
		obj.SetPowerLeft(obj.GetPowerLeft() + obj.GetPowerFactor())
	} else {
		obj.SetPowerLeft(0)
	}

	if obj.CheckRightRotate() {
		obj.SetPowerRight(obj.GetPowerRight() + obj.GetPowerFactor())
	} else {
		obj.SetPowerRight(0)
	}

	obj.SetPowerMove(math.Max(0, math.Min(obj.GetMoveMaxPower(), obj.GetPowerMove())))
	obj.SetReverse(math.Max(0, math.Min(obj.GetMaxReverse(), obj.GetReverse())))
	obj.SetPowerLeft(math.Max(0, math.Min(obj.GetMoveMaxPower(), obj.GetPowerLeft())))
	obj.SetPowerRight(math.Max(0, math.Min(obj.GetMoveMaxPower(), obj.GetPowerRight())))

	radRotate := game_math.DegToRadian(obj.GetRotate())
	obj.AddVelocity(
		game_math.Cos(radRotate)*(obj.GetPowerMove()-obj.GetReverse()),
		game_math.Sin(radRotate)*(obj.GetPowerMove()-obj.GetReverse()),
	)

	radRotate2 := game_math.DegToRadian(obj.GetRotate() - 90)
	obj.AddVelocity(
		game_math.Cos(radRotate2)*(obj.GetPowerLeft()-obj.GetPowerRight()),
		game_math.Sin(radRotate2)*(obj.GetPowerLeft()-obj.GetPowerRight()),
	)

	if obj.GetTypeControl() == 0 {
		wt := g.GetWeaponTarget(0)
		if wt != nil {
			needAngle := game_math.GetBetweenAngle(float64(wt.GetX()), float64(wt.GetY()), float64(obj.GetX()), float64(obj.GetY()))
			diffAngle := game_math.ShortestBetweenAngle(obj.GetRotate(), needAngle)

			if diffAngle > 2 {
				obj.SetAngularVelocity(obj.GetAngularVelocity() + obj.GetTurnSpeed())
			}

			if diffAngle < -2 {
				obj.SetAngularVelocity(obj.GetAngularVelocity() - obj.GetTurnSpeed())
			}
		}
	} else {
		if obj.GetAngleForClassicControl() >= 0 {
			diffAngle := game_math.ShortestBetweenAngle(obj.GetRotate(), obj.GetAngleForClassicControl())
			if diffAngle > 2 {
				obj.SetAngularVelocity(obj.GetAngularVelocity() + obj.GetTurnSpeed())
			}

			if diffAngle < -2 {
				obj.SetAngularVelocity(obj.GetAngularVelocity() - obj.GetTurnSpeed())
			}
		}
	}

	xV, yV := obj.GetVelocity()
	xR, yR := obj.GetRealPos()

	obj.SetNextPos(xR+xV, yR+yV)
}
