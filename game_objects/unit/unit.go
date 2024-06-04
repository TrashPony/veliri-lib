package unit

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/burst_of_shots"
	"github.com/TrashPony/veliri-lib/game_objects/coordinate"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"github.com/TrashPony/veliri-lib/game_objects/effect"
	"github.com/TrashPony/veliri-lib/game_objects/effects_store"
	"github.com/TrashPony/veliri-lib/game_objects/gunner"
	"github.com/TrashPony/veliri-lib/game_objects/inventory"
	"github.com/TrashPony/veliri-lib/game_objects/move_path"
	"github.com/TrashPony/veliri-lib/game_objects/physical_model"
	"github.com/TrashPony/veliri-lib/game_objects/target"
	"github.com/TrashPony/veliri-lib/game_objects/visible_objects"
	"sync"
	"time"
)

type Unit struct {
	ID            int    `json:"id"` // TODO int64
	Owner         string `json:"owner"`
	OwnerID       int    `json:"owner_id"`
	CorporationID int    `json:"Corporation_id"`
	OwnerFraction string `json:"of"`
	Ready         bool
	body          *detail.Body

	HP          int    `json:"hp"`
	Power       int    `json:"power"`
	Destroy     bool   `json:"destroy"`
	BodyTexture string `json:"body_texture"`
	MapID       int32  `json:"map_id"`

	evacuation      bool
	forceEvacuation bool
	inSky           bool /* отряд по той или иной причине летит Оо */

	movePath *move_path.MovePath // специальный обьект для пути

	Inventory           *inventory.Inventory            `json:"inventory"` // в роли ключей карты выступают номера слотов где содержиться итем
	AdditionalInventory map[string]*inventory.Inventory `json:"-"`

	weaponTarget    *target.Target
	reservoirTarget *target.Target

	followExit bool

	State *StateMS `json:"-"`

	CacheJson      []byte     `json:"-"`
	CreateJsonTime int64      `json:"-"`
	damageMX       sync.Mutex // специальный мьютекс для получения урона, todo возможно не нужен
	mx             sync.RWMutex
	AutoPilot      bool     `json:"auto_pilot"`
	damage         []Damage `json:"-"`
	Immortal       bool     `json:"-"`
	Interactive    bool     `json:"-"`

	LastDamageTime int64   `json:"-"` // время последнего урона неважно от кого
	LastFireTime   int64   `json:"-"` // время последнего выстрела, включая активные модули
	FearTimer      int     `json:"-"`
	FearAngle      float64 `json:"-"`

	effects        *effects_store.EffectsStore
	visibleObjects *visible_objects.VisibleObjectsStore
	gunner         *gunner.Gunner
	BurstOfShots   *burst_of_shots.BurstOfShots `json:"-"`
	physicalModel  *physical_model.PhysicalModel

	ghost              bool
	lockedControl      bool
	viewRange          int
	radarRange         int
	police             bool
	fractionWarrior    bool
	lights             bool
	Role               string `json:"-"`
	UnrepairableDamage int    `json:"-"`
	FractionByte       byte   `json:"-"`

	Decals []Decal
}

type Decal struct {
	X       int
	Y       int
	DecalID int
	Angle   int
}

func (u *Unit) CheckDecalSlot(x, y int) bool {
	for _, d := range u.Decals {
		if d.X == x && d.Y == y {
			return true
		}
	}

	return false
}

func (u *Unit) AddDecal(x, y, id, angle int) *Decal {
	if u.CheckDecalSlot(x, y) {
		return nil
	}

	if angle < 0 {
		angle += 360
	}

	decal := Decal{
		X:       x,
		Y:       y,
		DecalID: id,
		Angle:   angle / 2,
	}
	u.Decals = append(u.Decals, decal)
	return &decal
}

func (u *Unit) SetLights(lights bool) {
	u.lights = lights
}

func (u *Unit) Lights() bool {
	return u.lights
}

func (u *Unit) Ghost() bool {
	return u.ghost
}

