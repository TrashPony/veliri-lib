package game_math

import (
	"github.com/TrashPony/veliri-lib/game_objects/obstacle_point"
	"math"
)

func weaponCollisionReaction(collider1, collider2 collider, weaponPoint1, weaponPoint2 *obstacle_point.ObstaclePoint, weight1, weight2 float64) (int, int, float64) {

	// Определяем кто атакующий, а кто цель
	var attacker, target collider
	var attackPoint *obstacle_point.ObstaclePoint
	var attackWeight, targetWeight float64

	if weaponPoint1 != nil {
		// collider1 атакует collider2 своим оружием
		attacker, target = collider1, collider2
		attackPoint = weaponPoint1
		attackWeight, targetWeight = weight1, weight2
	} else {
		// collider2 атакует collider1 своим оружием
		attacker, target = collider2, collider1
		attackPoint = weaponPoint2
		attackWeight, targetWeight = weight2, weight1
	}

	// Берем реальные позиции
	targetX, targetY := target.GetRealPos()

	// Вектор от атакующего к цели через точку оружия
	weaponToTargetX := targetX - float64(attackPoint.X)
	weaponToTargetY := targetY - float64(attackPoint.Y)

	// Нормализуем
	length := math.Sqrt(weaponToTargetX*weaponToTargetX + weaponToTargetY*weaponToTargetY)
	if length > 0 {
		weaponToTargetX /= length
		weaponToTargetY /= length
	}

	// Сила удара основана на скорости атакующего
	attackSpeed := attacker.GetCurrentSpeed()
	impactForce := attackSpeed * 0.5 // Коэффициент можно настроить

	minForce := 1.0 // Минимальная сила отталкивания
	if impactForce < minForce {
		impactForce = minForce
	}

	// Применяем силу к цели - толкаем ОТ точки удара
	targetVX, targetVY := target.GetVelocity()
	target.SetVelocity(
		targetVX+weaponToTargetX*impactForce*(attackWeight/(attackWeight+targetWeight)),
		targetVY+weaponToTargetY*impactForce*(attackWeight/(attackWeight+targetWeight)),
	)

	// Отдача для атакующего (меньше)
	attackerVX, attackerVY := attacker.GetVelocity()
	recoilForce := impactForce * 0.3
	attacker.SetVelocity(
		attackerVX-weaponToTargetX*recoilForce*(targetWeight/(attackWeight+targetWeight)),
		attackerVY-weaponToTargetY*recoilForce*(targetWeight/(attackWeight+targetWeight)),
	)

	// Замедление атакующего при "застревании"
	attacker.SetPowerMove(attacker.GetPowerMove() * 0.8)

	// Расчет урона
	damage := calculateMeleeWeaponDamage(attackSpeed, attackWeight, attackPoint)

	// Определяем кто получил урон
	var damage1, damage2 int
	if weaponPoint1 != nil {
		damage1, damage2 = 0, damage // collider2 получает урон
	} else {
		damage1, damage2 = damage, 0 // collider1 получает урон
	}

	// Звук удара
	collisionPower := math.Min(impactForce/10.0, 1.0)

	return damage1, damage2, collisionPower
}

func calculateMeleeWeaponDamage(attackSpeed, attackWeight float64, attackPoint *obstacle_point.ObstaclePoint) int {
	// Базовый урон от скорости
	baseDamage := attackSpeed * 2.0

	// Модификатор от "остроты" оружия (радиус точки)
	sharpness := 1.0
	if attackPoint.Radius < 5 {
		sharpness = 1.5 // Тонкое оружие больнее
	}

	// Модификатор массы
	massFactor := math.Min(attackWeight/1000.0, 2.0)

	return int(baseDamage * sharpness * massFactor)
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
	// Проверяем коллизию круга оружия с телом
	// TODO "object"
	bodyX, bodyY := body.GetRealPos()
	dist := math.Sqrt(
		math.Pow(float64(weaponPoint.X)-bodyX, 2) +
			math.Pow(float64(weaponPoint.Y)-bodyY, 2),
	)
	return dist < float64(weaponPoint.Radius+body.GetRadius())
}
