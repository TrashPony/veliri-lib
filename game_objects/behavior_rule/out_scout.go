package behavior_rule

/* Эта роль свободного агента который пермещается между секторами по секьюрным и свободным секторам,
если он видит слабого мирного то он не побрезгует его атаковать с требованием выкупа или груза */

/* Если он видит вокруг себя брошеный итем и он помещается то забирает его и сливает все на базе */

func GetOutScoutRules() (*BehaviorRules, *BehaviorRules) {
	var BehaviorOutScoutInRules = BehaviorRules{
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
								Action: "check_end_mission",
								PassRule: &BehaviorRule{
									Action: "get_random_mission",
									PassRule: &BehaviorRule{
										Action: "fill_mission_rules",
										PassRule: &BehaviorRule{
											Action: "out_base",
										},
									},
								},
								StopRule: &BehaviorRule{
									Action: "out_base",
								},
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

	var BehaviorOutScoutOutRules = BehaviorRules{
		Rules: []*BehaviorRule{
			{
				Action: "check_pk",
				PassRule: &BehaviorRule{
					Action: "to_nearest_free_land",
					PassRule: &BehaviorRule{
						Action: "find_hostile_in_range_view",
						PassRule: &BehaviorRule{
							Action: "warrior_check_battle_solution",
							Meta:   &Meta{Type: "secure", Count: 35},
							PassRule: &BehaviorRule{
								Action: "send_npc_request",
								Meta:   &Meta{Type: "demand"},
								PassRule: &BehaviorRule{
									Action: "follow_attack_target",
								},
							},
							StopRule: &BehaviorRule{
								Action: "check_hp",
								Meta: &Meta{
									Parameter: "HP",
									Count:     50,
									Percent:   true,
								},
								PassRule: &BehaviorRule{
									Action: "send_npc_request",
									Meta:   &Meta{Type: "attack"},
									PassRule: &BehaviorRule{
										Action: "follow_attack_target",
									},
								},
								StopRule: &BehaviorRule{
									Action: "send_npc_request",
									Meta:   &Meta{Type: "defend"},
									PassRule: &BehaviorRule{
										Action: "follow_attack_target",
									},
								},
							},
						},
						StopRule: &BehaviorRule{
							Action: "scouting",
						},
					},
					StopRule: &BehaviorRule{
						Action: "to_sector_target",
						Meta:   &Meta{Type: "Fraction"},
						PassRule: &BehaviorRule{
							Action: "to_base",
						},
						StopRule: getBackRules(),
					},
				},
				StopRule: &BehaviorRule{
					Action:   "check_special_mission",
					Meta:     &Meta{Type: "mission_rules"},
					PassRule: nil,
					StopRule: &BehaviorRule{
						Action: "find_hostile_in_range_view",
						PassRule: &BehaviorRule{
							Action: "warrior_check_battle_solution",
							Meta:   &Meta{Type: "secure", Count: 35},
							PassRule: &BehaviorRule{
								Action: "send_npc_request",
								Meta:   &Meta{Type: "attack"},
								PassRule: &BehaviorRule{
									Action: "send_npc_request",
									Meta:   &Meta{Type: "demand"},
									PassRule: &BehaviorRule{
										Action: "follow_attack_target",
									},
								},
							},
							StopRule: &BehaviorRule{
								Action: "check_hp",
								Meta: &Meta{
									Parameter: "HP",
									Count:     50,
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
									Action: "send_npc_request",
									Meta:   &Meta{Type: "attack"},
									PassRule: &BehaviorRule{
										Action:   "send_npc_request",
										Meta:     &Meta{Type: "defend"},
										PassRule: getBackRules(),
									},
								},
							},
						},
						StopRule: &BehaviorRule{
							Action:   "check_cargo_full", // проверка трюма, если он заполнен на 50% то ливаем на базу
							Meta:     &Meta{Count: 80},
							PassRule: getBackRules(),
							StopRule: &BehaviorRule{
								Action: "find_unit_wreckage",
								PassRule: &BehaviorRule{
									Action: "pick_up_unit_wreckage",
								},
								StopRule: &BehaviorRule{
									Action: "find_drop_items", // смотри брошеные вещи которые можно поднять, если такие есть берем любой и кладем в мету
									PassRule: &BehaviorRule{
										Action: "get_drop_item", // идем к упавшему предмету и пытаемся его взять, ид предмета указать в мете
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

	return &BehaviorOutScoutInRules, &BehaviorOutScoutOutRules
}
