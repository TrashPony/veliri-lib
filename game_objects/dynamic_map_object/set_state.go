package dynamic_map_object

import (
	"github.com/TrashPony/veliri-lib/game_objects/target"
)

func (o *Object) SetPower(power int) {
	o.CurrentEnergy = power
}

func (o *Object) SetWeaponTarget(_ int, target *target.Target) {
	o.weaponTarget = target
}

func (o *Object) SetGunRotate(angle float64, slotNumber int) {
	o.GetWeaponSlot(slotNumber).SetGunRotate(angle)
}

func (o *Object) SetAnchorsEquip() {
	if o == nil || o.Equips[1] == nil || o.Equips[1].Equip == nil {
		return
	}

	o.Equips[1].SetAnchor()
}

func (o *Object) SetRun(run bool) {
	o.Run = run
}

func (o *Object) SetRadarWorkerExit(exit bool) {
	o.radarWorkerExit = exit
}

func (o *Object) SetRadarWorkerWork(work bool) {
	o.radarWorkerWork = work
}

func (o *Object) SetWork(work bool) {
	o.Work = work
}

func (o *Object) SetScale(scale int) {
	o.Scale = scale
}

func (o *Object) SetGrowTime(time int) {
	o.GrowTime = time
}

func (o *Object) SetHP(hp int) {
	if o.Immortal {
		return
	}

	o.HP = hp
}

func (o *Object) SetMapID(id int) {
	o.MapID = id
}

func (o *Object) SetComplete(complete float64) {
	o.Complete = complete
	if o.Complete > 100 {
		o.Complete = 100
	}

	if o.Complete < 0 {
		o.Complete = 0
	}
}
