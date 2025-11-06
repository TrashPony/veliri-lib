package drone

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/anchor"
	"github.com/TrashPony/veliri-lib/game_objects/burst_of_shots"
	"github.com/TrashPony/veliri-lib/game_objects/coordinate"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"github.com/TrashPony/veliri-lib/game_objects/effect"
	"github.com/TrashPony/veliri-lib/game_objects/effects_store"
	"github.com/TrashPony/veliri-lib/game_objects/gunner"
	"github.com/TrashPony/veliri-lib/game_objects/physical_model"
	"github.com/TrashPony/veliri-lib/game_objects/target"
	"github.com/TrashPony/veliri-lib/game_objects/visible_objects"
	uuid "github.com/satori/go.uuid"
	"math"
	"sync"
)

type Drone struct {
	UUID             string                           `json:"uuid"`
	ID               int                              `json:"id"`
	Sprite           string                           `json:"sprite"`
	EquipSlot        *detail.BodyEquipSlot            `json:"equip_slot"` // сылка на эквип который запустил дрона
	OwnerPlayerID    int                              `json:"owner_player_id"`
	CorporationID    int                              `json:"corporation_id"`
	OwnerType        string                           `json:"owner_type"`
	OwnerID          int                              `json:"owner_id"`
	MaxHP            int                              `json:"max_hp"`
	HP               int                              `json:"hp"`
	MapID            int                              `json:"map_id"`
	RangeView        int                              `json:"range_view"`
	Scale            int                              `json:"scale"`
	Fraction         string                           `json:"fraction"`
	FractionByte     byte                             `json:"-"`
	EngagePosNoScale map[string]coordinate.Coordinate `json:"wheels_pos_no_scale"`
	EngageAnchors    map[string]anchor.Anchor         `json:"wheel_anchors"`
	Independent      bool                             `json:"-"`
	WarpMode         WarpMode                         `json:"-"`

	followTarget *target.Target
	weaponTarget *target.Target
	targetMX     sync.Mutex

	CacheCreateData CacheData `json:"-"`
	CacheUpdateData CacheData `json:"-"`

	physicalModel *physical_model.PhysicalModel

	TransportUnitID int                            `json:"transport_unit_id"`
	LifeTime        int                            `json:"life_time"`
	WorkTime        int                            `json:"-"`
	Weapons         map[int]*detail.BodyWeaponSlot `json:"-"`
	effects         *effects_store.EffectsStore
	jobs            []*Job
	visibleObjects  *visible_objects.VisibleObjectsStore
	gunner          *gunner.Gunner
	burstOfShots    *burst_of_shots.BurstOfShots
	fractionWarrior bool
	mx              sync.RWMutex
}

type WarpMode struct {
	Status          string
	Path            []*coordinate.Coordinate
	CurrentPosition int
}

type CacheData struct {
	Data []byte `json:"-"`
	Time int64  `json:"-"`
}

func (d *Drone) GetOwnerPlayerID() int {
	return d.OwnerPlayerID
}

func (d *Drone) GetCorporationID() int {
	return d.CorporationID
}

func (d *Drone) SetCorporationID(id int) {
	d.CorporationID = id
}

func (d *Drone) AddEffect(newEffect *effect.Effect) bool {
	add := d.GetEffects().AddEffect(newEffect)
	if add {
		d.UpdatePhysicalModel()
		d.UpdateWeaponsState()
	}

	return add
}

func (d *Drone) RemoveEffect(uuid string) bool {
	remove, _ := d.GetEffects().RemoveEffect(uuid)
	d.UpdatePhysicalModel()
	return remove
}

func (d *Drone) GetEffects() *effects_store.EffectsStore {
	if d.effects == nil {
		d.effects = &effects_store.EffectsStore{}
	}

	return d.effects
}

func (d *Drone) GetMapHeight() float64 {
	return d.GetPhysicalModel().Z - 10
}

func (d *Drone) GetScale() int {
	return d.Scale
}

func (d *Drone) GetBurstOfShots() *burst_of_shots.BurstOfShots {
	if d.burstOfShots == nil {
		d.burstOfShots = &burst_of_shots.BurstOfShots{}
	}

	return d.burstOfShots
}

func (d *Drone) GetType() string {
	return "drone"
}

