package box

type Box struct {
	TypeID        int     `json:"type_id"`
	Name          string  `json:"name"`
	Specification string  `json:"specification"`
	Type          string  `json:"type"`
	CapacitySize  float64 `json:"capacity_size"`
	FoldSize      int     `json:"fold_size"`
	Protect       bool    `json:"protect"`
	ProtectLvl    int     `json:"protect_lvl"`
	Underground   bool    `json:"underground"`
	Height        int     `json:"height"`
	Width         int     `json:"width"`
	Length        int     `json:"length"`
	HP            int     `json:"hp"`
}

func (b *Box) GetName() string {
	return b.Name
}

func (b *Box) GetItemType() string {
	return ""
}

func (b *Box) GetItemName() string {
	return ""
}

func (b *Box) GetInMap() bool {
	return false
}

func (b *Box) GetIcon() string {
	return ""
}

func (b *Box) GetStandardSize() int {
	return 0
}

func (b *Box) GetTypeSlot() int {
	return 0
}
