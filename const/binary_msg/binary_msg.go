package binary_msg

import (
	"fmt"
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"github.com/TrashPony/veliri-lib/game_objects/inventory"
	"github.com/TrashPony/veliri-lib/game_objects/rope"
	"github.com/TrashPony/veliri-lib/game_objects/spawn"
	"github.com/TrashPony/veliri-lib/game_objects/target"
	"github.com/TrashPony/veliri-lib/game_objects/violator"
	"github.com/TrashPony/veliri-lib/game_objects/visible_anomaly"
	"math"
	"sort"
)

// [eventID, data]

func CreateBinaryUnitMoveMsg(unitID, speed, x, y, z, ms, rotate, angularVelocity int, animate, direction, sky, a, d, w, sp bool) []byte {
	// [1[eventID], 4[unitID], 4[speed], 4[x], 4[y], 4[z], 4[ms], 4[rotate], 4[angularVelocity], 4[mpID] 1[animate], 1[direction], 1[sky]]

	command := make([]byte, 37)

	command[0] = 1
	game_math.ReuseByteSlice(&command, 1, game_math.GetIntBytes(unitID))
	game_math.ReuseByteSlice(&command, 5, game_math.GetIntBytes(speed))
	game_math.ReuseByteSlice(&command, 9, game_math.GetIntBytes(x))
	game_math.ReuseByteSlice(&command, 13, game_math.GetIntBytes(y))
	game_math.ReuseByteSlice(&command, 17, game_math.GetIntBytes(z))
	command[21] = byte(ms)
	game_math.ReuseByteSlice(&command, 22, game_math.GetIntBytes(rotate))
	game_math.ReuseByteSlice(&command, 26, game_math.GetIntBytes(angularVelocity))
	command[30] = game_math.BoolToByte(animate)

	command[31] = game_math.BoolToByte(direction)
	command[32] = game_math.BoolToByte(sky)
	command[33] = game_math.BoolToByte(a)
	command[34] = game_math.BoolToByte(d)
	command[35] = game_math.BoolToByte(w)
	command[36] = game_math.BoolToByte(sp)

	return command
}

func CreateBinaryTransportMoveMsg(id, x, y, z, ms, rotate, mpID int) []byte {
	// [1[eventID], 4[ID], 4[x], 4[y], 4[z], 4[ms], 4[rotate], 4[mpID]]
	command := []byte{2}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(z)...)
	command = append(command, game_math.GetIntBytes(ms)...)
	command = append(command, game_math.GetIntBytes(rotate)...)
	command = append(command, game_math.GetIntBytes(mpID)...)

	return command
}

var typeObjects = map[string]byte{
	"box":      1,
	"object":   2,
	"map_item": 3,
}

func CreateBinaryObjectMove(id, x, y, ms, rotate int, typeObj string) []byte {
	// [1[eventID], 1[TypeObj] 4[ID], 4[x], 4[y], 4[ms], 4[rotate], 4[mpID]]
	command := []byte{3, typeObjects[typeObj]}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(ms)...)
	command = append(command, game_math.GetIntBytes(rotate)...)

	return command
}

func CreateBinaryDroneMove(id, x, y, z, ms, rotate int) []byte {
	// [1[eventID], 4[ID], 4[x], 4[y], 4[z], 4[ms], 4[rotate], 4[mpID]]
	command := []byte{4}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(z)...)
	command = append(command, byte(ms))
	command = append(command, game_math.GetIntBytes(rotate)...)

	return command
}

// TODO возможно надо передавать тип, хотя этим должен заниматся модуль view [radarMark.Type]
func CreateMarkBinaryMove(id, x, y, ms int) []byte {
	// [1[eventID] 4[ID], 4[x], 4[y], 4[ms]]

	command := []byte{5}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(ms)...)

	return command
}

func CreateBulletBinaryFly(typeID, id, x, y, z, ms, rotate int) []byte {
	// [1[eventID] 4[typeID], 4[id], 4[x], 4[y], 4[z], 4[ms], 4[rotate], 4[mpID]]

	command := []byte{6}

	command = append(command, byte(typeID))
	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(z)...)
	command = append(command, byte(ms))
	command = append(command, game_math.GetIntBytes(rotate)...)

	return command
}

func CreateBulletBinaryExplosion(typeID, x, y, z int) []byte {
	// [1[eventID] 4[typeID], 4[x], 4[y], 4[z], 4[mpID]]

	command := []byte{7}

	command = append(command, byte(typeID))
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(z)...)

	return command
}

func CreateRotateGunBinaryMsg(id, ms, rotate, slot int, sound bool) []byte {
	// [1[eventID], 4[ID], 4[ms], 4[rotate]
	command := []byte{8}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, byte(ms))
	command = append(command, game_math.GetIntBytes(rotate)...)
	command = append(command, byte(slot))
	command = append(command, game_math.BoolToByte(sound))

	return command
}

