package db

type DynObject struct {
	ID               int
	IDType           int
	X                int
	Y                int
	Rotate           int
	Scale            int
	HP               int
	Energy           int
	OwnerPlayerID    int
	OwnerType        string
	OwnerID          int
	OwnerBaseID      int
	MaxScale         int
	Complete         int
	GrowCycle        int
	GrowTime         int
	GrowLeftTime     int
	Fraction         string
	BodyID           int
	LifeTime         int
	DestroyLeftTimer bool
	DestroyTimer     int
	FractionWarrior  bool
	Attributes       []byte
	EquipID          int
	BoxID            int
}

type ReservoirOption struct {
	ID         int
	ResourceID int
	X          int
	Y          int
	Rotate     int
	Count      int
}
