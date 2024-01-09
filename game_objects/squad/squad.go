package squad

import (
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/coordinate"
	"github.com/TrashPony/veliri-lib/game_objects/dynamic_map_object"
	"github.com/TrashPony/veliri-lib/game_objects/skin"
	"github.com/TrashPony/veliri-lib/game_objects/special_hostiles"
	"github.com/TrashPony/veliri-lib/game_objects/target"
	"github.com/TrashPony/veliri-lib/game_objects/unit"
	"github.com/TrashPony/veliri-lib/game_objects/visible_objects"
	"github.com/TrashPony/veliri-lib/generate_ids"
	"math"
	"sync"
)

type Squad struct {
	ID                 int                    `json:"id"` // TODO int64
	Name               string                 `json:"name"`
	OwnerGame          bool                   `json:"-"` // если true то значит отряд привязан к игре UUIDGame
	UUIDGame           string                 `json:"-"` // если игрок не находится в UUIDGame то отряд перестает существовать
	ReturnSquadID      int                    `json:"-"` // после того как отряд перестал сущесовать возвращается ReturnSquadID
	Active             bool                   `json:"active"`
	BaseID             int                    `json:"base_id"`              /* если отряд не у игрока то он храниться на этой базе */
	LastGlobalPosition *coordinate.Coordinate `json:"last_global_position"` // если отряд был в данже, но по какой то причине сервак вырубило и данж пропал, эта та точка куда упадет отряд при заходе в игру
	softTransition     int

	matherShip *unit.Unit

	// запомненные динамические обьекты на карте [map_id][id]
	memoryDynamicObjects *visible_objects.VisibleObjectsStore

	// враги бота - например если это свой который открыл огонь по нему, то он будет защищатся
	// [typeHostile+id]points
	visibleObjects  *visible_objects.VisibleObjectsStore
	specialHostiles *special_hostiles.SpecialHostiles
	painThreshold   float64

	equipPanel             map[int]*EquipSell
	equipPanelRerender     bool
	countUpdateViewObjects int

	BodySkin    *skin.Skin         `json:"body_skin"`
	weaponSkins map[int]*skin.Skin `json:"weapon_skins"`
	Transfer    bool

	updateDB sync.Mutex
	tmx      sync.Mutex
	mx       sync.RWMutex
}

func (s *Squad) GetOwnerPlayerID() int {
	return s.GetOwnerID()
}

func (s *Squad) GetPainThreshold() float64 {
	if s.painThreshold == 0 {
		return 1
	}

	return s.painThreshold
}

func (s *Squad) SetPainThreshold(pain float64) {
	s.painThreshold = pain
}

func (s *Squad) AddHostile(typeHostile string, id, hatePoint int) {
	if s.specialHostiles == nil {
		s.specialHostiles = &special_hostiles.SpecialHostiles{}
	}

	s.specialHostiles.AddPoints(typeHostile, id, hatePoint)
}

func (s *Squad) RangeHostiles() <-chan *special_hostiles.SpecialHostile {
	if s.specialHostiles == nil {
		s.specialHostiles = &special_hostiles.SpecialHostiles{}
	}

	return s.specialHostiles.RangeHostiles()
}

func (s *Squad) AddIgnoreHostile(typeHostile string, id int) {
	if s.specialHostiles == nil {
		s.specialHostiles = &special_hostiles.SpecialHostiles{}
	}

	s.specialHostiles.AddIgnore(typeHostile, id)
}

func (s *Squad) SetHostilePoints(typeHostile string, id, hatePoint int) {
	if s.specialHostiles == nil {
		s.specialHostiles = &special_hostiles.SpecialHostiles{}
	}

	s.specialHostiles.SetPoints(typeHostile, id, hatePoint)
	// TODO NODES
}

func (s *Squad) ClearHostile() {

	if s == nil {
		return
	}

	s.specialHostiles = &special_hostiles.SpecialHostiles{}
}

func (s *Squad) GetSpecialHostileData() ([]byte, []byte) {
	if s.specialHostiles == nil {
		s.specialHostiles = &special_hostiles.SpecialHostiles{}
	}

	return s.specialHostiles.GetJsonData()
}

func (s *Squad) LoadSpecialHostileFromJson(ignoreHate []byte, specialHostiles []byte) {
	if s.specialHostiles == nil {
		s.specialHostiles = &special_hostiles.SpecialHostiles{}
	}

	s.specialHostiles.LoadFromJson(ignoreHate, specialHostiles)
}

