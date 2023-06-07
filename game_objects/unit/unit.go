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
	"github.com/TrashPony/veliri-lib/game_objects/physical_model"
	"github.com/TrashPony/veliri-lib/game_objects/target"
	"github.com/TrashPony/veliri-lib/game_objects/visible_objects"
	"sync"
	"time"
)

type Unit struct {
	ID            int    `json:"id"`
	Owner         string `json:"owner"`
	OwnerID       int    `json:"owner_id"`
	OwnerFraction string `json:"of"`

	body *detail.Body

	HP          int    `json:"hp"`
	Power       int    `json:"power"`
	Destroy     bool   `json:"destroy"`
	BodyTexture string `json:"body_texture"`
	MapID       int32  `json:"map_id"`

	evacuation      bool
	forceEvacuation bool
	inSky           bool /* отряд по той или иной причине летит Оо */

	movePath *MovePath // специальный обьект для пути
	moveMx   sync.Mutex

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
	Damage         []Damage `json:"-"`
	Immortal       bool     `json:"-"`

	LastDamageTime int64   `json:"-"` // время последнего урона неважно от кого
	LastFireTime   int64   `json:"-"` // время последнего выстрела, включая активные модули
	FearTimer      int     `json:"-"`
	FearAngle      float64 `json:"-"`

	effects        *effects_store.EffectsStore
	visibleObjects *visible_objects.VisibleObjectsStore
	gunner         *gunner.Gunner
	BurstOfShots   *burst_of_shots.BurstOfShots `json:"-"`
	physicalModel  *physical_model.PhysicalModel

	viewRange  int
	radarRange int

	fractionWarrior bool
}

func (unit *Unit) GetEffects() *effects_store.EffectsStore {
	if unit.effects == nil {
		unit.effects = &effects_store.EffectsStore{}
	}

	return unit.effects
}

func (unit *Unit) AddEffect(newEffect *effect.Effect) bool {
	add := unit.GetEffects().AddEffect(newEffect)
	if add {
		unit.UpdatePhysicalModel()
		unit.UpdateWeaponsState()
	}

	return add
}

func (unit *Unit) RemoveEffect(uuid string) bool {
	remove := unit.GetEffects().RemoveEffect(uuid)
	if remove {
		unit.UpdatePhysicalModel()
		unit.UpdateWeaponsState()
	}

	return remove
}

func (unit *Unit) RemoveBySlot(slotType, slotNumber int) bool {
	// TODO API
	remove := unit.GetEffects().RemoveBySlot(slotType, slotNumber)
	if remove {
		unit.UpdatePhysicalModel()
		unit.UpdateWeaponsState()
	}

	return remove
}

func (unit *Unit) GetEffectByUUID(uuid string) *effect.Effect {
	return unit.GetEffects().GetEffectByUUID(uuid)
}

func (unit *Unit) GetGunner() *gunner.Gunner {
	if unit.gunner == nil {
		unit.initGunner()
	}

	return unit.gunner
}

func (unit *Unit) initGunner() {
	unit.gunner = &gunner.Gunner{
		GunUser:          unit,
		WeaponSlotsState: make([]*gunner.WeaponSlotState, 0),
	}

	unit.UpdateWeaponsState()
}

func (unit *Unit) UpdateViewState() {
	unit.viewRange = unit.getRangeView()
	unit.radarRange = unit.getRadarRange()
}

func (unit *Unit) UpdateWeaponsState() {
	if unit.gunner == nil {
		unit.initGunner()
	}

	unit.viewRange = unit.getRangeView()
	unit.radarRange = unit.getRadarRange()

	unit.gunner.WeaponSlotsState = make([]*gunner.WeaponSlotState, 0)
	for _, wSlot := range unit.RangeWeaponSlots() {

		slotState := &gunner.WeaponSlotState{
			Number: wSlot.Number,
		}

		if wSlot.Weapon != nil {
			slotState.Accuracy = unit.getGunAccuracy(wSlot.Number)
			slotState.RotateSpeed = unit.getGunRotateSpeed(wSlot.Number)
			slotState.ReloadTime = unit.getWeaponReloadTime(wSlot.Number)
			slotState.ReloadAmmoTime = unit.getWeaponAmmoReloadTime(wSlot.Number)
		}

		if wSlot.Ammo != nil {
			slotState.MaxDamage = unit.getMaxDamage(wSlot.Number)
			slotState.MinDamage = unit.getMinDamage(wSlot.Number)
		}

		unit.gunner.WeaponSlotsState = append(unit.gunner.WeaponSlotsState, slotState)
	}
}

