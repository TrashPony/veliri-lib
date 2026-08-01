package bullet

import (
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/ammo"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"github.com/TrashPony/veliri-lib/game_objects/inventory"
	"github.com/TrashPony/veliri-lib/game_objects/target"
	"math"
	"sync"
	"time"
)

type Bullet struct {
	ID                   int            `json:"id"`
	Weapon               *detail.Weapon `json:"-"`
	Ammo                 *ammo.Ammo     `json:"ammo"`
	EquipID              int            `json:"-"`
	Rotate               float64        `json:"rotate"`
	Artillery            bool           `json:"artillery"`
	X                    int            `json:"x"`
	Y                    int            `json:"y"`
	Z                    float64        `json:"z"` // определяет "высоту" пули (сильнее отдалять тени)
	Speed                int            `json:"speed"`
	Target               *target.Target `json:"target"`
	ChaseTarget          *target.Target `json:"-"`
	OwnerID              int            `json:"owner_id"`   // какой игрок стрелял
	OwnerType            string         `json:"owner_type"` // unit, structure
	OwnerPlayerID        int            `json:"owner_player_id"`
	OwnerFractionWarrior bool           `json:"-"`
	OwnerGroupID         int            `json:"owner_group_id"`
	Fraction             string         `json:"-"`
	MaxRange             int            `json:"max_range"`
	FirePos              int            `json:"-"`
	MapID                int            `json:"map_id"`
	HP                   int            `json:"destroy"`
	LaunchX              int            `json:"-"`
	LaunchY              int            `json:"-"`
	StartX               int            `json:"-"`
	StartY               int            `json:"-"`
	StartZ               float64        `json:"-"`
	StartRadian          float64        `json:"start_radian"`
	Damage               int            `json:"-"`
	EquipDamage          int            `json:"-"`
	IgnoreOwner          bool           `json:"ignore_owner"`
	NoDamageNoMessage    bool           `json:"-"`
	OldDistToTarget      float64        `json:"-"`
	IgnoreShieldUnitID   int            `json:"-"`
	IgnoreAllShield      bool           `json:"-"`
	NoExplosion          bool           `json:"-"`
	AttachUnitID         int            `json:"-"`
	TimeOut              int            `json:"-"`
	UpdatePos            func()         `json:"-"`
	DrawTargetZone       bool           `json:"-"`
	PredictedX           int            `json:"-"`
	PredictedY           int            `json:"-"`
	AllPush              bool           `json:"-"`
	AllDamage            bool           `json:"-"`
	OwnerPlayer          interface{}    `json:"-"`
	Hide                 bool           `json:"-"`
	NoDamage             bool           `json:"-"`
	AmmoRadius           int            `json:"-"`
	AttachConfig         AttachConfig   `json:"-"`
	NoStopForCollision   bool           `json:"-"`

	ImmediateDestruction bool `json:"immediate_destruction"`
	end                  bool
	MaxAngle             float64 `json:"-"`
	MinAngle             float64 `json:"-"`
	//
	RealX            float64 `json:"-"`
	RealY            float64 `json:"-"`
	DistanceTraveled float64 `json:"-"`
	RealSpeed        float64 `json:"-"`
	RadRotate        float64 `json:"-"`
	AttackUnitID     int     `json:"-"`
	AttackStructID   int     `json:"-"`
	LaunchAngle      float64 `json:"-"`
	TicksAlive       int     `json:"-"`

	AngularVelocity float64 `json:"angular_velocity"`
	XVelocity       float64 `json:"x_velocity"`
	YVelocity       float64 `json:"y_velocity"`
	TargetXVelocity float64 `json:"-"`
	TargetYVelocity float64 `json:"-"`
	TargetRotate    float64 `json:"-"`

	CacheJson       []byte    `json:"-"`
	CacheUpdateData CacheData `json:"-"`
	CreateJsonTime  int64     `json:"-"`

	ForceExplosion      bool            `json:"-"`
	AutoActivate        bool            `json:"-"`
	DetonationDistance  int             `json:"-"`
	DetonationTimeOut   int             `json:"detonation_time_out"`
	Attributes          map[string]int  `json:"-"`
	ObjectID            int             `json:"-"` // ид обьекта которые вызывает снаряжения (турель/стена)
	DetonationForceView bool            `json:"-"` // все видят взрыв, независимо от тумана войны
	MapItem             *inventory.Slot `json:"-"`
	EquipType           int             `json:"-"`
	EquipNumber         int             `json:"-"`
	State               int             `json:"-"`
	ClientLag           float64         `json:"-"`
	ParentVelX          float64         `json:"-"`
	ParentVelY          float64         `json:"-"`
	ExcludeUnitIDs      []int           `json:"-"`
	ExcludeObjectIDs    []int           `json:"-"`
	RemoveTarget        bool            `json:"-"`

	BodyRotateValue     int // что бы на фронте пуля имела положение тела не по направлению а по значению
	BodyRotate          bool
	AccumulationPercent int `json:"accumulation_percent"`
	CallBack            func(b *Bullet, damageObjects interface{})
	ChainCount          int `json:"-"`
	ghost               bool
	stopTimeMS          int
	rWAttributes        map[string]int
	mx                  sync.RWMutex
}

