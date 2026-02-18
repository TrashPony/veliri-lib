package behavior_rule

import "math/rand"

func GetRustBucketPirateRules() (*BehaviorRules, *BehaviorRules) {
	// каждый третий пират у нас торговец)
	if rand.Intn(3) == 0 {
		return GetRustBucketPirateTradeRules()
	} else {
		return GetInScoutRules()
	}
}

func GetRustBucketPirateTradeRules() (*BehaviorRules, *BehaviorRules) {
	var BehaviorTransportInRules = BehaviorRules{
		Rules: []*BehaviorRule{
			{
				// если у бота есть товары и их покупает база по актуальной цене то продаем их и покупаем новые если таковые есть
				Action: "trade",
				PassRule: &BehaviorRule{
					Action: "check_hp",
					Meta: &Meta{
						Parameter: "HP",
						Count:     100,
						Percent:   true,
					},
					PassRule: &BehaviorRule{
						// ищем где купить ресурсы или продать те которые уже есть, маршрут задается заранее то есть он проверяется и актуализируется только на базах
						// если на базе куда ехал бот товаров больше нет то просто ищем другие, а если товаров стало больше 1000 то ищем где их можно слить
						Action: "check_trade_flight",
						PassRule: &BehaviorRule{
							Action: "out_base",
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

	transportLogig := &BehaviorRule{
		Action:   "find_hostile_in_range_view",
		PassRule: peacefulCheckBattleSolution(getBackRules()),
		StopRule: &BehaviorRule{
			Action:   "check_back_to_base",
			PassRule: getBackRules(),
			StopRule: &BehaviorRule{
				Action: "send_npc_request",
				Meta:   &Meta{Type: "attack"},
				PassRule: &BehaviorRule{
					Action: "send_npc_request",
					Meta:   &Meta{Type: "defend"},
					PassRule: &BehaviorRule{
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
	}

	var BehaviorTransportOutRules = BehaviorRules{
		Rules: []*BehaviorRule{
			{
				Action: "find_hostile_in_range_view",
				PassRule: &BehaviorRule{
					Action: "warrior_check_battle_solution",
					Meta:   &Meta{Type: "secure", Count: 60, Count2: 25}, // Count - minHPPercentToBase, Count2 (minK) - соотношение сил, торговцы трусливее обычных пиратов
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
								Action:   "send_npc_request",
								Meta:     &Meta{Type: "defend"},
								PassRule: transportLogig, // если не вывозим то идем торговать и забиваем на врагов
							},
						},
						StopRule: &BehaviorRule{
							Action:   "send_npc_request",
							Meta:     &Meta{Type: "defend"},
							PassRule: transportLogig, // если не вывозим то идем торговать и забиваем на врагов
						},
					},
				},
				StopRule: &BehaviorRule{
					Action:   "send_npc_request", // запрос о грабеже, какомунить рандомному мирняке
					Meta:     &Meta{Type: "demand"},
					PassRule: transportLogig, // если не удалось идем торговать)
				},
			},
		},
		Meta: &Meta{},
	}

	return &BehaviorTransportInRules, &BehaviorTransportOutRules
}
