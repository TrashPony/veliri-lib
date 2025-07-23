package damage_manager

import (
	"sync"
	"time"
)

type DamageManager struct {
	damage []Damage
	mx     sync.RWMutex
}

type Damage struct {
	PlayerID   int   `json:"player_id"`
	TimeDamage int64 `json:"time_damage"`
	Damage     int   `json:"damage"`
}

func (dm *DamageManager) FillLastDamage(d []Damage) {
	dm.mx.Lock()
	defer dm.mx.Unlock()
	dm.damage = d
}

func (dm *DamageManager) GetArrayLastDamage() []Damage {
	dm.mx.Lock()
	defer dm.mx.Unlock()
	d := make([]Damage, len(dm.damage))

	for i, v := range dm.damage {
		d[i] = v
	}

	return d
}

func (dm *DamageManager) SetLastDamage(playerID, damage int) *Damage {
	dm.mx.Lock()
	defer dm.mx.Unlock()

	newDamage := Damage{
		PlayerID:   playerID,
		TimeDamage: time.Now().Unix(),
		Damage:     damage,
	}
	dm.damage = append(dm.damage, newDamage)
	return &newDamage
}

func (dm *DamageManager) GetLastDamage(sec int64) int {
	dm.mx.Lock()
	defer dm.mx.Unlock()

	if len(dm.damage) == 0 {
		return 0
	}

	last := dm.damage[len(dm.damage)-1]
	// если урон наносился больше чем 30 сек назад то он не учитывается
	if time.Now().Unix()-last.TimeDamage > sec {
		dm.damage = dm.damage[:0]
		return 0
	}

	return last.PlayerID
}

func (dm *DamageManager) GetLastDamageByPlayerID(playerID, sec int) int {
	dm.mx.Lock()
	defer dm.mx.Unlock()

	if len(dm.damage) == 0 {
		return 0
	}

	allDamage := 0
	for _, d := range dm.damage {
		if d.PlayerID == playerID {
			if time.Now().Unix()-d.TimeDamage < int64(sec) {
				allDamage += d.Damage
			}
		}
	}

	return allDamage
}

func (dm *DamageManager) GetAllDamage(t int64) map[int]int {
	dm.mx.Lock()
	defer dm.mx.Unlock()

	r := make(map[int]int)

	if len(dm.damage) == 0 {
		return r
	}

	for _, d := range dm.damage {
		if time.Now().Unix()-d.TimeDamage < t {
			r[d.PlayerID] += d.Damage
		}
	}

	return r
}