func CreateFireGunBinaryMsg(typeID, x, y, z, rotate, accumulationPercent int, force bool, unit_id, type_slot, slot int) []byte {
	// [1[eventID] 4[typeID], 4[x], 4[y], 4[z], 4[rotate], 4[mpID]]
	command := []byte{9}

	command = append(command, game_math.GetIntBytes(typeID)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(z)...)
	command = append(command, game_math.GetIntBytes(rotate)...)
	command = append(command, byte(accumulationPercent))
	command = append(command, game_math.BoolToByte(force))
	command = append(command, game_math.GetIntBytes(unit_id)...)
	command = append(command, byte((type_slot*10)+slot))

	return command
}

func RotateTurretGunBinaryMsg(id, rotate, ms int) []byte {
	// [1[eventID], 4[ID], 4[ms], 4[rotate]
	command := []byte{10}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, game_math.GetIntBytes(ms)...)
	command = append(command, game_math.GetIntBytes(rotate)...)

	return command
}

func MoveCloudBinaryMsg(id, idType, x, y, mpID, alpha, rotate int) []byte {
	// [1[eventID], 4[id], 4[typeID], 4[x], 4[y], 4[mpID], 4[alpha], 4[rotate]]
	command := []byte{11}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, game_math.GetIntBytes(idType)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(mpID)...)
	command = append(command, game_math.GetIntBytes(alpha)...)
	command = append(command, game_math.GetIntBytes(rotate)...)

	return command
}

func RemoveCloudBinaryMsg(id int) []byte {
	// [1[eventID], 4[id]
	command := []byte{12}
	command = append(command, game_math.GetIntBytes(id)...)
	return command
}

func WeaponMouseTargetBinary(weaponSlot, x, y, radius, ammoCount, ammoAvailable, accumulationPercent int, reload, chase, hide bool, targetType string, targetID, fireX, fireY int) []byte {

	command := make([]byte, 33)

	command[0] = 13
	game_math.ReuseByteSlice(&command, 1, game_math.GetIntBytes(x))
	game_math.ReuseByteSlice(&command, 5, game_math.GetIntBytes(y))
	game_math.ReuseByteSlice(&command, 9, game_math.GetIntBytes(radius))
	command[13] = byte(ammoCount)
	command[14] = byte(ammoAvailable)
	command[15] = byte(accumulationPercent)
	command[16] = game_math.BoolToByte(reload)
	command[17] = game_math.BoolToByte(chase)
	command[18] = byte(_const.MapBinItems[targetType])
	game_math.ReuseByteSlice(&command, 19, game_math.GetIntBytes(targetID))
	game_math.ReuseByteSlice(&command, 23, game_math.GetIntBytes(fireX))
	game_math.ReuseByteSlice(&command, 27, game_math.GetIntBytes(fireY))
	command[31] = byte(weaponSlot)
	command[32] = game_math.BoolToByte(hide)

	return command
}

func StatusSquadBinaryMsg(hp, shieldHP, energy int, autopilot bool, slots map[int]*detail.ThoriumSlot) []byte {
	// [1[eventID], 4[hp], 4[energy], 1[autopilot]], 4[count_slot_install], 21 * slot_count[slots data]
	command := []byte{14}

	command = append(command, game_math.GetIntBytes(hp)...)
	command = append(command, game_math.GetIntBytes(shieldHP)...)
	command = append(command, game_math.GetIntBytes(energy)...)
	command = append(command, game_math.BoolToByte(autopilot))

	command = append(command, game_math.GetIntBytes(len(slots))...)

	keys := make([]int, 0, len(slots))

	for k := range slots {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, slot := range keys {

		worked := slots[slot].Worked
		if worked == 0 && slots[slot].Reload {
			worked = 1 // что бы на фронте не отображалось "нет энергии" когда запрос в пути
		}

		command = append(command, game_math.GetIntBytes(slots[slot].Number)...)
		command = append(command, game_math.GetIntBytes(worked)...)
		command = append(command, game_math.GetIntBytes(slots[slot].CurrentFuel.EnergyCap)...)
		command = append(command, byte(0))                     // TODO
		command = append(command, game_math.GetIntBytes(0)...) // TODO
		command = append(command, game_math.BoolToByte(false))
	}

	return command
}

func StatusEquipPanelBinaryMsg(slotEquipSelect int, rerender bool, equipPanelBin []byte) []byte {
	// [1[eventID], 1[typeSelect], 4[typeEquipSelect], 4[slotEquipSelect], ???[equipPanelBin]]
	command := []byte{15}

	command = append(command, byte(slotEquipSelect))
	command = append(command, game_math.BoolToByte(rerender))
	command = append(command, equipPanelBin...) // из за сложности конкуретного доступа бин формируется в обьекте отряда
	return command
}

func RotateEquipBinaryMsg(id, rotate, ms, typeSlot, slot int) []byte {
	// [1[eventID], 4[id], 4[r], 4[ms], 4[ts], 4[s]]
	command := []byte{16}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, game_math.GetIntBytes(rotate)...)
	command = append(command, game_math.GetIntBytes(ms)...)
	command = append(command, game_math.GetIntBytes(typeSlot)...)
	command = append(command, game_math.GetIntBytes(slot)...)

	return command
}

