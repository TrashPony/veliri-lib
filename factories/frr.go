package factories

import (
	"github.com/TrashPony/veliri-lib/game_objects/frr"
	"math/rand"
)

const sizeFrr = 100

var frrTypes = map[int]frr.FRR{
	// weapons
	1: {
		ID:   1,
		Name: "frr_arm_blt", // Field Research Report, Armament, Ballistic
		Size: sizeFrr,
	},
	2: {
		ID:   2,
		Name: "frr_arm_lsr", // Field Research Report, Armament, Laser
		Size: sizeFrr,
	},
	3: {
		ID:   3,
		Name: "frr_arm_msl", // Field Research Report, Armament, Missile
		Size: sizeFrr,
	},
	// bodies
	4: {
		ID:   4,
		Name: "frr_vhc_whl", // Field Research Report, Mobility, Wheeled
		Size: sizeFrr,
	},
	5: {
		ID:   5,
		Name: "frr_vhc_hov", // Field Research Report, Mobility, Wheeled
		Size: sizeFrr,
	},
	6: {
		ID:   6,
		Name: "frr_vhc_trd", // Field Research Report, Mobility, Wheeled
		Size: sizeFrr,
	},
	// equipment
	// если в будущем потребуется расширить SPT.EQP на разный эквип то мы сделаем так что надо будет 50 условно EQP и еще 50 какого то уже специального
	7: {
		ID:   7,
		Name: "frr_spt_eqp", // Field Research Report, Support, Equipment
		Size: sizeFrr,
	},
}

func GetFrr(id int) *frr.FRR {
	newFrr := frrTypes[id]
	return &newFrr
}

func GetAllFrr() map[int]frr.FRR {
	return frrTypes
}

func GetRandomFrr() frr.FRR {
	fs := make([]frr.FRR, 0)

	for _, p := range frrTypes {
		fs = append(fs, p)
	}

	return fs[rand.Intn(len(fs))]
}
