package game_math

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"math"
)

func GetReachAngle(xPlace, yPlace, xTarget, yTarget int, targetLvl, startLvlBullet, speed float64, artillery bool, weaponType string, g float64) float64 {

	// https://en.wikipedia.org/wiki/Projectile_motion#Angle_required_to_hit_coordinate_.28x.2Cy.29
	// # Angle required to hit coordinate (x,y)

	v := speed
	y := targetLvl - startLvlBullet
	if g == 0 {
		g = GetGravityByWeaponType(weaponType)
	}

	// делаем из 3х мерной модели двумерную, за счет оси Х у нас будет дистанция от обьекта до цели Y - высота пули
	// ху обьекта(0,0)
	x := GetBetweenDist(xPlace, yPlace, xTarget, yTarget)

	root := math.Pow(speed, 4) - g*(g*fastPow(x)+2*y*fastPow(v))
	root = math.Sqrt(root)

	y1 := fastPow(v) + root
	y2 := fastPow(v) - root
	gx := g * x

	if artillery {
		return math.Atan2(y1, gx)
	} else {
		return math.Atan2(y2, gx)
	}
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
