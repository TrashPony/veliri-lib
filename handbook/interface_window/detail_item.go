package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var DetailItem = map[string]map[string]string{
	// BonusTab
	"button_1": {
		_const.RU: `Навыки`,
		_const.EN: `Skills`,
	},
	// CharacteristicTab
	"small": {
		_const.RU: `малый`,
		_const.EN: `small	`,
	},
	"average": {
		_const.RU: `средний`,
		_const.EN: `average	`,
	},
	"large": {
		_const.RU: `большой`,
		_const.EN: `large`,
	},
	"firearms": {
		_const.RU: `огнестрельное`,
		_const.EN: `firearms`,
	},
	"missile": {
		_const.RU: `ракетное`,
		_const.EN: `missile`,
	},
	"laser": {
		_const.RU: `лазерное`,
		_const.EN: `laser`,
	},
	"m3": {
		_const.RU: `м^3`,
		_const.EN: `m^3	`,
	},
	"km": {
		_const.RU: `км/с`,
		_const.EN: `km/s`,
	},
	"insec": {
		_const.RU: `в сек`,
		_const.EN: `per sec`,
	},
	"unitsec": {
		_const.RU: `ед/с`,
		_const.EN: `units/s`,
	},
	"percentsec": {
		_const.RU: `%/с`,
		_const.EN: `%/s`,
	},
	"s": {
		_const.RU: `c`,
		_const.EN: `s`,
	},
	"deg": {
		_const.RU: `гр/с`,
		_const.EN: `deg/s`,
	},
	"rate_fire": {
		_const.RU: `выс./мин`,
		_const.EN: `r/min`,
	},
	"ch_1": {
		_const.RU: `Занимаемый объем`,
		_const.EN: `Occupied volume`,
	},
	"ch_2": {
		_const.RU: `Вместимость`,
		_const.EN: `Capacity`,
	},
	"ch_3": {
		_const.RU: `Прочность`,
		_const.EN: `Strength`,
	},
	"ch_4": {
		_const.RU: `Занимаемый объем в трюме`,
		_const.EN: `Occupied volume in the hold`,
	},
	"ch_5": {
		_const.RU: `Защищен паролем`,
		_const.EN: `Password protected`,
	},
	"ch_6": {
		_const.RU: `Находится под землей`,
		_const.EN: `Located underground`,
	},

	"ch_7": {
		_const.RU: `Размер корпуса`,
		_const.EN: `Body size`,
	},
	"ch_8": {
		_const.RU: `Вместимость трюма`,
		_const.EN: `Hold capacity`,
	},
	"ch_9": {
		_const.RU: `Дополнительные отсеки:`,
		_const.EN: `Additional compartments:`,
	},
	"ch_10": {
		_const.RU: `Вместимость`,
		_const.EN: `Capacity`,
	},
	"ch_11": {
		_const.RU: `Допустимые предметы`,
		_const.EN: `Acceptable Items`,
	},
	"ch_12": {
		_const.RU: `Ходовая`,
		_const.EN: `Running gear`,
	},
	"ch_13": {
		_const.RU: `Макс. скорость`,
		_const.EN: `Max speed`,
	},
	"ch_14": {
		_const.RU: `Скорость поворота`,
		_const.EN: `Turning speed`,
	},
	"ch_15": {
		_const.RU: `Навигация`,
		_const.EN: `Navigation`,
	},
	"ch_16": {
		_const.RU: `Дальность обзора`,
		_const.EN: `View range`,
	},
	"ch_17": {
		_const.RU: `Дальность радара`,
		_const.EN: `Radar range`,
	},
	"ch_18": {
		_const.RU: `Накопитель`,
		_const.EN: `Battery`,
	},
	"ch_19": {
		_const.RU: `Вместимость накопителя`,
		_const.EN: `Energy storage capacity`,
	},
	"ch_20": {
		_const.RU: `Скорость зарядки`,
		_const.EN: `Charging speed`,
	},
	"ch_21": {
		_const.RU: `Выживаемость:`,
		_const.EN: `Survival:`,
	},
	"ch_22": {
		_const.RU: `Структура`,
		_const.EN: `Structure`,
	},
	"ch_23": {
		_const.RU: `Защита:`,
		_const.EN: `Protection:`,
	},
	"ch_24": {
		_const.RU: `Броня и уязвимые места`,
		_const.EN: `Armor and vulnerabilities`,
	},
	"ch_164": {
		_const.RU: `Щит`,
		_const.EN: `Shield`,
	},
	"ch_165": {
		_const.RU: `Защита щита`,
		_const.EN: `Shield protection`,
	},
	"ch_166": {
		_const.RU: `Отсутствует`,
		_const.EN: `Missing`,
	},
	"ch_25": {
		_const.RU: `Размер боеприпаса`,
		_const.EN: `Ammunition size`,
	},
	"ch_26": {
		_const.RU: `Тип оружия`,
		_const.EN: `Weapon type`,
	},
	"ch_27": {
		_const.RU: `Урон`,
		_const.EN: `Damage`,
	},
	"ch_28": {
		_const.RU: `Тип атаки`,
		_const.EN: `Attack type`,
	},
	"ch_29": {
		_const.RU: `Радиус поражения`,
		_const.EN: `Damage radius`,
	},
	"ch_30": {
		_const.RU: `Скорость полета пули`,
		_const.EN: `Bullet speed`,
	},
	"ch_31": {
		_const.RU: `Занимаемый объем`,
		_const.EN: `Occupied volume`,
	},
	"ch_32": {
		_const.RU: `Особенность:`,
		_const.EN: `Feature:`,
	},
	"ch_33": {
		_const.RU: `Ручной захват цели`,
		_const.EN: `Target lock`,
	},
	"ch_34": {
		_const.RU: `Управляемый снаряд`,
		_const.EN: `Guided projectile`,
	},
	"ch_35": {
		_const.RU: `Автоматический захват цели`,
		_const.EN: `Automatic target lock`,
	},
	"ch_36": {
		_const.RU: `Макс. прочность`,
		_const.EN: `Max strength`,
	},
	"ch_37": {
		_const.RU: `Радиус щита`,
		_const.EN: `Shield radius`,
	},
	"ch_38": {
		_const.RU: `Восстановление прочности`,
		_const.EN: `Restoring Strength`,
	},
	"ch_39": {
		_const.RU: `Потребляемая энергия`,
		_const.EN: `Energy consumption`,
	},
	"ch_40": {
		_const.RU: `Потребляемая энергия в простое`,
		_const.EN: `Energy consumption when idle`,
	},
	"ch_41": {
		_const.RU: `Время восстановления после урона`,
		_const.EN: `Recovery time after damage`,
	},
	"ch_42": {
		_const.RU: `Радиус мемех`,
		_const.EN: `Interference radius`,
	},
	"ch_43": {
		_const.RU: `Потребляемая энергия`,
		_const.EN: `Energy consumption`,
	},
	"ch_44": {
		_const.RU: `Радиус защиты`,
		_const.EN: `Protection radius`,
	},
	"ch_45": {
		_const.RU: `Потребляемая энергия за сбитие`,
		_const.EN: `Energy consumption per shoot down`,
	},
	"ch_46": {
		_const.RU: `Время перезарядки`,
		_const.EN: `Recharge time`,
	},
	"ch_47": {
		_const.RU: `Дальность ремонта`,
		_const.EN: `Repair range`,
	},
	"ch_48": {
		_const.RU: `Время перезарядки`,
		_const.EN: `Recharge time`,
	},
	"ch_49": {
		_const.RU: `Потребляемая энергия`,
		_const.EN: `Energy consumption`,
	},
	"ch_50": {
		_const.RU: `Кол-во восстанавливаемой прочности`,
		_const.EN: `Amount of restored strength`,
	},
	"ch_51": {
		_const.RU: `Объем добываемой руды`,
		_const.EN: `Volume of ore mined`,
	},
	"ch_52": {
		_const.RU: `Разработка руды`,
		_const.EN: `Ore development`,
	},
	"ch_53": {
		_const.RU: `Дальность добычи`,
		_const.EN: `Mining range`,
	},
	"ch_54": {
		_const.RU: `Потребление энергии`,
		_const.EN: `Power consumption`,
	},
	"ch_55": {
		_const.RU: `Наполнение бочки`,
		_const.EN: `Filling the barrel`,
	},
	"ch_56": {
		_const.RU: `Дальность шланга`,
		_const.EN: `Hose range`,
	},
	"ch_57": {
		_const.RU: `Скорость перезарядки`,
		_const.EN: `Recharge speed`,
	},
	"ch_58": {
		_const.RU: `Потребление энергии за такт`,
		_const.EN: `Energy consumption per cycle`,
	},
	"ch_59": {
		_const.RU: `Дальность управления`,
		_const.EN: `Control range`,
	},
	"ch_60": {
		_const.RU: `Радиус раскопок`,
		_const.EN: `Excavation radius`,
	},
	"ch_61": {
		_const.RU: `Перезарядка дрона`,
		_const.EN: `Recharging the drone`,
	},
	"ch_62": {
		_const.RU: `Энергия для активации дрона`,
		_const.EN: `Energy to activate the drone`,
	},
	"ch_63": {
		_const.RU: `Дальность`,
		_const.EN: `Range`,
	},
	"ch_64": {
		_const.RU: `Скорость перезарядки`,
		_const.EN: `Recharge speed`,
	},
	"ch_65": {
		_const.RU: `Потребление энергии`,
		_const.EN: `Power consumption`,
	},
	"ch_66": {
		_const.RU: `Влияние на прочность`,
		_const.EN: `Impact on strength`,
	},
	"ch_67": {
		_const.RU: `Дальность передачи энергии`,
		_const.EN: `Energy transmission range`,
	},
	"ch_68": {
		_const.RU: `Передаваемая энергия за такт`,
		_const.EN: `Transmitted energy per cycle`,
	},
	"ch_69": {
		_const.RU: `Скорость перезарядки`,
		_const.EN: `Recharge speed`,
	},
	"ch_70": {
		_const.RU: `Чувствительность датчика`,
		_const.EN: `Sensor sensitivity`,
	},
	"ch_71": {
		_const.RU: `Потребление энергии за сканирование`,
		_const.EN: `Energy consumption per scan`,
	},
	"ch_72": {
		_const.RU: `Перезарядка`,
		_const.EN: `Reload`,
	},
	"ch_73": {
		_const.RU: `Урон`,
		_const.EN: `Damage`,
	},
	"ch_74": {
		_const.RU: `Радиус поражения`,
		_const.EN: `Damage radius`,
	},
	"ch_75": {
		_const.RU: `Задержка детонации после установки`,
		_const.EN: `Detonation delay after installation`,
	},
	"ch_76": {
		_const.RU: `Задержка детонации после срабатывания`,
		_const.EN: `Detonation delay after actuation`,
	},
	"ch_77": {
		_const.RU: `Радиус обнаружения цели`,
		_const.EN: `Target detection radius`,
	},
	"ch_78": {
		_const.RU: `Кол-во зарядов`,
		_const.EN: `Number of charges`,
	},
	"ch_79": {
		_const.RU: `Потребляемая энергия за использование`,
		_const.EN: `Energy consumption per use`,
	},
	"ch_80": {
		_const.RU: `Время перезарядки`,
		_const.EN: `Recharge time`,
	},
	"ch_81": {
		_const.RU: `Дальность обзора`,
		_const.EN: `View range`,
	},
	"ch_82": {
		_const.RU: `Структура`,
		_const.EN: `Structure`,
	},
	"ch_83": {
		_const.RU: `Время жизни`,
		_const.EN: `Lifetime`,
	},
	"ch_84": {
		_const.RU: `Дальность полета капсулы с турелью`,
		_const.EN: `Flight range of a capsule with a turret`,
	},
	"ch_85": {
		_const.RU: `Кол-во зарядов`,
		_const.EN: `Number of charges`,
	},
	"ch_86": {
		_const.RU: `Потребляемая энергия за использование`,
		_const.EN: `Energy consumption per use`,
	},
	"ch_87": {
		_const.RU: `Время перезарядки`,
		_const.EN: `Recharge time`,
	},
	"ch_88": {
		_const.RU: `Радиус закрытия обзора одного заряда`,
		_const.EN: `Radius of closing the view of one charge`,
	},
	"ch_89": {
		_const.RU: `Кол-во зарядов`,
		_const.EN: `Number of charges`,
	},
	"ch_90": {
		_const.RU: `Потребляемая энергия за использование`,
		_const.EN: `Energy consumption per use`,
	},
	"ch_91": {
		_const.RU: `Время перезарядки`,
		_const.EN: `Recharge time`,
	},
	"ch_92": {
		_const.RU: `Радиус закрытия обзора одного заряда`,
		_const.EN: `Radius of closing the view of one charge`,
	},
	"ch_93": {
		_const.RU: `Дальность полета капсул`,
		_const.EN: `Capsule flight range`,
	},
	"ch_94": {
		_const.RU: `Кол-во зарядов`,
		_const.EN: `Number of charges`,
	},
	"ch_95": {
		_const.RU: `Кол-во капсул в одном заряде`,
		_const.EN: `Number of capsules in one charge`,
	},
	"ch_96": {
		_const.RU: `Потребляемая энергия за использование`,
		_const.EN: `Energy consumption per use`,
	},
	"ch_97": {
		_const.RU: `Время перезарядки`,
		_const.EN: `Recharge time`,
	},
	"ch_98": {
		_const.RU: `Дальность обзора`,
		_const.EN: `View range`,
	},
	"ch_99": {
		_const.RU: `Время жизни`,
		_const.EN: `Lifetime`,
	},
	"ch_100": {
		_const.RU: `Потребляемая энергия за использование`,
		_const.EN: `Energy consumption per use`,
	},
	"ch_101": {
		_const.RU: `Время перезарядки`,
		_const.EN: `Recharge time`,
	},
	"ch_102": {
		_const.RU: `Дальность обзора`,
		_const.EN: `View range`,
	},
	"ch_103": {
		_const.RU: `Время жизни`,
		_const.EN: `Lifetime`,
	},
	"ch_104": {
		_const.RU: `Потребляемая энергия за использование`,
		_const.EN: `Energy consumption per use`,
	},
	"ch_105": {
		_const.RU: `Время перезарядки`,
		_const.EN: `Recharge time`,
	},
	"ch_106": {
		_const.RU: `Дальность захвата целей`,
		_const.EN: `Target acquisition range`,
	},
	"ch_107": {
		_const.RU: `Структура`,
		_const.EN: `Structure`,
	},
	"ch_108": {
		_const.RU: `Время жизни ловушки`,
		_const.EN: `Trap lifetime`,
	},
	"ch_109": {
		_const.RU: `Время жизни ловушки после активации`,
		_const.EN: `Trap lifetime after activation`,
	},
	"ch_110": {
		_const.RU: `Кол-во зарядов`,
		_const.EN: `Number of charges`,
	},
	"ch_111": {
		_const.RU: `Потребляемая энергия за использование`,
		_const.EN: `Energy consumption per use`,
	},
	"ch_112": {
		_const.RU: `Время перезарядки`,
		_const.EN: `Recharge time`,
	},
	"ch_113": {
		_const.RU: `Дальность полета`,
		_const.EN: `Range of flight`,
	},
	"ch_114": {
		_const.RU: `Кол-во зарядов`,
		_const.EN: `Number of charges`,
	},
	"ch_115": {
		_const.RU: `Потребляемая энергия за использование`,
		_const.EN: `Energy consumption per use`,
	},
	"ch_116": {
		_const.RU: `Время перезарядки`,
		_const.EN: `Recharge time`,
	},
	"ch_117": {
		_const.RU: `Дальность полета`,
		_const.EN: `Range of flight`,
	},
	"ch_118": {
		_const.RU: `Сила притяжения`,
		_const.EN: `Force of gravity`,
	},
	"ch_119": {
		_const.RU: `Радиус`,
		_const.EN: `Radius`,
	},
	"ch_120": {
		_const.RU: `Время жизни`,
		_const.EN: `Lifetime`,
	},
	"ch_121": {
		_const.RU: `Кол-во зарядов`,
		_const.EN: `Number of charges`,
	},
	"ch_122": {
		_const.RU: `Потребляемая энергия за использование`,
		_const.EN: `Energy consumption per use`,
	},
	"ch_123": {
		_const.RU: `Время перезарядки`,
		_const.EN: `Recharge time`,
	},
	"ch_124": {
		_const.RU: `Сила прыжка`,
		_const.EN: `Jump Power`,
	},
	"ch_125": {
		_const.RU: `Потребляемая энергия за использование`,
		_const.EN: `Energy consumption per use`,
	},
	"ch_126": {
		_const.RU: `Время перезарядки`,
		_const.EN: `Recharge time`,
	},
	"ch_127": {
		_const.RU: `Время активации`,
		_const.EN: `Activation time`,
	},
	"ch_128": {
		_const.RU: `Потребляемая энергия`,
		_const.EN: `Energy consumption`,
	},
	"ch_129": {
		_const.RU: `Время перезарядки`,
		_const.EN: `Recharge time`,
	},
	"ch_130": {
		_const.RU: `Структура`,
		_const.EN: `Structure`,
	},
	"ch_131": {
		_const.RU: `Дальность полета капсулы со стеной`,
		_const.EN: `Flight range of a capsule with a wall`,
	},
	"ch_132": {
		_const.RU: `Кол-во зарядов`,
		_const.EN: `Number of charges`,
	},
	"ch_133": {
		_const.RU: `Потребляемая энергия за использование`,
		_const.EN: `Energy consumption per use`,
	},
	"ch_134": {
		_const.RU: `Время перезарядки`,
		_const.EN: `Recharge time`,
	},
	"ch_135": {
		_const.RU: `Макс. прочность`,
		_const.EN: `Max strength`,
	},
	"ch_136": {
		_const.RU: `Радиус щита`,
		_const.EN: `Shield radius`,
	},
	"ch_137": {
		_const.RU: `Время жизни`,
		_const.EN: `Lifetime`,
	},
	"ch_138": {
		_const.RU: `Восстановление прочности`,
		_const.EN: `Restoring Strength`,
	},
	"ch_139": {
		_const.RU: `Время восстановления после урона`,
		_const.EN: `Recovery time after damage`,
	},
	"ch_140": {
		_const.RU: `Структура`,
		_const.EN: `Structure`,
	},
	"ch_141": {
		_const.RU: `Дальность полета капсулы`,
		_const.EN: `Capsule flight range`,
	},
	"ch_142": {
		_const.RU: `Кол-во зарядов`,
		_const.EN: `Number of charges`,
	},
	"ch_143": {
		_const.RU: `Потребляемая энергия за использование`,
		_const.EN: `Energy consumption per use`,
	},
	"ch_144": {
		_const.RU: `Время перезарядки`,
		_const.EN: `Recharge time`,
	},
	"ch_145": {
		_const.RU: `Радиус`,
		_const.EN: `Radius`,
	},
	"ch_146": {
		_const.RU: `Потребляемая энергия за использование`,
		_const.EN: `Energy consumption per use`,
	},
	"ch_147": {
		_const.RU: `Потребляя мощность`,
		_const.EN: `Consuming power`,
	},
	"ch_148": {
		_const.RU: `Занимаемый объем`,
		_const.EN: `Occupied volume`,
	},
	"ch_149": {
		_const.RU: `Размер оружия`,
		_const.EN: `Weapon size`,
	},
	"ch_150": {
		_const.RU: `Тип оружия`,
		_const.EN: `Weapon type`,
	},
	"ch_151": {
		_const.RU: `Скорость поворота`,
		_const.EN: `Swing speed`,
	},
	"ch_152": {
		_const.RU: `Модификатор урона`,
		_const.EN: `Damage modifier`,
	},
	"ch_153": {
		_const.RU: `Доступные углы атаки`,
		_const.EN: `Available angles of attack`,
	},
	"ch_154": {
		_const.RU: `Влияние на скорость пули`,
		_const.EN: `Effect on bullet speed`,
	},
	"ch_155": {
		_const.RU: `Дальность стрельбы`,
		_const.EN: `Range attack`,
	},
	"ch_156": {
		_const.RU: `зависит от снаряда`,
		_const.EN: `depends on the projectile`,
	},
	"ch_157": {
		_const.RU: `Очередь`,
		_const.EN: `Queue`,
	},
	"ch_158": {
		_const.RU: `Скорострельность`,
		_const.EN: `Rate of fire`,
	},
	"ch_159": {
		_const.RU: `Боезапас`,
		_const.EN: `Ammunition`,
	},
	"ch_160": {
		_const.RU: `Время перезарядки`,
		_const.EN: `Reload time`,
	},
	"ch_161": {
		_const.RU: `Потребляя мощность`,
		_const.EN: `Consuming power`,
	},
	"ch_162": {
		_const.RU: `Занимаемый объем`,
		_const.EN: `Occupied volume`,
	},
	"ch_163": {
		_const.RU: `артиллерия`,
		_const.EN: `artillery`,
	},
	// DetailItem
	"text_1": {
		_const.RU: `примерная цена: <span class="importantly">%credits_count%</span> cr.`,
		_const.EN: `approximate price: <span class="importantly">%credits_count%</span> cr.`,
	},
	"text_2": {
		_const.RU: `Описание`,
		_const.EN: `Description`,
	},
	"text_3": {
		_const.RU: `Особенности`,
		_const.EN: `Peculiarities`,
	},
	"text_4": {
		_const.RU: `Хар-ки`,
		_const.EN: `Specs`,
	},
	"text_5": {
		_const.RU: `Хар-ки оружия`,
		_const.EN: `Weapon specs.`,
	},
	"text_6": {
		_const.RU: `Хар-ки снаряда`,
		_const.EN: `Projectile specs.`,
	},
	"text_7": {
		_const.RU: `Варианты`,
		_const.EN: `Options`,
	},
	"text_8": {
		_const.RU: `Требования`,
		_const.EN: `Requirements`,
	},
	"text_9": {
		_const.RU: `Оснащение`,
		_const.EN: `Equipment`,
	},
	"text_10": {
		_const.RU: `Произ-во`,
		_const.EN: `Production`,
	},
	"broken_tip": {
		_const.RU: `Предмет сломан из за этого:<br>
          - его нельзя переработать или продать<br>
          - если выкинуть из трюма он уничтожиться<br>
          - его характеристики могут быть ухудшены<br>`,
		_const.EN: `The item is broken, which means that:<br>
          - it cannot be recycled or sold<br>
          - if you throw it away, it will be destroyed<br>
          - its characteristics may be reduced<br>`,
	},
	"resource_tip_1_1": {
		_const.RU: `Этот ресурс можно получить`,
		_const.EN: `This resource can be obtained by `,
	},
	"resource_tip_1_2": {
		_const.RU: `переработав`,
		_const.EN: `processing`,
	},
	"resource_tip_2_1": {
		_const.RU: `Этот ресурс можно`,
		_const.EN: `This resource can be`,
	},
	"resource_tip_2_2": {
		_const.RU: `добыть в пустошах`,
		_const.EN: `mined in the wasteland`,
	},
	"blueprints_tip_1_1": {
		_const.RU: `Чертежи можно получить из обломков или исследуя`,
		_const.EN: `Blueprints can be obtained from debris or by exploring`,
	},
	"blueprints_tip_1_2": {
		_const.RU: `аномалии в пустошах`,
		_const.EN: `anomalies in the wasteland`,
	},
	"detail_tip_1": {
		_const.RU: `Детали можно получить:`,
		_const.EN: `Details can be obtained:`,
	},
	"detail_tip_1_1": {
		_const.RU: `создав их на заводе полуфабрикатов`,
		_const.EN: `creating them at a semi-finished products factory`,
	},
	"detail_tip_1_2": {
		_const.RU: `создав их по чертежам`,
		_const.EN: `creating them according to drawings`,
	},
	"detail_tip_1_3": {
		_const.RU: `переработав обломки`,
		_const.EN: `recycling the wreckage`,
	},
	// EquipmentTab
	"text_11": {
		_const.RU: `Ячейки для топлива`,
		_const.EN: `Fuel cells`,
	},
	"text_12": {
		_const.RU: `Лимит мощности`,
		_const.EN: `Power limit`,
	},
	"text_13": {
		_const.RU: `Слотов оружия`,
		_const.EN: `Weapon slots`,
	},
	"text_14": {
		_const.RU: `Разъемы большой мощности`,
		_const.EN: `High Power Connectors`,
	},
	"text_15": {
		_const.RU: `Разъемы средней мощности`,
		_const.EN: `Medium Power Connectors`,
	},
	"text_16": {
		_const.RU: `Разъемы малой мощности`,
		_const.EN: `Low Power Connectors`,
	},
	"text_17": {
		_const.RU: `Боеприпасы`,
		_const.EN: `Ammunition`,
	},
	// ProductionTab
	"text_18": {
		_const.RU: `Продукция`,
		_const.EN: `Products`,
	},
	"text_19": {
		_const.RU: `Требуемые ресурсы <span style="font-size: 9px">(при 100% эф.)</span>`,
		_const.EN: `Required resources <span style="font-size: 9px">(at 100% ef.)</span>`,
	},
	"text_20": {
		_const.RU: `Можно получить переработав`,
		_const.EN: `Can be obtained by recycling`,
	},
	"text_21": {
		_const.RU: `Чертежи`,
		_const.EN: `Blueprints`,
	},
	"text_22": {
		_const.RU: `Компоненты после переработки`,
		_const.EN: `Components after processing`,
	},
	"text_24": {
		_const.RU: `Переработав <span style="color: yellow">%count%</span> ед. выйдет:`,
		_const.EN: `Having processed <span style="color: yellow">%count%</span> units. will come out:`,
	},
	// Requirements
	"text_23": {
		_const.RU: `Навыки`,
		_const.EN: `Skills`,
	},
}