type MovePath struct {
	typeFind     string
	angle        float64
	path         *[]*coordinate.Coordinate
	followTarget *target.Target
	currentPoint int
	needFindPath bool
	time         int64
}

func (unit *Unit) GetMovePathState() (string, float64, *target.Target, *[]*coordinate.Coordinate, int, bool, int64) {
	mp := unit.movePath
	if mp == nil {
		return "", 0, nil, nil, 0, false, 0
	}

	return mp.typeFind, mp.angle, mp.followTarget, mp.path, mp.currentPoint, mp.needFindPath, mp.time
}

func (unit *Unit) NextMovePoint() {
	unit.moveMx.Lock()
	defer unit.moveMx.Unlock()

	if unit.movePath == nil {
		return
	}

	unit.movePath.currentPoint++
}

func (unit *Unit) SetFindMovePath() {
	unit.moveMx.Lock()
	defer unit.moveMx.Unlock()

	if unit.movePath != nil {
		unit.movePath.needFindPath = true
	}
}

func (unit *Unit) RemoveMovePath() {
	unit.moveMx.Lock()
	defer unit.moveMx.Unlock()

	unit.movePath = nil
}

func (unit *Unit) SetMovePath(path *[]*coordinate.Coordinate) {
	unit.moveMx.Lock()
	defer unit.moveMx.Unlock()

	if unit.movePath == nil {
		return
	}

	unit.movePath.needFindPath = false
	unit.movePath.path = path
	unit.movePath.currentPoint = 0
	unit.movePath.time = time.Now().Unix()
}

func (unit *Unit) SetMovePathTarget(t *target.Target) {
	unit.moveMx.Lock()
	defer unit.moveMx.Unlock()

	unit.movePath = &MovePath{
		needFindPath: true,
		path:         &[]*coordinate.Coordinate{{X: t.X, Y: t.Y}},
		followTarget: t,
	}
}

func (unit *Unit) SetMovePathAngle(angle float64) {
	unit.moveMx.Lock()
	defer unit.moveMx.Unlock()

	unit.movePath = &MovePath{
		typeFind:     "angle",
		needFindPath: true,
		angle:        angle,
	}
}

type Damage struct {
	PlayerID   int   `json:"player_id"`
	TimeDamage int64 `json:"time_damage"`
	Damage     int   `json:"damage"`
}

func (unit *Unit) SetLastDamage(playerID, damage int) *Damage {
	if playerID == unit.OwnerID {
		return nil
	}

	newDamage := Damage{
		PlayerID:   playerID,
		TimeDamage: time.Now().Unix(),
		Damage:     damage,
	}
	unit.Damage = append(unit.Damage, newDamage)
	return &newDamage
}

func (unit *Unit) GetLastDamage() int {

	if len(unit.Damage) == 0 {
		return 0
	}

	last := unit.Damage[len(unit.Damage)-1]
	// если урон наносился больше чем 30 сек назад то он не учитывается
	if time.Now().Unix()-last.TimeDamage > 30 {
		unit.Damage = unit.Damage[:0]
		return 0
	}

	return last.PlayerID
}

func (unit *Unit) GetLastDamageByPlayerID(playerID, sec int) int {
	if len(unit.Damage) == 0 {
		return 0
	}

	allDamage := 0
	for _, d := range unit.Damage {
		if d.PlayerID == playerID {
			if time.Now().Unix()-d.TimeDamage < int64(sec) {
				allDamage += d.Damage
			}
		}
	}

	return allDamage
}

func (unit *Unit) GetAllDamage() map[int]int {
	r := make(map[int]int)

	if len(unit.Damage) == 0 {
		return r
	}

	for _, d := range unit.Damage {
		if time.Now().Unix()-d.TimeDamage < 30 {
			r[d.PlayerID] += d.Damage
		}
	}

	return r
}

func (unit *Unit) GetPhysicalModel() *physical_model.PhysicalModel {

	if unit.physicalModel == nil {
		unit.initPhysicalModel() // инициируем по умолчания, и ток в методе UpdatePhysicalModel уже докидываем скилы и тд
	}

	return unit.physicalModel
}

