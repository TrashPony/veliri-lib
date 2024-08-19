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

	AngularVelocity float64 `json:"angular_velocity"`
	XVelocity       float64 `json:"x_velocity"`
	YVelocity       float64 `json:"y_velocity"`

	CacheJson      []byte `json:"-"`
	CreateJsonTime int64  `json:"-"`

	ForceExplosion      bool            `json:"-"`
	AutoActivate        bool            `json:"-"`
	DetonationDistance  int             `json:"-"`
	DetonationTimeOut   int             `json:"detonation_time_out"`
	Attributes          map[string]int  `json:"-"`
	ObjectID            int             `json:"-"` // ид обьекта которые вызывает снаряжения (турель/стена)
	DetonationForceView bool            `json:"-"` // все видят взрыв, независимо от тумана войны
	MapItem             *inventory.Slot `json:"-"`

	BodyRotateValue     int // что бы на фронте пуля имела положение тела не по направлению а по значению
	BodyRotate          bool
	AccumulationPercent int `json:"accumulation_percent"`

	ghost      bool
	stopTimeMS int
	mx         sync.RWMutex
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

	b.CacheJson = command
	b.CreateJsonTime = mapTime

	return command
}

func (b *Bullet) GetUpdateData(mapTime int64) []byte {
	return []byte{}
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
