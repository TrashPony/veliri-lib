package drone

import (
	"math"
	"testing"

	"github.com/TrashPony/veliri-lib/game_objects/effect"
)

func directRangeView(d *Drone) int {
	return int(math.Ceil(d.GetEffects().GetAllBonus(float64(d.RangeView), "view")))
}

// Кэш дальности обзора не должен отставать от эффектов и базовой дальности.
func TestGetRangeViewFollowsEffectsAndBase(t *testing.T) {
	d := &Drone{RangeView: 800}

	check := func(step string) {
		t.Helper()
		if got, want := d.GetRangeView(), directRangeView(d); got != want {
			t.Fatalf("%s: GetRangeView()=%d, прямой расчёт %d", step, got, want)
		}
	}

	check("без эффектов")
	check("повторный вызов из кэша")

	d.GetEffects().AddEffect(&effect.Effect{UUID: "v1", Parameter: "view", Quantity: 20, Percentages: true})
	check("+20% обзора")
	if d.GetRangeView() != 960 {
		t.Fatalf("800 + 20%% = 960, получили %d", d.GetRangeView())
	}

	d.GetEffects().AddEffect(&effect.Effect{UUID: "v2", Parameter: "view", Quantity: 100})
	check("+100 абсолютных")

	d.RangeView = 500
	check("сменилась базовая дальность")

	d.GetEffects().RemoveEffect("v1")
	check("убрали один эффект")

	d.GetEffects().RemoveEffect("v2")
	check("убрали все эффекты")
	if d.GetRangeView() != 500 {
		t.Fatalf("без эффектов дальность = базовая 500, получили %d", d.GetRangeView())
	}
}

func BenchmarkGetRangeView(b *testing.B) {
	d := &Drone{RangeView: 800}
	d.GetEffects().AddEffect(&effect.Effect{UUID: "v1", Parameter: "view", Quantity: 20, Percentages: true})
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = d.GetRangeView()
	}
}
