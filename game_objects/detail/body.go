package detail

import (
	"encoding/json"
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/anchor"
	"github.com/TrashPony/veliri-lib/game_objects/coordinate"
	"github.com/TrashPony/veliri-lib/game_objects/effect"
	"math/rand"
	"sync"
)

type Body struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Texture       string `json:"texture"`
	Specification string `json:"specification"`
	MotherShip    bool   `json:"mother_ship"`

	Speed         float64 `json:"speed"`          // -- макс скорость вперед
	ReverseSpeed  float64 `json:"reverse_speed"`  //-- макс скорость назад
	PowerFactor   float64 `json:"power_factor"`   // -- сила ускорения вперед
	ReverseFactor float64 `json:"reverse_factor"` //-- сила ускорения назад
	TurnSpeed     float64 `json:"turn_speed"`     //-- скорость поворота в радианах

	MaxHP                 int `json:"max_hp"`
	ProtectionToKinetics  int `json:"protection_to_kinetics"`
	ProtectionToThermo    int `json:"protection_to_thermo"`
	ProtectionToExplosion int `json:"protection_to_explosion"`

	RangeView  int `json:"range_view"`
	RangeRadar int `json:"range_radar"`

	CapacitySize int `json:"capacity_size"` /* вместимость корпуса к кубо-метрах */
	StandardSize int `json:"standard_size"` /* small - 1, medium - 2, big - 3, размер корпуса (если корпус мс то неучитывается)*/

	// TODO передалать это на 1 мапу
	EquippingI   map[int]*BodyEquipSlot `json:"-"`
	EquippingII  map[int]*BodyEquipSlot `json:"-"`
	EquippingIII map[int]*BodyEquipSlot `json:"-"`
	EquippingIV  map[int]*BodyEquipSlot `json:"-"`
	EquippingV   map[int]*BodyEquipSlot `json:"-"`

	ThoriumSlots map[int]*ThoriumSlot `json:"-"` /* слоты в которых хранится топливо */

	Weapons map[int]*BodyWeaponSlot `json:"-"`

	Height int `json:"height"`
	Length int `json:"length"`
	Width  int `json:"width"`
	Radius int `json:"radius"`
	Scale  int `json:"scale"`

	Fraction     string      `json:"fraction"`
	Requirements map[int]int `json:"requirements"` // [id]lvl

	Bonuses          map[string]*Bonus                `json:"bonuses"`
	ChassisType      string                           `json:"chassis_type"`
	WheelsPosNoScale map[string]coordinate.Coordinate `json:"wheels_pos_no_scale"`
	WheelAnchors     map[string]anchor.Anchor         `json:"wheel_anchors"`

	MoveDrag    float64 `json:"move_drag"`
	AngularDrag float64 `json:"angular_drag"`

	Weight    float64 `json:"weight"`
	MaxEnergy int     `json:"max_energy"` // для установки снаряжения

	MaxPower                  int `json:"max_power"` // аккамулятор
	RecoveryPower             int `json:"recovery_power"`
	RecoveryPowerCycle        int `json:"recovery_power_cycle"`
	CurrentRecoveryPowerCycle int `json:"-"`

	Attributes          map[string]int                 `json:"attributes"`
	AdditionalInventory map[string]AdditionalInventory `json:"additional_inventory"`

	DamageZones []DamageZone `json:"damage_zones"`
	mx          sync.RWMutex
}

type AdditionalInventory struct {
	CapacitySize int              `json:"capacity_size"` /* вместимость корпуса к кубо-метрах */
	Filter       map[string][]int `json:"filter"`
	Key          string           `json:"key"`
}

type DamageZone struct {
	StartAngle int     `json:"start_angle"`
	EndAngle   int     `json:"end_angle"`
	DamageK    float64 `json:"damage_k"`
}

func (body *Body) GetDamageZoneByAngle(angle int) *DamageZone {
	for _, z := range body.DamageZones {
		if z.StartAngle <= angle && z.EndAngle >= angle {
			return &z
		}
	}

	return nil
}

