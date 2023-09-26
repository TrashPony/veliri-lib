package behavior_rule

func getBackRules() *BehaviorRule {
	return &BehaviorRule{
		Action: "check_in_hostile_sector",
		Meta:   &Meta{Parameter: "BackRule"},
		PassRule: &BehaviorRule{
			Action: "to_fraction_sector",
			StopRule: &BehaviorRule{
				Action: "find_random_sector",
				Meta:   &Meta{Parameter: "BackRule"},
				PassRule: &BehaviorRule{
					Action: "to_sector_target",
					Meta:   &Meta{Type: "Fraction"},
				},
			},
		},
		StopRule: &BehaviorRule{
			Action: "to_fraction_base",
			Meta:   &Meta{Parameter: "BackRule"},
		},
	}
}
