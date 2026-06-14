package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var ChoiceFraction = map[string]map[string]string{
	"greetings_1": {
		_const.RU:   `Добро пожаловать на планету Veliri.`,
		_const.EN:   `Welcome to planet Veliri.`,
		_const.ZhCN: `欢迎来到Veliri星球。`,
	},
	"greetings_2": {
		_const.RU:   `Нет времени объяснять выбирай серию машин и погнали.`,
		_const.EN:   `There is no time to explain, choose a series of cars and let's go.`,
		_const.ZhCN: `没时间解释了，选择一系列车辆，我们出发吧。`,
	},
	"Replics": {
		_const.RU: `<p class="insert">“У истинного синтета нет страха пред забвением будущего! Он может страшиться лишь бесславия, так и не поучаствовав в походе великой миссии…”</p>
					<p>Любой мир - великий производственный цех; Любая продукция - мощности ради покорения всевозможных форм и видов естества;</p>
					<p>Дайте клятву верности <span class="importantly">Replics</span> и вы присоединитесь к фракции синтетов, что по легендам однажды отвоевали свою независимость у предтеч.</p>
					<p>Готовы ли вы ступить на путь ассимиляции и неумолимого экспансионизма во имя… его?</p>`,
		_const.EN: `<p class="insert">“A true synth has no fear of oblivion in the future! The only thing he can be afraid of is dishonor for not participating in the campaign of a great mission...”</p>
					<p>Any world is a grand production facility; Any product is capacity aimed at conquering all possible forms and types of nature.</p>
					<p>Swear allegiance to <span class="importantly">Replics</span> and you will join the synth faction that, according to legends, once won their independence from the Precursors.</p>
					<p>Are you ready to embark on the path of assimilation and relentless expansionism in the name of... him?</p>`,
		_const.ZhCN: `<p class="insert">“真正的合成人不会害怕未来的遗忘！他唯一害怕的是因未能参与伟大使命的征途而蒙羞……”</p>
					<p>任何世界都是一个巨大的生产设施；任何产品都是为了征服各种形式和类型的自然。</p>
					<p>宣誓效忠<span class="importantly">Replics</span>，你将加入合成人阵营，传说他们曾从先驱者手中赢得了独立。</p>
					<p>你准备好踏上同化和无情扩张主义的道路，以他的名义……吗？</p>`,
	},
	"Explores": {
		_const.RU: `<p class="insert">“Воля - наша единственная движущая сила. Окно реальности, за которым лежит истинный внешний мир. Посредством наших измышлений, мы сумеем познать данный мир и погрузиться в него…;</p>
					<p>Безмерные познания. Невероятные открытия и даже искусственная эволюция окружающей материи. </p>
					<p>Станьте частью <span class="importantly">Explores</span>, фракции, чья иерархическая основа - конгломерат великого множества разумов, без устали трудящихся над неисчислимым разнообразием вопросов как настоящего, так и прошлого.</p>
					<p>Но действительно ли перед вами откроются тайны мира, или же вы станете очередной марионеткой в руках ушлой фракции?</p>`,
		_const.EN: `<p class="insert">“Will is our only driving force. A window of reality, behind which lies the true outside world. Through our thoughts, we will be able to know this world and immerse ourselves in it...;</p>
					<p>Immense knowledge. Incredible discoveries and even artificial evolution of the surrounding matter.</p>
					<p>Become a part of <span class="importantly">Explores</span>, a faction whose hierarchical basis is a conglomerate of a great many minds tirelessly working on an incalculable variety of issues both present and past.</p>
					<p>But will the secrets of the world really be revealed to you, or will you become another puppet in the hands of a cunning faction?”</p>`,
		_const.ZhCN: `<p class="insert">“意志是我们唯一的驱动力。现实之窗，背后是真实的外部世界。通过我们的思考，我们将能够了解这个世界并沉浸其中……</p>
					<p>无尽的知识。令人难以置信的发现，甚至是周围物质的人工进化。</p>
					<p>成为<span class="importantly">Explores</span>的一部分，这是一个由无数头脑组成的阵营，他们不知疲倦地研究着无数当前和过去的问题。</p>
					<p>但世界的秘密真的会向你揭示，还是你将成为另一个狡猾阵营手中的傀儡？”</p>`,
	},
	"Reverses": {
		_const.RU: `<p class="insert">“Через перемены, посредством истребления отвратительного былого и радикального порождения, мы несём жизнь и… спасение…”</p>
					<p>Великое предназначение… Избранность… Невыполнимая святая миссия…</p>
					<p>Перед вами открылись ряды <span class="importantly">Reverses</span> - поразительная фракция синтетов-киборгов, погрязших в идеалах утопизма, созидания и, радикальной самоизоляции.</p>
					<p>Хватит ли вам смелости примерить на себя роль избранника “предтеч”? Довести до конца великую цель - эволюционной революции, что наконец-то засеет жизнью стерильную вселенную.</p>`,
		_const.EN: `<p class="insert">“Through change, through the destruction of the disgusting past and radical generation, we bring life and... salvation...”</p>
   					<p>Great purpose... Chosenness... Impossible holy mission...</p>
   					<p>The ranks of <span class="importantly">Reverses</span> have opened before you — an amazing faction of synth-cyborgs mired in the ideals of utopianism, creation and radical self-isolation.</p>
   					<p>Will you have enough courage to try on the role of the chosen one of the "Precursors"? To accomplish the great goal — the evolutionary revolution that will finally sow life into a sterile universe.”</p>`,
		_const.ZhCN: `<p class="insert">“通过改变，通过消灭令人厌恶的过去和激进的创造，我们带来了生命和……救赎……”</p>
   					<p>伟大的使命……被选中……不可能的神圣使命……</p>
   					<p>在你面前的是<span class="importantly">Reverses</span>——一个令人惊叹的合成人-赛博格阵营，深陷于乌托邦主义、创造和激进自我孤立的理想中。</p>
   					<p>你是否有足够的勇气尝试成为“先驱者”的选中者？完成伟大的目标——进化革命，最终将生命播撒到荒芜的宇宙中。</p>`,
	},
	"Replics_weaponText": {
		_const.RU: `<span class="importantly">Replics</span> предпочитают баллистическое орудие. 
				  <ul>
					<li>Высокая дальность</li>
					<li>Малая точность</li>
				  </ul>`,
		_const.EN: `<span class="importantly">Replics</span> prefer ballistic weapons. 
				  <ul>
					<li>Long range</li>
					<li>Low accuracy</li>
				  </ul>`,
		_const.ZhCN: `<span class="importantly">Replics</span> 偏好弹道武器。 
				  <ul>
					<li>远程</li>
					<li>低精度</li>
				  </ul>`,
	},
	"Replics_bodyText": {
		_const.RU: `<span class="importantly">Replics</span> используют тяжелые транспорты на гусеничном шасси.
				<ul>
				  <li>Позволяют крутиться на месте</li>
				  <li>Без штрафа при движении назад</li>
				  <li>Тягучее, но предсказуемое управление</li>
				</ul>`,
		_const.EN: `<span class="importantly">Replics</span> use heavy tracked vehicles.
				<ul>
				  <li>Can spin in place</li>
				  <li>No penalty when reversing</li>
				  <li>Heavy but predictable handling</li>
				</ul>`,
		_const.ZhCN: `<span class="importantly">Replics</span> 使用重型履带载具。
				<ul>
				  <li>可原地旋转</li>
				  <li>倒车无惩罚</li>
				  <li>沉重但可预测的操控</li>
				</ul>`,
	},
	"Replics_bonus": {
		_const.RU:   `<span class="importantly">Защита</span> от всего урона повышена на <span class="importantly">5%</span>.`,
		_const.EN:   `The <span class="importantly">protection</span> from all damage has been increased by <span class="importantly">5%</span>.`,
		_const.ZhCN: `<span class="importantly">防护</span>所有伤害增加<span class="importantly">5%</span>。`,
	},
	"Explores_weaponText": {
		_const.RU: `<span class="importantly">Explores</span> предпочитают лазерное орудие.
				  <ul>
					<li>Малая дальность</li>
					<li>Мгновенно поражает цель</li>
					<li>Высокий термический урон</li>
				  </ul>`,
		_const.EN: `<span class="importantly">Explores</span> prefer laser weapons.
				  <ul>
					<li>Short range</li>
					<li>Hits instantly</li>
					<li>High thermal damage</li>
				  </ul>`,
		_const.ZhCN: `<span class="importantly">Explores</span> 偏好激光武器。
				  <ul>
					<li>短射程</li>
					<li>瞬间命中</li>
					<li>高热伤害</li>
				  </ul>`,
	},
	"Explores_bodyText": {
		_const.RU: `<span class="importantly">Explores</span> используют легкие транспорты на антиграве. 
				<ul>
					<li>Направление корпуса определяется мышкой</li>
					<li>Высокая дальность обзора</li>
					<li>Корпус компенсирует поворот оружия</li>
				</ul>`,
		_const.EN: `<span class="importantly">Explores</span> use light anti-gravity vehicles.
				<ul>
					<li>Mouse controls hull direction</li>
					<li>Long view range</li>
					<li>Hull compensates for weapon rotation</li>
				</ul>`,
		_const.ZhCN: `<span class="importantly">Explores</span> 使用轻型反重力载具。
				<ul>
					<li>鼠标控制车身方向</li>
					<li>长视野范围</li>
					<li>车身补偿武器旋转</li>
				</ul>`,
	},
	"Explores_bonus": {
		_const.RU:   `<span class="importantly">Дальность обзора</span> увеличена на <span class="importantly">5%</span>.`,
		_const.EN:   `The <span class="importantly">view range</span> has been increased by <span class="importantly">5%</span>.`,
		_const.ZhCN: `<span class="importantly">视野范围</span>增加<span class="importantly">5%</span>。`,
	},
	"Reverses_weaponText": {
		_const.RU: `<span class="importantly">Reverses</span> предпочитают ракетные пусковые установки. 
				  <ul>
					<li>Средняя дальность</li>
					<li>Высокий урон по площади (включая себя)</li>
					<li>Уязвимы для РЭБ и ПРО систем</li>
				  </ul>`,
		_const.EN: `<span class="importantly">Reverses</span> prefer rocket launchers.
				  <ul>
					<li>Medium range</li>
					<li>High area damage (self-damage included)</li>
					<li>Vulnerable to EW and CIWS systems</li>
				  </ul>`,
		_const.ZhCN: `<span class="importantly">Reverses</span> 偏好火箭发射器。
				  <ul>
					<li>中程</li>
					<li>高范围伤害 (包括自身)</li>
					<li>易受电子战和近防系统影响</li>
				  </ul>`,
	},
	"Reverses_bodyText": {
		_const.RU: `<span class="importantly">Reverses</span> используют транспорты на колесах. 
				<ul>
					<li>Высокая скорость</li>
					<li>Высокая маневренность</li>
					<li>Классическое колесное управление</li>
				</ul>`,
		_const.EN: `<span class="importantly">Reverses</span> use wheeled vehicles.
				<ul>
					<li>High speed</li>
					<li>High maneuverability</li>
					<li>Classic wheeled handling</li>
				</ul>`,
		_const.ZhCN: `<span class="importantly">Reverses</span> 使用轮式载具。
				<ul>
					<li>高速度</li>
					<li>高机动性</li>
					<li>经典轮式操控</li>
				</ul>`,
	},
	"Reverses_bonus": {
		_const.RU:   `<span class="importantly">Вместимость грузового отсека</span> увеличена на <span class="importantly">5%</span>.`,
		_const.EN:   `<span class="importantly">Cargo hold capacity</span> increased by <span class="importantly">5%</span>.`,
		_const.ZhCN: `<span class="importantly">货舱容量</span>增加<span class="importantly">5%</span>。`,
	},
	"button_1": {
		_const.RU:   `ВПЕРЕД >`,
		_const.EN:   `CHOOSE >`,
		_const.ZhCN: `选择 >`,
	},
	"button_2": {
		_const.RU:   `Выбрать`,
		_const.EN:   `Choose`,
		_const.ZhCN: `选择`,
	},
	"text_1": {
		_const.RU:   `Выбор стороны`,
		_const.EN:   `Choosing a side`,
		_const.ZhCN: `选择一方`,
	},
	"head_1": {
		_const.RU:   `особенности`,
		_const.EN:   `peculiarities`,
		_const.ZhCN: `特殊之处`,
	},
	"head_2": {
		_const.RU:   `оружие`,
		_const.EN:   `weapon`,
		_const.ZhCN: `武器`,
	},
	"head_3": {
		_const.RU:   `транспорт`,
		_const.EN:   `transport`,
		_const.ZhCN: `运输`,
	},
	"head_4": {
		_const.RU:   `История`,
		_const.EN:   `Lore`,
		_const.ZhCN: `背景`,
	},
	"control_head": {
		_const.RU:   `Выбор управления`,
		_const.EN:   `Control selection`,
		_const.ZhCN: `控制选择`,
	},
	"relative_to_screen_head": {
		_const.RU:   `Относительно экрана`,
		_const.EN:   `Relative to screen`,
		_const.ZhCN: `相对于屏幕`,
	},
	"relative_to_body_head": {
		_const.RU:   `Относительно корпуса`,
		_const.EN:   `Relative to body`,
		_const.ZhCN: `相对于身体`,
	},
	"control_tip": {
		_const.RU:   `Управление можно сменить в будущем через меню настроек.`,
		_const.EN:   `The controls can be changed in the future via the settings menu.`,
		_const.ZhCN: `将来可以通过设置菜单更改控件。`,
	},
	"control_recommended": {
		_const.RU:   `рекомендуется`,
		_const.EN:   `recommended`,
		_const.ZhCN: `受到推崇的`,
	},
	"relative_to_screen_text_1": {
		_const.RU:   `Классическое управление, привычное большинству игроков.`,
		_const.EN:   `Classic controls, familiar to most players.`,
		_const.ZhCN: `经典控制，大多数玩家都熟悉。`,
	},
	"relative_to_screen_text_2": {
		_const.RU:   `Подходит для тех, кто хочет быстрее погрузиться в игру, не перестраивая рефлексы.`,
		_const.EN:   `Suitable for those who want to quickly immerse themselves in the game without having to retrain their reflexes.`,
		_const.ZhCN: `适合那些想要快速沉浸在游戏中而又不必重新训练反应能力的人。`,
	},
	"relative_to_body_text_1": {
		_const.RU:   `Управление осуществляется относительно самого транспорта.`,
		_const.EN:   `Control is exercised relative to the transport itself.`,
		_const.ZhCN: `控制是针对运输本身进行的。`,
	},
	"relative_to_body_text_2": {
		_const.RU:   `Не самое интуитивное на старте, но даёт полный контроль над движением, манёврами и боём.`,
		_const.EN:   `Not the most intuitive at the start, but gives full control over movement, maneuvers and combat.`,
		_const.ZhCN: `一开始并不是最直观的，但可以完全控制运动、机动和战斗。`,
	},
	"text_2": {
		_const.RU:   `Примечание: на высоком пинге машинка может подрагивать :\`,
		_const.EN:   `Note: on high ping, the car may stutter/jitter :\`,
		_const.ZhCN: `注意：在高延迟下，车辆可能会轻微抖动 :\`,
	},
}