func DamageTextBinaryMsg(x, y, damage int, typeObject, typeDealer string, dealerID, ownerID, damageK int, unitShield, destroyShield bool, area int, tt byte) []byte {
	// [1[eventID], 4[x], 4[y], 4[d], 4[m], 1[t]]
	command := []byte{17}

	_, ok := _const.MapBinItems[typeObject]
	if !ok {
		fmt.Println("unknown type object: ", typeObject)
	}

	if area > math.MaxInt8 {
		area = math.MaxInt8
	}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(damage)...)
	command = append(command, byte(_const.MapBinItems[typeObject]))
	command = append(command, byte(_const.MapBinItems[typeDealer]))
	command = append(command, game_math.GetIntBytes(dealerID)...)
	command = append(command, game_math.GetIntBytes(ownerID)...)
	command = append(command, byte(damageK))
	command = append(command, game_math.BoolToByte(unitShield))
	command = append(command, game_math.BoolToByte(destroyShield))
	command = append(command, byte(area))
	command = append(command, byte(tt))

	return command
}

func ObjectDeadBinaryMsg(id, x, y int, typeObject string) []byte {
	// [1[eventID], 4[id], 4[x], 4[y], 4[m], 1[t]]
	command := []byte{18}

	_, ok := _const.MapBinItems[typeObject]
	if !ok {
		fmt.Println("unknown type object: ", typeObject)
	}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, byte(_const.MapBinItems[typeObject]))

	return command
}

func CreateBulletLaserFly(typeID, x, y, toX, toY, unitID, AccumulationPercent, equipType, equipNumber, state int) []byte {
	command := []byte{20}

	command = append(command, byte(typeID))
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(toX)...)
	command = append(command, game_math.GetIntBytes(toY)...)
	command = append(command, game_math.GetIntBytes(unitID)...)
	command = append(command, byte(AccumulationPercent))
	command = append(command, byte(equipType))
	command = append(command, byte(equipNumber))
	command = append(command, byte(state))

	return command
}

var MarksTypes = map[string]byte{
	"fly":       1,
	"ground":    2,
	"structure": 3,
	"resource":  4,
	"bullet":    5,
}

func CreateBinaryCreateRadarMarkMsg(id, x, y int, typeMark string) []byte {
	command := []byte{39}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, MarksTypes[typeMark])

	return command
}

func CreateBinaryRemoveRadarMarkMsg(id int) []byte {
	command := []byte{40}

	command = append(command, game_math.GetIntBytes(id)...)

	return command
}

func CreateBinaryRemoveRadarObjectMsg(idObj int, typeObj string) []byte {
	command := []byte{41}

	command = append(command, game_math.GetIntBytes(idObj)...)
	command = append(command, byte(_const.MapBinItems[typeObj]))

	return command
}

func CreateBinaryCreateObjMsg(typeObj string, data []byte) []byte {
	command := make([]byte, 6+len(data))

	_, ok := _const.MapBinItems[typeObj]
	if !ok {
		println(typeObj)
	}

	command[0] = 42
	command[1] = byte(_const.MapBinItems[typeObj])
	game_math.ReuseByteSlice(&command, 2, game_math.GetIntBytes(len(data)))
	game_math.ReuseByteSlice(&command, 6, data)

	return command
}

func CreateBinaryUpdateObjMsg(typeObj string, id int, data []byte) []byte {

	command := make([]byte, 10+len(data))

	command[0] = 43
	command[1] = byte(_const.MapBinItems[typeObj])
	game_math.ReuseByteSlice(&command, 2, game_math.GetIntBytes(id))
	game_math.ReuseByteSlice(&command, 6, game_math.GetIntBytes(len(data)))
	game_math.ReuseByteSlice(&command, 10, data)

	return command
}

func CreateBinaryJammerMsg(x, y, r, typeJamer int) []byte {
	command := []byte{21}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(r)...)
	command = append(command, byte(typeJamer))

	return command
}

func ErrorUseEquip(slotNumber int, noPower, reload bool) []byte {
	command := []byte{38}

	command = append(command, game_math.GetIntBytes(slotNumber)...)
	command = append(command, game_math.BoolToByte(noPower))
	command = append(command, game_math.BoolToByte(reload))

	return command
}

func CreateBinaryMissileDefenseMsg(x, y, toX, toY int) []byte {
	command := []byte{23}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(toX)...)
	command = append(command, game_math.GetIntBytes(toY)...)

	return command
}

func CreateBinaryPlaneMineMsg(x, y, r int) []byte {
	command := []byte{28}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(r)...)

	return command
}

func CreateBinaryFlyEquip(x, y, m int) []byte {
	// TODO в случае нужды переделать на 117 событие
	command := []byte{44}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(m)...)

	return command
}

func CreateBinaryLaunchDrone(x, y, m int) []byte {
	// TODO в случае нужды переделать на 117 событие
	command := []byte{45}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(m)...)

	return command
}

func CreateBodyCollision(x, y, m int) []byte {
	// TODO в случае нужды переделать на 117 событие
	command := []byte{46}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(m)...)

	return command
}

func CreateBinaryJumpMsg(x, y, a int) []byte {
	command := []byte{31}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(a)...)

	return command
}

func CreateBinaryInvisibilityMsg(unitID int) []byte {
	command := []byte{33}

	command = append(command, game_math.GetIntBytes(unitID)...)

	return command
}

