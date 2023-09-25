package behavior_rule

/* Роль пирата, вся его деятельность сводится к  атаке на мирные, агенский, игроков транспорты с требованиями выкинуть трюм или деньги */

/* Когда у пирата враждебные отношения с защитниками сектора он идет в другой сектор, где еше все хорошо, он должен меть журнал посещения
что бы не ходить в одни и теже сектора по кругу, может обитать в спорных и безопасных территориях*/

/* Если он видит вокруг себя брошеный итем и он помещается то забирает его и сливает все на базе */

func GetInScoutRules() (*BehaviorRules, *BehaviorRules) {
	var BehaviorInScoutInRules = BehaviorRules{
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

	var BehaviorInScoutOutRules = BehaviorRules{
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
						PassRule: &BehaviorRule{
							Action: "find_hostile_in_range_view",
							PassRule: &BehaviorRule{
								Action: "to_sector_target",
								PassRule: &BehaviorRule{
									Action: "to_base",
								},
								Meta:     &Meta{Type: "Fraction"},
								StopRule: getBackRules(),
							},
							StopRule: &BehaviorRule{
								Action: "to_sector_target",
								PassRule: &BehaviorRule{
									Action: "to_base",
								},
								Meta:     &Meta{Type: "Fraction"},
								StopRule: getBackRules(),
							},
						},
					},
				},
				StopRule: &BehaviorRule{
					Action: "find_hostile_in_range_view",
					PassRule: &BehaviorRule{
						Action: "warrior_check_battle_solution",
						Meta:   &Meta{Type: "secure", Count: 50, Count2: -15},
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
									Action: "send_npc_request",
									Meta:   &Meta{Type: "defend"},
									PassRule: &BehaviorRule{
										Action:   "check_profitability_sector",
										PassRule: getBackRules(),
										StopRule: &BehaviorRule{
											Action: "to_sector_target",
											Meta:   &Meta{Type: "Fraction"},
											PassRule: &BehaviorRule{
												Action: "to_base",
											},
											StopRule: getBackRules(),
										},
									},
								},
							},
							StopRule: &BehaviorRule{
								Action: "send_npc_request",
								Meta:   &Meta{Type: "defend"},
								PassRule: &BehaviorRule{
									Action:   "check_profitability_sector",
									PassRule: getBackRules(),
									StopRule: &BehaviorRule{
										Action: "to_sector_target",
										Meta:   &Meta{Type: "Fraction"},
										PassRule: &BehaviorRule{
											Action: "to_base",
										},
										StopRule: getBackRules(),
									},
								},
							},
						},
					},
					StopRule: &BehaviorRule{
						Action:   "check_back_to_base",
						PassRule: getBackRules(),
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
									StopRule: &BehaviorRule{
										Action: "check_hp",
										Meta: &Meta{
											Parameter: "HP",
											Count:     90,
											Percent:   true,
										},
										PassRule: &BehaviorRule{
											Action: "send_npc_request", // запрос о грабеже, какомунить рандомному мирняке
											Meta:   &Meta{Type: "demand"},
											PassRule: &BehaviorRule{
												// проверяем рентабельность сектора, если нет то выбираем любоей и следуем туда
												Action: "check_profitability_sector",
												PassRule: &BehaviorRule{
													Action: "fixed_to_current_sector",
													PassRule: &BehaviorRule{
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

	return &BehaviorInScoutInRules, &BehaviorInScoutOutRules
}