func (u *Unit) SetGhost(g bool) {
	u.ghost = g
}

func (u *Unit) SetPolice(p bool) {
	u.police = p
}

func (u *Unit) GetPolice() bool {
	return u.police
}

func (u *Unit) GetEffects() *effects_store.EffectsStore {
	if u.effects == nil {
		u.effects = &effects_store.EffectsStore{}
	}

	return u.effects
}

func (u *Unit) AddEffect(newEffect *effect.Effect) bool {
	add := u.GetEffects().AddEffect(newEffect)
	if add {
		u.UpdatePhysicalModel()
		u.UpdateWeaponsState()

		if newEffect.Parameter == "max_hp" && !newEffect.Percentages {
			u.HP += newEffect.Quantity
			if u.GetMaxHP() < u.HP {
				u.HP = u.GetMaxHP()
			}
		}
	}

	return add
}

func (u *Unit) RemoveEffect(uuid string) bool {
	remove, ef := u.GetEffects().RemoveEffect(uuid)
	if remove {
		u.UpdatePhysicalModel()
		u.UpdateWeaponsState()

		if ef != nil && ef.Parameter == "max_hp" && !ef.Percentages {
			u.HP -= ef.Quantity
			if u.HP < 1 {
				u.HP = 1
			}
		}
	}

	return remove
}

func (u *Unit) GetMinQuantityByParameter(parameter string) (bool, int) {
	return u.GetEffects().GetMinQuantityByParameter(parameter)
}

func (u *Unit) Invisibility() bool {
	ok, quantity := u.GetMinQuantityByParameter("invisibility")
	return ok && quantity <= 0
}

func (u *Unit) RemoveBySlot(slotType, slotNumber int) bool {
	remove, remEffects := u.GetEffects().RemoveBySlot(slotType, slotNumber)
	if remove {
		u.UpdatePhysicalModel()
		u.UpdateWeaponsState()

		for _, ef := range remEffects {
			if ef != nil && ef.Parameter == "max_hp" && !ef.Percentages {
				u.HP -= ef.Quantity
				if u.HP < 1 {
					u.HP = 1
				}
			}
		}
	}

	return remove
}

func (u *Unit) GetEffectByUUID(uuid string) *effect.Effect {
	return u.GetEffects().GetEffectByUUID(uuid)
}

func (u *Unit) GetGunner() *gunner.Gunner {
	if u.gunner == nil {
		u.initGunner()
	}

	return u.gunner
}

func (u *Unit) initGunner() {
	u.gunner = &gunner.Gunner{
		GunUser:          u,
		WeaponSlotsState: make([]*gunner.WeaponSlotState, 0),
	}

	u.UpdateWeaponsState()
}

func (u *Unit) UpdateViewState() {
	u.viewRange = u.getRangeView()
	u.radarRange = u.getRadarRange()
}

func (u *Unit) UpdateWeaponsState() {
	if u.gunner == nil {
		u.initGunner()
	}

	u.viewRange = u.getRangeView()
	u.radarRange = u.getRadarRange()

	u.gunner.WeaponSlotsState = make([]*gunner.WeaponSlotState, 0)
	for _, wSlot := range u.RangeWeaponSlots() {

		slotState := &gunner.WeaponSlotState{
			Number: wSlot.Number,
		}

		if wSlot.Weapon != nil {
			slotState.Accuracy = u.getGunAccuracy(wSlot.Number)
			slotState.RotateSpeed = u.getGunRotateSpeed(wSlot.Number)
			slotState.ReloadTime = u.getWeaponReloadTime(wSlot.Number)
			slotState.ReloadAmmoTime = u.getWeaponAmmoReloadTime(wSlot.Number)
		}

		if wSlot.Ammo != nil {
			slotState.MaxDamage = u.getMaxDamage(wSlot.Number)
			slotState.MinDamage = u.getMinDamage(wSlot.Number)
		}

		u.gunner.WeaponSlotsState = append(u.gunner.WeaponSlotsState, slotState)
	}
}

