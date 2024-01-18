package dynamic_map_object

import (
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/burst_of_shots"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"github.com/TrashPony/veliri-lib/game_objects/target"
	"github.com/TrashPony/veliri-lib/game_objects/user_skills"
)

func (o *Object) GetX() int {
	return o.GetPhysicalModel().X
}

func (o *Object) GetY() int {
	return o.GetPhysicalModel().Y
}

func (o *Object) GetOwnerID() int {
	if o.OwnerBaseID != 0 {
		return 0
	}

	return o.OwnerID
}

func (o *Object) GetOwnerPlayerID() int {
	if o.OwnerBaseID != 0 {
		return 0
	}

	return o.OwnerPlayerID
}

func (o *Object) GetPower() int {
	return o.CurrentEnergy
}

func (o *Object) GetHP() int {
	return o.HP
}

func (o *Object) GetScale() int {
	return o.Scale
}

func (o *Object) GetComplete() float64 {
	return o.Complete
}

func (o *Object) GetWork() bool {
	return o.Work
}

func (o *Object) GetEquipFirePos(typeEquip, numberSlot int) []*game_math.Positions {

	equipSlot := o.Equips[1] // TODO
	if equipSlot == nil || equipSlot.Equip == nil {
		return nil
	}

	return game_math.GetWeaponFirePositions(
		o.GetX(), o.GetY(), o.GetScale(), o.GetPhysicalModel().GetRotate(), equipSlot.Rotate,
		equipSlot.Equip.XAttach, equipSlot.Equip.YAttach,
		equipSlot.Equip.FirePositions, 128, 128,
	)
}

// у некоторых обьектов например стуктур для строительства статичный размер
func (o *Object) GetStartScale() {

	if o.Texture == "explores_antenna" {
		o.SetScale(100)
		return
	}

	if o.Texture == "explores_observatory" {
		o.SetScale(100)
		return
	}

	if o.Texture == "unknown_civilization_jammer" {
		o.SetScale(300)
		return
	}

	if o.Texture == "replic_gauss_gun" {
		o.SetScale(140)
		return
	}

	if o.Type == "turret" {
		o.SetScale(50)
	}

	if o.Texture == "shield_generator" {
		o.SetScale(50)
	}

	if o.Texture == "extractor" {
		o.SetScale(50)
	}

	if o.Texture == "energy_generator" {
		o.SetScale(74)
	}

	if o.Texture == "jammer_generator" {
		o.SetScale(60)
	}

	if o.Texture == "missile_defense" {
		o.SetScale(34)
	}

	if o.Texture == "meteorite_defense" {
		o.SetScale(50)
	}

	if o.Texture == "radar" {
		o.SetScale(74)
	}

	if o.Texture == "storage" {
		o.SetScale(74)
	}

	if o.Texture == "beacon" {
		o.SetScale(74)
	}

	if o.Texture == "repair_station" {
		o.SetScale(50)
	}

	if o.Texture == "zone_repair_station" {
		o.SetScale(75)
	}

	if o.Texture == "mini_turret_1" || o.Texture == "mini_turret_2" {
		o.SetScale(15)
	}

	if o.Texture == "rope_trap_1" {
		o.SetScale(25)
	}

	if o.Texture == "gravity_generator_1" {
		o.SetScale(20)
	}

	if o.Texture == "wall_1" {
		o.SetScale(40)
	}

	if o.Texture == "mini_shield_generator_1" {
		o.SetScale(25)
	}

	if o.Texture == "corporation_base_1" || o.Texture == "corporation_base_1_build" ||
		o.Texture == "corporation_office_1" || o.Texture == "corporation_office_1_build" ||
		o.Texture == "corporation_processing_1" || o.Texture == "corporation_processing_1_build" ||
		o.Texture == "corporation_prefabricated_1" || o.Texture == "corporation_prefabricated_1_build" {
		o.SetScale(100)
	}
}

func (o *Object) GetWeaponTarget() *target.Target {
	return o.weaponTarget
}

