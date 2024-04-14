package anchor

type Anchor struct {
	Type        string  `json:"type"`
	Scale       float64 `json:"scale"`
	Rotate      int     `json:"rotate"`
	XAnchor     float64 `json:"x_anchor"`
	YAnchor     float64 `json:"y_anchor"`
	RealXAttach int     `json:"real_x_attach"`
	RealYAttach int     `json:"real_y_attach"`
	Height      int     `json:"height"`
	Size        int     `json:"size"`
}