func CreateBinaryActiveInvisibilityMsg(unitID, ms int) []byte {
	command := []byte{34}

	command = append(command, game_math.GetIntBytes(unitID)...)
	command = append(command, game_math.GetIntBytes(ms)...)

	return command
}

func CreateBinaryDeactivateInvisibilityMsg(id, x, y int) []byte {
	command := []byte{35}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)

	return command
}

func CreateBinaryGravitySquareMsg(x, y, r int) []byte {
	command := []byte{22}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(r)...)

	return command
}

func CreateBinarySuicideProgressMsg(current, deadTime int) []byte {
	command := []byte{36}

	command = append(command, game_math.GetIntBytes(current)...)
	command = append(command, game_math.GetIntBytes(deadTime)...)

	return command
}

func CreateBinarySuicideMsg(unitID int) []byte {
	command := []byte{37}

	command = append(command, game_math.GetIntBytes(unitID)...)

	return command
}

func CreateBinaryAfterburnerMsg(unitID int, w bool) []byte {
	command := []byte{32}

	command = append(command, game_math.GetIntBytes(unitID)...)
	command = append(command, game_math.BoolToByte(w))

	return command
}

func RopeCatchMsg(x, y int) []byte {
	// TODO в случае нужды переделать на 117 событие
	command := []byte{29}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)

	return command
}

func ZoneHealRun(x, y, count int) []byte {
	command := []byte{52}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, byte(count))

	return command
}

func RopeRender(r *rope.Rope) []byte {
	command := []byte{53}

	command = append(command, byte(len(r.Points)))
	for _, p := range r.Points {
		command = append(command, game_math.GetIntBytes(int(p.GetX()))...)
		command = append(command, game_math.GetIntBytes(int(p.GetY()))...)
	}

	return command
}

func RopeMovePoint(point *rope.Point) []byte {
	command := []byte{54}

	command = append(command, game_math.GetIntBytes(int(point.ID))...)
	command = append(command, game_math.GetIntBytes(int(point.Position.X))...)
	command = append(command, game_math.GetIntBytes(int(point.Position.Y))...)

	return command
}

func CreateBinaryAnomaly(anomalies []visible_anomaly.VisibleAnomaly) []byte {
	command := []byte{55}

	for _, a := range anomalies {
		command = append(command, game_math.GetIntBytes(a.Rotate)...)
		command = append(command, byte(a.TypeAnomaly))
		command = append(command, byte(a.Signal))
	}

	return command
}

func CreateBinaryTransportBlock(id, x, y, z, r, toX, toY int) []byte {
	command := []byte{56}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(z)...)
	command = append(command, game_math.GetIntBytes(r)...)
	command = append(command, game_math.GetIntBytes(toX)...)
	command = append(command, game_math.GetIntBytes(toY)...)

	return command
}

func CreateInventoryBin(sourceType string, sourceID int, inventory *inventory.Inventory, forceOpen bool, cap int, objectOwner, remote bool) []byte {
	// [1[eventID], 1[source_type], 4[source_id] slots_data]
	command := []byte{58, byte(_const.SourceItemBin[sourceType])}

	command = append(command, game_math.GetIntBytes(sourceID)...)
	command = append(command, game_math.GetIntBytes(inventory.GetSize())...)
	command = append(command, game_math.GetIntBytes(cap)...)
	command = append(command, game_math.BoolToByte(forceOpen))
	command = append(command, game_math.BoolToByte(objectOwner))
	command = append(command, game_math.BoolToByte(remote))

	for slot := range inventory.GetSlotsChan() {
		// slots_data [4[quantity], 1[type], 4[item_id], 4[hp], 4[max_hp], 4[size], 4[number], 4[access_user_id], 1[infinite], 4[place_user_id], 4[tax], 1[find], 4[find_count],
		// 1[in_map], 1[item_type], 1[name_size], ???[name], 1[item_item_name_size], ???[item_item_name], 1[icon_size], ???[icon]]

		if slot.Item == nil {
			continue
		}
		command = append(command, game_math.GetIntBytes(slot.Quantity)...)
		command = append(command, byte(_const.ItemBinTypes[slot.Type]))
		command = append(command, game_math.GetIntBytes(slot.ItemID)...)
		command = append(command, game_math.GetIntBytes(slot.HP)...)
		command = append(command, game_math.GetIntBytes(slot.MaxHP)...)
		command = append(command, game_math.GetIntBytes(slot.GetSize())...)
		command = append(command, game_math.GetIntBytes(slot.Number)...)
		command = append(command, game_math.GetIntBytes(slot.AccessUserID)...)
		command = append(command, game_math.BoolToByte(slot.Infinite))
		command = append(command, game_math.GetIntBytes(slot.PlaceUserID)...)
		command = append(command, game_math.GetIntBytes(slot.Tax)...)
		command = append(command, game_math.BoolToByte(slot.Find))
		command = append(command, game_math.GetIntBytes(slot.FindCount)...)

		command = append(command, game_math.BoolToByte(slot.Item.InMap))
		command = append(command, byte(_const.ItemBinTypes[slot.Item.ItemType]))
		command = append(command, byte(slot.Item.TypeSlot))

		command = append(command, byte(len([]byte(slot.Item.Name))))
		command = append(command, []byte(slot.Item.Name)...)
		command = append(command, byte(len([]byte(slot.Item.ItemName))))
		command = append(command, []byte(slot.Item.ItemName)...)
		command = append(command, byte(len([]byte(slot.Item.Icon))))
		command = append(command, []byte(slot.Item.Icon)...)
	}

	return command
}

