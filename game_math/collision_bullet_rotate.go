package game_math

import "math"

func CollisionBulletRotate(collider1, collider2 collider, weight1, weight2, x2, y2 float64) {
	cX1, cY1 := collider1.GetNextPos()
	cX2, cY2 := collider2.GetNextPos()

	x1, y1 := collider1.GetRealPos()
	//x2, y2 := collider2.GetRealPos()

	m1, m2 := weight1, weight2

	startXV1, startYV1 := collider1.GetVelocity()
	startXV2, startYV2 := collider2.GetVelocity()

	m21 := m2 / m1
	x21 := cX2 - cX1
	y21 := cY2 - cY1
	vx21 := startXV2 - startXV1
	vy21 := startYV2 - startYV1

	vx_cm := (m1*startXV1 + m2*startXV2) / (m1 + m2)
	vy_cm := (m1*startYV1 + m2*startYV2) / (m1 + m2)

	a := y21 / x21
	dvx2 := -2 * (vx21 + a*vy21) / ((1 + a*a) * (1 + m21))
	vx2 := startXV2 + dvx2
	vy2 := startYV2 + a*dvx2
	vx1 := startXV1 - m21*dvx2
	vy1 := startYV1 - a*m21*dvx2

	R := 1.0

	vx1 = (vx1-vx_cm)*R + vx_cm
	vy1 = (vy1-vy_cm)*R + vy_cm
	vx2 = (vx2-vx_cm)*R + vx_cm
	vy2 = (vy2-vy_cm)*R + vy_cm

	collisionRotate(collider1, collider2, startXV1, startYV1, startXV2, startYV2, x1, x2, y1, y2, vx1, vy1)
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