func (u *Unit) GetMovePathTime() int64 {
	mp := u.movePath
	if mp == nil || !mp.GetNeedCalc() {
		return time.Now().UnixNano() + int64(time.Hour)
	}

	return mp.GetMovePathTime()
}

func (u *Unit) GetMovePathState() (bool, string, float64, *target.Target, *[]*coordinate.Coordinate, int, bool, int64) {
	mp := u.movePath
	if mp == nil {
		return false, "", 0, nil, nil, 0, false, 0
	}

	return mp.GetMovePathState()
}

func (u *Unit) NextMovePoint() {
	mp := u.movePath
	if mp == nil {
		return
	}

	mp.NextMovePoint()
}

func (u *Unit) SetFindMovePath() {
	mp := u.movePath
	if mp != nil {
		mp.SetFindMovePath()
	}
}

func (u *Unit) RemoveMovePath() {
	u.movePath = nil
}

func (u *Unit) SetMovePath(path *[]*coordinate.Coordinate) {
	mp := u.movePath
	if mp == nil {
		return
	}

	mp.SetMovePath(path)
}

func (u *Unit) SetMovePathTarget(t *target.Target) {
	mp := &move_path.MovePath{}
	mp.SetMovePathTarget(t, u.GetOwnerPlayerID(), u.ID)
	u.movePath = mp
}

func (u *Unit) SetMovePathAngle(angle float64) {
	mp := &move_path.MovePath{}
	mp.SetMovePathAngle(angle, u.GetOwnerPlayerID(), u.ID)
	u.movePath = mp
}

type Damage struct {
	PlayerID   int   `json:"player_id"`
	TimeDamage int64 `json:"time_damage"`
	Damage     int   `json:"damage"`
}

func (u *Unit) FillLastDamage(d []Damage) {
	u.mx.Lock()
	defer u.mx.Unlock()
	u.damage = d
}

func (u *Unit) GetArrayLastDamage() []Damage {
	u.mx.Lock()
	defer u.mx.Unlock()
	d := make([]Damage, len(u.damage))

	for i, v := range u.damage {
		d[i] = v
	}

	return d
}

func (u *Unit) SetLastDamage(playerID, damage int) *Damage {
	u.mx.Lock()
	defer u.mx.Unlock()

	if playerID == u.OwnerID {
		return nil
	}

	newDamage := Damage{
		PlayerID:   playerID,
		TimeDamage: time.Now().Unix(),
		Damage:     damage,
	}
	u.damage = append(u.damage, newDamage)
	return &newDamage
}

func (u *Unit) GetLastDamage() int {
	u.mx.Lock()
	defer u.mx.Unlock()

	if len(u.damage) == 0 {
		return 0
	}

	last := u.damage[len(u.damage)-1]
	// если урон наносился больше чем 30 сек назад то он не учитывается
	if time.Now().Unix()-last.TimeDamage > 30 {
		u.damage = u.damage[:0]
		return 0
	}

	return last.PlayerID
}

func (u *Unit) GetLastDamageByPlayerID(playerID, sec int) int {
	u.mx.Lock()
	defer u.mx.Unlock()

	if len(u.damage) == 0 {
		return 0
	}

	allDamage := 0
	for _, d := range u.damage {
		if d.PlayerID == playerID {
			if time.Now().Unix()-d.TimeDamage < int64(sec) {
				allDamage += d.Damage
			}
		}
	}

	return allDamage
}

func (u *Unit) GetAllDamage() map[int]int {
	u.mx.Lock()
	defer u.mx.Unlock()

	r := make(map[int]int)

	if len(u.damage) == 0 {
		return r
	}

	for _, d := range u.damage {
		if time.Now().Unix()-d.TimeDamage < 30 {
			r[d.PlayerID] += d.Damage
		}
	}

	return r
}

func (u *Unit) GetPhysicalModel() *physical_model.PhysicalModel {

	if u.physicalModel == nil {
		u.initPhysicalModel() // инициируем по умолчания, и ток в методе UpdatePhysicalModel уже докидываем скилы и тд
	}

	return u.physicalModel
}

