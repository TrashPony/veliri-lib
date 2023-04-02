package _const

var FractionQuickBase = map[string]int{
	Replicas: ReplicasQuickBaseID,
	Explores: ExploresQuickBaseID,
	Reverses: ReversesQuickBaseID,
}

var QuickBattleInfinitelyAmmo = map[int]bool{
	5:  true,
	13: true,
	3:  true,
}

var SessionBaseNames = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}

type quickBattleOption struct {
	CountPlayers int  `json:"count_players"`
	Group        bool `json:"group"`
	Groups       int  `json:"groups"`
}

var QuickBattleOptions = map[string]quickBattleOption{
	"battle_":             {CountPlayers: 16, Group: true, Groups: 2},
	"battle_ai":           {CountPlayers: 1, Group: false, Groups: 1},
	"training_":           {CountPlayers: 1, Group: false, Groups: 1},
	"team_death_match_":   {CountPlayers: 10, Group: true, Groups: 2},
	"team_death_match_ai": {CountPlayers: 1, Group: false, Groups: 1},
	"destroy_base_":       {CountPlayers: 20, Group: true, Groups: 2},
	"destroy_base_ai":     {CountPlayers: 1, Group: false, Groups: 1},
	"defense_":            {CountPlayers: 2, Group: true, Groups: 1},
	"convoy_":             {CountPlayers: 2, Group: true, Groups: 1},
	"breakthrough_":       {CountPlayers: 2, Group: true, Groups: 1},
	"defense_single":      {CountPlayers: 1, Group: true, Groups: 1},
	"breakthrough_single": {CountPlayers: 1, Group: true, Groups: 1},
}