func CreateMiningTargetsBin(ids map[int]bool) []byte {

	command := []byte{59}

	for id, allow := range ids {
		command = append(command, game_math.GetIntBytes(id)...)
		command = append(command, game_math.BoolToByte(allow))
	}

	return command
}

func CreateMiningBin(unitID, x, y, toX, toY int, crack bool) []byte {

	command := []byte{60}

	command = append(command, game_math.GetIntBytes(unitID)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(toX)...)
	command = append(command, game_math.GetIntBytes(toY)...)
	command = append(command, game_math.BoolToByte(crack))

	return command
}

func CreateMiningProgressBin(id, progress int) []byte {

	command := []byte{61}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, byte(progress))

	return command
}

func CreateHitItemBin(count int, typeItem string, idItem, x, y int) []byte {

	command := []byte{62}

	command = append(command, game_math.GetIntBytes(count)...)
	command = append(command, byte(_const.ItemBinTypes[typeItem]))
	command = append(command, game_math.GetIntBytes(idItem)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)

	return command
}

func CreateDiggerTargetBin(x, y, radius int) []byte {

	command := []byte{63}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(radius)...)

	return command
}

func CreateDismantlingTargetsBin(ids map[int]bool) []byte {

	command := []byte{64}

	for id, allow := range ids {
		command = append(command, game_math.GetIntBytes(id)...)
		command = append(command, game_math.BoolToByte(allow))
	}

	return command
}

func CreateBuildBin(unitID, x, y, toX, toY int) []byte {

	command := []byte{65}

	command = append(command, game_math.GetIntBytes(unitID)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(toX)...)
	command = append(command, game_math.GetIntBytes(toY)...)

	return command
}

func CreateGravityGunTargetBin(x, y, r, maxAngle, radius, catchX, catchY, catchRadius int) []byte {
	command := []byte{66}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(r)...)
	command = append(command, game_math.GetIntBytes(maxAngle)...)
	command = append(command, game_math.GetIntBytes(radius)...)
	command = append(command, game_math.GetIntBytes(catchX)...)
	command = append(command, game_math.GetIntBytes(catchY)...)
	command = append(command, game_math.GetIntBytes(catchRadius)...)

	return command
}

func CreateGravityGunRun(x, y, r, rd int) []byte {
	command := []byte{67}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(r)...)
	command = append(command, game_math.GetIntBytes(rd)...)
	return command
}

func CreateGravityGunDrop(x, y, r int) []byte {
	command := []byte{68}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(r)...)
	return command
}

func CreateRepairTargetsBin(targets []*target.Target) []byte {

	command := []byte{69}

	for _, t := range targets {
		command = append(command, byte(_const.MapBinItems[t.Type]))
		command = append(command, game_math.GetIntBytes(t.ID)...)
		command = append(command, game_math.BoolToByte(t.Attack))
	}

	return command
}

func CreateRepairLaserBin(unitID, x, y, toX, toY int) []byte {

	command := []byte{70}

	command = append(command, game_math.GetIntBytes(unitID)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(toX)...)
	command = append(command, game_math.GetIntBytes(toY)...)

	return command
}

func CreatePortalIntoBin(id, x, y int) []byte {

	command := []byte{71}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)

	return command
}

func CreatePortalOutBin(x, y int) []byte {

	command := []byte{72}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)

	return command
}

func CreatePortalDropExitBin(x, y, angle int) []byte {
	command := []byte{73}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(angle)...)

	return command
}

func CreateExitTimeBin(time int) []byte {
	command := []byte{74}
	command = append(command, game_math.GetIntBytes(time)...)
	return command
}

func CreateViolatorsBin(vs []*violator.Violator) []byte {
	command := []byte{75}

	for _, v := range vs {
		command = append(command, game_math.GetIntBytes(v.PlayerID)...)
		command = append(command, byte(getVioTypeInt(v.Type, v.GetTime())))
	}

	return command
}

func CreateSignalsBin(ss []*violator.DistressSignals) []byte {
	command := []byte{76}

	for _, s := range ss {
		command = append(command, game_math.GetIntBytes(s.X)...)
		command = append(command, game_math.GetIntBytes(s.Y)...)
	}

	return command
}

func CreatePumpAnimateBin(x, y, r int) []byte {
	command := []byte{77}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(r)...)

	return command
}

func CreateInventoryScannerTargetsBin(targets []*target.Target) []byte {

	command := []byte{78}

	for _, t := range targets {
		command = append(command, byte(_const.MapBinItems[t.Type]))
		command = append(command, game_math.GetIntBytes(t.ID)...)
		command = append(command, game_math.BoolToByte(t.Attack))
	}

	return command
}

func CreateInventoryScannerFailedBin(failed bool) []byte {
	command := []byte{79}
	command = append(command, game_math.BoolToByte(failed))
	return command
}