type Job struct {
	uuid     string
	Target   *target.Target `json:"target"`
	Job      string         `json:"job"`
	Complete bool           `json:"complete"`
}

func (d *Drone) AddJob(jobType string, target *target.Target) {
	if d.jobs == nil {
		d.jobs = make([]*Job, 0)
	}

	d.jobs = append(d.jobs, &Job{
		uuid:   uuid.NewV1().String(),
		Target: target,
		Job:    jobType,
	})
}

func (d *Drone) CompleteAllJobJob(jobType string) {
	for _, j := range d.jobs {
		if j.Job == jobType {
			j.Complete = true
		}
	}
}

func (d *Drone) CompleteJob(uuid string) {
	for _, j := range d.jobs {
		if j.uuid == uuid {
			j.Complete = true
		}
	}
}

func (d *Drone) GetJob() (string, string, *target.Target) {
	for _, j := range d.jobs {
		if !j.Complete {
			return j.uuid, j.Job, j.Target
		}
	}

	return "", "", nil
}

func (d *Drone) GetGunner() *gunner.Gunner {
	if d.gunner == nil {
		d.initGunner()
	}

	return d.gunner
}

func (d *Drone) GetMeleer() *gunner.Meleer {
	return nil
}

func (d *Drone) initGunner() {
	d.gunner = &gunner.Gunner{
		GunUser:          d,
		WeaponSlotsState: make([]*gunner.WeaponSlotState, 0),
	}

	d.UpdateWeaponsState()
}

func (d *Drone) UpdateWeaponsState() {
	if d.gunner == nil {
		d.initGunner()
	}

	d.gunner.WeaponSlotsState = make([]*gunner.WeaponSlotState, 0)
	for _, wSlot := range d.RangeWeaponSlots() {

		slotState := &gunner.WeaponSlotState{
			Number: wSlot.Number,
		}

		if wSlot.Weapon != nil {
			slotState.Accuracy = d.getGunAccuracy(wSlot.Number)
			slotState.RotateSpeed = d.getGunRotateSpeed(wSlot.Number)
			slotState.ReloadTime = d.getWeaponReloadTime(wSlot.Number)
			slotState.ReloadAmmoTime = d.getWeaponAmmoReloadTime(wSlot.Number)
		}

		if wSlot.Ammo != nil {
			slotState.MaxDamage = d.getMaxDamage(wSlot.Number)
			slotState.MinDamage = d.getMinDamage(wSlot.Number)
		}

		d.gunner.WeaponSlotsState = append(d.gunner.WeaponSlotsState, slotState)
	}
}

func (d *Drone) GetPhysicalModel() *physical_model.PhysicalModel {

	if d.physicalModel == nil {
		d.initPhysicalModel()
	}

	return d.physicalModel
}

func (d *Drone) LockedControl() bool {
	return false
}

func (d *Drone) Ghost() bool {
	return false
}

func (d *Drone) SetPhysicalModel(p *physical_model.PhysicalModel) {
	d.physicalModel = p
}

func (d *Drone) initPhysicalModel() {
	d.physicalModel = &physical_model.PhysicalModel{
		Speed:         0.75,
		ReverseSpeed:  0.5,
		PowerFactor:   0.15,
		ReverseFactor: 0.1,
		TurnSpeed:     0.01,
		MoveDrag:      0.9,
		AngularDrag:   0.9,
		Fly:           true,
		Type:          "drone",
		ID:            d.GetID(),
		ChassisType:   "fly",
	}
}

func (d *Drone) UpdatePhysicalModel() {
	// todo обонление параметров типо изменилась скорость из за скила например
}

func (d *Drone) GetID() int {
	return d.ID
}

func (d *Drone) SetAllGunRotate(addRotate float64) {
}

func (d *Drone) GetWeaponTarget(_ int) *target.Target {
	if d == nil {
		return nil
	}

	d.targetMX.Lock()
	defer d.targetMX.Unlock()
	return d.weaponTarget
}

func (d *Drone) SetWeaponTarget(_ int, target *target.Target) {
	d.targetMX.Lock()
	defer d.targetMX.Unlock()

	if target != nil {
		target.Attack = true
	}

	d.weaponTarget = target
}

func (d *Drone) GetX() int {
	return d.GetPhysicalModel().X
}

func (d *Drone) GetY() int {
	return d.GetPhysicalModel().Y
}