type AttachConfig struct {
	AttachedToType string  `json:"attached_to_type"` // "unit", "object", "flore"
	AttachedToID   int     `json:"attached_to_id"`
	AttachDist     float64 `json:"attach_dist"`
	AttachAngle    float64 `json:"attach_angle"`
	RelativeAngle  float64 `json:"relative_angle"`
}

type CacheData struct {
	Data []byte `json:"-"`
	Time int64  `json:"-"`
}

func (b *Bullet) GetFirePos(i int) *game_math.Positions {
	return &game_math.Positions{X: b.X, Y: b.Y}
}

func (b *Bullet) GetMapHeight() float64 {
	return b.Z
}

func (b *Bullet) Ghost() bool {
	return b.ghost
}

func (b *Bullet) SetGhost(g bool) {
	b.ghost = g
}

func (b *Bullet) StopRun(stopTimeMS int) {
	if b.stopTimeMS == 0 {
		go func() {
			for {
				b.stopTimeMS -= 100
				time.Sleep(time.Millisecond * 100)

				if b.stopTimeMS <= 0 {
					b.Speed = 1
					return
				}
			}
		}()
	}
	b.stopTimeMS = stopTimeMS
}

func (b *Bullet) AddVelocity(x float64, y float64) {
	b.XVelocity += x
	b.YVelocity += y
}

func (b *Bullet) GetVelocityRotate() float64 {
	return math.Atan2(b.YVelocity, b.XVelocity)
}

func (b *Bullet) GetX() int {
	return b.X
}

func (b *Bullet) SetX(x int) {
	b.X = x
}

func (b *Bullet) GetY() int {
	return b.Y
}

func (b *Bullet) SetY(y int) {
	b.Y = y
}

func (b *Bullet) GetZ() float64 {
	return b.Z
}

func (b *Bullet) SetZ(z float64) {
	b.Z = z
}

func (b *Bullet) GetRotate() float64 {
	return b.Rotate
}

func (b *Bullet) SetRotate(rotate float64) {
	b.Rotate = rotate
}

func (b *Bullet) GetEnd() bool {
	return b.end
}

func (b *Bullet) SetEnd(end bool) {
	b.end = end
}

func (b *Bullet) GetID() int {
	return b.ID
}

func (b *Bullet) SetID(id int) {
	b.ID = id
}

func (b *Bullet) GetJSON(mapTime int64) []byte {

	if b.CreateJsonTime == mapTime && len(b.CacheJson) > 0 {
		return b.CacheJson
	}

	command := []byte{}

	command = append(command, byte(b.Ammo.ID))
	command = append(command, game_math.GetIntBytes(b.ID)...)
	command = append(command, game_math.GetIntBytes(b.GetX())...)
	command = append(command, game_math.GetIntBytes(b.GetY())...)
	command = append(command, game_math.GetIntBytes(int(b.GetZ()))...)

	if b.BodyRotate {
		command = append(command, game_math.GetIntBytes(b.BodyRotateValue)...)
	} else {
		command = append(command, game_math.GetIntBytes(int(b.GetRotate()))...)
	}

	command = append(command, game_math.BoolToByte(b.Hide))

	b.CacheJson = command
	b.CreateJsonTime = mapTime

	return command
}

func (b *Bullet) GetUpdateData(mapTime int64) []byte {
	if b.CacheUpdateData.Time == mapTime && len(b.CacheUpdateData.Data) > 0 {
		return b.CacheUpdateData.Data
	}

	if b.CacheUpdateData.Data == nil {
		b.CacheUpdateData.Data = []byte{}
	}

	b.CacheUpdateData.Data = b.CacheUpdateData.Data[:0]

	b.CacheUpdateData.Data = append(b.CacheUpdateData.Data, game_math.BoolToByte(b.Hide))

	b.CacheUpdateData.Time = mapTime

	return b.CacheUpdateData.Data
}

func (b *Bullet) GetOwnerPlayerID() int {
	return b.OwnerPlayerID
}

func (b *Bullet) FractionWarrior() bool {
	return b.OwnerFractionWarrior
}

func (b *Bullet) OwnerFraction() string {
	return b.Fraction
}

func (b *Bullet) GetRWAttribute(key string) int {
	b.mx.RLock()
	defer b.mx.RUnlock()
	if b.rWAttributes == nil {
		b.rWAttributes = make(map[string]int)
	}
	return b.rWAttributes[key]
}

func (b *Bullet) SetRWAttribute(key string, value int) {
	b.mx.Lock()
	defer b.mx.Unlock()
	if b.rWAttributes == nil {
		b.rWAttributes = make(map[string]int)
	}
	b.rWAttributes[key] = value
}
