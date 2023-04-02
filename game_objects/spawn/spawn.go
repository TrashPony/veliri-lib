package spawn

type Spawn struct {
	ID           int         `json:"id"`
	X            int         `json:"x"`
	Y            int         `json:"y"`
	Name         string      `json:"name"`
	Radius       int         `json:"radius"`
	Rotate       int         `json:"-"`
	Type         string      `json:"-"`
	CaptureTeam  string      `json:"capture_team"` // или фракция
	Capture      float64     `json:"capture"`
	TypeGame     string      `json:"-"`
	SubTypeGame  int         `json:"-"`
	Owner        string      `json:"-"`
	CaptureFact  bool        `json:"capture_fact"`
	ObjectsOwner map[int]int `json:"-"` // [id обьекта в базе] id обьекта когда получен при создание игры
	SpawnOwner   map[int]int `json:"-"` // [id спавна в базе]
	NoUserResp   bool        `json:"-"`
	Fraction     string      `json:"-"`
	ReloadTime   int         `json:"-"`
	Count        int         `json:"-"`
	TickPoint    int
}