func CreateInventoryScannerHostileBin() []byte {
	command := []byte{80}
	return command
}

func RotateStructureEquipBinaryMsg(id, rotate, ms, slot int) []byte {
	// [1[eventID], 4[id], 4[r], 4[ms], 4[ts], 4[s]]
	command := []byte{81}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, game_math.GetIntBytes(rotate)...)
	command = append(command, game_math.GetIntBytes(ms)...)
	command = append(command, game_math.GetIntBytes(slot)...)

	return command
}

func WarSectorStateBinaryMsg(maxPoints int, fractionPoints map[string]int, bases []*spawn.Spawn, owner string, phase bool) []byte {

	command := []byte{82}

	if owner == "" {
		owner = "Empty_Fraction"
	}

	command = append(command, game_math.GetIntBytes(maxPoints)...)
	command = append(command, _const.FractionByte[owner])
	command = append(command, game_math.BoolToByte(phase))

	command = append(command, byte(len(fractionPoints)))

	for fr, p := range fractionPoints {
		command = append(command, _const.FractionByte[fr])
		command = append(command, game_math.GetIntBytes(p)...)
	}

	command = append(command, byte(len(bases)))
	for _, base := range bases {
		command = append(command, game_math.GetIntBytes(base.ID)...)
		command = append(command, _const.FractionByte[base.CaptureTeam])
		command = append(command, byte(int(math.Floor(base.Capture))))
		command = append(command, game_math.BoolToByte(base.CaptureFact))
	}

	return command
}

func PointsBinaryMsg(t string, count int) []byte {
	command := []byte{83}

	command = append(command, _const.PointType[t])
	command = append(command, game_math.GetIntBytes(count)...)

	return command
}

func KillBinaryMsg(deadName, deadFraction, killerName, killerFraction string, ammoID int) []byte {
	command := []byte{84}

	command = append(command, game_math.GetIntBytes(ammoID)...)
	command = append(command, _const.FractionByte[deadFraction])
	command = append(command, _const.FractionByte[killerFraction])

	command = append(command, byte(len(deadName)))
	command = append(command, deadName...)

	command = append(command, byte(len(killerName)))
	command = append(command, killerName...)

	return command
}

func CreatePlayerVioBin(vType string, time int, strikes int) []byte {
	command := []byte{85}

	command = append(command, game_math.GetIntBytes(time)...)
	command = append(command, byte(getVioTypeInt(vType, time)))
	command = append(command, byte(strikes))

	return command
}

func CreateUpRankBin(up bool) []byte {
	command := []byte{86}

	command = append(command, byte(game_math.BoolToByte(up)))

	return command
}

func ServerTime(unixTime int64) []byte {
	command := []byte{87}
	command = append(command, game_math.GetInt64Bytes(unixTime)...)
	return command
}

func getVioTypeInt(vType string, time int) byte {
	vTypeInt := byte(0)

	if vType == "" {
		return 0
	}

	if vType == _const.PK {
		vTypeInt = 1
	}

	if vType == _const.PVP {
		vTypeInt = 2
	}

	if vType == _const.PVP && time <= 3 {
		vTypeInt = 3
	}

	return vTypeInt
}

func OpenAnomaly(x, y, t int) []byte {
	command := []byte{88}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(t)...)

	return command
}

func DiableObject(id int, disable bool) []byte {
	command := []byte{89}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, game_math.BoolToByte(disable))

	return command
}

func SiegeBinaryMsg(active bool, baseHP, baseShield int, hostiles map[int]int, endTime int64) []byte {
	command := []byte{90}

	command = append(command, game_math.BoolToByte(active))
	command = append(command, game_math.GetInt64Bytes(endTime)...)
	command = append(command, game_math.GetIntBytes(baseHP)...)
	command = append(command, game_math.GetIntBytes(baseShield)...)

	for corporationID, count := range hostiles {
		command = append(command, game_math.GetIntBytes(corporationID)...)
		command = append(command, game_math.GetIntBytes(count)...)
	}

	return command
}

func DangerAnomaly(power int) []byte {
	command := []byte{91}

	command = append(command, game_math.GetIntBytes(power)...)

	return command
}

func TargetInfo(targetID, ownerID, weaponTargetID, x, y, ralation, corpID, fear int, weaponTargetType, fraction, targetType string) []byte {
	command := []byte{92}

	t, ok := _const.MapBinItems[weaponTargetType]
	if !ok {
		fmt.Println("unknown type object: ", weaponTargetType)
	}

	tt, ok := _const.MapBinItems[targetType]
	if !ok {
		fmt.Println("unknown type object: ", weaponTargetType)
	}

	command = append(command, game_math.GetIntBytes(targetID)...)
	command = append(command, game_math.GetIntBytes(ownerID)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(corpID)...)
	command = append(command, byte(t))
	command = append(command, game_math.GetIntBytes(weaponTargetID)...)
	command = append(command, _const.FractionByte[fraction])
	command = append(command, game_math.GetIntBytes(ralation)...)
	command = append(command, byte(fear))
	command = append(command, byte(tt))

	return command
}

