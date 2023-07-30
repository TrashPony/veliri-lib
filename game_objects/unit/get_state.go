package unit

import (
	"encoding/json"
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"github.com/TrashPony/veliri-lib/game_objects/effect"
	"github.com/TrashPony/veliri-lib/game_objects/obstacle_point"
	"github.com/TrashPony/veliri-lib/game_objects/target"
	"strconv"
	"sync/atomic"
)

type StateMS struct {
	UseEnergy    int `json:"use_energy"`
	MaxEnergy    int `json:"max_energy"`
	CapacitySize int `json:"capacity_size"`

	// живучесть
	MaxHP                 int `json:"max_hp"`
	ProtectionToKinetics  int `json:"protection_to_kinetics"`
	ProtectionToThermo    int `json:"protection_to_thermo"`
	ProtectionToExplosion int `json:"protection_to_explosion"`

	// движение
	Speed         float64 `json:"speed"`          // -- макс скорость вперед
	ReverseSpeed  float64 `json:"reverse_speed"`  //-- макс скорость назад
	PowerFactor   float64 `json:"power_factor"`   // -- сила ускорения вперед
	ReverseFactor float64 `json:"reverse_factor"` //-- сила ускорения назад
	TurnSpeed     float64 `json:"turn_speed"`     //-- скорость поворота в радианах
	CurrentSpeed  float64 `json:"current_speed"`

	// навигация
	RangeView  int `json:"range_view"`
	RangeRadar int `json:"range_radar"`

	// генератор
	MaxPower           int `json:"max_power"`
	RecoveryPower      int `json:"recovery_power"`
	RecoveryPowerCycle int `json:"recovery_power_cycle"`

	// оружие
	WeaponType        string  `json:"wt"`
	WeaponSize        int     `json:"ws"`
	WeaponInstall     bool    `json:"install"`
	WeaponMaxDamage   int     `json:"weapon_max_damage"`
	WeaponMinDamage   int     `json:"weapon_min_damage"`
	CountFireBullet   int     `json:"count_fire_bullet"`
	WeaponMinRange    int     `json:"weapon_min_range"`
	WeaponMaxRange    int     `json:"weapon_max_range"`
	WeaponAccuracy    int     `json:"weapon_accuracy"`
	GunRotateSpeed    int     `json:"gun_rotate_speed"`
	BulletSpeed       float64 `json:"bullet_speed"`
	ReloadTime        int     `json:"reload_time"`
	ReloadAmmoTime    int     `json:"reload_ammo_time"`
	EfficiencyReactor float64 `json:"efficiency_reactor"`
	KineticsDamage    int     `json:"kinetics_damage"`
	ExplosionDamage   int     `json:"explosion_damage"`
	ThermoDamage      int     `json:"thermo_damage"`
}

