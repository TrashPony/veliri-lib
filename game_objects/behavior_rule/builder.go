package behavior_rule

func GetBuilderRules() (*BehaviorRules, *BehaviorRules) {
	var BehaviorBuilderInRules = BehaviorRules{
		Rules: []*BehaviorRule{
			{
				Action: "check_hp",
				Meta: &Meta{
					Parameter: "HP",
					Count:     100,
					Percent:   true,
				},
				PassRule: &BehaviorRule{
					Action: "check_signal_about_building", // проверяем сигналы о запросе строительства,

					PassRule: &BehaviorRule{ //если есть то будет заполнена мета информация для пути
						Action: "get_resource",
						PassRule: &BehaviorRule{ // берем необходимые ресурсы для стрительства  (что брать должно быть заполненно в мета)
							Action: "out_base",
						},
					},
					StopRule: &BehaviorRule{
						Action: "find_not_complete_structure", // смотрим естли незаконченые здания у базы в округе (принадлежащие базе)
						PassRule: &BehaviorRule{
							Action: "out_base",
						},
					},
				},
				StopRule: &BehaviorRule{
					Action: "repair",
				},
			},
		},
		Meta: &Meta{},
	}

	// проверяем насколько мы превосходим противника
	// если сильно то выбиваем из него деньги или трюм
	// если не очень сильно то пытаемся помирится
	// если мы проигрываем то просим помощи у других ботов

	var BehaviorBuilderOutRules = BehaviorRules{
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
							Action: "find_not_complete_structure",
							PassRule: &BehaviorRule{
								Action: "build",
							},
							StopRule: &BehaviorRule{
								Action: "check_fraction_request",
								PassRule: &BehaviorRule{
									Action: "to_sector_target", // ищет путь до целли которая указана в мета (map_id, type_target, id_target)
									// мета информация должна быть заполнена уже т.к. задание выдает база
									Meta: &Meta{Type: "Fraction"},
									PassRule: &BehaviorRule{
										//ид, х, у - указано в мета
										Action: "place_object",
										StopRule: &BehaviorRule{
											Action: "global_target_scouting",
										},
									},
								},
								StopRule: getBackRules3(),
							},
						},
					},
				},
			},
		},
		Meta: &Meta{},
	}

	return &BehaviorBuilderInRules, &BehaviorBuilderOutRules
}
