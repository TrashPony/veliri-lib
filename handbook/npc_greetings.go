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
		_const.RU: "%UserName%, если у тебя есть соответствующая руда, попробуй наведаться на %BaseName%. Сейчас, там можно выгодно реализовать %ProductName%.",
		_const.EN: "%UserName%, if you have the appropriate ore, try visiting %BaseName%. Now, you can profitably sell %ProductName% there.",
	},
	_const.Replicas + "_miner_" + Good + "_5": {
		_const.RU: "Тоже решил заняться добычей руды, а по итогу не особо то и идёт дело? Попробуй-ка посетить %SectorName%.",
		_const.EN: "You also decided to start mining ore, but in the end things aren’t going well? Try visiting %SectorName%.",
	},
	_const.Replicas + "_miner_" + Good + "_6": {
		_const.RU: "Сегодня мне повезло как никогда ранее! Столько %ProductName% я ещё никогда за всё своё существование не добывал.",
		_const.EN: "Today I am luckier than ever before! I have never mined so much %ProductName% in my entire existence.",
	},
	_const.Replicas + "_miner_" + Good + "_7": {
		_const.RU: "Только между нами %UserName%: если рассчитываешь выгодно переработать %ProductName%, советую на время позабыть о %BaseName%. Ты ещё и останешься должным тамошнему диспетчеру.",
		_const.EN: "Just between us %UserName%: if you expect to profitably recycle %ProductName%, I advise you to forget about %BaseName% for a while. You will also remain indebted to the dispatcher there.",
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
		_const.RU: "Как тебе нынешние контракты у фракций? Бывал в %SectorName%? Я недавно выбил оттуда один пиратский кластер.",
		_const.EN: "What do you think of the current faction contracts? Have you been to %SectorName%? I recently knocked out one pirate cluster from there.",
	},
	_const.Replicas + "_out_scout_" + Good + "_6": {
		_const.RU: "Имей в виду, на %BaseName% можно заполучить неплохой контракт по сопровождению конвоя.",
		_const.EN: "Keep in mind that you can get a good convoy escort contract on %BaseName%.",
	},
	_const.Replicas + "_out_scout_" + Good + "_7": {
		_const.RU: "Если ввяжешься в какие-то неприятности, %UserName%, мой тебе совет - постарайся попросту сбежать.",
		_const.EN: "If you get into any trouble, %UserName%, my advice to you is to simply try to escape.",
	},
	_const.Replicas + "_out_scout_" + Good + "_8": {
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
		_const.RU: "Ищешь новых охотничьих угодий? Тогда наведайся в %SectorName%.",
		_const.EN: "Looking for new hunting grounds? Then check out %SectorName%.",
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

	_const.Replicas + "_builder_" + Good + "_1": {
		_const.RU: "Привет, как дела?",
		_const.EN: "Hi, how are you doing?",
	},
	_const.Replicas + "_warrior_" + Good + "_1": {
		_const.RU: "Мы отправляемся в сектор %SectorName%, можешь присоединиться!",
		_const.EN: "We're heading to sector %SectorName%, you're welcome to join!",
	},
	_const.Replicas + "_expedition_" + Good + "_1": {
		_const.RU: "Мы пробуем освоить сектор %SectorName%, будем рады содействию!",
		_const.EN: "We're trying to explore sector %SectorName%, we'd appreciate your help!",
	},
	_const.Replicas + "_expedition_" + Good + "_2": {
		_const.RU: "Наша задача — выбирать противника с территории пустоши в секторе %SectorName%! Будем рады содействию!",
		_const.EN: "Our task is to choose an enemy from the wasteland territory in the %SectorName% sector! We'll appreciate your help!",
	},

	_const.Replicas + "_miner_" + Bad + "_1": {
		_const.RU: "Ну, иди сюда, попробуй меня трахнуть — я тебя сам трахну!",
		_const.EN: "Well, come here, try to fuck me - I’ll fuck you myself!",
	},
	_const.Replicas + "_guard_" + Bad + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "Don't make it worse!",
	},
	_const.Replicas + "_out_scout_" + Bad + "_1": {
		_const.RU: "Лучше уходи, пока можешь!",
		_const.EN: "Better leave while you can!",
	},
	_const.Replicas + "_in_scout_" + Bad + "_1": {
		_const.RU: "Приветствую. Хочешь внести плату за проезд?",
		_const.EN: "Greetings. Do you want to pay the fare?",
	},
	_const.Replicas + "_transport_" + Bad + "_1": {
		_const.RU: "Если ты помешаешь мне торговать, то я за себя не ручаюсь!",
		_const.EN: "If you interfere with my trade, I won't be responsible for my actions!",
	},
	_const.Replicas + "_builder_" + Bad + "_1": {
		_const.RU: "Отвали, ты мне не интересен",
		_const.EN: "Get lost, you're not interesting to me",
	},
	_const.Replicas + "_warrior_" + Bad + "_1": {
		_const.RU: "Советую тебе избегать встречи со мной!",
		_const.EN: "I advise you to avoid meeting me!",
	},
	_const.Replicas + "_expedition_" + Bad + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "Don't come near!",
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
