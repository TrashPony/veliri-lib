package wasd_move

import (
	"github.com/TrashPony/veliri-lib/game_objects/obstacle_point"
	"github.com/TrashPony/veliri-lib/game_objects/physical_model"
)

type MoveObject interface {
	GetStateSenderPos() bool
	SetStateSenderPos(bool)
	GetX() int
	GetY() int
	GetVelocity() (float64, float64)
	SetVelocity(float64, float64)
	MultiplyVelocity(float64, float64)
	AddVelocity(float64, float64)
	GetRotate() float64
	GetID() int
	GetType() string
	SetPos(float64, float64, float64)
	GetRealPos() (float64, float64)
	GetWeight() float64
	GetDirection() bool
	GetCurrentSpeed() float64
	SetPowerMove(float64)
	GetPowerMove() float64
	GetLength() float64
	GetWidth() float64
	GetGeoData() []*obstacle_point.ObstaclePoint
	GetVelocityRotate() float64

	CheckGrowthPower() bool
	CheckGrowthRevers() bool
	CheckLeftRotate() bool
	CheckRightRotate() bool
	CheckHandBrake() bool
	GetTypeControl() int
	GetWasd() *physical_model.WASD
	GetAngleForClassicControl() float64

	GetPowerLeft() float64
	SetPowerLeft(float64)
	GetPowerRight() float64
	SetPowerRight(float64)

	GetAngularVelocity() float64
	SetAngularVelocity(float64)

	GetReverse() float64
	SetReverse(float64)

	GetMoveMaxPower() float64
	GetMaxReverse() float64

	GetPowerFactor() float64
	GetReverseFactor() float64
	GetTurnSpeed() float64

	GetMoveDrag() float64
	GetAngularDrag() float64

	SetWASD(bool, bool, bool, bool, bool, bool, bool, bool, bool)
	IsFly() bool
	GetZ() float64
	SetZ(float64)
	GetNeedZ() float64
	GetTransportUnitID() int
	SetPosFunc(func())
	GetPhysicalModel() *physical_model.PhysicalModel
	SetNextPos(x, y float64)
	GetNextPos() (float64, float64)
	GetChassisType() string
	GetMoveDestroyer() bool
	GetForceNitro() bool
}
