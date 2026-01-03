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

	if obj.CheckGrowthPower() {
		if obj.GetPowerMove() < obj.GetMoveMaxPower()/startWheelSpeedK {
			obj.SetPowerMove(obj.GetMoveMaxPower() / startWheelSpeedK)
		} else {
			obj.SetPowerMove(obj.GetPowerMove() + obj.GetPowerFactor())
		}
	} else {
		obj.SetPowerMove(obj.GetPowerMove() - obj.GetPowerFactor())
	}

	if obj.CheckGrowthRevers() {
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

	// ручной тормаз
	if obj.CheckHandBrake() {
		if obj.CheckGrowthPower() {
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

		if obj.GetPowerMove() <= obj.GetMoveMaxPower()/startWheelSpeedK {
			obj.SetPowerMove(obj.GetMoveMaxPower() / startWheelSpeedK / 2)
		}

		if obj.GetReverse() <= obj.GetMaxReverse()/startWheelSpeedK {
			obj.SetReverse(obj.GetMaxReverse() / startWheelSpeedK / 2)
		}
	}

	rad := game_math.DegToRadian(pm.GetRotate())
	forwardAccel := pm.GetPowerMove() - pm.GetReverse()

	// === 1. Ускорение вперёд (вдоль корпуса) → добавляем к ОСНОВНОЙ скорости
	pm.XVelocity += game_math.Cos(rad) * forwardAccel
	pm.YVelocity += game_math.Sin(rad) * forwardAccel

	// === 2. Рассчитываем дрифт-импульс ===
	driftAccelX, driftAccelY := 0.0, 0.0

	// Дрифт активен, если:
	isDrifting := false
	if pm.CheckHandBrake() && (pm.CheckLeftRotate() || pm.CheckRightRotate()) {
		isDrifting = true
	} else {
		// Инерционный дрифт: если угловая скорость велика, а продольная скорость мала
		currentSpeed := math.Hypot(pm.XVelocity, pm.YVelocity)
		if math.Abs(pm.GetAngularVelocity()) > 0.001 && currentSpeed > 5 {
			isDrifting = true
		}
	}

	if isDrifting {
		// Направление дрифта: перпендикулярно корпусу!
		// При повороте влево → зад уходит вправо → дрифт вправо (и наоборот)
		driftDir := 0.0
		if pm.CheckLeftRotate() {
			driftDir = 1
		} // вправо!
		if pm.CheckRightRotate() {
			driftDir = -1
		} // влево!

		// Сила дрифта: зависит от скорости и резкости
		as := (pm.GetAngularVelocity() * 8)
		if as < 0 {
			as *= -1
		}

		baseDrift := obj.GetCurrentSpeed() * as * massK
		if pm.CheckHandBrake() {
			baseDrift *= 2.0 // ручной тормоз = мощный дрифт
		}

		driftAccelX = -game_math.Sin(rad) * driftDir * baseDrift
		driftAccelY = game_math.Cos(rad) * driftDir * baseDrift
	}

	// === 3. Обновляем Drift-вектор (с инерцией!) ===
	// Новый импульс добавляется, старый — затухает
	cs := obj.GetCurrentSpeed() * 0.008 * massK
	pm.DriftX = pm.DriftX*0.9 + driftAccelX*cs
	pm.DriftY = pm.DriftY*0.9 + driftAccelY*cs

	// Ограничиваем, чтобы не улетал в космос
	driftMag := math.Hypot(pm.DriftX, pm.DriftY)
	if driftMag > 20 { // подбери под масштаб (20 юнит/сек — много)
		pm.DriftX = pm.DriftX / driftMag * 20
		pm.DriftY = pm.DriftY / driftMag * 20
	}

	// === 4. Итоговая скорость = основная + дрифт ===
	totalX := pm.XVelocity + pm.DriftX
	totalY := pm.YVelocity + pm.DriftY

	// === 5. Поворот (оставь как есть, но можно добавить grip-ослабление позже) ===
	direction := 1.0
	if pm.GetPowerMove() < pm.GetReverse() {
		direction = -1
	}

	if (obj.GetPowerMove() > 0 && direction > 0) || (obj.GetReverse() > 0 && direction < 0) {

		speedK := obj.GetPowerMove() / obj.GetMoveMaxPower()
		ts := direction * pm.GetTurnSpeed() * speedK
		if direction < 0 {
			// назад большая чувствительность ts = direction * pm.GetTurnSpeed() * speedK
			ts = pm.GetTurnSpeed()
		}

		if ts > pm.GetTurnSpeed() {
			ts = pm.GetTurnSpeed()
		}

		if ts < pm.GetTurnSpeed()/5 {
			ts = pm.GetTurnSpeed() / 5
		}

		if pm.CheckLeftRotate() {
			pm.SetAngularVelocity(pm.GetAngularVelocity() - direction*ts)
		}
		if pm.CheckRightRotate() {
			pm.SetAngularVelocity(pm.GetAngularVelocity() + direction*ts)
		}
	}

	// === 6. Следующая позиция — по ИТОГОВОЙ скорости ===
	xR, yR := pm.GetRealPos()
	pm.SetNextPos(xR+totalX, yR+totalY)
}
