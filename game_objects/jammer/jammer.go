package jammers

import "github.com/TrashPony/veliri-lib/game_objects/detail"

type Jammer struct {
	UUID string `json:"uuid"`
	ID   int    `json:"id"`

	// описание обьекта к которому закреплен щит, нужно для отрисовки на фронте
	OwnerID   int    `json:"owner_id"`
	OwnerType string `json:"owner_type"`
	TimeOut   int64  `json:"time_out"`
	Type      int    `json:"type"`

	CurrentReload    int `json:"current_reload"` // текущее время до востановления щита
	WorkCycle        int `json:"work_cycle"`     // время 1 цикла работы щита
	CurrentWorkCycle int `json:"-"`
	Radius           int `json:"radius"`    // радиус щита
	UsePower         int `json:"use_power"` // потребляемая энергия за цикл, доп. потребляемая энергия для востановления хп

	X     int `json:"x"` // позиция обекта владельца, что бы принимать его положение на карте
	Y     int `json:"y"`
	MapID int `json:"map_id"`

	off      bool
	Work     bool
	NoMsg    bool
	forceOff bool
	user     jammerUser

	equipSlot *detail.BodyEquipSlot
}

type jammerUser interface {
	GetID() int
	GetType() string
	GetX() int
	GetY() int
	GetPower() int
	SetPower(int)
	GetMapID() int
	SetMapID(int)
}

func (j *Jammer) GetOff() bool {
	return j.off
}

func (j *Jammer) SetOff(off bool) {
	j.off = off
}

func (j *Jammer) SetForceWorkState(off bool) {
	j.forceOff = off
}

func (j *Jammer) UpdatePos() {
	j.X = j.user.GetX()
	j.Y = j.user.GetY()
	j.MapID = j.user.GetMapID()
}

func (j *Jammer) GetForceWorkState() bool {
	return j.forceOff
}

func (j *Jammer) GetUser() jammerUser {
	return j.user
}

func (j *Jammer) SetUser(u jammerUser) {
	j.user = u
}

func (j *Jammer) GetEquipSlot() *detail.BodyEquipSlot {
	return j.equipSlot
}

func (j *Jammer) SetEquipSlot(equip *detail.BodyEquipSlot) {
	j.equipSlot = equip
}
