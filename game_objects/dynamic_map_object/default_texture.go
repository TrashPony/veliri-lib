package dynamic_map_object

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"strings"
)

func (o *Object) DefaultTexture() {
	body, weapon := _const.GetTextureKostil(o.Texture, o.Type)
	if body {
		o.Texture += "_" + strings.ToLower(_const.ALL)
	}

	if weapon {
		for _, wSlot := range o.RangeWeaponSlots() {
			if wSlot.Weapon != nil {
				wSlot.Weapon.Name += "_" + strings.ToLower(_const.ALL)
			}
		}
	}
}
