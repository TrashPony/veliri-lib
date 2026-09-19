package unit

import (
	"testing"
	"time"

	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"github.com/TrashPony/veliri-lib/game_objects/effect"
)

func newTestUnit() *Unit {
	u := &Unit{ID: 1, Fraction: "Replics"}
	u.SetBody(&detail.Body{RangeView: 900, RangeRadar: 1300, Fraction: "Replics", Scale: 2, StandardSize: 2, MaxHP: 1000})
	u.HP = 1000
	u.UpdateViewState()
	return u
}

func TestVisionLinkHasPlayerWithoutLinks(t *testing.T) {
	u := newTestUnit()

	if u.VisionLinkHasPlayer(7) {
		t.Fatal("без ссылок обзора ответ обязан быть false")
	}
	if u.viewCache.visionLinks.Load() != nil {
		t.Fatal("вопрос без ссылок не должен создавать менеджер (это и есть быстрый путь)")
	}

	u.VisionLinkAddPlayer(7, time.Hour)
	if !u.VisionLinkHasPlayer(7) || u.VisionLinkHasPlayer(8) {
		t.Fatal("ссылка на 7 должна быть, на 8 - нет")
	}

	u.VisionLinkAddPlayer(9, -time.Second) // уже протухла
	if u.VisionLinkHasPlayer(9) {
		t.Fatal("протухшая ссылка не считается")
	}
}

// Quantity эффекта invisibility уменьшается на месте каждый тик (это таймер включения) - кэш по тику обязан видеть
// это между тиками и не менять ответ внутри тика.
func TestInvisibilityAtFollowsInPlaceCountdown(t *testing.T) {
	u := newTestUnit()

	tick := int64(32)
	if u.InvisibilityAt(tick) {
		t.Fatal("без эффекта юнит видим")
	}

	eff := &effect.Effect{UUID: "slot3", Parameter: "invisibility", Quantity: 96}
	u.AddEffect(eff)

	// идёт обратный отсчёт, как в mechanics/equip/invisibility.go: пока Quantity > 0 юнит ещё виден
	for step := 0; step < 8; step++ {
		tick += 32
		want := u.Invisibility()
		if got := u.InvisibilityAt(tick); got != want {
			t.Fatalf("тик %d (Quantity=%d): InvisibilityAt=%v, Invisibility=%v", tick, eff.Quantity, got, want)
		}

		// внутри тика значение не пересчитывается, даже если Quantity поменяли
		before := u.InvisibilityAt(tick)
		eff.Quantity -= 32
		if u.InvisibilityAt(tick) != before {
			t.Fatal("внутри одного тика ответ должен браться из кэша")
		}
	}

	tick += 32
	if !u.InvisibilityAt(tick) {
		t.Fatalf("после отсчёта (Quantity=%d) юнит невидим", eff.Quantity)
	}

	u.RemoveEffect("slot3")
	tick += 32
	if u.InvisibilityAt(tick) {
		t.Fatal("эффект снят - юнит снова виден")
	}
}

func TestGetUpdateHashTracksData(t *testing.T) {
	u := newTestUnit()

	tick := int64(64)
	h1 := u.GetUpdateHash(tick)
	if h1 == 0 {
		t.Fatal("хэш не может быть 0")
	}
	if u.GetUpdateHash(tick) != h1 {
		t.Fatal("в пределах тика хэш обязан быть тем же")
	}

	// данные не менялись -> следующий тик даёт тот же хэш (это и есть 'обновления нет')
	tick += 32
	if u.GetUpdateHash(tick) != h1 {
		t.Fatal("данные те же - хэш обязан совпасть с прошлым тиком")
	}

	u.HP = 500
	tick += 32
	h2 := u.GetUpdateHash(tick)
	if h2 == h1 {
		t.Fatal("HP изменился - хэш обязан измениться")
	}

	// GetUpdateData после хэша в том же тике возвращает те же данные, а хэш соответствует им
	tick += 32
	data := u.GetUpdateData(tick)
	hData := u.GetUpdateHash(tick)
	if len(data) == 0 || hData == 0 {
		t.Fatal("данные/хэш пусты")
	}
}

func BenchmarkInvisibility(b *testing.B) {
	u := newTestUnit()
	b.Run("Invisibility_uncached", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = u.Invisibility()
		}
	})
	b.Run("InvisibilityAt_cached", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = u.InvisibilityAt(64)
		}
	})
}

// прежняя реализация VisionLinkHasPlayer (эксклюзивный Mutex юнита + ленивое создание менеджера) для сравнения
func (u *Unit) visionLinkHasPlayerOld(targetID int) bool {
	u.mx.Lock()
	defer u.mx.Unlock()

	m := u.viewCache.visionLinks.Load()
	if m == nil {
		m = NewVisionLinkManager()
		u.viewCache.visionLinks.Store(m)
	}

	return m.Has(targetID)
}

func BenchmarkVisionLinkHasPlayer(b *testing.B) {
	b.Run("old_lock_per_call", func(b *testing.B) {
		u := newTestUnit()
		for i := 0; i < b.N; i++ {
			_ = u.visionLinkHasPlayerOld(7)
		}
	})
	b.Run("new_no_links", func(b *testing.B) {
		u := newTestUnit()
		for i := 0; i < b.N; i++ {
			_ = u.VisionLinkHasPlayer(7)
		}
	})
	b.Run("old_parallel_contention", func(b *testing.B) {
		u := newTestUnit()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = u.visionLinkHasPlayerOld(7)
			}
		})
	})
	b.Run("new_parallel", func(b *testing.B) {
		u := newTestUnit()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = u.VisionLinkHasPlayer(7)
			}
		})
	})
}

func BenchmarkGetUpdateHash(b *testing.B) {
	u := newTestUnit()
	b.Run("hash_cached_per_tick", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = u.GetUpdateHash(96)
		}
	})
	b.Run("bytes_data_cached_per_tick", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = u.GetUpdateData(96)
		}
	})
}
