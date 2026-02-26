package game_math

import (
	"github.com/TrashPony/veliri-lib/game_objects/obstacle_point"
	"math"
)

func weaponCollisionReaction(collider1, collider2 collider, weaponPoint1, weaponPoint2 *obstacle_point.ObstaclePoint, weight1, weight2 float64) (int, int, float64) {

	// Вычисляем относительную скорость
	relativeDamageSpeed := calculateAttackerEfficiency(collider1, collider2)

	c2s := collider2.GetCurrentSpeed()
	c1s := collider1.GetCurrentSpeed()

	if c1s < 0.1 {
		c1s = 0.1
	}
	if c2s < 0.1 {
		c2s = 0.1
	}

	rawEf2 := relativeDamageSpeed / c2s
	rawEf1 := relativeDamageSpeed / c1s

	sum := rawEf2 + rawEf1
	efficiency1 := rawEf1 / sum
	efficiency2 := rawEf2 / sum

	// Расчет урона
	damage1 := calculateMeleeWeaponDamageWithEfficiency(collider2, collider1, weaponPoint2, weight2, weight1, efficiency1, relativeDamageSpeed)
	damage2 := calculateMeleeWeaponDamageWithEfficiency(collider1, collider2, weaponPoint1, weight1, weight2, efficiency2, relativeDamageSpeed)

	// Для физики используем центры масс (универсально для любого типа столкновения)
	targetX, targetY := collider2.GetRealPos()
	attackerX, attackerY := collider1.GetRealPos()

	// Вектор от collider1 к collider2
	dirX := targetX - attackerX
	dirY := targetY - attackerY

	// Нормализуем
	length := math.Sqrt(dirX*dirX + dirY*dirY)
	if length > 0 {
		dirX /= length
		dirY /= length
	}

	// Сила удара основана на ОТНОСИТЕЛЬНОЙ скорости
	speed1 := collider1.GetCurrentSpeed()
	speed2 := collider2.GetCurrentSpeed()
	relativeSpeed := math.Abs(speed1 - speed2)

	penetrationDepth := calculateObjectPenetrationDepth(collider1, collider2)

	ejectForce := penetrationDepth * 0.5 // Коэффициент можно настроить
	maxEjectForce := 3.0
	if ejectForce > maxEjectForce {
		ejectForce = maxEjectForce
	}

	impactForce := relativeSpeed * 0.5
	minForce := 1.0
	if impactForce < minForce {
		impactForce = minForce
	}

	totalForce := impactForce + ejectForce
	maxTotalForce := 5.0
	if totalForce > maxTotalForce {
		totalForce = maxTotalForce
	}

	// Применяем силу к ОБОИМ участникам
	vx1, vy1 := collider1.GetVelocity()
	vx2, vy2 := collider2.GetVelocity()

	// В твоем коде:
	if collider2.GetType() != "map_item" {
		powerLoss1 := calculatePowerLoss(collider1, collider2, dirX, dirY, totalForce, weight2)
		collider1.SetPowerMove(collider1.GetPowerMove() * powerLoss1)

		collider1.SetVelocity(
			vx1-dirX*totalForce*(weight2/(weight1+weight2)),
			vy1-dirY*totalForce*(weight2/(weight1+weight2)),
		)

		newVX, newVY := collider1.GetVelocity()
		cx, cy := collider1.GetRealPos()
		collider1.SetNextPos(cx+newVX, cy+newVY)
	}

	if collider1.GetType() != "map_item" {

		collider2.SetVelocity(
			vx2+dirX*totalForce*(weight1/(weight1+weight2)),
			vy2+dirY*totalForce*(weight1/(weight1+weight2)),
		)

		newVX, newVY := collider2.GetVelocity()
		cx, cy := collider2.GetRealPos()
		collider2.SetNextPos(cx+newVX, cy+newVY)
	}

	// Звук удара
	collisionPower := math.Min(totalForce/10.0, 1.0)

	applyCollisionRotation(collider1, collider2, dirX, dirY, totalForce, weight1, weight2)

	return damage1, damage2, collisionPower
}

