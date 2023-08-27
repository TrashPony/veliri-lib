package _const

import (
	"math/rand"
	"sort"
)

const (
	FreeLandMaxContAnomalies  = 20
	BattleMapMaxContAnomalies = 15
	SecureMapMaxContAnomalies = 10
)

var DetailsAnomalyOptions = map[int]*detailsAnomalyOption{
	1: {MaxCountDetails: 2, MaxDetails: 15, MaxCountRecycle: 2, MaxRecycle: 250},
	2: {MaxCountDetails: 3, MaxDetails: 20, MaxCountRecycle: 4, MaxRecycle: 500},
	3: {MaxCountDetails: 4, MaxDetails: 25, MaxCountRecycle: 6, MaxRecycle: 750},
}

type detailsAnomalyOption struct {
	MaxCountDetails int `json:"max_count_details"`
	MaxCountRecycle int `json:"max_count_recycle"`
	MaxDetails      int `json:"max_details"`
	MaxRecycle      int `json:"max_recycle"`
}

// алгоритм примерно такой,
// проходимся по всем бп и оставляем только те которые выпали
// если количество привышает показатель MaxCountBP то оставляем самые редки а самые частые выкидываем
// TODO написать тест который будет проверять что все блюпринты распихнуты по аномалиями что бы нечего не пропустить
// TODO написать тест с созданием и открыванием ящика что бы понимать как работает конфигурация

