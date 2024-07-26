package _const

import "time"

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

	PowerInWork = 1000 // 1 заряд равен 100 энергии

	// очки опыта
	// Оружия
	KillStructure      = 40
	KillUserExperience = 50 // * body_size

	// TODO начать их использовать)
	Replicas     = "Replics"
	ReplicasByte = 1
	Explores     = "Explores"
	ExploresByte = 2
	Reverses     = "Reverses"
	ReversesByte = 3
	APD          = "APD"
	APDByte      = 4
	FAUNA        = "Fauna"
	FAUNAByte    = 5
	FarGod       = "FarGod"
	FarGodByte   = 6
	Empty        = "Empty"
	EmptyByte    = 7

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
	CaptureSector       = 1000 * 100

	//MaxBaseTax = 50
	//MinBaseTax = 10

	BaseRecyclerLost       = 50 // 25% снимает общий навык, еще 25% специальный
	BaseFineProductionTime = 100

	AdminUserRole = "admin"

	PVP    = "pvp"
	PK     = "pk"
	Strike = "strike"

	TimePvpViolators = 30
	TimePkViolators  = 900
	TimeTakeStrike   = 180

	ItemSizeRadius = 5

	TerrainGrowTime = 3000

	// шанс 1 к chanceMeteoriteRain
	ChanceMeteoriteRain = 10000

	// шанс 1 к chanceSimpleMeteorite
	ChanceSimpleMeteorite = 10

	MaxRelations            = 200
	MinRelations            = -200
	DefaultUnionRelations   = 50
	DefaultHostileRelations = -50

	ChangeRelationFailTrade       = 0.0001   // 0.001% * (count + default_price)
	ChangeRelationTrade           = 0.000002 // 0.000001% * (count + default_price)
	ChangeRelationProcess         = 0.0001   // 0.01% * count
	ChangeRelationKill            = 0.006    // 6%
	ChangeRelationMissionComplete = 0.05     // 20%
	ChangeRelationSectorCapture   = 0.05     // 20%

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
	MaxPoints             = 300 * 3 // (300 сек * 3 базы) по 1 очку в секунду == 5 минут

	BaseMarketTax         = 0.1
	OfficeStorage         = 10000000
	OfficeRentalTime      = time.Hour * 24 * 30
	OfficeRentalBasePrice = 2500000

	WarPrice  = 1000000
	WarCycle  = 24 * 7 * time.Hour
	WarBorder = 24 * time.Hour

	BaseThoriumFillUp = 10000

	HandlerRadius     = 60
	HandlerDropRadius = 75
	HandlerSoftRadius = 120

	CorporationPolicyAll            = "all"             // все
	CorporationPolicyOnlyUnions     = "only-unions"     // союзники
	CorporationPolicyOnlyRivals     = "only-rivals"     // соперники
	CorporationPolicyOnlyOwner      = "only-owner"      // только кластер
	CorporationPolicyExceptUnions   = "except-unions"   // все кроме союзников
	CorporationPolicyExceptRivals   = "except-rivals"   // все кроме соперников
	CorporationPolicyExceptHostiles = "except-hostiles" // все кроме врагов

	CorporationTurretPolicyAll                   = "all"                  // Атакуют всех
	CorporationTurretPolicyExceptOwner           = "except-owner"         // Атакуют всех кроме кластера
	CorporationTurretPolicyExceptUnions          = "except-unions"        // Атакуют всех кроме кластера и союзников
	CorporationTurretPolicyExceptRivals          = "except-rivals"        // Атакуют всех кроме кластера и соперников
	CorporationTurretPolicyExceptRivalsAndUnions = "except-rivals-unions" // Атакуют всех кроме кластера, соперников и союзников (то есть нейтралов и врагов)
	CorporationTurretPolicyOnlyHostile           = "only-hostile"         // Атакуют только врагов (пк, апд, враги на войне) и агрессивных (те кто атакуют корпоративную собсвенность в секторе)
	CorporationTurretPolicyDisable               = "disable"              // Не работают

	CorporationRepairPolicyOnlyOwner     = "only-owner"     // Кластер
	CorporationRepairPolicyUnions        = "unions"         // Кластер и союзников
	CorporationRepairPolicyRivals        = "rivals"         // Кластер и соперников
	CorporationRepairPolicyExceptHostile = "except-hostile" // Всех кроме врагов
	CorporationRepairPolicyDisable       = "disable"        // Не работают

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
	"office_stock":         9,
	"hangar":               10,
	"hangar_aInv_resource": 11,
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
	"map":             14,
	"":                15,
}

var FractionByte = map[string]byte{
	Empty:    EmptyByte,
	Replicas: ReplicasByte,
	Explores: ExploresByte,
	Reverses: ReversesByte,
	APD:      APDByte,
	FAUNA:    FAUNAByte,
	FarGod:   FarGodByte,
}

var PointType = map[string]byte{
	"destroy":           1,
	"damage_hostile":    2,
	"assist_destroy":    3,
	"capture":           4,
	"destroy_structure": 5,
	"capture_sector":    6,
}
