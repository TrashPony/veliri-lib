package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var ChoiceFraction = map[string]map[string]string{
	"greetings_1": {
		_const.RU: `Добро пожаловать на планету Veliri.`,
		_const.EN: `Welcome to planet Veliri.`,
	},
	"greetings_2": {
		_const.RU: `Нет времени объяснять выбирай серию машин и погнали.`,
		_const.EN: `There is no time to explain, choose a series of cars and let's go.`,
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
	},
	"Replics_weaponText": {
		_const.RU: `<span class="importantly">Replics</span> предпочитают баллистическое орудие. Это оружие обладает высокую дальность поражения и вариативность боеприпасов.`,
		_const.EN: `<span class="importantly">Replics</span> prefer ballistic weapons. These weapons have a long range and offer a variety of ammunition options.`,
	},
	"Replics_bodyText": {
		_const.RU: `<span class="importantly">Replics</span> используют тяжелые транспорты на гусеничном шасси. <span class="importantly">Гусеницы</span> позволяют крутиться на месте и не имеют штрафа при движении назад.`,
		_const.EN: `<span class="importantly">Replics</span> use heavy vehicles on a tracked chassis. The tracks allow them to spin in place and there is no movement penalty when moving backward.`,
	},
	"Replics_bonus": {
		_const.RU: `<span class="importantly">Защита</span> от всего урона повышена на <span class="importantly">5%</span>.`,
		_const.EN: `The <span class="importantly">protection</span> from all damage has been increased by <span class="importantly">5%</span>.`,
	},
	"Explores_weaponText": {
		_const.RU: `<span class="importantly">Explores</span> предпочитают лазерное орудие. Это оружие обладает высокою точностью и способны мгновенно поражать цель, однако их дальность оставляет желать лучшего.`,
		_const.EN: `<span class="importantly">Explores</span> prefer laser weapons. These weapons have high accuracy and can instantly hit the target, but their range leaves much to be desired.`,
	},
	"Explores_bodyText": {
		_const.RU: `<span class="importantly">Explores</span> используют легкие транспорты на антиграве. <span class="importantly">Антиграв</span> дает вам большую свободу и гибкость. Оно позволяет вам управлять направлением корпуса мышкой и двигаться боком без поворота корпуса.`,
		_const.EN: `<span class="importantly">Explores</span> use light vehicles with antigravity. The antigravity gives you great freedom and flexibility. It allows you to control the direction of the hull with the mouse and move sideways without turning the hull.`,
	},
	"Explores_bonus": {
		_const.RU: `<span class="importantly">Дальность обзора</span> увеличена на <span class="importantly">5%</span>.`,
		_const.EN: `The <span class="importantly">view range</span> has been increased by <span class="importantly">5%</span>.`,
	},
	"Reverses_weaponText": {
		_const.RU: `<span class="importantly">Reverses</span> предпочитают ракетные пусковые установки. Ракеты обладают высокой поражающей способностью, однако существует снаряжение способное их сбивать.`,
		_const.EN: `The <span class="importantly">Reverses</span> prefer rocket launchers. Rockets have a high destructive capacity, but there is equipment that can shoot them down.`,
	},
	"Reverses_bodyText": {
		_const.RU: `<span class="importantly">Reverses</span> используют транспорты на колесах. <span class="importantly">Колеса</span> не имеет особых особенностей. Оно является самым простым и надежным типом шасси. Оно подходит для тех, кто предпочитает классический стиль управления.`,
		_const.EN: `The <span class="importantly">Reverses</span> use wheeled vehicles. The wheels do not have any special features. They are the simplest and most reliable type of chassis. They suit those who prefer a classic style of control.`,
	},
	"Reverses_bonus": {
		_const.RU: `<span class="importantly">Вместимость трюма</span> увеличена на <span class="importantly">5%</span>.`,
		_const.EN: `The <span class="importantly">hold capacity</span> has been increased by <span class="importantly">5%</span>.`,
	},
	"button_1": {
		_const.RU: `ВПЕРЕД >`,
		_const.EN: `CHOOSE >`,
	},
	"text_1": {
		_const.RU: `Выбери сторону`,
		_const.EN: `Choose a side.`,
	},
}
