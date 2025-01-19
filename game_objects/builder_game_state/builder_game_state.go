package builder_game_state

type BuilderGameState struct {
	MapID           int64       `json:"map_id"`
	Status          string      `json:"status"`
	ElevationStatus string      `json:"elevation_status"`
	UnitStates      []UnitState `json:"unit_states"`
}

type UnitState struct {
	UnitID int64  `json:"unit_id"`
	Status string `json:"status"`
}
