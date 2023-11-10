package visible_anomaly

type VisibleAnomaly struct {
	Signal      int  `json:"signal"`
	Rotate      int  `json:"rotate"`
	TypeAnomaly int  `json:"type_anomaly"`
	Open        bool `json:"open"`
	X           int  `json:"x"`
	Y           int  `json:"y"`
}
