package game_math

import (
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
}

// Вынесение сложных расчетов в отдельные функции
func calculateCollisionResponse(cX1, cY1, cX2, cY2, startXV1, startYV1, startXV2, startYV2, m1, m2 float64) (vx1, vy1, vx2, vy2 float64) {
	m21 := m2 / m1
	x21 := cX2 - cX1
	y21 := cY2 - cY1
	vx21 := startXV2 - startXV1
	vy21 := startYV2 - startYV1

	// Центр масс системы
	vx_cm := (m1*startXV1 + m2*startXV2) / (m1 + m2)
	vy_cm := (m1*startYV1 + m2*startYV2) / (m1 + m2)

	// Проверка направления столкновения
	if vx21*x21+vy21*y21 >= 0 {
		return startXV1, startYV1, startXV2, startYV2
	}

	// Обработка особого случая (деление на ноль)
	if math.Abs(x21) < 1e-12 {
		x21 = 1e-12
		if cX2 < cX1 {
			x21 = -x21
		}
	}

	a := y21 / x21
	dvx2 := -2 * (vx21 + a*vy21) / ((1 + a*a) * (1 + m21))

	vx2 = startXV2 + dvx2
	vy2 = startYV2 + a*dvx2
	vx1 = startXV1 - m21*dvx2
	vy1 = startYV1 - a*m21*dvx2

	// Коэффициент упругости
	R := 1.0
	vx1 = (vx1-vx_cm)*R + vx_cm
	vy1 = (vy1-vy_cm)*R + vy_cm
	vx2 = (vx2-vx_cm)*R + vx_cm
	vy2 = (vy2-vy_cm)*R + vy_cm

	return vx1, vy1, vx2, vy2
}

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

func CollisionReactionBallBall(collider1, collider2 collider, weight1, weight2, pf1, pf2, x2, y2 float64) (int, int, float64) {
	// Получаем позиции и скорости
	cX1, cY1 := collider1.GetNextPos()
	cX2, cY2 := collider2.GetNextPos()
	x1, y1 := collider1.GetRealPos()

	startXV1, startYV1 := collider1.GetVelocity()
	startXV2, startYV2 := collider2.GetVelocity()

	// Проверка на специальные случаи
	if direct := (startXV2-startXV1)*(cX2-cX1) + (startYV2-startYV1)*(cY2-cY1); direct >= 0 {
		if weight2 >= 20000 || (collider1.GetType() == "object" && collider2.GetType() == "object") {
			d1, d2, _, _ := fixAdhesion(collider1, collider2, startXV1, startXV2, startYV1, startYV2, weight1, weight2, x1, y1, x2, y2)
			return d1, d2, calculateCollisionSound(collider1, collider2, startXV1, startYV1, startXV2, startYV2, weight1, weight2)
		}
		return 0, 0, calculateCollisionSound(collider1, collider2, startXV1, startYV1, startXV2, startYV2, weight1, weight2)
	}

	// Расчет новых скоростей
	vx1, vy1, vx2, vy2 := calculateCollisionResponse(cX1, cY1, cX2, cY2,
		startXV1, startYV1, startXV2, startYV2, weight1, weight2)

	// Применение изменений
	applyCollisionResults(collider1, collider2, x1, y1, x2, y2, vx1, vy1, vx2, vy2)

	// Расчет вращения
	collisionRotate(collider1, collider2, startXV1, startYV1, startXV2, startYV2, x1, x2, y1, y2, vx1, vy1)

	d1, d2 := getDamage(startXV1, startYV1, vx1, vy1, startXV2, startYV2, vx2, vy2,
		weight1, weight2, collider1.GetRotate(), collider2.GetRotate(), collider1, collider2)
	return d1, d2, calculateCollisionSound(collider1, collider2, startXV1, startYV1, startXV2, startYV2, weight1, weight2)
}

func applyCollisionResults(c1, c2 collider, x1, y1, x2, y2, vx1, vy1, vx2, vy2 float64) {
	if c2.GetType() != "map_item" {
		c1.SetVelocity(vx1, vy1)
		c1.SetNextPos(x1+vx1, y1+vy1)
		c1.SetPowerMove(c1.GetPowerMove() / 1.5)
	}

	if c1.GetType() != "map_item" {
		c2.SetNextPos(x2+vx2, y2+vy2)
		c2.SetVelocity(vx2, vy2)
	}
}

