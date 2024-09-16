package descriptions

import _const "github.com/TrashPony/veliri-lib/const"

var EquipDescription = map[string]map[string]DescriptionItem{
	_const.EN: {
		"repair_kit":                     {Name: "Self-Repair Module", Description: "<p>About a million nanobots armed with microelements of spare parts and, when activated, begin immediate restoration work on the damaged hull and systems.</p>"},
		"energy_shield":                  {Name: "Force Field M-Shields", Description: "<p>Low-spectrum force emitters provide energy protection against any phenomenon trying to overcome the barrier. Significantly reduce damage, but do not guarantee one hundred percent protection.</p>"},
		"armored":                        {Name: "Armored Reinforcement Plate", Description: "<p>A standard hull armor booster that increases the overall percentage of damage resistance.</p>"},
		"digger":                         {Name: "«Phobos» Drone", Description: "<p>Remotely controlled drone used for surveying planetary anomalies and capable of attacking a variety of targets.</p>"},
		"geo_scanner":                    {Name: "Georadar", Description: "<p>Planetary analysis device that allows you to explore the surface for irregularities, voids and anomalous manifestations.</p>"},
		"jammer":                         {Name: "EMF Communication Jammer", Description: "<p>Advanced electromagnetic warfare device accumulating kinetic energy discharge, which temporarily jams all other electrical devices.</p>"},
		"missile_defense":                {Name: "ABM Complex", Description: "<p>A module that provides a no-fly zone for any types of missiles within a small radius due to overloading and failures of their internal guidance systems.</p>"},
		"repair_passive":                 {Name: "Nanite Restoration", Description: "<p>Passive support module that does not consume energy and, in case of damage to the hull or any systems, organizes immediate repair work by sending nanite robots.</p>"},
		"distance_repair":                {Name: "Remote Repair System", Description: "<p>The logical development of the «self-repair module» system with the ability to send remote repair nanobots to assist in restoring any allied targets.</p>"},
		"distance_generator":             {Name: "Energy Transfer Module", Description: "<p>A module specialized in the highly focused transfer of excess or reserved energy from its own hull to any target at a short distance.</p>"},
		"replic_miner_extracted":         {Name: "«Replic» Mining Projector", Description: "<p>Standardized device for the development and extraction of rocks on the planet Veliri-6. The device is adapted to the aggressive properties of the environment and is resistant to radiation.</p>"},
		"replic_builder":                 {Name: "«Replic» Engineering Module", Description: "<p>Standardized device for assembling and disassembling objects, as well as extracting useful parts from debris in great detail.</p>"},
		"explores_miner_extracted":       {Name: "«Explores» Mining Beam", Description: "<p>Standardized device for the development and extraction of rocks on the planet Veliri-6. The device is adapted to the aggressive properties of the environment and is resistant to radiation.</p>"},
		"explores_builder":               {Name: "«Explores» Engineering Module", Description: "<p>Standardized device for assembling and disassembling objects, as well as extracting useful parts from debris in great detail.</p>"},
		"reverses_miner_extracted":       {Name: "«Reverses» Ore Pulsator", Description: "<p>Standardized device for the development and extraction of rocks on the planet Veliri-6. The device is adapted to the aggressive properties of the environment and is resistant to radiation.</p>"},
		"reverses_builder":               {Name: "«Reverses» Engineering Module", Description: "<p>Standardized device for assembling and disassembling objects, as well as extracting useful parts from debris in great detail.</p>"},
		"gravity_gun":                    {Name: "MEP-0 Device", Description: "<p>A manipulator of gravitational influence used to attract absolutely any captured objects to the position of the device with a directed beam and, with a powerful impulse, send them in a different direction.</p>"},
		"gravity_square":                 {Name: "Gravitational Disturbance Resonator", Description: "<p>An accumulative device that creates waves of a gravitational nature and draws everything into the center of its own attraction, including combat shells.</p>"},
		"mini_turret_1":                  {Name: "Defense System «Sentinel»", Description: "<p>Deployable and stationary self-sustaining combat unit that acts as a «defensive tower» and fires at the enemy with a small-caliber machine gun.</p>"},
		"mini_turret_2":                  {Name: "Defense System «Destroyer»", Description: "<p>Deployable and stationary self-sustaining combat unit acting as a «defensive tower» and firing at the enemy using a rocket launcher with automated projectile supply.</p>"},
		"smoke_screen_local":             {Name: "Smoke Screen", Description: "<p>A lifesaver for careless spies and top management Synthetics of «factions». When activated, this device releases a thick cloud of smoke that follows the movement of the Synthetic. The smoke screen hides the Synthetic's movement and allows them to escape from visual pursuit by the enemy.</p>"},
		"smoke_screen_destination":       {Name: "Revolving Smoke Discharger", Description: "<p>A device used by sabotage units and merchants trying to protect their goods during raids. It makes three consecutive shots with specialized projectiles that, upon contact with any surface, explode and create an impenetrable smoke screen.</p>"},
		"drone_scout_1":                  {Name: "Z-C:H Observer", Description: "<p>Made of lightweight alloys and having an air-flowing shape, this screw-driven remote-controlled reconnaissance drone is designed to observe the terrain and detect potential enemy signatures.</p>"},
		"drone_defender_1":               {Name: "Autonomous Defender «Gunner»", Description: "<p>An independently moving bot that always follows its owner. Its main task is fire support using two fused machine gun turrets.</p></p>"},
		"drone_defender_2":               {Name: "Autonomous Defender «Grenadier»", Description: "<p>An independently moving bot that always follows its owner. Its main task is fire support using a rocket launcher with automated projectile supply.</p>"},
		"rope_trap_1":                    {Name: "Position Capture", Description: "<p>Remotely deployed trap made of metal-fiber, whose sole purpose is to shoot a welded harpoon at the nearest target and hold it for as long as possible.</p><p>Effective against light vehicles, but medium and heavy vehicles are likely to break the cable quickly.</p>"},
		"grenade_1":                      {Name: "Explosive Grenade", Description: "<p>A projectile with a fragmentation effect and parabolic flight path. It has moderate damage that spreads to all targets in the vicinity of the explosion epicenter. Additionally, it has a «repulsive effect» on everything that is close to the center of the projectile’s detonation.</p>"},
		"gravity_grenade_1":              {Name: "S-Gravity Converter", Description: "<p>Firing a specialized capsule charge containing a cluster of lightweight antimatter. At the point of impact and destruction of the antimatter capsule, magnetic vortex attraction forces immediately form. The affected area becomes a kind of «well» where artificially created gravity pulls any objects.</p>"},
		"jump_drive_1":                   {Name: "Amido-Pulse Booster Module", Description: "<p>Experimental device for «rapid» distance reduction by injecting a special fuel element into the propulsion system. Recreating the short-term effect of a «jump», the Synthetic is almost instantly able to move a short distance.</p>"},
		"invisibility_1":                 {Name: "Stealth System «Chameleon»", Description: "<p>A «temporary» invisibility tool designed for reconnaissance units and self-taught spies. Allows the wearer to become «transparent» by deflecting light rays and not showing their own thermal profile on radar systems.</p>"},
		"wall_1":                         {Name: "Mobile Barricade", Description: "<p>Remotely erected engineering structure that acts as a protective wall, allowing you to hide from enemy fire or even block their escape route.</p>"},
		"energy_shield_mini_structure_1": {Name: "Mobile Force Field Mark-2", Description: "<p>Remotely erected engineering structure acting as a power plant that projects an energy dome. Inside the impenetrable dome, you can hide from most types of small arms.</p>"},
		"mine_bomb_1":                    {Name: "Mine Layer", Description: "<p>A product of military engineering, a self-defense tool that drops a cumulative mine behind the Synthetic's hull, which reacts to approaching enemy signatures and self-detonates.</p>"},

		// пассивное
		"radar_booster_1":                {Name: "Sensor Repeater", Description: "<p>A set of amplifiers that provide increased range for an already installed radar system.</p><p>The radar allows you to detect and identify objects in the world, such as resources, structures, vehicles, drones/missiles.</p>"},
		"view_booster_1":                 {Name: "Collecting Lens", Description: "<p>Prototype observation equipment that provides guaranteed increase in range and quality of view.</p>"},
		"improved_inventory_capacity_1":  {Name: "W.H.T.C.", Description: "<p>'Warehouse Hull Trunk Controller' is a hull upgrade whose function is to expand the hold volume due to an advanced modular structuring system.</p>"},
		"antigravity_speed_booster_1":    {Name: "Injector", Description: "<p>A particle accelerator of a vacuum chamber that allows anti-gravity hulls to develop a higher speed due to an alternating electro-field.</p>"},
		"caterpillars_speed_booster_1":   {Name: "T-suspension", Description: "<p>Development of lightweight and, as a result, simpler tracked suspension technology, whose only advantage is an increase in the speed of tracked vehicles.</p>"},
		"antigravity_mobility_booster_1": {Name: "Maneuvering unit", Description: "<p>Module applicable only to anti-gravitational hulls, whose activation forms the formation of an ionized arc and increases maneuverability.</p>"},
		"wheels_mobility_booster_1":      {Name: "Axis kinematics", Description: "<p>Standardized improvement of wheeled vehicles. After installation, it improves the characteristics of the hull's maneuverability/sharpness of rotation.</p>"},
		"ballistic_rotate_booster_1":     {Name: "Rotator: Ballistics", Description: "<p>A tower rotation device that allows you to reduce the time it takes to turn ballistic weapons.</p>"},
		"ballistic_damage_booster_1":     {Name: "Acceleration Module", Description: "<p>An apparatus of a closed magnetic field that accelerates ballistic ammunition when fired, which increases the damage they inflict.</p>"},
		"missile_rotate_booster_1":       {Name: "Rotator: Missiles", Description: "<p>A device for tower rotation that reduces the time it takes for missile weapons to rotate.</p>"},
		"missile_damage_booster_1":       {Name: "Burning Tip", Description: "<p>Missile weapon enhancement module by replacing standard methods of destruction with narrowly focused cumulative impact, which increases damage.</p>"},
		"laser_rotate_booster_1":         {Name: "Rotator: Laser", Description: "<p>A tower rotation device that allows you to reduce the time it takes to turn laser weapons.</p>"},
		"laser_damage_booster_1":         {Name: "Thermal Storage Unit", Description: "<p>Emitter of a heating element that, when fired, releases a more concentrated clot of heat, which increases the damage of laser weapons.</p>"},
		"energy_capacity_1":              {Name: "Accumulator", Description: "<p>Technical equipment of an energy variety. Provides the hull with additional energy volume that can be used for various purposes.</p>"},
		"energy_charging_speed_1":        {Name: "Auxiliary Reactor", Description: "<p>Allows the hull to charge the accumulator faster and more efficiently from the main thorium reactor.</p>"},
		"power_booster_1":                {Name: "Converter Unit", Description: "<p>A support device for the energy variety. Reduces the pressure exerted on the thorium reactor and redistributes energy, which generally increases its efficiency. Allows you to install more powerful equipment or weapons on the hull.</p>"},
		"body_structure_booster_1":       {Name: "Composite Alloys", Description: "<p>Sprayed armor coating alloys designed to strengthen the hull. Provide greater transport strength.</p>"},

		"rig_armored":                       {Name: "Модификатор: броня общая", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_armored_explosion":             {Name: "Модификатор: броня от взрыва", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_armored_kinetics":              {Name: "Модификатор: броня от кинетики", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_armored_thermo":                {Name: "Модификатор: броня от термо", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_body_structure_booster":        {Name: "Модификатор: конструкция", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_caterpillars_speed_booster_1":  {Name: "Модификатор: скорость гусениц", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_antigravity_speed_booster_1":   {Name: "Модификатор: скорость антиграва", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_wheels_speed_booster_1":        {Name: "Модификатор: скорость колес", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_radar_booster_1":               {Name: "Модификатор: радар", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_view_booster_1":                {Name: "Модификатор: линзы", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_energy_capacity_1":             {Name: "Модификатор: емкость аккумулятора", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_energy_charging_speed_1":       {Name: "Модификатор: зарядка аккумулятора", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_improved_inventory_capacity_1": {Name: "Модификатор: объем трюма", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_power_booster_1":               {Name: "Модификатор: реактор", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_missile_damage_booster_1":      {Name: "Модификатор: урон ракет", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_laser_damage_booster_1":        {Name: "Модификатор: урон лазера", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_ballistic_damage_booster_1":    {Name: "Модификатор: урон снарядов", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_ballistic_rotate_booster_1":    {Name: "Модификатор: турель баллистики", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_missile_rotate_booster_1":      {Name: "Модификатор: турель пусковой установки", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_laser_rotate_booster_1":        {Name: "Модификатор: турель лазера", Description: "<p>Модуль уничтожается если его снять!</p>"},

		"dust_collector":      {Name: "Сборщик пыли", Description: "<p></p>"},
		"pump_extracted":      {Name: "Насос", Description: "Позволяет добывать жидкие ресурсы, включая нефть."},
		"inventory_scanner_1": {Name: "Сканер", Description: "Позволяет сканировать трюмы других транспортов."},
		"reverses_combine":    {Name: "Комбайн", Description: "<p>Стандартизированное устройство по сборке органики.</p>"},
	},
	_const.RU: {
		"repair_kit":                     {Name: "Модуль само-ремонта", Description: "<p>Около миллиона наноботов, вооружённых микроэлементами запчастей и при активации, приступающих к незамедлительным восстановительным работам повреждённого корпуса и систем.</p>"},
		"energy_shield":                  {Name: "М-щиты силового поля", Description: "<p>Силовые излучатели малого спектра энергетической защиты, противодействующих любому явлению, что пытается преодолеть барьер.  Существенно снижают урон, но не гарантируют стопроцентной защиты.</p>"},
		"armored":                        {Name: "Армированная броненакладка", Description: "<p>Усилитель стандартной брони корпуса, повышающий общий процент сопротивляемости к урону.</p>"},
		"digger":                         {Name: "Дрон \"Фобос\"", Description: "<p>Дрон на удалённом управлении, применяемый для работ изыскательного характера с планетарными аномалиями и возможностью атаки самых разных целей.</p>"},
		"geo_scanner":                    {Name: "Георадар", Description: "<p>Устройство планетарного анализа, позволяющее исследовать поверхность на наличие неоднородностей, пустот и проявлений аномального характера.</p>"},
		"jammer":                         {Name: "ЭМИ подавитель связи", Description: "<p>Передовое устройство электромагнитной войны, аккумулирующее кинетический разряд энергии, что на время глушит все прочие электротехнические устройства.</p>"},
		"missile_defense":                {Name: "Комплекс-ПРО", Description: "<p>Модуль, обеспечивающий в малом радиусе бесполётную зону для любых видов снарядов ракетно-залповых систем, за счёт перегрузки и сбоев их внутренних систем наведения.</p>"},
		"repair_passive":                 {Name: "Нанитовое восстановление", Description: "<p>Модуль пассивной поддержки, не расходующий энергии и в случае выявления повреждений корпуса или, каких-либо систем – организующий начало незамедлительных ремонтных работ, за счёт отправки роботов-нанитов.</p>"},
		"distance_repair":                {Name: "Д\\У система ремонта", Description: "<p>Логическое развитие системы «модуля само-ремонта» с возможностью отправки дистанциируемых ремонтных наноботов для оказания поддержки восстановления любых союзных целей.</p>"},
		"distance_generator":             {Name: "Модуль передачи энергии", Description: "<p>Модуль, специализирующийся на узконаправленной передаче излишка или зарезервированной доли энергии из собственного корпуса к любой цели на малой дистанции.</p>"},
		"replic_miner_extracted":         {Name: "Горнопроходческий проектор \"Replic\"", Description: "<p>Стандартизированное устройство по разработке и извлечению горных пород на планете Veliri-6. Устройство адаптировано под агрессивные свойства окружающей среды и устойчиво к радиации.</p>"},
		"replic_builder":                 {Name: "Инженерный модуль \"Replic\"", Description: "<p>Стандартизированное устройство по сборке и разборке объектов, а также детальнейшей добыче полезных частей из обломков.</p>"},
		"explores_miner_extracted":       {Name: "Добывающий луч \"Explores\"", Description: "<p>Стандартизированное устройство по разработке и извлечению горных пород на планете Veliri-6. Устройство адаптировано под агрессивные свойства окружающей среды и устойчиво к радиации.</p>"},
		"explores_builder":               {Name: "Инженерный модуль \"Explores\"", Description: "<p>Стандартизированное устройство по сборке и разборке объектов, а также детальнейшей добыче полезных частей из обломков.</p>"},
		"reverses_miner_extracted":       {Name: "Рудный пульсатор \"Reverses\"", Description: "<p>Стандартизированное устройство по разработке и извлечению горных пород на планете Veliri-6. Устройство адаптировано под агрессивные свойства окружающей среды и устойчиво к радиации.</p>"},
		"reverses_builder":               {Name: "Инженерный модуль \"Reverses\"", Description: "<p>Стандартизированное устройство по сборке и разборке объектов, а также детальнейшей добыче полезных частей из обломков.</p>"},
		"gravity_gun":                    {Name: "Устройство МЭП-0", Description: "<p>Манипулятор гравитационного воздействия, применяемый для притягивания направленным лучом абсолютно любых захваченных объектов к позиции работы устройства и, мощным импульсом – их отправки в иное направление.</p>"},
		"gravity_square":                 {Name: "Резонатор гравитационного возмущения", Description: "<p>Устройство накопительного действия, создающего волны гравитационной природы и затягивающего в центр собственного притяжения всякую вещь, включая боевые снаряды.</p>"},
		"mini_turret_1":                  {Name: "Система обороны “Сторож”", Description: "<p>Развёртываемая и неподвижная самообеспечивающаяся боевая единица, исполняющая функции “защитной башни” и обстреливающая противника посредством мелкокалиберного пулемёта.</p>"},
		"mini_turret_2":                  {Name: "Система обороны “Разрушитель”", Description: "<p>Развёртываемая и неподвижная самообеспечивающаяся боевая единица, выполняющая функции “защитной башни” и обстреливающая противника посредством ракетной установки с автоматизированной подачей снарядов.</p>"},
		"smoke_screen_local":             {Name: "Дымовое прикрытие", Description: "<p>Средство спасения нерадивых шпионов и высших управленческих чинов Синтетов «фракций». При активации, данное устройство выпускает густое и тянущееся прямиком вслед за движением Синтета облако завесы. Завеса скрывает перемещение Синтета и позволяет тому скрыться от визуального преследования неприятеля.</p>"},
		"smoke_screen_destination":       {Name: "Револьверная установка дымового выстрела", Description: "<p>Устройство диверсионных подразделений и торговцев старающихся сберечь свой товар при налётах. Делает три последовательных выстрела специализированными снарядами, что после соприкосновения с любой поверхностью разрываются и создают непроглядную дымовую завесу.</p>"},
		"drone_scout_1":                  {Name: "Наблюдатель Z-C:H", Description: "<p>Созданный из облегчённых сплавов и имеющий воздухообтекаемую форму, этот винтовой, дистанционно управляемый дрон-разведчик предназначен для наблюдения за местностью и выявлению потенциальных вражеских сигнатур.</p>"},
		"drone_defender_1":               {Name: "Автономный защитник “Стрелок”", Description: "<p>Самостоятельно передвигающийся, но всегда следующий за своим владельцем - бот, чья главная задача - огневая поддержка посредством двух спаянных пулемётных башен.</p>"},
		"drone_defender_2":               {Name: "Автономный защитник “Гренадёр”", Description: "<p>Самостоятельно передвигающийся, но всегда следующий за своим владельцем - бот, чья главная задача - огневая поддержка посредством ракетной установки с автоматизированной подачей снарядов.</p>"},
		"rope_trap_1":                    {Name: "Позиционный захват", Description: "<p>Дистанционно устанавливаемая ловушка из металла-волокна, чья единственная задача - выстреливать сварным гарпуном в ближайшую цель и удерживать её как можно дольше.</p> <p>Эффективно против легкой техники но средняя и тяжелая скорее всего быстро порвут трос.</p>"},
		"grenade_1":                      {Name: "Разрывная граната", Description: "<p>Снаряд с эффектом осколочного действия и параболической траекторией полёта. Обладает умеренным уроном, который распространяется на все близлежащие поблизости с эпицентром взрыва цели. Дополнительно, оказывает “отталкивающее действие” на всё, что окажется неподалёку от центра разрыва снаряда.</p>"},
		"gravity_grenade_1":              {Name: "Преобразователь С-Тяжести", Description: "<p>Выстрел специализированного заряда-капсулы, содержащего сгусток облегчённой антиматерии. На месте попадания и разрушения капсулы с антиматерией, немедленно образуются магнитно-вихревые силы притяжения. Область поражения становится своеобразным “колодцем” куда искусственно образованная гравитация стягивает любые объекты.</p>"},
		"jump_drive_1":                   {Name: "Амидо-Пульсовой разгонный модуль", Description: "<p>Экспериментальное устройство “стремительного” сокращения расстояния посредство впрыска особого топливного элемента в двигательную систему. Воссоздавая краткосрочный эффект “скачка” - Синтет чуть-ли не мгновенно способен переместиться на небольшое расстояние.</p>"},
		"invisibility_1":                 {Name: "Стелс система “Хамелеон”", Description: "<p>Средство “временной” невидимости предназначенное для разведывательных подразделений и шпионов-самоучек. Позволяет владельцу становится “прозрачным” отклоняя световые лучи и не выказывая собственного теплового профиля на радарных системах.</p>"},
		"wall_1":                         {Name: "Мобильная баррикада", Description: "<p>Дистанционно возводимое инженерное сооружение выполняющее функцию защитной стены, что позволяет укрыться от огня неприятеля или, даже заблокировать тому путь к отступлению.</p>"},
		"energy_shield_mini_structure_1": {Name: "Мобильное силовое поле Марк-2", Description: "<p>Дистанционно возводимое инженерное сооружение выполняющее функцию силовой установки, проектирующей энергетический купол. Внутри непроницаемого купола можно укрыться от большинства видов стрелкового вооружения. </p>"},
		"mine_bomb_1":                    {Name: "Миноукладчик", Description: "<p>Разработка военной инженерии - средство самообороны, которое откидывает позади корпуса Синтета кумулятивную мину, что реагирует на приближение вражеских сигнатур и производящая самоподрыв.</p>"},

		// пассивное
		"radar_booster_1":                {Name: "Репитер сенсоров", Description: "<p>Станция набора усилителей, обеспечивающих увеличенную дальность действия уже установленной радарной системы.</p><p>Радар позволяет обнаруживать и идентифицировать объекты в мире, такие как ресурсы, структуры, транспорты и дроны/ракеты.</p>"},
		"view_booster_1":                 {Name: "Собирающая линза", Description: "<p>Прототип наблюдательного оборудования, обеспечивающее гарантированное увеличение дальности и качества обзора.</p>"},
		"improved_inventory_capacity_1":  {Name: "К.Г.Р.Т", Description: "<p>«Контроллер Грузового Расширителя Трюма» - апгрейд корпуса, чья функция расширить объём трюма за счёт продвинутой модульной системы структуризации.</p>"},
		"antigravity_speed_booster_1":    {Name: "Инжектор", Description: "<p>Ускоритель частиц вакуумной камеры, позволяющий корпусам на антиграве развивать более высокую скорость за счёт переменного электро-поля.</p>"},
		"caterpillars_speed_booster_1":   {Name: "Т-подвеска", Description: "<p>Разработка облегчённой и как следствие более простой в технологиях гусеничной подвески, чьё единственное достоинство – увеличение скорости гусеничного транспорта.</p>"},
		"antigravity_mobility_booster_1": {Name: "Маневровая установка", Description: "<p>Модуль, применимый только к корпусам на антиграве, чья активация формирует образование ионизированной дуги и увеличения манёвренности.</p>"},
		"wheels_mobility_booster_1":      {Name: "Кинематика оси", Description: "<p>Стандартизированное улучшение колёсного транспорта. После установки, повышает характеристики манёвренности/резкости поворота корпуса.</p>"},
		"ballistic_rotate_booster_1":     {Name: "Ротатор: баллистика", Description: "<p>Устройство башенного вращения, позволяющее уменьшить сроки поворота баллистического оружия.</p>"},
		"ballistic_damage_booster_1":     {Name: "Разгонный модуль", Description: "<p>Аппарат замкнутого магнитного поля, при выстреле разгоняющий баллистические виды боеприпасов, чем повышает наносимый ими урон.</p>"},
		"missile_rotate_booster_1":       {Name: "Ротатор: ракеты", Description: "<p>Устройство башенного вращения, позволяющее уменьшить сроки поворота ракетного оружия.</p>"},
		"missile_damage_booster_1":       {Name: "Прожигающий наконечник", Description: "<p>Модуль усиления ракетного вооружения за счёт замены стандартных способов поражения на узконаправленное кумулятивное воздействие, чем повышает наносимый урон.</p>"},
		"laser_rotate_booster_1":         {Name: "Ротатор: лазер", Description: "<p>Устройство башенного вращения, позволяющее уменьшить сроки поворота лазерного оружия.</p>"},
		"laser_damage_booster_1":         {Name: "Термический накопитель", Description: "<p>Испускатель нагревательного элемента, при выстреле выделяющий более концентрированный сгусток тепла, чем повышает урон лазерного вооружения.</p>"},
		"energy_capacity_1":              {Name: "Аккумулятор", Description: "<p>Техническое оборудование энергетической разновидности. Обеспечивает корпус дополнительным объемом энергии, что можно использовать для различных целей.</p>"},
		"energy_charging_speed_1":        {Name: "Вспомогательный реактор", Description: "<p>Позволяет корпусу быстрее и эффективнее заряжать аккумулятор от главного ториевого реактора.</p>"},
		"power_booster_1":                {Name: "Установка преобразователя", Description: "<p>Устройство поддержки энергетической разновидности. Снижает оказываемое на ториевый реактор давление и перераспределяет энергию, чем в целом повышает его эффективность. Позволяет устанавливать на корпус более мощное снаряжение или оружие.</p>"},
		"body_structure_booster_1":       {Name: "Композитные сплавы", Description: "<p>Распылённые сплавы бронепокрытия, предназначенные для усиления корпуса. Обеспечивают большую прочность транспорта.</p>"},

		"rig_armored":                       {Name: "Модификатор: броня общая", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_armored_explosion":             {Name: "Модификатор: броня от взрыва", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_armored_kinetics":              {Name: "Модификатор: броня от кинетики", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_armored_thermo":                {Name: "Модификатор: броня от термо", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_body_structure_booster":        {Name: "Модификатор: конструкция", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_caterpillars_speed_booster_1":  {Name: "Модификатор: скорость гусениц", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_antigravity_speed_booster_1":   {Name: "Модификатор: скорость антиграва", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_wheels_speed_booster_1":        {Name: "Модификатор: скорость колес", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_radar_booster_1":               {Name: "Модификатор: радар", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_view_booster_1":                {Name: "Модификатор: линзы", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_energy_capacity_1":             {Name: "Модификатор: емкость аккумулятора", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_energy_charging_speed_1":       {Name: "Модификатор: зарядка аккумулятора", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_improved_inventory_capacity_1": {Name: "Модификатор: объем трюма", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_power_booster_1":               {Name: "Модификатор: реактор", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_missile_damage_booster_1":      {Name: "Модификатор: урон ракет", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_laser_damage_booster_1":        {Name: "Модификатор: урон лазера", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_ballistic_damage_booster_1":    {Name: "Модификатор: урон снарядов", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_ballistic_rotate_booster_1":    {Name: "Модификатор: турель баллистики", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_missile_rotate_booster_1":      {Name: "Модификатор: турель пусковой установки", Description: "<p>Модуль уничтожается если его снять!</p>"},
		"rig_laser_rotate_booster_1":        {Name: "Модификатор: турель лазера", Description: "<p>Модуль уничтожается если его снять!</p>"},

		"dust_collector":      {Name: "Сборщик пыли", Description: "<p></p>"},
		"pump_extracted":      {Name: "Насос", Description: "Позволяет добывать жидкие ресурсы, включая нефть."},
		"inventory_scanner_1": {Name: "Сканер", Description: "Позволяет сканировать трюмы других транспортов."},
		"reverses_combine":    {Name: "Комбайн", Description: "<p>Стандартизированное устройство по сборке органики.</p>"},
	},
}