func (u *Unit) LockedControl() bool {
	return u.lockedControl
}

func (u *Unit) SetLockedControl(l bool) {
	u.lockedControl = l
}

func (u *Unit) UpdatePhysicalModel() {

	if u.physicalModel == nil {
		u.initPhysicalModel() // инициируем по умолчания, и ток в методе UpdatePhysicalModel уже докидываем скилы и тд
	}

	u.physicalModel.ID = u.ID
	u.physicalModel.Speed = u.GetMoveMaxPower() / _const.ServerTickSecPart
	u.physicalModel.ReverseSpeed = u.GetMaxReverse() / _const.ServerTickSecPart
	u.physicalModel.PowerFactor = u.GetPowerFactor() / _const.ServerTickSecPart
	u.physicalModel.ReverseFactor = u.GetReverseFactor() / _const.ServerTickSecPart
	u.physicalModel.TurnSpeed = u.GetTurnSpeed() / _const.ServerTickSecPart
	u.physicalModel.Weight = u.GetWeight()
}

func (u *Unit) initPhysicalModel() {
	u.physicalModel = &physical_model.PhysicalModel{
		Speed:         u.GetMoveMaxPower() / _const.ServerTickSecPart,
		ReverseSpeed:  u.GetMaxReverse() / _const.ServerTickSecPart,
		PowerFactor:   u.GetPowerFactor() / _const.ServerTickSecPart,
		ReverseFactor: u.GetReverseFactor() / _const.ServerTickSecPart,
		TurnSpeed:     u.GetTurnSpeed() / _const.ServerTickSecPart,
		MoveDrag:      u.GetBody().MoveDrag,
		AngularDrag:   u.GetBody().AngularDrag,
		Weight:        float64(u.GetBody().CapacitySize),
		Type:          "unit",
		ID:            u.GetID(),
		ChassisType:   u.body.ChassisType,
		MoveDestroyer: u.body.Fraction == _const.FAUNA,
	}

	// применяем настройки размера к обьектам колизий
	sizeOffset := float64(u.GetBody().Scale) / 100
	u.physicalModel.Height = float64(u.GetBody().Height) * sizeOffset
	u.physicalModel.Length = float64(u.GetBody().Length) * sizeOffset
	u.physicalModel.Width = float64(u.GetBody().Width) * sizeOffset
	u.physicalModel.Radius = int(float64(u.GetBody().Radius) * sizeOffset)
}

func (u *Unit) GetCopyPhysicalModel() *physical_model.PhysicalModel {
	origPH := u.GetPhysicalModel()
	pm := *origPH
	return &pm
}

func (u *Unit) LockDamage() {
	u.damageMX.Lock()
}

func (u *Unit) UnlockDamage() {
	u.damageMX.Unlock()
}

func (u *Unit) SetAllGunRotate(addRotate float64) {
	for _, weaponSlot := range u.getBody().Weapons {
		if weaponSlot != nil {
			weaponSlot.SetGunRotate(weaponSlot.GetGunRotate() + addRotate)
		}
	}
}

func (u *Unit) GetTransportUnit() *Unit {
	return nil
}

func (u *Unit) GetNeedZ() float64 {
	return 0
}

func (u *Unit) GetZ() float64 {
	return 0
}

func (u *Unit) SetZ(float64) {
}

func (u *Unit) IsFly() bool {
	return false
}