var BlueprintAnomalyOptions = map[int]*bluePrintsAnomalyOption{
	1: {
		Bps: []BluePrintAnomalyOption{
			// AMMO +
			{BpID: 74, MaxCount: 30, PercentChance: 20},
			{BpID: 25, MaxCount: 30, PercentChance: 20},
			{BpID: 161, MaxCount: 30, PercentChance: 20},
			{BpID: 163, MaxCount: 30, PercentChance: 20},

			// BOXES +
			{BpID: 164, MaxCount: 10, PercentChance: 30},
			{BpID: 27, MaxCount: 5, PercentChance: 20},
			{BpID: 28, MaxCount: 2, PercentChance: 10},

			// DETAIL +
			{BpID: 112, MaxCount: 5, PercentChance: 25},
			{BpID: 113, MaxCount: 5, PercentChance: 25},
			{BpID: 115, MaxCount: 5, PercentChance: 25},
			{BpID: 116, MaxCount: 5, PercentChance: 25},
			{BpID: 118, MaxCount: 5, PercentChance: 25},

			// WEAPON lvl 1 +
			{BpID: 78, MaxCount: 2, PercentChance: 10, SectorFraction: Replicas},
			{BpID: 80, MaxCount: 2, PercentChance: 10, SectorFraction: Replicas},
			{BpID: 82, MaxCount: 2, PercentChance: 10, SectorFraction: Explores},
			{BpID: 83, MaxCount: 2, PercentChance: 10, SectorFraction: Explores},
			{BpID: 84, MaxCount: 2, PercentChance: 10, SectorFraction: Explores},
			{BpID: 162, MaxCount: 2, PercentChance: 10, SectorFraction: Reverses},
			{BpID: 95, MaxCount: 2, PercentChance: 10},

			// BODY lvl 1 +
			{BpID: 119, MaxCount: 2, PercentChance: 5, SectorFraction: Replicas},
			{BpID: 50, MaxCount: 2, PercentChance: 5, SectorFraction: Replicas},
			{BpID: 110, MaxCount: 2, PercentChance: 5, SectorFraction: Explores},
			{BpID: 55, MaxCount: 2, PercentChance: 5, SectorFraction: Explores},
			{BpID: 124, MaxCount: 2, PercentChance: 5, SectorFraction: Reverses},
			{BpID: 60, MaxCount: 2, PercentChance: 5, SectorFraction: Reverses},

			// EQUIP lvl 1 +
			{BpID: 103, MaxCount: 2, PercentChance: 15, SectorFraction: Replicas},
			{BpID: 104, MaxCount: 2, PercentChance: 15, SectorFraction: Explores},
			{BpID: 105, MaxCount: 2, PercentChance: 15, SectorFraction: Reverses},
			{BpID: 13, MaxCount: 2, PercentChance: 15},
			{BpID: 14, MaxCount: 2, PercentChance: 15},
			{BpID: 63, MaxCount: 2, PercentChance: 15},
			{BpID: 64, MaxCount: 2, PercentChance: 15},
			{BpID: 65, MaxCount: 2, PercentChance: 15},
			{BpID: 19, MaxCount: 2, PercentChance: 15},
			{BpID: 20, MaxCount: 2, PercentChance: 15},
			{BpID: 127, MaxCount: 2, PercentChance: 15},
			{BpID: 129, MaxCount: 2, PercentChance: 15},
			{BpID: 131, MaxCount: 2, PercentChance: 15},
			{BpID: 133, MaxCount: 2, PercentChance: 15},
			{BpID: 134, MaxCount: 2, PercentChance: 15},
			{BpID: 17, MaxCount: 2, PercentChance: 15},
			{BpID: 144, MaxCount: 2, PercentChance: 15},
			{BpID: 145, MaxCount: 2, PercentChance: 15},
			{BpID: 146, MaxCount: 2, PercentChance: 15},
			{BpID: 147, MaxCount: 2, PercentChance: 15},
			{BpID: 148, MaxCount: 2, PercentChance: 15},
			{BpID: 149, MaxCount: 2, PercentChance: 15},
			{BpID: 150, MaxCount: 2, PercentChance: 15},
			{BpID: 151, MaxCount: 2, PercentChance: 15},
			{BpID: 152, MaxCount: 2, PercentChance: 15},
			{BpID: 153, MaxCount: 2, PercentChance: 15},
			{BpID: 154, MaxCount: 2, PercentChance: 15},
			{BpID: 155, MaxCount: 2, PercentChance: 15},
			{BpID: 156, MaxCount: 2, PercentChance: 15},
			{BpID: 157, MaxCount: 2, PercentChance: 15},
			{BpID: 160, MaxCount: 2, PercentChance: 15},
		},
		MaxCountBP: 5, // сколько в коробе может быть разных БП
	},
	2: {
		Bps: []BluePrintAnomalyOption{
			// AMMO +
			{BpID: 23, MaxCount: 10, PercentChance: 20},
			{BpID: 75, MaxCount: 10, PercentChance: 20},
			{BpID: 21, MaxCount: 10, PercentChance: 20},
			{BpID: 26, MaxCount: 10, PercentChance: 20},

			// BOXES +
			{BpID: 164, MaxCount: 10, PercentChance: 30},
			{BpID: 27, MaxCount: 5, PercentChance: 20},
			{BpID: 28, MaxCount: 2, PercentChance: 10},

			// DETAIL +
			{BpID: 111, MaxCount: 5, PercentChance: 25},
			{BpID: 114, MaxCount: 5, PercentChance: 25},
			{BpID: 117, MaxCount: 5, PercentChance: 25},

			// WEAPON LVL 2 +
			{BpID: 79, MaxCount: 2, PercentChance: 10, SectorFraction: Replicas},
			{BpID: 81, MaxCount: 2, PercentChance: 10, SectorFraction: Explores},
			{BpID: 86, MaxCount: 2, PercentChance: 10, SectorFraction: Reverses},
			{BpID: 87, MaxCount: 2, PercentChance: 10, SectorFraction: Reverses},
			{BpID: 96, MaxCount: 2, PercentChance: 10},
			{BpID: 97, MaxCount: 2, PercentChance: 10},
			{BpID: 93, MaxCount: 2, PercentChance: 10},

			// BODY lvl 2 +
			{BpID: 120, MaxCount: 2, PercentChance: 5, SectorFraction: Replicas},
			{BpID: 100, MaxCount: 2, PercentChance: 5, SectorFraction: Replicas},
			{BpID: 51, MaxCount: 2, PercentChance: 5, SectorFraction: Replicas},
			{BpID: 122, MaxCount: 2, PercentChance: 5, SectorFraction: Explores},
			{BpID: 101, MaxCount: 2, PercentChance: 5, SectorFraction: Explores},
			{BpID: 56, MaxCount: 2, PercentChance: 5, SectorFraction: Explores},
			{BpID: 125, MaxCount: 2, PercentChance: 5, SectorFraction: Reverses},
			{BpID: 102, MaxCount: 2, PercentChance: 5, SectorFraction: Reverses},
			{BpID: 61, MaxCount: 2, PercentChance: 5, SectorFraction: Reverses},

			// EQUIP lvl 2 +
			{BpID: 69, MaxCount: 2, PercentChance: 15, SectorFraction: Replicas},
			{BpID: 71, MaxCount: 2, PercentChance: 15, SectorFraction: Explores},
			{BpID: 73, MaxCount: 2, PercentChance: 15, SectorFraction: Reverses},
			{BpID: 66, MaxCount: 2, PercentChance: 15},
			{BpID: 128, MaxCount: 2, PercentChance: 15},
			{BpID: 130, MaxCount: 2, PercentChance: 15},
			{BpID: 132, MaxCount: 2, PercentChance: 15},
			{BpID: 135, MaxCount: 2, PercentChance: 15},
			{BpID: 136, MaxCount: 2, PercentChance: 15},
			{BpID: 137, MaxCount: 2, PercentChance: 15},
			{BpID: 138, MaxCount: 2, PercentChance: 15},
			{BpID: 139, MaxCount: 2, PercentChance: 15},
			{BpID: 140, MaxCount: 2, PercentChance: 15},
			{BpID: 141, MaxCount: 2, PercentChance: 15},
			{BpID: 142, MaxCount: 2, PercentChance: 15},
			{BpID: 143, MaxCount: 2, PercentChance: 15},
			{BpID: 158, MaxCount: 2, PercentChance: 15},
			{BpID: 159, MaxCount: 2, PercentChance: 15},
		},
		MaxCountBP: 3, // сколько в коробе может быть разных БП
	},
	3: {
		Bps: []BluePrintAnomalyOption{
			// AMMO 3 +
			{BpID: 99999, MaxCount: 10, PercentChance: 20},
			{BpID: 99999, MaxCount: 10, PercentChance: 20},
			{BpID: 99999, MaxCount: 10, PercentChance: 20},

			// WEAPON lvl 3 +
			{BpID: 77, MaxCount: 2, PercentChance: 15},
			{BpID: 85, MaxCount: 2, PercentChance: 15},
			{BpID: 88, MaxCount: 2, PercentChance: 15},
			{BpID: 92, MaxCount: 2, PercentChance: 15},
			{BpID: 94, MaxCount: 2, PercentChance: 15},

			// BODY lvl 3 +
			{BpID: 52, MaxCount: 1, PercentChance: 5},
			{BpID: 48, MaxCount: 1, PercentChance: 5},
			{BpID: 121, MaxCount: 1, PercentChance: 5},
			{BpID: 57, MaxCount: 1, PercentChance: 5},
			{BpID: 53, MaxCount: 1, PercentChance: 5},
			{BpID: 123, MaxCount: 1, PercentChance: 5},
			{BpID: 62, MaxCount: 1, PercentChance: 5},
			{BpID: 58, MaxCount: 1, PercentChance: 5},
			{BpID: 126, MaxCount: 1, PercentChance: 5},
		},
		MaxCountBP: 2, // сколько в коробе может быть разных БП
	},
}

