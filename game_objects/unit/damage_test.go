package unit

import (
	"fmt"
	"github.com/TrashPony/veliri-lib/game_objects/detail"
	"testing"
)

func TestDamage(t *testing.T) {
	u := Unit{}

	u.SetBody(&detail.Body{
		ProtectionToKinetics:  75,
		ProtectionToThermo:    50,
		ProtectionToExplosion: 25,
	})

	u.SetHP(1000)
	_, _, passDamage, _ := u.SetDamage(1000, 1.0, 33, 33, 33, false, nil)
	fmt.Println(passDamage)
}
