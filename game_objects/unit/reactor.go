package unit

import (
	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"github.com/TrashPony/veliri-lib/game_objects/effect"
	"strconv"
)

func (u *Unit) WorkReactorPower(removeCount int, reloadCallBack func(u *Unit, slot *detail.ThoriumSlot)) {
	u.mx.Lock()
	defer u.mx.Unlock()

	// если эффективность реактор изменилась то применяем изменения
	defer func() {
		eff := u.EfficiencyReactor()
		if u.ReactorEfficiency != eff {
			u.ReactorEfficiency = eff
			u.AppendFuelModifier()
		}
	}()

	if removeCount == 0 {
		for _, slot := range u.getBody().ThoriumSlots {
			if slot.Worked == 0 {
				u.getNextFuel(slot, reloadCallBack)
			}
		}
	}

	for i := 0; i < removeCount; i++ {
		for _, slot := range u.getBody().ThoriumSlots {

			if slot.Worked == 0 {
				u.getNextFuel(slot, reloadCallBack)
			}

			if slot.Worked > 0 {
				slot.Worked--
				if slot.Worked <= 0 {
					slot.Worked = 0
					slot.Reload = true
					u.getNextFuel(slot, reloadCallBack)
				}
			}
		}
	}
}

func (u *Unit) getNextFuel(slot *detail.ThoriumSlot, reloadCallBack func(u *Unit, slot *detail.ThoriumSlot)) {
	if slot.SendRequest {
		return
	}

	slot.SendRequest = true

	go func() {
		defer func() {
			slot.SendRequest = false
		}()

		reloadCallBack(u, slot)
	}()
}

func (u *Unit) UpdateReactorState(slots map[int]*detail.ThoriumSlot, rs int) {
	u.mx.Lock()
	defer u.mx.Unlock()

	update := false
	for key, slot := range u.getBody().ThoriumSlots {
		s := slots[key]
		if s != nil {

			update = update || slot.CurrentFuel.ID != s.CurrentFuel.ID // другое топливо может давать бонусы

			slot.Worked = s.Worked
			slot.CurrentFuel = s.CurrentFuel
			slot.NextFuel = s.NextFuel

			if slot.Number == rs {
				slot.Reload = false
			}
		}
	}

	if update {
		u.AppendFuelModifier()
	}
}

func (u *Unit) EfficiencyReactor() float64 {
	full := 0.0
	for _, slot := range u.getBody().ThoriumSlots {
		if slot.Worked > 0 || slot.Reload {
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
		for _, slot := range u.getBody().ThoriumSlots {
			u.RemoveEffect("fuelBonus_" + parameter + strconv.Itoa(slot.Number)) // TODO производительность?
		}
	}

	for _, slot := range u.getBody().ThoriumSlots {
		if slot.Worked > 0 {
			// с каждой ячейки добавляем баф, то есть если ячейки 3 то можно сделать баф х3 или 3 разных
			for _, ef := range slot.CurrentFuel.Bonuses {
				u.AddEffect(&effect.Effect{
					UUID:        "fuelBonus_" + ef.Parameter + strconv.Itoa(slot.Number), // TODO производительность?
					Parameter:   ef.Parameter,
					Quantity:    ef.Quantity,
					Percentages: ef.Percentages,
					Subtract:    ef.Subtract,
				})
			}
		}
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
