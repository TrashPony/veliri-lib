package ammo

type Ammo struct {
	// если тут появяться ссылочные типы данных включая срезы, карты и тд, надо будет делать глубокое копирование в
	// factory/gameTypes/ammo
	ID                     int         `json:"id"`
	Name                   string      `json:"name"`
	Specification          string      `json:"specification"`
	Type                   string      `json:"type"`
	MinDamage              int         `json:"min_damage"`
	MaxDamage              int         `json:"max_damage"`
	AreaCovers             int         `json:"area_covers"`
	StandardSize           int         `json:"standard_size"` /* small - 1, medium - 2, big - 3 */
	EquipDamage            int         `json:"equip_damage"`
	EquipCriticalDamage    int         `json:"equip_critical_damage"`
	Size                   int         `json:"size"`
	ChaseTarget            bool        `json:"chase_target"`
	BulletSpeed            int         `json:"bullet_speed"`
	Requirements           map[int]int `json:"requirements"` // [id]lvl
	ChaseOption            string      `json:"chase_option"`
	Rotate                 float64     `json:"rotate"`
	Unit                   bool        `json:"unit"`
	ChaseCatchDestination  int         `json:"chase_catch_destination"`
	DetonationTimeOut      int         `json:"detonation_time_out"`
	DetonationStartTimeOut int         `json:"detonation_start_time_out"`
	DetonationDistance     int         `json:"detonation_distance"`
	PushingPower           int         `json:"pushing_power"`
	Gravity                float64     `json:"-"`
	ForceAnimate           bool        `json:"force_animate"`
	LifeTime               int         `json:"life_time"`
	ClusterID              int         `json:"cluster_id"`
	ClusterCount           int         `json:"cluster_count"`
	KineticsDamage         int         `json:"kinetics_damage"`
	ExplosionDamage        int         `json:"explosion_damage"`
	ThermoDamage           int         `json:"thermo_damage"`
	DistSpread             int         `json:"dist_spread"` // 1-100, 100 - нет рассеивания, 1 - рассеивание х100
	AdditionalAmmoID       int         `json:"additional_ammo_id"`
	Tech                   int         `json:"tech"`
}

func (a *Ammo) GetName() string {
	return a.Name
}

func (a *Ammo) GetSize() int {
	return a.Size
}

func (a *Ammo) GetItemType() string {
	return ""
}

func (a *Ammo) GetItemName() string {
	return ""
}

func (a *Ammo) GetInMap() bool {
	return false
}

func (a *Ammo) GetIcon() string {
	return ""
}

func (a *Ammo) GetStandardSize() int {
	return a.StandardSize
}

func (a *Ammo) GetTypeSlot() int {
	return 0
}

func (a *Ammo) GetTech() int {
	return a.Tech
}
