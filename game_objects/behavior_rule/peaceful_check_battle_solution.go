package behavior_rule

func peacefulCheckBattleSolution() *BehaviorRule {
	return &BehaviorRule{
		Action: "peaceful_check_battle_solution",
		PassRule: &BehaviorRule{
			Action: "send_npc_request",
			Meta:   &Meta{Type: "attack"},
			PassRule: &BehaviorRule{
				Action: "follow_attack_target",
			},
		},
		StopRule: &BehaviorRule{
			Action: "check_weapon",
			PassRule: &BehaviorRule{
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
	}
}