func (s *Squad) CheckHostile(typeHostile string, id int, ownerID int) (bool, int) {

	if s == nil {
		return false, 0
	}

	if s.specialHostiles == nil {
		return false, 0
	}

	if typeHostile == "unit" && s.GetMS().GetID() == id {
		return false, 0
	}

	return s.specialHostiles.CheckHostile(typeHostile, id)
}

func (s *Squad) SetMS(ms *unit.Unit) {

	if s.matherShip != nil {
		s.matherShip.SetVisibleObjectStore(nil)
	}

	s.matherShip = ms
	s.checkVisibleObjectStore()
	s.matherShip.SetVisibleObjectStore(s.visibleObjects)
}
func (s *Squad) GetMS() *unit.Unit {
	if s == nil {
		return nil
	}
	return s.matherShip
}

func (s *Squad) IgnoreSmoke() bool {
	if s.matherShip == nil {
		return false
	}

	return s.matherShip.GetEffectByUUID("smoke_blind") != nil
}

func (s *Squad) GetVisibleObjectByTypeAndID(typeObj string, id int) *visible_objects.VisibleObject {
	s.checkVisibleObjectStore()
	return s.visibleObjects.GetVisibleObjectByTypeAndID(typeObj, id)
}

func (s *Squad) InitVisibleObjects() {
	s.checkVisibleObjectStore()
	s.visibleObjects.InitVisibleObjects()
}

func (s *Squad) GetVisibleObjects() <-chan *visible_objects.VisibleObject {
	s.checkVisibleObjectStore()
	return s.visibleObjects.GetVisibleObjects()
}

func (s *Squad) AddVisibleObject(newObj *visible_objects.VisibleObject) {
	s.checkVisibleObjectStore()
	s.visibleObjects.AddVisibleObject(newObj)
}

func (s *Squad) RemoveVisibleObject(removeObj *visible_objects.VisibleObject) {
	s.checkVisibleObjectStore()
	s.visibleObjects.RemoveVisibleObject(removeObj)
}

func (s *Squad) UnsafeRangeVisibleObjects() ([]*visible_objects.VisibleObject, *sync.RWMutex) {
	s.checkMemoryObjectStore()
	return s.visibleObjects.UnsafeRangeMapDynamicObjects()
}

func (s *Squad) checkVisibleObjectStore() {
	if s.visibleObjects == nil {
		s.visibleObjects = &visible_objects.VisibleObjectsStore{}
	}
}

func (s *Squad) UpdateLock() {
	s.updateDB.Lock()
}

func (s *Squad) UpdateUnlock() {
	s.updateDB.Unlock()
}

func (s *Squad) ThoriumLock() {
	s.tmx.Lock()
}

func (s *Squad) ThoriumUnlock() {
	s.tmx.Unlock()
}

func (s *Squad) GetFormationCoordinate(x, y int) (int, int) {
	//хз почему +90 но работает)
	alpha := s.GetMS().GetRotate() + 90*math.Pi/180
	newX := float64(x)*game_math.Cos(alpha) - float64(y)*game_math.Sin(alpha) + float64(s.GetMS().GetX())
	newY := float64(x)*game_math.Sin(alpha) + float64(y)*game_math.Cos(alpha) + float64(s.GetMS().GetY())

	return int(newX), int(newY)
}

func (s *Squad) CheckViewCoordinate(x, y, radius int) (bool, bool) {

	if s == nil || s.GetMS() == nil {
		return false, false
	}

	view, radar := s.GetMS().CheckViewCoordinate(x, y, radius)
	if view {
		return true, true
	}

	return view, radar
}

// todo мега костыль что бы не обновлять обьекты каждые 32 мс а как укажешь
func (s *Squad) GetUpdateViewObjects() bool {
	if s.countUpdateViewObjects == 3 {
		s.countUpdateViewObjects = 0
		return true
	} else {
		s.countUpdateViewObjects++
		return false
	}
}

func (s *Squad) InitMemoryObjects() {
	s.memoryDynamicObjects.InitVisibleObjects()
}

