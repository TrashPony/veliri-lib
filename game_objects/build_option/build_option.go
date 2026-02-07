package build_option

type OutpostLevel int

const (
	LevelScouting    OutpostLevel = iota // Разведывательный лагерь
	LevelEstablished                     // Установившийся лагерь
	LevelFortified                       // Укрепленный лагерь
	LevelCommand                         // Командный центр (максимум)
)

const (
	MinPeaceTimeForGrowth = 300000 // 5 минут без атак для роста
	MaxLevel              = 3      // Максимальный уровень (0-3)
	BuildTimeOut          = 5 * 60
	DestroyTimeOut        = 30 * 60
)

type BuildOption struct {
	X        int    `json:"x"`
	Y        int    `json:"y"`
	Radius   int    `json:"radius"`
	ObjectID int    `json:"object_id"`
	Type     string `json:"type"`
}

type LevelRules struct {
	Turrets int `json:"turrets"`
	Radars  int `json:"radars"`
	Shields int `json:"shields"`
}

type BaseConfig struct {
	UUID            string         `json:"uuid"`
	Template        []*BuildOption `json:"template"`
	CurrentTemplate []*BuildOption `json:"current_template"`
	Radius          int            `json:"radius"`
	X               int            `json:"x"`
	Y               int            `json:"y"`
	Rotate          int            `json:"rotate"`
	BuildTimeOut    int64          `json:"build_time_out"`
	DestroyTimeOut  int64          `json:"destroy_time_out"`
	// прогрессия
	Level      OutpostLevel                `json:"level"`
	LevelRules map[OutpostLevel]LevelRules `json:"level_rules"`
	BuildStage int                         `json:"build_stage"` // должен накопить 5 минут (300000мс) без атак для роста
	MainBaseID int                         `json:"main_base"`
}
