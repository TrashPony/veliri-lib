package db

type DynObject struct {
	ID                int
	IDType            int
	X                 int
	Y                 int
	Rotate            int
	Scale             int
	HP                int
	Energy            int
	OwnerPlayerID     int
	OwnerType         string
	CorporationID     int
	OwnerID           int
	OwnerBaseID       int
	MaxScale          int
	Complete          int
	GrowTime          int
	GrowLeftTime      int
	Fraction          string
	BodyID            int
	LifeTime          int
	DestroyLeftTimer  bool
	DestroyTimer      int
	FractionWarrior   bool
	Attributes        []byte
	EquipID           int
	BoxID             int
	Immortal          bool
	Priority          int
	Filters           []byte
	Recipe            int
	NextRecipe        int
	CraftTime         int
	CraftRun          int
	CurrentBuildSlots []byte
}

type ReservoirOption struct {
	ID         int
	ResourceID int
	X          int
	Y          int
	Rotate     int
	Count      int
	Priority   int
}

type UnitOption struct {
	ID           int
	X            int
	Y            int
	OwnerID      int
	BodyID       int
	WorkPriority []byte
}