func (o *Object) GetDistWeaponToTarget() int {

	weaponTarget := o.GetWeaponTarget()
	if weaponTarget == nil || o == nil {
		return 9999
	}

	return int(game_math.GetBetweenDist(weaponTarget.GetX(), weaponTarget.GetY(), o.GetPhysicalModel().X, o.GetPhysicalModel().Y))
}

func (o *Object) GetWeaponSlot(slotNumber int) *detail.BodyWeaponSlot {
	return o.Weapons[slotNumber]
}

func (o *Object) GetID() int {
	return o.ID
}

func (o *Object) GetType() string {
	return "object"
}

func (o *Object) CheckViewCoordinate(x, y, radius int) (bool, bool) {
	if o.GetRangeView()+radius >= int(game_math.GetBetweenDist(o.GetPhysicalModel().X, o.GetPhysicalModel().Y, x, y)) {
		return true, true
	}

	if o.GetRadarRange()+radius >= int(game_math.GetBetweenDist(o.GetPhysicalModel().X, o.GetPhysicalModel().Y, x, y)) {
		return false, true
	}

	return false, false
}

func (o *Object) GetRangeView() int {
	if o == nil {
		return 0
	}

	if o.OwnerBaseID > 0 {
		return o.RangeView * 2
	}

	return o.RangeView
}

func (o *Object) GetRadarRange() int {
	radarRange := 0

	if o == nil {
		return radarRange
	}

	if o.Equips[1] != nil && o.Equips[1].Equip != nil && o.Equips[1].Equip.Applicable == "radar" && o.CurrentEnergy > 0 {
		radarRange = o.Equips[1].Equip.Radius
	}

	o.RangeRadar = radarRange

	if o.OwnerBaseID > 0 {
		return o.RangeView * 5
	}

	return o.RangeRadar
}

func (o *Object) GetGunRotate(slotNumber int) float64 {
	return o.GetWeaponSlot(slotNumber).GetGunRotate()
}

func (o *Object) GetGunRotateSpeed(slotNumber int, skills *user_skills.UserSkills) int {
	return o.GetWeaponSlot(slotNumber).Weapon.RotateSpeed
}

func (o *Object) GetWeaponAccuracy(slotNumber int, skillsStore *user_skills.UserSkills) int {
	return o.GetWeaponSlot(slotNumber).Weapon.Accuracy
}

func (o *Object) GetCapSize() int {
	if o.Inventory {
		return o.InventoryCapacity
	} else {
		return 0
	}
}

func (o *Object) GetPowerMove() float64 {
	return 0
}

func (o *Object) GetEquipPosInMap(typeEquip, numberSlot int) (int, int) {

	if o == nil || o.Equips[1] == nil || o.Equips[1].Equip == nil {
		return 0, 0
	}

	return game_math.GetWeaponPosInMap(
		o.GetX(), o.GetY(), o.GetScale(),
		float64(o.Equips[1].XAttach),
		float64(o.Equips[1].YAttach),
		o.GetPhysicalModel().Rotate)
}

func (o *Object) GetRun() bool {
	return o.Run
}

func (o *Object) GetRadarWorkerExit() bool {
	return o.radarWorkerExit
}

func (o *Object) GetRadarWorkerWork() bool {
	return o.radarWorkerWork
}

func (o *Object) GetGrowTime() int {
	return o.GrowTime
}

func (o *Object) GetMapID() int {
	return o.MapID
}

func (o *Object) GetRotate() float64 {
	return o.GetPhysicalModel().GetRotate()
}

func (o *Object) GetBurstOfShots() *burst_of_shots.BurstOfShots {
	if o.burstOfShots == nil {
		o.burstOfShots = &burst_of_shots.BurstOfShots{}
	}

	return o.burstOfShots
}

func (o *Object) RangeWeaponSlots() map[int]*detail.BodyWeaponSlot {
	// мы никогда не пишут в карту слотов оружия поэтому этот метод безопасен (по крайне мере пока)
	return o.Weapons
}
