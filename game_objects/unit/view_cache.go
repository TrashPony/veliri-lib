package unit

import (
	"sync/atomic"

	"github.com/TrashPony/veliri-lib/game_objects/visible_objects"
)

// unitViewCache - состояние юнита, которое читает система обзора (game_loop_view) для каждого наблюдателя каждый тик.
// Поля атомарные: пишет их цикл карты, а читать могут и другие горутины.
type unitViewCache struct {
	// ссылки обзора: nil, пока не добавлена ни одна (см. VisionLinkHasPlayer)
	visionLinks atomic.Pointer[VisionLinkManager]

	// результат Invisibility() для тика invisibilityTime
	invisibilityTime  atomic.Int64
	invisibilityValue atomic.Bool
}

// InvisibilityAt - Invisibility() с кэшем на тик mapTime. Invisibility() ходит в набор эффектов под блокировкой, а
// зовёт её вью для каждого юнита для каждого наблюдателя. Кэшировать по версии набора эффектов нельзя: Quantity
// эффекта "invisibility" уменьшается на месте каждый тик (mechanics/equip/invisibility.go), поэтому - по тику. Внутри
// одного тика вью выполняется после всех фаз, меняющих эффекты, значение не меняется.
func (u *Unit) InvisibilityAt(mapTime int64) bool {
	if mapTime != 0 && u.viewCache.invisibilityTime.Load() == mapTime {
		return u.viewCache.invisibilityValue.Load()
	}

	value := u.Invisibility()
	u.viewCache.invisibilityValue.Store(value)
	u.viewCache.invisibilityTime.Store(mapTime)

	return value
}

// GetUpdateHash - хэш GetUpdateData(mapTime), посчитанный один раз за тик и лежащий рядом с Time (одна кэш-линия):
// radar.CheckObjects сравнивает хэши вместо двух срезов байтов.
func (u *Unit) GetUpdateHash(mapTime int64) uint64 {
	if u.CacheUpdateData.Time == mapTime && u.CacheUpdateData.Hash != 0 {
		return u.CacheUpdateData.Hash
	}

	data := u.GetUpdateData(mapTime) // при пересчёте сбрасывает Hash
	hash := visible_objects.HashUpdateData(data)
	u.CacheUpdateData.Hash = hash

	return hash
}
