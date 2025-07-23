package dynamic_map_object

import "github.com/TrashPony/veliri-lib/game_objects/damage_manager"

func (o *Object) FillLastDamage(d []damage_manager.Damage) {
	o.damageManager.FillLastDamage(d)
}

func (o *Object) GetArrayLastDamage() []damage_manager.Damage {
	return o.damageManager.GetArrayLastDamage()
}

func (o *Object) SetLastDamage(playerID, damage int) *damage_manager.Damage {
	if playerID == o.OwnerPlayerID {
		return nil
	}

	return o.damageManager.SetLastDamage(playerID, damage)
}

func (o *Object) GetLastDamage(sec int64) int {
	return o.damageManager.GetLastDamage(sec)
}

func (o *Object) GetLastDamageByPlayerID(playerID, sec int) int {
	return o.damageManager.GetLastDamageByPlayerID(playerID, sec)
}

func (o *Object) GetAllDamage(t int64) map[int]int {
	return o.damageManager.GetAllDamage(t)
}
