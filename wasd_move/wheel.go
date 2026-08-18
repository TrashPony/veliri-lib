package wasd_move

import (
	"github.com/TrashPony/veliri-lib/game_math"
	"math"
)

const startWheelSpeedK = 3

func wheel(obj MoveObject) {

	pm := obj.GetPhysicalModel()

	massK := 8000.0 / pm.GetWeight()
	massK = math.Max(0.3, math.Min(1.2, massK))

	if obj.CheckGrowthPower() > 0 {
		if obj.GetPowerMove() < obj.GetMoveMaxPower()/startWheelSpeedK {
			obj.SetPowerMove(obj.GetMoveMaxPower() / startWheelSpeedK)
		} else {
			obj.SetPowerMove(obj.GetPowerMove() + obj.GetPowerFactor())
		}
	} else {
		obj.SetPowerMove(obj.GetPowerMove() - obj.GetPowerFactor())
	}

	if obj.CheckGrowthRevers() > 0 {
		if obj.GetReverse() < obj.GetMaxReverse()/startWheelSpeedK {
			obj.SetReverse(obj.GetMaxReverse() / startWheelSpeedK)
		} else {
			obj.SetReverse(obj.GetReverse() + obj.GetReverseFactor())
		}
	} else {
		obj.SetReverse(obj.GetReverse() - obj.GetReverseFactor())
	}

	obj.SetPowerMove(math.Max(0, math.Min(obj.GetMoveMaxPower(), obj.GetPowerMove())))
	obj.SetReverse(math.Max(0, math.Min(obj.GetMaxReverse(), obj.GetReverse())))

	if obj.CheckHandBrake() {
		if obj.CheckGrowthPower() > 0 {
			if obj.GetPowerMove() > obj.GetReverse() {
				obj.SetPowerMove(obj.GetPowerMove() - obj.GetPowerMove()/40)
			} else {
				obj.SetReverse(obj.GetReverse() - obj.GetReverse()/40)
			}
		} else {
			if obj.GetPowerMove() > obj.GetReverse() {
				obj.SetPowerMove(obj.GetPowerMove() - obj.GetPowerMove()/10)
			} else {
				obj.SetReverse(obj.GetReverse() - obj.GetReverse()/10)
			}
		}

		if obj.CheckGrowthPower() > 0 || obj.CheckGrowthRevers() > 0 {
			if obj.CheckGrowthPower() > 0 {
				if obj.GetPowerMove() <= obj.GetMoveMaxPower()/startWheelSpeedK {
					obj.SetPowerMove(obj.GetMoveMaxPower() / (startWheelSpeedK * 2))
				}
			}
			if obj.CheckGrowthRevers() > 0 {
				if obj.GetReverse() <= obj.GetMaxReverse()/startWheelSpeedK {
					obj.SetReverse(obj.GetMaxReverse() / (startWheelSpeedK * 2))
				}
			}
		} else {
			if obj.GetPowerMove() <= obj.GetMoveMaxPower()/startWheelSpeedK {
				obj.SetPowerMove(0)
			}
			if obj.GetReverse() <= obj.GetMaxReverse()/startWheelSpeedK {
				obj.SetReverse(0)
			}
		}
	}

	rad := game_math.DegToRadian(pm.GetRotate())
	cosRad := game_math.Cos(rad)
	sinRad := game_math.Sin(rad)

	forwardAccel := pm.GetPowerMove() - pm.GetReverse()

	pm.XVelocity += cosRad * forwardAccel
	pm.YVelocity += sinRad * forwardAccel

	currentSpeed := math.Hypot(pm.XVelocity, pm.YVelocity)

	inertialSlip := math.Abs(pm.GetAngularVelocity()) * currentSpeed * 0.75
	if inertialSlip > 0.9 {
		inertialSlip = 0.9
	}

	driftFactor := inertialSlip

	if pm.CheckHandBrake() && (pm.CheckLeftRotate() > 0 || pm.CheckRightRotate() > 0) {
		if driftFactor < 0.9 {
			driftFactor = 0.9
		}
	}

	driftAccelX, driftAccelY := 0.0, 0.0

	if driftFactor > 0.05 {
		driftDir := 0.0
		if pm.CheckLeftRotate() > 0 {
			driftDir = 1
		}
		if pm.CheckRightRotate() > 0 {
			driftDir = -1
		}

		if driftDir == 0 {
			if pm.GetAngularVelocity() > 0 {
				driftDir = -1
			} else if pm.GetAngularVelocity() < 0 {
				driftDir = 1
			}
		}

		as := math.Abs(pm.GetAngularVelocity() * 5)
		baseDrift := as * massK * 8.0
		if pm.CheckHandBrake() {
			baseDrift *= 1.5
		}

		driftAccelX = -sinRad * driftDir * baseDrift * currentSpeed * driftFactor
		driftAccelY = cosRad * driftDir * baseDrift * currentSpeed * driftFactor
	}

	driftDecay := 0.90
	if driftFactor > 0.3 {
		driftDecay = 0.96
	}

	pm.DriftX = pm.DriftX*driftDecay + driftAccelX*(0.01*massK)
	pm.DriftY = pm.DriftY*driftDecay + driftAccelY*(0.01*massK)

	driftMag := math.Hypot(pm.DriftX, pm.DriftY)
	if driftMag > 20 {
		pm.DriftX = pm.DriftX / driftMag * 20
		pm.DriftY = pm.DriftY / driftMag * 20
	}

	currentForwardSpeed := pm.XVelocity*cosRad + pm.YVelocity*sinRad

	direction := 1.0
	if currentForwardSpeed < 0 {
		direction = -1.0
	}

	if currentSpeed > 0.5 || pm.GetPowerMove() > 0 || pm.GetReverse() > 0 {

		ts := pm.GetTurnSpeed()

		// 3. Снижение чувствительности от скорости (currentSpeed)
		// speedThreshold: скорость, при которой чувствительность начнет заметно падать.
		speedThreshold := 20.0

		speedRatio := currentSpeed / speedThreshold
		if speedRatio > 1.0 {
			speedRatio = 1.0 // Ограничиваем максимум, чтобы руль не инвертировался
		}

		ts *= (1.0 - speedRatio*0.3)

		if direction < 0 {
			ts = pm.GetTurnSpeed()
		}

		ts = math.Max(pm.GetTurnSpeed()*0.7, ts)

		if pm.CheckLeftRotate() > 0 {
			pm.SetAngularVelocity(pm.GetAngularVelocity() - direction*getPercentF(ts, pm.CheckLeftRotate()))
		}
		if pm.CheckRightRotate() > 0 {
			pm.SetAngularVelocity(pm.GetAngularVelocity() + direction*getPercentF(ts, pm.CheckRightRotate()))
		}
	}

	totalVX := pm.XVelocity + pm.DriftX
	totalVY := pm.YVelocity + pm.DriftY

	forwardSpeed := totalVX*cosRad + totalVY*sinRad
	lateralSpeed := -totalVX*sinRad + totalVY*cosRad

	forwardDrag := obj.GetMoveDrag()

	baseLateralGrip := 0.2
	maxLateralSlip := forwardDrag

	lateralDrag := baseLateralGrip + (maxLateralSlip-baseLateralGrip)*driftFactor

	forwardSpeed *= forwardDrag
	lateralSpeed *= lateralDrag

	pm.XVelocity = forwardSpeed*cosRad - lateralSpeed*sinRad
	pm.YVelocity = forwardSpeed*sinRad + lateralSpeed*cosRad

	pm.DriftX *= driftFactor * 0.7
	pm.DriftY *= driftFactor * 0.7

	obj.SetAngularVelocity(obj.GetAngularVelocity() * obj.GetAngularDrag())

	totalX := pm.XVelocity
	totalY := pm.YVelocity

	xR, yR := pm.GetRealPos()
	pm.SetNextPos(xR+totalX, yR+totalY)
}
