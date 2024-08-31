package resource

type Resource struct {
	TypeID        int    `json:"type_id"`
	Name          string `json:"name"`
	Specification string `json:"specification"`
	Size          int    `json:"size"`

	// описывает что выходит из этих ресурсов при переработке
	EnrichedThorium int `json:"enriched_thorium"`
	Iron            int `json:"iron"`
	Copper          int `json:"copper"`
	Titanium        int `json:"titanium"`
	Silicon         int `json:"silicon"`
	Plastic         int `json:"plastic"`

	RecyclingAlgorithm string `json:"recycling_algorithm"`
}

func (r *Resource) GetName() string {
	return r.Name
}

func (r *Resource) GetSize() int {
	return r.Size
}

func (r *Resource) GetItemType() string {
	return ""
}

func (r *Resource) GetItemName() string {
	return ""
}

func (r *Resource) GetInMap() bool {
	return false
}

func (r *Resource) GetIcon() string {
	return ""
}

func (r *Resource) GetStandardSize() int {
	return 0
}

func (r *Resource) GetTypeSlot() int {
	return 0
}

func (r *Resource) GetEnrichedThorium() int {
	return r.EnrichedThorium
}

func (r *Resource) GetIron() int {
	return r.Iron
}

func (r *Resource) GetCopper() int {
	return r.Copper
}

func (r *Resource) GetTitanium() int {
	return r.Titanium
}

func (r *Resource) GetSilicon() int {
	return r.Silicon
}

func (r *Resource) GetPlastic() int {
	return r.Plastic
}

func (r *Resource) GetSteel() int {
	return 0
}

func (r *Resource) GetWire() int {
	return 0
}

func (r *Resource) GetElectronics() int {
	return 0
}

func (r *Resource) GetWires() int {
	return 0
}

func (r *Resource) GetGear() int {
	return 0
}

func (r *Resource) GetTitaniumPlate() int {
	return 0
}

func (r *Resource) GetBatteries() int {
	return 0
}

func (r *Resource) GetArmorItems() int {
	return 0
}

func (r *Resource) GetRecyclingAlgorithm() string {
	return r.RecyclingAlgorithm
}

type RecycledResource struct {
	TypeID        int    `json:"type_id"`
	Name          string `json:"name"`
	Specification string `json:"specification"`
	Size          int    `json:"size"`
	DefaultPrice  int64  `json:"default_price"`
}

func (r *RecycledResource) GetName() string {
	return r.Name
}

func (r *RecycledResource) GetSize() int {
	return r.Size
}

func (r *RecycledResource) GetItemType() string {
	return ""
}

func (r *RecycledResource) GetItemName() string {
	return ""
}

func (r *RecycledResource) GetInMap() bool {
	return false
}

func (r *RecycledResource) GetIcon() string {
	return ""
}

func (r *RecycledResource) GetStandardSize() int {
	return 0
}

func (r *RecycledResource) GetTypeSlot() int {
	return 0
}

type CraftDetail struct {
	TypeID        int    `json:"id"`
	Name          string `json:"name"`
	Specification string `json:"specification"`
	Tech          int    `json:"tech"`
	Size          int    `json:"size"`
}

func (c *CraftDetail) GetName() string {
	return c.Name
}

func (c *CraftDetail) GetSize() int {
	return c.Size
}

func (c *CraftDetail) GetItemType() string {
	return ""
}

func (c *CraftDetail) GetItemName() string {
	return ""
}

func (c *CraftDetail) GetInMap() bool {
	return false
}

func (c *CraftDetail) GetIcon() string {
	return ""
}

func (c *CraftDetail) GetStandardSize() int {
	return 0
}

func (c *CraftDetail) GetTypeSlot() int {
	return 0
}