func (unit *Unit) GetState() *StateMS {

	var weapon *detail.Weapon
	if unit.GetWeaponSlot(1) != nil {
		weapon = unit.GetWeaponSlot(1).Weapon
	}

	state := StateMS{
		UseEnergy:             unit.getBody().GetUseEnergy(),
		MaxEnergy:             unit.GetMaxEnergy(),
		MaxHP:                 unit.GetMaxHP(),
		ProtectionToKinetics:  unit.GetProtection("kinetics"),
		ProtectionToThermo:    unit.GetProtection("thermo"),
		ProtectionToExplosion: unit.GetProtection("explosion"),
		Speed:                 unit.GetMoveMaxPower(),
		ReverseSpeed:          unit.GetMaxReverse(),
		PowerFactor:           unit.GetPowerFactor(),
		ReverseFactor:         unit.GetReverseFactor(),
		TurnSpeed:             unit.GetTurnSpeed(),
		RangeView:             unit.GetRangeView(),
		RangeRadar:            unit.GetRadarRange(),
		WeaponInstall: unit.GetWeaponSlot(1) != nil &&
			weapon != nil &&
			unit.GetWeaponSlot(1).GetAmmo() != nil &&
			unit.GetWeaponSlot(1).GetAmmoQuantity() > 0,
		EfficiencyReactor:  unit.EfficiencyReactor(),
		CurrentSpeed:       unit.GetPhysicalModel().GetCurrentSpeed(),
		CapacitySize:       unit.GetCapSize(),
		MaxPower:           unit.GetMaxPower(),
		RecoveryPower:      unit.GetRecoveryPower(),
		RecoveryPowerCycle: unit.body.RecoveryPowerCycle,
	}

	if state.WeaponInstall {
		state.WeaponType = weapon.Type
		state.WeaponSize = weapon.StandardSize
		state.WeaponMaxDamage = unit.GetGunner().GetMaxDamage(unit.GetWeaponSlot(1).Number)
		state.WeaponMinDamage = unit.GetGunner().GetMinDamage(unit.GetWeaponSlot(1).Number)
		state.CountFireBullet = unit.GetGunner().GetCountFireBullet(unit.GetWeaponSlot(1).Number)
		state.WeaponMinRange, _ = unit.GetGunner().GetWeaponMinRange(0.0, unit.GetWeaponSlot(1).Number)
		state.WeaponMaxRange, _ = unit.GetGunner().GetWeaponMaxRange(0.0, unit.GetWeaponSlot(1).Number, false)
		state.WeaponAccuracy = unit.GetGunner().GetWeaponAccuracy(unit.GetWeaponSlot(1).Number)
		state.GunRotateSpeed = unit.GetGunner().GetGunRotateSpeed(unit.GetWeaponSlot(1).Number)
		state.BulletSpeed = unit.GetGunner().GetBulletSpeed(unit.GetWeaponSlot(1).Number)
		state.ReloadTime = unit.GetGunner().GetWeaponReloadTime(unit.GetWeaponSlot(1).Number)
		state.ReloadAmmoTime = unit.GetGunner().GetWeaponReloadAmmoTime(unit.GetWeaponSlot(1).Number)

		k, t, e := unit.GetGunner().GetDamageType(unit.GetWeaponSlot(1).Number)
		state.KineticsDamage = k
		state.ThermoDamage = t
		state.ExplosionDamage = e
	}

	return &state
}

// GetAllState метод возвращает полное состоиня машинки, влияние навыков, влияние снаряжения штрафы
func (unit *Unit) GetAllState() (all *StateMS, effects []*effect.Effect) {

	all = unit.GetState()
	effects = unit.GetEffects().GetAllEffects()

	return all, effects
}

func (unit *Unit) GetProtection(typeVulnerabilities string) int {

	var startValue float64

	if typeVulnerabilities == "thermo" {
		startValue = float64(unit.body.ProtectionToThermo)
	}

	if typeVulnerabilities == "kinetics" {
		startValue = float64(unit.body.ProtectionToKinetics)
	}

	if typeVulnerabilities == "explosion" {
		startValue = float64(unit.body.ProtectionToExplosion)
	}

	for _, e := range unit.GetEffects().GetAllEffects() {
		if e.Parameter == "protect_"+typeVulnerabilities {
			free := 100 - startValue
			startValue += free * (float64(e.Quantity) / 100)
		}
	}

	return int(startValue)
}

func (unit *Unit) GetOwnerID() int {
	return unit.OwnerID
}

func (unit *Unit) GetOwnerPlayerID() int {
	return unit.OwnerID
}

func (unit *Unit) GetPower() int {
	return unit.Power
}

func (unit *Unit) GetID() int {
	return unit.ID
}

func (unit *Unit) GetType() string {
	return "unit"
}

func (unit *Unit) GetRotate() float64 {
	return unit.GetPhysicalModel().Rotate
}

func (unit *Unit) GetMapHeight() float64 {
	return _const.UnitHeight
}

func (unit *Unit) GetStateSenderPos() bool {
	return false
}

func (unit *Unit) GetWidth() float64 {
	return float64(unit.body.Width) * 2
}

func (unit *Unit) GetLength() float64 {
	return float64(unit.body.Length) * 2
}

func (unit *Unit) GetHeight() float64 {
	return float64(unit.body.Height) * 2
}

