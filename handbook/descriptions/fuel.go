package descriptions

import _const "github.com/TrashPony/veliri-lib/const"

var FuelDescription = map[string]map[string]DescriptionItem{
	_const.EN: {
		"gas_fuel":          {Name: "Gas Cell Battery", Description: "<p>A rechargeable battery based on a gas fuel cell. Cheap, simple to operate, has low capacity.</p>"},
		"solid_fuel":        {Name: "Solid-State Cell Battery", Description: "<p>A rechargeable solid-state type battery. Cheap, simple to operate, has slightly increased capacity.</p><p>The cell's stable operation evenly improves all vehicle systems.</p>"},
		"liquid_fuel":       {Name: "Liquid Cell Battery", Description: "<p>A rechargeable liquid type battery. Common, has average capacity.</p><p>Electrolyte purity reduces interference, improving the operation of optical and sensor systems.</p>"},
		"nuclear_fuel":      {Name: "Nuclear Cell Battery", Description: "<p>A rechargeable battery with a nuclear power source. Has enormous capacity and high charging speed.</p>"},
		"synthetic_fuel":    {Name: "Synthetic Cell Battery", Description: "<p>A rechargeable battery based on synthetic materials. Has average capacity.</p><p>Highly purified composition practically eliminates interference, significantly improving the efficiency of detection systems.</p>"},
		"jet_fuel":          {Name: "Pulse Cell Battery", Description: "<p>A rechargeable battery optimized for pulse loads. Has low capacity, but significantly increases mobility.</p>"},
		"emergency_fuel_t0": {Name: "Small Energy Cartridge", Description: "<p>A disposable small chemical battery. Cheap, compact, has extremely low capacity.</p><p>Used for rapid recharging of main batteries. Slightly increases battery wear during recharging.</p><p>Can be used directly by the chassis, but this causes strong interference to equipment operation.</p>"},
		"emergency_fuel_t1": {Name: "Standard Energy Cartridge", Description: "<p>A disposable chemical battery. Compact, has low capacity.</p><p>Used for rapid recharging of main batteries. Slightly increases battery wear during recharging.</p><p>Can be used directly by the chassis, but this causes interference to equipment operation.</p>"},
		"emergency_fuel_t2": {Name: "Large Energy Cartridge", Description: "<p>A disposable chemical battery. Compact design with high capacity.</p><p>Used for rapid recharging of main batteries. Slightly increases battery wear during recharging.</p><p>Can be used directly by the frame, but this causes equipment interference.</p>"},
	},
	_const.RU: {
		"gas_fuel":          {Name: "Ячейка газового аккумулятора", Description: "<p>Перезаряжаемый аккумулятор на основе газового топливного элемента. Дёшев, прост в эксплуатации, обладает малой ёмкостью.</p>"},
		"solid_fuel":        {Name: "Ячейка твердотельного аккумулятора", Description: "<p>Перезаряжаемый аккумулятор твёрдотельного типа. Дёшев, прост в эксплуатации, обладает слегка повышенной ёмкостью.</p><p>Стабильная работа элемента равномерно улучшает все системы транспорта.</p>"},
		"liquid_fuel":       {Name: "Ячейка жидкостного аккумулятора", Description: "<p>Перезаряжаемый аккумулятор жидкостного типа. Распространён, имеет среднюю ёмкость.</p><p>Чистота электролита снижает помехи, улучшая работу оптических и сенсорных систем.</p>"},
		"nuclear_fuel":      {Name: "Ячейка ядерного аккумулятора", Description: "<p>Перезаряжаемый аккумулятор с ядерным источником питания. Обладает огромной ёмкостью и высокой скоростью зарядки.</p>"},
		"synthetic_fuel":    {Name: "Ячейка синтетического аккумулятора", Description: "<p>Перезаряжаемый аккумулятор на основе синтетических материалов. Имеет среднюю ёмкость.</p><p>Высокоочищенный состав практически исключает помехи, значительно улучшая эффективность систем обнаружения.</p>"},
		"jet_fuel":          {Name: "Ячейка импульсного аккумулятора", Description: "<p>Перезаряжаемый аккумулятор, оптимизированный для импульсных нагрузок. Имеет малую ёмкость, но значительно повышает подвижность.</p>"},
		"emergency_fuel_t0": {Name: "Малый энергопатрон", Description: "<p>Одноразовая химическая малая батарея. Дешёва, компактна, обладает крайне малой ёмкостью.</p><p>Используется для быстрой подзарядки основных аккумуляторов. При подзарядке немного увеличивает износ аккумулятора.</p> Может использоваться корпусом напрямую, но это вызывает сильные помехи в работе оборудования.</p>"},
		"emergency_fuel_t1": {Name: "Стандартный энергопатрон", Description: "<p>Одноразовая химическая батарея. Компактна, обладает малой ёмкостью.</p><p>Используется для быстрой подзарядки основных аккумуляторов. При подзарядке немного увеличивает износ аккумулятора.</p> Может использоваться корпусом напрямую, но это вызывает помехи в работе оборудования.</p>"},
		"emergency_fuel_t2": {Name: "Большой энергопатрон", Description: "<p>Одноразовая химическая батарея. Компактна, обладает большой ёмкостью.</p><p>Используется для быстрой подзарядки основных аккумуляторов. При подзарядке немного увеличивает износ аккумулятора.</p> Может использоваться корпусом напрямую, но это вызывает помехи в работе оборудования.</p>"},
	},
	_const.ZhCN: {
		"gas_fuel":          {Name: "气体电池单元", Description: "<p>基于气体燃料电池的可充电电池。成本低、操作简单、容量小。</p>"},
		"solid_fuel":        {Name: "固态电池单元", Description: "<p>固态可充电电池。成本低、操作简单、容量略有提高。</p><p>电池稳定运行能均衡提升载具所有系统性能。</p>"},
		"liquid_fuel":       {Name: "液体电池单元", Description: "<p>液体型可充电电池。常见，具有中等容量。</p><p>电解液纯度减少干扰，改善光学与传感器系统工作。</p>"},
		"nuclear_fuel":      {Name: "核能电池单元", Description: "<p>采用核动力源的可充电电池。容量极大且充电速度高。</p>"},
		"synthetic_fuel":    {Name: "合成材料电池单元", Description: "<p>基于合成材料的可充电电池。具有中等容量。</p><p>高纯度成分几乎杜绝干扰，显著提升探测系统效率。</p>"},
		"jet_fuel":          {Name: "脉冲电池单元", Description: "<p>针对脉冲负载优化的可充电电池。容量小，但显著提高机动性。</p>"},
		"emergency_fuel_t0": {Name: "小型能量弹", Description: "<p>一次性小型化学电池。廉价、紧凑，容量极低。</p><p>用于快速充能主电池。充能时会略微增加电池损耗。</p><p>可直接被机体使用，但会对设备运行造成严重干扰。</p>"},
		"emergency_fuel_t1": {Name: "标准能量弹", Description: "<p>一次性化学电池。紧凑，容量低。</p><p>用于快速充能主电池。充能时会略微增加电池损耗。</p><p>可直接被机体使用，但会对设备运行造成干扰。</p>"},
		"emergency_fuel_t2": {Name: "大型能量单元", Description: "<p>一次性化学电池。设计紧凑，容量大。</p><p>用于快速给主电池充电。充电时会略微增加电池损耗。</p><p>可被机体框架直接使用，但这会导致设备干扰。</p>"},
	},
}
