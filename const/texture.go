package _const

func GetTextureKostil(texture, objectType string) (bool, bool) {
	// TODO костыль :(
	if !(texture == "replic_gauss_gun" || texture == "unknown_civilization_jammer" ||
		texture == "explores_antenna" || texture == "explores_observatory" || texture == "mini_turret_1" ||
		texture == "mini_turret_2" || texture == "mini_shield_generator_1") {

		if objectType == "turret" || objectType == "shield" || objectType == "generator" ||
			objectType == "jammer" || objectType == "missile_defense" || objectType == "meteorite_defense" ||
			objectType == "radar" || objectType == "storage" || objectType == "beacon" ||
			objectType == "extractor" || objectType == "repair_station" {

			return true, true
		}

		return false, true
	}

	return false, false
}
