package behavior_rule

func GetDismantlingRules() (*BehaviorRules, *BehaviorRules) {
	var BehaviorDismantlingInRules = BehaviorRules{
		Rules: []*BehaviorRule{
			{
				Action: "simply_trade",
				PassRule: &BehaviorRule{
					Action: "process_cargo",
					PassRule: &BehaviorRule{
						Action: "drop_cargo_in_base",
						PassRule: &BehaviorRule{
							Action: "check_hp",
							Meta: &Meta{
								Parameter: "HP",
								Count:     100,
								Percent:   true,
							},
							PassRule: &BehaviorRule{
								Action: "out_base",
							},
							StopRule: &BehaviorRule{
								Action: "repair",
							},
						},
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

	var BehaviorDismantlingOutRules = BehaviorRules{
		Rules: []*BehaviorRule{
			{
				Action:   "find_hostile_in_range_view",
				PassRule: peacefulCheckBattleSolution(getBackRules3()),
				StopRule: &BehaviorRule{
					Action:   "check_back_to_base",
					PassRule: getBackRules3(),
					StopRule: &BehaviorRule{
						Action: "send_npc_request",
						Meta:   &Meta{Type: "defend"},
						PassRule: &BehaviorRule{
							Action: "to_wreckage",
							StopRule: &BehaviorRule{
								Action: "check_hp",
								Meta: &Meta{
									Parameter: "HP",
									Count:     90,
									Percent:   true,
								},
								PassRule: &BehaviorRule{
									Action: "find_random_sector",
									PassRule: &BehaviorRule{
										Action: "to_sector_target",
										Meta:   &Meta{Type: "Fraction"},
									},
									StopRule: getBackRules(),
								},
								StopRule: getBackRules(),
							},
						},
					},
				},
			},
		},
		Meta: &Meta{},
	}

	return &BehaviorDismantlingInRules, &BehaviorDismantlingOutRules
}
