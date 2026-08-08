package rope

import (
	_const "github.com/TrashPony/veliri-lib/const"
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
	stiffness float64, spriteID byte, maxVelocity float64) *Rope {

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
		posX := float64(x)
		posY := float64(y)

		newPoint := &Point{
			ID:            getPointID(),
			RopeID:        newRope.ID,
			RopePosition:  i,
			MaxVelocity:   maxVelocity,
			Position:      &game_math.Vector{X: posX, Y: posY},
			PrevPosition:  &game_math.Vector{X: posX, Y: posY},
			LastSentX:     x,
			LastSentY:     y,
			SegmentLength: float64(segmentLength),
			Stiffness:     stiffness, // ЗАМЕНЕНО
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
			P1:         newRope.Points[i],
			P2:         newRope.Points[i+1],
			Length:     float64(segmentLength),
			excludes:   make(map[string][]int),
			RopeLength: segmentLength * points,
		}

		newRope.Constraints = append(newRope.Constraints, newConstraint)
	}

	return newRope
}

func (r *Rope) SetStiffness(stiffness float64) {
	for _, p := range r.Points {
		p.Stiffness = stiffness
	}
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
	LifeTime      int            `json:"life_time"`
	mx            sync.Mutex
}

type DestroyOption struct {
	IndexPoint int
	Direction  bool
	Time       int
	Interval   int
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
	if i < 0 || i >= len(r.Points) {
		return // Защита от некорректного индекса
	}

	r.Points = append(r.Points[:i], r.Points[i+1:]...)

	// Пересчитываем позиции
	for idx, p := range r.Points {
		p.RopePosition = idx
	}

	// Безопасное удаление ограничений
	if i > 0 && i < len(r.Constraints) {
		r.Constraints = append(r.Constraints[:i-1], r.Constraints[i:]...)
	} else if i == 0 && len(r.Constraints) > 0 {
		r.Constraints = r.Constraints[1:]
	} else if i == len(r.Points) && len(r.Constraints) > 0 {
		r.Constraints = r.Constraints[:len(r.Constraints)-1]
	}

	// Перепривязываем оставшиеся ограничения к новым соседям
	for idx, c := range r.Constraints {
		if idx < len(r.Points)-1 {
			c.P1 = r.Points[idx]
			c.P2 = r.Points[idx+1]
		}
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
	} else {
		p.PinnedPosition = nil
	}

	if pos != nil {
		p.PinnedPosition = func() *game_math.Vector {
			return pos
		}
	}

	if p.Pinned == nil && pos == nil {
		p.PinnedPosition = nil
		return
	}

	neighborsCount := 5
	neighbors := make([]*Point, 0)

	for _, neighborPoint := range r.Points {
		if neighborPoint.RopePosition >= p.RopePosition-neighborsCount && neighborPoint.RopePosition <= p.RopePosition+neighborsCount {
			neighbors = append(neighbors, neighborPoint)
		}
	}

	// что бы обновить клиент
	p.LastSentX = -1
	p.LastSentY = -1

	if ph != nil {
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
}

type Point struct {
	ID             int                           `json:"id"`
	RopeID         int                           `json:"rope_id"`
	RopePosition   int                           `json:"rope_position"`
	MaxVelocity    float64                       `json:"max_velocity"`
	Pinned         *physical_model.PhysicalModel `json:"pinned"`
	PinnedPosition func() *game_math.Vector      `json:"-"` // Скрываем функцию из JSON
	Position       *game_math.Vector             `json:"position"`

	// ФИЗИКА: Неявная скорость Верле (Position - PrevPosition). Скрываем из JSON.
	PrevPosition *game_math.Vector `json:"-"`

	// СЕТЬ: Последние отправленные клиенту координаты. Скрываем из JSON.
	LastSentX int `json:"-"`
	LastSentY int `json:"-"`

	SegmentLength  float64 `json:"segment_length"`
	Stiffness      float64 `json:"stiffness"`
	SpriteID       byte    `json:"sprite_id"`
	AddSprite      bool
	CacheJson      []byte  `json:"-"`
	CreateJsonTime int64   `json:"-"`
	OldRadius      int     `json:"-"`
	CurrentAccel   float64 `json:"-"`
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
	command = append(command, game_math.BoolToByte(p.AddSprite))

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
	if p.PrevPosition == nil {
		p.PrevPosition = p.Position.Copy()
	}
	p.PrevPosition.X = p.Position.X - velocity.X
	p.PrevPosition.Y = p.Position.Y - velocity.Y
}

func (p *Point) GetVelocity() *game_math.Vector {
	return p.PrevPosition.VecTo(p.Position)
}

func (p *Point) UpdateFriction() {
	vel := p.GetVelocity()

	// Безопасный коэффициент затухания (например, 0.95 означает потерю 5% скорости за тик).
	// Это предотвращает бесконечные колебания веревки, но не дает ей "взорваться",
	// даже если Stiffness > 1.0
	damping := 0.95

	vel.X *= damping
	vel.Y *= damping
	p.SetVelocity(vel)
}

type Constraints struct {
	P1         *Point  `json:"p_1"`
	P2         *Point  `json:"p_2"`
	Length     float64 `json:"length"`
	RopeLength int     `json:"rope_length"`
	excludes   map[string][]int
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

func (c *Constraints) ApplyConstraint(currentLen float64, source *Point, updatePinned bool) {

	radius := c.Length / 2
	midPoint := c.P1.Position.VecTo(c.P2.Position).Scale(0.5)
	midPoint = c.P1.Position.Add(midPoint)

	apply := func(p, anchor *Point) {
		// СЦЕНАРИЙ 1: Абсолютный якорь (координаты или статичный объект). Не двигаем.
		if (p.Pinned == nil && p.PinnedPosition != nil) || (p.Pinned != nil && p.Pinned.Static) {
			return
		}

		// СЦЕНАРИЙ 2: Свободная точка веревки. Стабильный Верле с лимитером.
		if p.Pinned == nil && p.PinnedPosition == nil {
			targetPos := midPoint.Add(midPoint.VecTo(p.Position).Resize(radius))
			deltaX := targetPos.X - p.Position.X
			deltaY := targetPos.Y - p.Position.Y
			p.Position.X += deltaX
			p.Position.Y += deltaY
		}

		if p.Pinned != nil && !p.Pinned.Static {
			dist := game_math.GetBetweenDistFloat(p.Position.X, p.Position.Y, source.Position.X, source.Position.Y)
			effectiveRopeLength := float64(c.RopeLength) * 0.9
			// Трещотка: если сближаемся или веревка провисает, не мешаем
			if dist < float64(p.OldRadius) || dist < effectiveRopeLength {
				p.OldRadius = int(dist)
				return
			}
			p.OldRadius = int(dist)

			totalStretch := dist - effectiveRopeLength
			if totalStretch <= 0 {
				return
			}

			stretch := currentLen - c.Length

			pullDir := p.Position.VecTo(anchor.Position).Norm()

			mass := p.Pinned.GetWeight()
			if mass <= 0 {
				mass = _const.MaxWeight / 4
			}

			rawFactor := (_const.MaxWeight / 4) / mass
			massFactor := math.Pow(rawFactor, 2.5)

			if massFactor > 7.5 {
				massFactor = 7.5
			}
			if massFactor < 0.05 {
				massFactor = 0.05
			}

			objVelX, objVelY := p.Pinned.GetVelocity()
			velAlongRope := (objVelX * pullDir.X) + (objVelY * pullDir.Y)

			var accel float64

			// Базовая сила натяжения
			basePull := stretch * 0.5 * p.Stiffness * massFactor

			// Демпфирование: гасит скорость, если объект уже движется в направлении натяжения
			damping := (-velAlongRope) * 0.8 * p.Stiffness * massFactor
			accel = basePull + damping

			// 4. Учет сопротивления движению (Drag)
			drag := p.Pinned.GetMoveDrag()
			if drag < 0 {
				drag = 0
			}
			if drag > 0.99 {
				drag = 0.99
			}

			accel = accel * (1.0 - drag)

			// 6. Финальные ограничители
			if accel < 0 {
				accel = 0
			}
			if accel > 100.0 {
				accel = 100.0
			}

			p.CurrentAccel += (accel - p.CurrentAccel) * 0.3
			p.Pinned.AddVelocity(pullDir.X*p.CurrentAccel, pullDir.Y*p.CurrentAccel)
		}
	}

	apply(c.P1, c.P2)
	apply(c.P2, c.P1)
}