func (unit *Unit) GetGeoData() []*obstacle_point.ObstaclePoint {
	return nil
}

func (unit *Unit) GetWeight() float64 {
	return unit.body.Weight
}

func (unit *Unit) GetMoveMaxPower() float64 {
	return unit.GetEffects().GetAllBodyBonus(unit.body.Speed, "speed", unit.getBody().ChassisType, unit.getBody().StandardSize)
}

func (unit *Unit) GetMaxReverse() float64 {
	return unit.GetEffects().GetAllBodyBonus(unit.body.ReverseSpeed, "reverse_speed", unit.getBody().ChassisType, unit.getBody().StandardSize)
}

func (unit *Unit) GetPowerFactor() float64 {
	return unit.GetEffects().GetAllBodyBonus(unit.body.PowerFactor, "power_factor", unit.getBody().ChassisType, unit.getBody().StandardSize)
}

func (unit *Unit) GetReverseFactor() float64 {
	return unit.GetEffects().GetAllBodyBonus(unit.body.ReverseFactor, "reverse_factor", unit.getBody().ChassisType, unit.getBody().StandardSize)
}

func (unit *Unit) GetTurnSpeed() float64 {
	return unit.GetEffects().GetAllBodyBonus(unit.body.TurnSpeed, "turn_speed", unit.getBody().ChassisType, unit.getBody().StandardSize)
}

func (unit *Unit) GetCapSize() int {
	if unit == nil {
		return 0
	}

	return int(unit.GetEffects().GetAllBodyBonus(float64(unit.body.CapacitySize), "cap_size", unit.getBody().ChassisType, unit.getBody().StandardSize))
}

func (unit *Unit) GetAdditionalCapSize(key string) int {
	if unit == nil {
		return 0
	}

	inv := unit.body.AdditionalInventory[key]
	return int(unit.GetEffects().GetAllBodyBonus(float64(inv.GetCapSize()), key+"_cap_size", unit.getBody().ChassisType, unit.getBody().StandardSize))
}

func (unit *Unit) GetFullFreeCapSize(itemType string, itemID int) int {
	freeSize := unit.GetCapSize() - unit.Inventory.GetSize()

	for key, inv := range unit.AdditionalInventory {
		filter := unit.body.AdditionalInventory[key].Filter
		for fType, ids := range filter {
			if fType == itemType {
				for _, id := range ids {
					if id == itemID {
						freeSize += unit.GetAdditionalCapSize(key) - inv.GetSize()
					}
				}
			}
		}
	}

	return freeSize
}

func (unit *Unit) GetRangeView() int {
	return unit.viewRange
}

func (unit *Unit) getRangeView() int {
	return int(unit.GetEffects().GetAllBodyBonus(float64(unit.body.RangeView), "view", unit.getBody().ChassisType, unit.getBody().StandardSize))
}

func (unit *Unit) GetRadarRange() int {
	return unit.radarRange
}

func (unit *Unit) getRadarRange() int {
	return int(unit.GetEffects().GetAllBodyBonus(float64(unit.body.RangeRadar), "radar", unit.getBody().ChassisType, unit.getBody().StandardSize))
}

func (unit *Unit) GetMaxHP() int {
	return int(unit.GetEffects().GetAllBodyBonus(float64(unit.body.MaxHP), "max_hp", unit.getBody().ChassisType, unit.getBody().StandardSize))
}

func (unit *Unit) GetMaxPower() int {
	return int(unit.GetEffects().GetAllBodyBonus(float64(unit.body.MaxPower), "max_power", unit.getBody().ChassisType, unit.getBody().StandardSize))
}

func (unit *Unit) GetMaxEnergy() int {
	return int(unit.GetEffects().GetAllBodyBonus(float64(unit.body.MaxEnergy), "max_energy", unit.getBody().ChassisType, unit.getBody().StandardSize))
}

// weapon
func (unit *Unit) getGunAccuracy(weaponSlotNumber int) int {
	weaponSlot := unit.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return 0
	}
	return int(unit.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Weapon.Accuracy), "accuracy", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize))
}

