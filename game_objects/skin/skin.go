package skin

type Skin struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ItemType string `json:"item_type"`
	ItemID   int    `json:"item_id"`
	Default  bool   `json:"default"`
	Price    int64  `json:"price"`
	Access   bool   `json:"access"`
	Path     string `json:"path"`
	Fraction string `json:"fraction"`
}
