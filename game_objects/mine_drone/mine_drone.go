package mine_drone

type MineDrone struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Size   int    `json:"size"`
	BodyID int    `json:"body_id"`
	Tech   int    `json:"tech"`
}

func (md *MineDrone) GetName() string {
	return md.Name
}

func (md *MineDrone) GetSize() int {
	return md.Size
}

func (md *MineDrone) GetItemType() string {
	return ""
}

func (md *MineDrone) GetItemName() string {
	return ""
}

func (md *MineDrone) GetInMap() bool {
	return false
}

func (md *MineDrone) GetIcon() string {
	return ""
}

func (md *MineDrone) GetStandardSize() int {
	return 0
}

func (md *MineDrone) GetTypeSlot() int {
	return 0
}

func (md *MineDrone) GetTech() int {
	return md.Tech
}