func (body *Body) GetName() string {
	return body.Name
}

func (body *Body) GetItemType() string {
	return ""
}

func (body *Body) GetItemName() string {
	return ""
}

func (body *Body) GetInMap() bool {
	return false
}

func (body *Body) GetIcon() string {
	return ""
}

func (body *Body) GetStandardSize() int {
	return body.StandardSize
}

func (body *Body) GetTypeSlot() int {
	return 0
}

type Bonus struct {
	ID                  int            `json:"'id'"`
	Bonus               *effect.Effect `json:"bonus"`
	OverrideDescription string         `json:"od"`
	Requirements        map[int]int    `json:"requirements"` // [skill_name] level
}

func (body *Body) GetAllEquips() []*BodyEquipSlot {
	equips := make([]*BodyEquipSlot, 0)

	var addEquips = func(equip map[int]*BodyEquipSlot, typeSlot int) {
		for slot, s := range equip {
			if s.Equip != nil {
				s.TypeSlot = typeSlot
				s.Slot = slot
				equips = append(equips, s)
			}
		}
	}

	addEquips(body.EquippingI, 1) // TODO передалать это на 1 мапу
	addEquips(body.EquippingII, 2)
	addEquips(body.EquippingIII, 3)
	addEquips(body.EquippingIV, 4)
	addEquips(body.EquippingV, 5)

	return equips
}

func (body *Body) GetAllEquipSlots() []*BodyEquipSlot {
	equips := make([]*BodyEquipSlot, 0)
	var addEquips = func(equip map[int]*BodyEquipSlot, typeSlot int) {
		for slot, s := range equip {
			s.TypeSlot = typeSlot
			s.Slot = slot
			equips = append(equips, s)
		}
	}

	addEquips(body.EquippingI, 1) // TODO передалать это на 1 мапу
	addEquips(body.EquippingII, 2)
	addEquips(body.EquippingIII, 3)
	addEquips(body.EquippingIV, 4)
	addEquips(body.EquippingV, 5)

	return equips
}

func (body *Body) GetApplicableEquips(applicable string) []*BodyEquipSlot {

	equips := make([]*BodyEquipSlot, 0)

	var findEquip = func(equip map[int]*BodyEquipSlot, typeSlot int) {
		for slot, s := range equip {
			if s.Equip != nil && s.Equip.Applicable == applicable {
				s.TypeSlot = typeSlot
				s.Slot = slot
				equips = append(equips, s)
			}
		}
	}

	findEquip(body.EquippingI, 1) // TODO передалать это на 1 мапу
	findEquip(body.EquippingII, 2)
	findEquip(body.EquippingIII, 3)
	findEquip(body.EquippingIV, 4)
	findEquip(body.EquippingV, 5)

	return equips
}

func (body *Body) GetEquipSlot(typeSlot, numberSlot int) *BodyEquipSlot {
	if typeSlot == 1 { // TODO передалать это на 1 мапу
		return body.EquippingI[numberSlot]
	}
	if typeSlot == 2 {
		return body.EquippingII[numberSlot]
	}
	if typeSlot == 3 {
		return body.EquippingIII[numberSlot]
	}
	if typeSlot == 4 {
		return body.EquippingIV[numberSlot]
	}
	if typeSlot == 5 {
		return body.EquippingV[numberSlot]
	}

	return nil
}

func (body *Body) GetRandomEquip() *BodyEquipSlot {

	typeSlot := rand.Intn(5) + 1

	randEquip := func(equips map[int]*BodyEquipSlot) *BodyEquipSlot {
		count := 0

		if equips == nil || len(equips) == 0 {
			return body.GetRandomEquip()
		}

		count2 := rand.Intn(len(equips))
		for _, slot := range equips {
			if count == count2 {
				return slot
			}
			count++
		}
		return body.GetRandomEquip()
	}

	if typeSlot == 1 { // TODO передалать это на 1 мапу
		return randEquip(body.EquippingI)
	}
	if typeSlot == 2 {
		return randEquip(body.EquippingII)
	}
	if typeSlot == 3 {
		return randEquip(body.EquippingIII)
	}
	if typeSlot == 4 {
		return randEquip(body.EquippingIV)
	}
	if typeSlot == 5 {
		return randEquip(body.EquippingV)
	}

	return body.GetRandomEquip()
}

