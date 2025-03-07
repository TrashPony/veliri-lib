package relation_bonuses

type RelationBonuses struct {
	FractionDroneView        bool `json:"fraction_drone_view"`
	AttackFractionWarrior    bool `json:"attack_fraction_warrior"`
	Evacuation               bool `json:"evacuation"`
	FractionWarriorView      bool `json:"fraction_warrior_view"`
	ReducedRecycleTax        int  `json:"reduced_recycle_tax"`
	ReducedWorkTax           int  `json:"reduced_work_tax"`
	ReducedDetailWorkTax     int  `json:"reduced_detail_work_tax"`
	ReducedMarketTax         int  `json:"reduced_market_tax"`
	AdditionalMissionCredits int  `json:"additional_mission_credits"`

	// negative
	LockedUpNPCTrade bool `json:"locked_up_npc_trade"`
	LockedUpWork     bool `json:"locked_up_work"`
	LockedUpRecycle  bool `json:"locked_up_recycle"`
	FractionHostile  bool `json:"fraction_hostile"`
}