func (s *Squad) AddDynamicObject(object *dynamic_map_object.Object, mapID int, view, radar bool, mapTime int64) {
	s.checkMemoryObjectStore()
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
		OwnerID:       object.GetOwnerID(),
		OwnerPlayerID: object.OwnerPlayerID,
		Object:        object,
		ObjectJSON:    object.GetJSON(mapTime),
		Work:          object.GetWork(),
	}

	vObj.SetView(view)
	vObj.SetRadar(radar)

	s.memoryDynamicObjects.AddDynamicObject(vObj)
}

func (s *Squad) AddDynamicMemoryObject(object *visible_objects.VisibleObject) {
	s.checkMemoryObjectStore()
	s.memoryDynamicObjects.AddDynamicMemoryObject(object)
}

func (s *Squad) RemoveDynamicObject(id int) {
	s.checkMemoryObjectStore()
	s.memoryDynamicObjects.RemoveDynamicObject(id)
}

func (s *Squad) GetMapDynamicObjectByID(id int) *visible_objects.VisibleObject {
	s.checkMemoryObjectStore()
	return s.memoryDynamicObjects.GetMapDynamicObjectByID(id)
}

func (s *Squad) GetMapDynamicObjects(mapID int) <-chan *visible_objects.VisibleObject {
	s.checkMemoryObjectStore()
	return s.memoryDynamicObjects.GetMapDynamicObjects(mapID)
}

func (s *Squad) UnsafeRangeMapDynamicObjects() ([]*visible_objects.VisibleObject, *sync.RWMutex) {
	s.checkMemoryObjectStore()
	return s.memoryDynamicObjects.UnsafeRangeMapDynamicObjects()
}

func (s *Squad) checkMemoryObjectStore() {
	if s.memoryDynamicObjects == nil {
		s.memoryDynamicObjects = &visible_objects.VisibleObjectsStore{}
	}
}

func (s *Squad) GetX() int {
	if s.GetMS() != nil {
		return s.GetMS().GetX()
	}
	return 0
}

func (s *Squad) GetY() int {
	if s.GetMS() != nil {
		return s.GetMS().GetY()
	}
	return 0
}

func (s *Squad) GetOwnerID() int {
	if s.GetMS() != nil {
		return s.GetMS().GetOwnerID()
	}
	return 0
}

func (s *Squad) GetRangeView() int {
	if s.GetMS() != nil {
		return s.GetMS().GetRangeView()
	}
	return 0
}

func (s *Squad) GetID() int {
	return s.ID
}

func (s *Squad) GetType() string {
	if s.GetMS() != nil {
		return s.GetMS().GetType()
	}
	return ""
}

func (s *Squad) GetWeaponTarget() *target.Target {
	if s.GetMS() != nil {
		return s.GetMS().GetWeaponTarget()
	}

	return nil
}

func (s *Squad) SetWeaponTarget(target *target.Target) {
	if s.GetMS() != nil {
		s.GetMS().SetWeaponTarget(target)
	}
}

func (s *Squad) GetSoftTransition() int {
	return s.softTransition
}

func (s *Squad) SetSoftTransition(transition int) {
	s.softTransition = transition
}

func (s *Squad) AddSoftTransition(transition int) {
	s.softTransition += transition
}

func (s *Squad) GetName() string {
	return s.Name
}

func (s *Squad) SetFractionWarrior(ok bool) {
	s.GetMS().SetFractionWarrior(ok)
}

func (s *Squad) FractionWarrior() bool {
	return s.GetMS().FractionWarrior()
}

func (s *Squad) GetFraction() string {
	return s.GetMS().OwnerFraction
}

func (s *Squad) SetWeaponSkin(slotNumber int, sk *skin.Skin) {
	s.mx.Lock()
	defer s.mx.Unlock()

	if s.weaponSkins == nil {
		s.weaponSkins = make(map[int]*skin.Skin)
	}

	s.weaponSkins[slotNumber] = sk
}

func (s *Squad) GetWeaponSkin(slotNumber int) *skin.Skin {
	s.mx.Lock()
	defer s.mx.Unlock()

	return s.weaponSkins[slotNumber]
}

func (s *Squad) RangeWeaponSkins() map[int]*skin.Skin {
	s.mx.Lock()
	defer s.mx.Unlock()

	weaponSkins := make(map[int]*skin.Skin)

	for k, v := range s.weaponSkins {
		weaponSkins[k] = v
	}

	return weaponSkins
}

func (s *Squad) GetCorporationID() int {
	if s.GetMS() != nil {
		return s.GetMS().GetCorporationID()
	}
	return 0
}
