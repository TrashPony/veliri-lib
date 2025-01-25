package _package

type Package struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Size      int    `json:"size"`
	ItemType  string `json:"item_type"`
	ItemId    int    `json:"item_id"`
	ItemName  string `json:"item_name"`
	ItemCount int    `json:"item_count"`
}

func (p *Package) GetName() string {
	return p.Name
}

func (p *Package) GetSize() int {
	return p.Size
}

func (p *Package) GetItemType() string {
	return p.ItemType
}

func (p *Package) GetItemName() string {
	return p.ItemName
}

func (p *Package) GetInMap() bool {
	return false
}

func (p *Package) GetIcon() string {
	return "box"
}

func (p *Package) GetStandardSize() int {
	return 0
}

func (p *Package) GetTypeSlot() int {
	return 0
}
