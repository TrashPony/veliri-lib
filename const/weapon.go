package _const

// fraction/size/ []names
var fractionWeapon = map[string][]string{
	"Replics":  {"replic_weapon_1", "replic_weapon_2", "replic_weapon_3", "replic_weapon_4"},
	"Explores": {"explores_weapon_1", "explores_weapon_2", "explores_weapon_3", "explores_weapon_4"},
	"Reverses": {"reverses_weapon_1", "reverses_weapon_2", "reverses_weapon_3", "reverses_weapon_4", "reverses_weapon_5"},
	"APD":      {"small_laser", "big_laser", "small_missile", "tank_gun", "artillery", "big_missile"},
}

var weaponDefaultSkin = map[string]string{
	"replic_weapon_1":   "_skin_1",
	"replic_weapon_2":   "_skin_1",
	"replic_weapon_3":   "_skin_1",
	"replic_weapon_4":   "_skin_1",
	"explores_weapon_1": "_skin_1",
	"explores_weapon_2": "_skin_1",
	"explores_weapon_3": "_skin_1",
	"explores_weapon_4": "_skin_1",
	"reverses_weapon_1": "_skin_1",
	"reverses_weapon_2": "_skin_1",
	"reverses_weapon_3": "_skin_1",
	"reverses_weapon_4": "_skin_1",
	"reverses_weapon_5": "_skin_1",
	"artillery":         "_skin_1",
	"big_laser":         "_skin_1",
	"big_missile":       "_skin_1",
	"small_laser":       "_skin_1",
	"small_missile":     "_skin_1",
	"tank_gun":          "_skin_1",
}

func GetFractionWeapon(fraction string) []string {
	weapons := make([]string, 0)

	for _, w := range fractionWeapon[fraction] {
		weapons = append(weapons, w)
	}

	return weapons
}

func GetDefaultWeaponSkin(weapon string) string {
	return weaponDefaultSkin[weapon]
}