type ShortUnitInfo struct {
	ID int `json:"id"`
	/* позиция */
	GunRotate float64 `json:"gun_rotate"`
	Rotate    float64 `json:"rotate"`
	X         int     `json:"x"`
	Y         int     `json:"y"`

	/*видимый фит*/
	Body        detail.Body                   `json:"body"`
	WeaponSlots map[int]detail.BodyWeaponSlot `json:"weapon_slots"`
	EquipSlots  []detail.BodyEquipSlot        `json:"equip_slots"`

	HP int `json:"hp"`

	/* покраска юнитов */
	BodyColor1   string `json:"body_color_1"`
	BodyColor2   string `json:"body_color_2"`
	WeaponColor1 string `json:"weapon_color_1"`
	WeaponColor2 string `json:"weapon_color_2"`

	/* путь к файлу готовой покраске, пока не реализовано */
	BodyTexture   string `json:"body_texture"`
	WeaponTexture string `json:"weapon_texture"`

	OwnerID         int       `json:"owner_id"`
	Owner           string    `json:"owner"`
	OwnerFraction   string    `json:"of"`
	MapID           int       `json:"map_id"`
	Evacuation      bool      `json:"evacuation"`
	ForceEvacuation bool      `json:"force_evacuation"`
	InSky           bool      `json:"in_sky"` /* отряд по той или иной причине летит Оо */
	MoveChecker     bool      `json:"move_checker"`
	ActualPathCell  *PathUnit `json:"actual_path_cell"`

	CacheJson      string `json:"-"`
	CreateJsonTime int64  `json:"-"`
}

type PathUnit struct {
	X               int     `json:"x"`
	Y               int     `json:"y"`
	Z               float64 `json:"z"`
	Rotate          float64 `json:"rotate"`
	RotateGun       float64 `json:"rotate_gun"`
	Millisecond     int     `json:"millisecond"`
	Speed           float64
	Direction       bool    `json:"direction"`
	AngularVelocity float64 `json:"angular_velocity"`
	Traversed       bool    `json:"traversed"`
	Animate         bool    `json:"animate"` // включить или нет анимацию движения гусениц
}

type Slot struct {
	Unit       *Unit `json:"unit"`
	NumberSlot int   `json:"number_slot"`
}

