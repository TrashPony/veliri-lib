package obstacle_point

import (
	"encoding/json"
)

type ObstaclePoint struct {
	ID         int     `json:"-"`
	X          int     `json:"x"`
	Y          int     `json:"y"`
	Radius     int     `json:"radius"`
	Move       bool    `json:"move"`     // если тру то это только для пуль
	Resource   bool    `json:"resource"` // не влияет на колизии с миром, но влияет на то что низя строить, или куда навести мыш
	ParentID   int     `json:"-"`
	ParentType string  `json:"-"`
	Key        string  `json:"-"`
	Height     float64 `json:"height"`
	K          int     `json:"-"`
}

func (o *ObstaclePoint) GetID() int {
	return o.ID
}

func (o *ObstaclePoint) SetID(id int) {
	o.ID = id
}

func (o *ObstaclePoint) GetX() int {
	return o.X
}

func (o *ObstaclePoint) SetX(x int) {
	o.X = x
}

func (o *ObstaclePoint) GetY() int {
	return o.Y
}

func (o *ObstaclePoint) SetY(y int) {
	o.Y = y
}

func (o *ObstaclePoint) GetRadius() int {
	return int(o.Radius)
}

func (o *ObstaclePoint) SetRadius(radius int) {
	o.Radius = radius
}

func (o *ObstaclePoint) GetMove() bool {
	return o.Move
}

func (o *ObstaclePoint) SetMove(move bool) {
	o.Move = move
}

func (o *ObstaclePoint) GetResource() bool {
	return o.Resource
}

func (o *ObstaclePoint) SetResource(resource bool) {
	o.Resource = resource
}

func (o *ObstaclePoint) GetParentID() int {
	return int(o.ParentID)
}

func (o *ObstaclePoint) SetParentID(parentID int) {
	o.ParentID = parentID

}

func (o *ObstaclePoint) GetParentType() string {
	return o.ParentType
}

func (o *ObstaclePoint) SetParentType(parentType string) {
	o.ParentType = parentType
}

func (o *ObstaclePoint) GetKey() string {
	return o.Key
}

func (o *ObstaclePoint) SetKey(key string) {
	o.Key = key
}

func (o *ObstaclePoint) GetHeight() float64 {
	return o.Height
}

func (o *ObstaclePoint) SetHeight(height float64) {
	o.Height = height
}

func (o *ObstaclePoint) GetJSON() string {
	jsonPoint, err := json.Marshal(struct {
		X        int     `json:"x"`
		Y        int     `json:"y"`
		Radius   int     `json:"radius"`
		Move     bool    `json:"move"`
		Resource bool    `json:"resource"`
		Height   float64 `json:"height"`
	}{
		X:        o.X,
		Y:        o.Y,
		Radius:   o.Radius,
		Move:     o.Move,
		Resource: o.Resource,
		Height:   o.Height,
	})
	if err != nil {
		println("geo point to json: ", err.Error())
	}

	return string(jsonPoint)
}