func applyCollisionRotation(collider1, collider2 collider, dirX, dirY, totalForce, weight1, weight2 float64) {
	if collider1.GetType() == "unit" && collider2.GetType() == "object" {
		return
	}

	x1, y1 := collider1.GetRealPos()
	x2, y2 := collider2.GetRealPos()

	impactVector := Vector{X: x2 - x1, Y: y2 - y1}
	impactLength := math.Sqrt(impactVector.X*impactVector.X + impactVector.Y*impactVector.Y)
	if impactLength > 0 {
		impactVector.X /= impactLength
		impactVector.Y /= impactLength
	}

	normalVector := Vector{X: -impactVector.Y, Y: impactVector.X}
	vx1, vy1 := collider1.GetVelocity()
	tangentialDot := (vx1*normalVector.X + vy1*normalVector.Y) * 0.2

	rotationMultiplier := 0.033
	rotationForce := 0.03 + (tangentialDot * totalForce * rotationMultiplier)

	massFactor := math.Max(weight2/weight1, 2.0) / 3
	rotationForce *= massFactor

	maxRotation := 0.5
	if math.Abs(rotationForce) > maxRotation {
		rotationForce = math.Copysign(maxRotation, rotationForce)
	}

	if collider2.GetType() != "map_item" {
		collider1.SetAngularVelocity(collider1.GetAngularVelocity() + rotationForce)
	}

	if collider1.GetType() != "map_item" && collider1.GetType() != "object" {
		collider2.SetAngularVelocity(collider2.GetAngularVelocity() - rotationForce*0.5)
	}
}

func calculateObjectPenetrationDepth(collider1, collider2 collider) float64 {
	// Находим максимальную глубину проникновения между всеми кругами
	maxPenetration := 0.0

	// Определяем кто object, а кто простой коллайдер
	var objectCollider, simpleCollider collider
	if collider1.GetType() == "object" {
		objectCollider, simpleCollider = collider1, collider2
	} else {
		objectCollider, simpleCollider = collider2, collider1
	}

	// Для составного объекта проверяем все круги
	for _, objPoint := range objectCollider.GetGeoData() {
		simpleX, simpleY := simpleCollider.GetRealPos()

		dist := math.Sqrt(
			math.Pow(float64(objPoint.X)-simpleX, 2) +
				math.Pow(float64(objPoint.Y)-simpleY, 2),
		)

		sumRadii := float64(objPoint.Radius + simpleCollider.GetRadius())
		if dist < sumRadii {
			penetration := sumRadii - dist
			if penetration > maxPenetration {
				maxPenetration = penetration
			}
		}
	}

	return maxPenetration
}

func calculatePowerLoss(collider, other collider, dirX, dirY, impactForce, otherWeight float64) float64 {
	currentVX, currentVY := collider.GetVelocity()
	currentSpeed := math.Sqrt(currentVX*currentVX + currentVY*currentVY)

	if currentSpeed == 0 {
		return 0.8
	}

	// Вектор движения текущего коллайдера
	currentDirX := currentVX / currentSpeed
	currentDirY := currentVY / currentSpeed

	// Косинус угла между движением и направлением удара
	dotProduct := currentDirX*dirX + currentDirY*dirY

	// Для подвижных целей проверяем ОТНОСИТЕЛЬНОЕ движение
	if otherWeight < 20000 {
		otherVX, otherVY := other.GetVelocity()

		// Проекции скоростей на вектор удара
		mySpeedProj := currentVX*dirX + currentVY*dirY
		otherSpeedProj := otherVX*dirX + otherVY*dirY

		// УНИВЕРСАЛЬНОЕ условие для догоняющих ударов:
		// Я движусь быстрее цели ВДОЛЬ вектора удара И в том же направлении
		if mySpeedProj > otherSpeedProj && dotProduct > 0 {
			// Догоняющий удар - минимальные потери
			return 0.9 // Всего 10% потерь
		}

		// Если мы движемся в противоположных направлениях - это встречный удар
		if mySpeedProj > 0 && otherSpeedProj < 0 {
			// Встречный удар - максимальные потери
			return 0.2
		}
	}

	// Базовые потери от угла
	basePowerLoss := 0.5 - (dotProduct * 0.3)

	// Статичные объекты дают дополнительные потери
	if otherWeight >= 20000 {
		basePowerLoss += 0.3
	}

	return math.Max(0.1, math.Min(0.9, basePowerLoss))
}

