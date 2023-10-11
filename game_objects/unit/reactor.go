package unit

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"github.com/TrashPony/veliri-lib/game_objects/effect"
)

func (u *Unit) WorkReactorPower(removeCount int) {
	u.mx.Lock()
	defer u.mx.Unlock()

	// если эффективность реактор изменилась то применяем изменения
	startEfficiency := u.EfficiencyReactor()
	defer func() {
		if startEfficiency != u.EfficiencyReactor() {
			u.AppendFuelModifier()
		}
	}()

	for i := 0; i < removeCount; i++ {
		for _, slot := range u.getBody().ThoriumSlots {
			if slot.GetCount() > 0 {
				slot.WorkedOut++
				if slot.WorkedOut >= _const.PowerInWork {
					slot.WorkedOut = 0
					slot.SetCount(slot.GetCount() - 1)
				}
			}
		}
	}
}

func (u *Unit) UpdateReactorState(slots map[int]*detail.ThoriumSlot) {
	u.mx.Lock()
	defer u.mx.Unlock()

	for key, slot := range u.getBody().ThoriumSlots {
		s := slots[key]
		if s != nil {
			slot.Count = s.Count
			slot.WorkedOut = s.WorkedOut
		}
	}
}

func (u *Unit) EfficiencyReactor() float64 {
	full := 0.0
	for _, slot := range u.getBody().ThoriumSlots {
		if slot.GetCount() > 0 {
			full++
		}
	}

	if full == 0 {
		return 0
	}

	return full / float64(len(u.getBody().ThoriumSlots))
}

var reactorChangeParameters = []string{"radar", "view", "turn_speed", "reverse_factor", "power_factor", "reverse_speed", "speed", "charging_speed"}

func (u *Unit) AppendFuelModifier() {
	for _, parameter := range reactorChangeParameters {
		u.RemoveEffect("fuel_" + parameter)
	}

	percentFuel := u.EfficiencyReactor() * 100
	if int(percentFuel) == 100 {
		return
	}

	doublePercentFuel := float64(50) * (percentFuel / 100.0)
	for _, parameter := range reactorChangeParameters {
		u.AddEffect(&effect.Effect{
			UUID:        "fuel_" + parameter,
			Parameter:   parameter,
			Quantity:    50 - int(doublePercentFuel),
			Percentages: true,
			Subtract:    true,
		})
	}
}
