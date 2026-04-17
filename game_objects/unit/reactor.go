package unit

import (
	"github.com/TrashPony/veliri-lib/game_objects/effect"
	"math"
	"strconv"
)

type Fuel struct {
	CurrentFuel int  `json:"current_fuel"`
	Reload      bool `json:"-"`
	SendRequest bool `json:"-"`
}

func (u *Unit) GetFuel() *Fuel {
	u.mx.Lock()
	defer u.mx.Unlock()
	return u.getFuel()
}

func (u *Unit) getFuel() *Fuel {
	if u.Fuel == nil {
		u.Fuel = &Fuel{}
	}

	return u.Fuel
}

func (u *Unit) GetCapFuel() int {
	u.mx.Lock()
	defer u.mx.Unlock()
	return u.getCapFuel()
}

func (u *Unit) getCapFuel() int {
	cap := u.GetMaxPower() * 10 // базовое вместилище

	for _, slot := range u.getBody().ThoriumSlots {
		if slot.CurrentFuel.ID > 0 {
			cap += slot.MaxCap
		}
	}

	return cap
}

func (u *Unit) WorkReactorPower(removeCount int, reloadCallBack func(u *Unit, slot *Fuel)) {
	u.mx.Lock()
	defer u.mx.Unlock()

	// если эффективность реактор изменилась то применяем изменения
	defer func() {
		eff, lowPower := u.EfficiencyReactor()
		if u.ReactorEfficiency != eff || u.LowPower != lowPower {
			u.AppendFuelModifier()
		}
	}()

	if removeCount == 0 {
		if u.getFuel().CurrentFuel == 0 {
			u.getNextFuel(reloadCallBack)
		}
	}

	for i := 0; i < removeCount; i++ {
		if u.getFuel().CurrentFuel == 0 {
			u.getNextFuel(reloadCallBack)
		}

		if u.getFuel().CurrentFuel > 0 {
			u.getFuel().CurrentFuel--
			if u.getFuel().CurrentFuel <= 0 {
				u.getFuel().CurrentFuel = 0
				u.getFuel().Reload = true
				u.getNextFuel(reloadCallBack)
			}
		}
	}
}

func (u *Unit) ChargeReactorPower(chage int) bool {
	u.mx.Lock()
	defer u.mx.Unlock()

	if u.getFuel().CurrentFuel > u.getCapFuel()/2 {
		return false
	}

	u.getFuel().CurrentFuel += chage
	u.getFuel().Reload = false

	if u.getFuel().CurrentFuel > (u.getCapFuel()/2)+(u.getCapFuel()/10) {
		u.getFuel().CurrentFuel = (u.getCapFuel() / 2) + (u.getCapFuel() / 10)
	}

	return true
}

func (u *Unit) getNextFuel(reloadCallBack func(u *Unit, slot *Fuel)) {
	if u.Fuel.SendRequest {
		return
	}

	u.Fuel.SendRequest = true

	go func() {
		defer func() {
			u.Fuel.SendRequest = false
		}()

		reloadCallBack(u, u.getFuel())
	}()
}

func (u *Unit) UpdateReactorState(fuelAmount int) {
	u.mx.Lock()
	defer u.mx.Unlock()

	u.getFuel().CurrentFuel = fuelAmount
	u.getFuel().Reload = false
	u.AppendFuelModifier()
}

func (u *Unit) EfficiencyReactor() (float64, bool) {
	if u.getFuel().CurrentFuel == 0 && !u.getFuel().Reload {
		// считаем что енергия кончается когда ее осталось меньше 33%, начинает мигать красным и все вот это
		return 0, u.GetPower() < (u.GetMaxPower() / 3)
	}

	fuelPercent := math.Ceil((float64(u.getFuel().CurrentFuel) / float64(u.getCapFuel())) * 100)
	return fuelPercent, false
}

// параметры для дебафа когда у реактора кончается топливо
var reactorChangeParameters = []string{"charging_speed"}

// параметры для дебафа когда кончилось топливо и не осталось енергии
var lowPowerParameters = []string{"radar", "view", "turn_speed", "reverse_factor", "power_factor", "reverse_speed", "speed"}

func (u *Unit) AppendFuelModifier() {
	for _, parameter := range reactorChangeParameters {
		u.RemoveEffect("fuel_" + parameter)
		for _, slot := range u.getBody().ThoriumSlots {
			u.RemoveEffect("fuelBonus_" + parameter + strconv.Itoa(slot.Number))
		}
	}

	for _, parameter := range lowPowerParameters {
		u.RemoveEffect("fuel_" + parameter)
		for _, slot := range u.getBody().ThoriumSlots {
			u.RemoveEffect("fuelBonus_" + parameter + strconv.Itoa(slot.Number))
		}
	}

	percentFuel, lowPower := u.EfficiencyReactor() // TODO общий пул енергии а не по ячейкам
	u.ReactorEfficiency = percentFuel
	u.LowPower = lowPower

	percentFuel = percentFuel * 100
	if int(percentFuel) > 0 {
		for _, slot := range u.getBody().ThoriumSlots {
			for _, ef := range slot.CurrentFuel.Bonuses {
				u.AddEffect(&effect.Effect{
					UUID:        "fuelBonus_" + ef.Parameter + strconv.Itoa(slot.Number),
					Parameter:   ef.Parameter,
					Quantity:    ef.Quantity,
					Percentages: ef.Percentages,
					Subtract:    ef.Subtract,
				})
			}
		}

		return
	}

	doublePercentFuel := float64(90) * (percentFuel / 100.0)
	for _, parameter := range reactorChangeParameters {
		u.AddEffect(&effect.Effect{
			UUID:        "fuel_" + parameter,
			Parameter:   parameter,
			Quantity:    90 - int(doublePercentFuel),
			Percentages: true,
			Subtract:    true,
		})
	}

	if u.LowPower {
		doublePercentFuel = float64(50) * (percentFuel / 100.0)
		for _, parameter := range lowPowerParameters {
			u.AddEffect(&effect.Effect{
				UUID:        "fuel_" + parameter,
				Parameter:   parameter,
				Quantity:    50 - int(doublePercentFuel),
				Percentages: true,
				Subtract:    true,
			})
		}
	}
}
