package blueprints

import "github.com/TrashPony/veliri-lib/game_objects/inventory"

type Blueprint struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Specification string `json:"specification"`
	ItemType      string `json:"item_type"`
	ItemId        int    `json:"item_id"`
	ItemName      string `json:"item_name"`
	Icon          string `json:"icon"`
	CraftTime     int    `json:"craft_time"`
	Original      bool   `json:"original"`
	Copies        int    `json:"copies"`
	Count         int    `json:"count"`

	EnrichedThorium int `json:"enriched_thorium"`
	Iron            int `json:"iron"`
	Copper          int `json:"copper"`
	Titanium        int `json:"titanium"`
	Silicon         int `json:"silicon"`
	Plastic         int `json:"plastic"`
	Carbon          int `json:"carbon"`
	Organic         int `json:"organic"`
	Oil             int `json:"oil"`
	Spices          int `json:"spices"`
	Disputes1       int `json:"disputes_1"`
	Disputes2       int `json:"disputes_2"`
	Resin           int `json:"resin"`
	Flower          int `json:"Flower"`
	Steel           int `json:"steel"`
	Wire            int `json:"wire"`
	Electronics     int `json:"electronics"`
	Wires           int `json:"wires"`
	Gear            int `json:"gear"`
	TitaniumPlate   int `json:"titanium_plate"`
	Batteries       int `json:"batteries"`
	ArmorItems      int `json:"armor_items"`
	CarbonPlate     int `json:"carbon_plate"`
	//T2 details
	AdvancedElectronics int              `json:"advanced_electronics"`
	ReinforcedFrame     int              `json:"reinforced_frame"`
	ReinforcedDrives    int              `json:"reinforced_drives"`
	CompositeArmor      int              `json:"composite_armor"`
	AdvancedItems       []inventory.Slot `json:"advanced_items"`
	Variants            []int            `json:"variants"`
	NeedRff             []map[string]int `json:"need_rff"`

	// products
	CivilianWeapon      int `json:"civilian_weapon"`
	FoxCrystals         int `json:"fox_crystals"`
	NeutrinoProcessors  int `json:"neutrino_processors"`
	WeaponParts         int `json:"weapon_parts"`
	TechnologicalMaps   int `json:"technological_maps"`
	SubatomicComposites int `json:"subatomic_composites"`

	InMap         bool `json:"in_map"`
	BuildObjectID int  `json:"build_object_id"`
	DontRecycle   bool `json:"dont_recycle"`
	Tech          int  `json:"tech"`
}

func (b *Blueprint) GetAdvancedItems() []inventory.Slot {
	return b.AdvancedItems
}

func (b *Blueprint) GetAdvancedElectronics() int {
	return b.AdvancedElectronics
}

func (b *Blueprint) GetReinforcedFrame() int {
	return b.ReinforcedFrame
}

func (b *Blueprint) GetReinforcedDrives() int {
	return b.ReinforcedDrives
}

func (b *Blueprint) GetCompositeArmor() int {
	return b.CompositeArmor
}

func (b *Blueprint) GetCivilianWeapon() int {
	return b.CivilianWeapon
}

func (b *Blueprint) GetFoxCrystals() int {
	return b.FoxCrystals

}

func (b *Blueprint) GetNeutrinoProcessors() int {
	return b.NeutrinoProcessors

}

func (b *Blueprint) GetWeaponParts() int {
	return b.WeaponParts
}

func (b *Blueprint) GetTechnologicalMaps() int {
	return b.TechnologicalMaps
}

func (b *Blueprint) GetSubatomicComposites() int {
	return b.SubatomicComposites
}

func (b *Blueprint) GetName() string {
	return b.Name
}

func (b *Blueprint) GetSize() int {
	return 1
}

func (b *Blueprint) GetItemType() string {
	return b.ItemType
}

func (b *Blueprint) GetItemName() string {
	return b.ItemName
}

func (b *Blueprint) GetInMap() bool {
	return b.InMap
}

func (b *Blueprint) GetIcon() string {
	return b.Icon
}

func (b *Blueprint) GetStandardSize() int {
	return 0
}

func (b *Blueprint) GetTypeSlot() int {
	return 0
}

func (b *Blueprint) GetTech() int {
	return b.Tech
}

func (b *Blueprint) GetNeedCount() int {
	return 0
}

func (b *Blueprint) GetEnrichedThorium() int {
	return b.EnrichedThorium
}

func (b *Blueprint) GetIron() int {
	return b.Iron
}

func (b *Blueprint) GetCopper() int {
	return b.Copper
}

func (b *Blueprint) GetTitanium() int {
	return b.Titanium
}

func (b *Blueprint) GetSilicon() int {
	return b.Silicon
}

func (b *Blueprint) GetPlastic() int {
	return b.Plastic
}

func (b *Blueprint) GetCarbon() int {
	return b.Carbon
}

func (b *Blueprint) GetOrganic() int {
	return b.Organic
}

func (b *Blueprint) GetSpices() int {
	return b.Spices
}

func (b *Blueprint) GetDisputes1() int {
	return b.Disputes1
}

func (b *Blueprint) GetDisputes2() int {
	return b.Disputes2
}

func (b *Blueprint) GetResin() int {
	return b.Resin
}

func (b *Blueprint) GetFlower() int {
	return b.Flower
}

func (b *Blueprint) GetOil() int {
	return b.Oil
}

func (b *Blueprint) GetSteel() int {
	return b.Steel
}

func (b *Blueprint) GetWire() int {
	return b.Wire
}

func (b *Blueprint) GetElectronics() int {
	return b.Electronics
}

func (b *Blueprint) GetWires() int {
	return b.Wires
}

func (b *Blueprint) GetGear() int {
	return b.Gear
}

func (b *Blueprint) GetTitaniumPlate() int {
	return b.TitaniumPlate
}

func (b *Blueprint) GetBatteries() int {
	return b.Batteries
}

func (b *Blueprint) GetArmorItems() int {
	return b.ArmorItems
}

func (b *Blueprint) GetCarbonPlate() int {
	return b.CarbonPlate
}

func (b *Blueprint) GetRecyclingAlgorithm() string {
	return ""
}