func (d *Drone) GetRotate() float64 {
	return d.GetPhysicalModel().Rotate
}

func (d *Drone) GetGunRotate() float64 {
	return 0
}

func (d *Drone) SetGunRotate(gunRotate float64) {
}

func (d *Drone) SetFollowTarget(target *target.Target) {
	d.targetMX.Lock()
	d.followTarget = target
	d.targetMX.Unlock()
}

func (d *Drone) GetFollowTarget() *target.Target {
	return d.followTarget
}

func (d *Drone) GetZ() float64 {
	return d.GetPhysicalModel().GetZ()
}

func (d *Drone) GetJSON(mapTime int64) []byte {

	if d.CacheCreateData.Time == mapTime && len(d.CacheCreateData.Data) > 0 {
		return d.CacheCreateData.Data
	}

	lifePercent := 0
	equipType := 0
	equipSlot := 0

	if d.EquipSlot != nil {

		equipType = d.EquipSlot.Type
		equipSlot = d.EquipSlot.Number

		if d.WorkTime > 0 || d.LifeTime > 0 {
			lifePercent = int((float64(d.WorkTime) / float64(d.EquipSlot.Equip.LifeTime)) * 100)
		}
	}

	if d.CacheCreateData.Data == nil {
		d.CacheCreateData.Data = []byte{}
	}

	d.CacheCreateData.Data = d.CacheCreateData.Data[:0]

	d.CacheCreateData.Data = append(d.CacheCreateData.Data, game_math.GetIntBytes(d.ID)...)
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, game_math.GetIntBytes(d.GetX())...)
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, game_math.GetIntBytes(d.GetY())...)
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, game_math.GetIntBytes(int(d.GetZ()))...)
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, game_math.GetIntBytes(int(d.GetRotate()))...)
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, game_math.GetIntBytes(d.MaxHP)...)
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, game_math.GetIntBytes(d.HP)...)
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, game_math.GetIntBytes(d.RangeView)...)
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, byte(d.Scale))
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, game_math.GetIntBytes(d.OwnerPlayerID)...)
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, byte(lifePercent)) // todo update
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, byte(equipType))
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, byte(equipSlot))
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, game_math.GetIntBytes(d.OwnerID)...)
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, byte(_const.MapBinItems[d.OwnerType]))
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, game_math.BoolToByte(d.fractionWarrior))
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, _const.FractionByte[d.Fraction])
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, game_math.GetIntBytes(d.CorporationID)...)
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, byte(len(d.Sprite)))
	d.CacheCreateData.Data = append(d.CacheCreateData.Data, []byte(d.Sprite)...)

	d.CacheCreateData.Time = mapTime

	return d.CacheCreateData.Data
}

func (d *Drone) GetUpdateData(mapTime int64) []byte {

	if d.CacheUpdateData.Time == mapTime && len(d.CacheUpdateData.Data) > 0 {
		return d.CacheUpdateData.Data
	}

	if d.CacheUpdateData.Data == nil {
		d.CacheUpdateData.Data = []byte{}
	}

	d.CacheUpdateData.Data = d.CacheUpdateData.Data[:0]

	d.CacheUpdateData.Data = append(d.CacheUpdateData.Data, game_math.GetIntBytes(d.HP)...)
	d.CacheUpdateData.Data = append(d.CacheUpdateData.Data, byte(d.GetPercentLife()))

	d.CacheUpdateData.Time = mapTime

	return d.CacheUpdateData.Data
}

func (d *Drone) GetPercentLife() int {
	if d.EquipSlot != nil {
		if d.WorkTime > 0 || d.LifeTime > 0 {
			return int((float64(d.WorkTime) / float64(d.EquipSlot.Equip.LifeTime)) * 100)
		}
	}

	return -1
}

func (d *Drone) GetWeaponSlot(slotNumber int) *detail.BodyWeaponSlot {
	return d.Weapons[slotNumber]
}

func (d *Drone) RangeWeaponSlots() map[int]*detail.BodyWeaponSlot {
	// мы никогда не пишут в карту слотов оружия поэтому этот метод безопасен (по крайне мере пока)
	return d.Weapons
}

func (d *Drone) RangeMeleeWeaponSlots() map[int]*detail.BodyWeaponSlot {
	return map[int]*detail.BodyWeaponSlot{}
}