func (unit *Unit) UpdatePhysicalModel() {

	if unit.physicalModel == nil {
		unit.initPhysicalModel() // инициируем по умолчания, и ток в методе UpdatePhysicalModel уже докидываем скилы и тд
	}

	unit.physicalModel.ID = unit.ID
	unit.physicalModel.Speed = unit.GetMoveMaxPower() / _const.ServerTickSecPart
	unit.physicalModel.ReverseSpeed = unit.GetMaxReverse() / _const.ServerTickSecPart
	unit.physicalModel.PowerFactor = unit.GetPowerFactor() / _const.ServerTickSecPart
	unit.physicalModel.ReverseFactor = unit.GetReverseFactor() / _const.ServerTickSecPart
	unit.physicalModel.TurnSpeed = unit.GetTurnSpeed() / _const.ServerTickSecPart
	unit.physicalModel.Weight = unit.GetWeight()
}

func (unit *Unit) initPhysicalModel() {
	unit.physicalModel = &physical_model.PhysicalModel{
		Speed:         unit.GetMoveMaxPower() / _const.ServerTickSecPart,
		ReverseSpeed:  unit.GetMaxReverse() / _const.ServerTickSecPart,
		PowerFactor:   unit.GetPowerFactor() / _const.ServerTickSecPart,
		ReverseFactor: unit.GetReverseFactor() / _const.ServerTickSecPart,
		TurnSpeed:     unit.GetTurnSpeed() / _const.ServerTickSecPart,
		WASD:          physical_model.WASD{},
		MoveDrag:      unit.GetBody().MoveDrag,
		AngularDrag:   unit.GetBody().AngularDrag,
		Weight:        float64(unit.GetBody().CapacitySize),
		Type:          "unit",
		ID:            unit.GetID(),
		ChassisType:   unit.body.ChassisType,
		MoveDestroyer: unit.body.Fraction == _const.FAUNA,
	}

	// применяем настройки размера к обьектам колизий
	sizeOffset := float64(unit.GetBody().Scale) / 100
	unit.physicalModel.Height = float64(unit.GetBody().Height) * sizeOffset
	unit.physicalModel.Length = float64(unit.GetBody().Length) * sizeOffset
	unit.physicalModel.Width = float64(unit.GetBody().Width) * sizeOffset
	unit.physicalModel.Radius = int(float64(unit.GetBody().Radius) * sizeOffset)
}

func (unit *Unit) GetCopyPhysicalModel() *physical_model.PhysicalModel {
	origPH := unit.GetPhysicalModel()
	pm := *origPH
	return &pm
}

func (unit *Unit) LockDamage() {
	unit.damageMX.Lock()
}

func (unit *Unit) UnlockDamage() {
	unit.damageMX.Unlock()
}

func (unit *Unit) SetAllGunRotate(addRotate float64) {
	for _, weaponSlot := range unit.getBody().Weapons {
		if weaponSlot != nil {
			weaponSlot.SetGunRotate(weaponSlot.GetGunRotate() + addRotate)
		}
	}
}

func (unit *Unit) GetTransportUnit() *Unit {
	return nil
}

func (unit *Unit) GetNeedZ() float64 {
	return 0
}

func (unit *Unit) GetZ() float64 {
	return 0
}

func (unit *Unit) SetZ(float64) {
}

