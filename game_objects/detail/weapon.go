package detail

import "github.com/TrashPony/veliri-lib/game_objects/coordinate"

type Weapon struct {
	// если тут появяться ссылочные типы данных включая срезы, карты и тд, надо будет делать глубокое копирование в
	// factory/gameTypes/weapons
	ID                 int                      `json:"id"`
	Name               string                   `json:"name"`
	Specification      string                   `json:"specification"`
	MinRange           int                      `json:"min_range"`
	MaxRange           int                      `json:"max_range"`
	MinAngle           int                      `json:"min_angle"`
	MaxAngle           int                      `json:"max_angle"`
	Accuracy           int                      `json:"accuracy"`
	AmmoCapacity       int                      `json:"ammo_capacity"`
	Artillery          bool                     `json:"artillery"`
	Power              int                      `json:"power"`
	MaxHP              int                      `json:"max_hp"`
	Type               string                   `json:"type"`          /* firearms, missile_weapon, laser_weapon */
	StandardSize       int                      `json:"standard_size"` /* small - 1, medium - 2, big - 3 */
	Size               int                      `json:"size"`          /* занимаемый обьем в кубо метрах */
	XAttach            int                      `json:"x_attach"`
	YAttach            int                      `json:"y_attach"`
	RotateSpeed        int                      `json:"rotate_speed"`
	CountFireBullet    int                      `json:"count_fire_bullet"`
	BulletSpeed        int                      `json:"bullet_speed"`
	ReloadAmmoTime     int                      `json:"reload_ammo_time"`
	ReloadTime         int                      `json:"reload_time"`
	DelayFollowingFire int                      `json:"delay_following_fire"`
	FirePositions      []*coordinate.Coordinate `json:"fire_positions"`
	Unit               bool                     `json:"unit"`
	Structure          bool                     `json:"structure"`
	Requirements       map[int]int              `json:"requirements"` // [id]lvl
	DamageModifier     float64                  `json:"damage_modifier"`
	AllowAmmo          []int                    `json:"allow_ammo"`

	// оружие накопительного типа (зажатая мышка коит энергию, отпускание выпускает снаряд)
	AccumulationFirePower   bool    `json:"accumulation_fire_power"`
	AccumulationFull        float64 `json:"accumulation_full"`
	AccumulationFullTimeOut int     `json:"accumulation_full_time_out"` // время в мс когда произойдет перегрев или автовыстрел
	PushingPower            int     `json:"pushing_power"`
	Variants                []int   `json:"variants"`
	PassAngle               int     `json:"pass_angle"`
}

func (w *Weapon) GetName() string {
	return w.Name
}

func (w *Weapon) GetItemType() string {
	return ""
}

func (w *Weapon) GetItemName() string {
	return ""
}

func (w *Weapon) GetInMap() bool {
	return false
}

func (w *Weapon) GetIcon() string {
	return ""
}

func (w *Weapon) GetStandardSize() int {
	return w.StandardSize
}

func (w *Weapon) GetTypeSlot() int {
	return 0
}
