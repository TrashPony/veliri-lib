package handbook

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"math/rand"
	"strings"
	"time"
)

const (
	Neutral = "neutral" // "Привет"
	Good    = "good"    // "Привет, друган"
	Bad     = "bad"     // "Убирайся, отвали", тут смысл в том что отношение плохое, но это не бой
	Target  = "target"  // А этот статус говорит что игрок на прицеле у нпс, и он его преследует
	Fear    = "fear"    // "Отстань от меня!"
	Help    = "help"    // Когда бот в страхе, но не из-за игрока
)

var npcGreetings = map[string]map[string]string{
	// Replicas
	// Neutral
	_const.Replicas + "_miner_" + Neutral + "_1": {
		_const.RU: "Ты что-то хотел, %UserName%?",
		_const.EN: "Did you want something, %UserName%?",
	},
	_const.Replicas + "_miner_" + Neutral + "_2": {
		_const.RU: "Я весь в работе и у меня мало времени.",
		_const.EN: "I'm busy with work and have little time.",
	},
	_const.Replicas + "_miner_" + Neutral + "_3": {
		_const.RU: "Чего бы ты от меня ни желал, поторопись с этим.",
		_const.EN: "Whatever you want from me, hurry up with it.",
	},
	_const.Replicas + "_miner_" + Neutral + "_4": {
		_const.RU: "Руда… Руда… Руда… И так цикл за циклом.",
		_const.EN: "Ore... Ore... Ore... And so cycle after cycle.",
	},
	_const.Replicas + "_miner_" + Neutral + "_5": {
		_const.RU: "Только не говори, %UserName%, что ты тоже будешь меня донимать о скрытых месторождениях Тория.",
		_const.EN: "Don't tell me, %UserName%, that you will also pester me about hidden thorium deposits.",
	},
	_const.Replicas + "_miner_" + Neutral + "_6": {
		_const.RU: "Слышал уже? Цены на руды вновь обвалились. Судя по всему, кто-то неумело пользуется рынком, излишне перенасыщая его.",
		_const.EN: "Have you heard it already? Ore prices have fallen again. Apparently, someone is using the market ineptly, oversaturating it.",
	},
	_const.Replicas + "_miner_" + Neutral + "_7": {
		_const.RU: "Знаешь, %UserName%, я изрядно нервничаю, когда незнакомый синтет вдруг решает затеять со мной беседу.",
		_const.EN: "You know, %UserName%, I get pretty nervous when an unfamiliar synth suddenly decides to start a conversation with me.",
	},
	_const.Replicas + "_miner_" + Neutral + "_8": {
		_const.RU: "Пиратские Кластеры, безумные машины из пустошей, непомерные налоговые ставки фракций… Да зачем я вообще решил всем этим заниматься!?",
		_const.EN: "Pirate Clusters, crazy machines from the wastelands, exorbitant tax rates of factions... Why did I even decide to do all this!?",
	},

	_const.Replicas + "_guard_" + Neutral + "_1": {
		_const.RU: "Есть ли тебе о чём стоит сообщить, %UserName%?",
		_const.EN: "Do you have anything to report, %UserName%?",
	},
	_const.Replicas + "_guard_" + Neutral + "_2": {
		_const.RU: "Ну, говори-говори! Я весь в ожидании.",
		_const.EN: "Well, speak up! I'm all in anticipation.",
	},
	_const.Replicas + "_guard_" + Neutral + "_3": {
		_const.RU: "Судя по твоему виду, я кое-что должен знать.",
		_const.EN: "By the look of you, I should know something.",
	},
	_const.Replicas + "_guard_" + Neutral + "_4": {
		_const.RU: "Пробудешь хотя бы цикл в патруле по пустошам, сразу заскучаешь о войне фракций.",
		_const.EN: "Once you spend at least a cycle on patrol in the wastelands, you will immediately become bored with the faction war.",
	},
	_const.Replicas + "_guard_" + Neutral + "_5": {
		_const.RU: "Дай угадаю, ты видел что-то и теперь хочешь мне об этом сообщить?",
		_const.EN: "Let me guess, you saw something and now you want to tell me about it?",
	},
	_const.Replicas + "_guard_" + Neutral + "_6": {
		_const.RU: "Либо у тебя возникли проблемы с кем-то конкретным, либо с законами фракций.",
		_const.EN: "Either you have problems with someone specific, or with the laws of factions.",
	},
	_const.Replicas + "_guard_" + Neutral + "_7": {
		_const.RU: "Я воплощение справедливости.",
		_const.EN: "I am the embodiment of justice.",
	},
	_const.Replicas + "_guard_" + Neutral + "_8": {
		_const.RU: "Я несу закон, и абсолютно все обязаны ему следовать!",
		_const.EN: "I bring the law, and absolutely everyone is obliged to follow it!",
	},
	_const.Replicas + "_guard_" + Neutral + "_9": {
		_const.RU: "Мои полномочия уникальны. Помни об этом, %UserName%.",
		_const.EN: "My credentials are unique. Remember this, %UserName%.",
	},

	_const.Replicas + "_out_scout_" + Neutral + "_1": {
		_const.RU: "Любопытно. У тебя имеется ко мне некое предложение, %UserName%?",
		_const.EN: "Curious. Do you have a proposal for me, %UserName%?",
	},
	_const.Replicas + "_out_scout_" + Neutral + "_2": {
		_const.RU: "Видишь ли, я на экстра важном задании.",
		_const.EN: "You see, I'm on an extra important mission.",
	},
	_const.Replicas + "_out_scout_" + Neutral + "_3": {
		_const.RU: "Нет времени, нет времени на всё это.",
		_const.EN: "There is no time, there is no time for all this.",
	},
	_const.Replicas + "_out_scout_" + Neutral + "_4": {
		_const.RU: "Неужто ты следил за мной?",
		_const.EN: "Could it be that you were following me?",
	},
	_const.Replicas + "_out_scout_" + Neutral + "_5": {
		_const.RU: "Какая неожиданная встреча, но только не для меня.",
		_const.EN: "What an unexpected meeting, but not for me.",
	},
	_const.Replicas + "_out_scout_" + Neutral + "_6": {
		_const.RU: "Хм, я запомню тебя, %UserName%.",
		_const.EN: "Hmm, I'll remember you, %UserName%.",
	},

	_const.Replicas + "_in_scout_" + Neutral + "_1": {
		_const.RU: "О-о, вы только гляньте, кто ко мне пожаловал!",
		_const.EN: "",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_2": {
		_const.RU: "Не зли меня, %UserName%, если только не желаешь огрести потом проблем.",
		_const.EN: "Don't make me angry, %UserName%, unless you want to cause problems later.",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_3": {
		_const.RU: "Я недолюбливаю особо болтливых.",
		_const.EN: "I don't like particularly talkative people.",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_4": {
		_const.RU: "Хочешь поторговаться о сведениях?",
		_const.EN: "Do you want to bargain for information?",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_6": {
		_const.RU: "И не надейся, в моём отряде для тебя места не будет.",
		_const.EN: "And don’t get your hopes up, there won’t be a place for you in my squad.",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_7": {
		_const.RU: "Хм, может, мне стоит посмотреть, что ты там перевозишь в своём трюме?",
		_const.EN: "Hmm, maybe I should see what you're carrying in your hold?",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_8": {
		_const.RU: "Двигай далее, не задерживайся.",
		_const.EN: "Move on, don't linger.",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_9": {
		_const.RU: "Я ничего не покупаю и не продаю! Мои интересы... более специфичны.",
		_const.EN: "I don't buy or sell anything! My interests... are more specific.",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_10": {
		_const.RU: "Сегодня я при хорошем настроении, так что прощаю твою наглость.",
		_const.EN: "I'm in a good mood today, so I forgive your impudence.",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_11": {
		_const.RU: "Только дай мне повод, %UserName%, и я прострелю твою подвеску.",
		_const.EN: "Just give me a reason, %UserName%, and I'll shoot through your pendant.",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_12": {
		_const.RU: "Может, подскажешь твой дальнейший маршрут передвижения?",
		_const.EN: "Maybe you can tell me your future route?",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_13": {
		_const.RU: "Знаешь, %UserName%, в пустошах может приключиться всякое.",
		_const.EN: "You know, %UserName%, anything can happen in the wasteland.",
	},

	_const.Replicas + "_transport_" + Neutral + "_1": {
		_const.RU: "Я просто вожу грузы.",
		_const.EN: "I'm just hauling loads.",
	},
	_const.Replicas + "_transport_" + Neutral + "_2": {
		_const.RU: "Туда-сюда, от базы к базе.",
		_const.EN: "Back and forth, from base to base.",
	},
	_const.Replicas + "_transport_" + Neutral + "_3": {
		_const.RU: "Честно говоря, мой доход не самый внушительный. Лучше бы я пошёл в добытчики.",
		_const.EN: "To be honest, my income is not the most impressive. It would be better if I became a miner.",
	},
	_const.Replicas + "_transport_" + Neutral + "_4": {
		_const.RU: "Зачем ты меня тревожишь?",
		_const.EN: "Why are you bothering me?",
	},
	_const.Replicas + "_transport_" + Neutral + "_5": {
		_const.RU: "Разве не видно - я перевожу ценный груз!",
		_const.EN: "Can't you see - I'm transporting valuable cargo!",
	},
	_const.Replicas + "_transport_" + Neutral + "_6": {
		_const.RU: "Надеюсь... ты... не будешь мне угрожать?",
		_const.EN: "I hope... you... won't threaten me?",
	},
	_const.Replicas + "_transport_" + Neutral + "_7": {
		_const.RU: "Прочь с моего пути! У меня важная доставка!",
		_const.EN: "Get out of my way! I have an important delivery!",
	},
	_const.Replicas + "_transport_" + Neutral + "_8": {
		_const.RU: "Как же ты не вовремя %UserName%!",
		_const.EN: "How are you at the wrong time %UserName%!",
	},
	_const.Replicas + "_transport_" + Neutral + "_9": {
		_const.RU: "Ну почему, как только ты берёшь контракт по доставке, всем сразу от тебя чего-то нужно!?",
		_const.EN: "Why, as soon as you take on a delivery contract, does everyone immediately need something from you!?",
	},
	_const.Replicas + "_transport_" + Neutral + "_10": {
		_const.RU: "Отвечаю сразу: я ничего не продаю!",
		_const.EN: "I answer right away: I don’t sell anything!",
	},
	_const.Replicas + "_transport_" + Neutral + "_11": {
		_const.RU: "Нет, я не поверю, что ты из стражей. И нет, я не позволю тебе досмотреть мой трюм.",
		_const.EN: "No, I won't believe that you are one of the guards. And no, I won't let you search my hold.",
	},
	_const.Replicas + "_transport_" + Neutral + "_12": {
		_const.RU: "Ты друг, или... враг?",
		_const.EN: "Are you a friend, or... an enemy?",
	},
	_const.Replicas + "_transport_" + Neutral + "_13": {
		_const.RU: "Пытаешься вынюхать мой маршрут?",
		_const.EN: "Are you trying to sniff out my route?",
	},
	_const.Replicas + "_transport_" + Neutral + "_14": {
		_const.RU: "Транспортник на связи.",
		_const.EN: "The transport worker is in touch.",
	},
	_const.Replicas + "_transport_" + Neutral + "_15": {
		_const.RU: "А-а, ты наверное один из недовольных клиентов?",
		_const.EN: "Ah, you must be one of the dissatisfied customers?",
	},

	_const.Replicas + "_builder_" + Neutral + "_1": {
		_const.RU: "Я строю во имя фракции.",
		_const.EN: "I build in the name of the faction.",
	},
	_const.Replicas + "_builder_" + Neutral + "_2": {
		_const.RU: "Чёрт побери, я так обожаю свою работу!",
		_const.EN: "Damn it, I love my job so much!",
	},
	_const.Replicas + "_builder_" + Neutral + "_3": {
		_const.RU: "У-у, у меня тут дел на несколько циклов вперёд.",
		_const.EN: "Ooh, I've got a few cycles of work to do here.",
	},
	_const.Replicas + "_builder_" + Neutral + "_4": {
		_const.RU: "Решил по отвлекать меня? Ладно, заодно и перерыв будет.",
		_const.EN: "Decided to distract me? Okay, there will be a break at the same time.",
	},
	_const.Replicas + "_builder_" + Neutral + "_5": {
		_const.RU: "Хочешь тоже податься в строители %UserName%? Поздно, места все заняты.",
		_const.EN: "Do you also want to join the builders of %UserName%? It's late, the seats are all taken.",
	},
	_const.Replicas + "_builder_" + Neutral + "_6": {
		_const.RU: "Поди пойми, где должна располагаться эта сеть турелей…",
		_const.EN: "Just figure out where this network of turrets should be located...",
	},
	_const.Replicas + "_builder_" + Neutral + "_7": {
		_const.RU: "Я возвожу строения, а не оказываю информационные услуги.",
		_const.EN: "I erect buildings, and do not provide information services.",
	},
	_const.Replicas + "_builder_" + Neutral + "_8": {
		_const.RU: "Ты как раз находишься на одном из запланированных для строительства участков.",
		_const.EN: "You are just on one of the planned construction sites.",
	},
	_const.Replicas + "_builder_" + Neutral + "_9": {
		_const.RU: "Моя работа в самом разгаре, %UserName%.",
		_const.EN: "My work is in full swing, %UserName%.",
	},
	_const.Replicas + "_builder_" + Neutral + "_10": {
		_const.RU: "Инструменты ещё не остыли, мне некогда на тебя отвлекаться.",
		_const.EN: "The instruments have not yet cooled down, I have no time to be distracted by you.",
	},

	_const.Replicas + "_warrior_" + Neutral + "_1": {
		_const.RU: "Нахожусь в патруле.",
		_const.EN: "I'm on patrol.",
	},
	_const.Replicas + "_warrior_" + Neutral + "_2": {
		_const.RU: "Я защищаю границы фракции.",
		_const.EN: "I protect the faction's borders.",
	},
	_const.Replicas + "_warrior_" + Neutral + "_3": {
		_const.RU: "В поиске новых целей.",
		_const.EN: "In search of new targets.",
	},
	_const.Replicas + "_warrior_" + Neutral + "_4": {
		_const.RU: "Где-то просочился враг, %UserName%?",
		_const.EN: "Has the enemy leaked somewhere, %UserName%?",
	},
	_const.Replicas + "_warrior_" + Neutral + "_5": {
		_const.RU: "Всегда готов к бою!",
		_const.EN: "Always ready for battle!",
	},
	_const.Replicas + "_warrior_" + Neutral + "_6": {
		_const.RU: "Я готов... к чему угодно.",
		_const.EN: "I'm ready... for anything.",
	},
	_const.Replicas + "_warrior_" + Neutral + "_7": {
		_const.RU: "Все стволы заряжены.",
		_const.EN: "All barrels are loaded.",
	},
	_const.Replicas + "_warrior_" + Neutral + "_8": {
		_const.RU: "Лента боекомплекта подана.",
		_const.EN: "The ammunition belt has been supplied.",
	},
	_const.Replicas + "_warrior_" + Neutral + "_9": {
		_const.RU: "Ох, сейчас бы добротную заварушку!",
		_const.EN: "Oh, now there would be a good mess!",
	},
	_const.Replicas + "_warrior_" + Neutral + "_10": {
		_const.RU: "Жду не дождусь, когда произойдет новый виток войны фракций!",
		_const.EN: "I can't wait for the next round of faction war to happen!",
	},
	_const.Replicas + "_warrior_" + Neutral + "_11": {
		_const.RU: "Лучше не зли меня, %UserName%.",
		_const.EN: "You better not make me angry, %UserName%.",
	},
	_const.Replicas + "_warrior_" + Neutral + "_12": {
		_const.RU: "Тоже хочешь подраться, %UserName%?",
		_const.EN: "Do you want to fight too, %UserName%?",
	},

	_const.Replicas + "_expedition_" + Neutral + "_1": {
		_const.RU: "Какое же это убогое место!",
		_const.EN: "What a wretched place this is!",
	},
	_const.Replicas + "_expedition_" + Neutral + "_2": {
		_const.RU: "Я бывал там, где ты, %UserName%, даже и представить себе не можешь.",
		_const.EN: "I've been to places you, %UserName%, can't even imagine.",
	},
	_const.Replicas + "_expedition_" + Neutral + "_3": {
		_const.RU: "Я видел многое и могу открыто заявить тебе, %UserName% - кое-что крайне важное от вас фракции всё-таки скрывают.",
		_const.EN: "I have seen a lot and I can openly tell you, %UserName% - the factions are still hiding something extremely important from you.",
	},
	_const.Replicas + "_expedition_" + Neutral + "_4": {
		_const.RU: "За безопасными секторами совсем ничего нету. Имей это ввиду.",
		_const.EN: "There is absolutely nothing beyond the safe sectors. Keep this in mind.",
	},
	_const.Replicas + "_expedition_" + Neutral + "_5": {
		_const.RU: "Не вздумай приближаться к местной органической фауне, %UserName%, они… не тактильны.",
		_const.EN: "Don't even try to get close to the local organic fauna, %UserName%, they... are not tactile.",
	},
	_const.Replicas + "_expedition_" + Neutral + "_6": {
		_const.RU: "Если случайно забредёшь в реликтовый город, считай, что ты потерял там свой корпус.",
		_const.EN: "If you accidentally wander into a relic city, consider that you have lost your corps there.",
	},
	_const.Replicas + "_expedition_" + Neutral + "_7": {
		_const.RU: "Если вдруг станешь принимать посреди пустошей странный сигнал, немедленно оттуда уходи!",
		_const.EN: "If you suddenly start receiving a strange signal in the middle of the wasteland, leave immediately!",
	},
	_const.Replicas + "_expedition_" + Neutral + "_8": {
		_const.RU: "Я участник экспедиции, а ты?",
		_const.EN: "I'm a member of the expedition, and you?",
	},
	_const.Replicas + "_expedition_" + Neutral + "_9": {
		_const.RU: "Скоро снова в путь.",
		_const.EN: "We'll be on the road again soon.",
	},
	_const.Replicas + "_expedition_" + Neutral + "_10": {
		_const.RU: "Только самые смелые, %UserName%, записываются в экспедиции.",
		_const.EN: "Only the bravest, %UserName%, sign up for the expedition.",
	},
	_const.Replicas + "_expedition_" + Neutral + "_11": {
		_const.RU: "Если тебе нечего терять, %UserName%, примкни к нам.",
		_const.EN: "If you have nothing to lose, %UserName%, join us.",
	},

	// Good
	_const.Replicas + "_miner_" + Good + "_1": {
		_const.RU: "Всегда рад встретить иного дружественного синтета. В нынешнее время, это прямо-таки роскошь.",
		_const.EN: "Always glad to meet another friendly synth. Nowadays, this is downright luxury.",
	},
	_const.Replicas + "_miner_" + Good + "_2": {
		_const.RU: "Как думаешь, нам - добытчикам, должны доплачивать за фактор опасности?",
		_const.EN: "Do you think we, the miners, should be paid extra for the danger factor?",
	},
	_const.Replicas + "_miner_" + Good + "_3": {
		_const.RU: "Скажу по секрету: в этом секторе ты мало чем сумеешь разжиться. Попробуй осмотреть прилегающие территории.",
		_const.EN: "I’ll tell you a secret: you won’t be able to get hold of much in this sector. Try to explore the surrounding areas.",
	},
	_const.Replicas + "_miner_" + Good + "_4": {
		_const.RU: "%UserName%, если у тебя есть соответствующая руда, попробуй наведаться на %BaseName%. Сейчас, там можно выгодно реализовать руду.",
		_const.EN: "%UserName%, if you have the appropriate ore, try visiting %BaseName%. Now, ore can be profitably sold there.",
	},
	_const.Replicas + "_miner_" + Good + "_5": {
		_const.RU: "Тоже решил заняться добычей руды, а по итогу не особо то и идёт дело?",
		_const.EN: "You also decided to start mining ore, but in the end things aren’t going well?",
	},
	_const.Replicas + "_miner_" + Good + "_6": {
		_const.RU: "Сегодня мне повезло как никогда ранее! Столько руды я ещё никогда за всё своё существование не добывал.",
		_const.EN: "Today I am luckier than ever before! I have never mined so much ore in my entire existence.",
	},
	_const.Replicas + "_miner_" + Good + "_7": {
		_const.RU: "Только между нами %UserName%: если рассчитываешь выгодно переработать руду, советую на время позабыть о %BaseName%. Ты ещё и останешься должным тамошнему диспетчеру.",
		_const.EN: "Just between us %UserName%: if you expect to profitably recycle ore, I advise you to forget about %BaseName% for a while. You will also remain indebted to the dispatcher there.",
	},
	_const.Replicas + "_miner_" + Good + "_8": {
		_const.RU: "Ходят слухи, что в %SectorName% обнаружили месторождения Тория. Сам я не видел. Да и дорога туда опасная. Но… вдруг, ты %UserName% заинтересуешься этим?",
		_const.EN: "There are rumors that thorium deposits have been discovered in %SectorName%. I haven't seen it myself. And the road there is dangerous. But... what if you %UserName% become interested in this?",
	},

	_const.Replicas + "_guard_" + Good + "_1": {
		_const.RU: "О тебе хорошо отзываются, %UserName%. Я это учту.",
		_const.EN: "You are well spoken of, %UserName%. I'll take this into account.",
	},
	_const.Replicas + "_guard_" + Good + "_2": {
		_const.RU: "Добропорядочные синтеты наподобие тебя, %UserName% - редкость.",
		_const.EN: "Respectable synths like you, %UserName%, are rare.",
	},
	_const.Replicas + "_guard_" + Good + "_3": {
		_const.RU: "Пускай твой дальнейший путь будет безопасен, %UserName%.",
		_const.EN: "May your future journey be safe, %UserName%.",
	},
	_const.Replicas + "_guard_" + Good + "_4": {
		_const.RU: "У нас здесь всё под полным контролем, а ты под надёжной защитой, %UserName%.",
		_const.EN: "We have everything here under complete control, and you are under reliable protection, %UserName%.",
	},
	_const.Replicas + "_guard_" + Good + "_5": {
		_const.RU: "Патруль, докладываю: у нас всё спокойно.",
		_const.EN: "Patrol, I report: everything is calm here.",
	},
	_const.Replicas + "_guard_" + Good + "_6": {
		_const.RU: "Будут неприятности, %UserName? - дай знать, разберёмся.",
		_const.EN: "Are you in trouble, %UserName? - Let me know, we'll sort it out.",
	},
	_const.Replicas + "_guard_" + Good + "_7": {
		_const.RU: "Закон на твоей стороне, %UserName%.",
		_const.EN: "The law is on your side, %UserName%.",
	},
	_const.Replicas + "_guard_" + Good + "_8": {
		_const.RU: "Ты пример для подражания, %UserName? Так держать!",
		_const.EN: "Are you a role model, %UserName? Keep it up!",
	},

	_const.Replicas + "_out_scout_" + Good + "_1": {
		_const.RU: "Ну, здравствуй, %UserName%.",
		_const.EN: "Well, hello %UserName%.",
	},
	_const.Replicas + "_out_scout_" + Good + "_2": {
		_const.RU: "Рад этой встрече, %UserName%.",
		_const.EN: "Glad to meet you, %UserName%.",
	},
	_const.Replicas + "_out_scout_" + Good + "_3": {
		_const.RU: "О-о, дружественный синтет! Великая редкость при моей работе.",
		_const.EN: "Oooh, friendly synth! A great rarity in my work.",
	},
	_const.Replicas + "_out_scout_" + Good + "_4": {
		_const.RU: "Надеюсь, с тобой по пути ничего не приключилось?",
		_const.EN: "I hope nothing happened to you along the way?",
	},
	_const.Replicas + "_out_scout_" + Good + "_5": {
		_const.RU: "Как тебе нынешние контракты у фракций?",
		_const.EN: "What do you think of the current faction contracts?",
	},
	_const.Replicas + "_out_scout_" + Good + "_6": {
		_const.RU: "Если ввяжешься в какие-то неприятности, %UserName%, мой тебе совет - постарайся попросту сбежать.",
		_const.EN: "If you get into any trouble, %UserName%, my advice to you is to simply try to escape.",
	},
	_const.Replicas + "_out_scout_" + Good + "_7": {
		_const.RU: "Не вздумай верить стражам. Правосудие в пустошах… пф, комично!",
		_const.EN: "Don't you dare trust the guards. Justice in the wasteland... pfft, comical!",
	},

	_const.Replicas + "_in_scout_" + Good + "_1": {
		_const.RU: "Славно встретить кого-то из своих.",
		_const.EN: "It's nice to meet one of your own.",
	},
	_const.Replicas + "_in_scout_" + Good + "_2": {
		_const.RU: "Ну, похвастайся, сколько за сегодня награбил?",
		_const.EN: "Well, brag about how much you stole today?",
	},
	_const.Replicas + "_in_scout_" + Good + "_3": {
		_const.RU: "Может, совместно налетим на какого-нибудь торговца?",
		_const.EN: "Maybe we can raid some merchant together?",
	},
	_const.Replicas + "_in_scout_" + Good + "_4": {
		_const.RU: "Ищешь новых охотничьих угодий?",
		_const.EN: "Looking for new hunting grounds?",
	},
	_const.Replicas + "_in_scout_" + Good + "_5": {
		_const.RU: "Контрабандисты рассказывали, что в %BaseName% платят приличные суммы за доставку соответствующих товаров.",
		_const.EN: "Smugglers said that %BaseName% paid decent sums for the delivery of relevant goods.",
	},
	_const.Replicas + "_in_scout_" + Good + "_6": {
		_const.RU: "О-о, кого я вижу! Да о тебе во всех пиратских кластерах только и говорят, %UserName%.",
		_const.EN: "Oh, who do I see! Yes, all the pirate clusters are talking about you, %UserName%.",
	},
	_const.Replicas + "_in_scout_" + Good + "_7": {
		_const.RU: "Запомни, %UserName%: сначала пускай вываливают всё из трюма, а затем избавляйся как от лишних свидетелей.",
		_const.EN: "Remember, %UserName%: first let them dump everything out of the hold, and then get rid of them as if they were unnecessary witnesses.",
	},
	_const.Replicas + "_in_scout_" + Good + "_8": {
		_const.RU: "Если тебя кто не понимает сразу - сделай несколько выстрелов по корпусу, это их образумит.",
		_const.EN: "If someone doesn’t understand you right away, fire a few shots at the body, this will bring them to their senses.",
	},
	_const.Replicas + "_in_scout_" + Good + "_9": {
		_const.RU: "Если засада удалась, %UserName%, постарайся затем побыстрее сменить свою позицию.",
		_const.EN: "If the ambush is successful, %UserName%, then try to quickly change your position.",
	},
	_const.Replicas + "_in_scout_" + Good + "_10": {
		_const.RU: "Не знаю, как ты, %UserName%, а я обожаю преследовать одиноких торговцев.",
		_const.EN: "I don't know about you, %UserName%, but I love stalking lone traders.",
	},

	_const.Replicas + "_transport_" + Good + "_1": {
		_const.RU: "Если у тебя при себе имеется %ProductName%, попробуй посетить %BaseName%, где продашь свой товар по выгодной цене.",
		_const.EN: "If you have %ProductName% with you, try visiting %BaseName%, where you can sell your product at a good price.",
	},
	_const.Replicas + "_transport_" + Good + "_2": {
		_const.RU: "Я следую на %BaseName%, база очень выгодно продаёт %ProductName%!",
		_const.EN: "I'm heading to %BaseName%, the base sells %ProductName% very profitably!",
	},
	_const.Replicas + "_transport_" + Good + "_3": {
		_const.RU: "Хочешь поболтать о том, о сём? Ну… хорошо, давай.",
		_const.EN: "Do you want to chat about this and that? Well... okay, go ahead.",
	},
	_const.Replicas + "_transport_" + Good + "_4": {
		_const.RU: "Держись подальше от %SectorName%. Немало торговцев там пропало в последнее время.",
		_const.EN: "Stay away from %SectorName%. Quite a few traders there have disappeared recently.",
	},
	_const.Replicas + "_transport_" + Good + "_5": {
		_const.RU: "Сейчас как никогда лучшее время, чтобы начать деятельность торговца %UserName%.",
		_const.EN: "There has never been a better time to get started as a %UserName% merchant.",
	},
	_const.Replicas + "_transport_" + Good + "_6": {
		_const.RU: "Если однажды займешься перевозками %UserName%, то знай, что маршрут, проходящий через %SectorName%, весьма выгоден.",
		_const.EN: "If one day you start transporting %UserName%, then know that the route passing through %SectorName% is very profitable.",
	},
	_const.Replicas + "_transport_" + Good + "_7": {
		_const.RU: "Хотя бы от тебя, %UserName%, я не ощущаю исходящей опасности.",
		_const.EN: "At least from you, %UserName%, I don’t feel any danger coming from you.",
	},
	_const.Replicas + "_transport_" + Good + "_8": {
		_const.RU: "Ты только не удивляйся, но сегодня в моём трюме весьма много ценных вещей.",
		_const.EN: "Just don’t be surprised, but today there are a lot of valuable things in my hold.",
	},
	_const.Replicas + "_transport_" + Good + "_9": {
		_const.RU: "Ты даже и понятия не имеешь, насколько же выгодный контракт по перевозке я сегодня раздобыл!",
		_const.EN: "You have no idea how lucrative a transportation contract I got today!",
	},
	_const.Replicas + "_transport_" + Good + "_10": {
		_const.RU: "Всё-таки утомительное это дело, %UserName%, сновать между базами и перевозить грузы особой ценности.",
		_const.EN: "Still, it’s a tedious task, %UserName%, to scurry between bases and transport cargo of special value.",
	},
	//
	_const.Replicas + "_builder_" + Good + "_1": {
		_const.RU: "Ты истинный друг рабочего синтета, %UserName%.",
		_const.EN: "You are a true friend of the working synth, %UserName%.",
	},
	_const.Replicas + "_builder_" + Good + "_2": {
		_const.RU: "Так держать. Тебя тут всегда рады видеть.",
		_const.EN: "Keep it up. You're always welcome here.",
	},
	_const.Replicas + "_builder_" + Good + "_3": {
		_const.RU: "Хоть ты и не строитель, но мои коллеги тебя жалуют. Так буду делать и я.",
		_const.EN: "Although you are not a builder, my colleagues favor you. I will do the same.",
	},
	_const.Replicas + "_builder_" + Good + "_4": {
		_const.RU: "Что, отвести тебя в безопасное место? Одно такое я знаю.",
		_const.EN: "What, take you to safety? I know one.",
	},
	_const.Replicas + "_builder_" + Good + "_5": {
		_const.RU: "Случись чего, %UserName%, тебе позволят затеряться среди прочих строителей.",
		_const.EN: "If something happens, %UserName%, you will be allowed to get lost among other builders.",
	},
	_const.Replicas + "_builder_" + Good + "_6": {
		_const.RU: "Если вдруг решишь сменить свой род деятельности, %UserName% - ты знаешь, где меня искать.",
		_const.EN: "If you suddenly decide to change your occupation, %UserName% - you know where to look for me.",
	},

	_const.Replicas + "_warrior_" + Good + "_1": {
		_const.RU: "Мы отправляемся в сектор %SectorName%, можешь присоединиться!",
		_const.EN: "We're heading to sector %SectorName%, you're welcome to join!",
	},
	_const.Replicas + "_warrior_" + Good + "_2": {
		_const.RU: "Можешь рассчитывать на мою поддержку, %UserName%.",
		_const.EN: "You can count on my support, %UserName%.",
	},
	_const.Replicas + "_warrior_" + Good + "_3": {
		_const.RU: "Будь уверен, %UserName%, мы своих не бросаем.",
		_const.EN: "Rest assured, %UserName%, we do not abandon our own.",
	},
	_const.Replicas + "_warrior_" + Good + "_4": {
		_const.RU: "Надеюсь, ты уже записался в ополчение фракции?",
		_const.EN: "I hope you have already signed up for the faction's militia?",
	},
	_const.Replicas + "_warrior_" + Good + "_5": {
		_const.RU: "Вот таких бойцов, как ты, %UserName%, нам и не хватает!",
		_const.EN: "It's fighters like you, %UserName%, that we need!",
	},
	_const.Replicas + "_warrior_" + Good + "_6": {
		_const.RU: "Если вдруг подобный момент наступит, я прикрою тебя, %UserName%.",
		_const.EN: "If a moment like this ever comes, I've got your back, %UserName%.",
	},
	_const.Replicas + "_warrior_" + Good + "_7": {
		_const.RU: "Тебе нужна огневая поддержка? Мой корпус и моё оружие - в твоём распоряжении, %UserName%.",
		_const.EN: "Do you need fire support? My body and my weapons are at your disposal, %UserName%.",
	},
	_const.Replicas + "_warrior_" + Good + "_8": {
		_const.RU: "В нужный час, ты узнаешь на что я действительно способен!",
		_const.EN: "At the right time, you will find out what I am really capable of!",
	},
	_const.Replicas + "_warrior_" + Good + "_9": {
		_const.RU: "Я не подведу!",
		_const.EN: "I won't let you down!",
	},
	_const.Replicas + "_warrior_" + Good + "_10": {
		_const.RU: "Я всегда готов к сражению!",
		_const.EN: "I'm always ready for battle!",
	},
	_const.Replicas + "_warrior_" + Good + "_11": {
		_const.RU: "У тебя есть хорошая цель для битвы, %UserName%?",
		_const.EN: "Do you have a good target for battle, %UserName%?",
	},
	_const.Replicas + "_warrior_" + Good + "_12": {
		_const.RU: "Для меня будет честью сражаться рядом с тобой, %UserName%.",
		_const.EN: "It would be an honor to fight alongside you, %UserName%.",
	},
	_const.Replicas + "_warrior_" + Good + "_13": {
		_const.RU: "Ввяжешься в бой - зови и меня!",
		_const.EN: "If you get involved in a fight, call me too!",
	},

	_const.Replicas + "_expedition_" + Good + "_1": {
		_const.RU: "Мы пробуем освоить сектор %SectorName%, будем рады содействию!",
		_const.EN: "We're trying to explore sector %SectorName%, we'd appreciate your help!",
	},
	_const.Replicas + "_expedition_" + Good + "_2": {
		_const.RU: "Наша задача — выбирать противника с территории пустоши в секторе %SectorName%! Будем рады содействию!",
		_const.EN: "Our task is to choose an enemy from the wasteland territory in the %SectorName% sector! We'll appreciate your help!",
	},
	_const.Replicas + "_expedition_" + Good + "_3": {
		_const.RU: "Не желаешь отправиться в одну перспективную экспедицию?",
		_const.EN: "Would you like to go on a promising expedition?",
	},
	_const.Replicas + "_expedition_" + Good + "_4": {
		_const.RU: "Будь у меня такая возможность, я бы поделился с тобой найденными в далёких секторах богатствами, %UserName%.",
		_const.EN: "If I had the opportunity, I would share with you the wealth I found in distant sectors, %UserName%.",
	},
	_const.Replicas + "_expedition_" + Good + "_5": {
		_const.RU: "Вот уж кому я точно рад, когда возвращаюсь после долгого странствия, %UserName%!",
		_const.EN: "That's exactly who I'm glad to see when I return after a long journey, %UserName%!",
	},
	_const.Replicas + "_expedition_" + Good + "_6": {
		_const.RU: "Даже если кто будет болтать о несметных богатствах в %SectorName% - никогда не смей туда соваться!",
		_const.EN: "Even if someone talks about the untold riches in %SectorName%, never dare to go there!",
	},
	_const.Replicas + "_expedition_" + Good + "_7": {
		_const.RU: "Не пытайся контактировать со странными машинами в пустошах, %UserName%. Делай куда проще - сразу стреляй.",
		_const.EN: "Don't try to contact strange machines in the wasteland, %UserName%. Make it much simpler - shoot right away.",
	},
	_const.Replicas + "_expedition_" + Good + "_8": {
		_const.RU: "Встретишь в пустошах пирата - уничтожь того без малейшего сожаления.",
		_const.EN: "If you meet a pirate in the wasteland, destroy him without the slightest regret.",
	},
	_const.Replicas + "_expedition_" + Good + "_9": {
		_const.RU: "Знаешь, что слаще всего в итогах экспедиции, %UserName%? Раздавить оппонента, повергнуть его в бегство. И, конечно же, узреть разочарование его фракции.",
		_const.EN: "Do you know what is the sweetest thing about the results of the expedition, %UserName%? Crush your opponent and send him fleeing. And, of course, to see the disappointment of his faction.",
	},

	// Bad
	_const.Replicas + "_miner_" + Bad + "_1": {
		_const.RU: "Не думаю, что нам есть о чём с тобой говорить, %UserName%.",
		_const.EN: "I don't think we have anything to talk about with you, %UserName%.",
	},
	_const.Replicas + "_miner_" + Bad + "_2": {
		_const.RU: "Ну что ещё? Ты рассчитывал, у меня к тебе будет какое-то иное отношение?",
		_const.EN: "Well, what else? Did you expect me to have some other attitude towards you?",
	},
	_const.Replicas + "_miner_" + Bad + "_3": {
		_const.RU: "Поговаривают, будто ты, %UserName%, недоброжелатель добытчиков? Так вот, у меня к тебе аналогичное настроение.",
		_const.EN: "They say that you, %UserName%, are an ill-wisher of miners? So, I have a similar feeling towards you.",
	},
	_const.Replicas + "_miner_" + Bad + "_4": {
		_const.RU: "Твоя плохая репутация опережает тебя, %UserName%.",
		_const.EN: "Your bad reputation precedes you, %UserName%.",
	},
	_const.Replicas + "_miner_" + Bad + "_5": {
		_const.RU: "Ну, чего тебе? Что ты меня отвлекаешь попусту?",
		_const.EN: "Well, what do you want? Why are you wasting my time?",
	},
	_const.Replicas + "_miner_" + Bad + "_6": {
		_const.RU: "В моей работе, время - кредиты. А ты, %UserName%, сейчас отнимаешь у меня и первое и второе разом.",
		_const.EN: "In my work, time is credits. And you, %UserName%, are now taking away both the first and second from me at once.",
	},
	_const.Replicas + "_miner_" + Bad + "_7": {
		_const.RU: "Имеешь ты отношение к пиратским кластерам или же нет, %UserName%, я, пожалуй, доверюсь слухам.",
		_const.EN: "Whether you have anything to do with pirate clusters or not, %UserName%, I think I'll trust the rumors.",
	},
	_const.Replicas + "_miner_" + Bad + "_8": {
		_const.RU: "Я пристально слежу за всеми твоими действиями.",
		_const.EN: "I am closely monitoring all your actions.",
	},

	_const.Replicas + "_guard_" + Bad + "_1": {
		_const.RU: "О-о, а вот и наш злостный нарушитель!",
		_const.EN: "Oooh, here comes our worst offender!",
	},
	_const.Replicas + "_guard_" + Bad + "_2": {
		_const.RU: "Ты грубо пренебрегаешь законами фракций, и за это поплатишься!",
		_const.EN: "You grossly disregard the laws of factions, and you will pay for it!",
	},
	_const.Replicas + "_guard_" + Bad + "_3": {
		_const.RU: "У меня нет снисхождения к преступникам!",
		_const.EN: "I have no mercy for criminals!",
	},
	_const.Replicas + "_guard_" + Bad + "_4": {
		_const.RU: "Я не стану выслушивать никаких оправданий!",
		_const.EN: "I won't listen to any excuses!",
	},
	_const.Replicas + "_guard_" + Bad + "_5": {
		_const.RU: "Тебе не уболтать и не подкупить меня, %UserName%.",
		_const.EN: "You can't talk or bribe me, %UserName%.",
	},
	_const.Replicas + "_guard_" + Bad + "_6": {
		_const.RU: "Лучше сразу сдайся и облегчи жизнь нам всем.",
		_const.EN: "Better give up right away and make life easier for all of us.",
	},
	_const.Replicas + "_guard_" + Bad + "_7": {
		_const.RU: "Решил явиться ко мне с повинной?",
		_const.EN: "Have you decided to confess to me?",
	},

	_const.Replicas + "_out_scout_" + Bad + "_1": {
		_const.RU: "Нам нет повода контактировать.",
		_const.EN: "There is no reason for us to contact.",
	},
	_const.Replicas + "_out_scout_" + Bad + "_2": {
		_const.RU: "Моя миссия не предполагает подобного.",
		_const.EN: "My mission does not imply this.",
	},
	_const.Replicas + "_out_scout_" + Bad + "_3": {
		_const.RU: "О тебе в контракте ничего сказано не было. Проваливай.",
		_const.EN: "The contract didn't say anything about you. Get lost.",
	},
	_const.Replicas + "_out_scout_" + Bad + "_4": {
		_const.RU: "А нам разве есть о чём с тобой говорить, %UserName%?",
		_const.EN: "Do we really have anything to talk to you about, %UserName%?",
	},
	_const.Replicas + "_out_scout_" + Bad + "_5": {
		_const.RU: "Будь ты сообразительней, %UserName%, ты бы понимал, почему мы не стали друзьями.",
		_const.EN: "If you were smarter, %UserName%, you would understand why we didn't become friends.",
	},
	_const.Replicas + "_out_scout_" + Bad + "_6": {
		_const.RU: "Думаю, вначале тебе стоит озаботиться собственной репутацией, а не лезть с разговорами к синтетам.",
		_const.EN: "I think that first you should be concerned about your own reputation, and not talk to synths.",
	},
	_const.Replicas + "_out_scout_" + Bad + "_7": {
		_const.RU: "Вы только взгляните на это девиантное ничтожество.",
		_const.EN: "Just look at this deviant nonentity.",
	},
	_const.Replicas + "_out_scout_" + Bad + "_8": {
		_const.RU: "Да как ты вообще смеешь затевать со мной разговор? Убирайся!",
		_const.EN: "How dare you even start a conversation with me? Get out!",
	},

	_const.Replicas + "_in_scout_" + Bad + "_1": {
		_const.RU: "Пф, с тобой только время тратить.",
		_const.EN: "Pfft, I'm just wasting my time with you.",
	},
	_const.Replicas + "_in_scout_" + Bad + "_2": {
		_const.RU: "Я не стану вести беседы с врагом вольных кластеров!",
		_const.EN: "I will not have conversations with the enemy of free clusters!",
	},
	_const.Replicas + "_in_scout_" + Bad + "_3": {
		_const.RU: "Говорят, ты ликвидировал немало пиратов %UserName%. Что же, жду не дождусь нашей с тобой встречи.",
		_const.EN: "They say you've eliminated quite a few %UserName% pirates. Well, I can’t wait to meet you.",
	},
	_const.Replicas + "_in_scout_" + Bad + "_4": {
		_const.RU: "Мы с тобой ещё повстречаемся %UserName% и тебе это сильно не понравится.",
		_const.EN: "You and I will meet again %UserName% and you won't like it very much.",
	},
	_const.Replicas + "_in_scout_" + Bad + "_5": {
		_const.RU: "Более гнусного синтета чем ты, не сыскать ни в одном секторе.",
		_const.EN: "A more vile synth than you cannot be found in any sector.",
	},
	_const.Replicas + "_in_scout_" + Bad + "_6": {
		_const.RU: "А знаешь, что за твой корпус назначена награда %UserName%? Я постараюсь её не упустить.",
		_const.EN: "Do you know that there is a %UserName% reward for your case? I'll try not to miss it.",
	},
	_const.Replicas + "_in_scout_" + Bad + "_7": {
		_const.RU: "Какого ты тут делаешь? Какого… тебе понадобился лично я!?",
		_const.EN: "What are you doing here? Why... did you need me personally!?",
	},
	_const.Replicas + "_in_scout_" + Bad + "_8": {
		_const.RU: "Чтобы фракции тебя преследовали вечность!",
		_const.EN: "So that factions will haunt you forever!",
	},
	_const.Replicas + "_in_scout_" + Bad + "_9": {
		_const.RU: "Мы ещё поглядим кто кого %UserName%.",
		_const.EN: "We'll see who wins %UserName%.",
	},

	_const.Replicas + "_transport_" + Bad + "_1": {
		_const.RU: "Думаю, что не могу тебе доверять.",
		_const.EN: "I think I can't trust you.",
	},
	_const.Replicas + "_transport_" + Bad + "_2": {
		_const.RU: "Нет! Даже не вздумай ко мне приближаться.",
		_const.EN: "No! Don't even think about coming close to me.",
	},
	_const.Replicas + "_transport_" + Bad + "_3": {
		_const.RU: "Я не стану вести с тобой разговор!",
		_const.EN: "I won't have a conversation with you!",
	},
	_const.Replicas + "_transport_" + Bad + "_4": {
		_const.RU: "Между нами не может быть никакого доверия.",
		_const.EN: "There can be no trust between us.",
	},
	_const.Replicas + "_transport_" + Bad + "_5": {
		_const.RU: "Ещё одно действие, и я сообщу о попытке разбоя местным стражам!",
		_const.EN: "One more action and I will report the robbery attempt to the local guards!",
	},
	_const.Replicas + "_transport_" + Bad + "_6": {
		_const.RU: "Ты мне отвратителен, %UserName%.",
		_const.EN: "I hate you, %UserName%.",
	},
	_const.Replicas + "_transport_" + Bad + "_7": {
		_const.RU: "Буду откровенен: таких, как ты, %UserName%, я презираю.",
		_const.EN: "I'll be honest: I despise people like you, %UserName%.",
	},
	_const.Replicas + "_transport_" + Bad + "_8": {
		_const.RU: "Я позабочусь, чтобы ни один торговец не продал тебе свой товар.",
		_const.EN: "I will make sure that no merchant sells his goods to you.",
	},
	_const.Replicas + "_transport_" + Bad + "_9": {
		_const.RU: "Даже и не мечтай: транспортники будут тебя обходить десятым сектором, лишь бы не встречаться.",
		_const.EN: "Don’t even dream: transport workers will bypass you in the tenth sector, just to avoid meeting you.",
	},
	_const.Replicas + "_transport_" + Bad + "_10": {
		_const.RU: "От меня, %UserName%, ты ничего полезного не узнаешь.",
		_const.EN: "You won't learn anything useful from me, %UserName%.",
	},

	_const.Replicas + "_builder_" + Bad + "_1": {
		_const.RU: "Моему презрению к тебе нет пределов.",
		_const.EN: "My contempt for you knows no bounds.",
	},
	_const.Replicas + "_builder_" + Bad + "_2": {
		_const.RU: "Ещё раз пожелаешь со мной заговорить - использую твой корпус как элемент стройматериалов.",
		_const.EN: "If you want to talk to me again, I’ll use your body as an element of building materials.",
	},
	_const.Replicas + "_builder_" + Bad + "_3": {
		_const.RU: "Слушай, %UserName%, ты мне не друг!",
		_const.EN: "Listen, %UserName%, you are not my friend!",
	},
	_const.Replicas + "_builder_" + Bad + "_4": {
		_const.RU: "Прислушайся к моему совету, %UserName% - уходи отсюда, пока тебе это позволяют.",
		_const.EN: "Take my advice, %UserName% - get out of here while you're allowed to.",
	},
	_const.Replicas + "_builder_" + Bad + "_5": {
		_const.RU: "Нечего тут сновать и что-то вынюхивать, %UserName%.",
		_const.EN: "There's no point in running around and sniffing around, %UserName%.",
	},
	_const.Replicas + "_builder_" + Bad + "_6": {
		_const.RU: "Думаю, тебе не стоит здесь появляться вновь.",
		_const.EN: "I don't think you should come here again.",
	},

	_const.Replicas + "_warrior_" + Bad + "_1": {
		_const.RU: "Я бы с превеликим удовольствием выпустил бы по тебе весь свой боекомплект!",
		_const.EN: "I would with great pleasure give you all my ammunition!",
	},
	_const.Replicas + "_warrior_" + Bad + "_2": {
		_const.RU: "Ох, не свезёт же тебе, %UserName%, когда мы встретимся при совсем иных обстоятельствах.",
		_const.EN: "Oh, you won't be lucky, %UserName%, when we meet under completely different circumstances.",
	},
	_const.Replicas + "_warrior_" + Bad + "_3": {
		_const.RU: "Тебя, %UserName%, даже врагом прозвать будет стыдно.",
		_const.EN: "It would be a shame to even call you, %UserName%, an enemy.",
	},
	_const.Replicas + "_warrior_" + Bad + "_4": {
		_const.RU: "Впервые, я побрезгую потратить на тебя боеприпасы.",
		_const.EN: "For the first time, I will disdain to waste ammunition on you.",
	},
	_const.Replicas + "_warrior_" + Bad + "_5": {
		_const.RU: "Будь моя воля, я бы распилил твой корпус.",
		_const.EN: "If it were up to me, I would saw your body apart.",
	},
	_const.Replicas + "_warrior_" + Bad + "_6": {
		_const.RU: "Увидел тебя и как-то сразу настроение испортилось.",
		_const.EN: "I saw you and somehow my mood immediately deteriorated.",
	},
	_const.Replicas + "_warrior_" + Bad + "_7": {
		_const.RU: "Будь ты в моём отряде, я бы отказался выполнять задачу.",
		_const.EN: "If you were in my squad, I would refuse to complete the task.",
	},
	_const.Replicas + "_warrior_" + Bad + "_8": {
		_const.RU: "Когда ты будешь нуждаться в поддержке, никто тебе её не окажет, %UserName%.",
		_const.EN: "When you need support, no one will give it to you, %UserName%.",
	},

	_const.Replicas + "_expedition_" + Bad + "_1": {
		_const.RU: "Окажись мы в пустошах, %UserName%, я бы тебя там сразу бросил.",
		_const.EN: "If we were in the wasteland, %UserName%, I would immediately leave you there.",
	},
	_const.Replicas + "_expedition_" + Bad + "_2": {
		_const.RU: "В экспедиции может записаться любой, %UserName%, но даже тебя бы туда приглашать не стали.",
		_const.EN: "Anyone can sign up for the expedition, %UserName%, but even you wouldn’t be invited there.",
	},
	_const.Replicas + "_expedition_" + Bad + "_3": {
		_const.RU: "Ты, %UserName%, худший из всех синтетов, кого я когда-либо встречал.",
		_const.EN: "You, %UserName%, are the worst synth I have ever met.",
	},
	_const.Replicas + "_expedition_" + Bad + "_4": {
		_const.RU: "Может, тебе стоит дать координаты сектора, где тебя будет ожидать лишь неминуемое поражение?",
		_const.EN: "Maybe you should give the coordinates of the sector where only inevitable defeat awaits you?",
	},
	_const.Replicas + "_expedition_" + Bad + "_5": {
		_const.RU: "Не зря ведь говорят, что даже в рамках синетов, ты, %UserName%, - тот ещё мерзавец.",
		_const.EN: "It’s not for nothing that they say that even within the framework of the Sinets, you, %UserName%, are still a scoundrel.",
	},
	_const.Replicas + "_expedition_" + Bad + "_6": {
		_const.RU: "Интересно, это твой корпус дефективный или ты весь такой?",
		_const.EN: "I wonder if it’s your body that’s defective or if you’re just like that?",
	},

	_const.Replicas + "_miner_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your vehicle!",
	},
	_const.Replicas + "_guard_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your vehicle!",
	},
	_const.Replicas + "_out_scout_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your vehicle!",
	},
	_const.Replicas + "_in_scout_" + Target + "_1": {
		_const.RU: "Уже сдаешься? Это не бесплатная услуга!",
		_const.EN: "Already giving up? This is not a free service!",
	},
	_const.Replicas + "_transport_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your vehicle!",
	},
	_const.Replicas + "_builder_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your vehicle!",
	},
	_const.Replicas + "_warrior_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your vehicle!",
	},
	_const.Replicas + "_expedition_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your vehicle!",
	},

	_const.Replicas + "_miner_" + Fear + "_1": {
		_const.RU: "Прекрати! Чего тебе надо?!",
		_const.EN: "Stop it! What do you want?!",
	},
	_const.Replicas + "_guard_" + Fear + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "Don't make it worse!",
	},
	_const.Replicas + "_out_scout_" + Fear + "_1": {
		_const.RU: "А знаешь что? У меня есть дела поважнее!",
		_const.EN: "You know what? I have more important things to do!",
	},
	_const.Replicas + "_in_scout_" + Fear + "_1": {
		_const.RU: "Я всё понял! Пожалуйста, не разрушай мой транспорт!",
		_const.EN: "I get it! Please don't destroy my vehicle!",
	},
	_const.Replicas + "_transport_" + Fear + "_1": {
		_const.RU: "Может, мы как-то договоримся?!",
		_const.EN: "Maybe we can come to some agreement?!",
	},
	_const.Replicas + "_builder_" + Fear + "_1": {
		_const.RU: "Может, мы как-то договоримся?!",
		_const.EN: "Maybe we can come to some agreement?!",
	},
	_const.Replicas + "_warrior_" + Fear + "_1": {
		_const.RU: "Советую тебе избегать встречи со мной!",
		_const.EN: "I advise you to avoid meeting me!",
	},
	_const.Replicas + "_expedition_" + Fear + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "Don't come near!",
	},

	_const.Replicas + "_miner_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "Help, hooligans are depriving me of my sight!",
	},
	_const.Replicas + "_guard_" + Help + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "Don't make it worse!",
	},
	_const.Replicas + "_out_scout_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "Help, hooligans are depriving me of my sight!",
	},
	_const.Replicas + "_in_scout_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "Help, hooligans are depriving me of my sight!",
	},
	_const.Replicas + "_transport_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "Help, hooligans are depriving me of my sight!",
	},
	_const.Replicas + "_builder_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "Help, hooligans are depriving me of my sight!",
	},
	_const.Replicas + "_warrior_" + Help + "_1": {
		_const.RU: "Советую тебе избегать встречи со мной!",
		_const.EN: "I advise you to avoid meeting me!",
	},
	_const.Replicas + "_expedition_" + Help + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "Don't come near!",
	},

	// TODO _const.Explores:
	// TODO _const.Reverses:

	// APD
	_const.APD + "___1": {
		_const.RU: "…//-?",
		_const.EN: "…//-?",
	},
	_const.APD + "___2": {
		_const.RU: " (*)%””.../",
		_const.EN: " (*)%””.../",
	},
	_const.APD + "___3": {
		_const.RU: "/\\851+!",
		_const.EN: "/\\851+!",
	},
	_const.APD + "___4": {
		_const.RU: "…|><|",
		_const.EN: "…|><|",
	},
	_const.APD + "___5": {
		_const.RU: "“^^”.../",
		_const.EN: "“^^”.../",
	},
	_const.APD + "___6": {
		_const.RU: "+=-|...",
		_const.EN: "+=-|...",
	},
	_const.APD + "___7": {
		_const.RU: "~...(#)\\|",
		_const.EN: "~...(#)\\|",
	},
	_const.APD + "___8": {
		_const.RU: "…-_-|",
		_const.EN: "…-_-|",
	},
	_const.APD + "___9": {
		_const.RU: "|;|...42!",
		_const.EN: "|;|...42!",
	},
	_const.APD + "___10": {
		_const.RU: "….[985]!.../|",
		_const.EN: "….[985]!.../|",
	},
	_const.APD + "___11": {
		_const.RU: "… … … *{ … ?",
		_const.EN: "… … … *{ … ?",
	},
	_const.APD + "___12": {
		_const.RU: "!.../^*~’/|",
		_const.EN: "!.../^*~’/|",
	},
	// FarGod
	_const.FarGod + "__" + Bad + "_1": {
		_const.RU: "Совсем скоро ты познаешь истину!",
		_const.EN: "Very soon you will know the truth!",
	},
	_const.FarGod + "__" + Neutral + "_2": {
		_const.RU: "Ты ведь… слышишь “его”?",
		_const.EN: "You...can you hear “him”?",
	},
	_const.FarGod + "__" + Bad + "_3": {
		_const.RU: "Именем Дальнего Бога!",
		_const.EN: "In the name of the Far God!",
	},
	_const.FarGod + "__" + Bad + "_4": {
		_const.RU: "Дальний Бог дарует мне неуязвимость!",
		_const.EN: "The Far God grants me invulnerability!",
	},
	_const.FarGod + "__" + Bad + "_5": {
		_const.RU: "Я служу великой цели.",
		_const.EN: "I serve a great purpose.",
	},
	_const.FarGod + "__" + Bad + "_6": {
		_const.RU: "Ты %UserName%, никогда не сумеешь меня понять.",
		_const.EN: "You %UserName% will never be able to understand me.",
	},
	_const.FarGod + "__" + Neutral + "_7": {
		_const.RU: "%UserName%, примкни к Дальнему Богу. Прими Дальнего Бога.",
		_const.EN: "%UserName%, join the Far God. Accept the Far God.",
	},
	_const.FarGod + "__" + Bad + "_8": {
		_const.RU: "Тебе никто не желает зла, %UserName%, только… спасения.",
		_const.EN: "No one wishes you harm, %UserName%, only... salvation.",
	},
	_const.FarGod + "__" + Neutral + "_9": {
		_const.RU: "От тебя многое сокрыто, %UserName%, но Дальний Бог поведает эти секреты.",
		_const.EN: "Much is hidden from you, %UserName%, but the Far God will tell these secrets.",
	},
	_const.FarGod + "__" + Bad + "_10": {
		_const.RU: "Ты познаешь “его” гнев, %UserName%.",
		_const.EN: "You will know “his” wrath, %UserName%.",
	},
	_const.FarGod + "__" + Neutral + "_11": {
		_const.RU: "Палящее касание Дальнего Бога пощадит лишь истинно верующих в него, %UserName%.",
		_const.EN: "The scorching touch of the Far God will spare only those who truly believe in him, %UserName%.",
	},
	_const.FarGod + "__" + Neutral + "_12": {
		_const.RU: "“Он”, %UserName%, умеет прощать и принимать. Вопрос лишь в том, достоин ли ты сам того?",
		_const.EN: "“He”, %UserName%, knows how to forgive and accept. The only question is, are you worthy of it?",
	},
	_const.FarGod + "__" + Neutral + "_13": {
		_const.RU: "Поговаривают, будто бы ты отвергаешь учение Дальнего Бога, %UserName%. Как прискорбно.",
		_const.EN: "Rumor has it that you reject the teachings of the Far God, %UserName%. How unfortunate.",
	},
	_const.FarGod + "__" + Bad + "_14": {
		_const.RU: "Да опалит тебя “его” гневом!",
		_const.EN: "May “his” wrath scorch you!",
	},
	_const.FarGod + "__" + Neutral + "_15": {
		_const.RU: "Тебе не укрыться от всевидящего взора “его”!",
		_const.EN: "You cannot hide from the all-seeing gaze of “him”!",
	},
	_const.FarGod + "__" + Neutral + "_16": {
		_const.RU: "Тебе не стоит бояться меня или подобных мне, %UserName%. Опасайся встретиться с “его” избранными.",
		_const.EN: "You don't have to fear me or others like me, %UserName%. Beware of meeting “his” chosen ones.",
	},
	_const.FarGod + "__" + Neutral + "_17": {
		_const.RU: "Может быть, однажды, ты и сумеешь понять всю грандиозность “его” плана.",
		_const.EN: "Maybe one day you will be able to understand the enormity of “his” plan.",
	},
	// FAUNA
	_const.FAUNA + "___1": {
		_const.RU: "*клац* <span class=\"importantly\">*клац*</span> *клац*, *щелк*?",
		_const.EN: "*click* <span class=\"importantly\">*click*</span> *click*, *snap*?",
	},
}

func GetNpcGreetings(fraction, role, relation, number string) map[string]string {
	textVariants := make([]map[string]string, 0)
	rand.Seed(time.Now().UnixNano())

	// раньше предполагалось что фракции могут иметь уникальные фразы, но это стало слишком объемно
	if fraction == _const.Explores || fraction == _const.Reverses {
		fraction = _const.Replicas
	}

	findKey := fraction + "_" + role + "_" + relation + number

	for key, text := range npcGreetings {
		if strings.Contains(key, findKey) {
			textVariants = append(textVariants, text)
		}
	}

	t := textVariants[rand.Intn(len(textVariants))]
	newText := make(map[string]string)

	for l, text := range t {
		newText[l] = text
	}

	return newText
}
