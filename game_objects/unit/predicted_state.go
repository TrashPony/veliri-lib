package unit

import (
	"sync"
)

type FireInputState struct {
	PendingInputs []*FireInput `json:"-"`
	//FireInputStateSnapshot [][]*FireInputStateSnapshot `json:"-"`
	mx sync.Mutex
}

//type FireInputStateSnapshot struct {
//	UnitID              int // защита
//	ServerInputSeq      byte
//	WeaponSlot          byte
//	UnitX               int
//	UnitY               int
//	UnitRotate          float64
//	WeaponRotate        float64
//	WeaponFirePos       []*game_math.Positions
//	Spread              int
//	WeaponSlotCopy      *detail.BodyWeaponSlot
//	WeaponPosInMap      game_math.Positions
//	FirePos             *game_math.Positions
//	WeaponFirePosOne    *game_math.Positions
//	WeaponTarget        *target.Target
//	ServerTime          int64
//	AccumulationCurrent float64
//	AccumulationTimeOut int
//}

//func (u *Unit) CreateFireStateSnapshot(serverInputSeq byte, serverTime int64) {
//	if u.FireInputState == nil {
//		u.FireInputState = &FireInputState{PendingInputs: make([]*FireInput, 0, 256), FireInputStateSnapshot: make([][]*FireInputStateSnapshot, 256)}
//	}
//
//	u.FireInputState.mx.Lock()
//	defer u.FireInputState.mx.Unlock()
//
//	snapshotByWeapon := make([]*FireInputStateSnapshot, 0, 4)
//	for _, ws := range u.RangeWeaponSlots() {
//		if ws == nil || ws.Weapon == nil {
//			continue
//		}
//
//		weaponPosX, weaponPosY := u.GetGunner().GetWeaponPosInMap(ws.Number)
//		fp := u.GetGunner().GetFirePos(ws.Number) // тут мы прокручиваем, поэтому выполняем заранее LastFirePosition
//		wt := u.GetWeaponTarget(ws.Number)
//
//		var copyWT *target.Target
//		if wt != nil {
//			copyWT = wt.GetCopy()
//		}
//
//		snapshotByWeapon = append(snapshotByWeapon, &FireInputStateSnapshot{
//			UnitID:              u.GetID(),
//			ServerInputSeq:      serverInputSeq,
//			WeaponSlot:          byte(ws.Number),
//			UnitX:               u.GetX(),
//			UnitY:               u.GetY(),
//			UnitRotate:          u.GetRotate(),
//			WeaponRotate:        u.GetGunner().GetGunRotate(ws.Number),
//			WeaponFirePos:       u.GetGunner().GetWeaponFirePos(ws.Number),
//			Spread:              ws.Spread,
//			WeaponSlotCopy:      ws.GetCopy(),
//			WeaponPosInMap:      game_math.Positions{X: weaponPosX, Y: weaponPosY},
//			FirePos:             fp,
//			WeaponFirePosOne:    u.GetGunner().GetWeaponFirePosOne(ws.Number, ws.GetLastFirePosition()),
//			WeaponTarget:        copyWT,
//			ServerTime:          serverTime,
//			AccumulationCurrent: ws.AccumulationCurrent,
//			AccumulationTimeOut: ws.AccumulationTimeOut,
//		})
//	}
//
//	u.FireInputState.FireInputStateSnapshot[serverInputSeq] = snapshotByWeapon
//}

type FireInput struct {
	ServerInputSeq byte `json:"-"`
	WeaponSlot     byte `json:"-"`
	ClientLag      byte `json:"-"`
	Tick           byte `json:"-"`
}

func (u *Unit) AddFireInput(clientSeq, serverInputSeq, currentInputSeq, weaponNumber byte, x, y int, serverTime int64) (bool, string) {
	if u.FireInputState == nil {
		// u.FireInputState = &FireInputState{PendingInputs: make([]*FireInput, 0, 256), FireInputStateSnapshot: make([][]*FireInputStateSnapshot, 256)}
		u.FireInputState = &FireInputState{PendingInputs: make([]*FireInput, 0, 256)}
	}

	u.FireInputState.mx.Lock()
	defer u.FireInputState.mx.Unlock()

	ws := u.GetGunner().GetWeaponSlot(int(weaponNumber))
	if ws == nil || ws.Weapon == nil {
		return false, "no weapon"
	}

	data := &FireInput{
		ServerInputSeq: serverInputSeq,
		WeaponSlot:     weaponNumber,
		ClientLag:      calculateClientLag(currentInputSeq, serverInputSeq),
	}

	u.mx.Lock()
	defer u.mx.Unlock()

	for i, _ := range u.FireInputState.PendingInputs {
		if i == 0 {
			u.FireInputState.PendingInputs[i] = data
			return true, "update"
		}
		//if input.WeaponSlot == weaponNumber && input.ServerInputSeq == serverInputSeq {
		//	u.FireInputState.PendingInputs[i] = data
		//	return true, "update"
		//}
	}

	u.FireInputState.PendingInputs = append(u.FireInputState.PendingInputs, data)
	return true, "add"
}

func (u *Unit) GetFireInputs(ws int) []*FireInput {
	u.mx.Lock()
	defer u.mx.Unlock()

	if u.FireInputState == nil || len(u.FireInputState.PendingInputs) == 0 {
		return nil
	}

	var toProcess []*FireInput
	for _, input := range u.FireInputState.PendingInputs {
		if byte(ws) == input.WeaponSlot {
			toProcess = append(toProcess, input)
		}
	}

	return toProcess
}

func (u *Unit) ClearFireInput() {
	u.mx.Lock()
	defer u.mx.Unlock()

	if u.FireInputState == nil {
		// u.FireInputState = &FireInputState{PendingInputs: make([]*FireInput, 0, 256), FireInputStateSnapshot: make([][]*FireInputStateSnapshot, 256)}
		u.FireInputState = &FireInputState{PendingInputs: make([]*FireInput, 0, 256)}
	}

	u.FireInputState.PendingInputs = u.FireInputState.PendingInputs[:0]
}

func calculateClientLag(serverInputSeq, clientSeq byte) byte {
	lag := int(serverInputSeq) - int(clientSeq)
	if lag < 0 {
		lag += 256
	}

	if lag > 2 {
		lag = 2
	}

	return byte(lag)
}
