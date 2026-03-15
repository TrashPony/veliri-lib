package game_math

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"math"
)

func GetReachAngleWithVelocity(
	xPlace, yPlace int,
	vxUnit, vyUnit float64, // скорость носителя (пикс/тик)
	xTarget, yTarget int,
	targetLvl, startLvlBullet float64,
	bulletSpeed float64, // скорость пули (пикс/сек)
	artillery bool,
	weaponType string,
	g float64,
) float64 {

	// Константа
	const serverTickMs = 32.0

	// 1. Переводим скорость пули из пикс/сек в пикс/тик
	bulletSpeedPerTick := bulletSpeed * (serverTickMs / 1000.0)

	// 2. Гравитация (если не задана)
	if g == 0 {
		g = GetGravityByWeaponType(weaponType)
	}
	// Переводим гравитацию в пикс/тик²
	gPerTick := g * (serverTickMs / 1000.0) * (serverTickMs / 1000.0)

	// 3. Относительные координаты цели
	dx := float64(xTarget - xPlace)
	dy := float64(yTarget - yPlace)
	dist2D := math.Hypot(dx, dy) // горизонтальная дистанция

	// 4. Разница высот
	deltaZ := targetLvl - startLvlBullet

	// 5. Начальное приближение угла (без учёта скорости носителя)
	// Используем стандартную формулу
	v := bulletSpeedPerTick
	x := dist2D
	y := deltaZ

	// Формула: root = v^4 - g*(g*x^2 + 2*y*v^2)
	v2 := v * v
	v4 := v2 * v2
	x2 := x * x

	root := v4 - gPerTick*(gPerTick*x2+2*y*v2)

	// Защита от отрицательного корня
	if root < 0 {
		root = 0
	}
	root = math.Sqrt(root)

	// Два возможных угла
	y1 := v2 + root
	y2 := v2 - root
	gx := gPerTick * x

	var angle float64
	if artillery {
		angle = math.Atan2(y1, gx)
	} else {
		angle = math.Atan2(y2, gx)
	}

	// 6. Коррекция угла с учётом скорости носителя
	// Раскладываем скорость носителя на компоненты относительно направления стрельбы
	cosAngle := math.Cos(angle)
	sinAngle := math.Sin(angle)

	// Проекция скорости носителя на направление стрельбы
	vUnitParallel := vxUnit*cosAngle + vyUnit*sinAngle

	// Корректируем эффективную скорость пули
	vEffective := bulletSpeedPerTick + vUnitParallel

	// 7. Пересчитываем угол с эффективной скоростью
	if math.Abs(vEffective) > 0.001 {
		v2Eff := vEffective * vEffective
		v4Eff := v2Eff * v2Eff

		rootEff := v4Eff - gPerTick*(gPerTick*x2+2*y*v2Eff)
		if rootEff < 0 {
			rootEff = 0
		}
		rootEff = math.Sqrt(rootEff)

		y1Eff := v2Eff + rootEff
		y2Eff := v2Eff - rootEff

		var angleEff float64
		if artillery {
			angleEff = math.Atan2(y1Eff, gx)
		} else {
			angleEff = math.Atan2(y2Eff, gx)
		}

		// Плавно смешиваем углы (можно просто вернуть скорректированный)
		angle = angleEff
	}

	return angle
}

func GetMaxDistBallisticWeapon(speed, targetLvl, startLvlBullet, radian float64, weaponType string, g float64) float64 {
	// https://en.m.wikipedia.org/wiki/Range_of_a_projectile

	y := startLvlBullet - targetLvl
	if g == 0 {
		g = GetGravityByWeaponType(weaponType)
	}
	g2 := 2 * g

	if radian < 0 {
		// формула из вики не работает с отрицательными углами
		// https://math.stackexchange.com/questions/2838938/correct-formula-to-find-the-range-of-a-projectile-given-angle-possibly-negativ/2839024#2839024?newreg=46afc5af54f44bbdb7ec018492791891
		Vx := speed * Cos(radian)
		Vy := speed * Sin(radian)
		g = g * -1

		root := math.Sqrt(math.Abs(fastPow(Vy) - (4 * g / 2 * y)))
		dist := Vx * ((Vy + root) / g)

		return dist * -1

	} else {
		a := fastPow(speed) / g2
		b := (g2 * y) / (fastPow(speed) * fastPow(Sin(radian)))
		dist := a * (1 + math.Sqrt(1+b)) * Sin(2*radian)
		return dist
	}
}

func GetSpeedByMaxRange(maxRange, targetLvl, startLvlBullet, radian float64, weaponType string, gravity float64) float64 {
	// TODO

	return 0
}

func GetMaxHeightBulletBallisticWeapon(speed, startRadian, startY float64, weaponType string, g float64) float64 {

	// https://en.wikipedia.org/wiki/Projectile_motion#Angle_required_to_hit_coordinate_.28x.2Cy.29
	// # Maximum height of projectile

	// startY => bullet.StartZ
	if g == 0 {
		g = GetGravityByWeaponType(weaponType)
	}

	v2 := fastPow(speed)
	maxHeight := ((v2 * fastPow(Sin(startRadian))) / (2 * g)) + startY
	return maxHeight
}

func GetZBulletByXPath(startZ, startRadian, speed float64, startX, startY, currentX, currentY int, ammoType string, gravity float64) float64 {

	/*
				1 angle
		(__)===D
				-1 angle
	*/

	// делаем из 3х мерной модели двумерную, за счет оси Х у нас будет дистанция от обьекта до цели
	// Y - высота пули
	// ху обьекта(0,0)

	// https://en.wikipedia.org/wiki/Projectile_motion#Angle_required_to_hit_coordinate_.28x.2Cy.29
	// #Displacement

	x := GetBetweenDist(startX, startY, currentX, currentY)
	return GetZBulletByX(startZ, startRadian, x, speed, ammoType, gravity)
}

func GetZBulletByX(startZ, startRadian, x, speed float64, weaponType string, gravity float64) float64 {

	if gravity == 0 {
		gravity = GetGravityByWeaponType(weaponType)
	}

	offsetZ := math.Tan(startRadian)*x - (gravity/(2*(fastPow(speed))*(Cos(startRadian)*Cos(startRadian))))*fastPow(x)
	return startZ + offsetZ
}

func GetGravityByWeaponType(weaponType string) float64 {
	g := _const.Gravity

	if weaponType == "missile" {
		g = 1
	}

	if weaponType == "meteorite" {
		g = _const.Gravity / 2
	}

	return g
}