type bluePrintsAnomalyOption struct {
	Bps        []BluePrintAnomalyOption `json:"bps"`
	MaxCountBP int                      `json:"max_count_bp"`
}

type BluePrintAnomalyOption struct {
	BpID          int `json:"bp_id"`
	MaxCount      int `json:"max_count"`
	PercentChance int `json:"chance"`

	// только для безопасной зоны, если фракция сектора не совпадает с этой опцией то опция игнорируется
	SectorFraction string `json:"sector_fraction"`
}

var AmmoAnomalyOptions = map[int]*ammoAnomalyOptions{
	1: {
		Ammo: []AmmoAnomalyOption{
			{ID: 5, MaxCount: 8 * 10, PercentChance: 20},
			{ID: 13, MaxCount: 20 * 10, PercentChance: 20},
			{ID: 18, MaxCount: 100 * 10, PercentChance: 20},
			{ID: 33, MaxCount: 10 * 10, PercentChance: 20},

			{ID: 1, MaxCount: 4 * 10, PercentChance: 7},
			{ID: 3, MaxCount: 8 * 10, PercentChance: 7},
			{ID: 6, MaxCount: 4 * 10, PercentChance: 7},
			{ID: 14, MaxCount: 4 * 10, PercentChance: 7},

			{ID: 2, MaxCount: 2 * 10, PercentChance: 3},
			{ID: 4, MaxCount: 2 * 10, PercentChance: 3},
			{ID: 15, MaxCount: 2 * 10, PercentChance: 3},
		},
	},
	2: {
		Ammo: []AmmoAnomalyOption{
			{ID: 5, MaxCount: 8 * 10, PercentChance: 5},
			{ID: 13, MaxCount: 20 * 10, PercentChance: 5},
			{ID: 18, MaxCount: 100 * 10, PercentChance: 5},
			{ID: 33, MaxCount: 10 * 10, PercentChance: 5},

			{ID: 1, MaxCount: 4 * 10, PercentChance: 20},
			{ID: 3, MaxCount: 8 * 10, PercentChance: 20},
			{ID: 6, MaxCount: 4 * 10, PercentChance: 20},
			{ID: 14, MaxCount: 4 * 10, PercentChance: 20},

			{ID: 2, MaxCount: 2 * 10, PercentChance: 5},
			{ID: 4, MaxCount: 2 * 10, PercentChance: 5},
			{ID: 15, MaxCount: 2 * 10, PercentChance: 5},
		},
	},
	3: {
		Ammo: []AmmoAnomalyOption{
			{ID: 5, MaxCount: 8 * 10, PercentChance: 5},
			{ID: 13, MaxCount: 20 * 10, PercentChance: 5},
			{ID: 18, MaxCount: 100 * 10, PercentChance: 5},
			{ID: 33, MaxCount: 10 * 10, PercentChance: 5},

			{ID: 2, MaxCount: 2 * 10, PercentChance: 20},
			{ID: 4, MaxCount: 2 * 10, PercentChance: 20},
			{ID: 15, MaxCount: 2 * 10, PercentChance: 20},

			{ID: 2, MaxCount: 2 * 10, PercentChance: 5},
			{ID: 4, MaxCount: 2 * 10, PercentChance: 5},
			{ID: 15, MaxCount: 2 * 10, PercentChance: 5},
		},
	},
}