func (u *Unit) GetJSON(mapTime int64) []byte {

	if u.CreateJsonTime == mapTime && len(u.CacheJson) > 0 {
		return u.CacheJson
	}

	if u.CacheJson == nil {
		u.CacheJson = []byte{}
	}

	u.CacheJson = u.CacheJson[:0]
	u.CacheJson = append(u.CacheJson, game_math.GetIntBytes(u.GetID())...)
	u.CacheJson = append(u.CacheJson, game_math.GetIntBytes(u.OwnerID)...)
	u.CacheJson = append(u.CacheJson, game_math.GetIntBytes(u.HP)...)
	u.CacheJson = append(u.CacheJson, game_math.GetIntBytes(u.GetBody().ID)...)

	// position data
	u.CacheJson = append(u.CacheJson, game_math.GetIntBytes(u.GetX())...)
	u.CacheJson = append(u.CacheJson, game_math.GetIntBytes(u.GetY())...)
	u.CacheJson = append(u.CacheJson, game_math.GetIntBytes(int(u.GetRotate()))...)

	u.CacheJson = append(u.CacheJson, game_math.GetIntBytes(u.GetMaxHP())...)
	u.CacheJson = append(u.CacheJson, game_math.GetIntBytes(u.GetRangeView())...)
	u.CacheJson = append(u.CacheJson, game_math.GetIntBytes(u.GetRadarRange())...)

	u.CacheJson = append(u.CacheJson, game_math.BoolToByte(u.fractionWarrior))
	u.CacheJson = append(u.CacheJson, _const.FractionByte[u.OwnerFraction])

	u.CacheJson = append(u.CacheJson, game_math.GetIntBytes(u.GetMaxPower())...)
	u.CacheJson = append(u.CacheJson, game_math.BoolToByte(u.Invisibility()))
	u.CacheJson = append(u.CacheJson, game_math.BoolToByte(u.Interactive))

	u.CacheJson = append(u.CacheJson, game_math.GetIntBytes(u.CorporationID)...)
	u.CacheJson = append(u.CacheJson, game_math.BoolToByte(u.ghost))
	u.CacheJson = append(u.CacheJson, game_math.BoolToByte(u.police))
	u.CacheJson = append(u.CacheJson, game_math.BoolToByte(u.lights))

	u.CacheJson = append(u.CacheJson, byte(len([]byte(u.GetBody().Texture))))
	u.CacheJson = append(u.CacheJson, []byte(u.GetBody().Texture)...)

	u.CacheJson = append(u.CacheJson, byte(len([]byte(u.Owner))))
	u.CacheJson = append(u.CacheJson, []byte(u.Owner)...)

	// decal data
	decalData := []byte{}
	for _, d := range u.Decals {
		decalData = append(decalData, byte(d.X))
		decalData = append(decalData, byte(d.Y))
		decalData = append(decalData, byte(d.DecalID))
		decalData = append(decalData, byte(d.Angle))
	}

	u.CacheJson = append(u.CacheJson, game_math.GetIntBytes(len(decalData))...)
	u.CacheJson = append(u.CacheJson, decalData...)

	// weapon data
	weaponData := []byte{}
	for _, unitWeaponSlot := range u.RangeWeaponSlots() {

		weaponData = append(weaponData, byte(unitWeaponSlot.Number))
		weaponData = append(weaponData, game_math.GetIntBytes(unitWeaponSlot.GetRealXAttach())...)
		weaponData = append(weaponData, game_math.GetIntBytes(unitWeaponSlot.GetRealYAttach())...)
		weaponData = append(weaponData, byte(unitWeaponSlot.GetXAnchor()*100))
		weaponData = append(weaponData, byte(unitWeaponSlot.GetYAnchor()*100))
		weaponData = append(weaponData, game_math.GetIntBytes(int(unitWeaponSlot.GetGunRotate()))...)

		if unitWeaponSlot.Weapon != nil {
			weaponData = append(weaponData, byte(len([]byte(unitWeaponSlot.WeaponTexture))))
			weaponData = append(weaponData, []byte(unitWeaponSlot.WeaponTexture)...)
		} else {
			weaponData = append(weaponData, byte(0))
		}
	}

	u.CacheJson = append(u.CacheJson, game_math.GetIntBytes(len(weaponData))...)
	u.CacheJson = append(u.CacheJson, weaponData...)

	// equip data
	equipData := []byte{}
	for _, equipSlot := range u.GetBody().GetAllEquips() {
		if equipSlot.Equip != nil && (equipSlot.Equip.XAttach != 0 || equipSlot.Equip.YAttach != 0) {

			equipData = append(equipData, byte(equipSlot.Number))
			equipData = append(equipData, byte(equipSlot.Type))
			equipData = append(equipData, game_math.GetIntBytes(equipSlot.GetRealXAttach())...)
			equipData = append(equipData, game_math.GetIntBytes(equipSlot.GetRealYAttach())...)
			equipData = append(equipData, byte(equipSlot.GetXAnchor()*100))
			equipData = append(equipData, byte(equipSlot.GetYAnchor()*100))

			equipData = append(equipData, byte(len([]byte(equipSlot.Equip.Name))))
			equipData = append(equipData, []byte(equipSlot.Equip.Name)...)
		}
	}

	u.CacheJson = append(u.CacheJson, game_math.GetIntBytes(len(equipData))...)
	u.CacheJson = append(u.CacheJson, equipData...)

	return u.CacheJson
}

func (u *Unit) GetUpdateData(mapTime int64) []byte {

	command := []byte{}
	command = append(command, game_math.GetIntBytes(u.HP)...)
	command = append(command, game_math.GetIntBytes(u.GetRangeView())...)
	command = append(command, game_math.GetIntBytes(u.GetRadarRange())...)
	command = append(command, game_math.BoolToByte(u.ghost))
	command = append(command, game_math.BoolToByte(u.lights))

	return command
}

