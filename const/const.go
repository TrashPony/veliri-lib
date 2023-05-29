package _const

const (
	CellSize              = 32
	DiscreteSize          = 4 * CellSize
	DiscreteFindPathSize  = 8 * CellSize
	Gravity               = 980.0 / 2
	UnitHeight            = 35
	AmmoRadius            = 3
	CompleteStructureWork = 50

	DistOpenBox       = 90 // дистания на которой юниты могут взаимодействовать с ящиками на карте
	DistInstallBox    = 75
	DistInitBuild     = 150 // дистания на которой юниты могут начинать строительство
	ServerTick        = 32
	ServerBulletTick  = 32
	ServerTickSecPart = float64(int(1000 / ServerTick))
	AiTick            = 1000 // если меняешь посмотри что везде правильно сработает таймаут

	// очки опыта
	// Оружия
	KillStructure      = 40
	KillUserExperience = 20 // * body_size

	// TODO начать их использовать)
	Replicas = "Replics"
	Explores = "Explores"
	Reverses = "Reverses"
	APD      = "APD"
	FAUNA    = "Fauna"
	ALL      = "ALL"

	QuickBattle = "quick_battle"
	OpenWorld   = "open_world"

	ReplicasQuickBaseID = -1
	ExploresQuickBaseID = -2
	ReversesQuickBaseID = -3

	// session exp multipliers
	KillsK              = 0.5 * 100
	AssistK             = 0.25 * 100
	Damage              = 0.3 * 100
	IntelligenceDamageK = 1 * 100
	DestroyStructure    = 200 * 100
	CaptureK            = 250 * 100
	CaptureSector       = 1500 * 100

	//MaxBaseTax = 50
	//MinBaseTax = 10

	BaseRecyclerLost       = 50 // 25% снимает общий навык, еще 25% специальный
	BaseFineProductionTime = 100

	AdminUserRole = "admin"

	PVP              = "pvp"
	PK               = "pk"
	TimePvpViolators = 30
	TimePkViolators  = 300

	ItemSizeRadius = 5

	GrowCycle       = 25
	TerrainGrowTime = 1000

	// шанс 1 к chanceMeteoriteRain
	ChanceMeteoriteRain = 10000

	// шанс 1 к chanceSimpleMeteorite
	ChanceSimpleMeteorite = 10

	MaxRelations                  = 200
	MinRelations                  = -200
	DefaultUnionRelations         = 50
	DefaultHostileRelations       = -50
	ChangeRelationKill            = 0.06 // 6%
	ChangeRelationMissionComplete = 0.2  // 20%

	LaserWeapon    = "laser"
	FirearmsWeapon = "firearms"
	MissileWeapon  = "missile"
	Meteorite      = "meteorite"
	MineWeapon     = "mine"
	BuildBullet    = "builder"

	ObjTypeStaticGeo     = 1
	ObjTypeStaticObject  = 2
	ObjTypeDynamicObject = 3
	ObjTypeReservoir     = 4
	ObjTypeUnit          = 5
	ObjTypeBox           = 6
	ObjTypeFlore         = 7
	ObjTypeMeteorite     = 8

	EN = "EN"
	RU = "RU"

	SpriteSize  = 128.00
	SpriteSize2 = SpriteSize / 2

	AfterburnerSlotNumber = -1
	MaxPoints             = 60 * 3 // (600 сек * 3 базы) по 1 очку в секунду == 10 минут
)

var ItemBinTypes = map[string]int{
	"":           0,
	"weapon":     1,
	"ammo":       2,
	"equip":      3,
	"body":       4,
	"resource":   5,
	"recycle":    6,
	"detail":     7,
	"boxes":      8,
	"blueprints": 9,
	"trash":      10,
	"product":    11,
}

var SourceItemBin = map[string]int{
	"squadInventory":       1,
	"Constructor":          2,
	"empty":                3,
	"box":                  4,
	"storage":              5,
	"scanner":              6,
	"aInv:resource":        7,
	"scannerAInv:resource": 8,
}

var MapBinItems = map[string]int{
	"transport":       1,
	"box":             2,
	"unit":            3,
	"drone":           4,
	"reservoir":       5,
	"dynamic_objects": 6,
	"shield":          7,
	"item":            8,
	"object":          9,
	"bullet":          10,
	"mark":            11,
	"rope_element":    12,
	"base":            13,
}

var FractionByte = map[string]byte{
	"Empty_Fraction": 0,
	Replicas:         1,
	Explores:         2,
	Reverses:         3,
	APD:              4,
}

var PointType = map[string]byte{
	"destroy":           1,
	"damage_hostile":    2,
	"assist_destroy":    3,
	"capture":           4,
	"destroy_structure": 5,
	"capture_sector":    6,
}
