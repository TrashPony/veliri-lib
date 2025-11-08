package game_math

import (
	"github.com/TrashPony/veliri-lib/game_objects/obstacle_point"
	"math"
)

func weaponCollisionReaction(collider1, collider2 collider, weaponPoint1, weaponPoint2 *obstacle_point.ObstaclePoint, weight1, weight2 float64) (int, int, float64) {

	// Расчет урона для ОБОИХ участников
	damage1 := calculateMeleeWeaponDamage(collider2, collider1, weaponPoint2, weight2, weight1)
	damage2 := calculateMeleeWeaponDamage(collider1, collider2, weaponPoint1, weight1, weight2)

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
		// Для collider2 вектор направления противоположный
		powerLoss2 := calculatePowerLoss(collider2, collider1, -dirX, -dirY, totalForce, weight1)
		collider2.SetPowerMove(collider2.GetPowerMove() * powerLoss2)

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

func calculatePowerLoss(collider1, collider2 collider, dirX, dirY float64, impactForce float64, weight float64) float64 {
	// 1. Получаем текущее направление движения collider1
	currentVX, currentVY := collider1.GetVelocity()
	currentSpeed := math.Sqrt(currentVX*currentVX + currentVY*currentVY)

	if currentSpeed == 0 {
		return 0.8 // Максимальные потери если стоял на месте
	}

	// 2. Нормализуем текущее направление
	currentDirX := currentVX / currentSpeed
	currentDirY := currentVY / currentSpeed

	// 3. Вычисляем косинус угла между текущим движением и вектором столкновения
	dotProduct := currentDirX*dirX + currentDirY*dirY

	// Чем больше угол - тем больше потери энергии
	// dotProduct = 1.0 → движение вперед → минимальные потери (0.2)
	// dotProduct = -1.0 → лобовое столкновение → максимальные потери (0.8)
	// dotProduct = 0.0 → боковой удар → средние потери (0.5)

	// 4. Базовые потери от угла
	basePowerLoss := 0.5 - (dotProduct * 0.3)

	// 5. УЧИТЫВАЕМ МАССУ: если collider2 статичный (weight2 >= 20000)
	if weight >= 20000 {
		// Удар о статичный объект - дополнительные потери
		staticObjectPenalty := 0.3
		basePowerLoss += staticObjectPenalty
	} else {
		// Удар о подвижный объект - меньшие потери
		mobileObjectBonus := 0.1
		basePowerLoss -= mobileObjectBonus
	}

	// 6. Ограничиваем диапазон
	powerLoss := math.Max(0.1, math.Min(0.9, basePowerLoss))

	return powerLoss
}

func calculateMeleeWeaponDamage(attacker, target collider, weaponPoint *obstacle_point.ObstaclePoint, attackWeight, targetWeight float64) int {

	// Базовые параметры (для случая без оружия)
	k := 10
	sharpness := 1.0

	// Если есть оружие - используем его параметры
	if weaponPoint != nil {
		k = weaponPoint.K
	}

	// Базовый урон от скорости
	attackSpeed := attacker.GetCurrentSpeed()

	minSpeed := 2.0
	if attackSpeed < minSpeed {
		attackSpeed = minSpeed
	}

	baseDamage := attackSpeed * 10.0

	// Модификатор массы
	massFactor := (math.Max((attackWeight / 1000.0), 1.0)) / 3.0

	// Модификатор усиления урона из weaponPoint.K
	damageBoost := 1.0 + float64(k)/100.0 // K=100 → +100% урона

	// Итоговый урон
	totalDamage := baseDamage * sharpness * massFactor * damageBoost

	return int(totalDamage * 0.2)
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
