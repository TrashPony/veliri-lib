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
	"sync"
)

type Drone struct {
	UUID             string                           `json:"uuid"`
	ID               int                              `json:"id"`
	Sprite           string                           `json:"sprite"`
	EquipSlot        *detail.BodyEquipSlot            `json:"equip_slot"` // сылка на эквип который запустил дрона
	OwnerPlayerID    int                              `json:"owner_player_id"`
	OwnerType        string                           `json:"owner_type"`
	OwnerID          int                              `json:"owner_id"`
	MaxHP            int                              `json:"max_hp"`
	HP               int                              `json:"hp"`
	MapID            int                              `json:"map_id"`
	RangeView        int                              `json:"range_view"`
	Scale            int                              `json:"scale"`
	Fraction         string                           `json:"fraction"`
	EngagePosNoScale map[string]coordinate.Coordinate `json:"wheels_pos_no_scale"`
	EngageAnchors    map[string]anchor.Anchor         `json:"wheel_anchors"`

	followTarget *target.Target
	weaponTarget *target.Target
	targetMX     sync.Mutex

	CacheJson      []byte `json:"-"`
	CreateJsonTime int64  `json:"-"`

	physicalModel *physical_model.PhysicalModel

	TransportUnitID int                            `json:"transport_unit_id"`
	LifeTime        int                            `json:"life_time"`
	WorkTime        int                            `json:"-"`
	Weapons         map[int]*detail.BodyWeaponSlot `json:"-"`
	effects         *effects_store.EffectsStore
	jobs            []*job
	visibleObjects  *visible_objects.VisibleObjectsStore
	gunner          *gunner.Gunner
	burstOfShots    *burst_of_shots.BurstOfShots
	fractionWarrior bool
	mx              sync.RWMutex
}

func (d *Drone) AddEffect(newEffect *effect.Effect) bool {
	add := d.GetEffects().AddEffect(newEffect)
	d.UpdatePhysicalModel()
	return add
}

func (d *Drone) RemoveEffect(uuid string) bool {
	remove := d.GetEffects().RemoveEffect(uuid)
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

type job struct {
	uuid     string
	target   *target.Target
	job      string
	complete bool
}

func (d *Drone) AddJob(jobType string, target *target.Target) {
	if d.jobs == nil {
		d.jobs = make([]*job, 0)
	}

	d.jobs = append(d.jobs, &job{
		uuid:   uuid.NewV1().String(),
		target: target,
		job:    jobType,
	})
}

func (d *Drone) CompleteJob(uuid string) {
	for _, j := range d.jobs {
		if j.uuid == uuid {
			j.complete = true
		}
	}
}

func (d *Drone) GetJob() (string, string, *target.Target) {
	for _, j := range d.jobs {
		if !j.complete {
			return j.uuid, j.job, j.target
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
			slotState.Accuracy = wSlot.Weapon.Accuracy
			slotState.RotateSpeed = wSlot.Weapon.RotateSpeed
			slotState.ReloadTime = wSlot.Weapon.ReloadTime
			slotState.ReloadAmmoTime = wSlot.Weapon.ReloadAmmoTime
		}

		if wSlot.Ammo != nil {
			slotState.MaxDamage = int(float64(wSlot.Ammo.MaxDamage) * wSlot.Weapon.DamageModifier)
			slotState.MinDamage = int(float64(wSlot.Ammo.MinDamage) * wSlot.Weapon.DamageModifier)
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

func (d *Drone) GetWeaponTarget() *target.Target {
	if d == nil {
		return nil
	}

	d.targetMX.Lock()
	defer d.targetMX.Unlock()
	return d.weaponTarget
}

func (d *Drone) SetWeaponTarget(target *target.Target) {
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

	if d.CreateJsonTime == mapTime && len(d.CacheJson) > 0 {
		return d.CacheJson
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

	command := []byte{}

	command = append(command, game_math.GetIntBytes(d.ID)...)
	command = append(command, game_math.GetIntBytes(d.GetX())...)
	command = append(command, game_math.GetIntBytes(d.GetY())...)
	command = append(command, game_math.GetIntBytes(int(d.GetZ()))...)
	command = append(command, game_math.GetIntBytes(int(d.GetRotate()))...)
	command = append(command, game_math.GetIntBytes(d.MaxHP)...)
	command = append(command, game_math.GetIntBytes(d.HP)...)
	command = append(command, game_math.GetIntBytes(d.RangeView)...)
	command = append(command, byte(d.Scale))
	command = append(command, game_math.GetIntBytes(d.OwnerPlayerID)...)
	command = append(command, byte(lifePercent)) // todo update
	command = append(command, byte(equipType))
	command = append(command, byte(equipSlot))
	command = append(command, game_math.GetIntBytes(d.OwnerID)...)
	command = append(command, byte(_const.MapBinItems[d.OwnerType]))
	command = append(command, game_math.BoolToByte(d.fractionWarrior))
	command = append(command, _const.FractionByte[d.Fraction])
	command = append(command, byte(len(d.Sprite)))
	command = append(command, []byte(d.Sprite)...)

	d.CacheJson = command
	d.CreateJsonTime = mapTime

	return d.CacheJson
}

func (d *Drone) GetUpdateData(mapTime int64) []byte {
	command := []byte{}
	command = append(command, game_math.GetIntBytes(d.HP)...)
	command = append(command, byte(d.GetPercentLife()))
	return command
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

func (d *Drone) UnsafeRangeVisibleObjects() ([]*visible_objects.VisibleObject, *sync.RWMutex) {
	return d.visibleObjectStore().UnsafeRangeMapDynamicObjects()
}

func (d *Drone) GetVisibleObjectByTypeAndID(typeObj string, id int) *visible_objects.VisibleObject {
	return d.visibleObjectStore().GetVisibleObjectByTypeAndID(typeObj, id)
}

func (d *Drone) GetRangeView() int {
	return int(d.GetEffects().GetAllBonus(float64(d.RangeView), "view"))
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
