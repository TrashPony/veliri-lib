package behavior_rule

// торговец ищет где купить подешевле и продать по дороже
func GetTransportRules() (*BehaviorRules, *BehaviorRules) {
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

	// проверяем насколько мы превосходим противника
	// если сильно то выбиваем из него деньги или трюм
	// если не очень сильно то пытаемся помирится
	// если мы проигрываем то просим помощи у других ботов
	// если у транспорт мало хп, то против защитников или воинов его сопроводить

	var BehaviorTransportOutRules = BehaviorRules{
		Rules: []*BehaviorRule{
			{
				Action:   "find_hostile_in_range_view",
				PassRule: peacefulCheckBattleSolution(),
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
							},
						},
					},
				},
			},
		},
		Meta: &Meta{},
	}

	return &BehaviorTransportInRules, &BehaviorTransportOutRules
}