func GameTime(h, m int) []byte {
	command := []byte{93}

	command = append(command, byte(h))
	command = append(command, byte(m))

	return command
}

func AddDecal(unitID, x, y, id, angle int) []byte {
	command := []byte{94}

	command = append(command, game_math.GetIntBytes(unitID)...)
	command = append(command, byte(x))
	command = append(command, byte(y))
	command = append(command, byte(id))
	command = append(command, byte(angle))

	return command
}

func SetUnrepairableDamage(unrepairableDamage int) []byte {
	command := []byte{95}

	command = append(command, byte(unrepairableDamage))

	return command
}

func ShieldAnimate(id, r, x, y int) []byte {
	command := []byte{96}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, byte(r))
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)

	return command
}

func BodyShieldOn(id int) []byte {
	command := []byte{97}
	command = append(command, game_math.GetIntBytes(id)...)
	return command
}

func PumpMiniGame(performance byte, current, border, need int, reload, destroy bool) []byte {
	command := []byte{98}

	command = append(command, performance)
	command = append(command, game_math.GetIntBytes(current)...)
	command = append(command, byte(border))
	command = append(command, game_math.GetIntBytes(need)...)
	command = append(command, game_math.BoolToByte(reload))
	command = append(command, game_math.BoolToByte(destroy))

	return command
}

func CreateCombineTargetsBin(ids map[int]bool) []byte {

	command := []byte{101}

	for id, allow := range ids {
		command = append(command, game_math.GetIntBytes(id)...)
		command = append(command, game_math.BoolToByte(allow))
	}

	return command
}

func CreateBinaryPollenMove(id, x, y, z, r int) []byte {
	command := []byte{102}

	command = append(command, game_math.GetIntBytes(id)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(z)...)
	command = append(command, game_math.GetIntBytes(r)...)

	return command
}

type BuildPointer interface {
	GetX() int
	GetY() int
	GetPass() bool
	GetStructure() bool
}

func CreateBinaryZonePlace(objData []byte, points []BuildPointer) []byte {
	command := []byte{103}

	command = append(command, game_math.GetIntBytes(len(objData))...)
	command = append(command, objData...)

	pointData := make([]byte, 0, len(points)*3)
	for _, p := range points {
		pointData = append(pointData, byte(p.GetX()))
		pointData = append(pointData, byte(p.GetY()))
		pointData = append(pointData, game_math.BoolToByte(p.GetPass()))
		pointData = append(pointData, game_math.BoolToByte(p.GetStructure()))
	}

	command = append(command, game_math.GetIntBytes(len(pointData))...)
	command = append(command, pointData...)

	return command
}

type Pointer interface {
	GetX() int
	GetY() int
}

func CreateBinaryDigPlace(points []Pointer) []byte {
	command := []byte{104}

	pointData := make([]byte, 0, len(points)*2)
	for _, p := range points {
		pointData = append(pointData, byte(p.GetX()))
		pointData = append(pointData, byte(p.GetY()))
	}

	command = append(command, game_math.GetIntBytes(len(pointData))...)
	command = append(command, pointData...)

	return command
}

type EnergyPointer interface {
	GetX() int
	GetY() int
	GetEnergy() int
	GetPass() bool
	GetEven() bool
}

func CreateBinaryPipePlace(points []Pointer, current []BuildPointer, energy []EnergyPointer, open bool) []byte {
	command := []byte{105}

	pointData := make([]byte, 0, len(points)*2)
	for _, p := range points {
		pointData = append(pointData, byte(p.GetX()))
		pointData = append(pointData, byte(p.GetY()))
	}

	command = append(command, game_math.GetIntBytes(len(pointData))...)
	command = append(command, pointData...)

	pointData = pointData[:0]
	for _, p := range current {
		pointData = append(pointData, byte(p.GetX()))
		pointData = append(pointData, byte(p.GetY()))
		pointData = append(pointData, game_math.BoolToByte(p.GetStructure()))
	}

	command = append(command, game_math.GetIntBytes(len(pointData))...)
	command = append(command, pointData...)

	pointData = pointData[:0]
	for _, p := range energy {
		pointData = append(pointData, byte(p.GetX()))
		pointData = append(pointData, byte(p.GetY()))
		pointData = append(pointData, game_math.GetIntBytes(p.GetEnergy())...)
		pointData = append(pointData, game_math.BoolToByte(p.GetPass()))
		pointData = append(pointData, game_math.BoolToByte(p.GetEven()))
	}

	command = append(command, game_math.GetIntBytes(len(pointData))...)
	command = append(command, pointData...)

	command = append(command, game_math.BoolToByte(open))

	return command
}

func RTSTargetInfo(targetID int, targetType string, ownerID, x, y, ralation, ce, ct int) []byte {
	command := []byte{106}

	tt, ok := _const.MapBinItems[targetType]
	if !ok {
		fmt.Println("unknown type object: ", targetType)
	}

	command = append(command, game_math.GetIntBytes(targetID)...)
	command = append(command, byte(tt))
	command = append(command, game_math.GetIntBytes(ownerID)...)
	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(ralation)...)
	command = append(command, game_math.GetIntBytes(ce)...)
	command = append(command, game_math.GetIntBytes(ct)...)

	return command
}

