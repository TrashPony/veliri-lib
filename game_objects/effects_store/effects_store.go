package effects_store

import (
	"github.com/TrashPony/veliri-lib/game_objects/effect"
	"sync"
)

type EffectsStore struct {
	Effects []*effect.Effect `json:"-"`
	mx      sync.RWMutex
}

func (e *EffectsStore) AddEffect(newEffect *effect.Effect) bool {

	if e.GetEffectByUUID(newEffect.UUID) != nil {
		return false
	}

	if newEffect.Parameter == "reload" || newEffect.Parameter == "reload_ammo" || newEffect.Parameter == "accuracy" {
		newEffect.Subtract = !newEffect.Subtract
	}

	e.mx.Lock()
	defer e.mx.Unlock()

	if e.Effects == nil {
		e.Effects = make([]*effect.Effect, 0)
	}

	e.Effects = append(e.Effects, newEffect)
	return true
}

func (e *EffectsStore) RemoveEffect(uuid string) (bool, *effect.Effect) {
	e.mx.Lock()
	defer e.mx.Unlock()

	if e.Effects == nil {
		e.Effects = make([]*effect.Effect, 0)
	}

	index := -1
	var removeEffect *effect.Effect
	for i, ef := range e.Effects {
		if ef.UUID == uuid {
			index = i
			removeEffect = ef
			break
		}
	}

	if index >= 0 {
		e.Effects = append(e.Effects[:index], e.Effects[index+1:]...)
	}

	return index >= 0, removeEffect
}

func (e *EffectsStore) RemoveBySlot(slotType, slotNumber int) (bool, []*effect.Effect) {
	var uuids []string

	e.mx.RLock()
	for _, ef := range e.Effects {
		if ef.SlotType == slotType && ef.SlotNumber == slotNumber {
			uuids = append(uuids, ef.UUID)
		}
	}
	e.mx.RUnlock()

	remove := false
	removeEffects := make([]*effect.Effect, 0)
	for _, u := range uuids {

		rem, remEf := e.RemoveEffect(u)
		if remEf != nil {
			removeEffects = append(removeEffects, remEf)
		}

		remove = rem || remove
	}

	return remove, removeEffects
}

func (e *EffectsStore) GetEffectByUUID(uuid string) *effect.Effect {
	e.mx.Lock()
	defer e.mx.Unlock()

	for _, ef := range e.Effects {
		if ef.UUID == uuid {
			return ef
		}
	}

	return nil
}

func (e *EffectsStore) GetCountByName(parameter string, percent bool) int {
	e.mx.Lock()
	defer e.mx.Unlock()

	allValue := 0
	for _, ef := range e.Effects {

		if ef.Percentages != percent {
			continue
		}

		if ef.Parameter == parameter {
			if ef.Subtract {
				allValue -= ef.Quantity
			} else {
				allValue += ef.Quantity
			}
		}
	}

	return allValue
}

func (e *EffectsStore) GetCountByWeaponTypeAndSize(parameter string, percent bool, typeWeapon string, sizeWeapon int) int {
	e.mx.Lock()
	defer e.mx.Unlock()

	allValue := 0
	for _, ef := range e.Effects {

		if ef.Percentages != percent {
			continue
		}

		if ef.StandardSize > 0 && ef.StandardSize != sizeWeapon {
			continue
		}

		if ef.WeaponType != "" && ef.WeaponType != typeWeapon {
			continue
		}

		if ef.Parameter == parameter {
			if ef.Subtract {
				allValue -= ef.Quantity
			} else {
				allValue += ef.Quantity
			}
		}
	}

	return allValue
}

func (e *EffectsStore) GetCountByBodyTypeAndSize(parameter string, percent bool, typeBody string, sizeBody int) int {
	e.mx.Lock()
	defer e.mx.Unlock()

	allValue := 0
	for _, ef := range e.Effects {

		if ef.Percentages != percent {
			continue
		}

		if ef.StandardSize > 0 && ef.StandardSize != sizeBody {
			continue
		}

		if ef.BodyType != "" && ef.BodyType != typeBody {
			continue
		}

		if ef.Parameter == parameter {
			if ef.Subtract {
				allValue -= ef.Quantity
			} else {
				allValue += ef.Quantity
			}
		}
	}

	return allValue
}

func (e *EffectsStore) GetAllBonus(startValue float64, parameterName string) float64 {
	sumEffects := effect.Effect{
		Quantity:    e.GetCountByName(parameterName, false),
		Parameter:   parameterName,
		Percentages: false,
	}

	absoluteValue := sumEffects.ToAccept(startValue, parameterName)

	sumEffectsPercent := effect.Effect{
		Quantity:    e.GetCountByName(parameterName, true),
		Parameter:   parameterName,
		Percentages: true,
	}

	if sumEffectsPercent.Quantity < -90 {
		sumEffectsPercent.Quantity = -90 // что бы не уйти в минус)
	}

	return sumEffectsPercent.ToAccept(absoluteValue, parameterName)
}

func (e *EffectsStore) GetAllWeaponBonus(startValue float64, parameterName, typeWeapon string, sizeWeapon int) float64 {
	sumEffects := effect.Effect{
		Quantity:    e.GetCountByWeaponTypeAndSize(parameterName, false, typeWeapon, sizeWeapon),
		Parameter:   parameterName,
		Percentages: false,
	}

	absoluteValue := sumEffects.ToAccept(startValue, parameterName)

	sumEffectsPercent := effect.Effect{
		Quantity:    e.GetCountByWeaponTypeAndSize(parameterName, true, typeWeapon, sizeWeapon),
		Parameter:   parameterName,
		Percentages: true,
	}

	if sumEffectsPercent.Quantity < -90 {
		sumEffectsPercent.Quantity = -90 // что бы не уйти в минус)
	}

	return sumEffectsPercent.ToAccept(absoluteValue, parameterName)
}

func (e *EffectsStore) GetAllBodyBonus(startValue float64, parameterName, typeBody string, sizeBody int) float64 {
	sumEffects := effect.Effect{
		Quantity:    e.GetCountByBodyTypeAndSize(parameterName, false, typeBody, sizeBody),
		Parameter:   parameterName,
		Percentages: false,
	}

	absoluteValue := sumEffects.ToAccept(startValue, parameterName)

	sumEffectsPercent := effect.Effect{
		Quantity:    e.GetCountByBodyTypeAndSize(parameterName, true, typeBody, sizeBody),
		Parameter:   parameterName,
		Percentages: true,
	}

	if sumEffectsPercent.Quantity < -90 {
		sumEffectsPercent.Quantity = -90 // что бы не уйти в минус)
	}

	return sumEffectsPercent.ToAccept(absoluteValue, parameterName)
}

func (e *EffectsStore) GetAllEffects() []*effect.Effect {
	effects := make([]*effect.Effect, 0, len(e.Effects))
	for _, e := range e.Effects {
		effects = append(effects, e)
	}

	return effects
}

func (e *EffectsStore) GetMinQuantityByParameter(parameter string) (bool, int) {
	e.mx.Lock()
	defer e.mx.Unlock()

	find := false
	minQuantity := 0

	for _, ef := range e.Effects {
		if ef.Parameter == parameter {
			if ef.Quantity < minQuantity || !find {
				find = true
				minQuantity = ef.Quantity
			}
		}
	}

	return find, minQuantity
}
