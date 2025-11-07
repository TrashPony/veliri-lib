package game_math

import (
	"github.com/TrashPony/veliri-lib/game_objects/obstacle_point"
	"math"
)

type collider interface {
	GetX() int
	GetY() int
	GetVelocity() (float64, float64)
	GetVelocityRotate() float64
	GetCurrentSpeed() float64
	GetDirection() bool
	SetPowerMove(float64)
	SetVelocity(float64, float64)
	AddVelocity(float64, float64)
	GetPowerMove() float64
	SetNextPos(x, y float64)
	GetNextPos() (float64, float64)
	GetRealPos() (float64, float64)
	GetRotate() float64
	GetAngularVelocity() float64
	SetAngularVelocity(float64)
	GetType() string
	GetRadius() int
	GetGeoData() []*obstacle_point.ObstaclePoint
}

const maxWeight = 20000

func calculateCollisionSound(c1, c2 collider, startXV1, startYV1, startXV2, startYV2, m1, m2 float64) float64 {
	// 1. Рассчитываем относительную скорость столкновения
	relSpeedX := startXV1 - startXV2
	relSpeedY := startYV1 - startYV2
	impactSpeed := math.Sqrt(relSpeedX*relSpeedX + relSpeedY*relSpeedY)

	// 2. Минимальная скорость для воспроизведения звука
	const minImpactSpeed = 0.3
	if impactSpeed < minImpactSpeed {
		return 0 // Слишком слабое столкновение
	}

	// 3. Рассчитываем "силу" столкновения
	massFactor := math.Min((m1+m2)/1000, 1.0)
	speedFactor := math.Min(impactSpeed/10.0, 1.0)
	collisionPower := massFactor * speedFactor

	return collisionPower
}

func CollisionReactionBallBall(collider1, collider2 collider, meleeData1, meleeData2 []*obstacle_point.ObstaclePoint, weight1, weight2, pf1, pf2, x2, y2 float64) (int, int, float64) {

	// 1. Сначала проверяем столкновение оружия
	weaponCollisionPoint1, weaponCollisionPoint2 := findWeaponCollision(meleeData1, meleeData2, collider1, collider2)

	return weaponCollisionReaction(collider1, collider2, weaponCollisionPoint1, weaponCollisionPoint2, weight1, weight2)
}

func GetCollisionPoint(c1, c2 collider) (float64, float64) {
	x1, y1 := c1.GetNextPos()
	x2, y2 := c2.GetNextPos()
	r1 := float64(c1.GetRadius())

	// Вектор от c1 к c2
	dx := x2 - x1
	dy := y2 - y1

	// Нормализуем вектор (длина = 1)
	length := math.Sqrt(dx*dx + dy*dy)
	if length == 0 {
		return x1, y1 // На случай совпадения позиций
	}

	// Точка на границе первого коллайдера
	collisionX := x1 + (dx/length)*r1
	collisionY := y1 + (dy/length)*r1

	return collisionX, collisionY
}
