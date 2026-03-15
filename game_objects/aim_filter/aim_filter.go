package aim_filter

import (
	"math"
	"time"
)

type AimFilter struct {
	lastX, lastY       int
	lastTime           int64
	historyX, historyY []int
	alpha              float64 // коэффициент сглаживания (0-1)
}

func NewAimFilter() *AimFilter {
	return &AimFilter{
		historyX: make([]int, 0, 5),
		historyY: make([]int, 0, 5),
		alpha:    0.3, // 0.3 - 30% новой точки, 70% истории
	}
}

func (f *AimFilter) Filter(rawX, rawY int) (int, int) {
	now := time.Now().UnixMilli()

	// Если это первая точка
	if len(f.historyX) == 0 {
		f.lastX, f.lastY = rawX, rawY
		f.lastTime = now
		f.historyX = append(f.historyX, rawX)
		f.historyY = append(f.historyY, rawY)
		return rawX, rawY
	}

	// Считаем скорость (пикселей в секунду)
	dt := float64(now-f.lastTime) / 1000.0
	if dt < 0.001 {
		dt = 0.001
	}
	dx := float64(rawX - f.lastX)
	dy := float64(rawY - f.lastY)
	speed := math.Hypot(dx, dy) / dt

	// Динамическая альфа: если скорость аномальная - больше сглаживания
	alpha := f.alpha
	if speed > 2000 { // подозрительно быстро
		alpha = f.alpha / 2 // почти не двигаемся
	} else if speed < 100 { // медленно
		alpha = f.alpha * 2 // больше доверия
	}

	// Сглаживание
	smoothX := int(float64(f.lastX)*(1-alpha) + float64(rawX)*alpha)
	smoothY := int(float64(f.lastY)*(1-alpha) + float64(rawY)*alpha)

	// Обновляем историю
	f.lastX, f.lastY = smoothX, smoothY
	f.lastTime = now
	f.historyX = append(f.historyX, smoothX)
	f.historyY = append(f.historyY, smoothY)

	// Держим историю небольшой
	if len(f.historyX) > 5 {
		f.historyX = f.historyX[1:]
		f.historyY = f.historyY[1:]
	}

	return smoothX, smoothY
}

func (f *AimFilter) Reset(x, y int) {
	f.lastX, f.lastY = x, y
	f.lastTime = time.Now().UnixMilli()
	f.historyX = []int{x}
	f.historyY = []int{y}
}
