package trash_item

type TrashItem struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Size          int    `json:"size"`
	Specification string `json:"specification"`
}

func (t *TrashItem) GetName() string {
	return t.Name
}

func (t *TrashItem) GetSize() int {
	return t.Size
}

func (t *TrashItem) GetItemType() string {
	return ""
}

func (t *TrashItem) GetItemName() string {
	return ""
}

func (t *TrashItem) GetInMap() bool {
	return false
}

func (t *TrashItem) GetIcon() string {
	return ""
}

func (t *TrashItem) GetStandardSize() int {
	return 0
}

func (t *TrashItem) GetTypeSlot() int {
	return 0
}

func (t *TrashItem) GetTech() int {
	return 0
}