type RecipePointer interface {
	GetX() int
	GetY() int
	GetRecipe() int
	GetEven() bool
}

func RecipeStructure(rp []RecipePointer) []byte {
	command := []byte{107}

	for _, p := range rp {
		command = append(command, byte(p.GetX()))
		command = append(command, byte(p.GetY()))
		command = append(command, byte(p.GetRecipe()))
		command = append(command, game_math.BoolToByte(p.GetEven()))
	}

	return command
}

func ZeroPriority(rp []RecipePointer) []byte {
	command := []byte{108}

	pointData := make([]byte, 0, len(rp)*3)
	for _, p := range rp {
		pointData = append(pointData, byte(p.GetX()))
		pointData = append(pointData, byte(p.GetY()))
		pointData = append(pointData, game_math.BoolToByte(p.GetEven()))
	}

	command = append(command, pointData...)

	return command
}

type ItemCount struct {
	ItemType string
	ItemID   int
	Quantity int
}

func CountResource(resourceCounts []ItemCount) []byte {
	command := []byte{109}

	for _, rc := range resourceCounts {
		itemTypeByte, exists := _const.ItemBinTypes[rc.ItemType]
		if !exists {
			continue
		}

		command = append(command, byte(itemTypeByte))
		command = append(command, game_math.GetIntBytes(rc.ItemID)...)
		command = append(command, game_math.GetIntBytes(rc.Quantity)...)
	}

	return command
}

func ElevationStatus(status string) []byte {
	command := []byte{110}

	if status == "down" {
		command = append(command, 1)
	} else {
		command = append(command, 0)
	}

	return command
}

func MouseSelectMod(mod string) []byte {
	command := []byte{111}

	if mod == "Pipes" {
		command = append(command, 1)
	} else {
		command = append(command, 0)
	}

	return command
}

func CreateDropItemBin(x, y, toX, toY int) []byte {
	command := []byte{112}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(toX)...)
	command = append(command, game_math.GetIntBytes(toY)...)

	return command
}

func CreateUnitPathBin(points []Pointer) []byte {
	command := []byte{113}

	pointData := make([]byte, 0, len(points)*8)
	for _, p := range points {
		pointData = append(pointData, game_math.GetIntBytes(p.GetX())...)
		pointData = append(pointData, game_math.GetIntBytes(p.GetY())...)
	}

	command = append(command, pointData...)
	return command
}

type PlacePointer interface {
	GetX() int
	GetY() int
	GetID() int
	GetType() string
	GetRotate() int
}

var matchingIDs = map[int]byte{
	584: 15,
	583: 14,
	582: 13,
	581: 12,
	580: 11,
	579: 10,
	578: 9,
	577: 8,
	576: 7,
	575: 6,
	574: 5,
	573: 4,
	572: 3,
	571: 2,
	570: 1,
}

var matchingType = map[string]byte{
	"mine_place": 1,
	"mine_dirt":  2,
	"mine_lava":  3,
}

func CreatePlaceDataBin(points []PlacePointer, size int) []byte {
	command := []byte{114}

	pointData := make([]byte, 0, len(points)*5)
	for _, p := range points {
		byteID, ok := matchingIDs[p.GetID()]
		if !ok {
			continue
		}

		byteType, ok := matchingType[p.GetType()]
		if !ok {
			continue
		}

		pointData = append(pointData, byte(p.GetX()/size))
		pointData = append(pointData, byte(p.GetY()/size))
		pointData = append(pointData, byteID)
		pointData = append(pointData, byteType)
		pointData = append(pointData, byte(p.GetRotate()/2))
	}

	command = append(command, pointData...)
	return command
}

func CreateWarpBin(x, y, angle int, in bool) []byte {
	command := []byte{115}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(angle)...)
	command = append(command, game_math.BoolToByte(in))

	return command
}

func CreateBinaryEvacuationTimeMsg(t int) []byte {
	command := []byte{116}
	command = append(command, game_math.GetIntBytes(t)...)
	return command
}

func CreateSoundMsg(x, y, soundID, k int, force bool) []byte {
	command := []byte{117}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, byte(soundID))
	command = append(command, byte(k))
	command = append(command, game_math.BoolToByte(force))

	return command
}

func CreateBossStateMsg(pID, bossID int, currentHP, maxHP int, phases, percentShield int) []byte {
	command := []byte{118}

	command = append(command, game_math.GetIntBytes(pID)...)
	command = append(command, byte(bossID))
	command = append(command, game_math.GetIntBytes(currentHP)...)
	command = append(command, game_math.GetIntBytes(maxHP)...)
	command = append(command, byte(phases))
	command = append(command, byte(percentShield))

	return command
}

func CreateWindStateMsg(a float64, f, t int) []byte {
	command := []byte{119}

	command = append(command, byte(a/2))
	command = append(command, byte(f))
	command = append(command, byte(t))

	return command
}

func CreateCamTargetMsg(x, y, radius int) []byte {
	command := []byte{120}

	command = append(command, game_math.GetIntBytes(x)...)
	command = append(command, game_math.GetIntBytes(y)...)
	command = append(command, game_math.GetIntBytes(radius)...)

	return command
}