func (unit *Unit) getGunRotateSpeed(weaponSlotNumber int) int {
	weaponSlot := unit.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return 0
	}
	return int(unit.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Weapon.RotateSpeed), "gun_speed_rotate", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize))
}

func (unit *Unit) getWeaponReloadTime(weaponSlotNumber int) int {
	weaponSlot := unit.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return 0
	}
	return int(unit.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Weapon.ReloadTime), "reload", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize))
}

func (unit *Unit) getWeaponAmmoReloadTime(weaponSlotNumber int) int {
	weaponSlot := unit.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return 0
	}
	return int(unit.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Weapon.ReloadAmmoTime), "reload_ammo", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize))
}

func (unit *Unit) getMaxDamage(weaponSlotNumber int) int {
	weaponSlot := unit.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil || weaponSlot.Ammo == nil {
		return 0
	}
	return int(unit.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Ammo.MaxDamage), "damage", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize))
}

func (unit *Unit) getMinDamage(weaponSlotNumber int) int {
	weaponSlot := unit.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil || weaponSlot.Ammo == nil {
		return 0
	}
	return int(unit.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Ammo.MinDamage), "damage", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize))
}

func (unit *Unit) AppendWeaponDamageModifier(weaponSlotNumber int) {
	unit.RemoveEffect("weapon_damage_modifier")
	weaponSlot := unit.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot != nil && weaponSlot.Weapon != nil && weaponSlot.Weapon.DamageModifier != 1 {
		damageModifier := 100 - (weaponSlot.Weapon.DamageModifier * 100)
		subtract := damageModifier > 0
		if damageModifier < 0 {
			damageModifier *= -1
		}

		unit.AddEffect(&effect.Effect{
			UUID:        "weapon_damage_modifier",
			Parameter:   "damage",
			Quantity:    int(damageModifier),
			Percentages: true,
			Subtract:    subtract,
		})
	}
}

func (unit *Unit) AppendPassiveEquipModifier() {

	for _, slot := range unit.body.GetAllEquipSlots() {
		unit.RemoveBySlot(slot.Type, slot.Number)
	}

	for _, slot := range unit.body.GetAllEquips() {
		if slot.Equip != nil && !slot.Equip.Active && len(slot.Equip.Effects) > 0 {
			for _, e := range slot.Equip.Effects {
				unit.AddEffect(&effect.Effect{
					UUID:        "equip_passive_" + strconv.Itoa(slot.Type) + ":" + strconv.Itoa(slot.Number) + e.Name,
					Parameter:   e.Parameter,
					Quantity:    e.Quantity,
					Percentages: e.Percentages,
					Subtract:    e.Subtract,
					SlotType:    slot.Type,
					SlotNumber:  slot.Number,
				})
			}
		}
	}
}

func (unit *Unit) GetEvacuation() bool {
	return unit.evacuation
}

func (unit *Unit) GetForceEvacuation() bool {
	return unit.forceEvacuation
}

func (unit *Unit) GetHP() int {
	return unit.HP
}

func (unit *Unit) GetX() int {
	return unit.GetPhysicalModel().GetX()
}

func (unit *Unit) GetY() int {
	return unit.GetPhysicalModel().GetY()
}

func (unit *Unit) GetMapID() int {
	return int(atomic.LoadInt32(&unit.MapID))
}

func (unit *Unit) InSky() bool {
	return unit.inSky
}

func (unit *Unit) GetFullJSON() string {
	jsonUnit, err := json.Marshal(unit)
	if err != nil {
		println("unit to json: ", err.Error())
	}

	return string(jsonUnit)
}

func (unit *Unit) GetFollowTarget() *target.Target {
	mp := unit.movePath
	if mp != nil {
		return mp.GetFollowTarget()
	}

	return nil
}

func (unit *Unit) GetRecoveryPower() int {
	return int(unit.GetEffects().GetAllBodyBonus(float64(unit.getBody().RecoveryPower), "charging_speed", unit.getBody().ChassisType, unit.getBody().StandardSize))
}
