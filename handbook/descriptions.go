package handbook

import (
	"github.com/TrashPony/veliri-lib/handbook/descriptions"
)

//func init() {
//
//	for key, v := range ProductsDescription[_const.RU] {
//
//		time.Sleep(time.Second)
//		var err error
//
//		if v.Name != "" {
//			v.Name, err = Translate(v.Name, _const.RU, _const.EN)
//			if err != nil {
//				panic(err)
//			}
//
//			v.Name = strings.ReplaceAll(v.Name, `\"`, `\\\"`)
//			v.Name = strings.ReplaceAll(v.Name, `"`, `\"`)
//			v.Name = strings.ReplaceAll(v.Name, "\n", "")
//		}
//
//		if v.Description != "" {
//			v.Description, err = Translate(v.Description, _const.RU, _const.EN)
//			if err != nil {
//				panic(err)
//			}
//
//			v.Description = strings.ReplaceAll(v.Description, `\"`, `\\\"`)
//			v.Description = strings.ReplaceAll(v.Description, `"`, `\"`)
//			v.Description = strings.ReplaceAll(v.Description, "\n", "")
//		}
//
//		fmt.Println(`"` + key + `":{ID: ` + strconv.Itoa(v.ID) + `, Name: "` + v.Name + `", Description: "` + v.Description + `"},`)
//	}
//}

var AllDescription = map[string]map[string]map[string]descriptions.DescriptionItem{
	"weapon":     descriptions.WeaponDescription,
	"ammo":       descriptions.AmmoDescription,
	"equip":      descriptions.EquipDescription,
	"body":       descriptions.BodyDescription,
	"resource":   descriptions.ResourceDescription,
	"recycle":    descriptions.RecycleDescription,
	"detail":     descriptions.DetailDescription,
	"boxes":      descriptions.BoxDescription,
	"trash":      descriptions.TrashDescription,
	"skill":      descriptions.SkillDescription,
	"structure":  descriptions.StructureDescription,
	"categories": descriptions.CategoriesDescription,
	"sector":     descriptions.SectorDescription,
	"base":       descriptions.BaseDescription,
	"region":     descriptions.RegionsDescription,
	"product":    descriptions.ProductsDescription,
	"rank":       descriptions.RankDescription,
	"effects":    descriptions.EffectsDescription,
	"fuel":       descriptions.FuelDescription,
	"mine_drone": descriptions.MineDrones,
	"package":    descriptions.Packages,
	"book":       descriptions.BookDescription,
}

func GetInfo(t, l, n string) descriptions.DescriptionItem {
	return AllDescription[t][l][n]
}
