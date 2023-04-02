package _const

type rank struct {
	name        string
	lvl         []int
	RewardItems map[string]rewardItem
}

type rewardItem struct {
	t  string
	id int
	c  int
}

var ranks = map[int]rank{
	0: {name: "civil", lvl: []int{0}},

	1: {name: "volunteer", lvl: []int{1, 2}},
	2: {name: "rookie", lvl: []int{3, 4},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	3: {name: "soldier", lvl: []int{5, 6},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	4: {name: "legionary", lvl: []int{7, 8},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	5: {name: "veteran", lvl: []int{9, 10},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	6: {name: "corporal", lvl: []int{11, 12},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	7: {name: "sergeant", lvl: []int{13, 14},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	8: {name: "staff_sergeant", lvl: []int{15, 16},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	9: {name: "lieutenant", lvl: []int{17, 18},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	10: {name: "captain", lvl: []int{19, 20},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	11: {name: "major", lvl: []int{21, 22},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	12: {name: "centurion", lvl: []int{23, 24},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	13: {name: "colonel", lvl: []int{25, 26},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	14: {name: "tribune", lvl: []int{27, 28},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	15: {name: "brigadier_general", lvl: []int{29, 30},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	16: {name: "prefect", lvl: []int{31, 32},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	17: {name: "praetorian", lvl: []int{33, 34},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	18: {name: "palatine", lvl: []int{35, 36},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	19: {name: "honored_palatine", lvl: []int{37, 38},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	20: {name: "legate", lvl: []int{39, 40},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	21: {name: "general", lvl: []int{41, 42},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	22: {name: "warlord", lvl: []int{43, 44},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	23: {name: "supreme_warlord", lvl: []int{45, 46},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	24: {name: "commander", lvl: []int{47, 48},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
	25: {name: "commander_in_chief", lvl: []int{49, 50},
		RewardItems: map[string]rewardItem{
			Replicas: {t: "blueprints", id: 13, c: 3},
			Explores: {t: "blueprints", id: 13, c: 3},
			Reverses: {t: "blueprints", id: 13, c: 3},
		},
	},
}

func GetMaxRankLvl(rank int) int {
	r, ok := ranks[rank]
	if ok {
		m := 0

		for i, e := range r.lvl {
			if i == 0 || e > m {
				m = e
			}
		}

		return m
	}

	return -1
}

func CheckLvl(rank, lvl int) bool {
	r, ok := ranks[rank]
	if ok {
		for _, l := range r.lvl {
			if l == lvl {
				return true
			}
		}
	}

	return false
}

func GetRankName(rank int) string {
	r, ok := ranks[rank]
	if ok {
		return r.name
	}

	return ""
}

func GetRewardRank(rank int, fraction string) (string, int, int) {
	r, ok := ranks[rank]
	if ok {
		return r.RewardItems[fraction].t, r.RewardItems[fraction].id, r.RewardItems[fraction].c
	}

	return "", 0, 0
}
