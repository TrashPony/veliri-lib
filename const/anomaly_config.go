package _const

const (
	FreeLandMaxContAnomalies  = 20
	BattleMapMaxContAnomalies = 15
	SecureMapMaxContAnomalies = 10
)

var DetailsAnomalyOptions = map[int]*detailsAnomalyOption{
	1: {TechDetail: 1, MaxCountDetails: 2, MaxDetails: 15, MaxCountRecycle: 2, MaxRecycle: 250},
	2: {TechDetail: 1, MaxCountDetails: 3, MaxDetails: 20, MaxCountRecycle: 4, MaxRecycle: 500},
	3: {TechDetail: 2, MaxCountDetails: 4, MaxDetails: 25, MaxCountRecycle: 6, MaxRecycle: 750},
}

type detailsAnomalyOption struct {
	MaxCountDetails int `json:"max_count_details"`
	MaxCountRecycle int `json:"max_count_recycle"`
	MaxDetails      int `json:"max_details"`
	TechDetail      int `json:"tech_detail"`
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
			// AMMO LVL 1
			{BpID: 25, MaxCount: 20, PercentChance: 25},
			{BpID: 74, MaxCount: 20, PercentChance: 25},
			{BpID: 21, MaxCount: 10, PercentChance: 15},
			{BpID: 23, MaxCount: 10, PercentChance: 15},
			{BpID: 75, MaxCount: 5, PercentChance: 5},

			// BOXES LVL 1
			{BpID: 27, MaxCount: 5, PercentChance: 10},

			// DETAIL lvl 1
			{BpID: 30, MaxCount: 15, PercentChance: 25},
			{BpID: 31, MaxCount: 15, PercentChance: 25},
			{BpID: 33, MaxCount: 15, PercentChance: 25},
			{BpID: 34, MaxCount: 15, PercentChance: 25},

			// EQUIP lvl 1
			{BpID: 13, MaxCount: 3, PercentChance: 15},
			{BpID: 17, MaxCount: 3, PercentChance: 15},
			{BpID: 19, MaxCount: 3, PercentChance: 15},
			{BpID: 20, MaxCount: 3, PercentChance: 15},
			{BpID: 66, MaxCount: 3, PercentChance: 15},
			{BpID: 103, MaxCount: 3, PercentChance: 15, SectorFraction: Replicas},
			{BpID: 69, MaxCount: 3, PercentChance: 15, SectorFraction: Replicas},
			{BpID: 104, MaxCount: 3, PercentChance: 15, SectorFraction: Explores},
			{BpID: 71, MaxCount: 3, PercentChance: 15, SectorFraction: Explores},
			{BpID: 105, MaxCount: 3, PercentChance: 15, SectorFraction: Reverses},
			{BpID: 73, MaxCount: 3, PercentChance: 15, SectorFraction: Reverses},

			// WEAPON lvl 1
			{BpID: 78, MaxCount: 3, PercentChance: 15, SectorFraction: Replicas},
			{BpID: 80, MaxCount: 3, PercentChance: 15, SectorFraction: Replicas},
			{BpID: 82, MaxCount: 3, PercentChance: 15, SectorFraction: Explores},
			{BpID: 83, MaxCount: 3, PercentChance: 15, SectorFraction: Explores},
			{BpID: 86, MaxCount: 3, PercentChance: 15, SectorFraction: Reverses},
			{BpID: 87, MaxCount: 3, PercentChance: 15, SectorFraction: Reverses},
			{BpID: 95, MaxCount: 3, PercentChance: 15},
			{BpID: 96, MaxCount: 3, PercentChance: 15},
			{BpID: 97, MaxCount: 3, PercentChance: 15},

			// BODY lvl 1
			{BpID: 48, MaxCount: 2, PercentChance: 10, SectorFraction: Replicas},
			{BpID: 100, MaxCount: 2, PercentChance: 10, SectorFraction: Replicas},
			{BpID: 53, MaxCount: 2, PercentChance: 10, SectorFraction: Explores},
			{BpID: 101, MaxCount: 2, PercentChance: 10, SectorFraction: Explores},
			{BpID: 58, MaxCount: 2, PercentChance: 10, SectorFraction: Reverses},
			{BpID: 102, MaxCount: 2, PercentChance: 10, SectorFraction: Reverses},
			{BpID: 89, MaxCount: 3, PercentChance: 15},
			{BpID: 90, MaxCount: 3, PercentChance: 15},
		},
		MaxCountBP: 5, // сколько в коробе может быть разных БП
	},
	2: {
		Bps: []BluePrintAnomalyOption{
			// AMMO LVL 2
			{BpID: 22, MaxCount: 5, PercentChance: 10},
			{BpID: 24, MaxCount: 5, PercentChance: 10},
			{BpID: 76, MaxCount: 5, PercentChance: 10},
			{BpID: 26, MaxCount: 5, PercentChance: 10},

			// BOXES LVL 2
			{BpID: 28, MaxCount: 2, PercentChance: 5},

			// DETAIL lvl 2
			{BpID: 29, MaxCount: 10, PercentChance: 25},
			{BpID: 32, MaxCount: 10, PercentChance: 25},
			{BpID: 35, MaxCount: 10, PercentChance: 25},
			{BpID: 36, MaxCount: 10, PercentChance: 25},

			// EQUIP lvl 2
			{BpID: 14, MaxCount: 3, PercentChance: 15},
			{BpID: 63, MaxCount: 3, PercentChance: 5},
			{BpID: 64, MaxCount: 3, PercentChance: 5},
			{BpID: 65, MaxCount: 3, PercentChance: 5},
			{BpID: 67, MaxCount: 3, PercentChance: 5},
			{BpID: 108, MaxCount: 3, PercentChance: 5},
			{BpID: 109, MaxCount: 3, PercentChance: 5},

			// BUILD OBJECTS
			{BpID: 37, MaxCount: 3, PercentChance: 5},
			{BpID: 38, MaxCount: 3, PercentChance: 5},
			{BpID: 39, MaxCount: 3, PercentChance: 5},
			{BpID: 40, MaxCount: 3, PercentChance: 5},
			{BpID: 41, MaxCount: 3, PercentChance: 10},
			{BpID: 42, MaxCount: 3, PercentChance: 10},
			{BpID: 43, MaxCount: 3, PercentChance: 10},
			{BpID: 45, MaxCount: 3, PercentChance: 10},
			{BpID: 46, MaxCount: 3, PercentChance: 25},
			{BpID: 47, MaxCount: 3, PercentChance: 5},
			{BpID: 98, MaxCount: 3, PercentChance: 10},
			{BpID: 99, MaxCount: 3, PercentChance: 10},

			// WEAPON LVL 2
			{BpID: 79, MaxCount: 3, PercentChance: 15},
			{BpID: 84, MaxCount: 3, PercentChance: 15},
			{BpID: 85, MaxCount: 3, PercentChance: 15},
			{BpID: 92, MaxCount: 3, PercentChance: 15},
			{BpID: 93, MaxCount: 3, PercentChance: 15},
			{BpID: 94, MaxCount: 3, PercentChance: 15},

			// BODY lvl 2
			{BpID: 50, MaxCount: 2, PercentChance: 10},
			{BpID: 51, MaxCount: 2, PercentChance: 10},
			{BpID: 55, MaxCount: 2, PercentChance: 10},
			{BpID: 56, MaxCount: 2, PercentChance: 10},
			{BpID: 60, MaxCount: 2, PercentChance: 10},
			{BpID: 61, MaxCount: 2, PercentChance: 10},
			{BpID: 91, MaxCount: 2, PercentChance: 10},
			{BpID: 52, MaxCount: 2, PercentChance: 10},
			{BpID: 57, MaxCount: 2, PercentChance: 10},
			{BpID: 62, MaxCount: 2, PercentChance: 10},
		},
		MaxCountBP: 3, // сколько в коробе может быть разных БП
	},
	3: {
		Bps: []BluePrintAnomalyOption{
			// WEAPON lvl 3
			{BpID: 77, MaxCount: 3, PercentChance: 15},
			{BpID: 81, MaxCount: 3, PercentChance: 15},
			{BpID: 88, MaxCount: 3, PercentChance: 15},

			// BODY lvl 3
			{BpID: 107, MaxCount: 1, PercentChance: 3},
			{BpID: 106, MaxCount: 1, PercentChance: 3},
			{BpID: 3, MaxCount: 1, PercentChance: 3},
			{BpID: 4, MaxCount: 1, PercentChance: 3},
			{BpID: 5, MaxCount: 1, PercentChance: 3},
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
			{ID: 1, MaxCount: 50, PercentChance: 10},
			{ID: 3, MaxCount: 50, PercentChance: 10},
			{ID: 5, MaxCount: 50, PercentChance: 10},
			{ID: 13, MaxCount: 250, PercentChance: 10},
			{ID: 14, MaxCount: 50, PercentChance: 10},
		},
	},
	2: {
		Ammo: []AmmoAnomalyOption{
			{ID: 2, MaxCount: 50, PercentChance: 10},
			{ID: 4, MaxCount: 50, PercentChance: 10},
			{ID: 6, MaxCount: 50, PercentChance: 10},
			{ID: 15, MaxCount: 50, PercentChance: 10},
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