func collisionRotate(collider1, collider2 collider, startXV1, startYV1, startXV2, startYV2, x1, x2, y1, y2, vx1, vy1 float64) {
	// Пропускаем специальные случаи
	if collider1.GetType() == "unit" && collider2.GetType() == "object" {
		return
	}

	// Рассчитываем векторы
	impactVector := Vector{X: x2 - x1, Y: y2 - y1}
	velocityVector := Vector{X: startXV2, Y: startYV2}
	normalVector := velocityVector.Norm().Sub(impactVector.Norm())

	// Рассчитываем угол воздействия
	angle := calculateImpactAngle(x1, y1, x2, y2, normalVector)
	if math.IsNaN(angle) {
		return
	}

	// Рассчитываем силу вращения
	velocityChange := &Vector{X: startXV1 - vx1, Y: startYV1 - vy1}
	rotationForce := calculateRotationForce(velocityChange, normalVector)

	if rotationForce == 0 {
		return
	}

	// Применяем вращение в зависимости от угла
	if angle > 0 && angle < 180 {
		collider1.SetAngularVelocity(collider1.GetAngularVelocity() + rotationForce)
	} else {
		collider1.SetAngularVelocity(collider1.GetAngularVelocity() - rotationForce)
	}
}

func calculateImpactAngle(x1, y1, x2, y2 float64, normal *Vector) float64 {
	angle1 := GetBetweenAngle(x1+normal.X, normal.Y+y1, x1, y1)
	angle2 := GetBetweenAngle(x2, y2, x1, y1)
	angle := angle1 - angle2
	PrepareAngle(&angle)
	return angle
}

func calculateRotationForce(velocityChange, normal *Vector) float64 {
	lenNormal := 2 - normal.Len()
	if lenNormal < 0.02 {
		return 0
	}
	return velocityChange.Len()/100 + normal.Len()/100
}

// иногда модели слипают и не могут разлипнуть
func fixAdhesion(collider1, collider2 collider, XVelocity1, XVelocity2, YVelocity1, YVelocity2, m1, m2, cX1, cY1, cX2, cY2 float64) (int, int, float64, float64) {

	speed1 := math.Sqrt((XVelocity1 * XVelocity1) + (YVelocity1 * YVelocity1))
	speed2 := math.Sqrt((XVelocity2 * XVelocity2) + (YVelocity2 * YVelocity2))

	minSpeed := 3.0
	if speed1 < minSpeed {
		speed1 = minSpeed
	}

	if speed2 < minSpeed {
		speed2 = minSpeed
	}

	if m1 > m2 {
		speed1 = (speed1) * m2 / m1
	} else {
		speed2 = (speed2) * m1 / m2
	}

	angle1 := DegToRadian(GetBetweenAngle(cX1, cY1, cX2, cY2))
	angle2 := DegToRadian(GetBetweenAngle(cX2, cY2, cX1, cY1))

	vx1, vy1 := speed1*Cos(angle1), speed1*Sin(angle1)
	vx2, vy2 := speed2*Cos(angle2), speed2*Sin(angle2)

	if collider2.GetType() != "map_item" {
		collider1.SetVelocity(vx1, vy1)
		collider1.SetPowerMove(collider1.GetPowerMove() / 1.5)
	}

	if collider1.GetType() != "map_item" {
		collider2.SetVelocity(vx2, vy2)
		collider2.SetPowerMove(collider1.GetPowerMove() / 1.5)
	}

	return 0, 0, vx1, vy1
	//return getDamage(XVelocity1, YVelocity1, vx1, vy1, XVelocity2, YVelocity2, vx2, vy2, m1, m2)
}

func getDamage(
	startXV1, startYV1, XVelocity1, YVelocity1 float64,
	startXV2, startYV2, XVelocity2, YVelocity2 float64,
	m1, m2, rotate1Deg, rotate2Deg float64,
	collider1, collider2 collider,
) (int, int) {
	const (
		damageMultiplier = 0.0002 // Общий множитель урона
		minDamageEnergy  = 5.0    // Минимальная энергия для урона
	)

	relVx := startXV1 - startXV2
	relVy := startYV1 - startYV2

	nx := startXV2 - startXV1
	ny := startYV2 - startYV1
	normLength := math.Sqrt(nx*nx + ny*ny)
	if normLength == 0 {
		return 0, 0
	}
	nx /= normLength
	ny /= normLength

	velocityProjection := relVx*nx + relVy*ny
	if velocityProjection > 0 {
		return 0, 0
	}

	// Новый код: определяем тыловой удар
	isRearAttack := velocityProjection < -10.0 // Порог скорости для "сильного" тылового удара
	energyBoost := 1.0
	if isRearAttack {
		energyBoost = 1.0 + math.Min(math.Abs(velocityProjection)/20.0, 3.0)
	}

	reducedMass := (m1 * m2) / (m1 + m2)
	energy := 0.5 * reducedMass * velocityProjection * velocityProjection * energyBoost

	if energy < minDamageEnergy {
		return 0, 0
	}

	baseDamage := energy * damageMultiplier

	damage1 := int(baseDamage * (m2 / (m1 + m2)))
	damage2 := int(baseDamage * (m1 / (m1 + m2)))

	return damage1, damage2
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
