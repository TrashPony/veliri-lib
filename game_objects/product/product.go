package product

type Product struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Size         int      `json:"size"`
	DefaultPrice int      `json:"default_price"`
	FairTrade    []string `json:"fair_trade"`
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetSize() int {
	return p.Size
}

func (p *Product) GetItemType() string {
	return ""
}

func (p *Product) GetItemName() string {
	return ""
}

func (p *Product) GetInMap() bool {
	return false
}

func (p *Product) GetIcon() string {
	return ""
}

func (p *Product) GetStandardSize() int {
	return 0
}

func (p *Product) GetTypeSlot() int {
	return 0
}

func (p *Product) GetTech() int {
	return 0
}