func (unit *Unit) IsFly() bool {
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

func (unit *Unit) GetJSON(mapTime int64) []byte {

	if unit.CreateJsonTime == mapTime && len(unit.CacheJson) > 0 {
		return unit.CacheJson
	}

	if unit.CacheJson == nil {
		unit.CacheJson = []byte{}
	}

	unit.CacheJson = unit.CacheJson[:0]
	unit.CacheJson = append(unit.CacheJson, game_math.GetIntBytes(unit.GetID())...)
	unit.CacheJson = append(unit.CacheJson, game_math.GetIntBytes(unit.OwnerID)...)
	unit.CacheJson = append(unit.CacheJson, game_math.GetIntBytes(unit.HP)...)
	unit.CacheJson = append(unit.CacheJson, byte(unit.GetBody().ID))

	// position data
	unit.CacheJson = append(unit.CacheJson, game_math.GetIntBytes(unit.GetX())...)
	unit.CacheJson = append(unit.CacheJson, game_math.GetIntBytes(unit.GetY())...)
	unit.CacheJson = append(unit.CacheJson, game_math.GetIntBytes(int(unit.GetRotate()))...)

	unit.CacheJson = append(unit.CacheJson, game_math.GetIntBytes(unit.GetBody().MaxHP)...)
	unit.CacheJson = append(unit.CacheJson, game_math.GetIntBytes(unit.GetRangeView())...)
	unit.CacheJson = append(unit.CacheJson, game_math.GetIntBytes(unit.GetRadarRange())...)

	unit.CacheJson = append(unit.CacheJson, game_math.BoolToByte(unit.fractionWarrior))
	unit.CacheJson = append(unit.CacheJson, _const.FractionByte[unit.OwnerFraction])

	unit.CacheJson = append(unit.CacheJson, byte(len([]byte(unit.GetBody().Texture))))
	unit.CacheJson = append(unit.CacheJson, []byte(unit.GetBody().Texture)...)

	unit.CacheJson = append(unit.CacheJson, byte(len([]byte(unit.Owner))))
	unit.CacheJson = append(unit.CacheJson, []byte(unit.Owner)...)

	// weapon data
	weaponData := []byte{}
	for _, unitWeaponSlot := range unit.RangeWeaponSlots() {

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

	unit.CacheJson = append(unit.CacheJson, game_math.GetIntBytes(len(weaponData))...)
	unit.CacheJson = append(unit.CacheJson, weaponData...)

	// equip data
	equipData := []byte{}
	for _, equipSlot := range unit.GetBody().GetAllEquips() {
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

	unit.CacheJson = append(unit.CacheJson, game_math.GetIntBytes(len(equipData))...)
	unit.CacheJson = append(unit.CacheJson, equipData...)

	return unit.CacheJson
}

func (unit *Unit) GetUpdateData(mapTime int64) []byte {

	command := []byte{}
	command = append(command, game_math.GetIntBytes(unit.HP)...)
	command = append(command, game_math.GetIntBytes(unit.GetRangeView())...)
	command = append(command, game_math.GetIntBytes(unit.GetRadarRange())...)

	return command
}

func (unit *Unit) GetShortInfo() *ShortUnitInfo {
	if unit == nil || unit.getBody() == nil {
		return nil
	}

	var hostile ShortUnitInfo

	hostile.X = unit.GetPhysicalModel().GetX()
	hostile.Y = unit.GetPhysicalModel().GetY()

	hostile.GunRotate = unit.GetGunner().GetGunRotate(1)
	hostile.Rotate = unit.GetRotate()
	hostile.MapID = unit.GetMapID()
	hostile.Evacuation = unit.GetEvacuation()
	hostile.ForceEvacuation = unit.GetForceEvacuation()
	hostile.InSky = unit.InSky()
	hostile.OwnerFraction = unit.OwnerFraction

	hostile.Body = detail.Body{
		Name:       unit.getBody().Name,
		MaxHP:      unit.getBody().MaxHP,
		RangeView:  unit.getBody().RangeView,
		RangeRadar: unit.getBody().RangeRadar,
		Scale:      unit.getBody().Scale,
		Length:     int(unit.GetPhysicalModel().GetLength()),
		Width:      int(unit.GetPhysicalModel().GetWidth()),
	}

	hostile.WeaponSlots = make(map[int]detail.BodyWeaponSlot)
	for number, wSlot := range unit.RangeWeaponSlots() {
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

	hostile.OwnerID = unit.GetOwnerID()
	hostile.Owner = unit.Owner
	hostile.ID = unit.GetID()
	hostile.BodyTexture = unit.BodyTexture

	hostile.HP = unit.GetHP()

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

	copyEquips(unit.getBody().EquippingI)
	copyEquips(unit.getBody().EquippingII)
	copyEquips(unit.getBody().EquippingIII)
	copyEquips(unit.getBody().EquippingIV)
	copyEquips(unit.getBody().EquippingV)

	return &hostile
}

func (unit *Unit) DelEquip() {

}

func (unit *Unit) DelAmmo() {

}

func (unit *Unit) SetEquip() {

}

func (unit *Unit) SetAmmo() {

}

func (unit *Unit) GetOwnerUser() string {
	return unit.Owner
}

func (unit *Unit) GetAmmoCount() int { // по диз доку оружие в юните может быть только одно
	for _, weaponSlot := range unit.getBody().Weapons {
		if weaponSlot.Weapon != nil {
			return weaponSlot.GetAmmoQuantity()
		}
	}

	return 0
}

func (unit *Unit) SetAnchorsEquip() {
	// метод расчитывает правильные положения якорец и точек прикрепления для спрайтов снаряжения на корпусе
	if unit == nil {
		return
	}

	// распологаем оружие
	unit.GetWeaponSlot(1).SetAnchor()

	// распологаем снарягу (не вся снаряга может быть на корпусе но это разруливается не тут)
	for _, slot := range unit.getBody().GetAllEquips() {
		slot.SetAnchor()
	}
}

func (unit *Unit) GetWeaponSlot(slotNumber int) *detail.BodyWeaponSlot { // по диз доку оружие в юните может быть только одно

	if unit == nil || unit.getBody() == nil {
		return nil
	}

	// TODO пока только 1 оружие это +- применимо, а так нет
	for _, weaponSlot := range unit.getBody().Weapons {
		return weaponSlot
	}

	return nil
}

func (unit *Unit) GetEquipPosInMap(typeEquip, numberSlot int) (int, int) {

	equipSlot := unit.GetBodyEquipSlot(typeEquip, numberSlot)
	if equipSlot == nil || equipSlot.Equip == nil {
		return 0, 0
	}

	return game_math.GetWeaponPosInMap(
		unit.GetX(), unit.GetY(), unit.getBody().Scale,
		float64(equipSlot.XAttach),
		float64(equipSlot.YAttach),
		unit.GetRotate())
}

func (unit *Unit) GetEquipFirePos(typeEquip, numberSlot int) []*game_math.Positions {

	equipSlot := unit.GetBodyEquipSlot(typeEquip, numberSlot)
	if equipSlot == nil || equipSlot.Equip == nil {
		return nil
	}

	return game_math.GetWeaponFirePositions(
		unit.GetX(), unit.GetY(), unit.GetBodyScale(), unit.GetRotate(), equipSlot.Rotate,
		equipSlot.Equip.XAttach, equipSlot.Equip.YAttach,
		equipSlot.Equip.FirePositions,
		float64(equipSlot.XAttach),
		float64(equipSlot.YAttach),
	)
}

func (unit *Unit) CheckViewCoordinate(x, y, radius int) (bool, bool) {

	if unit == nil || unit.getBody() == nil {
		return false, false
	}

	if unit.GetRangeView()+radius >= int(game_math.GetBetweenDist(unit.GetX(), unit.GetY(), x, y)) {
		return true, true
	}

	if unit.GetRadarRange()+radius >= int(game_math.GetBetweenDist(unit.GetX(), unit.GetY(), x, y)) {
		return false, true
	}

	return false, false
}

func (unit *Unit) RangeWeaponSlots() map[int]*detail.BodyWeaponSlot {
	// мы никогда не пишут в карту слотов оружия поэтому этот метод безопасен (по крайне мере пока)
	return unit.getBody().Weapons
}

func (unit *Unit) GetScale() int {
	return unit.getBody().Scale
}

func (unit *Unit) GetBurstOfShots() *burst_of_shots.BurstOfShots {
	if unit.BurstOfShots == nil {
		unit.BurstOfShots = &burst_of_shots.BurstOfShots{}
	}

	return unit.BurstOfShots
}

func (unit *Unit) SetVisibleObjectStore(v *visible_objects.VisibleObjectsStore) {
	unit.visibleObjects = v
}

func (unit *Unit) GetVisibleObjectStore() *visible_objects.VisibleObjectsStore {
	return unit.visibleObjects
}

func (unit *Unit) visibleObjectStore() *visible_objects.VisibleObjectsStore {
	if unit.visibleObjects == nil {
		unit.visibleObjects = &visible_objects.VisibleObjectsStore{}
	}

	return unit.visibleObjects
}

func (unit *Unit) GetVisibleObjects() <-chan *visible_objects.VisibleObject {
	return unit.visibleObjectStore().GetVisibleObjects()
}

func (unit *Unit) UnsafeRangeVisibleObjects() ([]*visible_objects.VisibleObject, *sync.RWMutex) {
	return unit.visibleObjectStore().UnsafeRangeMapDynamicObjects()
}

func (unit *Unit) GetVisibleObjectByTypeAndID(typeObj string, id int) *visible_objects.VisibleObject {
	return unit.visibleObjectStore().GetVisibleObjectByTypeAndID(typeObj, id)
}

func (unit *Unit) SetFractionWarrior(ok bool) {
	unit.fractionWarrior = ok
}

func (unit *Unit) FractionWarrior() bool {
	return unit.fractionWarrior
}
