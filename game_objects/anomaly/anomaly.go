package anomaly

type Anomaly struct {
	x            int
	y            int
	MapID        int `json:"map_id"`
	power        int // сила это растояние на котором можно увидить аномалию на сканере со сканером нулевого радиуса, общая дальность сила аномалии + радиус сканера
	Type         int
	RealType     int
	Lvl          int
	UserIDAccess int  `json:"user_id_access"` // видит и может откапать только игрок с этим ИД, если 0 то общая
	Open         bool `json:"-"`
	box          bool
	resource     bool
	text         bool
	bots         []string // apd боты [role]count
	botsCount    int
}

func (a *Anomaly) GetX() int {
	return a.x
}

func (a *Anomaly) SetX(x int) {
	a.x = x
}

func (a *Anomaly) GetY() int {
	return a.y
}

func (a *Anomaly) SetY(y int) {
	a.y = y
}

func (a *Anomaly) GetPower() int {
	return a.power
}

func (a *Anomaly) SetPower(power int) {
	a.power = power
}

func (a *Anomaly) SetBox() {
	a.box = true
}

func (a *Anomaly) SetRes() {
	a.resource = true
}

func (a *Anomaly) SetDialog() {
	a.text = true
}

func (a *Anomaly) SetBots(roles []string, count int) {
	a.bots = roles
	a.botsCount = count
}

func (a *Anomaly) GetBots() ([]string, int) {
	return a.bots, a.botsCount
}
