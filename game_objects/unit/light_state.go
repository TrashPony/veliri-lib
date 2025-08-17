package unit

import "time"

type LightState struct {
	On             bool
	WorkLights     bool
	LastChangeTime int64
	NextInterval   int
	IsActive       bool
	CooldownUntil  int64
}

func (u *Unit) GetLightState() *LightState {
	u.initLights()
	return u.lights
}

func (u *Unit) SetLights(lights bool) {
	u.initLights()
	u.lights.On = lights
}

func (u *Unit) Lights() bool {
	u.initLights()
	return u.lights.On
}

func (u *Unit) LightsWork() bool {
	u.initLights()
	return u.lights.On && u.lights.WorkLights
}

func (u *Unit) initLights() {
	if u.lights == nil {
		u.lights = &LightState{
			LastChangeTime: time.Now().UnixMilli(),
			NextInterval:   500,
			IsActive:       true,
			CooldownUntil:  0,
		}
	}
}
