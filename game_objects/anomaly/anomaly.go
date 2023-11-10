package anomaly

type Anomaly struct {
	Key          string   `json:"key"`
	X            int      `json:"x"`
	Y            int      `json:"y"`
	MapID        int      `json:"map_id"`
	Power        int      `json:"power"` // сила это растояние на котором можно увидить аномалию на сканере со сканером нулевого радиуса, общая дальность сила аномалии + радиус сканера
	Type         int      `json:"type"`
	RealType     int      `json:"realType"`
	Lvl          int      `json:"lvl"`
	BoxID        int      `json:"boxID"`
	UserIDAccess int      `json:"user_id_access"` // видит и может откапать только игрок с этим ИД, если 0 то общая
	Open         bool     `json:"-"`
	Box          bool     `json:"box"`
	Resource     bool     `json:"resource"`
	Text         bool     `json:"text"`
	Bots         []string `json:"bots"` // apd боты [role]count
	BotsCount    int      `json:"bots_count"`
	Remove       bool     `json:"remove"`
}

func (a *Anomaly) GetX() int {
	return a.X
}

func (a *Anomaly) SetX(x int) {
	a.X = x
}

func (a *Anomaly) GetY() int {
	return a.Y
}

func (a *Anomaly) SetY(y int) {
	a.Y = y
}

func (a *Anomaly) GetPower() int {
	return a.Power
}

func (a *Anomaly) SetPower(power int) {
	a.Power = power
}

func (a *Anomaly) SetBox() {
	a.Box = true
}

func (a *Anomaly) SetBoxID(id int) {
	a.BoxID = id
}

func (a *Anomaly) SetRes() {
	a.Resource = true
}

func (a *Anomaly) SetDialog() {
	a.Text = true
}

func (a *Anomaly) SetBots(roles []string, count int) {
	a.Bots = roles
	a.BotsCount = count
}

func (a *Anomaly) GetBots() ([]string, int) {
	return a.Bots, a.BotsCount
}
