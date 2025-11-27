package frr

type FRR struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Size int    `json:"size"`
}

func (f *FRR) GetName() string {
	return f.Name
}

func (f *FRR) GetSize() int {
	return f.Size
}

func (f *FRR) GetItemType() string {
	return ""
}

func (f *FRR) GetItemName() string {
	return ""
}

func (f *FRR) GetInMap() bool {
	return false
}

func (f *FRR) GetIcon() string {
	return ""
}

func (f *FRR) GetStandardSize() int {
	return 0
}

func (f *FRR) GetTypeSlot() int {
	return 0
}

func (f *FRR) GetTech() int {
	return 0
}
