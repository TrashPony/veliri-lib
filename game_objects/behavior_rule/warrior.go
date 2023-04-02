package behavior_rule

func GetWarriorRules() (*BehaviorRules, *BehaviorRules) {
	// TODO эти правила не используются
	var BehaviorWarriorInRules = BehaviorRules{
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

	var BehaviorWarriorOutRules = BehaviorRules{
		Rules: []*BehaviorRule{
			{
				Action: "find_hostile_in_range_view", // главная задача воина - убивать, поэтому если мы видим что убивать, убиваем
				PassRule: &BehaviorRule{
					Action: "warrior_check_battle_solution",
					PassRule: &BehaviorRule{
						Action: "follow_attack_target",
					},
					StopRule: &BehaviorRule{
						Action: "check_hp",
						Meta: &Meta{
							Parameter: "HP",
							Count:     33,
							Percent:   true,
						},
						PassRule: &BehaviorRule{
							Action: "send_npc_request",
							Meta:   &Meta{Type: "attack"},
							PassRule: &BehaviorRule{
								Action:   "send_npc_request",
								Meta:     &Meta{Type: "defend"},
								PassRule: getBackRules(),
							},
						},
						StopRule: &BehaviorRule{
							Action:   "send_npc_request",
							Meta:     &Meta{Type: "defend"},
							PassRule: getBackRules(),
						},
					},
				},
				StopRule: &BehaviorRule{
					Action:   "check_back_to_base",
					PassRule: getBackRules(),
					StopRule: &BehaviorRule{
						Action: "check_in_hostile_sector", // нахожус ли я в секторе врага?
						PassRule: &BehaviorRule{
							Action: "capture_sector", // следуем на территорию захвата
							PassRule: &BehaviorRule{
								Action: "to_capture_base", // следуем на территорию захвата
							},
							StopRule: &BehaviorRule{
								Action: "to_capture_base", // следуем на территорию захвата
							},
						},
						StopRule: &BehaviorRule{
							Action: "check_hp",
							Meta: &Meta{
								Parameter: "HP",
								Count:     60,
								Percent:   true,
							},
							PassRule: &BehaviorRule{
								Action: "find_hostile_sector", // ищем приоритетный сектор
								PassRule: &BehaviorRule{
									Action: "to_sector_target", // следуем на территорию захвата
								},
								StopRule: &BehaviorRule{
									Action: "to_hostile_sector",
									StopRule: &BehaviorRule{ // если такого нет то мечемся в панике и умираем в рандомном секторе
										Action: "to_fraction_sector",
									},
								},
							},
							StopRule: getBackRules(),
						},
					},
				},
			},
		},
		Meta: &Meta{},
	}

	return &BehaviorWarriorInRules, &BehaviorWarriorOutRules
}
