package game_math

import (
	"math"
)

func VectorSub(p1, p2 *Point) *Point {
	return &Point{
		X: p2.X - p1.X,
		Y: p2.Y - p1.Y,
	}
}

func IntersectVectorToCircle(a, b, centerCircle *Point, radius int) (intersect bool, x1, y1, x2, y2 float64) {
	// https://stackoverflow.com/questions/1073336/circle-line-segment-collision-detection-algorithm
	// вычисляем расстояние между A и B
	var LAB = math.Sqrt(fastPow(b.X-a.X) + fastPow(b.Y-a.Y))

	// вычислить вектор направления D от A до B
	var Dx = (b.X - a.X) / LAB
	var Dy = (b.Y - a.Y) / LAB

	// compute the value t of the closest point to the circle center (Cx, Cy)
	var t = Dx*(centerCircle.X-a.X) + Dy*(centerCircle.Y-a.Y)

	// This is the projection of C on the line from A to B.

	// вычислить координаты точки E на прямой и ближайшей к C
	xE, yE := (t*Dx)+a.X, (t*Dy)+a.Y

	// высчитывает растояние от E до центра круга
	var LEC = math.Sqrt(fastPow(xE-centerCircle.X) + fastPow(yE-centerCircle.Y))

	// проверяем что бы проекционная точка была ближе радиуса
	if int(LEC) < radius {
		// compute distance from t to circle intersection point
		var dt = math.Sqrt(fastPow(float64(radius)) - fastPow(LEC))
		// ищем первую точку пересечения
		x1, y1 := (t-dt)*Dx+a.X, (t-dt)*Dy+a.Y
		// и вторую
		x2, y2 := (t+dt)*Dx+a.X, (t+dt)*Dy+a.Y

		intersect := PointInVector(a, b, x1, y1) || PointInVector(a, b, x2, y2)
		return intersect, x1, y1, x2, y2
	}

	if int(LEC) == radius { // else test if the line is tangent to circle
		// прямая прилегает к окружности 1 точка пересечени
		return PointInVector(a, b, xE, yE), xE, yE, 0, 0
	}

	return false, 0, 0, 0, 0
}

func fastPow(a float64) float64 {
	return a * a
}

func IntersectVectorToCircleFast(a, b, centerCircle *Point, radius int) (intersect bool) {
	// Быстрая проверка на nil
	if a == nil || b == nil || centerCircle == nil {
		return false
	}

	// Конвертируем float64 из Point в int для быстрой версии
	ax, ay := int(a.X), int(a.Y)
	bx, by := int(b.X), int(b.Y)
	cx, cy := int(centerCircle.X), int(centerCircle.Y)

	return intersectVectorToCircleFast(ax, ay, bx, by, cx, cy, radius)
}

// intersectVectorToCircleFast проверяет, пересекает ли отрезок [A, B] окружность.
// Все координаты - целые числа (int).
func intersectVectorToCircleFast(ax, ay, bx, by, cx, cy, radius int) bool {
	// Переводим в int64 для безопасности умножения
	Ax, Ay := int64(ax), int64(ay)
	Bx, By := int64(bx), int64(by)
	Cx, Cy := int64(cx), int64(cy)
	r := int64(radius)

	// Вектор AB
	ABx := Bx - Ax
	ABy := By - Ay

	// Вектор AC
	ACx := Cx - Ax
	ACy := Cy - Ay

	// Квадрат длины отрезка AB (|AB|^2). Используем int64.
	AB2 := ABx*ABx + ABy*ABy

	// Если отрезок - это точка (A == B)
	if AB2 == 0 {
		dist2ToA := ACx*ACx + ACy*ACy
		return dist2ToA <= r*r
	}

	// Проекция вектора AC на AB. t - это параметр точки E (проекции C на прямую AB),
	// умноженный на |AB|^2. Это позволяет нам работать в целых числах.
	// t = (AC · AB)
	t := ACx*ABx + ACy*ABy

	var dist2ToLine int64 // Квадрат расстояния от C до отрезка AB

	if t < 0 {
		// Точка C проецируется за точку A, ближайшая точка отрезка - A
		dist2ToLine = ACx*ACx + ACy*ACy
	} else if t > AB2 {
		// Точка C проецируется за точку B, ближайшая точка отрезка - B
		BCx := Cx - Bx
		BCy := Cy - By
		dist2ToLine = BCx*BCx + BCy*BCy
	} else {
		// Проекция C попадает на отрезок AB. Ближайшая точка - E.
		// Квадрат расстояния от C до E (LEC^2) считается по формуле:
		// |CE|^2 = |AC|^2 - t^2/|AB|^2
		// Чтобы избежать деления, домножим на AB2:
		// CE2 * AB2 = AC2 * AB2 - t^2
		AC2 := ACx*ACx + ACy*ACy
		// Используем int128 или проверку на переполнение, если числа очень большие.
		// В Go нет встроенного int128, но для большинства игровых карт (координаты до 1e6) int64 хватит.
		// dist2ToLineAB2 = AC2 * AB2 - t*t
		dist2ToLineTimesAB2 := AC2*AB2 - t*t
		if dist2ToLineTimesAB2 < 0 {
			dist2ToLineTimesAB2 = 0 // от небольших погрешностей вычислений
		}
		// Сравниваем (LEC^2) с (r^2). Так как у нас LEC^2 = dist2ToLineTimesAB2 / AB2,
		// то условие LEC^2 <= r^2 эквивалентно dist2ToLineTimesAB2 <= r^2 * AB2
		return dist2ToLineTimesAB2 <= r*r*AB2
	}

	// Сравниваем dist2ToLine (квадрат расстояния до конца отрезка) с r^2
	return dist2ToLine <= r*r
}
