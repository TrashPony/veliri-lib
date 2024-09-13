package descriptions

import _const "github.com/TrashPony/veliri-lib/const"

var AmmoDescription = map[string]map[string]DescriptionItem{
	_const.EN: {
		"piu-piu":                       {Name: "B-a medium-caliber projectile", Description: "<p> A standardized type of medium-caliber projectiles for kinetic and ballistic weapons that provide significant damage, but, as a rule, are not always distinguished by accuracy and high rate of fire. </p>"},
		"piu-piu_2":                     {Name: "B-a small-caliber projectile", Description: "<p> A standardized type of small-bore projectiles for kinetic and ballistic weapons that do not provide high damage, but, as a rule, contribute to high accuracy and rate of fire. </p>"},
		"piu-piu_3":                     {Name: "Бронебойный мелкокалиберный снаряд", Description: "<p></p>"},
		"piu-piu_4":                     {Name: "Разрывной мелкокалиберный снаряд", Description: "<p></p>"},
		"piu-piu_5":                     {Name: "Бронебойный среднекалиберной снаряд", Description: "<p></p>"},
		"piu-piu_6":                     {Name: "Кумулятивный среднекалиберной снаряд", Description: "<p></p>"},
		"medium_lens":                   {Name: "Medium focusing cartridge", Description: "<p> Standardized cartridge for charging medium and large laser beam impact equipment. </p>"},
		"aim_medium_missile_bullet":     {Name: "Cicada - Homing Medium Missile", Description: "<p> Cicada - development of a line of unguided missiles of the Sistema family with the aim of creating homing shells. Provides average damage to any captured targets. </p>"},
		"aim_big_missile_bullet":        {Name: "Vindicator - Homing Large Rocket", Description: "<p> The Vindicator is a continuation of ideas started in Cicada and the Sistema unguided rockets family. This homing missile, which can only be used in large rocket firing systems, is a real terror to the enemies it captures. </p>"},
		"ballistics_artillery_bullet":   {Name: "B-a large-caliber projectile", Description: "<p> A standardized type of large-caliber projectiles for kinetic and ballistic weapons that guarantee devastating damage, but cannot be used with rapid-fire or extremely accurate weapons. </p>"},
		"ballistics_artillery_bullet_3": {Name: "Бронебойный крупнокалиберный снаряд", Description: "<p></p>"},
		"ballistics_artillery_bullet_4": {Name: "Кумулятивный крупнокалиберный снаряд", Description: "<p></p>"},
		"medium_missile_bullet":         {Name: "System.A-2 - Medium missile", Description: "<p> Massive and military-standardized missile of medium destructive power. Cannot be used in high-dimensional missile system weapons. </p>"},
		"big_missile_bullet":            {Name: "System.O-5 - Large rocket", Description: "<p> Massive and military-standardized missile of great destructive power. Cannot be used in missile weapons with an average overall weight or less. </p>"},
		"small_lens":                    {Name: "Small focusing cartridge", Description: "<p> Standardized cartridge for charging small and medium-sized laser beam impact tools. </p>"},
		"small_lens_beam":               {Name: "Малый лучевой картридж", Description: "<p></p>"},
		"small_missile_bullet":          {Name: "Маля ракета", Description: "<p></p>"},
		"big_lens_1":                    {Name: "big_lens_1", Description: "<p></p>"},
		"medium_lens_beam":              {Name: "medium_lens_beam", Description: "<p></p>"},
	},
	_const.RU: {
		"piu-piu":                       {Name: "Б-а среднекалиберной снаряд", Description: "<p>Стандартизированный тип среднекалиберных снарядов для кинетических и баллистических орудий, обеспечивающих ощутимый урон, но, как правило, не всегда отличающихся кучностью и высокой скорострельностью.</p>"},
		"ballistics_artillery_bullet":   {Name: "Б-а крупнокалиберный снаряд", Description: "<p>Стандартизированный тип крупнокалиберных снарядов для кинетических и баллистических орудий, гарантирующих сокрушительный урон, но невозможных к применению в скорострельных или крайне точных орудиях.</p>"},
		"ballistics_artillery_bullet_3": {Name: "Бронебойный крупнокалиберный снаряд", Description: "<p>Бронебойный снаряд, обладает высокой скоростью и точностью на дальнюю дистанцию, однако лишен урона по площади.</p>"},             // TODO
		"ballistics_artillery_bullet_4": {Name: "Кумулятивный крупнокалиберный снаряд", Description: "<p>Кумулятивный снаряд, обладает огромным уроном и большой точностью, однако дальность полета и площадь поражения минимальны.</p>"}, // TODO
		"medium_missile_bullet":         {Name: "Система.А-2 - Средняя ракета", Description: "<p>Массовая и стандартизированная под военные нужды ракета средней поражающей силы. Невозможна к использования в орудиях ракетной системы высоко-габаритного образца.</p>"},
		"big_missile_bullet":            {Name: "Система.О-5 - Большая ракета", Description: "<p>Массовая и стандартизированная под военные нужды ракета большой поражающей силы. Невозможна к использования в ракетных орудиях средней габаритной массы или меньше.</p>"},
		"small_lens":                    {Name: "Малый фокусирующий картридж", Description: "<p>Стандартизированный картридж для зарядки в малые и средние орудия лазерно-лучевого воздействия.</p>"},
		"medium_lens":                   {Name: "Средний фокусирующий картридж", Description: "<p>Стандартизированный картридж для зарядки в средние и крупные орудия лазерно-лучевого воздействия.</p>"},
		"piu-piu_2":                     {Name: "Б-а мелкокалиберный снаряд", Description: "<p>Стандартизированный тип мелкокалиберных снарядов для кинетических и баллистических орудий, не обеспечивающих высокий урон, но, как правило, способствующие высокой кучности и скорострельности.</p>"},
		"piu-piu_3":                     {Name: "Бронебойный мелкокалиберный снаряд", Description: "<p>Обладает высоким уроном и высокой скорость полета пули, однако малой дальностью</p>"}, // TODO
		"piu-piu_4":                     {Name: "Разрывной мелкокалиберный снаряд", Description: "<p>Обладает высоким уроном, уроном по площади, высоким разбросом и малой дальностью</p>"},  // TODO
		"piu-piu_5":                     {Name: "Бронебойный среднекалиберной снаряд", Description: "<p>Высокая скорость полета и урон, нет урона по площади</p>"},                           // TODO
		"piu-piu_6":                     {Name: "Кумулятивный среднекалиберной снаряд", Description: "<p>Огромный урон, малая дальность и высокий разброс.</p>"},                             // TODO
		"aim_medium_missile_bullet":     {Name: "Цикада - Самонаводящиеся средняя ракета", Description: "<p>Цикада – развитие линейки неуправляемых ракет семейства «Система» с целью создания самонаводящихся на цель снарядов. Обеспечивает усреднённый урон по любым захваченным целям.</p>"},
		"aim_big_missile_bullet":        {Name: "Воздаятель - Самонаводящиеся большая ракета", Description: "<p>Воздаятель – продолжение идей, начатых в «Цикада» и линейке семейства неуправляемых ракет «Система». Эта самонаводящаяся ракета, что может быть использована только в крупногабаритных ракетных огневых системах, истинный ужас для врагов, что были ею захвачены.</p>"},
		"small_lens_beam":               {Name: "Малый лучевой картридж", Description: "<p></p>"}, // TODO
		"small_missile_bullet":          {Name: "Малая ракета", Description: "<p></p>"},           // TODO
		"big_lens_1":                    {Name: "big_lens_1", Description: "<p></p>"},
		"medium_lens_beam":              {Name: "medium_lens_beam", Description: "<p></p>"},
	},
}
