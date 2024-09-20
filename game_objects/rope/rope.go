package rope

import (
	"github.com/TrashPony/veliri-lib/game_math"
	"github.com/TrashPony/veliri-lib/game_objects/physical_model"
	"math"
	"sync"
)

var lastPointID = 0
var mx sync.Mutex

func getPointID() int {
	mx.Lock()
	defer mx.Unlock()
	lastPointID++
	return lastPointID
}

func CreateRope(points, segmentLength, mapID, x, y int, radian float64, dist int, ownerType string, ownerID int,
	friction float64, spriteID byte) *Rope {

	speed := float64(dist) / float64(points)

	newRope := &Rope{
		ID:            getPointID(),
		Points:        make([]*Point, 0),
		Constraints:   make([]*Constraints, 0),
		SegmentLength: float64(segmentLength),
		MapID:         mapID,
		OwnerType:     ownerType,
		OwnerID:       ownerID,
	}

	for i := 0; i < points; i++ {
		newPoint := &Point{
			ID:           getPointID(),
			RopeID:       newRope.ID,
			RopePosition: i,
			MaxVelocity:  40,
			Position: &game_math.Vector{
				X: float64(x),
				Y: float64(y),
			},
			SegmentLength: float64(segmentLength),
			Friction:      friction,
			SpriteID:      spriteID,
		}

		newPoint.SetVelocity(&game_math.Vector{
			X: (math.Cos(radian) * (speed / (float64(i - points)))) * -1,
			Y: (math.Sin(radian) * (speed / (float64(i - points)))) * -1,
		})
		newRope.Points = append(newRope.Points, newPoint)
	}

	for i := 0; i < points-1; i++ {
		newConstraint := &Constraints{
			P1:       newRope.Points[i],
			P2:       newRope.Points[i+1],
			Length:   float64(segmentLength),
			excludes: make(map[string][]int),
		}

		newRope.Constraints = append(newRope.Constraints, newConstraint)
	}

	return newRope
}

type Rope struct {
	ID            int            `json:"id"`
	Destroy       *DestroyOption `json:"-"`
	OwnerID       int            `json:"owner_id"`
	OwnerType     string         `json:"owner_type"`
	Points        []*Point       `json:"points"`
	Constraints   []*Constraints `json:"constraints"`
	SegmentLength float64        `json:"segment_length"`
	MapID         int            `json:"map_id"`
	mx            sync.Mutex
}

type DestroyOption struct {
	IndexPoint int
	Direction  bool
	Time       int
}

func (r *Rope) SetID(id int) {
	r.ID = id
}

func (r *Rope) GetRopeEnd() *Point {
	return r.Points[len(r.Points)-1]
}

func (r *Rope) GetPointByPosition(i int) *Point {
	for _, p := range r.Points {
		if p.RopePosition == i {
			return p
		}
	}

	return nil
}

//func (r *Rope) Attach(point, toPoint *Point) {
//	newConstraint := &Constraints{P1: toPoint, P2: point}
//	newConstraint.setLength(r.SegmentLength)
//	r.Constraints = append(r.Constraints, newConstraint)
//}

func (r *Rope) RemovePoint(i int) {
	r.Points = append(r.Points[:i], r.Points[i+1:]...)
	for i, p := range r.Points {
		p.RopePosition = i
	}

	r.Constraints = append(r.Constraints[:i], r.Constraints[i+1:]...)

	r.Constraints[i-1].P1 = r.Points[i-1]
	r.Constraints[i-1].P2 = r.Points[i]

	if len(r.Constraints) > i {
		r.Constraints[i].P1 = r.Points[i]
		r.Constraints[i].P2 = r.Points[i+1]
	}
}

func (r *Rope) SetPinned(p *Point, ph *physical_model.PhysicalModel, pos *game_math.Vector) {
	r.mx.Lock()
	defer r.mx.Unlock()

	p.Pinned = ph
	if p.Pinned != nil {
		p.PinnedPosition = func() *game_math.Vector {
			return &game_math.Vector{X: float64(p.Pinned.X), Y: float64(p.Pinned.Y)}
		}
	}

	if pos != nil {
		p.PinnedPosition = func() *game_math.Vector {
			return pos
		}
	}

	neighborsCount := 5
	neighbors := make([]*Point, 0)

	for _, neighborPoint := range r.Points {
		if neighborPoint.RopePosition >= p.RopePosition-neighborsCount && neighborPoint.RopePosition <= p.RopePosition+neighborsCount {
			neighbors = append(neighbors, neighborPoint)
		}
	}

	for _, c := range r.Constraints {
		for _, neighborPoint := range neighbors {
			if c.P1.ID == neighborPoint.ID || c.P2.ID == neighborPoint.ID {
				if c.excludes[ph.Type] == nil {
					c.excludes[ph.Type] = make([]int, 0)
				}

				c.excludes[ph.Type] = append(c.excludes[ph.Type], ph.ID)
			}
		}
	}
}