func calculateMeleeWeaponDamageWithEfficiency(attacker, target collider, weaponPoint *obstacle_point.ObstaclePoint, attackWeight, targetWeight, efficiency, relativeDamageSpeed float64) int {
	if efficiency <= 0 || math.IsNaN(efficiency) {
		efficiency = 0.1
	}

	if efficiency > 1 {
		efficiency = 1.0
	}

	if targetWeight > maxWeight {
		targetWeight = maxWeight
	}

	if attackWeight > maxWeight {
		attackWeight = maxWeight
	}

	// Базовые параметры
	k := 1
	sharpness := 1.0

	if weaponPoint != nil {
		k = weaponPoint.K
	}

	// 1. Скорость атакующего
	attackSpeed := relativeDamageSpeed * efficiency
	if attackSpeed <= 1 {
		attackSpeed = 1
	}

	// 3. Базовый урон от излишка скорости
	baseDamage := attackSpeed * attackSpeed

	baseMinDamage := math.Log1p(attackWeight/1000) * 2
	hardnessFactor := attackWeight / targetWeight
	speedFactor := relativeDamageSpeed * 0.2

	minDamage := baseMinDamage * hardnessFactor * speedFactor
	if baseDamage < minDamage {
		baseDamage = minDamage
	}

	// Модификатор массы
	massRatio := attackWeight / targetWeight
	massFactor := math.Pow(math.Max(0.01, massRatio), 0.5)

	// Модификатор усиления урона
	damageBoost := 1.0 + float64(k)/100.0

	// Итоговый урон
	totalDamage := baseDamage * sharpness * massFactor * damageBoost

	return int(totalDamage)
}

func calculateAttackerEfficiency(attacker, target collider) float64 {
	// Позиции и скорости
	attackerX, attackerY := attacker.GetRealPos()
	targetX, targetY := target.GetRealPos()
	attackVX, attackVY := attacker.GetVelocity()
	targetVX, targetVY := target.GetVelocity()

	// Вектор от атакующего к цели
	collisionDirX := targetX - attackerX
	collisionDirY := targetY - attackerY

	// Нормализуем направление удара
	collisionLength := math.Sqrt(collisionDirX*collisionDirX + collisionDirY*collisionDirY)
	if collisionLength == 0 {
		// Если объекты в одной точке — используем направление по умолчанию (например, вправо)
		collisionDirX, collisionDirY = 1.0, 0.0
	} else {
		collisionDirX /= collisionLength
		collisionDirY /= collisionLength
	}

	// Проекции скоростей на линию атаки
	attackSpeedProj := attackVX*collisionDirX + attackVY*collisionDirY
	targetSpeedProj := targetVX*collisionDirX + targetVY*collisionDirY

	// Относительная скорость: насколько быстро атакующий "врезается" в цель
	relativeSpeed := attackSpeedProj - targetSpeedProj

	// Не даём отрицательной скорости (удаление → меньше урона, но не "отрицательный урон")
	if relativeSpeed < 0 {
		relativeSpeed = 0
	}

	return relativeSpeed
}

func findWeaponCollision(meleeData1, meleeData2 []*obstacle_point.ObstaclePoint, collider1, collider2 collider) (*obstacle_point.ObstaclePoint, *obstacle_point.ObstaclePoint) {

	// Оружие 1 vs Корпус 2
	for _, weaponPoint := range meleeData1 {
		if checkPointToBodyCollision(weaponPoint, collider2) {
			return weaponPoint, nil
		}
	}

	// Оружие 2 vs Корпус 1
	for _, weaponPoint := range meleeData2 {
		if checkPointToBodyCollision(weaponPoint, collider1) {
			return nil, weaponPoint
		}
	}

	return nil, nil
}

func checkPointToBodyCollision(weaponPoint *obstacle_point.ObstaclePoint, body collider) bool {
	// Особый случай: если body - это "object" с составной геометрией
	if body.GetType() == "object" {
		// Проверяем коллизию с каждым кругом объекта
		for _, bodyPoint := range body.GetGeoData() {
			dist := math.Sqrt(
				math.Pow(float64(weaponPoint.X)-float64(bodyPoint.X), 2) +
					math.Pow(float64(weaponPoint.Y)-float64(bodyPoint.Y), 2),
			)
			if dist < float64(weaponPoint.Radius+bodyPoint.Radius) {
				return true
			}
		}
		return false
	}

	// Стандартная проверка для других типов (unit, box и т.д.)
	bodyX, bodyY := body.GetRealPos()
	dist := math.Sqrt(
		math.Pow(float64(weaponPoint.X)-bodyX, 2) +
			math.Pow(float64(weaponPoint.Y)-bodyY, 2),
	)

	return dist < float64(weaponPoint.Radius+body.GetRadius())
}
