package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var DetailItem = map[string]map[string]string{
	// BonusTab
	"button_1": {
		_const.RU:   `Навыки`,
		_const.EN:   `Skills`,
		_const.ZhCN: `技能`,
	},
	"text_25": {
		_const.RU:   `Профильные особенности`,
		_const.EN:   `Profile features`,
		_const.ZhCN: `配置文件特性`,
	},
	"text_26": {
		_const.RU:   `За каждые <span class="importantly">25%</span> синхронизации с транспортом:`,
		_const.EN:   `For every <span class="importantly">25%</span> synchronization with transport:`,
		_const.ZhCN: `每与运输工具同步 <span class="importantly">25%</span>：`,
	},
	// CharacteristicTab
	"small": {
		_const.RU:   `малый`,
		_const.EN:   `small`,
		_const.ZhCN: `小`,
	},
	"average": {
		_const.RU:   `средний`,
		_const.EN:   `average`,
		_const.ZhCN: `中`,
	},
	"large": {
		_const.RU:   `большой`,
		_const.EN:   `large`,
		_const.ZhCN: `大`,
	},
	"firearms": {
		_const.RU:   `огнестрельное`,
		_const.EN:   `firearms`,
		_const.ZhCN: `火器`,
	},
	"missile": {
		_const.RU:   `ракетное`,
		_const.EN:   `missile`,
		_const.ZhCN: `导弹`,
	},
	"laser": {
		_const.RU:   `лазерное`,
		_const.EN:   `laser`,
		_const.ZhCN: `激光`,
	},
	"melee": {
		_const.RU:   "ближнего боя",
		_const.EN:   "melee",
		_const.ZhCN: "近战",
	},
	"m3": {
		_const.RU:   `м^3`,
		_const.EN:   `m^3`,
		_const.ZhCN: `立方米`,
	},
	"km": {
		_const.RU:   `км/с`,
		_const.EN:   `km/s`,
		_const.ZhCN: `公里/秒`,
	},
	"insec": {
		_const.RU:   `в сек`,
		_const.EN:   `per sec`,
		_const.ZhCN: `每秒`,
	},
	"unitsec": {
		_const.RU:   `ед/с`,
		_const.EN:   `units/s`,
		_const.ZhCN: `单位/秒`,
	},
	"percentsec": {
		_const.RU:   `%/с`,
		_const.EN:   `%/s`,
		_const.ZhCN: `%/秒`,
	},
	"s": {
		_const.RU:   `c`,
		_const.EN:   `s`,
		_const.ZhCN: `秒`,
	},
	"deg": {
		_const.RU:   `гр/с`,
		_const.EN:   `deg/s`,
		_const.ZhCN: `度/秒`,
	},
	"rate_fire": {
		_const.RU:   `выс./мин`,
		_const.EN:   `r/min`,
		_const.ZhCN: `转/分钟`,
	},
	"ch_1": {
		_const.RU:   `Занимаемый объем`,
		_const.EN:   `Occupied volume`,
		_const.ZhCN: `占用体积`,
	},
	"ch_2": {
		_const.RU:   `Вместимость`,
		_const.EN:   `Capacity`,
		_const.ZhCN: `容量`,
	},
	"ch_3": {
		_const.RU:   `Прочность`,
		_const.EN:   `Strength`,
		_const.ZhCN: `强度`,
	},
	"ch_4": {
		_const.RU:   `Занимаемый объем в грузовом отсеке`,
		_const.EN:   `Occupied volume in cargo hold`,
		_const.ZhCN: `货舱占用容量`,
	},
	"ch_5": {
		_const.RU:   `Защищен паролем`,
		_const.EN:   `Password protected`,
		_const.ZhCN: `密码保护`,
	},
	"ch_6": {
		_const.RU:   `Находится под землей`,
		_const.EN:   `Located underground`,
		_const.ZhCN: `位于地下`,
	},
	"ch_7": {
		_const.RU:   `Размер корпуса`,
		_const.EN:   `Body size`,
		_const.ZhCN: `机身尺寸`,
	},
	"ch_8": {
		_const.RU:   `Вместимость грузового отсека`,
		_const.EN:   `Cargo hold capacity`,
		_const.ZhCN: `货舱容量`,
	},
	"ch_9": {
		_const.RU:   `Дополнительные отсеки:`,
		_const.EN:   `Additional compartments:`,
		_const.ZhCN: `附加舱室：`,
	},
	"ch_10": {
		_const.RU:   `Вместимость`,
		_const.EN:   `Capacity`,
		_const.ZhCN: `容量`,
	},
	"ch_11": {
		_const.RU:   `Допустимые предметы`,
		_const.EN:   `Acceptable Items`,
		_const.ZhCN: `可接受物品`,
	},
	"ch_12": {
		_const.RU:   `Ходовая`,
		_const.EN:   `Running gear`,
		_const.ZhCN: `行走机构`,
	},
	"ch_13": {
		_const.RU:   `Макс. скорость`,
		_const.EN:   `Max speed`,
		_const.ZhCN: `最大速度`,
	},
	"ch_14": {
		_const.RU:   `Скорость поворота`,
		_const.EN:   `Turning speed`,
		_const.ZhCN: `转弯速度`,
	},
	"ch_15": {
		_const.RU:   `Навигация`,
		_const.EN:   `Navigation`,
		_const.ZhCN: `导航`,
	},
	"ch_16": {
		_const.RU:   `Дальность обзора`,
		_const.EN:   `View range`,
		_const.ZhCN: `视野范围`,
	},
	"ch_17": {
		_const.RU:   `Дальность радара`,
		_const.EN:   `Radar range`,
		_const.ZhCN: `雷达范围`,
	},
	"ch_18": {
		_const.RU:   `Накопитель`,
		_const.EN:   `Battery`,
		_const.ZhCN: `电池`,
	},
	"ch_19": {
		_const.RU:   `Вместимость накопителя`,
		_const.EN:   `Energy storage capacity`,
		_const.ZhCN: `能量存储容量`,
	},
	"ch_20": {
		_const.RU:   `Скорость зарядки`,
		_const.EN:   `Charging speed`,
		_const.ZhCN: `充电速度`,
	},
	"ch_21": {
		_const.RU:   `Выживаемость:`,
		_const.EN:   `Survival:`,
		_const.ZhCN: `生存能力：`,
	},
	"ch_22": {
		_const.RU:   `Структура`,
		_const.EN:   `Structure`,
		_const.ZhCN: `结构`,
	},
	"ch_23": {
		_const.RU:   `Защита:`,
		_const.EN:   `Protection:`,
		_const.ZhCN: `保护：`,
	},
	"ch_24": {
		_const.RU:   `Броня и уязвимые места`,
		_const.EN:   `Armor and vulnerabilities`,
		_const.ZhCN: `装甲和弱点`,
	},
	"ch_164": {
		_const.RU:   `Щит`,
		_const.EN:   `Shield`,
		_const.ZhCN: `护盾`,
	},
	"ch_165": {
		_const.RU:   `Защита щита`,
		_const.EN:   `Shield protection`,
		_const.ZhCN: `护盾保护`,
	},
	"ch_166": {
		_const.RU:   `Отсутствует`,
		_const.EN:   `Missing`,
		_const.ZhCN: `缺失`,
	},
	"ch_25": {
		_const.RU:   `Размер боеприпаса`,
		_const.EN:   `Ammunition size`,
		_const.ZhCN: `弹药尺寸`,
	},
	"ch_26": {
		_const.RU:   `Тип оружия`,
		_const.EN:   `Weapon type`,
		_const.ZhCN: `武器类型`,
	},
	"ch_27": {
		_const.RU:   `Урон`,
		_const.EN:   `Damage`,
		_const.ZhCN: `伤害`,
	},
	"ch_28": {
		_const.RU:   `Тип атаки`,
		_const.EN:   `Attack type`,
		_const.ZhCN: `攻击类型`,
	},
	"ch_29": {
		_const.RU:   `Радиус поражения`,
		_const.EN:   `Damage radius`,
		_const.ZhCN: `伤害半径`,
	},
	"ch_30": {
		_const.RU:   `Скорость полета пули`,
		_const.EN:   `Bullet speed`,
		_const.ZhCN: `子弹速度`,
	},
	"ch_31": {
		_const.RU:   `Занимаемый объем`,
		_const.EN:   `Occupied volume`,
		_const.ZhCN: `占用体积`,
	},
	"ch_32": {
		_const.RU:   `Особенность:`,
		_const.EN:   `Feature:`,
		_const.ZhCN: `特性：`,
	},
	"ch_33": {
		_const.RU:   `Ручной захват цели`,
		_const.EN:   `Target lock`,
		_const.ZhCN: `目标锁定`,
	},
	"ch_34": {
		_const.RU:   `Управляемый снаряд`,
		_const.EN:   `Guided projectile`,
		_const.ZhCN: `制导炮弹`,
	},
	"ch_35": {
		_const.RU:   `Автоматический захват цели`,
		_const.EN:   `Automatic target lock`,
		_const.ZhCN: `自动目标锁定`,
	},
	"ch_36": {
		_const.RU:   `Макс. прочность`,
		_const.EN:   `Max strength`,
		_const.ZhCN: `最大强度`,
	},
	"ch_37": {
		_const.RU:   `Радиус щита`,
		_const.EN:   `Shield radius`,
		_const.ZhCN: `护盾半径`,
	},
	"ch_38": {
		_const.RU:   `Восстановление прочности`,
		_const.EN:   `Restoring Strength`,
		_const.ZhCN: `恢复强度`,
	},
	"ch_39": {
		_const.RU:   `Потребляемая энергия`,
		_const.EN:   `Energy consumption`,
		_const.ZhCN: `能量消耗`,
	},
	"ch_40": {
		_const.RU:   `Потребляемая энергия в простое`,
		_const.EN:   `Energy consumption when idle`,
		_const.ZhCN: `空闲时能量消耗`,
	},
	"ch_41": {
		_const.RU:   `Время восстановления после урона`,
		_const.EN:   `Recovery time after damage`,
		_const.ZhCN: `损坏后恢复时间`,
	},
	"ch_42": {
		_const.RU:   `Радиус мемех`,
		_const.EN:   `Interference radius`,
		_const.ZhCN: `干扰半径`,
	},
	"ch_43": {
		_const.RU:   `Потребляемая энергия`,
		_const.EN:   `Energy consumption`,
		_const.ZhCN: `能量消耗`,
	},
	"ch_44": {
		_const.RU:   `Радиус защиты`,
		_const.EN:   `Protection radius`,
		_const.ZhCN: `保护半径`,
	},
	"ch_45": {
		_const.RU:   `Потребляемая энергия за сбитие`,
		_const.EN:   `Energy consumption per shoot down`,
		_const.ZhCN: `每次击落能量消耗`,
	},
	"ch_46": {
		_const.RU:   `Время перезарядки`,
		_const.EN:   `Recharge time`,
		_const.ZhCN: `重新装填时间`,
	},
	"ch_47": {
		_const.RU:   `Дальность ремонта`,
		_const.EN:   `Repair range`,
		_const.ZhCN: `修复范围`,
	},
	"ch_48": {
		_const.RU:   `Время перезарядки`,
		_const.EN:   `Recharge time`,
		_const.ZhCN: `重新装填时间`,
	},
	"ch_49": {
		_const.RU:   `Потребляемая энергия`,
		_const.EN:   `Energy consumption`,
		_const.ZhCN: `能量消耗`,
	},
	"ch_50": {
		_const.RU:   `Кол-во восстанавливаемой прочности`,
		_const.EN:   `Amount of restored strength`,
		_const.ZhCN: `恢复强度量`,
	},
	"ch_51": {
		_const.RU:   `Объем добываемой руды`,
		_const.EN:   `Volume of ore mined`,
		_const.ZhCN: `开采矿石量`,
	},
	"ch_52": {
		_const.RU:   `Разработка руды`,
		_const.EN:   `Ore development`,
		_const.ZhCN: `矿石开发`,
	},
	"ch_53": {
		_const.RU:   `Дальность добычи`,
		_const.EN:   `Mining range`,
		_const.ZhCN: `采矿范围`,
	},
	"ch_54": {
		_const.RU:   `Потребление энергии`,
		_const.EN:   `Power consumption`,
		_const.ZhCN: `能量消耗`,
	},
	"ch_55": {
		_const.RU:   `Наполнение бочки`,
		_const.EN:   `Filling the barrel`,
		_const.ZhCN: `填充桶`,
	},
	"ch_56": {
		_const.RU:   `Дальность шланга`,
		_const.EN:   `Hose range`,
		_const.ZhCN: `软管范围`,
	},
	"ch_57": {
		_const.RU:   `Скорость перезарядки`,
		_const.EN:   `Recharge speed`,
		_const.ZhCN: `重新装填速度`,
	},
	"ch_58": {
		_const.RU:   `Потребление энергии за такт`,
		_const.EN:   `Energy consumption per cycle`,
		_const.ZhCN: `每周期能量消耗`,
	},
	"ch_59": {
		_const.RU:   `Дальность управления`,
		_const.EN:   `Control range`,
		_const.ZhCN: `控制范围`,
	},
	"ch_60": {
		_const.RU:   `Радиус раскопок`,
		_const.EN:   `Excavation radius`,
		_const.ZhCN: `挖掘半径`,
	},
	"ch_61": {
		_const.RU:   `Перезарядка дрона`,
		_const.EN:   `Recharging the drone`,
		_const.ZhCN: `无人机充电`,
	},
	"ch_62": {
		_const.RU:   `Энергия для активации дрона`,
		_const.EN:   `Energy to activate the drone`,
		_const.ZhCN: `激活无人机能量`,
	},
	"ch_63": {
		_const.RU:   `Дальность`,
		_const.EN:   `Range`,
		_const.ZhCN: `范围`,
	},
	"ch_64": {
		_const.RU:   `Скорость перезарядки`,
		_const.EN:   `Recharge speed`,
		_const.ZhCN: `重新装填速度`,
	},
	"ch_65": {
		_const.RU:   `Потребление энергии`,
		_const.EN:   `Power consumption`,
		_const.ZhCN: `能量消耗`,
	},
	"ch_66": {
		_const.RU:   `Влияние на прочность`,
		_const.EN:   `Impact on strength`,
		_const.ZhCN: `对强度的影响`,
	},
	"ch_67": {
		_const.RU:   `Дальность передачи энергии`,
		_const.EN:   `Energy transmission range`,
		_const.ZhCN: `能量传输范围`,
	},
	"ch_68": {
		_const.RU:   `Передаваемая энергия за такт`,
		_const.EN:   `Transmitted energy per cycle`,
		_const.ZhCN: `每周期传输能量`,
	},
	"ch_69": {
		_const.RU:   `Скорость перезарядки`,
		_const.EN:   `Recharge speed`,
		_const.ZhCN: `重新装填速度`,
	},
	"ch_70": {
		_const.RU:   `Чувствительность датчика`,
		_const.EN:   `Sensor sensitivity`,
		_const.ZhCN: `传感器灵敏度`,
	},
	"ch_71": {
		_const.RU:   `Потребление энергии за сканирование`,
		_const.EN:   `Energy consumption per scan`,
		_const.ZhCN: `每次扫描的能量消耗`,
	},
	"ch_72": {
		_const.RU:   `Перезарядка`,
		_const.EN:   `Reload`,
		_const.ZhCN: `重新装填`,
	},
	"ch_73": {
		_const.RU:   `Урон`,
		_const.EN:   `Damage`,
		_const.ZhCN: `伤害`,
	},
	"ch_74": {
		_const.RU:   `Радиус поражения`,
		_const.EN:   `Damage radius`,
		_const.ZhCN: `伤害半径`,
	},
	"ch_75": {
		_const.RU:   `Задержка детонации после установки`,
		_const.EN:   `Detonation delay after installation`,
		_const.ZhCN: `安装后引爆延迟`,
	},
	"ch_76": {
		_const.RU:   `Задержка детонации после срабатывания`,
		_const.EN:   `Detonation delay after actuation`,
		_const.ZhCN: `触发后引爆延迟`,
	},
	"ch_77": {
		_const.RU:   `Радиус обнаружения цели`,
		_const.EN:   `Target detection radius`,
		_const.ZhCN: `目标探测半径`,
	},
	"ch_78": {
		_const.RU:   `Кол-во зарядов`,
		_const.EN:   `Number of charges`,
		_const.ZhCN: `充能次数`,
	},
	"ch_79": {
		_const.RU:   `Потребляемая энергия за использование`,
		_const.EN:   `Energy consumption per use`,
		_const.ZhCN: `每次使用的能量消耗`,
	},
	"ch_80": {
		_const.RU:   `Время перезарядки`,
		_const.EN:   `Recharge time`,
		_const.ZhCN: `重新充能时间`,
	},
	"ch_81": {
		_const.RU:   `Дальность обзора`,
		_const.EN:   `View range`,
		_const.ZhCN: `视野范围`,
	},
	"ch_82": {
		_const.RU:   `Структура`,
		_const.EN:   `Structure`,
		_const.ZhCN: `结构`,
	},
	"ch_83": {
		_const.RU:   `Время жизни`,
		_const.EN:   `Lifetime`,
		_const.ZhCN: `寿命`,
	},
	"ch_84": {
		_const.RU:   `Дальность полета капсулы с турелью`,
		_const.EN:   `Flight range of a capsule with a turret`,
		_const.ZhCN: `带炮塔的胶囊飞行范围`,
	},
	"ch_85": {
		_const.RU:   `Кол-во зарядов`,
		_const.EN:   `Number of charges`,
		_const.ZhCN: `充能次数`,
	},
	"ch_86": {
		_const.RU:   `Потребляемая энергия за использование`,
		_const.EN:   `Energy consumption per use`,
		_const.ZhCN: `每次使用的能量消耗`,
	},
	"ch_87": {
		_const.RU:   `Время перезарядки`,
		_const.EN:   `Recharge time`,
		_const.ZhCN: `重新充能时间`,
	},
	"ch_88": {
		_const.RU:   `Радиус закрытия обзора одного заряда`,
		_const.EN:   `Radius of closing the view of one charge`,
		_const.ZhCN: `单次充能的视野关闭半径`,
	},
	"ch_89": {
		_const.RU:   `Кол-во зарядов`,
		_const.EN:   `Number of charges`,
		_const.ZhCN: `充能次数`,
	},
	"ch_90": {
		_const.RU:   `Потребляемая энергия за использование`,
		_const.EN:   `Energy consumption per use`,
		_const.ZhCN: `每次使用的能量消耗`,
	},
	"ch_91": {
		_const.RU:   `Время перезарядки`,
		_const.EN:   `Recharge time`,
		_const.ZhCN: `重新充能时间`,
	},
	"ch_92": {
		_const.RU:   `Радиус закрытия обзора одного заряда`,
		_const.EN:   `Radius of closing the view of one charge`,
		_const.ZhCN: `单次充能的视野关闭半径`,
	},
	"ch_93": {
		_const.RU:   `Дальность полета капсул`,
		_const.EN:   `Capsule flight range`,
		_const.ZhCN: `胶囊飞行范围`,
	},
	"ch_94": {
		_const.RU:   `Кол-во зарядов`,
		_const.EN:   `Number of charges`,
		_const.ZhCN: `充能次数`,
	},
	"ch_95": {
		_const.RU:   `Кол-во капсул в одном заряде`,
		_const.EN:   `Number of capsules in one charge`,
		_const.ZhCN: `单次充能的胶囊数量`,
	},
	"ch_96": {
		_const.RU:   `Потребляемая энергия за использование`,
		_const.EN:   `Energy consumption per use`,
		_const.ZhCN: `每次使用的能量消耗`,
	},
	"ch_97": {
		_const.RU:   `Время перезарядки`,
		_const.EN:   `Recharge time`,
		_const.ZhCN: `重新充能时间`,
	},
	"ch_98": {
		_const.RU:   `Дальность обзора`,
		_const.EN:   `View range`,
		_const.ZhCN: `视野范围`,
	},
	"ch_99": {
		_const.RU:   `Время жизни`,
		_const.EN:   `Lifetime`,
		_const.ZhCN: `寿命`,
	},
	"ch_100": {
		_const.RU:   `Потребляемая энергия за использование`,
		_const.EN:   `Energy consumption per use`,
		_const.ZhCN: `每次使用的能量消耗`,
	},
	"ch_101": {
		_const.RU:   `Время перезарядки`,
		_const.EN:   `Recharge time`,
		_const.ZhCN: `重新充能时间`,
	},
	"ch_102": {
		_const.RU:   `Дальность обзора`,
		_const.EN:   `View range`,
		_const.ZhCN: `视野范围`,
	},
	"ch_103": {
		_const.RU:   `Время жизни`,
		_const.EN:   `Lifetime`,
		_const.ZhCN: `寿命`,
	},
	"ch_104": {
		_const.RU:   `Потребляемая энергия за использование`,
		_const.EN:   `Energy consumption per use`,
		_const.ZhCN: `每次使用的能量消耗`,
	},
	"ch_105": {
		_const.RU:   `Время перезарядки`,
		_const.EN:   `Recharge time`,
		_const.ZhCN: `重新充能时间`,
	},
	"ch_106": {
		_const.RU:   `Дальность захвата целей`,
		_const.EN:   `Target acquisition range`,
		_const.ZhCN: `目标捕获范围`,
	},
	"ch_107": {
		_const.RU:   `Структура`,
		_const.EN:   `Structure`,
		_const.ZhCN: `结构`,
	},
	"ch_108": {
		_const.RU:   `Время жизни ловушки`,
		_const.EN:   `Trap lifetime`,
		_const.ZhCN: `陷阱寿命`,
	},
	"ch_109": {
		_const.RU:   `Время жизни ловушки после активации`,
		_const.EN:   `Trap lifetime after activation`,
		_const.ZhCN: `激活后陷阱寿命`,
	},
	"ch_110": {
		_const.RU:   `Кол-во зарядов`,
		_const.EN:   `Number of charges`,
		_const.ZhCN: `充能次数`,
	},
	"ch_111": {
		_const.RU:   `Потребляемая энергия за использование`,
		_const.EN:   `Energy consumption per use`,
		_const.ZhCN: `每次使用的能量消耗`,
	},
	"ch_112": {
		_const.RU:   `Время перезарядки`,
		_const.EN:   `Recharge time`,
		_const.ZhCN: `重新充能时间`,
	},
	"ch_113": {
		_const.RU:   `Дальность полета`,
		_const.EN:   `Range of flight`,
		_const.ZhCN: `飞行范围`,
	},
	"ch_114": {
		_const.RU:   `Кол-во зарядов`,
		_const.EN:   `Number of charges`,
		_const.ZhCN: `充能次数`,
	},
	"ch_115": {
		_const.RU:   `Потребляемая энергия за использование`,
		_const.EN:   `Energy consumption per use`,
		_const.ZhCN: `每次使用的能量消耗`,
	},
	"ch_116": {
		_const.RU:   `Время перезарядки`,
		_const.EN:   `Recharge time`,
		_const.ZhCN: `重新充能时间`,
	},
	"ch_117": {
		_const.RU:   `Дальность полета`,
		_const.EN:   `Range of flight`,
		_const.ZhCN: `飞行范围`,
	},
	"ch_118": {
		_const.RU:   `Сила притяжения`,
		_const.EN:   `Force of gravity`,
		_const.ZhCN: `重力强度`,
	},
	"ch_119": {
		_const.RU:   `Радиус`,
		_const.EN:   `Radius`,
		_const.ZhCN: `半径`,
	},
	"ch_120": {
		_const.RU:   `Время жизни`,
		_const.EN:   `Lifetime`,
		_const.ZhCN: `寿命`,
	},
	"ch_121": {
		_const.RU:   `Кол-во зарядов`,
		_const.EN:   `Number of charges`,
		_const.ZhCN: `充能次数`,
	},
	"ch_122": {
		_const.RU:   `Потребляемая энергия за использование`,
		_const.EN:   `Energy consumption per use`,
		_const.ZhCN: `每次使用的能量消耗`,
	},
	"ch_123": {
		_const.RU:   `Время перезарядки`,
		_const.EN:   `Recharge time`,
		_const.ZhCN: `重新充能时间`,
	},
	"ch_124": {
		_const.RU:   `Сила прыжка`,
		_const.EN:   `Jump Power`,
		_const.ZhCN: `跳跃力量`,
	},
	"ch_125": {
		_const.RU:   `Потребляемая энергия за использование`,
		_const.EN:   `Energy consumption per use`,
		_const.ZhCN: `每次使用的能量消耗`,
	},
	"ch_126": {
		_const.RU:   `Время перезарядки`,
		_const.EN:   `Recharge time`,
		_const.ZhCN: `重新充能时间`,
	},
	"ch_127": {
		_const.RU:   `Время активации`,
		_const.EN:   `Activation time`,
		_const.ZhCN: `激活时间`,
	},
	"ch_128": {
		_const.RU:   `Потребляемая энергия`,
		_const.EN:   `Energy consumption`,
		_const.ZhCN: `能量消耗`,
	},
	"ch_129": {
		_const.RU:   `Время перезарядки`,
		_const.EN:   `Recharge time`,
		_const.ZhCN: `重新充能时间`,
	},
	"ch_130": {
		_const.RU:   `Структура`,
		_const.EN:   `Structure`,
		_const.ZhCN: `结构`,
	},
	"ch_131": {
		_const.RU:   `Дальность полета капсулы со стеной`,
		_const.EN:   `Flight range of a capsule with a wall`,
		_const.ZhCN: `带墙的胶囊飞行范围`,
	},
	"ch_132": {
		_const.RU:   `Кол-во зарядов`,
		_const.EN:   `Number of charges`,
		_const.ZhCN: `充能次数`,
	},
	"ch_133": {
		_const.RU:   `Потребляемая энергия за использование`,
		_const.EN:   `Energy consumption per use`,
		_const.ZhCN: `每次使用的能量消耗`,
	},
	"ch_134": {
		_const.RU:   `Время перезарядки`,
		_const.EN:   `Recharge time`,
		_const.ZhCN: `重新充能时间`,
	},
	"ch_135": {
		_const.RU:   `Макс. прочность`,
		_const.EN:   `Max strength`,
		_const.ZhCN: `最大强度`,
	},
	"ch_136": {
		_const.RU:   `Радиус щита`,
		_const.EN:   `Shield radius`,
		_const.ZhCN: `护盾半径`,
	},
	"ch_137": {
		_const.RU:   `Время жизни`,
		_const.EN:   `Lifetime`,
		_const.ZhCN: `寿命`,
	},
	"ch_138": {
		_const.RU:   `Восстановление прочности`,
		_const.EN:   `Restoring Strength`,
		_const.ZhCN: `恢复强度`,
	},
	"ch_139": {
		_const.RU:   `Время восстановления после урона`,
		_const.EN:   `Recovery time after damage`,
		_const.ZhCN: `损坏后恢复时间`,
	},
	"ch_140": {
		_const.RU:   `Структура`,
		_const.EN:   `Structure`,
		_const.ZhCN: `结构`,
	},
	"ch_141": {
		_const.RU:   `Дальность полета капсулы`,
		_const.EN:   `Capsule flight range`,
		_const.ZhCN: `胶囊飞行范围`,
	},
	"ch_142": {
		_const.RU:   `Кол-во зарядов`,
		_const.EN:   `Number of charges`,
		_const.ZhCN: `充能次数`,
	},
	"ch_143": {
		_const.RU:   `Потребляемая энергия за использование`,
		_const.EN:   `Energy consumption per use`,
		_const.ZhCN: `每次使用的能量消耗`,
	},
	"ch_144": {
		_const.RU:   `Время перезарядки`,
		_const.EN:   `Recharge time`,
		_const.ZhCN: `重新充能时间`,
	},
	"ch_145": {
		_const.RU:   `Радиус`,
		_const.EN:   `Radius`,
		_const.ZhCN: `半径`,
	},
	"ch_146": {
		_const.RU:   `Потребляемая энергия за использование`,
		_const.EN:   `Energy consumption per use`,
		_const.ZhCN: `每次使用的能量消耗`,
	},
	"ch_147": {
		_const.RU:   `Потребляя мощность`,
		_const.EN:   `Consuming power`,
		_const.ZhCN: `消耗功率`,
	},
	"ch_148": {
		_const.RU:   `Занимаемый объем`,
		_const.EN:   `Occupied volume`,
		_const.ZhCN: `占用体积`,
	},
	"ch_149": {
		_const.RU:   `Размер оружия`,
		_const.EN:   `Weapon size`,
		_const.ZhCN: `武器尺寸`,
	},
	"ch_150": {
		_const.RU:   `Тип оружия`,
		_const.EN:   `Weapon type`,
		_const.ZhCN: `武器类型`,
	},
	"ch_151": {
		_const.RU:   `Скорость поворота`,
		_const.EN:   `Swing speed`,
		_const.ZhCN: `摆动速度`,
	},
	"ch_152": {
		_const.RU:   `Модификатор урона`,
		_const.EN:   `Damage modifier`,
		_const.ZhCN: `伤害修正`,
	},
	"ch_153": {
		_const.RU:   `Доступные углы атаки`,
		_const.EN:   `Available angles of attack`,
		_const.ZhCN: `可用攻击角度`,
	},
	"ch_154": {
		_const.RU:   `Влияние на скорость пули`,
		_const.EN:   `Effect on bullet speed`,
		_const.ZhCN: `对子弹速度的影响`,
	},
	"ch_155": {
		_const.RU:   `Дальность стрельбы`,
		_const.EN:   `Range attack`,
		_const.ZhCN: `攻击范围`,
	},
	"ch_156": {
		_const.RU:   `зависит от снаряда`,
		_const.EN:   `depends on the projectile`,
		_const.ZhCN: `取决于弹丸`,
	},
	"ch_157": {
		_const.RU:   `Очередь`,
		_const.EN:   `Queue`,
		_const.ZhCN: `队列`,
	},
	"ch_158": {
		_const.RU:   `Скорострельность`,
		_const.EN:   `Rate of fire`,
		_const.ZhCN: `射速`,
	},
	"ch_159": {
		_const.RU:   `Боезапас`,
		_const.EN:   `Ammunition`,
		_const.ZhCN: `弹药`,
	},
	"ch_160": {
		_const.RU:   `Время перезарядки`,
		_const.EN:   `Reload time`,
		_const.ZhCN: `重新装填时间`,
	},
	"ch_161": {
		_const.RU:   `Потребляя мощность`,
		_const.EN:   `Consuming power`,
		_const.ZhCN: `消耗功率`,
	},
	"ch_162": {
		_const.RU:   `Занимаемый объем`,
		_const.EN:   `Occupied volume`,
		_const.ZhCN: `占用体积`,
	},
	"ch_163": {
		_const.RU:   `артиллерия`,
		_const.EN:   `artillery`,
		_const.ZhCN: `炮兵`,
	},
	"ch_167": {
		_const.RU:   `Энергоемкость`,
		_const.EN:   `Energy intensity`,
		_const.ZhCN: `能量强度`,
	},
	"ch_168": {
		_const.RU:   `Кол-во предметов:`,
		_const.EN:   `Number of items:`,
		_const.ZhCN: `物品数量：`,
	},
	"ch_169": {
		_const.RU:   "Усиление урона при столкновении",
		_const.EN:   "Collision damage boost",
		_const.ZhCN: "碰撞伤害增强",
	},
	"ch_170": {
		_const.RU:   "Периодический урон при контакте",
		_const.EN:   "Periodic contact damage",
		_const.ZhCN: "接触周期性伤害",
	},
	// DetailItem
	"text_1": {
		_const.RU:   `примерная цена: <span class="importantly">%credits_count%</span> cr.`,
		_const.EN:   `approximate price: <span class="importantly">%credits_count%</span> cr.`,
		_const.ZhCN: `预估价格：<span class="importantly">%credits_count%</span> cr.`,
	},
	"text_2": {
		_const.RU:   `Описание`,
		_const.EN:   `Description`,
		_const.ZhCN: `描述`,
	},
	"text_3": {
		_const.RU:   `Особенности`,
		_const.EN:   `Peculiarities`,
		_const.ZhCN: `特性`,
	},
	"text_4": {
		_const.RU:   `Хар-ки`,
		_const.EN:   `Specs`,
		_const.ZhCN: `规格`,
	},
	"text_5": {
		_const.RU:   `Хар-ки оружия`,
		_const.EN:   `Weapon specs.`,
		_const.ZhCN: `武器规格`,
	},
	"text_6": {
		_const.RU:   `Хар-ки снаряда`,
		_const.EN:   `Projectile specs.`,
		_const.ZhCN: `弹丸规格`,
	},
	"text_7": {
		_const.RU:   `Варианты`,
		_const.EN:   `Options`,
		_const.ZhCN: `选项`,
	},
	"text_8": {
		_const.RU:   `Требования`,
		_const.EN:   `Requirements`,
		_const.ZhCN: `要求`,
	},
	"text_9": {
		_const.RU:   `Оснащение`,
		_const.EN:   `Equipment`,
		_const.ZhCN: `装备`,
	},
	"text_10": {
		_const.RU:   `Произ-во`,
		_const.EN:   `Production`,
		_const.ZhCN: `生产`,
	},
	"text_28": {
		_const.RU:   `Цены`,
		_const.EN:   `Prices`,
		_const.ZhCN: `价格`,
	},
	"broken_tip": {
		_const.RU: `Предмет сломан из за этого:<br>
          - его нельзя переработать или продать<br>
          - если выкинуть из грузового отсека он уничтожиться<br>
          - его характеристики могут быть ухудшены<br>`,
		_const.EN: `Item is broken because of this:<br>
		  - It cannot be recycled or sold<br>
		  - If ejected from cargo hold, it will be destroyed<br>
		  - Its characteristics may be degraded<br>`,
		_const.ZhCN: `物品因此损坏：<br>
		  - 不可回收或出售<br>
		  - 从货舱丢弃将被销毁<br>
		  - 性能参数可能降低<br>`,
	},
	"book_tip_1_1": {
		_const.RU: `Чипы можно получить:<br>
          - в фракционном магазине<br>
          - исследуя аномалии<br>
          - купив на рынке`,
		_const.EN: `Skill chips can be obtained by:<br>
          - purchasing from the faction store<br>
          - exploring anomalies<br>
          - buying from the market`,
		_const.ZhCN: `技能芯片获取方式：<br>
          - 在阵营商店购买<br>
          - 探索异常现象<br>
          - 在市场上交易`,
	},
	"resource_tip_1_1": {
		_const.RU:   `Этот ресурс можно получить`,
		_const.EN:   `This resource can be obtained by `,
		_const.ZhCN: `可以通过以下方式获取此资源：`,
	},
	"resource_tip_1_2": {
		_const.RU:   `переработав`,
		_const.EN:   `processing`,
		_const.ZhCN: `加工`,
	},
	"resource_tip_2_1": {
		_const.RU:   `Этот ресурс можно`,
		_const.EN:   `This resource can be`,
		_const.ZhCN: `此资源可以`,
	},
	"resource_tip_2_2": {
		_const.RU:   `добыть в пустошах`,
		_const.EN:   `mined in the wasteland`,
		_const.ZhCN: `在荒地中开采`,
	},
	"blueprints_tip_1_1": {
		_const.RU:   `Чертежи можно получить из обломков или исследуя`,
		_const.EN:   `Blueprints can be obtained from debris or by exploring`,
		_const.ZhCN: `可以从残骸中或通过探索获取蓝图`,
	},
	"blueprints_tip_1_2": {
		_const.RU:   `аномалии в пустошах`,
		_const.EN:   `anomalies in the wasteland`,
		_const.ZhCN: `荒地中的异常现象`,
	},
	"detail_tip_1": {
		_const.RU:   `Детали можно получить:`,
		_const.EN:   `Details can be obtained:`,
		_const.ZhCN: `可以通过以下方式获取零件：`,
	},
	"detail_tip_1_1": {
		_const.RU:   `создав их на заводе полуфабрикатов`,
		_const.EN:   `creating them at a semi-finished products factory`,
		_const.ZhCN: `在半成品工厂中制造`,
	},
	"detail_tip_1_2": {
		_const.RU:   `создав их по чертежам`,
		_const.EN:   `creating them according to drawings`,
		_const.ZhCN: `根据图纸制造`,
	},
	"detail_tip_1_3": {
		_const.RU:   `переработав обломки`,
		_const.EN:   `recycling the wreckage`,
		_const.ZhCN: `回收残骸`,
	},
	// EquipmentTab
	"text_11": {
		_const.RU:   `Ячейки для топлива`,
		_const.EN:   `Fuel cells`,
		_const.ZhCN: `燃料槽`,
	},
	"text_12": {
		_const.RU:   `Лимит мощности`,
		_const.EN:   `Power limit`,
		_const.ZhCN: `功率限制`,
	},
	"text_13": {
		_const.RU:   `Слотов оружия`,
		_const.EN:   `Weapon slots`,
		_const.ZhCN: `武器槽`,
	},
	"text_14": {
		_const.RU:   `Разъемы большой мощности`,
		_const.EN:   `High power connectors`,
		_const.ZhCN: `高功率连接器`,
	},
	"text_15": {
		_const.RU:   `Разъемы средней мощности`,
		_const.EN:   `Medium power connectors`,
		_const.ZhCN: `中功率连接器`,
	},
	"text_16": {
		_const.RU:   `Разъемы малой мощности`,
		_const.EN:   `Low power connectors`,
		_const.ZhCN: `低功率连接器`,
	},
	"text_17": {
		_const.RU:   `Боеприпасы`,
		_const.EN:   `Ammunition`,
		_const.ZhCN: `弹药`,
	},
	"text_27": {
		_const.RU:   `Разъемы оружия ближнего боя:`,
		_const.EN:   `Melee weapon connectors:`,
		_const.ZhCN: `近战武器插槽：`,
	},
	// ProductionTab
	"text_18": {
		_const.RU:   `Продукция`,
		_const.EN:   `Products`,
		_const.ZhCN: `产品`,
	},
	"text_19": {
		_const.RU:   `Требуемые ресурсы <span style="font-size: 9px">(при 100% эф.)</span>`,
		_const.EN:   `Required resources <span style="font-size: 9px">(at 100% ef.)</span>`,
		_const.ZhCN: `所需资源 <span style="font-size: 9px">(效率 100%)</span>`,
	},
	"text_20": {
		_const.RU:   `Можно получить переработав`,
		_const.EN:   `Can be obtained by recycling`,
		_const.ZhCN: `可以通过回收获得`,
	},
	"text_21": {
		_const.RU:   `Чертежи`,
		_const.EN:   `Blueprints`,
		_const.ZhCN: `蓝图`,
	},
	"text_22": {
		_const.RU:   `Компоненты после переработки`,
		_const.EN:   `Components after processing`,
		_const.ZhCN: `加工后的组件`,
	},
	"text_24": {
		_const.RU:   `Переработав <span style="color: yellow">%count%</span> ед. выйдет:`,
		_const.EN:   `Having processed <span style="color: yellow">%count%</span> units. will come out:`,
		_const.ZhCN: `处理 <span style="color: yellow">%count%</span> 单位后，将得到：`,
	},
	// Requirements
	"text_23": {
		_const.RU:   `Навыки`,
		_const.EN:   `Skills`,
		_const.ZhCN: `技能`,
	},
}
