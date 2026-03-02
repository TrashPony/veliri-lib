package wasd_move

import (
	"github.com/TrashPony/veliri-lib/game_math"
	"math"
)

const startCaterpillarsSpeedK = 5

func caterpillars(obj MoveObject) {

	if obj.CheckGrowthPower() > 0 {
		if obj.GetPowerMove() < obj.GetMoveMaxPower()/startCaterpillarsSpeedK {
			obj.SetPowerMove(obj.GetMoveMaxPower() / startCaterpillarsSpeedK)
		} else {
			obj.SetPowerMove(obj.GetPowerMove() + obj.GetPowerFactor())
		}
	} else {
		obj.SetPowerMove(obj.GetPowerMove() - obj.GetPowerFactor())
	}

	if obj.CheckGrowthRevers() > 0 {
		if obj.GetReverse() < obj.GetMaxReverse()/startCaterpillarsSpeedK {
			obj.SetReverse(obj.GetMaxReverse() / startCaterpillarsSpeedK)
		} else {
			obj.SetReverse(obj.GetReverse() + obj.GetReverseFactor())
		}
	} else {
		obj.SetReverse(obj.GetReverse() - obj.GetReverseFactor())
	}

	obj.SetPowerMove(math.Max(0, math.Min(obj.GetMoveMaxPower(), obj.GetPowerMove())))
	obj.SetReverse(math.Max(0, math.Min(obj.GetMaxReverse(), obj.GetReverse())))

	// ручной тормаз
	if obj.CheckHandBrake() {

		if obj.GetPowerMove() > obj.GetReverse() {
			obj.SetPowerMove(obj.GetPowerMove() - obj.GetPowerMove()/8)
		} else {
			obj.SetReverse(obj.GetReverse() - obj.GetReverse()/8)
		}

		if obj.GetPowerMove() <= obj.GetMoveMaxPower()/startCaterpillarsSpeedK {
			obj.SetPowerMove(0)
		}

		if obj.GetReverse() <= obj.GetMaxReverse()/startCaterpillarsSpeedK {
			obj.SetReverse(0)
		}
	}

	direction := 1.0
	if obj.GetPowerMove() < obj.GetReverse() {
		direction = -1
	}

	move := obj.GetPowerMove() > 0 || obj.GetReverse() > 0

	if obj.GetChassisType() == "caterpillars" || obj.GetChassisType() == "fly" || ((obj.GetChassisType() == "wheels" || obj.GetChassisType() == "fly-2") && move) {
		if obj.CheckLeftRotate() > 0 {
			obj.SetAngularVelocity(obj.GetAngularVelocity() - (direction * getPercentF(obj.GetTurnSpeed(), obj.CheckLeftRotate())))
		}

		if obj.CheckRightRotate() > 0 {
			obj.SetAngularVelocity(obj.GetAngularVelocity() + (direction * getPercentF(obj.GetTurnSpeed(), obj.CheckRightRotate())))
		}
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