func (u *Unit) GetShortInfo() *ShortUnitInfo {
	if u == nil || u.getBody() == nil {
		return nil
	}

	var hostile ShortUnitInfo

	hostile.X = u.GetPhysicalModel().GetX()
	hostile.Y = u.GetPhysicalModel().GetY()

	hostile.GunRotate = u.GetGunner().GetGunRotate(1)
	hostile.Rotate = u.GetRotate()
	hostile.MapID = u.GetMapID()
	hostile.Evacuation = u.GetEvacuation()
	hostile.ForceEvacuation = u.GetForceEvacuation()
	hostile.InSky = u.InSky()
	hostile.OwnerFraction = u.OwnerFraction

	hostile.Body = detail.Body{
		Name:       u.getBody().Name,
		MaxHP:      u.getBody().MaxHP,
		RangeView:  u.getBody().RangeView,
		RangeRadar: u.getBody().RangeRadar,
		Scale:      u.getBody().Scale,
		Length:     int(u.GetPhysicalModel().GetLength()),
		Width:      int(u.GetPhysicalModel().GetWidth()),
	}

	hostile.WeaponSlots = make(map[int]detail.BodyWeaponSlot)
	for number, wSlot := range u.RangeWeaponSlots() {
		if wSlot != nil {
			hostile.WeaponSlots[number] = detail.BodyWeaponSlot{
				Number:        wSlot.Number,
				Weapon:        wSlot.Weapon,
				XAttach:       wSlot.XAttach,
				YAttach:       wSlot.YAttach,
				RealXAttach:   wSlot.GetRealXAttach(),
				RealYAttach:   wSlot.GetRealYAttach(),
				XAnchor:       wSlot.GetXAnchor(),
				YAnchor:       wSlot.GetYAnchor(),
				GunRotate:     wSlot.GetGunRotate(),
				WeaponColor1:  wSlot.WeaponColor1,
				WeaponColor2:  wSlot.WeaponColor2,
				WeaponTexture: wSlot.WeaponTexture,
			}
		}
	}

	hostile.OwnerID = u.GetOwnerID()
	hostile.Owner = u.Owner
	hostile.ID = u.GetID()
	hostile.BodyTexture = u.BodyTexture

	hostile.HP = u.GetHP()

	hostile.EquipSlots = make([]detail.BodyEquipSlot, 0)

	copyEquips := func(realEquips map[int]*detail.BodyEquipSlot) {
		for _, equipSlot := range realEquips {

			if equipSlot.GetXAnchor() > 0 || equipSlot.GetYAnchor() > 0 {
				hostile.EquipSlots = append(hostile.EquipSlots, detail.BodyEquipSlot{
					Type:         equipSlot.Type,
					Number:       equipSlot.Number,
					Equip:        equipSlot.Equip,
					StandardSize: equipSlot.StandardSize,
					RealXAttach:  equipSlot.GetRealXAttach(),
					RealYAttach:  equipSlot.GetRealYAttach(),
					XAnchor:      equipSlot.GetXAnchor(),
					YAnchor:      equipSlot.GetYAnchor(),
				})
			}
		}
	}

	copyEquips(u.getBody().EquippingI)
	copyEquips(u.getBody().EquippingII)
	copyEquips(u.getBody().EquippingIII)
	copyEquips(u.getBody().EquippingIV)
	copyEquips(u.getBody().EquippingV)

	return &hostile
}

func (u *Unit) DelEquip() {

}

func (u *Unit) DelAmmo() {

}

func (u *Unit) SetEquip() {

}

func (u *Unit) SetAmmo() {

}

func (u *Unit) GetOwnerUser() string {
	return u.Owner
}

func (u *Unit) GetAmmoCount() int { // по диз доку оружие в юните может быть только одно
	for _, weaponSlot := range u.getBody().Weapons {
		if weaponSlot.Weapon != nil {
			return weaponSlot.GetAmmoQuantity()
		}
	}

	return 0
}

func (u *Unit) SetAnchorsEquip() {
	// метод расчитывает правильные положения якорец и точек прикрепления для спрайтов снаряжения на корпусе
	if u == nil {
		return
	}

	// распологаем оружие
	for _, ws := range u.RangeWeaponSlots() {
		ws.SetAnchor()
	}

	// распологаем снарягу (не вся снаряга может быть на корпусе но это разруливается не тут)
	for _, slot := range u.getBody().GetAllEquips() {
		slot.SetAnchor()
	}
}

