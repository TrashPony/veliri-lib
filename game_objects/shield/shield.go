package shield

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_math"
	"sync"
)

type shieldUser interface {
	GetID() int
	GetType() string
	GetX() int
	GetY() int
	GetPower() int
	SetPower(int)
	GetMapID() int
	SetMapID(int)
}

type Shield struct {
	UUID string `json:"uuid"`
	ID   int    `json:"id"`

	// описание обьекта к которому закреплен щит, нужно для отрисовки на фронте
	OwnerID   int    `json:"owner_id"`
	OwnerType string `json:"owner_type"`

	MaxHP            int `json:"max_hp"`         // максимальное хп
	CurrentHP        int `json:"current_hp"`     // текущее хп
	RecoveryHP       int `json:"recovery_hp"`    // востанавливаемое хп за цикл
	Reload           int `json:"reload"`         // время запуска, через которое хп начнет востанавливается если щит перестал получать урон или его сбили
	CurrentReload    int `json:"current_reload"` // текущее время до востановления щита
	WorkCycle        int `json:"work_cycle"`     // время 1 цикла работы щита
	CurrentWorkCycle int `json:"-"`
	Radius           int `json:"radius"`    // радиус щита
	UsePower         int `json:"use_power"` // потребляемая энергия за цикл, доп. потребляемая энергия для востановления хп

	X     int  `json:"x"` // позиция обекта владельца, что бы принимать его положение на карте
	Y     int  `json:"y"`
	MapID int  `json:"map_id"`
	Block bool `json:"block"` // сбивает саряды в обе стороны

	user     shieldUser
	off      bool
	work     bool
	forceOff bool

	CacheJson      []byte `json:"-"`
	CreateJsonTime int64  `json:"-"`
	Key            string `json:"-"`

	mx sync.RWMutex
}

func (s *Shield) UpdatePos(x, y, mapId int) {
	s.X = x
	s.Y = y
	s.MapID = mapId
}

func (s *Shield) GetX() int {
	return s.X
}

func (s *Shield) GetY() int {
	return s.Y
}

func (s *Shield) GetMapID() int {
	return s.MapID
}

func (s *Shield) FindSidePoint(x, y int) bool {
	// метод смотрим находится ли точка внутри щита или снаружи
	// если дистанция от точки больше чем радиус то это снаружи

	dist := int(game_math.GetBetweenDist(x, y, s.GetX(), s.GetY()))
	// true - снаружи
	// false - внутри
	return dist > s.Radius
}

func (s *Shield) CheckPointsBetweenShield(x1, y1, x2, y2, mapID int) bool {
	// метод узнает есть ли между 2мя точками на карте щит
	dist1 := int(game_math.GetBetweenDist(s.GetX(), s.GetY(), x1, y1))
	dist2 := int(game_math.GetBetweenDist(s.GetX(), s.GetY(), x2, y2))

	// если 1 точка ближе чем радиус а другая точка дальше радиуса значит между ними щит shield
	if (dist1 > s.Radius && dist2 < s.Radius) || (dist1 < s.Radius && dist2 > s.Radius) {
		return true
	}

	return false
}

func (s *Shield) SetForceWorkState(off bool) {
	s.forceOff = off
}

func (s *Shield) GetForceWorkState() bool {
	return s.forceOff
}

func (s *Shield) GetWork() bool {
	return s.work
}

func (s *Shield) SetWork(work bool) {
	s.work = work
}

func (s *Shield) GetUser() shieldUser {
	return s.user
}

func (s *Shield) SetUser(u shieldUser) {
	s.user = u
}

func (s *Shield) GetHP() int {
	return s.CurrentHP
}

func (s *Shield) SetHP(hp int) {
	s.CurrentHP = hp
}

func (s *Shield) GetCurrentReload() int {
	return s.CurrentReload
}

func (s *Shield) SetCurrentReload(reload int) {
	s.CurrentReload = reload
}

func (s *Shield) GetOff() bool {
	return s.off
}

func (s *Shield) SetOff(off bool) {
	s.off = off
}

func (s *Shield) Damage(damage int) {
	if s == nil {
		return
	}

	s.SetHP(s.GetHP() - damage)
	if s.GetHP() <= 0 {
		s.SetHP(0)
		s.SetWork(false)
	}

	s.SetCurrentReload(s.Reload) // включаем перегрузку на щит
}

func (s *Shield) GetJSON(mapTime int64) []byte {

	if s.CreateJsonTime == mapTime && len(s.CacheJson) > 0 {
		return s.CacheJson
	}

	command := []byte{}

	command = append(command, game_math.GetIntBytes(s.ID)...)
	command = append(command, game_math.GetIntBytes(s.GetX())...)
	command = append(command, game_math.GetIntBytes(s.GetY())...)
	command = append(command, game_math.GetIntBytes(s.Radius)...)
	command = append(command, game_math.BoolToByte(s.Block))
	command = append(command, game_math.GetIntBytes(s.CurrentHP)...)
	command = append(command, game_math.GetIntBytes(s.MaxHP)...)
	command = append(command, byte(_const.MapBinItems[s.OwnerType]))
	command = append(command, game_math.GetIntBytes(s.OwnerID)...)

	s.CacheJson = command
	s.CreateJsonTime = mapTime

	return s.CacheJson
}

func (s *Shield) GetUpdateData(mapTime int64) []byte {
	// TODO оптимизация и совмещение с src/game_loop/game_loop_view/update_messages.go
	command := []byte{}
	command = append(command, game_math.GetIntBytes(s.CurrentHP)...)
	return command
}