type Point struct {
	ID             int                           `json:"id"`
	RopeID         int                           `json:"rope_id"`
	RopePosition   int                           `json:"rope_position"`
	MaxVelocity    float64                       `json:"max_velocity"`
	Pinned         *physical_model.PhysicalModel `json:"pinned"`
	PinnedPosition func() *game_math.Vector      `json:"pinned_position"`
	Position       *game_math.Vector             `json:"position"`
	PrevPosition   *game_math.Vector             `json:"prevPosition"`
	SegmentLength  float64                       `json:"segment_length"`
	Friction       float64                       `json:"friction"`
	SpriteID       byte                          `json:"sprite_id"`
	CacheJson      []byte                        `json:"-"`
	CreateJsonTime int64                         `json:"-"`
}

func (p *Point) GetJSON(mapTime int64) []byte {
	if p.CreateJsonTime == mapTime && len(p.CacheJson) > 0 {
		return p.CacheJson
	}

	var command []byte

	command = append(command, game_math.GetIntBytes(p.ID)...)
	command = append(command, game_math.GetIntBytes(p.RopeID)...)
	command = append(command, game_math.GetIntBytes(p.RopePosition)...)
	command = append(command, game_math.GetIntBytes(int(p.Position.X))...)
	command = append(command, game_math.GetIntBytes(int(p.Position.Y))...)
	command = append(command, p.SpriteID)

	p.CacheJson = command
	p.CreateJsonTime = mapTime

	return command
}

func (p *Point) GetUpdateData(mapTime int64) []byte {
	return []byte{}
}

func (p *Point) GetX() float64 {
	return p.Position.X
}

func (p *Point) GetY() float64 {
	return p.Position.Y
}

func (p *Point) setX(x float64) {
	p.Position.X = x
}

func (p *Point) setY(y float64) {
	p.Position.Y = y
}

func (p *Point) SetVelocity(velocity *game_math.Vector) {
	p.PrevPosition = p.Position.Sub(velocity)
}

func (p *Point) GetVelocity() *game_math.Vector {
	return p.PrevPosition.VecTo(p.Position)
}

func (p *Point) UpdateFriction() {
	vel := p.GetVelocity()
	vel.X *= p.Friction
	vel.Y *= p.Friction
	p.SetVelocity(vel)
}

type Constraints struct {
	P1       *Point  `json:"p_1"`
	P2       *Point  `json:"p_2"`
	Length   float64 `json:"length"`
	excludes map[string][]int
}

func (c *Constraints) AddExcludes(typeExclude string, id int) {
	if c.excludes[typeExclude] == nil {
		c.excludes[typeExclude] = make([]int, 0)
	}

	c.excludes[typeExclude] = append(c.excludes[typeExclude], id)
}

func (c *Constraints) GetExcludes(typeExcludes string) []int {
	return c.excludes[typeExcludes]
}

func (c *Constraints) setLength(length float64) {
	c.Length = length
}

func (c *Constraints) ApplyConstraint(currentLen float64) {

	radius := c.Length / 2
	midPoint := c.P1.Position.VecTo(c.P2.Position).Scale(0.5)
	midPoint = c.P1.Position.Add(midPoint)

	apply := func(p *Point) {
		if p.Pinned == nil {
			p.Position = midPoint.Add(midPoint.VecTo(p.Position).Resize(radius))
		} else if !p.Pinned.Static {

			a1 := midPoint.VecTo(p.Position)
			a2 := a1.Resize((currentLen - c.Length) / 5)

			k := (float64(p.Pinned.Weight) / 2000) * 0.5
			if k < 1 {
				k = 1
			}

			if k > 3 {
				k = 3
			}

			p.Pinned.SubVelocity((a2.X)/k, (a2.Y)/k)
		}
	}

	apply(c.P1)
	apply(c.P2)
}
