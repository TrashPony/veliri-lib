package dynamic_map_object

import (
	"encoding/json"
	"fmt"
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/access_manager"
	"github.com/TrashPony/veliri-lib/game_objects/ammo"
	"github.com/TrashPony/veliri-lib/game_objects/burst_of_shots"
	"github.com/TrashPony/veliri-lib/game_objects/coordinate"
	"github.com/TrashPony/veliri-lib/game_objects/damage_manager"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"github.com/TrashPony/veliri-lib/game_objects/gunner"
	"github.com/TrashPony/veliri-lib/game_objects/inventory"
	"github.com/TrashPony/veliri-lib/game_objects/move_path"
	"github.com/TrashPony/veliri-lib/game_objects/obstacle_point"
	"github.com/TrashPony/veliri-lib/game_objects/physical_model"
	"github.com/TrashPony/veliri-lib/game_objects/special_hostiles"
	"github.com/TrashPony/veliri-lib/game_objects/target"
	"github.com/TrashPony/veliri-lib/game_objects/visible_objects"
	"github.com/TrashPony/veliri-lib/generate_ids"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Object struct {
	// везде где есть приставка Type это оригиналдьные данные типа, все остальное сформированые
	ID                  int    `json:"id"`
	TypeID              int    `json:"type_id"`
	Type                string `json:"type"`
	MapID               int    `json:"map_id"`
	Texture             string `json:"texture"`
	AnimateSpriteSheets string `json:"animate_sprite_sheets"`
	AnimateLoop         bool   `json:"animate_loop"`
	Name                string `json:"name"`
	Inventory           bool   `json:"inventory"`
	InventoryCapacity   int    `json:"-"`
	BoxID               int    `json:"-"`
	MaxHP               int    `json:"max_hp"`
	TypeMaxHP           int    `json:"-"`
	HP                  int    `json:"hp"`
	Scale               int    `json:"scale"`
	MaxScale            int    `json:"-"` // определяется рандомно для растений максимальный размер куста
	GrowCollision       bool   `json:"-"`
	GrowTime            int    `json:"grow_time"` // говорит время цикла когда растение росло для гладкой отрисовки
	Shadow              bool   `json:"shadow"`
	AnimationSpeed      int    `json:"animation_speed"`
	Priority            int    `json:"priority"`
	Weight              int    `json:"weight"`

	TypeXShadowOffset int `json:"-"`
	XShadowOffset     int `json:"x_shadow_offset"`
	TypeYShadowOffset int `json:"-"`
	YShadowOffset     int `json:"y_shadow_offset"`
	ShadowIntensity   int `json:"shadow_intensity"`

	TypeGeoData   []*obstacle_point.ObstaclePoint `json:"-"`
	TypeBaseData  []*coordinate.Coordinate        `json:"type_base_data"`
	AccessManager access_manager.AccessManager    `json:"-"`
	HeightType    float64                         `json:"-"`

	Fraction     string `json:"fraction"`
	FractionByte byte   `json:"-"`
	RangeView    int    `json:"range_view"`
	RangeRadar   int    `json:"range_radar"`
	XAttach      int    `json:"x_attach"`
	YAttach      int    `json:"y_attach"`

	/* постройка */
	Build           bool              `json:"build"` // если билд true, то обьект считает завершенным если Complete == 100, иначе он не работает
	Immortal        bool              `json:"immortal"`
	AutoBuild       bool              `json:"auto_build"`
	AutoDismantling bool              `json:"auto_dismantling"`
	Complete        float64           `json:"complete"`    // процент завершенности
	StartItems      []*inventory.Slot `json:"start_items"` // необхоидимые ресурсы для производства на старте
	NeedItems       []*inventory.Slot `json:"need_items"`  // необходимо ресурсов для завершения,

	/* турели и оружие*/
	WeaponID      int  `json:"weapon_id"`
	AmmoCell      bool `json:"ammo_cell"`
	weaponTarget  *target.Target
	OwnerPlayerID int    `json:"owner_player_id"`
	CorporationID int    `json:"corporation_id"`
	OwnerType     string `json:"owner_type"`
	OwnerID       int    `json:"owner_id"` // ид игрока владельца

	/* энергия и снаряжение */
	EnergyCell           *detail.ThoriumSlot `json:"energy_cell"`
	Energy               int                 `json:"energy"`
	EnergyOn             bool                `json:"energy_on"`
	CurrentEnergy        int                 `json:"current_energy"`
	MaxEnergy            int                 `json:"max_energy"`
	EquipID              int                 `json:"equip_id"`
	SpecialCell          bool                `json:"special_cell"`
	Interactive          bool                `json:"interactive"`
	placeUserSpecialCell map[string]bool

	Weapons map[int]*detail.BodyWeaponSlot `json:"weapons"`
	Equips  map[int]*detail.BodyEquipSlot  `json:"equips"`

	viewRange  int
	radarRange int

	/* Экстрактор ресурсов */
	ReservoirID int `json:"-"`
	/* Для кусков туш */
	BodyID int `json:"-"`
	/* ловушки */
	TrapUnits []*TrapUnit `json:"trap_units"`

	MoveDestroyer bool `json:"-"`

	OwnerBaseID int  `json:"-"`
	Run         bool `json:"-"`
	Work        bool `json:"work"`
	Disable     bool `json:"-"`
	Static      bool `json:"-"`
	LifeTime    int  `json:"-"`

	specialHostiles *special_hostiles.SpecialHostiles

	radarWorkerWork bool
	radarWorkerExit bool

	memoryDynamicObjects *visible_objects.VisibleObjectsStore
	visibleObjects       *visible_objects.VisibleObjectsStore

	PosFunc      func() `json:"-"`
	MemoryID     int    `json:"-"`
	GrowLeftTime int    `json:"-"`

	CacheJson              []byte `json:"-"`
	CreateJsonTime         int64  `json:"-"`
	countUpdateViewObjects int

	DestroyLeftTimer bool   `json:"-"`
	DestroyTimer     int    `json:"-"`
	CacheGeoData     []byte `json:"-"`

	attributes              map[string]int
	countUpdateWeaponTarget int
	fractionWarrior         bool
	physicalModel           *physical_model.PhysicalModel
	NoAutoDestroy           bool `json:"-"`
	gunner                  *gunner.Gunner
	damageManager           damage_manager.DamageManager
	burstOfShots            *burst_of_shots.BurstOfShots `json:"-"`
	ToPath                  move_path.To                 `json:"-"`
	BossWreckage            bool                         `json:"-"`
	ForceView               bool                         `json:"-"` // видим обьект даже через туман войны, но не получаем его обзор
	MissionKey              string                       `json:"-"`
	mx                      sync.RWMutex
}

type TrapUnit struct {
	ID       int  `json:"id"`
	BreakOut int  `json:"breakOut"`
	Skip     int  `json:"skip"`
	Attack   bool `json:"attack"`
}

func (o *Object) GetFraction() string {
	return o.Fraction
}

func (o *Object) Reset() {
	o.gunner = nil
	o.physicalModel = nil
	o.burstOfShots = nil
	o.visibleObjects = nil
	o.memoryDynamicObjects = nil
	o.specialHostiles = nil
	o.updateViewState()
	o.mx = sync.RWMutex{}
}

func (o *Object) GetGeoDataBin() []byte {
	return o.CacheGeoData
}

func (o *Object) GetMapHeight() float64 {
	return o.GetPhysicalModel().GetHeight()
}

func (o *Object) GetGunner() *gunner.Gunner {
	if o.gunner == nil {
		o.initGunner()
	}

	return o.gunner
}

func (o *Object) initGunner() {
	o.gunner = &gunner.Gunner{
		GunUser:          o,
		WeaponSlotsState: make([]*gunner.WeaponSlotState, 0),
	}

	o.UpdateWeaponsState()
}

func (o *Object) updateViewState() {
	o.viewRange = o.getRangeView()
	o.radarRange = o.getRadarRange()
}

func (o *Object) UpdateWeaponsState() {
	if o.gunner == nil {
		o.initGunner()
	}

	o.updateViewState()
	o.gunner.WeaponSlotsState = make([]*gunner.WeaponSlotState, 0)
	for _, wSlot := range o.RangeWeaponSlots() {

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

		o.gunner.WeaponSlotsState = append(o.gunner.WeaponSlotsState, slotState)
	}
}

func (o *Object) GetPhysicalModel() *physical_model.PhysicalModel {
	if o.physicalModel == nil {
		o.initPhysicalModel()
	}

	return o.physicalModel
}

func (o *Object) LockedControl() bool {
	return false
}

func (o *Object) Ghost() bool {
	return false
}

func (o *Object) initPhysicalModel() {

	o.updateViewState()
	weight := float64(o.Weight)
	if o.Static {
		weight = float64(math.MaxInt32)
	}

	o.physicalModel = &physical_model.PhysicalModel{
		MoveDrag:    0.7,
		AngularDrag: 0.7,
		Type:        "object",
		ID:          o.ID,
		Static:      o.Static,
		Weight:      weight,
	}
}

func (o *Object) IgnoreSmoke() bool {
	return false
}

func (o *Object) GetRole() string {
	return "object"
}

func (o *Object) SetAllGunRotate(addRotate float64) {
	o.GetWeaponSlot(1).SetGunRotate(o.GetWeaponSlot(1).GetGunRotate() + addRotate)
}

func (o *Object) GetVisibleObjectByTypeAndID(typeObj string, id int) *visible_objects.VisibleObject {
	o.checkVisibleObjectStore()
	return o.visibleObjects.GetVisibleObjectByTypeAndID(typeObj, id)
}

func (o *Object) InitVisibleObjects() {
	o.checkVisibleObjectStore()
	o.visibleObjects.InitVisibleObjects()
}

func (o *Object) GetVisibleObjects() <-chan *visible_objects.VisibleObject {
	o.checkVisibleObjectStore()
	return o.visibleObjects.GetVisibleObjects()
}

func (o *Object) AddVisibleObject(newObj *visible_objects.VisibleObject) {
	o.checkVisibleObjectStore()
	o.visibleObjects.AddVisibleObject(newObj)
}

func (o *Object) RemoveVisibleObject(removeObj *visible_objects.VisibleObject) {
	o.checkVisibleObjectStore()
	o.visibleObjects.RemoveVisibleObject(removeObj)
}

func (o *Object) UnsafeRangeVisibleObjects() ([][]*visible_objects.VisibleObject, *sync.RWMutex) {
	o.checkVisibleObjectStore()
	return o.visibleObjects.UnsafeRangeMapDynamicObjects()
}

func (o *Object) checkVisibleObjectStore() {
	if o.visibleObjects == nil {
		o.visibleObjects = &visible_objects.VisibleObjectsStore{}
	}
}

func (o *Object) AddDynamicObject(object *Object, mapID int, view, radar, force bool, mapTime int64) {
	o.checkMemoryObjectStore()
	if object.MemoryID == 0 {
		object.MemoryID = generate_ids.GetMarkID()
	}

	vObj := &visible_objects.VisibleObject{
		ID:            object.MemoryID,
		IDObject:      object.ID,
		TypeObject:    "object",
		Scale:         object.GetScale(),
		Energy:        object.GetPower(),
		Complete:      object.GetComplete(),
		HP:            object.GetHP(),
		MapID:         mapID,
		X:             object.GetX(),
		Y:             object.GetY(),
		Rotate:        int(object.GetRotate()),
		OwnerID:       object.GetOwnerID(),
		OwnerPlayerID: object.OwnerPlayerID,
		CorporationID: object.CorporationID,
		Object:        object,
		ObjectJSON:    object.GetJSON(mapTime),
		Work:          object.GetWork(),
		ForceView:     force,
	}

	vObj.SetView(view || force)
	vObj.SetRadar(radar || force)

	o.memoryDynamicObjects.AddDynamicObject(vObj)
}

func (o *Object) GetMapDynamicObjectByID(id int) *visible_objects.VisibleObject {
	o.checkMemoryObjectStore()
	return o.memoryDynamicObjects.GetMapDynamicObjectByID(id)
}

func (o *Object) RemoveDynamicObject(id int) {
	o.checkMemoryObjectStore()
	o.memoryDynamicObjects.RemoveDynamicObject(id)
}
func (o *Object) GetMapDynamicObjects(mapID int) <-chan *visible_objects.VisibleObject {
	o.checkMemoryObjectStore()
	return o.memoryDynamicObjects.GetMapDynamicObjects(mapID)
}

func (o *Object) UnsafeRangeMapDynamicObjects() ([][]*visible_objects.VisibleObject, *sync.RWMutex) {
	o.checkMemoryObjectStore()
	return o.memoryDynamicObjects.UnsafeRangeMapDynamicObjects()
}

func (o *Object) checkMemoryObjectStore() {
	if o.memoryDynamicObjects == nil {
		o.memoryDynamicObjects = &visible_objects.VisibleObjectsStore{}
	}
}

func (o *Object) GetUpdateViewObjects() bool {
	if o.countUpdateViewObjects == 3 {
		o.countUpdateViewObjects = 0
		return true
	} else {
		o.countUpdateViewObjects++
		return false
	}
}

func (o *Object) ExitAllWorker() {
	// работа обзора и радара отряда
	if o.GetRadarWorkerWork() {
		o.SetRadarWorkerExit(true)
		for o.GetRadarWorkerWork() {
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (o *Object) IsFly() bool {
	return false
}

func (o *Object) CheckGrowthPower() bool {
	return false
}

func (o *Object) CheckGrowthRevers() bool {
	return false
}

func (o *Object) CheckLeftRotate() bool {
	return false
}

func (o *Object) CheckRightRotate() bool {
	return false
}

func (o *Object) AddHostile(typeHostile string, id int, hatePoint int, mod string) {
	if o.specialHostiles == nil {
		o.specialHostiles = &special_hostiles.SpecialHostiles{}
	}

	o.specialHostiles.AddPoints(typeHostile, id, hatePoint, mod)
}

func (o *Object) SetHostilePoints(typeHostile string, id, hatePoint int) {
	if o.specialHostiles == nil {
		o.specialHostiles = &special_hostiles.SpecialHostiles{}
	}

	o.specialHostiles.SetPoints(typeHostile, id, hatePoint)
}

func (o *Object) CheckHostile(typeHostile string, id int, ownerID int) (bool, int) {

	if o == nil {
		return false, 0
	}

	if o.specialHostiles == nil {
		return false, 0
	}

	// хозяина ненадо бить)
	if typeHostile == "unit" && ownerID > 0 && ownerID == o.OwnerID {
		return false, 0
	}

	return o.specialHostiles.CheckHostile(typeHostile, id)
}

func (o *Object) SetPosFunc(fun func()) {
	o.PosFunc = fun
}

func (o *Object) GetTurretAmmo(turretAmmo *ammo.Ammo) {
	o.GetWeaponSlot(1).SetAmmo(turretAmmo)
	o.GetWeaponSlot(1).AmmoQuantity = o.GetWeaponSlot(1).Weapon.AmmoCapacity
}

func (o *Object) GetJSON(mapTime int64) []byte {

	if o.CreateJsonTime == mapTime && len(o.CacheJson) > 0 {
		return o.CacheJson
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in GetJSON object", r)
		}
	}()

	/*

				let template = {
				  id: 0,
			      x: 0,
			      y: 0,
				  rotate: 0,
		          height: 0,

				  hp: 0,
				  max_hp: 0,
				  current_energy: 0,
				  max_energy: 0,
				  view_range:0,

				  x_shadow_offset: 0,
				  y_shadow_offset: 0,
				  shadow_intensity: 0,
				  owner_id: 0,
				  priority: 0,

				  work: false,         // byte
				  build: false,        // byte
				  scale: 0,            // byte
				  shadow: false,       // byte
				  animate: false,      // byte
				  animation_speed: 0,  // byte
				  animate_loop: false, // byte

				  type: "",
				  texture: "",

				  weapons: [{
				    gun_rotate: 0,
				    real_x_attach: 0,
				    real_y_attach: 0,
				    number: 0,   // byte
				    x_anchor: 0, // byte
				    y_anchor: 0, // byte
				  }],
				  geo_data: [{
				    x: 0,
				    y: 0,
				    radius: 0,
				  }],
				}

	*/

	command := []byte{}

	command = append(command, game_math.GetIntBytes(o.ID)...)
	command = append(command, game_math.GetIntBytes(o.GetX())...)
	command = append(command, game_math.GetIntBytes(o.GetY())...)
	command = append(command, game_math.GetIntBytes(int(o.GetRotate()))...)
	command = append(command, game_math.GetIntBytes(int(o.GetMapHeight()))...)

	command = append(command, game_math.GetIntBytes(o.HP)...)
	command = append(command, game_math.GetIntBytes(o.MaxHP)...)
	command = append(command, game_math.GetIntBytes(o.CurrentEnergy)...)
	command = append(command, game_math.GetIntBytes(o.MaxEnergy)...)
	command = append(command, game_math.GetIntBytes(o.GetRangeView())...)

	command = append(command, game_math.GetIntBytes(o.XShadowOffset)...)
	command = append(command, game_math.GetIntBytes(o.YShadowOffset)...)
	command = append(command, game_math.GetIntBytes(o.ShadowIntensity)...)
	command = append(command, game_math.GetIntBytes(o.OwnerPlayerID)...)
	command = append(command, game_math.GetIntBytes(o.Priority)...)

	command = append(command, game_math.BoolToByte(o.Work))
	command = append(command, game_math.BoolToByte(o.Build))
	command = append(command, byte(o.Scale))
	command = append(command, game_math.BoolToByte(o.Shadow))
	command = append(command, game_math.BoolToByte(o.AnimateSpriteSheets != ""))
	command = append(command, byte(o.AnimationSpeed))
	command = append(command, game_math.BoolToByte(o.AnimateLoop))
	command = append(command, byte(o.Complete))
	command = append(command, game_math.BoolToByte(o.Static))
	command = append(command, game_math.BoolToByte(o.fractionWarrior))
	command = append(command, _const.FractionByte[o.Fraction])
	command = append(command, game_math.BoolToByte(o.Inventory || o.SpecialCell || o.Interactive || (o.Build && o.Complete < 100))) // interactive
	command = append(command, game_math.GetIntBytes(o.CorporationID)...)
	command = append(command, game_math.BoolToByte(o.ForceView))

	command = append(command, byte(len(o.Type)))
	command = append(command, []byte(o.Type)...)

	if o.AnimateSpriteSheets != "" {
		command = append(command, byte(len(o.AnimateSpriteSheets)))
		command = append(command, []byte(o.AnimateSpriteSheets)...)
	} else {
		command = append(command, byte(len(o.Texture)))
		command = append(command, []byte(o.Texture)...)
	}

	// weapons
	command = append(command, byte(len(o.Weapons)))
	for _, weaponSlot := range o.RangeWeaponSlots() {
		command = append(command, game_math.GetIntBytes(int(weaponSlot.GetGunRotate()))...)
		command = append(command, game_math.GetIntBytes(weaponSlot.GetRealXAttach())...)
		command = append(command, game_math.GetIntBytes(weaponSlot.GetRealYAttach())...)
		command = append(command, byte(weaponSlot.Number))
		command = append(command, byte(weaponSlot.GetXAnchor()*100))
		command = append(command, byte(weaponSlot.GetYAnchor()*100))

		if weaponSlot.Weapon != nil {
			command = append(command, byte(len([]byte(weaponSlot.Weapon.Name))))
			command = append(command, []byte(weaponSlot.Weapon.Name)...)
		} else {
			command = append(command, byte(0))
		}
	}

	// geo data
	geoData := o.GetGeoDataBin()
	command = append(command, game_math.GetIntBytes(len(geoData))...)
	command = append(command, geoData...)

	// equip data
	equipData := []byte{}
	for _, equipSlot := range o.Equips {
		if equipSlot.Equip != nil && (equipSlot.Equip.XAttach != 0 || equipSlot.Equip.YAttach != 0) {

			equipData = append(equipData, byte(equipSlot.Number))
			equipData = append(equipData, game_math.GetIntBytes(int(equipSlot.Rotate))...)
			equipData = append(equipData, game_math.GetIntBytes(equipSlot.GetRealXAttach())...)
			equipData = append(equipData, game_math.GetIntBytes(equipSlot.GetRealYAttach())...)
			equipData = append(equipData, byte(equipSlot.GetXAnchor()*100))
			equipData = append(equipData, byte(equipSlot.GetYAnchor()*100))

			equipData = append(equipData, byte(len([]byte(equipSlot.Equip.Name))))
			equipData = append(equipData, []byte(equipSlot.Equip.Name)...)
		}
	}

	command = append(command, game_math.GetIntBytes(len(equipData))...)
	command = append(command, equipData...)

	o.CacheJson = command
	o.CreateJsonTime = mapTime

	return o.CacheJson
}

func (o *Object) GetUpdateData(mapTime int64) []byte {
	command := []byte{}

	command = append(command, game_math.GetIntBytes(o.HP)...)
	command = append(command, game_math.GetIntBytes(o.CurrentEnergy)...)
	command = append(command, game_math.GetIntBytes(o.Scale)...)
	command = append(command, game_math.GetIntBytes(o.GetRangeView())...)
	command = append(command, game_math.BoolToByte(o.Work))
	command = append(command, byte(o.Complete))

	return command
}

type Flore struct {
	TextureOverFlore string `json:"texture_over_flore"`
	TexturePriority  int    `json:"texture_priority"`
	X                int    `json:"x"`
	Y                int    `json:"y"`
	Rotate           int    `json:"r"`
}

func (o *Object) SetVisibleObjectStore(v *visible_objects.VisibleObjectsStore) {
	o.visibleObjects = v
}

func (o *Object) GetVisibleObjectStore() *visible_objects.VisibleObjectsStore {
	return o.visibleObjects
}

func (o *Object) GetUpdateWeaponTarget() bool {
	if o.countUpdateWeaponTarget == 30 {
		o.countUpdateWeaponTarget = 0
		return true
	} else {
		o.countUpdateWeaponTarget++
		return false
	}
}

func (o *Object) GetMaxHP() int {
	if o.Build {
		return int((o.GetComplete() / float64(100)) * float64(o.MaxHP))
	}

	return o.MaxHP
}

func (o *Object) PlaceItemToSpecialCell(playerID, actionNumber int) {
	o.mx.Lock()
	defer o.mx.Unlock()

	if o.placeUserSpecialCell == nil {
		o.placeUserSpecialCell = make(map[string]bool)
	}

	o.placeUserSpecialCell[strconv.Itoa(playerID)+":"+strconv.Itoa(actionNumber)] = true
}

func (o *Object) GetPlayersPlaceItemToSpecialCell(playerID, actionNumber int) bool {
	o.mx.Lock()
	defer o.mx.Unlock()

	if o.placeUserSpecialCell == nil {
		o.placeUserSpecialCell = make(map[string]bool)
	}

	return o.placeUserSpecialCell[strconv.Itoa(playerID)+":"+strconv.Itoa(actionNumber)]
}

func (o *Object) SetFractionWarrior(ok bool) {
	o.fractionWarrior = ok
}

func (o *Object) FractionWarrior() bool {
	return o.fractionWarrior
}

func (o *Object) BuildMXLock() {
	o.mx.Lock()
}

func (o *Object) BuildMXUnlock() {
	o.mx.Unlock()
}

func (o *Object) GetCorporationID() int {
	return o.CorporationID
}

func (o *Object) SetCorporationID(id int) {
	o.CorporationID = id
}

func (o *Object) OwnerFraction() string {
	return o.Fraction
}

func (o *Object) RangeHostiles() <-chan *special_hostiles.SpecialHostile {
	if o.specialHostiles == nil {
		o.specialHostiles = &special_hostiles.SpecialHostiles{}
	}

	return o.specialHostiles.RangeHostiles()
}

// TODO hardCode
func (o *Object) GetTurretAmmoID() int {
	o.mx.Lock()
	defer o.mx.Unlock()

	if o.attributes != nil && o.attributes["ammo_id"] > 0 {
		return o.attributes["ammo_id"]
	}

	for _, w := range o.Weapons {
		for _, aa := range w.Weapon.AllowAmmo {
			return aa
		}
	}

	if strings.Contains(o.Texture, "laser_turret") {
		return 6
	}
	if strings.Contains(o.Texture, "tank_turret") {
		return 2
	}
	if strings.Contains(o.Texture, "replic_gauss_gun") {
		return 16
	}
	if strings.Contains(o.Texture, "artillery_turret") {
		return 39
	}

	return 0
}

func (o *Object) UnmarshalAttributes(data []byte) {
	o.mx.Lock()
	defer o.mx.Unlock()
	_ = json.Unmarshal(data, &o.attributes)
}

func (o *Object) SetAttributes(attributes map[string]int) {
	o.mx.Lock()
	defer o.mx.Unlock()
	o.attributes = attributes
}

func (o *Object) GetAttributes() map[string]int {
	o.mx.Lock()
	defer o.mx.Unlock()

	attr := make(map[string]int)
	for k, v := range o.attributes {
		attr[k] = v
	}
	return attr
}

func (o *Object) GetAttribute(key string) int {
	o.mx.Lock()
	defer o.mx.Unlock()

	if o.attributes == nil {
		o.attributes = make(map[string]int)
	}

	return o.attributes[key]
}

func (o *Object) HasKey(key string) bool {
	o.mx.Lock()
	defer o.mx.Unlock()

	_, ok := o.attributes[key]
	return ok
}

func (o *Object) SetAttribute(key string, v int) {
	o.mx.Lock()
	defer o.mx.Unlock()

	if o.attributes == nil {
		o.attributes = make(map[string]int)
	}

	o.attributes[key] = v
}