func (d *Drone) SetVisibleObjectStore(v *visible_objects.VisibleObjectsStore) {
	d.visibleObjects = v
}

func (d *Drone) GetVisibleObjectStore() *visible_objects.VisibleObjectsStore {
	return d.visibleObjects
}

func (d *Drone) visibleObjectStore() *visible_objects.VisibleObjectsStore {
	if d.visibleObjects == nil {
		d.visibleObjects = &visible_objects.VisibleObjectsStore{}
	}

	return d.visibleObjects
}

func (d *Drone) GetVisibleObjects() <-chan *visible_objects.VisibleObject {
	return d.visibleObjectStore().GetVisibleObjects()
}

func (d *Drone) UnsafeRangeVisibleObjects() ([][]*visible_objects.VisibleObject, *sync.RWMutex) {
	return d.visibleObjectStore().UnsafeRangeMapDynamicObjects()
}

func (d *Drone) GetVisibleObjectByTypeAndID(typeObj string, id int) *visible_objects.VisibleObject {
	return d.visibleObjectStore().GetVisibleObjectByTypeAndID(typeObj, id)
}

func (d *Drone) GetRangeView() int {
	return int(math.Ceil(d.GetEffects().GetAllBonus(float64(d.RangeView), "view")))
}

func (d *Drone) CheckViewCoordinate(x, y, radius int) (bool, bool) {
	if d.GetRangeView()+radius >= int(game_math.GetBetweenDist(d.GetPhysicalModel().X, d.GetPhysicalModel().Y, x, y)) {
		return true, true
	}

	return false, false
}

func (d *Drone) SetFractionWarrior(ok bool) {
	d.fractionWarrior = ok
}

func (d *Drone) FractionWarrior() bool {
	return d.fractionWarrior
}

func (d *Drone) SetEngagePositions() {

	if d.EngageAnchors == nil {
		d.EngageAnchors = make(map[string]anchor.Anchor)
	}

	for key, pos := range d.EngagePosNoScale {
		xAnchor, yAnchor, realXAttach, realYAttach := game_math.GetAnchorWeapon(64, 64, pos.X, pos.Y)
		newAnc := anchor.Anchor{
			XAnchor:     xAnchor,
			YAnchor:     yAnchor,
			RealXAttach: realXAttach,
			RealYAttach: realYAttach,
		}

		ancState, ok := d.EngageAnchors[key]
		if ok {
			newAnc.Type = ancState.Type
			newAnc.Scale = ancState.Scale
			newAnc.Rotate = ancState.Rotate
			newAnc.Height = ancState.Height
		}

		d.EngageAnchors[key] = newAnc
	}
}

func (d *Drone) getGunAccuracy(weaponSlotNumber int) int {
	weaponSlot := d.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return 0
	}
	return int(math.Ceil(d.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Weapon.Accuracy), "accuracy", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (d *Drone) getGunRotateSpeed(weaponSlotNumber int) int {
	weaponSlot := d.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return 0
	}
	return int(math.Ceil(d.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Weapon.RotateSpeed), "gun_speed_rotate", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (d *Drone) getWeaponReloadTime(weaponSlotNumber int) int {
	weaponSlot := d.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return 0
	}
	return int(math.Ceil(d.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Weapon.ReloadTime), "reload", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (d *Drone) getWeaponAmmoReloadTime(weaponSlotNumber int) int {
	weaponSlot := d.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil {
		return 0
	}
	return int(math.Ceil(d.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Weapon.ReloadAmmoTime), "reload_ammo", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (d *Drone) getMaxDamage(weaponSlotNumber int) int {
	weaponSlot := d.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil || weaponSlot.Ammo == nil {
		return 0
	}
	return int(math.Ceil(d.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Ammo.MaxDamage), "damage", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (d *Drone) getMinDamage(weaponSlotNumber int) int {
	weaponSlot := d.GetWeaponSlot(weaponSlotNumber)
	if weaponSlot == nil || weaponSlot.Weapon == nil || weaponSlot.Ammo == nil {
		return 0
	}
	return int(math.Ceil(d.GetEffects().GetAllWeaponBonus(float64(weaponSlot.Ammo.MinDamage), "damage", weaponSlot.Weapon.Type, weaponSlot.Weapon.StandardSize)))
}

func (d *Drone) OwnerFraction() string {
	return d.Fraction
}
