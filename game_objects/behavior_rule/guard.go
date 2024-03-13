package behavior_rule

func GetGuardRules() (*BehaviorRules, *BehaviorRules) {
	var BehaviorGuardInRules = BehaviorRules{
		Rules: []*BehaviorRule{
			{
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
		Meta: &Meta{},
	}

	// проверяем насколько мы превосходим противника
	// если мы проигрываем то просим помощи у других ботов

	var BehaviorGuardOutRules = BehaviorRules{
		Rules: []*BehaviorRule{
			{
				Action: "police_logic", // ищем нарушителей в нашей зоне видимости
				PassRule: &BehaviorRule{
					Action: "find_fraction_hostile",
					PassRule: &BehaviorRule{
						Action: "find_hostile_in_range_view",
						PassRule: &BehaviorRule{
							Action: "warrior_check_battle_solution",
							PassRule: &BehaviorRule{
								Action: "send_npc_request",
								Meta:   &Meta{Type: "attack"},
								PassRule: &BehaviorRule{
									Action: "follow_attack_target",
								},
							},
							StopRule: &BehaviorRule{
								Action: "send_npc_request",
								Meta:   &Meta{Type: "attack"},
								PassRule: &BehaviorRule{
									Action:   "send_npc_request",
									Meta:     &Meta{Type: "defend"},
									PassRule: getBackRules(),
								},
							},
						},
						StopRule: &BehaviorRule{
							Action: "back_to_parent_sector", // ищем дорогу домой
							PassRule: &BehaviorRule{
								Action: "to_sector_target",
								Meta:   &Meta{Type: "Fraction"},
								PassRule: &BehaviorRule{
									Action: "to_base",
								},
								StopRule: getBackRules(),
							},
							StopRule: &BehaviorRule{
								Action:   "check_back_to_base",
								PassRule: getBackRules(),
								StopRule: &BehaviorRule{
									Action: "check_hp",
									Meta: &Meta{
										Parameter: "HP",
										Count:     90,
										Percent:   true,
									},
									PassRule: &BehaviorRule{
										Action: "check_distress_signals",
										StopRule: &BehaviorRule{
											Action: "scouting",
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
		Meta: &Meta{},
	}

	return &BehaviorGuardInRules, &BehaviorGuardOutRules
}