func (body *Body) GetUseEnergy() int {
	var allPower int

	var counter = func(equip map[int]*BodyEquipSlot) int {
		var power int
		for _, slot := range equip {
			if slot.Equip != nil {
				power = power + slot.Equip.Power
			}
		}
		return power
	}

	allPower = allPower + counter(body.EquippingI) // TODO передалать это на 1 мапу
	allPower = allPower + counter(body.EquippingII)
	allPower = allPower + counter(body.EquippingIII)
	allPower = allPower + counter(body.EquippingIV)
	allPower = allPower + counter(body.EquippingV)

	for _, slot := range body.Weapons {
		if slot.Weapon != nil {
			allPower = allPower + slot.Weapon.Power
		}
	}

	return allPower
}

func (body *Body) GetUseCapacitySize() int {
	var allSize int

	var counter = func(equip map[int]*BodyEquipSlot) int {
		var size int
		for _, slot := range equip {
			if slot.Equip != nil {
				size = size + slot.Equip.Size
			}
		}
		return size
	}

	allSize = allSize + counter(body.EquippingI) // TODO передалать это на 1 мапу
	allSize = allSize + counter(body.EquippingII)
	allSize = allSize + counter(body.EquippingIII)
	allSize = allSize + counter(body.EquippingIV)
	allSize = allSize + counter(body.EquippingV)

	for _, slot := range body.Weapons {
		if slot.Weapon != nil {
			allSize = allSize + slot.Weapon.Size

			if slot.GetAmmo() != nil {
				allSize = allSize + slot.GetAmmo().Size*slot.GetAmmoQuantity()
			}
		}
	}

	return allSize
}

func (body *Body) GetThoriumSlot(number int) *ThoriumSlot {

	body.mx.RLock()
	defer body.mx.RUnlock()

	return body.ThoriumSlots[number]
}

func (body *Body) GetJSONBodyEquipSlots() []string {

	equips := make([]string, 0)

	for _, slot := range body.GetAllEquipSlots() {
		equips = append(equips, slot.GetJSON())
	}

	return equips
}

func (body *Body) GetJSONWeaponSlots() []string {

	weapons := make([]string, 0)

	for _, weaponSlot := range body.Weapons {
		weapons = append(weapons, weaponSlot.GetJSON())
	}

	return weapons
}

func (body *Body) GetJSONThoriumSlots() []string {

	thorium := make([]string, 0)

	for _, thoriumSlot := range body.ThoriumSlots {
		thorium = append(thorium, thoriumSlot.GetJSON())
	}

	return thorium
}

func (body *Body) GetJSON() string {
	body.mx.Lock()
	defer body.mx.Unlock()

	jsonBody, err := json.Marshal(body)
	if err != nil {
		println("body to json: ", err.Error())
	}

	return string(jsonBody)
}

func (body *Body) SetWheelsPositions() {

	if body.WheelAnchors == nil {
		body.WheelAnchors = make(map[string]anchor.Anchor)
	}

	for key, pos := range body.WheelsPosNoScale {
		xAnchor, yAnchor, realXAttach, realYAttach := game_math.GetAnchorWeapon(64, 64, int(pos.X), int(pos.Y))
		newAnc := anchor.Anchor{
			XAnchor:     xAnchor,
			YAnchor:     yAnchor,
			RealXAttach: realXAttach,
			RealYAttach: realYAttach,
		}

		ancState, ok := body.WheelAnchors[key]
		if ok {
			newAnc.Type = ancState.Type
			newAnc.Scale = ancState.Scale
			newAnc.Rotate = ancState.Rotate
			newAnc.Height = ancState.Height
		}

		body.WheelAnchors[key] = newAnc
	}
}
