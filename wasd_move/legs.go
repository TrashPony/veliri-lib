package wasd_move

import "github.com/TrashPony/veliri-lib/game_math"

func legs(obj MoveObject) {

	if obj.CheckGrowthPower() > 0 {
		obj.SetPowerMove(obj.GetMoveMaxPower())
	} else {
		obj.SetPowerMove(0)
	}

	if obj.CheckGrowthRevers() > 0 {
		obj.SetReverse(obj.GetMaxReverse())
	} else {
		obj.SetReverse(0)
	}

	direction := 1.0
	if obj.GetPowerMove() < obj.GetReverse() {
		direction = -1
	}

	if obj.CheckLeftRotate() > 0 {
		obj.SetAngularVelocity(obj.GetAngularVelocity() - (direction * obj.GetTurnSpeed()))
	}

	if obj.CheckRightRotate() > 0 {
		obj.SetAngularVelocity(obj.GetAngularVelocity() + (direction * obj.GetTurnSpeed()))
	}

	radRotate := game_math.DegToRadian(obj.GetRotate())
	obj.AddVelocity(
		game_math.Cos(radRotate)*(obj.GetPowerMove()-obj.GetReverse()),
		game_math.Sin(radRotate)*(obj.GetPowerMove()-obj.GetReverse()),
	)

	xV, yV := obj.GetVelocity()
	xR, yR := obj.GetRealPos()

	obj.SetNextPos(xR+xV, yR+yV)
}
