package unit

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_objects/effect"
)

func (unit *Unit) WorkReactorPower(removeCount int) {

	// если эффективность реактор изменилась то применяем изменения
	startEfficiency := unit.EfficiencyReactor()
	defer func() {
		if startEfficiency != unit.EfficiencyReactor() {
			unit.AppendFuelModifier()
		}
	}()

	for i := 0; i < removeCount; i++ {
		for _, slot := range unit.getBody().ThoriumSlots {
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

func (unit *Unit) EfficiencyReactor() float64 {
	full := 0.0
	for _, slot := range unit.getBody().ThoriumSlots {
		if slot.GetCount() > 0 {
			full++
		}
	}

	if full == 0 {
		return 0.2
	}

	return full / float64(len(unit.getBody().ThoriumSlots))
}

var reactorChangeParameters = []string{"radar", "view", "turn_speed", "reverse_factor", "power_factor", "reverse_speed", "speed", "charging_speed"}

func (unit *Unit) AppendFuelModifier() {
	for _, parameter := range reactorChangeParameters {
		unit.RemoveEffect("fuel_" + parameter)
	}

	percentFuel := int(unit.EfficiencyReactor() * 100)
	if percentFuel == 100 {
		return
	}

	for _, parameter := range reactorChangeParameters {
		unit.AddEffect(&effect.Effect{
			UUID:        "fuel_" + parameter,
			Parameter:   parameter,
			Quantity:    100 - percentFuel,
			Percentages: true,
			Subtract:    true,
		})
	}
}
