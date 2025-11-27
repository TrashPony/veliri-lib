package book

type Book struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Size    int    `json:"size"`
	SkillID int    `json:"skill_id"`
}

func (b *Book) GetName() string {
	return b.Name
}

func (b *Book) GetSize() int {
	return b.Size
}

func (b *Book) GetItemType() string {
	return ""
}

func (b *Book) GetItemName() string {
	return ""
}

func (b *Book) GetInMap() bool {
	return false
}

func (b *Book) GetIcon() string {
	return ""
}

func (b *Book) GetStandardSize() int {
	return 0
}

func (b *Book) GetTypeSlot() int {
	return b.SkillID
}

func (b *Book) GetTech() int {
	return 0
}
