package behavior_rule

func GetMinerRules() (*BehaviorRules, *BehaviorRules) {
	var BehaviorMinerInRules = BehaviorRules{
		Rules: []*BehaviorRule{
			{
				Action: "simply_trade",
				PassRule: &BehaviorRule{
					Action: "check_hp",
					Meta: &Meta{
						Parameter: "HP",
						Count:     100,
						Percent:   true,
					},
					PassRule: &BehaviorRule{
						Action: "check_cargo_empty",
						PassRule: &BehaviorRule{
							Action: "out_base",
							Exit:   true,
						},
						StopRule: &BehaviorRule{
							Action: "process_cargo",
							Meta:   &Meta{Type: "processing_plant"},
							PassRule: &BehaviorRule{
								Action: "drop_cargo_in_base",
								PassRule: &BehaviorRule{
									Action: "out_base",
								},
								Exit: true,
							},
							StopRule: &BehaviorRule{
								Action: "out_base",
							},
						},
					},
					StopRule: &BehaviorRule{
						Action: "repair",
					},
				},
			},
		},
		Meta: &Meta{},
	}

	// проверяем насколько мы превосходим противника
	// если сильно то выбиваем из него деньги или трюм
	// если не очень сильно то пытаемся помирится
	// если мы проигрываем то просим помощи у других ботов

	var BehaviorMinerOutRules = BehaviorRules{
		Rules: []*BehaviorRule{
			{
				Action:   "find_hostile_in_range_view",
				PassRule: peacefulCheckBattleSolution(),
				StopRule: &BehaviorRule{
					Action: "send_npc_request",
					Meta:   &Meta{Type: "attack"},
					PassRule: &BehaviorRule{
						Action:   "check_back_to_base",
						PassRule: getBackRules(),
						StopRule: &BehaviorRule{
							Action: "check_cargo_full",
							Meta:   &Meta{Count: 75},
							PassRule: &BehaviorRule{
								Action: "find_fraction_processing_plant",
								PassRule: &BehaviorRule{
									Action: "to_sector_target",
									Meta:   &Meta{Type: "Fraction"},
									PassRule: &BehaviorRule{
										Action: "to_base",
									},
									StopRule: getBackRules(),
								},
							},
							StopRule: &BehaviorRule{
								Action: "find_drop_items", // смотри брошеные вещи которые можно поднять, если такие есть берем любой и кладем в мету
								PassRule: &BehaviorRule{
									Action: "get_drop_item", // идем к упавшему предмету и пытаемся его взять, ид предмета указать в мете
								},
								StopRule: &BehaviorRule{
									Action: "to_reservoir_mine",
									StopRule: &BehaviorRule{
										Action: "check_hp",
										Meta: &Meta{
											Parameter: "HP",
											Count:     90,
											Percent:   true,
										},
										PassRule: &BehaviorRule{
											Action: "to_fraction_sector",
											Meta:   &Meta{SetBackToBaseWhenChangeSector: true},
											StopRule: &BehaviorRule{
												Action: "find_random_sector",
												PassRule: &BehaviorRule{
													Action: "to_sector_target",
													Meta:   &Meta{Type: "Fraction"},
												},
												StopRule: getBackRules(),
											},
										},
										StopRule: getBackRules(),
									},
								},
							},
						},
					},
				},
			},
		},
		Meta: &Meta{},
	}
	return &BehaviorMinerInRules, &BehaviorMinerOutRules
}
