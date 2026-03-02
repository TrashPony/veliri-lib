package wasd_move

import (
	"github.com/TrashPony/veliri-lib/game_objects/gunner"
)

func WasdMove(obj MoveObject, g *gunner.Gunner) {

	if obj.GetPhysicalModel().Fly && !(obj.GetChassisType() == "fly" || obj.GetChassisType() == "fly-2" || obj.GetChassisType() == "fly-3" || obj.GetChassisType() == "wind") {
		return
	}

	if obj.GetChassisType() == "wind" {
		wind(obj)
	}

	if obj.GetChassisType() == "antigravity" || obj.GetChassisType() == "fly-3" {
		antigravity(obj, g)
	}

	if obj.GetChassisType() == "" || obj.GetChassisType() == "caterpillars" || obj.GetChassisType() == "fly" || obj.GetChassisType() == "fly-2" {
		caterpillars(obj)
	}

	if obj.GetChassisType() == "wheels" {
		wheel(obj)
	}

	if obj.GetChassisType() == "legs" {
		legs(obj)
	}
}

func wind(obj MoveObject) {
	xV, yV := obj.GetVelocity()
	xR, yR := obj.GetRealPos()

	obj.SetNextPos(xR+xV, yR+yV)
}

func getPercentF(f float64, percent byte) float64 {
	return f * float64(percent) / 100
}
