package build_option

var BaseConfigTemplates = []BaseConfig{
	// Вариант 1: Классическая симметричная база (твой текущий)
	{
		Template: []*BuildOption{
			{X: 0, Y: 0, Radius: 50, Type: "core"},
			{X: -200, Y: -200, Radius: 50, Type: "turret"},
			{X: 200, Y: 200, Radius: 50, Type: "turret"},
			{X: -200, Y: 200, Radius: 50, Type: "turret"},
			{X: 200, Y: -200, Radius: 50, Type: "turret"},
			{X: -300, Y: 0, Radius: 50, Type: "turret"},
			{X: 0, Y: 300, Radius: 50, Type: "turret"},
			{X: 100, Y: 100, Radius: 50, Type: "radar"},
			{X: -100, Y: -100, Radius: 50, Type: "radar"},
			{X: -150, Y: 150, Radius: 50, Type: "shield"},
			{X: 150, Y: -150, Radius: 50, Type: "shield"},
			{X: 150, Y: 150, Radius: 50, Type: "shield"},
			{X: -150, Y: -150, Radius: 50, Type: "shield"},
		},
		LevelRules: map[OutpostLevel]LevelRules{
			LevelScouting:    {Turrets: 1, Radars: 1, Shields: 0},
			LevelEstablished: {Turrets: 2, Radars: 1, Shields: 1},
			LevelFortified:   {Turrets: 4, Radars: 1, Shields: 2},
			LevelCommand:     {Turrets: 6, Radars: 2, Shields: 4},
		},
		Radius: 350,
	},

	// Вариант 2: Компактная база (все ближе к центру)
	{
		Template: []*BuildOption{
			{X: 0, Y: 0, Radius: 50, Type: "core"},
			{X: -150, Y: -150, Radius: 50, Type: "turret"},
			{X: 150, Y: 150, Radius: 50, Type: "turret"},
			{X: -150, Y: 150, Radius: 50, Type: "turret"},
			{X: 150, Y: -150, Radius: 50, Type: "turret"},
			{X: -250, Y: 0, Radius: 50, Type: "turret"},
			{X: 0, Y: 250, Radius: 50, Type: "turret"},
			{X: 100, Y: 0, Radius: 50, Type: "radar"},
			{X: -100, Y: 0, Radius: 50, Type: "radar"},
			{X: -200, Y: 100, Radius: 50, Type: "shield"},
			{X: 200, Y: -100, Radius: 50, Type: "shield"},
			{X: 100, Y: 200, Radius: 50, Type: "shield"},
			{X: -100, Y: -200, Radius: 50, Type: "shield"},
		},
		LevelRules: map[OutpostLevel]LevelRules{
			LevelScouting:    {Turrets: 1, Radars: 1, Shields: 0},
			LevelEstablished: {Turrets: 3, Radars: 1, Shields: 1}, // +1 турель
			LevelFortified:   {Turrets: 5, Radars: 1, Shields: 2}, // +1 турель
			LevelCommand:     {Turrets: 6, Radars: 2, Shields: 4},
		},
		Radius: 300, // Меньший радиус
	},

	// Вариант 3: Растянутая база (по кругу)
	{
		Template: []*BuildOption{
			{X: 0, Y: 0, Radius: 50, Type: "core"},
			// Турели по кругу
			{X: -250, Y: -150, Radius: 50, Type: "turret"},
			{X: 250, Y: 150, Radius: 50, Type: "turret"},
			{X: -250, Y: 150, Radius: 50, Type: "turret"},
			{X: 250, Y: -150, Radius: 50, Type: "turret"},
			{X: -300, Y: 0, Radius: 50, Type: "turret"},
			{X: 0, Y: 300, Radius: 50, Type: "turret"},
			// Радары ближе
			{X: 50, Y: 150, Radius: 50, Type: "radar"},
			{X: -50, Y: -150, Radius: 50, Type: "radar"},
			// Щиты между турелями
			{X: -200, Y: 200, Radius: 50, Type: "shield"},
			{X: 200, Y: -200, Radius: 50, Type: "shield"},
			{X: -200, Y: -200, Radius: 50, Type: "shield"},
			{X: 200, Y: 200, Radius: 50, Type: "shield"},
		},
		LevelRules: map[OutpostLevel]LevelRules{
			LevelScouting:    {Turrets: 2, Radars: 1, Shields: 0}, // +1 турель на старте
			LevelEstablished: {Turrets: 3, Radars: 1, Shields: 1},
			LevelFortified:   {Turrets: 5, Radars: 2, Shields: 2}, // +1 радар
			LevelCommand:     {Turrets: 6, Radars: 2, Shields: 4},
		},
		Radius: 400, // Больший радиус
	},

	// Вариант 4: Оборонительная база (больше щитов)
	{
		Template: []*BuildOption{
			{X: 0, Y: 0, Radius: 50, Type: "core"},
			// Меньше турелей
			{X: -200, Y: 0, Radius: 50, Type: "turret"},
			{X: 200, Y: 0, Radius: 50, Type: "turret"},
			{X: 0, Y: -200, Radius: 50, Type: "turret"},
			{X: 0, Y: 200, Radius: 50, Type: "turret"},
			{X: -150, Y: -150, Radius: 50, Type: "turret"},
			{X: 150, Y: 150, Radius: 50, Type: "turret"},
			// Больше щитов
			{X: -100, Y: 100, Radius: 50, Type: "radar"},
			{X: 100, Y: -100, Radius: 50, Type: "radar"},
			{X: -250, Y: 0, Radius: 50, Type: "shield"},
			{X: 250, Y: 0, Radius: 50, Type: "shield"},
			{X: 0, Y: 250, Radius: 50, Type: "shield"},
			{X: 0, Y: -250, Radius: 50, Type: "shield"},
			{X: -150, Y: 150, Radius: 50, Type: "shield"},
			{X: 150, Y: -150, Radius: 50, Type: "shield"},
		},
		LevelRules: map[OutpostLevel]LevelRules{
			LevelScouting:    {Turrets: 2, Radars: 1, Shields: 2}, // Щиты с самого начала
			LevelEstablished: {Turrets: 3, Radars: 1, Shields: 3},
			LevelFortified:   {Turrets: 4, Radars: 2, Shields: 4},
			LevelCommand:     {Turrets: 6, Radars: 2, Shields: 6}, // Все щиты
		},
		Radius: 350,
	},

	// Вариант 5: Агрессивная база (больше турелей)
	{
		Template: []*BuildOption{
			{X: 0, Y: 0, Radius: 50, Type: "core"},
			// 8 позиций для турелей!
			{X: -250, Y: -100, Radius: 50, Type: "turret"},
			{X: 250, Y: 100, Radius: 50, Type: "turret"},
			{X: -250, Y: 100, Radius: 50, Type: "turret"},
			{X: 250, Y: -100, Radius: 50, Type: "turret"},
			{X: -100, Y: -250, Radius: 50, Type: "turret"},
			{X: 100, Y: 250, Radius: 50, Type: "turret"},
			{X: -100, Y: 250, Radius: 50, Type: "turret"},
			{X: 100, Y: -250, Radius: 50, Type: "turret"},
			// Меньше щитов и радаров
			{X: 150, Y: 0, Radius: 50, Type: "radar"},
			{X: -150, Y: 0, Radius: 50, Type: "radar"},
			{X: 0, Y: 150, Radius: 50, Type: "shield"},
			{X: 0, Y: -150, Radius: 50, Type: "shield"},
		},
		LevelRules: map[OutpostLevel]LevelRules{
			LevelScouting:    {Turrets: 2, Radars: 1, Shields: 0},
			LevelEstablished: {Turrets: 4, Radars: 1, Shields: 1},
			LevelFortified:   {Turrets: 6, Radars: 1, Shields: 1},
			LevelCommand:     {Turrets: 8, Radars: 2, Shields: 2}, // Все турели!
		},
		Radius: 400,
	},
}