func (u *Unit) GetWeaponSlot(slotNumber int) *detail.BodyWeaponSlot { // по диз доку оружие в юните может быть только одно

	if u == nil || u.getBody() == nil {
		return nil
	}

	return u.getBody().Weapons[slotNumber]
}

func (u *Unit) GetEquipPosInMap(typeEquip, numberSlot int) (int, int) {

	equipSlot := u.GetBodyEquipSlot(typeEquip, numberSlot)
	if equipSlot == nil || equipSlot.Equip == nil {
		return 0, 0
	}

	return game_math.GetWeaponPosInMap(
		u.GetX(), u.GetY(), u.getBody().Scale,
		float64(equipSlot.XAttach),
		float64(equipSlot.YAttach),
		u.GetRotate())
}

func (u *Unit) GetEquipFirePos(typeEquip, numberSlot int) []*game_math.Positions {

	equipSlot := u.GetBodyEquipSlot(typeEquip, numberSlot)
	if equipSlot == nil || equipSlot.Equip == nil {
		return nil
	}

	return game_math.GetWeaponFirePositions(
		u.GetX(), u.GetY(), u.GetBodyScale(), u.GetRotate(), equipSlot.Rotate,
		equipSlot.Equip.XAttach, equipSlot.Equip.YAttach,
		equipSlot.Equip.FirePositions,
		float64(equipSlot.XAttach),
		float64(equipSlot.YAttach),
	)
}

func (u *Unit) CheckViewCoordinate(x, y, radius int) (bool, bool) {

	if u == nil || u.getBody() == nil {
		return false, false
	}

	if u.GetRangeView()+radius >= int(game_math.GetBetweenDist(u.GetX(), u.GetY(), x, y)) {
		return true, true
	}

	if u.GetRadarRange()+radius >= int(game_math.GetBetweenDist(u.GetX(), u.GetY(), x, y)) {
		return false, true
	}

	return false, false
}

func (u *Unit) RangeWeaponSlots() map[int]*detail.BodyWeaponSlot {
	// мы никогда не пишут в карту слотов оружия поэтому этот метод безопасен (по крайне мере пока)
	return u.getBody().Weapons
}

func (u *Unit) GetScale() int {
	return u.getBody().Scale
}

func (u *Unit) GetBurstOfShots() *burst_of_shots.BurstOfShots {
	if u.BurstOfShots == nil {
		u.BurstOfShots = &burst_of_shots.BurstOfShots{}
	}

	return u.BurstOfShots
}

func (u *Unit) SetVisibleObjectStore(v *visible_objects.VisibleObjectsStore) {
	u.visibleObjects = v
}

func (u *Unit) GetVisibleObjectStore() *visible_objects.VisibleObjectsStore {
	return u.visibleObjects
}

func (u *Unit) visibleObjectStore() *visible_objects.VisibleObjectsStore {
	if u.visibleObjects == nil {
		u.visibleObjects = &visible_objects.VisibleObjectsStore{}
	}

	return u.visibleObjects
}

func (u *Unit) GetVisibleObjects() <-chan *visible_objects.VisibleObject {
	return u.visibleObjectStore().GetVisibleObjects()
}

func (u *Unit) UnsafeRangeVisibleObjects() ([]*visible_objects.VisibleObject, *sync.RWMutex) {
	return u.visibleObjectStore().UnsafeRangeMapDynamicObjects()
}

func (u *Unit) GetVisibleObjectByTypeAndID(typeObj string, id int) *visible_objects.VisibleObject {
	return u.visibleObjectStore().GetVisibleObjectByTypeAndID(typeObj, id)
}

func (u *Unit) SetFractionWarrior(ok bool) {
	u.fractionWarrior = ok
}

func (u *Unit) FractionWarrior() bool {
	return u.fractionWarrior
}

func (u *Unit) GetCorporationID() int {
	return u.CorporationID
}

func (u *Unit) SetCorporationID(id int) {
	u.CorporationID = id
}
