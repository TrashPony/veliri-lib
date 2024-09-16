package descriptions

import _const "github.com/TrashPony/veliri-lib/const"

var WeaponDescription = map[string]map[string]DescriptionItem{
	_const.EN: {
		"small_missile":     {Name: "Centurion", Description: "<p> Independent development of the \"APD\" machine intelligence in order to create new weapon systems. \"Centurion\" - represented by a rocket system, assembled from samples of missile and salvo systems that once existed on the planet and trying to accommodate their best sides. In particular, it is possible to fire both \"homing\" and \"unguided\" missiles. </p>"},
		"replic_weapon_2":   {Name: "Aplex", Description: "<p> Aplex is another representative of the Replic weapon variety, which will appeal to all those who prefer to shower their enemies with an endless barrage of bullets. Aplex is a mass series of machine guns, which does not have sufficient lethality, however, the low cost of production and design features will help find their loyal followers. </p>"},
		"explores_weapon_2": {Name: "Filaz", Description: "<p> Cheap, unpretentious and massive - the best that can be said about the line of beam launchers - \"Filaz\". An outlet in this niche product from Explores is the high precision of the cutting beam hitting the target. </p>"},
		"artillery":         {Name: "S-2500", Description: "<p> A relict development, already improved by the \"APD\" machine intelligence and representing a ballistic mortar with not the best characteristics. Regardless of the damage, it mostly suffers from low hitting accuracy without preliminary zeroing. </p>"},
		"explores_weapon_4": {Name: "O-Keot", Description: "<p> O-Keot is another development along with the \"EzE-T\" capabilities of laser weapons. This time, Explores weapons designers went in a different direction and managed to develop a massive beam splitter that works on the principle of a shotgun and literally melts the target in the affected area. </p>"},
		"reverses_weapon_4": {Name: "Pin", Description: "<p> “An annoying enemy? Tired of the battle? \" - The pin will end any dispute. </p> <p> This rocket launcher of devastating power, capable of firing both homing and unguided missiles. Sacrificing in exchange for damage with your gun traverse speed, reloading and total magazine capacity. </p> <p> Well, the price will additionally turn away from the means of mass destruction many \"amateurs\". </p>"},
		"tank_gun":          {Name: "Inhibitor", Description: "<p> Independent development of the \"APD\" machine intelligence in order to create new weapon systems. \"Inhibitor\" is the most common and effective enough gun within its framework. It is capable of hitting various targets, is not too expensive to manufacture and is competitive with analogues of \"Synthets\". </p>"},
		"replic_weapon_4":   {Name: "Crusher", Description: "<p> \"Wipe your opponents to dust!\" - so the lie of the assault shotgun - \"Crusher\". This is a formidable weapon, most effective at close range, which will be able to fully reveal its potential only for those who are not afraid to face the enemy \"face to face\". </p>"},
		"explores_weapon_1": {Name: "Ocelot", Description: "<p> The Ocelot is a standardized, off-the-shelf, medium penetration laser from Explores. If you are looking for a quality product, not the highest price, and most importantly - the ability to effectively hit most targets, Ocelot will handle it. </p>"},
		"explores_weapon_3": {Name: "EzE-T", Description: "<p> If you want a machine gun, but do not want to buy tons of ammunition - \"EzE-T\" is a specialized development by Explores regarding the development of the potential of combat lasers. This frequency emitter, if it does not hit the target, will at least demoralize anyone who tries to oppose you. </p>"},
		"big_missile":       {Name: "Agra", Description: "<p> A relict development, already improved by the \"APD\" machine intelligence and representing a large rocket launcher with colossal damage, but a small magazine and a slow gun traverse speed. Against the background with analogues of \"Synthets\", \"Agra\" loses in almost all parameters, except for tangible damage. </p>"},
		"replic_weapon_1":   {Name: "Furiosa", Description: "<p> Would you like to appreciate the power of Replic? Furiosa - will give you that last opportunity. This mortar deals so devastating damage that it will not leave a trace of your opponents. The ideal solution for those who prefer to hit targets at a long distance along a hinged trajectory, sacrificing accuracy and speed of rotation of the gun. </p>"},
		"reverses_weapon_1": {Name: "Judge", Description: "<p> Annihilate space. Don't give the enemy a chance to get away. Crush with a huge salvo of eight consecutive missiles and deliver the final judgment. The Judge is the Reverses' advanced missile armament. </p>"},
		"big_laser":         {Name: "P-0 / A", Description: "<p> A relict development, already improved by the APD machine intelligence and representing a combat beam launching unit with mediocre characteristics of the destructive properties. </p>"},
		"small_laser":       {Name: "VHF", Description: "<p> Independent development of the \"APD\" machine intelligence in order to create new weapon systems. But due to the insufficient material base, as well as the use of schemes of the previous era, the product cannot boast of something outstanding, although it is not one step lower in comparison with the analogues of \"Synthets\" regarding laser-beam armament. </p>"},
		"replic_weapon_3":   {Name: "Condor", Description: "<p> If you are looking for something in between - Condor - the recommendation from Replic. This gun has managed to establish itself as reliable, powerful enough and comfortable enough to firmly take a place in the Replic army. </p>"},
		"reverses_weapon_2": {Name: "Discharge", Description: "<p> Standardized for military needs, Reverses cannon - capable of hitting targets by firing two swift missiles. The average rate of fire and the small volume of charged charges are compensated for by the quick turn of the gun. </p>"},
		"reverses_weapon_3": {Name: "Rank-Mark 2", Description: "<p> Discharge-Mark 2 is a continuation of the ideas laid down by the previous version of Discharge. This is a more advanced weapon that combines the advantages of its predecessor. Those who like to stand aside and watch an explosive show will love this product twice. </p>"},
		"far_god_1":         {Name: "LBL-X", Description: "<p>Officially, previously, never mentioned anywhere, an energy weapon that is a particle projector cannon by its characteristics. Apparently, LBL-X was created under artisanal conditions by some extremely talented synth gunsmith.</p>"},
		"far_god_2":         {Name: "Leffy", Description: "<p>Volley fire rocket launcher presented by an easily manipulable slipway with many connected barrels. It is believed that Leffy is not the most successful design solution of Reversec. It is known that a batch of this weapon was destroyed, but Terreus still managed to steal one of its samples.</p>"},
		"far_god_3":         {Name: "Ultra-10", Description: "<p>An anti-aircraft machine gun that was adapted for firing at ground targets and created by Replics. Subsequently, the 'Ultra' model was artisanal upgraded by Iair. After several unsuccessful iterations, they received the most effective variation.</p>"},
		"far_god_4":         {Name: "Impulse", Description: "<p>A creation of Replics, which is a ballistic weapon of impressive destructive power, but it did not enter mass production due to its size and overwhelming force of reverse inertia.</p>"},

		"reverses_weapon_5": {Name: "Малая пусковая установка", Description: "<p></p>"}, // TODO
		"explores_weapon_5": {Name: "explores_weapon_5", Description: "<p>Тяжелый лазер, при перегреве сбрасывает накопленную энергию создавая взрыв по направлению луча.</p>"},
		"explores_weapon_6": {Name: "explores_weapon_6", Description: "<p>Средний лучевой лазер, при перегреве увеличивает урон</p>"},
	},
	_const.RU: {
		"replic_weapon_1":   {Name: "Фуриоса", Description: "<p>Желаете по достоинству оценить мощь Replic? Фуриоса – даст вам такую последнюю возможность. Эта мортира наносит настолько сокрушительный урон, что не оставит от ваших соперников и следа. Идеальное решение для тех, кто предпочитает поражать цели на большом расстоянии по навесной траектории, пожертвовав взамен точностью и скоростью поворота орудия.</p>"},
		"replic_weapon_2":   {Name: "Аплекс", Description: "<p>Аплекс – очередной представитель оружейного разнообразия Replic, который придётся по душе всем тем, кто предпочитает поливать своих врагов нескончаемым градом пуль. Аплекс – массовая серия пулемётов, что не обладает достаточной убойностью, однако, дешевизна производства и особенности конструкции, помогут найти своих преданных приверженцев.</p>"},
		"replic_weapon_3":   {Name: "Кондор", Description: "<p>Если вы ищете нечто среднее во всём – Кондор – рекомендация от Replic. Эта пушка сумела зарекомендовать себя как безотказная, в меру мощная и достаточно удобная, чтобы прочно занять место в армии Replic.</p>"},
		"replic_weapon_4":   {Name: "Сокрушитель", Description: "<p>«Сотрите своих противников в пыль!» – так гласить солган штурмового дробовика – «Сокрушитель».  Это грозное оружие, максимально эффективное на близких дистанциях, что сумеют полностью раскрыть свой потенциал только у тех, кто не боится столкнуться с врагом «лицом к лицу».</p>"},
		"explores_weapon_1": {Name: "Оцелот", Description: "<p>Оцелот – это стандартизированный под множество нужд и серийно выпускаемый лазер средней пробивной силы от Explores. Если вы ищете качественного изделия, не самой большой цены и главное – возможности эффективно поражать большинство целей, «Оцелот» справится с этим.</p>"},
		"explores_weapon_2": {Name: "Филаз", Description: "<p>Дешёвый, неприхотливый и массовый – самое лучшее, что можно сказать о линейке лучепускателей – «Филаз».  Отдушиной в данном нишевом товаре от Explores можно считать высокую точность попадания режущего луча по цели.</p>"},
		"explores_weapon_3": {Name: "EzE-T", Description: "<p>Хотите пулемёт, но не желаете закупать тонны амуниции - «EzE-T» специализированная разработка Explores касательно развития потенциала боевых лазеров. Данный частотный испускатель, если и не попадёт в цель, то как минимум деморализует любого, кто попытается выступить против вас.</p>"},
		"explores_weapon_4": {Name: "О-Кеот", Description: "<p>О-Кеот – ещё одно развитие наряду с «EzE-T» возможностей лазерного вооружения. На сей раз, оружейные конструкторы Explores пошли по иному направлению и сумели разработать массовый делитель лучей, работающий по принципу дробовика и буквально расплавляющий цель в зоне поражения.</p>"},
		"reverses_weapon_1": {Name: "Судья", Description: "<p>Аннигилируйте пространство. Не дайте шанса врагу уйти.  Сокрушите огромным залпом из восьми последовательных ракет и вынесите окончательный приговор. Судья – передовое ракетное вооружение Reverses. </p>"},
		"reverses_weapon_2": {Name: "Разряд", Description: "<p>Стандартизированное под войсковые нужды Reverses орудие – способное путём выпуска двух стремительных ракет, поражать цели. Средний темп стрельбы и маленький объём заряжаемых зарядов, компенсируется быстротой поворота орудия.</p>"},
		"reverses_weapon_3": {Name: "Разряд-Марк 2", Description: "<p>Разряд-Марк 2 -  продолжение идей, заложенных предыдущей версией «Разряд». Это более усовершенствованное орудие, что собрало в себе плюсы предшественника. Любителям стоять в сторонке и наблюдать за взрывоопасным шоу – это изделие понравится вдвойне.</p>"},
		"reverses_weapon_4": {Name: "Булавка", Description: "<p>«Назойливый противник? Утомило сражение?» - Булавка окончит любой спор. </p> <p>Эта ракетная установка сокрушительной силы, способная вести огонь, как самонаводящимися, так и неуправляемыми ракетами. Жертвуя в обмен на урон своей скоростью поворота орудия, перезарядкой и общей вместимостью магазина.</p> <p>Ну а цена, дополнительно отвернёт от средства массового уничтожения многих «любителей».</p>"},
		"artillery":         {Name: "С-2500", Description: "<p>Реликтовая разработка, улучшенная уже машинным разумом \"APD\" и представляющая из себя баллистическую мортиру с не самыми лучшими характеристиками. Не взирая на урон, в основном страдает низкой точностью попадания без предварительной пристрелки.</p>"},
		"big_laser":         {Name: "П-0/А", Description: "<p>Реликтовая разработка, улучшенная уже машинным разумом \"APD\" и представляющая из себя боевой лучепускательный агрегат с посредственными характеристиками поражающего свойства.</p>"},
		"big_missile":       {Name: "Агра", Description: "<p>Реликтовая разработка, улучшенная уже машинным разумом \"APD\" и представляющая из себя большую ракетную установку с колоссальным уроном, но малым магазином и малой скоростью поворота орудия. На фоне с аналогами «Синтетов», «Агра» проигрывает практически по всем параметрам кроме ощутимых повреждений.</p>"},
		"small_laser":       {Name: "УКВ", Description: "<p>Самостоятельная разработка машинного разума \"APD\" с целью создания новых оружейных систем.  Но из-за недостаточной материальной базы, а также задействования схем предыдущей эпохи, изделие не может похвастаться чем-то выдающимся, хотя и не стоит на ступени ниже по сравнению с аналогами «Синтетов» касательно лазерно-лучевого вооружения.</p>"},
		"small_missile":     {Name: "Центурион", Description: "<p>Самостоятельная разработка машинного разума \"APD\" с целью создания новых оружейных систем.  «Центурион» - представлен ракетной системой, собранной из образцов некогда существовавших на планете ракетно-залповых систем и, пытающийся вместить в себе лучшие их стороны.  В частности, возможно ведения огня как «самонаводящимися», так и «неуправляемыми» ракетами.</p>"},
		"tank_gun":          {Name: "Ингибитор", Description: "<p>Самостоятельная разработка машинного разума \"APD\" с целью создания новых оружейных систем.  «Ингибитор» - самая обыденная и достаточная эффективная в своих рамках пушка. Она способна поражать различные цели, не слишком дорога в производстве и конкурентоспособна с аналогами «Синтетов».</p>"},
		"far_god_1":         {Name: "LBL-X", Description: "<p>Официально, ранее, нигде не упоминаемое энергетическое оружие, что по своим характеристикам является - пушкой-проектором частиц. Судя по всему, LBL-X был создан при кустарных условиях неким крайне талантливым синтетом-оружейником.</p>"},
		"far_god_2":         {Name: "Леффи", Description: "<p>Ракетное вооружение залпового огня, представленное легко манипулируемым стапелем со множеством соединённых стволов. Принято считать, что Леффи - не самое удачное конструкторское решение Reversec. Известно, что партия данного оружия была подвергнута уничтожению, однако Терреусу, всё-таки удалось выкрасть один из его образцов.</p>"},
		"far_god_3":         {Name: "Ультра-10", Description: "<p>Зенитный пулемёт, что был адаптирован для ведения огня по наземным целям и созданный Replics. Впоследствии, модель “Ультра” была кустарно модернизирована Иаиром. Спустя несколько неудачных итераций, получив максимально эффективную вариацию.</p>"},
		"far_god_4":         {Name: "Импульс", Description: "<p>Творение Replics, представляющее из себя баллистическое орудие внушительной поражающей силы, которое не вошло в массовое производство из-за своих габаритов и непреодолимой силы обратной инерции.</p>"},

		"reverses_weapon_5": {Name: "Малая пусковая установка", Description: "<p></p>"}, // TODO
		"explores_weapon_5": {Name: "explores_weapon_5", Description: "<p>Тяжелый лазер, при перегреве сбрасывает накопленную энергию создавая взрыв по направлению луча.</p>"},
		"explores_weapon_6": {Name: "explores_weapon_6", Description: "<p>Средний лучевой лазер, при перегреве увеличивает урон</p>"},
	},
}