type ammoAnomalyOptions struct {
	Ammo []AmmoAnomalyOption `json:"ammo"`
}

type AmmoAnomalyOption struct {
	ID            int `json:"bp_id"`
	MaxCount      int `json:"max_count"`
	PercentChance int `json:"chance"`
}

func GetLoopBlueprint(lvl int, mpFraction string, freeLand bool) map[int]int { // [id] count

	result := make(map[int]int)

	option := BlueprintAnomalyOptions[lvl]
	if option == nil {
		return result
	}

	// высчитаем все чертежи которые выпадают игроку
	passBlueprints := make([]BluePrintAnomalyOption, 0)
	for _, bpOption := range option.Bps {
		if bpOption.PercentChance >= rand.Intn(100) && (bpOption.SectorFraction == mpFraction || freeLand) {
			passBlueprints = append(passBlueprints, bpOption)
		}
	}

	// сортируем массив, что бы самые редкие были сверху
	sort.SliceStable(passBlueprints, func(i, j int) bool {
		return passBlueprints[i].PercentChance < passBlueprints[j].PercentChance
	})

	for i := 0; i < option.MaxCountBP; i++ {
		if len(passBlueprints) > i {
			result[passBlueprints[i].BpID] = rand.Intn(option.MaxCountBP) + 1
		}
	}

	return result
}
