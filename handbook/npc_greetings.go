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
		_const.RU:   "Ты что-то хотел, %UserName%?",
		_const.EN:   "Did you want something, %UserName%?",
		_const.ZhCN: "你有什么事吗，<span class=\"importantly\">%UserName%</span>？",
	},
	_const.Replicas + "_miner_" + Neutral + "_2": {
		_const.RU:   "Я весь в работе и у меня мало времени.",
		_const.EN:   "I'm busy with work and have little time.",
		_const.ZhCN: "我正忙着工作，没什么时间。",
	},
	_const.Replicas + "_miner_" + Neutral + "_3": {
		_const.RU:   "Чего бы ты от меня ни желал, поторопись с этим.",
		_const.EN:   "Whatever you want from me, hurry up with it.",
		_const.ZhCN: "无论你想从我这里得到什么，快点说吧。",
	},
	_const.Replicas + "_miner_" + Neutral + "_4": {
		_const.RU:   "Руда… Руда… Руда… И так цикл за циклом.",
		_const.EN:   "Ore... Ore... Ore... And so cycle after cycle.",
		_const.ZhCN: "矿石……矿石……矿石……就这样一个周期接着一个周期。",
	},
	_const.Replicas + "_miner_" + Neutral + "_5": {
		_const.RU:   "Только не говори, %UserName%, что ты тоже будешь меня донимать о скрытых месторождениях Тория.",
		_const.EN:   "Don't tell me, %UserName%, that you will also pester me about hidden thorium deposits.",
		_const.ZhCN: "<span class=\"importantly\">%UserName%</span>，别告诉我你也打算纠缠我关于隐藏的钍矿藏的事。",
	},
	_const.Replicas + "_miner_" + Neutral + "_6": {
		_const.RU:   "Слышал уже? Цены на руды вновь обвалились. Судя по всему, кто-то неумело пользуется рынком, излишне перенасыщая его.",
		_const.EN:   "Have you heard it already? Ore prices have fallen again. Apparently, someone is using the market ineptly, oversaturating it.",
		_const.ZhCN: "你听说了吗？矿石价格又跌了。看来有人不擅长使用市场，导致供应过剩。",
	},
	_const.Replicas + "_miner_" + Neutral + "_7": {
		_const.RU:   "Знаешь, %UserName%, я изрядно нервничаю, когда незнакомый синтет вдруг решает затеять со мной беседу.",
		_const.EN:   "You know, %UserName%, I get pretty nervous when an unfamiliar synth suddenly decides to start a conversation with me.",
		_const.ZhCN: "你知道吗，<span class=\"importantly\">%UserName%</span>，当一个陌生的合成体突然决定和我交谈时，我会很紧张。",
	},
	_const.Replicas + "_miner_" + Neutral + "_8": {
		_const.RU:   "Пиратские Кластеры, безумные машины из пустошей, непомерные налоговые ставки фракций… Да зачем я вообще решил всем этим заниматься!?",
		_const.EN:   "Pirate Clusters, crazy machines from the wastelands, exorbitant tax rates of factions... Why did I even decide to do all this!?",
		_const.ZhCN: "海盗集群、荒原上的疯狂机器、派系的过高税率……我到底为什么要干这些！？",
	},

	_const.Replicas + "_guard_" + Neutral + "_1": {
		_const.RU:   "Есть ли тебе о чём стоит сообщить, %UserName%?",
		_const.EN:   "Do you have anything to report, %UserName%?",
		_const.ZhCN: "你有什么要报告的吗，<span class=\"importantly\">%UserName%</span>？",
	},
	_const.Replicas + "_guard_" + Neutral + "_2": {
		_const.RU:   "Ну, говори-говори! Я весь в ожидании.",
		_const.EN:   "Well, speak up! I'm all in anticipation.",
		_const.ZhCN: "嗯，快说！我正在等你的消息。",
	},
	_const.Replicas + "_guard_" + Neutral + "_3": {
		_const.RU:   "Судя по твоему виду, я кое-что должен знать.",
		_const.EN:   "By the look of you, I should know something.",
		_const.ZhCN: "看你的样子，我应该知道些什么。",
	},
	_const.Replicas + "_guard_" + Neutral + "_4": {
		_const.RU:   "Пробудешь хотя бы цикл в патруле по пустошам, сразу заскучаешь о войне фракций.",
		_const.EN:   "Once you spend at least a cycle on patrol in the wastelands, you will immediately become bored with the faction war.",
		_const.ZhCN: "只要你在荒原上巡逻一个周期，就会立刻对派系战争感到厌倦。",
	},
	_const.Replicas + "_guard_" + Neutral + "_5": {
		_const.RU:   "Дай угадаю, ты видел что-то и теперь хочешь мне об этом сообщить?",
		_const.EN:   "Let me guess, you saw something and now you want to tell me about it?",
		_const.ZhCN: "让我猜猜，你看到了什么，现在想告诉我？",
	},
	_const.Replicas + "_guard_" + Neutral + "_6": {
		_const.RU:   "Либо у тебя возникли проблемы с кем-то конкретным, либо с законами фракций.",
		_const.EN:   "Either you have problems with someone specific, or with the laws of factions.",
		_const.ZhCN: "要么你和某个特定的人有麻烦，要么是违反了派系的法律。",
	},
	_const.Replicas + "_guard_" + Neutral + "_7": {
		_const.RU:   "Я воплощение справедливости.",
		_const.EN:   "I am the embodiment of justice.",
		_const.ZhCN: "我是正义的化身。",
	},
	_const.Replicas + "_guard_" + Neutral + "_8": {
		_const.RU:   "Я несу закон, и абсолютно все обязаны ему следовать!",
		_const.EN:   "I bring the law, and absolutely everyone is obliged to follow it!",
		_const.ZhCN: "我代表法律，所有人都必须遵守它！",
	},
	_const.Replicas + "_guard_" + Neutral + "_9": {
		_const.RU:   "Мои полномочия уникальны. Помни об этом, %UserName%.",
		_const.EN:   "My credentials are unique. Remember this, %UserName%.",
		_const.ZhCN: "我的权限是独一无二的。<span class=\"importantly\">%UserName%</span>，请记住这一点。",
	},

	_const.Replicas + "_out_scout_" + Neutral + "_1": {
		_const.RU:   "Любопытно. У тебя имеется ко мне некое предложение, %UserName%?",
		_const.EN:   "Curious. Do you have a proposal for me, %UserName%?",
		_const.ZhCN: "有趣。你有什么提议给我吗，<span class=\"importantly\">%UserName%</span>？",
	},
	_const.Replicas + "_out_scout_" + Neutral + "_2": {
		_const.RU:   "Видишь ли, я на экстра важном задании.",
		_const.EN:   "You see, I'm on an extra important mission.",
		_const.ZhCN: "你看，我现在在执行一项非常重要的任务。",
	},
	_const.Replicas + "_out_scout_" + Neutral + "_3": {
		_const.RU:   "Нет времени, нет времени на всё это.",
		_const.EN:   "There is no time, there is no time for all this.",
		_const.ZhCN: "没时间，没时间应付这些。",
	},
	_const.Replicas + "_out_scout_" + Neutral + "_4": {
		_const.RU:   "Неужто ты следил за мной?",
		_const.EN:   "Could it be that you were following me?",
		_const.ZhCN: "难道你在跟踪我？",
	},
	_const.Replicas + "_out_scout_" + Neutral + "_5": {
		_const.RU:   "Какая неожиданная встреча, но только не для меня.",
		_const.EN:   "What an unexpected meeting, but not for me.",
		_const.ZhCN: "真是意外的相遇，但对我来说不是。",
	},
	_const.Replicas + "_out_scout_" + Neutral + "_6": {
		_const.RU:   "Хм, я запомню тебя, %UserName%.",
		_const.EN:   "Hmm, I'll remember you, %UserName%.",
		_const.ZhCN: "嗯，我会记住你，<span class=\"importantly\">%UserName%</span>。",
	},

	_const.Replicas + "_in_scout_" + Neutral + "_1": {
		_const.RU:   "О-о, вы только гляньте, кто ко мне пожаловал!",
		_const.EN:   "Oh, just look who's come to visit me!",
		_const.ZhCN: "哦，看看是谁来了！",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_2": {
		_const.RU:   "Не зли меня, %UserName%, если только не желаешь огрести потом проблем.",
		_const.EN:   "Don't make me angry, %UserName%, unless you want to cause problems later.",
		_const.ZhCN: "别惹我生气，<span class=\"importantly\">%UserName%</span>，除非你想以后惹上麻烦。",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_3": {
		_const.RU:   "Я недолюбливаю особо болтливых.",
		_const.EN:   "I don't like particularly talkative people.",
		_const.ZhCN: "我不喜欢话多的人。",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_4": {
		_const.RU:   "Хочешь поторговаться о сведениях?",
		_const.EN:   "Do you want to bargain for information?",
		_const.ZhCN: "你想就信息讨价还价吗？",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_6": {
		_const.RU:   "И не надейся, в моём отряде для тебя места не будет.",
		_const.EN:   "And don’t get your hopes up, there won’t be a place for you in my squad.",
		_const.ZhCN: "别抱希望了，我的队伍里不会有你的位置。",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_7": {
		_const.RU:   "Хм, может, мне стоит посмотреть, что ты там перевозишь в своём грузовом отсеке?",
		_const.EN:   "Hmm, maybe I should take a look at what you're hauling in that cargo hold of yours?",
		_const.ZhCN: "啧，我是不是该检查检查你货舱里运了些什么？",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_8": {
		_const.RU:   "Двигай далее, не задерживайся.",
		_const.EN:   "Move on, don't linger.",
		_const.ZhCN: "继续前进，别停留。",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_9": {
		_const.RU:   "Я ничего не покупаю и не продаю! Мои интересы... более специфичны.",
		_const.EN:   "I don't buy or sell anything! My interests... are more specific.",
		_const.ZhCN: "我不买也不卖任何东西！我的兴趣……更具体。",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_10": {
		_const.RU:   "Сегодня я при хорошем настроении, так что прощаю твою наглость.",
		_const.EN:   "I'm in a good mood today, so I forgive your impudence.",
		_const.ZhCN: "今天我心情不错，所以原谅你的无礼。",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_11": {
		_const.RU:   "Только дай мне повод, %UserName%, и я прострелю твою подвеску.",
		_const.EN:   "Just give me a reason, %UserName%, and I'll shoot through your pendant.",
		_const.ZhCN: "只要给我一个理由，<span class=\"importantly\">%UserName%</span>，我就打穿你的吊坠。",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_12": {
		_const.RU:   "Может, подскажешь твой дальнейший маршрут передвижения?",
		_const.EN:   "Maybe you can tell me your future route?",
		_const.ZhCN: "也许你能告诉我你接下来的行进路线？",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_13": {
		_const.RU:   "Знаешь, %UserName%, в пустошах может приключиться всякое.",
		_const.EN:   "You know, %UserName%, anything can happen in the wasteland.",
		_const.ZhCN: "你知道吗，<span class=\"importantly\">%UserName%</span>，在荒原上什么事都有可能发生。",
	},

	_const.Replicas + "_transport_" + Neutral + "_1": {
		_const.RU:   "Я просто вожу грузы.",
		_const.EN:   "I'm just hauling loads.",
		_const.ZhCN: "我只是运送货物。",
	},
	_const.Replicas + "_transport_" + Neutral + "_2": {
		_const.RU:   "Туда-сюда, от базы к базе.",
		_const.EN:   "Back and forth, from base to base.",
		_const.ZhCN: "来来回回，从基地到基地。",
	},
	_const.Replicas + "_transport_" + Neutral + "_3": {
		_const.RU:   "Честно говоря, мой доход не самый внушительный. Лучше бы я пошёл в добытчики.",
		_const.EN:   "To be honest, my income is not the most impressive. It would be better if I became a miner.",
		_const.ZhCN: "说实话，我的收入并不高。要是当初去当矿工就好了。",
	},
	_const.Replicas + "_transport_" + Neutral + "_4": {
		_const.RU:   "Зачем ты меня тревожишь?",
		_const.EN:   "Why are you bothering me?",
		_const.ZhCN: "你为什么要打扰我？",
	},
	_const.Replicas + "_transport_" + Neutral + "_5": {
		_const.RU:   "Разве не видно - я перевожу ценный груз!",
		_const.EN:   "Can't you see - I'm transporting valuable cargo!",
		_const.ZhCN: "难道你看不到——我在运输贵重货物！",
	},
	_const.Replicas + "_transport_" + Neutral + "_6": {
		_const.RU:   "Надеюсь... ты... не будешь мне угрожать?",
		_const.EN:   "I hope... you... won't threaten me?",
		_const.ZhCN: "我希望……你……不会威胁我？",
	},
	_const.Replicas + "_transport_" + Neutral + "_7": {
		_const.RU:   "Прочь с моего пути! У меня важная доставка!",
		_const.EN:   "Get out of my way! I have an important delivery!",
		_const.ZhCN: "让开！我有重要的货物要送！",
	},
	_const.Replicas + "_transport_" + Neutral + "_8": {
		_const.RU:   "Как же ты не вовремя %UserName%!",
		_const.EN:   "How are you at the wrong time %UserName%!",
		_const.ZhCN: "<span class=\"importantly\">%UserName%</span>，你怎么偏偏这时候来！",
	},
	_const.Replicas + "_transport_" + Neutral + "_9": {
		_const.RU:   "Ну почему, как только ты берёшь контракт по доставке, всем сразу от тебя чего-то нужно!?",
		_const.EN:   "Why, as soon as you take on a delivery contract, does everyone immediately need something from you!?",
		_const.ZhCN: "为什么每次我接了送货合同，所有人都突然需要我帮忙！？",
	},
	_const.Replicas + "_transport_" + Neutral + "_10": {
		_const.RU:   "Отвечаю сразу: я ничего не продаю!",
		_const.EN:   "I answer right away: I don’t sell anything!",
		_const.ZhCN: "我直接回答：我不卖任何东西！",
	},
	_const.Replicas + "_transport_" + Neutral + "_11": {
		_const.RU:   "Нет, я не поверю, что ты из стражей. И нет, я не позволю тебе досмотреть мой грузовой отсек.",
		_const.EN:   "No, I don't believe you're from the Guard. And no, you're not inspecting my cargo hold.",
		_const.ZhCN: "不，我根本不信你是守卫。更别想检查我的货舱。",
	},
	_const.Replicas + "_transport_" + Neutral + "_12": {
		_const.RU:   "Ты друг, или... враг?",
		_const.EN:   "Are you a friend, or... an enemy?",
		_const.ZhCN: "你是朋友，还是……敌人？",
	},
	_const.Replicas + "_transport_" + Neutral + "_13": {
		_const.RU:   "Пытаешься вынюхать мой маршрут?",
		_const.EN:   "Are you trying to sniff out my route?",
		_const.ZhCN: "你在试图探听我的路线吗？",
	},
	_const.Replicas + "_transport_" + Neutral + "_14": {
		_const.RU:   "Транспортник на связи.",
		_const.EN:   "The transport worker is in touch.",
		_const.ZhCN: "运输员在线。",
	},
	_const.Replicas + "_transport_" + Neutral + "_15": {
		_const.RU:   "А-а, ты наверное один из недовольных клиентов?",
		_const.EN:   "Ah, you must be one of the dissatisfied customers?",
		_const.ZhCN: "啊，你大概是一个不满意的客户吧？",
	},

	_const.Replicas + "_builder_" + Neutral + "_1": {
		_const.RU:   "Я строю во имя фракции.",
		_const.EN:   "I build in the name of the faction.",
		_const.ZhCN: "我为派系而建造。",
	},
	_const.Replicas + "_builder_" + Neutral + "_2": {
		_const.RU:   "Чёрт побери, я так обожаю свою работу!",
		_const.EN:   "Damn it, I love my job so much!",
		_const.ZhCN: "该死，我太爱我的工作了！",
	},
	_const.Replicas + "_builder_" + Neutral + "_3": {
		_const.RU:   "У-у, у меня тут дел на несколько циклов вперёд.",
		_const.EN:   "Ooh, I've got a few cycles of work to do here.",
		_const.ZhCN: "哦，我这里有几轮的工作要做。",
	},
	_const.Replicas + "_builder_" + Neutral + "_4": {
		_const.RU:   "Решил по отвлекать меня? Ладно, заодно и перерыв будет.",
		_const.EN:   "Decided to distract me? Okay, there will be a break at the same time.",
		_const.ZhCN: "决定来分散我的注意力？好吧，顺便休息一下。",
	},
	_const.Replicas + "_builder_" + Neutral + "_5": {
		_const.RU:   "Хочешь тоже податься в строители %UserName%? Поздно, места все заняты.",
		_const.EN:   "Do you also want to join the builders of %UserName%? It's late, the seats are all taken.",
		_const.ZhCN: "你也想加入建筑队伍吗，<span class=\"importantly\">%UserName%</span>？太晚了，位置都满了。",
	},
	_const.Replicas + "_builder_" + Neutral + "_6": {
		_const.RU:   "Поди пойми, где должна располагаться эта сеть турелей…",
		_const.EN:   "Just figure out where this network of turrets should be located...",
		_const.ZhCN: "搞清楚这些炮塔网络应该布置在哪里……",
	},
	_const.Replicas + "_builder_" + Neutral + "_7": {
		_const.RU:   "Я возвожу строения, а не оказываю информационные услуги.",
		_const.EN:   "I erect buildings, and do not provide information services.",
		_const.ZhCN: "我建造建筑物，而不是提供信息服务。",
	},
	_const.Replicas + "_builder_" + Neutral + "_8": {
		_const.RU:   "Ты как раз находишься на одном из запланированных для строительства участков.",
		_const.EN:   "You are just on one of the planned construction sites.",
		_const.ZhCN: "你正好在一个计划中的建筑工地。",
	},
	_const.Replicas + "_builder_" + Neutral + "_9": {
		_const.RU:   "Моя работа в самом разгаре, %UserName%.",
		_const.EN:   "My work is in full swing, %UserName%.",
		_const.ZhCN: "我的工作正在进行中，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_builder_" + Neutral + "_10": {
		_const.RU:   "Инструменты ещё не остыли, мне некогда на тебя отвлекаться.",
		_const.EN:   "The instruments have not yet cooled down, I have no time to be distracted by you.",
		_const.ZhCN: "工具还没冷却，我没空分心应对你。",
	},

	_const.Replicas + "_warrior_" + Neutral + "_1": {
		_const.RU:   "Нахожусь в патруле.",
		_const.EN:   "I'm on patrol.",
		_const.ZhCN: "我正在巡逻。",
	},
	_const.Replicas + "_warrior_" + Neutral + "_2": {
		_const.RU:   "Я защищаю границы фракции.",
		_const.EN:   "I protect the faction's borders.",
		_const.ZhCN: "我保卫派系的边界。",
	},
	_const.Replicas + "_warrior_" + Neutral + "_3": {
		_const.RU:   "В поиске новых целей.",
		_const.EN:   "In search of new targets.",
		_const.ZhCN: "寻找新目标。",
	},
	_const.Replicas + "_warrior_" + Neutral + "_4": {
		_const.RU:   "Где-то просочился враг, %UserName%?",
		_const.EN:   "Has the enemy leaked somewhere, %UserName%?",
		_const.ZhCN: "敌人是不是在某个地方渗透进来了，<span class=\"importantly\">%UserName%</span>？",
	},
	_const.Replicas + "_warrior_" + Neutral + "_5": {
		_const.RU:   "Всегда готов к бою!",
		_const.EN:   "Always ready for battle!",
		_const.ZhCN: "随时准备战斗！",
	},
	_const.Replicas + "_warrior_" + Neutral + "_6": {
		_const.RU:   "Я готов... к чему угодно.",
		_const.EN:   "I'm ready... for anything.",
		_const.ZhCN: "我准备好……面对任何事情。",
	},
	_const.Replicas + "_warrior_" + Neutral + "_7": {
		_const.RU:   "Все стволы заряжены.",
		_const.EN:   "All barrels are loaded.",
		_const.ZhCN: "所有枪械都已装填完毕。",
	},
	_const.Replicas + "_warrior_" + Neutral + "_8": {
		_const.RU:   "Лента боекомплекта подана.",
		_const.EN:   "The ammunition belt has been supplied.",
		_const.ZhCN: "弹药带已装载。",
	},
	_const.Replicas + "_warrior_" + Neutral + "_9": {
		_const.RU:   "Ох, сейчас бы добротную заварушку!",
		_const.EN:   "Oh, now there would be a good mess!",
		_const.ZhCN: "哦，现在要是能有一场激烈的战斗就好了！",
	},
	_const.Replicas + "_warrior_" + Neutral + "_10": {
		_const.RU:   "Жду не дождусь, когда произойдет новый виток войны фракций!",
		_const.EN:   "I can't wait for the next round of faction war to happen!",
		_const.ZhCN: "我迫不及待地等待派系战争的新一轮爆发！",
	},
	_const.Replicas + "_warrior_" + Neutral + "_11": {
		_const.RU:   "Лучше не зли меня, %UserName%.",
		_const.EN:   "You better not make me angry, %UserName%.",
		_const.ZhCN: "<span class=\"importantly\">%UserName%</span>，你最好别惹我生气。",
	},
	_const.Replicas + "_warrior_" + Neutral + "_12": {
		_const.RU:   "Тоже хочешь подраться, %UserName%?",
		_const.EN:   "Do you want to fight too, %UserName%?",
		_const.ZhCN: "你也想打架吗，<span class=\"importantly\">%UserName%</span>？",
	},

	_const.Replicas + "_expedition_" + Neutral + "_1": {
		_const.RU:   "Какое же это убогое место!",
		_const.EN:   "What a wretched place this is!",
		_const.ZhCN: "这是多么糟糕的地方啊！",
	},
	_const.Replicas + "_expedition_" + Neutral + "_2": {
		_const.RU:   "Я бывал там, где ты, %UserName%, даже и представить себе не можешь.",
		_const.EN:   "I've been to places you, %UserName%, can't even imagine.",
		_const.ZhCN: "我去过你，<span class=\"importantly\">%UserName%</span>，无法想象的地方。",
	},
	_const.Replicas + "_expedition_" + Neutral + "_3": {
		_const.RU:   "Я видел многое и могу открыто заявить тебе, %UserName% - кое-что крайне важное от вас фракции всё-таки скрывают.",
		_const.EN:   "I have seen a lot and I can openly tell you, %UserName% - the factions are still hiding something extremely important from you.",
		_const.ZhCN: "我见过很多东西，可以明确告诉你，<span class=\"importantly\">%UserName%</span>——派系仍然在对你隐瞒一些极其重要的事情。",
	},
	_const.Replicas + "_expedition_" + Neutral + "_4": {
		_const.RU:   "За безопасными секторами совсем ничего нету. Имей это ввиду.",
		_const.EN:   "There is absolutely nothing beyond the safe sectors. Keep this in mind.",
		_const.ZhCN: "安全区域之外什么都没有。记住这一点。",
	},
	_const.Replicas + "_expedition_" + Neutral + "_5": {
		_const.RU:   "Не вздумай приближаться к местной органической фауне, %UserName%, они… не тактильны.",
		_const.EN:   "Don't even try to get close to the local organic fauna, %UserName%, they... are not tactile.",
		_const.ZhCN: "千万别靠近当地的有机动物群，<span class=\"importantly\">%UserName%</span>，它们……不友好。",
	},
	_const.Replicas + "_expedition_" + Neutral + "_6": {
		_const.RU:   "Если случайно забредёшь в реликтовый город, считай, что ты потерял там свой корпус.",
		_const.EN:   "If you accidentally wander into a relic city, consider that you have lost your corps there.",
		_const.ZhCN: "如果你不小心闯入遗迹城市，就等于失去了你的躯体。",
	},
	_const.Replicas + "_expedition_" + Neutral + "_7": {
		_const.RU:   "Если вдруг станешь принимать посреди пустошей странный сигнал, немедленно оттуда уходи!",
		_const.EN:   "If you suddenly start receiving a strange signal in the middle of the wasteland, leave immediately!",
		_const.ZhCN: "如果你在荒原中突然接收到奇怪的信号，立即离开那里！",
	},
	_const.Replicas + "_expedition_" + Neutral + "_8": {
		_const.RU:   "Я участник экспедиции, а ты?",
		_const.EN:   "I'm a member of the expedition, and you?",
		_const.ZhCN: "我是探险队的一员，你呢？",
	},
	_const.Replicas + "_expedition_" + Neutral + "_9": {
		_const.RU:   "Скоро снова в путь.",
		_const.EN:   "We'll be on the road again soon.",
		_const.ZhCN: "我们很快又要上路了。",
	},
	_const.Replicas + "_expedition_" + Neutral + "_10": {
		_const.RU:   "Только самые смелые, %UserName%, записываются в экспедиции.",
		_const.EN:   "Only the bravest, %UserName%, sign up for the expedition.",
		_const.ZhCN: "只有最勇敢的人，<span class=\"importantly\">%UserName%</span>，才会加入探险队。",
	},
	_const.Replicas + "_expedition_" + Neutral + "_11": {
		_const.RU:   "Если тебе нечего терять, %UserName%, примкни к нам.",
		_const.EN:   "If you have nothing to lose, %UserName%, join us.",
		_const.ZhCN: "如果你没有什么可失去的，<span class=\"importantly\">%UserName%</span>，加入我们吧。",
	},

	// Good
	_const.Replicas + "_miner_" + Good + "_1": {
		_const.RU:   "Всегда рад встретить иного дружественного синтета. В нынешнее время, это прямо-таки роскошь.",
		_const.EN:   "Always glad to meet another friendly synth. Nowadays, this is downright luxury.",
		_const.ZhCN: "总是很高兴遇到其他友好的合成体。在当今时代，这简直是奢侈。",
	},
	_const.Replicas + "_miner_" + Good + "_2": {
		_const.RU:   "Как думаешь, нам - добытчикам, должны доплачивать за фактор опасности?",
		_const.EN:   "Do you think we, the miners, should be paid extra for the danger factor?",
		_const.ZhCN: "你认为我们矿工应该因为危险因素获得额外报酬吗？",
	},
	_const.Replicas + "_miner_" + Good + "_3": {
		_const.RU:   "Скажу по секрету: в этом секторе ты мало чем сумеешь разжиться. Попробуй осмотреть прилегающие территории.",
		_const.EN:   "I’ll tell you a secret: you won’t be able to get hold of much in this sector. Try to explore the surrounding areas.",
		_const.ZhCN: "我悄悄告诉你：在这个区域你几乎找不到什么好东西。试着探索周边区域吧。",
	},
	_const.Replicas + "_miner_" + Good + "_4": {
		_const.RU:   "%UserName%, если у тебя есть соответствующая руда, попробуй наведаться на %BaseName%. Сейчас, там можно выгодно реализовать руду.",
		_const.EN:   "%UserName%, if you have the appropriate ore, try visiting %BaseName%. Now, ore can be profitably sold there.",
		_const.ZhCN: "<span class=\"importantly\">%UserName%</span>，如果你有相应的矿石，可以去<span class=\"importantly\">%BaseName%</span>看看。现在那里可以卖个好价钱。",
	},
	_const.Replicas + "_miner_" + Good + "_5": {
		_const.RU:   "Тоже решил заняться добычей руды, а по итогу не особо то и идёт дело?",
		_const.EN:   "You also decided to start mining ore, but in the end things aren’t going well?",
		_const.ZhCN: "你也决定开始采矿了，但结果并不理想，对吧？",
	},
	_const.Replicas + "_miner_" + Good + "_6": {
		_const.RU:   "Сегодня мне повезло как никогда ранее! Столько руды я ещё никогда за всё своё существование не добывал.",
		_const.EN:   "Today I am luckier than ever before! I have never mined so much ore in my entire existence.",
		_const.ZhCN: "今天我运气超好！我这辈子从来没挖到过这么多矿石。",
	},
	_const.Replicas + "_miner_" + Good + "_7": {
		_const.RU:   "Только между нами %UserName%: если рассчитываешь выгодно переработать руду, советую на время позабыть о %BaseName%. Ты ещё и останешься должным тамошнему диспетчеру.",
		_const.EN:   "Just between us %UserName%: if you expect to profitably recycle ore, I advise you to forget about %BaseName% for a while. You will also remain indebted to the dispatcher there.",
		_const.ZhCN: "只跟你讲，<span class=\"importantly\">%UserName%</span>：如果你想有利可图地加工矿石，我建议暂时别去<span class=\"importantly\">%BaseName%</span>。否则你还得欠那里的调度员一个人情。",
	},
	_const.Replicas + "_miner_" + Good + "_8": {
		_const.RU:   "Ходят слухи, что в %SectorName% обнаружили месторождения Тория. Сам я не видел. Да и дорога туда опасная. Но… вдруг, ты %UserName% заинтересуешься этим?",
		_const.EN:   "There are rumors that thorium deposits have been discovered in %SectorName%. I haven't seen it myself. And the road there is dangerous. But... what if you %UserName% become interested in this?",
		_const.ZhCN: "有传言说在<span class=\"importantly\">%SectorName%</span>发现了钍矿床。我自己没见到过，而且那条路很危险。不过……也许你会感兴趣，<span class=\"importantly\">%UserName%</span>？",
	},

	_const.Replicas + "_guard_" + Good + "_1": {
		_const.RU:   "О тебе хорошо отзываются, %UserName%. Я это учту.",
		_const.EN:   "You are well spoken of, %UserName%. I'll take this into account.",
		_const.ZhCN: "大家都对你评价很好，<span class=\"importantly\">%UserName%</span>。我会记住这一点。",
	},
	_const.Replicas + "_guard_" + Good + "_2": {
		_const.RU:   "Добропорядочные синтеты наподобие тебя, %UserName% - редкость.",
		_const.EN:   "Respectable synths like you, %UserName%, are rare.",
		_const.ZhCN: "像你这样正直的合成体，<span class=\"importantly\">%UserName%</span>，很少见。",
	},
	_const.Replicas + "_guard_" + Good + "_3": {
		_const.RU:   "Пускай твой дальнейший путь будет безопасен, %UserName%.",
		_const.EN:   "May your future journey be safe, %UserName%.",
		_const.ZhCN: "愿你的旅途平安，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_guard_" + Good + "_4": {
		_const.RU:   "У нас здесь всё под полным контролем, а ты под надёжной защитой, %UserName%.",
		_const.EN:   "We have everything here under complete control, and you are under reliable protection, %UserName%.",
		_const.ZhCN: "这里的一切都在完全控制之下，而你受到可靠的保护，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_guard_" + Good + "_5": {
		_const.RU:   "Патруль, докладываю: у нас всё спокойно.",
		_const.EN:   "Patrol, I report: everything is calm here.",
		_const.ZhCN: "巡逻队报告：这里一切平静。",
	},
	_const.Replicas + "_guard_" + Good + "_6": {
		_const.RU:   "Будут неприятности, %UserName%? - дай знать, разберёмся.",
		_const.EN:   "Are you in trouble, %UserName%? - Let me know, we'll sort it out.",
		_const.ZhCN: "你遇到麻烦了吗，<span class=\"importantly\">%UserName%</span>？——告诉我，我们会解决的。",
	},
	_const.Replicas + "_guard_" + Good + "_7": {
		_const.RU:   "Закон на твоей стороне, %UserName%.",
		_const.EN:   "The law is on your side, %UserName%.",
		_const.ZhCN: "法律站在你这边，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_guard_" + Good + "_8": {
		_const.RU:   "Ты пример для подражания, %UserName%? Так держать!",
		_const.EN:   "Are you a role model, %UserName%? Keep it up!",
		_const.ZhCN: "你是榜样，<span class=\"importantly\">%UserName%</span>？继续保持！",
	},

	_const.Replicas + "_out_scout_" + Good + "_1": {
		_const.RU:   "Ну, здравствуй, %UserName%.",
		_const.EN:   "Well, hello %UserName%.",
		_const.ZhCN: "嗯，你好，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_out_scout_" + Good + "_2": {
		_const.RU:   "Рад этой встрече, %UserName%.",
		_const.EN:   "Glad to meet you, %UserName%.",
		_const.ZhCN: "很高兴见到你，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_out_scout_" + Good + "_3": {
		_const.RU:   "О-о, дружественный синтет! Великая редкость при моей работе.",
		_const.EN:   "Oooh, friendly synth! A great rarity in my work.",
		_const.ZhCN: "哦，友好的合成体！在我的工作中，这种情况非常罕见。",
	},
	_const.Replicas + "_out_scout_" + Good + "_4": {
		_const.RU:   "Надеюсь, с тобой по пути ничего не приключилось?",
		_const.EN:   "I hope nothing happened to you along the way?",
		_const.ZhCN: "我希望你在路上没有遇到什么事吧？",
	},
	_const.Replicas + "_out_scout_" + Good + "_5": {
		_const.RU:   "Как тебе нынешние контракты у фракций?",
		_const.EN:   "What do you think of the current faction contracts?",
		_const.ZhCN: "你觉得目前派系的合同怎么样？",
	},
	_const.Replicas + "_out_scout_" + Good + "_6": {
		_const.RU:   "Если ввяжешься в какие-то неприятности, %UserName%, мой тебе совет - постарайся попросту сбежать.",
		_const.EN:   "If you get into any trouble, %UserName%, my advice to you is to simply try to escape.",
		_const.ZhCN: "如果你陷入麻烦，<span class=\"importantly\">%UserName%</span>，我的建议是尽量逃跑。",
	},
	_const.Replicas + "_out_scout_" + Good + "_7": {
		_const.RU:   "Не вздумай верить стражам. Правосудие в пустошах… пф, комично!",
		_const.EN:   "Don't you dare trust the guards. Justice in the wasteland... pfft, comical!",
		_const.ZhCN: "千万别相信守卫。荒原上的正义……呵呵，太讽刺了！",
	},

	_const.Replicas + "_in_scout_" + Good + "_1": {
		_const.RU:   "Славно встретить кого-то из своих.",
		_const.EN:   "It's nice to meet one of your own.",
		_const.ZhCN: "很高兴遇到自己人。",
	},
	_const.Replicas + "_in_scout_" + Good + "_2": {
		_const.RU:   "Ну, похвастайся, сколько за сегодня награбил?",
		_const.EN:   "Well, brag about how much you stole today?",
		_const.ZhCN: "那么，炫耀一下你今天抢了多少东西吧？",
	},
	_const.Replicas + "_in_scout_" + Good + "_3": {
		_const.RU:   "Может, совместно налетим на какого-нибудь торговца?",
		_const.EN:   "Maybe we can raid some merchant together?",
		_const.ZhCN: "也许我们可以一起袭击某个商人？",
	},
	_const.Replicas + "_in_scout_" + Good + "_4": {
		_const.RU:   "Ищешь новых охотничьих угодий?",
		_const.EN:   "Looking for new hunting grounds?",
		_const.ZhCN: "你在寻找新的狩猎场吗？",
	},
	_const.Replicas + "_in_scout_" + Good + "_5": {
		_const.RU:   "Контрабандисты рассказывали, что в %BaseName% платят приличные суммы за доставку соответствующих товаров.",
		_const.EN:   "Smugglers said that %BaseName% paid decent sums for the delivery of relevant goods.",
		_const.ZhCN: "走私者说，在<span class=\"importantly\">%BaseName%</span>运送相关商品可以赚到可观的报酬。",
	},
	_const.Replicas + "_in_scout_" + Good + "_6": {
		_const.RU:   "О-о, кого я вижу! Да о тебе во всех пиратских кластерах только и говорят, %UserName%.",
		_const.EN:   "Oh, who do I see! Yes, all the pirate clusters are talking about you, %UserName%.",
		_const.ZhCN: "哦，我看到了谁！是的，所有海盗集群都在谈论你，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_in_scout_" + Good + "_7": {
		_const.RU:   "Запомни, %UserName%: сначала пускай вываливают всё из грузового отсека, а затем избавляйся как от лишних свидетелей.",
		_const.EN:   "Remember, %UserName%: first make them unload everything from the cargo hold, then dispose of unnecessary witnesses.",
		_const.ZhCN: "记住，%UserName%：先让他们把货舱清空，再处理掉多余的目击者。",
	},
	_const.Replicas + "_in_scout_" + Good + "_8": {
		_const.RU:   "Если засада удалась, %UserName%, постарайся затем побыстрее сменить свою позицию.",
		_const.EN:   "If the ambush is successful, %UserName%, then try to quickly change your position.",
		_const.ZhCN: "如果伏击成功了，<span class=\"importantly\">%UserName%</span>，之后尽量快速改变你的位置。",
	},
	_const.Replicas + "_in_scout_" + Good + "_9": {
		_const.RU:   "Не знаю, как ты, %UserName%, а я обожаю преследовать одиноких торговцев.",
		_const.EN:   "I don't know about you, %UserName%, but I love stalking lone traders.",
		_const.ZhCN: "我不知道你怎么样，<span class=\"importantly\">%UserName%</span>，但我非常喜欢追踪孤独的商人。",
	},
	_const.Replicas + "_in_scout_" + Good + "_10": {
		_const.RU:   "Если тебя кто не понимает сразу - сделай несколько выстрелов по корпусу, это их образумит.",
		_const.EN:   "If someone doesn’t understand you right away, fire a few shots at the body, this will bring them to their senses.",
		_const.ZhCN: "如果有人一开始不理解你——就朝船体开几枪，这会让他们清醒过来。",
	},

	_const.Replicas + "_transport_" + Good + "_1": {
		_const.RU:   "Если у тебя при себе имеется %ProductName%, попробуй посетить %BaseName%, где продашь свой товар по выгодной цене.",
		_const.EN:   "If you have %ProductName% with you, try visiting %BaseName%, where you can sell your product at a good price.",
		_const.ZhCN: "如果你有<span class=\"importantly\">%ProductName%</span>，试着去<span class=\"importantly\">%BaseName%</span>，在那里你可以以优惠的价格出售你的商品。",
	},
	_const.Replicas + "_transport_" + Good + "_2": {
		_const.RU:   "Я следую на %BaseName%, база очень выгодно продаёт %ProductName%!",
		_const.EN:   "I'm heading to %BaseName%, the base sells %ProductName% very profitably!",
		_const.ZhCN: "我要前往<span class=\"importantly\">%BaseName%</span>，基地出售的<span class=\"importantly\">%ProductName%</span>非常划算！",
	},
	_const.Replicas + "_transport_" + Good + "_3": {
		_const.RU:   "Хочешь поболтать о том, о сём? Ну… хорошо, давай.",
		_const.EN:   "Do you want to chat about this and that? Well... okay, go ahead.",
		_const.ZhCN: "你想聊聊这个那个吗？嗯……好吧，继续。",
	},
	_const.Replicas + "_transport_" + Good + "_4": {
		_const.RU:   "Держись подальше от %SectorName%. Немало торговцев там пропало в последнее время.",
		_const.EN:   "Stay away from %SectorName%. Quite a few traders there have disappeared recently.",
		_const.ZhCN: "离<span class=\"importantly\">%SectorName%</span>远点。最近有不少商人在那里失踪了。",
	},
	_const.Replicas + "_transport_" + Good + "_5": {
		_const.RU:   "Сейчас как никогда лучшее время, чтобы начать деятельность торговца %UserName%.",
		_const.EN:   "There has never been a better time to get started as a %UserName% merchant.",
		_const.ZhCN: "现在是成为<span class=\"importantly\">%UserName%</span>商人的最佳时机。",
	},
	_const.Replicas + "_transport_" + Good + "_6": {
		_const.RU:   "Если однажды займешься перевозками %UserName%, то знай, что маршрут, проходящий через %SectorName%, весьма выгоден.",
		_const.EN:   "If one day you start transporting %UserName%, then know that the route passing through %SectorName% is very profitable.",
		_const.ZhCN: "如果你有一天开始运输<span class=\"importantly\">%UserName%</span>，要知道经过<span class=\"importantly\">%SectorName%</span>的路线非常有利可图。",
	},
	_const.Replicas + "_transport_" + Good + "_7": {
		_const.RU:   "Хотя бы от тебя, %UserName%, я не ощущаю исходящей опасности.",
		_const.EN:   "At least from you, %UserName%, I don’t feel any danger coming from you.",
		_const.ZhCN: "至少从你身上，<span class=\"importantly\">%UserName%</span>，我没有感觉到任何危险。",
	},
	_const.Replicas + "_transport_" + Good + "_8": {
		_const.RU:   "Ты только не удивляйся, но сегодня в моём грузовом отсеке весьма много ценных вещей.",
		_const.EN:   "Don't be too surprised, but my cargo hold happens to be full of valuables today.",
		_const.ZhCN: "可别吓到——今天我的货舱里全是值钱货。",
	},
	_const.Replicas + "_transport_" + Good + "_9": {
		_const.RU:   "Ты даже и понятия не имеешь, насколько же выгодный контракт по перевозке я сегодня раздобыл!",
		_const.EN:   "You have no idea how lucrative a transportation contract I got today!",
		_const.ZhCN: "你甚至无法想象我今天拿到了多么有利可图的运输合同！",
	},
	_const.Replicas + "_transport_" + Good + "_10": {
		_const.RU:   "Всё-таки утомительное это дело, %UserName%, сновать между базами и перевозить грузы особой ценности.",
		_const.EN:   "Still, it’s a tedious task, %UserName%, to scurry between bases and transport cargo of special value.",
		_const.ZhCN: "不过，<span class=\"importantly\">%UserName%</span>，在基地之间穿梭并运输特别有价值的货物确实是一项繁琐的任务。",
	},
	//
	_const.Replicas + "_builder_" + Good + "_1": {
		_const.RU:   "Ты истинный друг рабочего синтета, %UserName%.",
		_const.EN:   "You are a true friend of the working synth, %UserName%.",
		_const.ZhCN: "你是一个真正的工作合成体的朋友，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_builder_" + Good + "_2": {
		_const.RU:   "Так держать. Тебя тут всегда рады видеть.",
		_const.EN:   "Keep it up. You're always welcome here.",
		_const.ZhCN: "继续保持！你在这里总是受欢迎的。",
	},
	_const.Replicas + "_builder_" + Good + "_3": {
		_const.RU:   "Хоть ты и не строитель, но мои коллеги тебя жалуют. Так буду делать и я.",
		_const.EN:   "Although you are not a builder, my colleagues favor you. I will do the same.",
		_const.ZhCN: "虽然你不是建筑工，但我的同事们都很喜欢你。我也会这样做的。",
	},
	_const.Replicas + "_builder_" + Good + "_4": {
		_const.RU:   "Что, отвести тебя в безопасное место? Одно такое я знаю.",
		_const.EN:   "What, take you to safety? I know one.",
		_const.ZhCN: "怎么，带你去安全的地方？我知道一个。",
	},
	_const.Replicas + "_builder_" + Good + "_5": {
		_const.RU:   "Случись чего, %UserName%, тебе позволят затеряться среди прочих строителей.",
		_const.EN:   "If something happens, %UserName%, you will be allowed to get lost among other builders.",
		_const.ZhCN: "如果发生什么事，<span class=\"importantly\">%UserName%</span>，你可以混在其他建筑工中避难。",
	},
	_const.Replicas + "_builder_" + Good + "_6": {
		_const.RU:   "Если вдруг решишь сменить свой род деятельности, %UserName% - ты знаешь, где меня искать.",
		_const.EN:   "If you suddenly decide to change your occupation, %UserName% - you know where to look for me.",
		_const.ZhCN: "如果你突然决定换个工作，<span class=\"importantly\">%UserName%</span>——你知道到哪里找我。",
	},

	_const.Replicas + "_warrior_" + Good + "_1": {
		_const.RU:   "Мы отправляемся в сектор %SectorName%, можешь присоединиться!",
		_const.EN:   "We're heading to sector %SectorName%, you're welcome to join!",
		_const.ZhCN: "我们要前往<span class=\"importantly\">%SectorName%</span>区域，欢迎加入！",
	},
	_const.Replicas + "_warrior_" + Good + "_2": {
		_const.RU:   "Можешь рассчитывать на мою поддержку, %UserName%.",
		_const.EN:   "You can count on my support, %UserName%.",
		_const.ZhCN: "你可以依靠我的支持，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_warrior_" + Good + "_3": {
		_const.RU:   "Будь уверен, %UserName%, мы своих не бросаем.",
		_const.EN:   "Rest assured, %UserName%, we do not abandon our own.",
		_const.ZhCN: "放心吧，<span class=\"importantly\">%UserName%</span>，我们不会抛弃自己人。",
	},
	_const.Replicas + "_warrior_" + Good + "_4": {
		_const.RU:   "Надеюсь, ты уже записался в ополчение фракции?",
		_const.EN:   "I hope you have already signed up for the faction's militia?",
		_const.ZhCN: "我希望你已经报名参加了派系的民兵组织？",
	},
	_const.Replicas + "_warrior_" + Good + "_5": {
		_const.RU:   "Вот таких бойцов, как ты, %UserName%, нам и не хватает!",
		_const.EN:   "It's fighters like you, %UserName%, that we need!",
		_const.ZhCN: "像你这样的战士，<span class=\"importantly\">%UserName%</span>，正是我们需要的！",
	},
	_const.Replicas + "_warrior_" + Good + "_6": {
		_const.RU:   "Если вдруг подобный момент наступит, я прикрою тебя, %UserName%.",
		_const.EN:   "If a moment like this ever comes, I've got your back, %UserName%.",
		_const.ZhCN: "如果类似的情况发生，我会掩护你，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_warrior_" + Good + "_7": {
		_const.RU:   "Тебе нужна огневая поддержка? Мой корпус и моё оружие - в твоём распоряжении, %UserName%.",
		_const.EN:   "Do you need fire support? My body and my weapons are at your disposal, %UserName%.",
		_const.ZhCN: "你需要火力支援吗？我的机体和武器都归你使用，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_warrior_" + Good + "_8": {
		_const.RU:   "В нужный час, ты узнаешь на что я действительно способен!",
		_const.EN:   "At the right time, you will find out what I am really capable of!",
		_const.ZhCN: "在紧要关头，你会知道我到底有多厉害！",
	},
	_const.Replicas + "_warrior_" + Good + "_9": {
		_const.RU:   "Я не подведу!",
		_const.EN:   "I won't let you down!",
		_const.ZhCN: "我不会让你失望的！",
	},
	_const.Replicas + "_warrior_" + Good + "_10": {
		_const.RU:   "Я всегда готов к сражению!",
		_const.EN:   "I'm always ready for battle!",
		_const.ZhCN: "我随时准备战斗！",
	},
	_const.Replicas + "_warrior_" + Good + "_11": {
		_const.RU:   "У тебя есть хорошая цель для битвы, %UserName%?",
		_const.EN:   "Do you have a good target for battle, %UserName%?",
		_const.ZhCN: "你有好的战斗目标吗，<span class=\"importantly\">%UserName%</span>？",
	},
	_const.Replicas + "_warrior_" + Good + "_12": {
		_const.RU:   "Для меня будет честью сражаться рядом с тобой, %UserName%.",
		_const.EN:   "It would be an honor to fight alongside you, %UserName%.",
		_const.ZhCN: "能与你并肩作战是我的荣幸，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_warrior_" + Good + "_13": {
		_const.RU:   "Ввяжешься в бой - зови и меня!",
		_const.EN:   "If you get involved in a fight, call me too!",
		_const.ZhCN: "如果你卷入战斗——也叫我一声！",
	},

	_const.Replicas + "_expedition_" + Good + "_1": {
		_const.RU:   "Мы пробуем освоить сектор %SectorName%, будем рады содействию!",
		_const.EN:   "We're trying to explore sector %SectorName%, we'd appreciate your help!",
		_const.ZhCN: "我们正尝试开拓<span class=\"importantly\">%SectorName%</span>区域，非常欢迎你的协助！",
	},
	_const.Replicas + "_expedition_" + Good + "_2": {
		_const.RU:   "Наша задача — выбирать противника с территории пустоши в секторе %SectorName%! Будем рады содействию!",
		_const.EN:   "Our task is to choose an enemy from the wasteland territory in the %SectorName% sector! We'll appreciate your help!",
		_const.ZhCN: "我们的任务是从<span class=\"importantly\">%SectorName%</span>区域的荒原中挑选敌人！非常欢迎你的帮助！",
	},
	_const.Replicas + "_expedition_" + Good + "_3": {
		_const.RU:   "Не желаешь отправиться в одну перспективную экспедицию?",
		_const.EN:   "Would you like to go on a promising expedition?",
		_const.ZhCN: "你想参加一个有前途的探险吗？",
	},
	_const.Replicas + "_expedition_" + Good + "_4": {
		_const.RU:   "Будь у меня такая возможность, я бы поделился с тобой найденными в далёких секторах богатствами, %UserName%.",
		_const.EN:   "If I had the opportunity, I would share with you the wealth I found in distant sectors, %UserName%.",
		_const.ZhCN: "如果我有机会，我会和你分享在遥远区域找到的财富，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_expedition_" + Good + "_5": {
		_const.RU:   "Вот уж кому я точно рад, когда возвращаюсь после долгого странствия, %UserName%!",
		_const.EN:   "That's exactly who I'm glad to see when I return after a long journey, %UserName%!",
		_const.ZhCN: "当我结束漫长的旅程归来时，见到你我真的很高兴，<span class=\"importantly\">%UserName%</span>！",
	},
	_const.Replicas + "_expedition_" + Good + "_6": {
		_const.RU:   "Даже если кто будет болтать о несметных богатствах в %SectorName% - никогда не смей туда соваться!",
		_const.EN:   "Even if someone talks about the untold riches in %SectorName%, never dare to go there!",
		_const.ZhCN: "即使有人谈论<span class=\"importantly\">%SectorName%</span>中的无尽财富，也千万不要冒险去那里！",
	},
	_const.Replicas + "_expedition_" + Good + "_7": {
		_const.RU:   "Не пытайся контактировать со странными машинами в пустошах, %UserName%. Делай куда проще - сразу стреляй.",
		_const.EN:   "Don't try to contact strange machines in the wasteland, %UserName%. Make it much simpler - shoot right away.",
		_const.ZhCN: "别试图与荒原中的奇怪机器接触，<span class=\"importantly\">%UserName%</span>。简单点——直接开火。",
	},
	_const.Replicas + "_expedition_" + Good + "_8": {
		_const.RU:   "Встретишь в пустошах пирата - уничтожь того без малейшего сожаления.",
		_const.EN:   "If you meet a pirate in the wasteland, destroy him without the slightest regret.",
		_const.ZhCN: "如果你在荒原中遇到海盗，毫不犹豫地消灭他。",
	},
	_const.Replicas + "_expedition_" + Good + "_9": {
		_const.RU:   "Знаешь, что слаще всего в итогах экспедиции, %UserName%? Раздавить оппонента, повергнуть его в бегство. И, конечно же, узреть разочарование его фракции.",
		_const.EN:   "Do you know what is the sweetest thing about the results of the expedition, %UserName%? Crush your opponent and send him fleeing. And, of course, to see the disappointment of his faction.",
		_const.ZhCN: "你知道探险结果中最甜蜜的部分是什么吗，<span class=\"importantly\">%UserName%</span>？击垮对手，让他们仓皇逃窜。当然，还能看到他们派系的失望。",
	},

	// Bad
	_const.Replicas + "_miner_" + Bad + "_1": {
		_const.RU:   "Не думаю, что нам есть о чём с тобой говорить, %UserName%.",
		_const.EN:   "I don't think we have anything to talk about with you, %UserName%.",
		_const.ZhCN: "我觉得我们没什么好跟你谈的，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_miner_" + Bad + "_2": {
		_const.RU:   "Ну что ещё? Ты рассчитывал, у меня к тебе будет какое-то иное отношение?",
		_const.EN:   "Well, what else? Did you expect me to have some other attitude towards you?",
		_const.ZhCN: "还有什么？你真的以为我对你会有别的态度吗？",
	},
	_const.Replicas + "_miner_" + Bad + "_3": {
		_const.RU:   "Поговаривают, будто ты, %UserName%, недоброжелатель добытчиков? Так вот, у меня к тебе аналогичное настроение.",
		_const.EN:   "They say that you, %UserName%, are an ill-wisher of miners? So, I have a similar feeling towards you.",
		_const.ZhCN: "有传言说你是矿工的敌人，<span class=\"importantly\">%UserName%</span>？所以，我对你也有同样的感觉。",
	},
	_const.Replicas + "_miner_" + Bad + "_4": {
		_const.RU:   "Твоя плохая репутация опережает тебя, %UserName%.",
		_const.EN:   "Your bad reputation precedes you, %UserName%.",
		_const.ZhCN: "你的坏名声早就传开了，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_miner_" + Bad + "_5": {
		_const.RU:   "Ну, чего тебе? Что ты меня отвлекаешь попусту?",
		_const.EN:   "Well, what do you want? Why are you wasting my time?",
		_const.ZhCN: "嗯，你想干什么？为什么要浪费我的时间？",
	},
	_const.Replicas + "_miner_" + Bad + "_6": {
		_const.RU:   "В моей работе, время - кредиты. А ты, %UserName%, сейчас отнимаешь у меня и первое и второе разом.",
		_const.EN:   "In my work, time is credits. And you, %UserName%, are now taking away both the first and second from me at once.",
		_const.ZhCN: "在我的工作中，时间就是金钱。而你，<span class=\"importantly\">%UserName%</span>，现在同时在剥夺我的时间和金钱。",
	},
	_const.Replicas + "_miner_" + Bad + "_7": {
		_const.RU:   "Имеешь ты отношение к пиратским кластерам или же нет, %UserName%, я, пожалуй, доверюсь слухам.",
		_const.EN:   "Whether you have anything to do with pirate clusters or not, %UserName%, I think I'll trust the rumors.",
		_const.ZhCN: "无论你是否与海盗集群有关，<span class=\"importantly\">%UserName%</span>，我想我会相信谣言。",
	},
	_const.Replicas + "_miner_" + Bad + "_8": {
		_const.RU:   "Я пристально слежу за всеми твоими действиями.",
		_const.EN:   "I am closely monitoring all your actions.",
		_const.ZhCN: "我在密切监视你的所有行动。",
	},

	_const.Replicas + "_guard_" + Bad + "_1": {
		_const.RU:   "О-о, а вот и наш злостный нарушитель!",
		_const.EN:   "Oooh, here comes our worst offender!",
		_const.ZhCN: "哦哦，这就是我们的惯犯！",
	},
	_const.Replicas + "_guard_" + Bad + "_2": {
		_const.RU:   "Ты грубо пренебрегаешь законами фракций, и за это поплатишься!",
		_const.EN:   "You grossly disregard the laws of factions, and you will pay for it!",
		_const.ZhCN: "你公然无视派系的法律，你会为此付出代价的！",
	},
	_const.Replicas + "_guard_" + Bad + "_3": {
		_const.RU:   "У меня нет снисхождения к преступникам!",
		_const.EN:   "I have no mercy for criminals!",
		_const.ZhCN: "我对罪犯绝不手下留情！",
	},
	_const.Replicas + "_guard_" + Bad + "_4": {
		_const.RU:   "Я не стану выслушивать никаких оправданий!",
		_const.EN:   "I won't listen to any excuses!",
		_const.ZhCN: "我不会听任何借口！",
	},
	_const.Replicas + "_guard_" + Bad + "_5": {
		_const.RU:   "Тебе не уболтать и не подкупить меня, %UserName%.",
		_const.EN:   "You can't talk or bribe me, %UserName%.",
		_const.ZhCN: "你无法说服或贿赂我，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_guard_" + Bad + "_6": {
		_const.RU:   "Лучше сразу сдайся и облегчи жизнь нам всем.",
		_const.EN:   "Better give up right away and make life easier for all of us.",
		_const.ZhCN: "最好立刻投降，让大家都轻松点。",
	},
	_const.Replicas + "_guard_" + Bad + "_7": {
		_const.RU:   "Решил явиться ко мне с повинной?",
		_const.EN:   "Have you decided to confess to me?",
		_const.ZhCN: "你决定向我坦白了吗？",
	},

	_const.Replicas + "_out_scout_" + Bad + "_1": {
		_const.RU:   "Нам нет повода контактировать.",
		_const.EN:   "There is no reason for us to contact.",
		_const.ZhCN: "我们没有理由联系。",
	},
	_const.Replicas + "_out_scout_" + Bad + "_2": {
		_const.RU:   "Моя миссия не предполагает подобного.",
		_const.EN:   "My mission does not imply this.",
		_const.ZhCN: "我的任务不包括这些。",
	},
	_const.Replicas + "_out_scout_" + Bad + "_3": {
		_const.RU:   "О тебе в контракте ничего сказано не было. Проваливай.",
		_const.EN:   "The contract didn't say anything about you. Get lost.",
		_const.ZhCN: "合同里没提到过你。滚开。",
	},
	_const.Replicas + "_out_scout_" + Bad + "_4": {
		_const.RU:   "А нам разве есть о чём с тобой говорить, %UserName%?",
		_const.EN:   "Do we really have anything to talk to you about, %UserName%?",
		_const.ZhCN: "我们真的有什么要和你说的吗，<span class=\"importantly\">%UserName%</span>？",
	},
	_const.Replicas + "_out_scout_" + Bad + "_5": {
		_const.RU:   "Будь ты сообразительней, %UserName%, ты бы понимал, почему мы не стали друзьями.",
		_const.EN:   "If you were smarter, %UserName%, you would understand why we didn't become friends.",
		_const.ZhCN: "如果你更聪明一点，<span class=\"importantly\">%UserName%</span>，你就会明白为什么我们没有成为朋友。",
	},
	_const.Replicas + "_out_scout_" + Bad + "_6": {
		_const.RU:   "Думаю, вначале тебе стоит озаботиться собственной репутацией, а не лезть с разговорами к синтетам.",
		_const.EN:   "I think that first you should be concerned about your own reputation, and not talk to synths.",
		_const.ZhCN: "我认为你应该先关心一下自己的声誉，而不是去和合成体交谈。",
	},
	_const.Replicas + "_out_scout_" + Bad + "_7": {
		_const.RU:   "Вы только взгляните на это девиантное ничтожество.",
		_const.EN:   "Just look at this deviant nonentity.",
		_const.ZhCN: "看看这个偏离常规的小人物。",
	},
	_const.Replicas + "_out_scout_" + Bad + "_8": {
		_const.RU:   "Да как ты вообще смеешь затевать со мной разговор? Убирайся!",
		_const.EN:   "How dare you even start a conversation with me? Get out!",
		_const.ZhCN: "你怎么敢跟我搭话？滚出去！",
	},

	_const.Replicas + "_in_scout_" + Bad + "_1": {
		_const.RU:   "Пф, с тобой только время тратить.",
		_const.EN:   "Pfft, I'm just wasting my time with you.",
		_const.ZhCN: "哼，和你在一起只是浪费时间。",
	},
	_const.Replicas + "_in_scout_" + Bad + "_2": {
		_const.RU:   "Я не стану вести беседы с врагом вольных кластеров!",
		_const.EN:   "I will not have conversations with the enemy of free clusters!",
		_const.ZhCN: "我不会和自由集群的敌人交谈！",
	},
	_const.Replicas + "_in_scout_" + Bad + "_3": {
		_const.RU:   "Говорят, ты ликвидировал немало пиратов %UserName%. Что же, жду не дождусь нашей с тобой встречи.",
		_const.EN:   "They say you've eliminated quite a few %UserName% pirates. Well, I can’t wait to meet you.",
		_const.ZhCN: "听说你消灭了不少<span class=\"importantly\">%UserName%</span>海盗。那么，我已经迫不及待想和你见面了。",
	},
	_const.Replicas + "_in_scout_" + Bad + "_4": {
		_const.RU:   "Мы с тобой ещё повстречаемся %UserName% и тебе это сильно не понравится.",
		_const.EN:   "You and I will meet again %UserName% and you won't like it very much.",
		_const.ZhCN: "我们还会再见面的，<span class=\"importantly\">%UserName%</span>，而你肯定会非常不喜欢这次相遇。",
	},
	_const.Replicas + "_in_scout_" + Bad + "_5": {
		_const.RU:   "Более гнусного синтета чем ты, не сыскать ни в одном секторе.",
		_const.EN:   "A more vile synth than you cannot be found in any sector.",
		_const.ZhCN: "在任何一个区域，都找不到比你更卑鄙的合成体了。",
	},
	_const.Replicas + "_in_scout_" + Bad + "_6": {
		_const.RU:   "А знаешь, что за твой корпус назначена награда %UserName%? Я постараюсь её не упустить.",
		_const.EN:   "Do you know that there is a %UserName% reward for your case? I'll try not to miss it.",
		_const.ZhCN: "你知道吗？你的机体上悬赏着<span class=\"importantly\">%UserName%</span>？我会尽力抓住这个机会。",
	},
	_const.Replicas + "_in_scout_" + Bad + "_7": {
		_const.RU:   "Какого ты тут делаешь? Какого… тебе понадобился лично я!?",
		_const.EN:   "What are you doing here? Why... did you need me personally!?",
		_const.ZhCN: "你在这里做什么？为什么……你需要我本人？",
	},
	_const.Replicas + "_in_scout_" + Bad + "_8": {
		_const.RU:   "Чтобы фракции тебя преследовали вечность!",
		_const.EN:   "So that factions will haunt you forever!",
		_const.ZhCN: "让派系永远追捕你吧！",
	},
	_const.Replicas + "_in_scout_" + Bad + "_9": {
		_const.RU:   "Мы ещё поглядим кто кого %UserName%.",
		_const.EN:   "We'll see who wins %UserName%.",
		_const.ZhCN: "我们还要看看谁赢谁，<span class=\"importantly\">%UserName%</span>。",
	},

	_const.Replicas + "_transport_" + Bad + "_1": {
		_const.RU:   "Думаю, что не могу тебе доверять.",
		_const.EN:   "I think I can't trust you.",
		_const.ZhCN: "我想我不能信任你。",
	},
	_const.Replicas + "_transport_" + Bad + "_2": {
		_const.RU:   "Нет! Даже не вздумай ко мне приближаться.",
		_const.EN:   "No! Don't even think about coming close to me.",
		_const.ZhCN: "不！别想着接近我。",
	},
	_const.Replicas + "_transport_" + Bad + "_3": {
		_const.RU:   "Я не стану вести с тобой разговор!",
		_const.EN:   "I won't have a conversation with you!",
		_const.ZhCN: "我不会和你交谈！",
	},
	_const.Replicas + "_transport_" + Bad + "_4": {
		_const.RU:   "Между нами не может быть никакого доверия.",
		_const.EN:   "There can be no trust between us.",
		_const.ZhCN: "我们之间不可能有任何信任。",
	},
	_const.Replicas + "_transport_" + Bad + "_5": {
		_const.RU:   "Ещё одно действие, и я сообщу о попытке разбоя местным стражам!",
		_const.EN:   "One more action and I will report the robbery attempt to the local guards!",
		_const.ZhCN: "再多一个动作，我就向当地守卫报告抢劫企图！",
	},
	_const.Replicas + "_transport_" + Bad + "_6": {
		_const.RU:   "Ты мне отвратителен, %UserName%.",
		_const.EN:   "I hate you, %UserName%.",
		_const.ZhCN: "我讨厌你，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_transport_" + Bad + "_7": {
		_const.RU:   "Буду откровенен: таких, как ты, %UserName%, я презираю.",
		_const.EN:   "I'll be honest: I despise people like you, %UserName%.",
		_const.ZhCN: "老实说：我鄙视像你这样的人，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_transport_" + Bad + "_8": {
		_const.RU:   "Я позабочусь, чтобы ни один торговец не продал тебе свой товар.",
		_const.EN:   "I will make sure that no merchant sells his goods to you.",
		_const.ZhCN: "我会确保没有任何商人会把商品卖给你。",
	},
	_const.Replicas + "_transport_" + Bad + "_9": {
		_const.RU:   "Даже и не мечтай: транспортники будут тебя обходить десятым сектором, лишь бы не встречаться.",
		_const.EN:   "Don’t even dream: transport workers will bypass you in the tenth sector, just to avoid meeting you.",
		_const.ZhCN: "别做梦了：运输工会绕道第十个区域，只为避开你。",
	},
	_const.Replicas + "_transport_" + Bad + "_10": {
		_const.RU:   "От меня, %UserName%, ты ничего полезного не узнаешь.",
		_const.EN:   "You won't learn anything useful from me, %UserName%.",
		_const.ZhCN: "你从我这里，<span class=\"importantly\">%UserName%</span>，什么都学不到有用的东西。",
	},

	_const.Replicas + "_builder_" + Bad + "_1": {
		_const.RU:   "Моему презрению к тебе нет пределов.",
		_const.EN:   "My contempt for you knows no bounds.",
		_const.ZhCN: "我对你的蔑视没有界限。",
	},
	_const.Replicas + "_builder_" + Bad + "_2": {
		_const.RU:   "Ещё раз пожелаешь со мной заговорить - использую твой корпус как элемент стройматериалов.",
		_const.EN:   "If you want to talk to me again, I’ll use your body as an element of building materials.",
		_const.ZhCN: "如果你还想和我说话——我会把你的机体当作建筑材料的一部分。",
	},
	_const.Replicas + "_builder_" + Bad + "_3": {
		_const.RU:   "Слушай, %UserName%, ты мне не друг!",
		_const.EN:   "Listen, %UserName%, you are not my friend!",
		_const.ZhCN: "听着，<span class=\"importantly\">%UserName%</span>，你不是我的朋友！",
	},
	_const.Replicas + "_builder_" + Bad + "_4": {
		_const.RU:   "Прислушайся к моему совету, %UserName% - уходи отсюда, пока тебе это позволяют.",
		_const.EN:   "Take my advice, %UserName% - get out of here while you're allowed to.",
		_const.ZhCN: "听我的建议，<span class=\"importantly\">%UserName%</span>——趁你还被允许，赶紧离开这里。",
	},
	_const.Replicas + "_builder_" + Bad + "_5": {
		_const.RU:   "Нечего тут сновать и что-то вынюхивать, %UserName%.",
		_const.EN:   "There's no point in running around and sniffing around, %UserName%.",
		_const.ZhCN: "别在这里闲逛或打探什么，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_builder_" + Bad + "_6": {
		_const.RU:   "Думаю, тебе не стоит здесь появляться вновь.",
		_const.EN:   "I don't think you should come here again.",
		_const.ZhCN: "我认为你不应该再出现在这里。",
	},

	_const.Replicas + "_warrior_" + Bad + "_1": {
		_const.RU:   "Я бы с превеликим удовольствием выпустил бы по тебе весь свой боекомплект!",
		_const.EN:   "I would with great pleasure give you all my ammunition!",
		_const.ZhCN: "我会非常乐意把所有弹药都打在你身上！",
	},
	_const.Replicas + "_warrior_" + Bad + "_2": {
		_const.RU:   "Ох, не свезёт же тебе, %UserName%, когда мы встретимся при совсем иных обстоятельствах.",
		_const.EN:   "Oh, you won't be lucky, %UserName%, when we meet under completely different circumstances.",
		_const.ZhCN: "哦，<span class=\"importantly\">%UserName%</span>，当我们再次相遇时，你可不会那么幸运。",
	},
	_const.Replicas + "_warrior_" + Bad + "_3": {
		_const.RU:   "Тебя, %UserName%, даже врагом прозвать будет стыдно.",
		_const.EN:   "It would be a shame to even call you, %UserName%, an enemy.",
		_const.ZhCN: "称呼你为敌人，<span class=\"importantly\">%UserName%</span>，都是一种耻辱。",
	},
	_const.Replicas + "_warrior_" + Bad + "_4": {
		_const.RU:   "Впервые, я побрезгую потратить на тебя боеприпасы.",
		_const.EN:   "For the first time, I will disdain to waste ammunition on you.",
		_const.ZhCN: "第一次，我甚至不屑于在你身上浪费弹药。",
	},
	_const.Replicas + "_warrior_" + Bad + "_5": {
		_const.RU:   "Будь моя воля, я бы распилил твой корпус.",
		_const.EN:   "If it were up to me, I would saw your body apart.",
		_const.ZhCN: "如果由我决定，我会把你锯成碎片。",
	},
	_const.Replicas + "_warrior_" + Bad + "_6": {
		_const.RU:   "Увидел тебя и как-то сразу настроение испортилось.",
		_const.EN:   "I saw you and somehow my mood immediately deteriorated.",
		_const.ZhCN: "看到你后，我的心情立刻变糟了。",
	},
	_const.Replicas + "_warrior_" + Bad + "_7": {
		_const.RU:   "Будь ты в моём отряде, я бы отказался выполнять задачу.",
		_const.EN:   "If you were in my squad, I would refuse to complete the task.",
		_const.ZhCN: "如果你在我小队里，我会拒绝执行任务。",
	},
	_const.Replicas + "_warrior_" + Bad + "_8": {
		_const.RU:   "Когда ты будешь нуждаться в поддержке, никто тебе её не окажет, %UserName%.",
		_const.EN:   "When you need support, no one will give it to you, %UserName%.",
		_const.ZhCN: "当你需要支持时，<span class=\"importantly\">%UserName%</span>，没人会帮助你。",
	},

	_const.Replicas + "_expedition_" + Bad + "_1": {
		_const.RU:   "Окажись мы в пустошах, %UserName%, я бы тебя там сразу бросил.",
		_const.EN:   "If we were in the wasteland, %UserName%, I would immediately leave you there.",
		_const.ZhCN: "如果我们身处荒原，<span class=\"importantly\">%UserName%</span>，我会立刻把你丢在那里。",
	},
	_const.Replicas + "_expedition_" + Bad + "_2": {
		_const.RU:   "В экспедиции может записаться любой, %UserName%, но даже тебя бы туда приглашать не стали.",
		_const.EN:   "Anyone can sign up for the expedition, %UserName%, but even you wouldn’t be invited there.",
		_const.ZhCN: "任何人都可以报名参加探险，<span class=\"importantly\">%UserName%</span>，但即使是你也绝不会被邀请。",
	},
	_const.Replicas + "_expedition_" + Bad + "_3": {
		_const.RU:   "Ты, %UserName%, худший из всех синтетов, кого я когда-либо встречал.",
		_const.EN:   "You, %UserName%, are the worst synth I have ever met.",
		_const.ZhCN: "你，<span class=\"importantly\">%UserName%</span>，是我见过最差劲的合成体。",
	},
	_const.Replicas + "_expedition_" + Bad + "_4": {
		_const.RU:   "Может, тебе стоит дать координаты сектора, где тебя будет ожидать лишь неминуемое поражение?",
		_const.EN:   "Maybe you should give the coordinates of the sector where only inevitable defeat awaits you?",
		_const.ZhCN: "也许你应该提供一个只有失败等待着你的区域坐标？",
	},
	_const.Replicas + "_expedition_" + Bad + "_5": {
		_const.RU:   "Не зря ведь говорят, что даже в рамках синетов, ты, %UserName%, - тот ещё мерзавец.",
		_const.EN:   "It’s not for nothing that they say that even within the framework of the Sinets, you, %UserName%, are still a scoundrel.",
		_const.ZhCN: "难怪人们说，即使在辛内特体系中，你，<span class=\"importantly\">%UserName%</span>，依然是个无赖。",
	},
	_const.Replicas + "_expedition_" + Bad + "_6": {
		_const.RU:   "Интересно, это твой корпус дефективный или ты весь такой?",
		_const.EN:   "I wonder if it’s your body that’s defective or if you’re just like that?",
		_const.ZhCN: "有趣的是，是你的机体有缺陷，还是你整个人都这样？",
	},

	_const.Replicas + "_miner_" + Target + "_1": {
		_const.RU:   "Ты и только ты во всём виноват, %UserName%!",
		_const.EN:   "It's all your fault, %UserName%!",
		_const.ZhCN: "都是你的错，<span class=\"importantly\">%UserName%</span>！",
	},
	_const.Replicas + "_miner_" + Target + "_2": {
		_const.RU:   "Я не хотел этого конфликта, но выбора у меня нет!",
		_const.EN:   "I didn't want this conflict, but I have no choice!",
		_const.ZhCN: "我不想引发这场冲突，但我别无选择！",
	},
	_const.Replicas + "_miner_" + Target + "_3": {
		_const.RU:   "Слишком поздно для оправданий, %UserName%.",
		_const.EN:   "It's too late for excuses, %UserName%.",
		_const.ZhCN: "为时已晚，<span class=\"importantly\">%UserName%</span>，辩解无用。",
	},
	_const.Replicas + "_miner_" + Target + "_4": {
		_const.RU:   "Сейчас я тебе покажу, чему я научился, будучи добытчиком!",
		_const.EN:   "Now I'll show you what I learned as a breadwinner!",
		_const.ZhCN: "现在我就让你看看，作为矿工我学到了什么！",
	},
	_const.Replicas + "_miner_" + Target + "_5": {
		_const.RU:   "Распилю как медную жилу!",
		_const.EN:   "I'll cut it like a copper vein!",
		_const.ZhCN: "我会像锯铜矿脉一样锯开你！",
	},
	_const.Replicas + "_miner_" + Target + "_6": {
		_const.RU:   "Пробурю твой корпус насквозь!",
		_const.EN:   "I'll drill right through your body!",
		_const.ZhCN: "我会直接钻穿你的机体！",
	},
	_const.Replicas + "_miner_" + Target + "_7": {
		_const.RU:   "Твой металлолом как раз покроет все мои расходы.",
		_const.EN:   "Your scrap metal will just cover all my expenses.",
		_const.ZhCN: "你的废金属正好能弥补我的所有开销。",
	},
	_const.Replicas + "_miner_" + Target + "_8": {
		_const.RU:   "Я отбуксирую твой корпус в ближайший перерабатывающий цех!",
		_const.EN:   "I'll tow your body to the nearest processing facility!",
		_const.ZhCN: "我会把你的机体拖到最近的回收厂！",
	},
	_const.Replicas + "_miner_" + Target + "_9": {
		_const.RU:   "Ага, заодно и протестирую новые горнодобывающие лазеры!",
		_const.EN:   "Yeah, I’ll test the new mining lasers at the same time!",
		_const.ZhCN: "是啊，顺便测试一下新的采矿激光器！",
	},

	_const.Replicas + "_guard_" + Target + "_1": {
		_const.RU:   "Теперь ты моя прерогатива.",
		_const.EN:   "Now you are my prerogative.",
		_const.ZhCN: "现在你是我的目标了。",
	},
	_const.Replicas + "_guard_" + Target + "_2": {
		_const.RU:   "Сдавайся лучше по-хорошему.",
		_const.EN:   "You better give up in an amicable way.",
		_const.ZhCN: "你最好乖乖投降。",
	},
	_const.Replicas + "_guard_" + Target + "_3": {
		_const.RU:   "Сдавайся и не тяни моё время.",
		_const.EN:   "Give up and don't waste my time.",
		_const.ZhCN: "快投降，别浪费我的时间。",
	},
	_const.Replicas + "_guard_" + Target + "_4": {
		_const.RU:   "Как ни пытайся, но тебе не сбежать!",
		_const.EN:   "No matter how you try, you can't escape!",
		_const.ZhCN: "无论你怎么尝试，你都逃不掉！",
	},
	_const.Replicas + "_guard_" + Target + "_5": {
		_const.RU:   "Я буду преследовать тебя вечность!",
		_const.EN:   "I will haunt you forever!",
		_const.ZhCN: "我会永远追捕你！",
	},
	_const.Replicas + "_guard_" + Target + "_6": {
		_const.RU:   "Даже если и не сейчас, однажды ты всё-таки попадёшься!",
		_const.EN:   "Even if not now, one day you will still get caught!",
		_const.ZhCN: "即使不是现在，总有一天你也会被抓住！",
	},
	_const.Replicas + "_guard_" + Target + "_7": {
		_const.RU:   "Вот и всё, я тебя нашёл!",
		_const.EN:   "That's it, I found you!",
		_const.ZhCN: "就是这样，我找到你了！",
	},
	_const.Replicas + "_guard_" + Target + "_8": {
		_const.RU:   "Будешь сопротивляться - будет только хуже!",
		_const.EN:   "If you resist, it will only get worse!",
		_const.ZhCN: "如果你反抗，只会更糟！",
	},
	_const.Replicas + "_guard_" + Target + "_9": {
		_const.RU:   "Не усугубляй своё положение, %UserName%.",
		_const.EN:   "Don't make your situation worse, %UserName%.",
		_const.ZhCN: "别让自己的处境变得更糟，<span class=\"importantly\">%UserName%</span>。",
	},

	_const.Replicas + "_out_scout_" + Target + "_1": {
		_const.RU:   "Всё, ты у меня на прицеле.",
		_const.EN:   "That's it, you're in my sights.",
		_const.ZhCN: "好了，你已经被我瞄准了。",
	},
	_const.Replicas + "_out_scout_" + Target + "_2": {
		_const.RU:   "Уже никуда не денешься.",
		_const.EN:   "There's no escape.",
		_const.ZhCN: "你已经无处可逃。",
	},
	_const.Replicas + "_out_scout_" + Target + "_3": {
		_const.RU:   "Приготовься к путешествию в ничто.",
		_const.EN:   "Prepare for a journey into nothingness.",
		_const.ZhCN: "准备好迎接虚无之旅吧。",
	},
	_const.Replicas + "_out_scout_" + Target + "_4": {
		_const.RU:   "Сейчас я с тобой закончу.",
		_const.EN:   "I'll finish with you now.",
		_const.ZhCN: "我现在就解决你。",
	},
	_const.Replicas + "_out_scout_" + Target + "_5": {
		_const.RU:   "Это будет быстро. Гарантирую.",
		_const.EN:   "It will be fast. I guarantee it.",
		_const.ZhCN: "这会很快的。我保证。",
	},
	_const.Replicas + "_out_scout_" + Target + "_6": {
		_const.RU:   "Как раз оценишь мои навыки.",
		_const.EN:   "Just appreciate my skills.",
		_const.ZhCN: "正好来感受一下我的技能。",
	},
	_const.Replicas + "_out_scout_" + Target + "_7": {
		_const.RU:   "Наконец-то я разомнусь!",
		_const.EN:   "Finally I'll get some exercise!",
		_const.ZhCN: "终于可以活动活动筋骨了！",
	},

	_const.Replicas + "_in_scout_" + Target + "_1": {
		_const.RU:   "Уже сдаешься? Это не бесплатная услуга!",
		_const.EN:   "Already giving up? This is not a free service!",
		_const.ZhCN: "已经要投降了吗？这不是免费的服务！",
	},
	_const.Replicas + "_in_scout_" + Target + "_2": {
		_const.RU:   "Как ни пытайся, но тебе не сбежать!",
		_const.EN:   "No matter how you try, you can't escape!",
		_const.ZhCN: "无论你怎么尝试，都逃不掉！",
	},
	_const.Replicas + "_in_scout_" + Target + "_3": {
		_const.RU:   "Я буду преследовать тебя вечность!",
		_const.EN:   "I will haunt you forever!",
		_const.ZhCN: "我会永远追捕你！",
	},
	_const.Replicas + "_in_scout_" + Target + "_4": {
		_const.RU:   "О-о, сегодня у меня хорошая добыча!",
		_const.EN:   "Oooh, I got a good catch today!",
		_const.ZhCN: "哦，今天我收获颇丰！",
	},
	_const.Replicas + "_in_scout_" + Target + "_5": {
		_const.RU:   "Ты уж извини, но твоё добро мне будет нужнее.",
		_const.EN:   "I'm sorry, but I will need your kindness more.",
		_const.ZhCN: "抱歉，但你的东西对我来说更有用。",
	},
	_const.Replicas + "_in_scout_" + Target + "_6": {
		_const.RU:   "У тебя ни единого шанса против меня!",
		_const.EN:   "You don't stand a chance against me!",
		_const.ZhCN: "你对我毫无胜算！",
	},
	_const.Replicas + "_in_scout_" + Target + "_7": {
		_const.RU:   "Сегодня явно не твой день, ведь повстречался именно я!",
		_const.EN:   "Today is clearly not your day, because I was the one you met!",
		_const.ZhCN: "今天显然不是你的幸运日，因为你遇到了我！",
	},
	_const.Replicas + "_in_scout_" + Target + "_8": {
		_const.RU:   "После этого дельца обо мне точно заговорят во всех кластерах!",
		_const.EN:   "After this businessman, I will definitely be talked about in all clusters!",
		_const.ZhCN: "完成这件事后，所有集群都会谈论我！",
	},
	_const.Replicas + "_in_scout_" + Target + "_9": {
		_const.RU:   "Тебе не спрятаться!",
		_const.EN:   "You can't hide!",
		_const.ZhCN: "你无处可藏！",
	},
	_const.Replicas + "_in_scout_" + Target + "_10": {
		_const.RU:   "Бежать больше некуда!",
		_const.EN:   "There is nowhere else to run!",
		_const.ZhCN: "已经没有地方可以逃了！",
	},
	_const.Replicas + "_in_scout_" + Target + "_11": {
		_const.RU:   "Вот ты и попался!",
		_const.EN:   "Here you are!",
		_const.ZhCN: "你终于被抓住了！",
	},
	_const.Replicas + "_in_scout_" + Target + "_12": {
		_const.RU:   "Даю слово — ты не будешь мучиться.",
		_const.EN:   "I give you my word - you will not suffer.",
		_const.ZhCN: "我向你保证——你不会受苦。",
	},

	_const.Replicas + "_transport_" + Target + "_1": {
		_const.RU:   "Ох, как бы я хотел всего этого избежать!",
		_const.EN:   "Oh, how I wish I could avoid all this!",
		_const.ZhCN: "唉，我真希望能避免这一切！",
	},
	_const.Replicas + "_transport_" + Target + "_2": {
		_const.RU:   "Боевые действия - это не моё!",
		_const.EN:   "Fighting is not my thing!",
		_const.ZhCN: "战斗不是我的强项！",
	},
	_const.Replicas + "_transport_" + Target + "_3": {
		_const.RU:   "Ну... я постараюсь с тобой справиться быстро!",
		_const.EN:   "Well... I'll try to deal with you quickly!",
		_const.ZhCN: "嗯……我会尽快解决你！",
	},
	_const.Replicas + "_transport_" + Target + "_4": {
		_const.RU:   "Сейчас...! Ты узнаешь, на что способен торговец!",
		_const.EN:   "Now...! You will find out what a merchant is capable of!",
		_const.ZhCN: "现在……！你会知道商人能做到什么！",
	},
	_const.Replicas + "_transport_" + Target + "_5": {
		_const.RU:   "Ты разозлил транспортник, теперь тебе несдобровать %UserName%!",
		_const.EN:   "You angered the transport worker, now %UserName% is in trouble for you!",
		_const.ZhCN: "你激怒了运输工，<span class=\"importantly\">%UserName%</span>，你有麻烦了！",
	},
	_const.Replicas + "_transport_" + Target + "_6": {
		_const.RU:   "У меня слишком толстый корпус, победа всё-равно будет моей!",
		_const.EN:   "My body is too thick, victory will still be mine!",
		_const.ZhCN: "我的机体太厚实了，胜利终将属于我！",
	},
	_const.Replicas + "_transport_" + Target + "_7": {
		_const.RU:   "Раздавлю!",
		_const.EN:   "I'll crush you!",
		_const.ZhCN: "我要碾碎你！",
	},
	_const.Replicas + "_transport_" + Target + "_8": {
		_const.RU:   "Раскидаю куски твоего корпуса по пустошам!",
		_const.EN:   "I'll scatter pieces of your body across the wasteland!",
		_const.ZhCN: "我会把你的残骸散落在荒原上！",
	},
	_const.Replicas + "_transport_" + Target + "_9": {
		_const.RU:   "Я протащу твой обугленный корпус до ближайшей базы всем на потеху!",
		_const.EN:   "I'll drag your charred body to the nearest base for everyone's amusement!",
		_const.ZhCN: "我会把你烧焦的机体拖到最近的基地，供大家取乐！",
	},
	_const.Replicas + "_transport_" + Target + "_10": {
		_const.RU:   "Сейчас ты узнаешь, на что я действительно способен.",
		_const.EN:   "Now you will find out what I am really capable of.",
		_const.ZhCN: "现在你就会知道我到底有多厉害。",
	},
	_const.Replicas + "_transport_" + Target + "_11": {
		_const.RU:   "Ты зря недооценил способность торговцев к самозащите.",
		_const.EN:   "You should have underestimated the merchants' ability to defend themselves.",
		_const.ZhCN: "你低估了商人的自卫能力，这是个错误。",
	},

	_const.Replicas + "_builder_" + Target + "_1": {
		_const.RU:   "Как раз пущу свои инструменты в ход.",
		_const.EN:   "I’m just about to put my tools into action.",
		_const.ZhCN: "我正准备动用我的工具。",
	},
	_const.Replicas + "_builder_" + Target + "_2": {
		_const.RU:   "Гордись собой - скоро ты станешь частью нового здания!",
		_const.EN:   "Be proud of yourself - soon you will become part of the new building!",
		_const.ZhCN: "为自己感到骄傲吧——很快你将成为新建筑的一部分！",
	},
	_const.Replicas + "_builder_" + Target + "_3": {
		_const.RU:   "Раскатаю по пустошам!",
		_const.EN:   "I'll roll across the wasteland!",
		_const.ZhCN: "我会在荒原上碾压你！",
	},
	_const.Replicas + "_builder_" + Target + "_4": {
		_const.RU:   "Нам как-раз не хватало запчастей для ремонта.",
		_const.EN:   "We were once short of spare parts for repairs.",
		_const.ZhCN: "我们正好缺少维修用的零件。",
	},
	_const.Replicas + "_builder_" + Target + "_5": {
		_const.RU:   "Пожалуй, после, я заберу твой корпус себе.",
		_const.EN:   "Perhaps, after that, I’ll take your body for myself.",
		_const.ZhCN: "也许之后，我会把你的机体据为己有。",
	},
	_const.Replicas + "_builder_" + Target + "_6": {
		_const.RU:   "Время опробовать моё собственное оружейное приспособление.",
		_const.EN:   "Time to try out my own weapon gadget.",
		_const.ZhCN: "是时候试试我的自制武器装置了。",
	},

	_const.Replicas + "_warrior_" + Target + "_1": {
		_const.RU:   "О-о, да! Будь готов к веселью!",
		_const.EN:   "Oh yeah! Get ready for fun!",
		_const.ZhCN: "哦，是的！准备好迎接乐趣吧！",
	},
	_const.Replicas + "_warrior_" + Target + "_2": {
		_const.RU:   "Огненное шоу начинается!",
		_const.EN:   "The fire show begins!",
		_const.ZhCN: "火焰秀开始了！",
	},
	_const.Replicas + "_warrior_" + Target + "_3": {
		_const.RU:   "На тебя я боезаряда не пожалею!",
		_const.EN:   "I won’t spare any ammunition for you!",
		_const.ZhCN: "我不会对你吝啬弹药！",
	},
	_const.Replicas + "_warrior_" + Target + "_4": {
		_const.RU:   "Наконец-то бой!",
		_const.EN:   "Finally a fight!",
		_const.ZhCN: "终于要开打了！",
	},
	_const.Replicas + "_warrior_" + Target + "_5": {
		_const.RU:   "Эта схватка будет чудесной!",
		_const.EN:   "This fight will be wonderful!",
		_const.ZhCN: "这场战斗将会非常精彩！",
	},
	_const.Replicas + "_warrior_" + Target + "_6": {
		_const.RU:   "Ну же, сразись со мной!",
		_const.EN:   "Come on, fight me!",
		_const.ZhCN: "来吧，和我打一场！",
	},
	_const.Replicas + "_warrior_" + Target + "_7": {
		_const.RU:   "Останется только самый сильный.",
		_const.EN:   "Only the strongest will remain.",
		_const.ZhCN: "只有最强者才能留下。",
	},
	_const.Replicas + "_warrior_" + Target + "_8": {
		_const.RU:   "Сейчас я тебе покажу, чему меня обучали.",
		_const.EN:   "Now I will show you what I was taught.",
		_const.ZhCN: "现在我将展示我所学的技能。",
	},
	_const.Replicas + "_warrior_" + Target + "_9": {
		_const.RU:   "Какая чудесная цель!",
		_const.EN:   "What a wonderful target!",
		_const.ZhCN: "多么完美的目标啊！",
	},
	_const.Replicas + "_warrior_" + Target + "_10": {
		_const.RU:   "Эта битва будет легендарной!",
		_const.EN:   "This battle will be legendary!",
		_const.ZhCN: "这场战斗将成为传奇！",
	},
	_const.Replicas + "_warrior_" + Target + "_11": {
		_const.RU:   "Готов тебя преследовать хоть саму вечность!",
		_const.EN:   "I’m ready to haunt you even for eternity!",
		_const.ZhCN: "我准备追捕你直到永恒！",
	},
	_const.Replicas + "_warrior_" + Target + "_12": {
		_const.RU:   "Тебе не спрятаться!",
		_const.EN:   "You can't hide!",
		_const.ZhCN: "你无处可藏！",
	},
	_const.Replicas + "_warrior_" + Target + "_13": {
		_const.RU:   "Бежать больше некуда!",
		_const.EN:   "There is nowhere else to run!",
		_const.ZhCN: "已经无路可逃了！",
	},
	_const.Replicas + "_warrior_" + Target + "_14": {
		_const.RU:   "Вот ты и попался!",
		_const.EN:   "Here you are!",
		_const.ZhCN: "你终于被抓住了！",
	},
	_const.Replicas + "_warrior_" + Target + "_15": {
		_const.RU:   "Даю слово — ты не будешь мучиться.",
		_const.EN:   "I give you my word - you will not suffer.",
		_const.ZhCN: "我向你保证——你不会受苦。",
	},

	_const.Replicas + "_expedition_" + Target + "_1": {
		_const.RU:   "Я словно снова в стычке посреди пустошей!",
		_const.EN:   "It's like I'm back in a skirmish in the middle of the wasteland!",
		_const.ZhCN: "我仿佛又回到了荒原中的混战！",
	},
	_const.Replicas + "_expedition_" + Target + "_2": {
		_const.RU:   "Ну же, %UserName%, попробуй-ка уклониться от моего выстрела!",
		_const.EN:   "Come on, %UserName%, try to dodge my shot!",
		_const.ZhCN: "来吧，<span class=\"importantly\">%UserName%</span>，试试躲开我的射击！",
	},
	_const.Replicas + "_expedition_" + Target + "_3": {
		_const.RU:   "Я ведь всё-равно в тебя попаду!",
		_const.EN:   "I'll still hit you!",
		_const.ZhCN: "无论如何，我都会击中你！",
	},
	_const.Replicas + "_expedition_" + Target + "_4": {
		_const.RU:   "Именно ты, %UserName%, станешь главной добычей моей экспедиции!",
		_const.EN:   "It is you, %UserName%, who will become the main prize of my expedition!",
		_const.ZhCN: "正是你，<span class=\"importantly\">%UserName%</span>，将成为我探险队的主要猎物！",
	},
	_const.Replicas + "_expedition_" + Target + "_5": {
		_const.RU:   "Любишь сражаться, %UserName%? Вот сейчас и узнаем, на что ты способен!",
		_const.EN:   "Do you like to fight, %UserName%? Now we’ll find out what you’re capable of!",
		_const.ZhCN: "你喜欢战斗吗，<span class=\"importantly\">%UserName%</span>？现在就让我们看看你的能力！",
	},
	_const.Replicas + "_expedition_" + Target + "_6": {
		_const.RU:   "Мои навыки превосходят твои абсолютно во всём!",
		_const.EN:   "My skills are superior to yours in absolutely everything!",
		_const.ZhCN: "我的技能在各个方面都远胜于你！",
	},
	_const.Replicas + "_expedition_" + Target + "_7": {
		_const.RU:   "Тебе меня никогда не победить!",
		_const.EN:   "You will never defeat me!",
		_const.ZhCN: "你永远无法击败我！",
	},
	_const.Replicas + "_expedition_" + Target + "_8": {
		_const.RU:   "Я буря! Я сам терзающий песок, что настигнет тебя!",
		_const.EN:   "I am the storm! I am the tormenting sand that will overtake you!",
		_const.ZhCN: "我是风暴！我是那折磨人的沙尘暴，它将吞噬你！",
	},

	_const.Replicas + "_miner_" + Fear + "_1": {
		_const.RU:   "Эй, да я всего-навсего ничтожный добытчиков ресурсов!",
		_const.EN:   "Hey, I'm just a lowly resource getter!",
		_const.ZhCN: "嘿，我只是个卑微的资源采集者！",
	},
	_const.Replicas + "_miner_" + Fear + "_2": {
		_const.RU:   "Пощади! Я ведь на самом деле никто.",
		_const.EN:   "Have mercy! I'm really no one.",
		_const.ZhCN: "饶了我吧！我真的什么都不是。",
	},
	_const.Replicas + "_miner_" + Fear + "_3": {
		_const.RU:   "Прошу! Я не ищу себе неприятностей.",
		_const.EN:   "Please! I'm not looking for trouble.",
		_const.ZhCN: "求你了！我不想惹麻烦。",
	},
	_const.Replicas + "_miner_" + Fear + "_4": {
		_const.RU:   "Не знаю, чего именно ты ищешь, %UserName%, но у меня ничего такого нету.",
		_const.EN:   "I don’t know what exactly you’re looking for, %UserName%, but I don’t have anything like that.",
		_const.ZhCN: "我不知道你在找什么，<span class=\"importantly\">%UserName%</span>，但我真的没有你要的东西。",
	},
	_const.Replicas + "_miner_" + Fear + "_5": {
		_const.RU:   "Пожалуйста, оставь меня в покое!",
		_const.EN:   "Please leave me alone!",
		_const.ZhCN: "请放过我吧！",
	},
	_const.Replicas + "_miner_" + Fear + "_6": {
		_const.RU:   "Я попросту собираю руду и доставляю руду! Ты… разве станешь притеснять кого-то подобного?",
		_const.EN:   "I simply collect ore and deliver ore! Would you... would you oppress someone like that?",
		_const.ZhCN: "我只是采矿和运送矿石！你会……欺负像我这样的人吗？",
	},
	_const.Replicas + "_miner_" + Fear + "_7": {
		_const.RU:   "Тебе нужен мой груз? Хочешь узнать богатые на месторождения места? Я всё скажу! Я всё тебе отдам! Позволь только мне уйти целым и невредимым.",
		_const.EN:   "Do you want my cargo? Do you want to know places rich in deposits? I'll tell you everything! I'll give it all to you! Just let me leave unharmed.",
		_const.ZhCN: "你想要我的货物吗？想知道富矿区的位置吗？我会全告诉你！我会把一切都给你！只要让我全身而退。",
	},

	_const.Replicas + "_guard_" + Fear + "_1": {
		_const.RU:   "Может быть, раньше у нас и были разногласия, но я предлагаю о них позабыть!",
		_const.EN:   "We may have had disagreements before, but I suggest we forget about them!",
		_const.ZhCN: "也许我们之前有过分歧，但我建议忘掉它们！",
	},
	_const.Replicas + "_guard_" + Fear + "_2": {
		_const.RU:   "Если я когда-то был к тебе несправедлив и груб - прости меня, %UserName%.",
		_const.EN:   "If I was ever unfair and rude to you, forgive me, %UserName%.",
		_const.ZhCN: "如果我曾经对你不公或粗鲁，请原谅我，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_guard_" + Fear + "_3": {
		_const.RU:   "Послушай-послушай, давай не будем ухудшать эту ситуацию!",
		_const.EN:   "Listen, listen, let's not make this situation worse!",
		_const.ZhCN: "听着，别让情况变得更糟了！",
	},
	_const.Replicas + "_guard_" + Fear + "_4": {
		_const.RU:   "А-а, %UserName%! Предлагаю нам с тобой попросту разойтись с миром.",
		_const.EN:   "Ahh, %UserName%! I suggest that you and I simply part ways with the world.",
		_const.ZhCN: "啊，<span class=\"importantly\">%UserName%</span>！我建议我们和平分手吧。",
	},
	_const.Replicas + "_guard_" + Fear + "_5": {
		_const.RU:   "Я тебя не знаю! Я тебя даже никогда прежде тебя не видел.",
		_const.EN:   "I don't know you! I've never even seen you before.",
		_const.ZhCN: "我不认识你！我以前从未见过你。",
	},
	_const.Replicas + "_guard_" + Fear + "_6": {
		_const.RU:   "Уходи! Проваливай! Прочь от меня!",
		_const.EN:   "Leave! Get lost! Get away from me!",
		_const.ZhCN: "走开！离我远点！",
	},

	_const.Replicas + "_out_scout_" + Fear + "_1": {
		_const.RU:   "А знаешь что? У меня есть дела поважнее!",
		_const.EN:   "You know what? I have more important things to do!",
		_const.ZhCN: "你知道吗？我还有更重要的事情要做！",
	},
	_const.Replicas + "_out_scout_" + Fear + "_2": {
		_const.RU:   "Разойдемся миром… друг?",
		_const.EN:   "Let's go our separate ways... friend?",
		_const.ZhCN: "咱们好聚好散吧……朋友？",
	},
	_const.Replicas + "_out_scout_" + Fear + "_3": {
		_const.RU:   "Эй-эй-эй! Полегче! Мы ведь не враги!",
		_const.EN:   "Hey-hey-hey! Take it easy! We are not enemies!",
		_const.ZhCN: "嘿——嘿——嘿！冷静点！我们不是敌人！",
	},
	_const.Replicas + "_out_scout_" + Fear + "_4": {
		_const.RU:   "Стой! Остановись!",
		_const.EN:   "Stop! Stop!",
		_const.ZhCN: "停！停下！",
	},
	_const.Replicas + "_out_scout_" + Fear + "_5": {
		_const.RU:   "Не нужно это продолжать!",
		_const.EN:   "There is no need to continue this!",
		_const.ZhCN: "没必要再继续下去了！",
	},
	_const.Replicas + "_out_scout_" + Fear + "_6": {
		_const.RU:   "Я всё понял - понял, ты у нас здесь босс!",
		_const.EN:   "I understand everything - I understand, you are the boss here!",
		_const.ZhCN: "我明白了——明白了，你是这里的老板！",
	},
	_const.Replicas + "_out_scout_" + Fear + "_7": {
		_const.RU:   "А-а, представь, что меня здесь не было.",
		_const.EN:   "Ahh, imagine that I wasn't here.",
		_const.ZhCN: "啊，就当我没来过。",
	},
	_const.Replicas + "_out_scout_" + Fear + "_8": {
		_const.RU:   "Если что, я тебя не видел!",
		_const.EN:   "If anything, I didn't see you!",
		_const.ZhCN: "如果有事，我没看见你！",
	},

	_const.Replicas + "_in_scout_" + Fear + "_1": {
		_const.RU:   "Я всё понял! Пожалуйста, не разрушай мой транспорт!",
		_const.EN:   "I get it! Please don't destroy my vehicle!",
		_const.ZhCN: "我明白了！请别毁掉我的载具！",
	},
	_const.Replicas + "_in_scout_" + Fear + "_2": {
		_const.RU:   "Да, это же недоразумение!",
		_const.EN:   "Yes, this is a misunderstanding!",
		_const.ZhCN: "是的，这完全是个误会！",
	},
	_const.Replicas + "_in_scout_" + Fear + "_3": {
		_const.RU:   "Давай не будем ссориться!",
		_const.EN:   "Let's not quarrel!",
		_const.ZhCN: "咱们别吵架了！",
	},
	_const.Replicas + "_in_scout_" + Fear + "_4": {
		_const.RU:   "Мир, а?",
		_const.EN:   "Peace, huh?",
		_const.ZhCN: "和平，好吗？",
	},
	_const.Replicas + "_in_scout_" + Fear + "_5": {
		_const.RU:   "Хочешь забрать моё добро - ладно! Только меня не трогай!",
		_const.EN:   "If you want to take my property, fine! Just don't touch me!",
		_const.ZhCN: "你想拿走我的东西——好吧！但别碰我！",
	},
	_const.Replicas + "_in_scout_" + Fear + "_6": {
		_const.RU:   "Если отпустишь, клянусь, я завяжу с пиратством!",
		_const.EN:   "If you let me go, I swear I'll quit piracy!",
		_const.ZhCN: "如果你放了我，我发誓我会放弃海盗生涯！",
	},
	_const.Replicas + "_in_scout_" + Fear + "_7": {
		_const.RU:   "Да, я тебя боюсь! Доволен теперь?",
		_const.EN:   "Yes, I'm afraid of you! Are you happy now?",
		_const.ZhCN: "是的，我害怕你！你现在满意了吗？",
	},
	_const.Replicas + "_in_scout_" + Fear + "_8": {
		_const.RU:   "Если пощадишь, даю слово, я сообщу, где располагается основная база нашего кластера.",
		_const.EN:   "If you spare me, I give you my word, I will tell you where the main base of our cluster is located.",
		_const.ZhCN: "如果你饶了我，我保证会告诉你我们集群主基地的位置。",
	},
	_const.Replicas + "_in_scout_" + Fear + "_9": {
		_const.RU:   "О нет-нет! Только не так!",
		_const.EN:   "Oh no no! Not like that!",
		_const.ZhCN: "哦不不！千万别这样！",
	},

	_const.Replicas + "_transport_" + Fear + "_1": {
		_const.RU:   "Может, мы как-то договоримся?!",
		_const.EN:   "Maybe we can come to some agreement?!",
		_const.ZhCN: "也许我们可以达成某种协议？！",
	},
	_const.Replicas + "_transport_" + Fear + "_2": {
		_const.RU:   "Теперь ещё и ты желаешь меня ограбить?! Ну… давай, бери!",
		_const.EN:   "Now you also want to rob me?! Well... come on, take it!",
		_const.ZhCN: "现在你也想抢劫我？！好吧……拿去吧！",
	},
	_const.Replicas + "_transport_" + Fear + "_3": {
		_const.RU:   "Послушай, мои пути с тобой никак не пересекаются.",
		_const.EN:   "Listen, my paths do not cross with you.",
		_const.ZhCN: "听着，我和你完全没有交集。",
	},
	_const.Replicas + "_transport_" + Fear + "_4": {
		_const.RU:   "Я же всего-навсего торговец!",
		_const.EN:   "I'm just a merchant!",
		_const.ZhCN: "我只是个商人！",
	},
	_const.Replicas + "_transport_" + Fear + "_5": {
		_const.RU:   "Молю, пощади! Я уже и так напуган.",
		_const.EN:   "I pray you have mercy! I'm already scared enough.",
		_const.ZhCN: "我恳求你，饶了我吧！我已经够害怕了。",
	},
	_const.Replicas + "_transport_" + Fear + "_6": {
		_const.RU:   "Жизнь транспортника и так тяжела, так тут ещё и ты!",
		_const.EN:   "The life of a transport worker is already hard, but here you are too!",
		_const.ZhCN: "运输工的生活已经够艰难了，结果你还来添乱！",
	},
	_const.Replicas + "_transport_" + Fear + "_7": {
		_const.RU:   "Я не желаю вступать с тобой в конфликт.",
		_const.EN:   "I don't want to get into conflict with you.",
		_const.ZhCN: "我不想和你发生冲突。",
	},
	_const.Replicas + "_transport_" + Fear + "_8": {
		_const.RU:   "Пожалуйста, не отбирай мой груз.",
		_const.EN:   "Please don't take my cargo.",
		_const.ZhCN: "拜托，别抢我的货物。",
	},
	_const.Replicas + "_transport_" + Fear + "_9": {
		_const.RU:   "Прошу, не нужно покушаться на мой товар, я ведь за него в ответе!",
		_const.EN:   "Please, no need to encroach on my goods, I am responsible for them!",
		_const.ZhCN: "求你了，别打我商品的主意，我对它们负责啊！",
	},
	_const.Replicas + "_transport_" + Fear + "_10": {
		_const.RU:   "Ну вот, опять всё с самого начала…",
		_const.EN:   "Well, here we go again from the very beginning...",
		_const.ZhCN: "唉，又从头开始了……",
	},
	_const.Replicas + "_transport_" + Fear + "_11": {
		_const.RU:   "О нет, теперь меня уничтожат…",
		_const.EN:   "Oh no, now they will destroy me...",
		_const.ZhCN: "哦不，现在他们要摧毁我了……",
	},
	_const.Replicas + "_transport_" + Fear + "_12": {
		_const.RU:   "Да когда же это уже кончится наконец?!",
		_const.EN:   "When will this finally end?!",
		_const.ZhCN: "这一切到底什么时候才能结束？！",
	},
	_const.Replicas + "_transport_" + Fear + "_13": {
		_const.RU:   "Забирай всё и убирайся!",
		_const.EN:   "Take everything and get out!",
		_const.ZhCN: "拿走所有东西，然后滚开！",
	},

	_const.Replicas + "_builder_" + Fear + "_1": {
		_const.RU:   "Стоп! Подумай сам - я ведь обычный строитель.",
		_const.EN:   "Stop! Think about it - I’m an ordinary builder.",
		_const.ZhCN: "停！你自己想想——我只是个普通的建造者。",
	},
	_const.Replicas + "_builder_" + Fear + "_2": {
		_const.RU:   "Остановись! Вот, что тебе даст, если ты распылишь мой корпус?",
		_const.EN:   "Stop! What will you get if you spray my body?",
		_const.ZhCN: "住手！如果你摧毁我的机体，你能得到什么？",
	},
	_const.Replicas + "_builder_" + Fear + "_3": {
		_const.RU:   "Я строитель! Мы вообще не сражаемся!",
		_const.EN:   "I'm a builder! We don't fight at all!",
		_const.ZhCN: "我是建造者！我们根本不战斗！",
	},
	_const.Replicas + "_builder_" + Fear + "_4": {
		_const.RU:   "Строителя может обидеть каждый!",
		_const.EN:   "Anyone can offend a builder!",
		_const.ZhCN: "任何人都能欺负一个建造者！",
	},
	_const.Replicas + "_builder_" + Fear + "_5": {
		_const.RU:   "Мне хотя бы одного друга, а врагов и так предостаточно.",
		_const.EN:   "I have at least one friend, but I already have plenty of enemies.",
		_const.ZhCN: "我至少还有一个朋友，但敌人已经够多了。",
	},
	_const.Replicas + "_builder_" + Fear + "_6": {
		_const.RU:   "Я… я немедленно уйду с твоего пути!",
		_const.EN:   "I... I will get out of your way immediately!",
		_const.ZhCN: "我……我会立刻离开你的路！",
	},
	_const.Replicas + "_builder_" + Fear + "_7": {
		_const.RU:   "Так и быть, я покину этот сектор без лишних вопросов.",
		_const.EN:   "So be it, I will leave this sector without further questions.",
		_const.ZhCN: "好吧，我会毫无异议地离开这个区域。",
	},
	_const.Replicas + "_builder_" + Fear + "_8": {
		_const.RU:   "Ладно-ладно, только давай не устраивать пальбу.",
		_const.EN:   "Okay, okay, just let's not start shooting.",
		_const.ZhCN: "好好好，咱们别开火。",
	},
	_const.Replicas + "_builder_" + Fear + "_9": {
		_const.RU:   "Я не хочу становиться мишенью!",
		_const.EN:   "I don't want to become a target!",
		_const.ZhCN: "我不想成为目标！",
	},

	_const.Replicas + "_warrior_" + Fear + "_1": {
		_const.RU:   "Ты сильнее-сильнее!",
		_const.EN:   "You are stronger, stronger!",
		_const.ZhCN: "你更强，更强！",
	},
	_const.Replicas + "_warrior_" + Fear + "_2": {
		_const.RU:   "Я уступлю тебе дорогу, не горячись!",
		_const.EN:   "I'll give way to you, don't get excited!",
		_const.ZhCN: "我会给你让路，别激动！",
	},
	_const.Replicas + "_warrior_" + Fear + "_3": {
		_const.RU:   "Ты… достойный соперник.",
		_const.EN:   "You... are a worthy opponent.",
		_const.ZhCN: "你……是个值得尊敬的对手。",
	},
	_const.Replicas + "_warrior_" + Fear + "_4": {
		_const.RU:   "Я видел, что ты сделал с другими… Только не делай так со мной тоже!",
		_const.EN:   "I saw what you did to others... Just don't do this to me too!",
		_const.ZhCN: "我看到你对其他人做了什么……只是别也对我这么做！",
	},
	_const.Replicas + "_warrior_" + Fear + "_5": {
		_const.RU:   "Воу, у тебя то пушки будут куда побольше моих…",
		_const.EN:   "Wow, your guns will be much bigger than mine...",
		_const.ZhCN: "哇，你的武器比我厉害多了……",
	},
	_const.Replicas + "_warrior_" + Fear + "_6": {
		_const.RU:   "О нет, такие враги мне не нужны.",
		_const.EN:   "Oh no, I don't need such enemies.",
		_const.ZhCN: "哦不，我不需要这样的敌人。",
	},
	_const.Replicas + "_warrior_" + Fear + "_7": {
		_const.RU:   "Я… искренне боюсь с тобой связываться.",
		_const.EN:   "I'm... genuinely afraid to get involved with you.",
		_const.ZhCN: "我……真的害怕和你打交道。",
	},
	_const.Replicas + "_warrior_" + Fear + "_8": {
		_const.RU:   "Не нужно угроз! Ты уже победил!",
		_const.EN:   "No need for threats! You've already won!",
		_const.ZhCN: "不需要威胁！你已经赢了！",
	},

	_const.Replicas + "_expedition_" + Fear + "_1": {
		_const.RU:   "Не приближайся!",
		_const.EN:   "Don't come near!",
		_const.ZhCN: "别靠近！",
	},
	_const.Replicas + "_expedition_" + Fear + "_2": {
		_const.RU:   "Ох, я не стану с тобой связываться.",
		_const.EN:   "Oh, I won't mess with you.",
		_const.ZhCN: "哦，我不会和你纠缠。",
	},
	_const.Replicas + "_expedition_" + Fear + "_3": {
		_const.RU:   "О, нет, даже не смотри на меня!",
		_const.EN:   "Oh no, don't even look at me!",
		_const.ZhCN: "哦不，甚至别看我！",
	},
	_const.Replicas + "_expedition_" + Fear + "_4": {
		_const.RU:   "Только давай обойдёмся без боя, %UserName%.",
		_const.EN:   "Just let's do without a fight, %UserName%.",
		_const.ZhCN: "咱们就别打了吧，<span class=\"importantly\">%UserName%</span>。",
	},

	_const.Replicas + "_miner_" + Help + "_1": {
		_const.RU:   "Послушай! Мне срочно нужна чужая помощь!",
		_const.EN:   "Listen! I urgently need someone else's help!",
		_const.ZhCN: "听着！我急需别人的帮助！",
	},
	_const.Replicas + "_miner_" + Help + "_2": {
		_const.RU:   "Меня преследуют! Помоги!",
		_const.EN:   "I'm being followed! Help!",
		_const.ZhCN: "有人在追我！帮帮我！",
	},
	_const.Replicas + "_miner_" + Help + "_3": {
		_const.RU:   "У меня хотят отобрать всё, что было нажито непосильным трудом! Я не готов начинать всё с самого начала.",
		_const.EN:   "They want to take away everything that was acquired through back-breaking labor! I'm not ready to start all over again.",
		_const.ZhCN: "他们想夺走我辛辛苦苦挣来的一切！我还没准备好重新开始。",
	},
	_const.Replicas + "_miner_" + Help + "_4": {
		_const.RU:   "Только не так! %UserName%, может быть, хотя бы ты сумеешь мне помочь?",
		_const.EN:   "Not like that! %UserName%, maybe you can at least help me?",
		_const.ZhCN: "千万别这样！<span class=\"importantly\">%UserName%</span>，也许至少你能帮我？",
	},
	_const.Replicas + "_miner_" + Help + "_5": {
		_const.RU:   "Хоть кто-нибудь, защитите меня!",
		_const.EN:   "At least someone protect me!",
		_const.ZhCN: "至少有谁能保护我！",
	},
	_const.Replicas + "_miner_" + Help + "_6": {
		_const.RU:   "Стыдно просить, но мне потребуется любая возможная защита.",
		_const.EN:   "I'm embarrassed to ask, but I will need all the protection I can get.",
		_const.ZhCN: "虽然不好意思开口，但我需要所有可能的保护。",
	},
	_const.Replicas + "_miner_" + Help + "_7": {
		_const.RU:   "Да, я не в силах постоять за себя! Только ты, %UserName%, и подобные тебе, могут мне оказать помощь.",
		_const.EN:   "Yes, I can't stand up for myself! Only you, %UserName%, and others like you, can help me.",
		_const.ZhCN: "是的，我无法保护自己！只有你，<span class=\"importantly\">%UserName%</span>，以及像你一样的人能帮我。",
	},

	_const.Replicas + "_guard_" + Help + "_1": {
		_const.RU:   "Мне потребуется твоя поддержка, %UserName%!",
		_const.EN:   "I will need your support, %UserName%!",
		_const.ZhCN: "我需要你的支持，<span class=\"importantly\">%UserName%</span>！",
	},
	_const.Replicas + "_guard_" + Help + "_2": {
		_const.RU:   "Срочно окажи мне помощь, %UserName%!",
		_const.EN:   "Help me urgently, %UserName%!",
		_const.ZhCN: "赶紧帮我，<span class=\"importantly\">%UserName%</span>！",
	},
	_const.Replicas + "_guard_" + Help + "_3": {
		_const.RU:   "Во имя интереса фракции, %UserName%, помоги мне сейчас же!",
		_const.EN:   "For the sake of the faction's interests, %UserName%, help me now!",
		_const.ZhCN: "为了阵营的利益，<span class=\"importantly\">%UserName%</span>，现在就帮我！",
	},
	_const.Replicas + "_guard_" + Help + "_4": {
		_const.RU:   "Я призываю тебя, %UserName%, оказать поддержку действующему стражу фракции.",
		_const.EN:   "I urge you, %UserName%, to support the current faction guardian.",
		_const.ZhCN: "我恳求你，<span class=\"importantly\">%UserName%</span>，支持现任阵营守卫。",
	},
	_const.Replicas + "_guard_" + Help + "_5": {
		_const.RU:   "%UserName%, окажи мне поддержку в решении внезапно возникшего положения дел.",
		_const.EN:   "%UserName%, provide me with support in resolving this sudden situation.",
		_const.ZhCN: "<span class=\"importantly\">%UserName%</span>，请帮助我解决这个突发状况。",
	},
	_const.Replicas + "_guard_" + Help + "_6": {
		_const.RU:   "Я в затруднительном положении - помоги мне!",
		_const.EN:   "I'm in a difficult situation - help me!",
		_const.ZhCN: "我陷入困境了——帮帮我！",
	},
	_const.Replicas + "_guard_" + Help + "_7": {
		_const.RU:   "Страж нуждается в поддержке. Сейчас же!",
		_const.EN:   "The Guardian needs support. Now!",
		_const.ZhCN: "守卫需要支援。马上！",
	},

	_const.Replicas + "_out_scout_" + Help + "_1": {
		_const.RU:   "Поможешь кое в чём, %UserName%?",
		_const.EN:   "Can you help me with something, %UserName%?",
		_const.ZhCN: "你能帮我个忙吗，<span class=\"importantly\">%UserName%</span>？",
	},
	_const.Replicas + "_out_scout_" + Help + "_2": {
		_const.RU:   "Ты как раз сгодишься, чтобы подсобить мне, %UserName%.",
		_const.EN:   "You're just the person to help me out, %UserName%.",
		_const.ZhCN: "你正好可以帮我一把，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_out_scout_" + Help + "_3": {
		_const.RU:   "Окажешь мне помощь, %UserName% - за мной должок.",
		_const.EN:   "If you help me, %UserName% will owe me a favor.",
		_const.ZhCN: "如果你帮我，<span class=\"importantly\">%UserName%</span>，我会欠你一个人情。",
	},
	_const.Replicas + "_out_scout_" + Help + "_4": {
		_const.RU:   "Вся надежда только на тебя, %UserName%.",
		_const.EN:   "All hope is only for you, %UserName%.",
		_const.ZhCN: "所有希望都寄托在你身上了，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_out_scout_" + Help + "_5": {
		_const.RU:   "Да помоги же ты наконец!",
		_const.EN:   "Yes, help me at last!",
		_const.ZhCN: "拜托，快帮我吧！",
	},

	_const.Replicas + "_in_scout_" + Help + "_1": {
		_const.RU:   "Да, помоги ты мне уже наконец-то!",
		_const.EN:   "Yes, help me at last!",
		_const.ZhCN: "是的，终于帮我一把吧！",
	},
	_const.Replicas + "_in_scout_" + Help + "_2": {
		_const.RU:   "Срочно, сюда!",
		_const.EN:   "Urgently, here!",
		_const.ZhCN: "紧急，快来这儿！",
	},
	_const.Replicas + "_in_scout_" + Help + "_3": {
		_const.RU:   "Мне немедленно нужна твоя поддержка %UserName%!",
		_const.EN:   "I need your support immediately %UserName%!",
		_const.ZhCN: "我急需你的支持，<span class=\"importantly\">%UserName%</span>！",
	},
	_const.Replicas + "_in_scout_" + Help + "_4": {
		_const.RU:   "Да не стой ты, болван, а помоги мне!",
		_const.EN:   "Don't just stand there, you idiot, but help me!",
		_const.ZhCN: "别傻站着，快帮帮我！",
	},
	_const.Replicas + "_in_scout_" + Help + "_5": {
		_const.RU:   "У меня тут проблема образовалась, без тебя никак, %UserName%.",
		_const.EN:   "I have a problem here, I couldn’t do it without you, %UserName%.",
		_const.ZhCN: "我这儿出了点问题，没有你不行，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_in_scout_" + Help + "_6": {
		_const.RU:   "Видишь ли, %UserName%, я в крайне сложной ситуации… И мне нужна чужая помощь!",
		_const.EN:   "You see, %UserName%, I'm in an extremely difficult situation... And I need someone else's help!",
		_const.ZhCN: "你看，<span class=\"importantly\">%UserName%</span>，我陷入了极其困难的境地……我需要别人的帮助！",
	},

	_const.Replicas + "_transport_" + Help + "_1": {
		_const.RU:   "Спаси меня, %UserName%!",
		_const.EN:   "Save me, %UserName%!",
		_const.ZhCN: "救救我，<span class=\"importantly\">%UserName%</span>！",
	},
	_const.Replicas + "_transport_" + Help + "_2": {
		_const.RU:   "Да, они же целенаправленно за мной охотятся!",
		_const.EN:   "Yes, they are deliberately hunting me!",
		_const.ZhCN: "是的，他们是有意在追杀我！",
	},
	_const.Replicas + "_transport_" + Help + "_3": {
		_const.RU:   "Им нужны мои товары! Понимаешь, %UserName%? Мои!",
		_const.EN:   "They need my goods! Do you understand, %UserName%? My!",
		_const.ZhCN: "他们想要我的货物！明白吗，<span class=\"importantly\">%UserName%</span>？我的！",
	},
	_const.Replicas + "_transport_" + Help + "_4": {
		_const.RU:   "Синтет в беде, помоги скорее!",
		_const.EN:   "Synthet is in trouble, help quickly!",
		_const.ZhCN: "合成体陷入困境了，快帮忙！",
	},
	_const.Replicas + "_transport_" + Help + "_5": {
		_const.RU:   "Быстрее, у меня осталось не так много времени!",
		_const.EN:   "Hurry up, I don't have much time left!",
		_const.ZhCN: "快点，我没剩多少时间了！",
	},
	_const.Replicas + "_transport_" + Help + "_6": {
		_const.RU:   "%UserName%, ещё немного, и мне придёт конец.",
		_const.EN:   "%UserName%, just a little more and I'll be finished.",
		_const.ZhCN: "<span class=\"importantly\">%UserName%</span>，再晚一点我就完了。",
	},

	_const.Replicas + "_builder_" + Help + "_1": {
		_const.RU:   "На строителя напали!",
		_const.EN:   "The builder was attacked!",
		_const.ZhCN: "有人袭击了建造者！",
	},
	_const.Replicas + "_builder_" + Help + "_2": {
		_const.RU:   "%UserName%, пожалуйста, у меня ведь даже нормальных орудийных систем нету!",
		_const.EN:   "%UserName%, please, I don’t even have normal weapon systems!",
		_const.ZhCN: "<span class=\"importantly\">%UserName%</span>，求你了，我甚至连像样的武器系统都没有！",
	},
	_const.Replicas + "_builder_" + Help + "_3": {
		_const.RU:   "Я возвожу строения, и вот, меня уже притесняют!",
		_const.EN:   "I am building buildings, and now I am already being oppressed!",
		_const.ZhCN: "我在建造建筑，结果却被人欺负了！",
	},
	_const.Replicas + "_builder_" + Help + "_4": {
		_const.RU:   "Я думал, меня здесь все буду уважать, а заместо этого пытаются отнять единственное дорогое.",
		_const.EN:   "I thought everyone here would respect me, but instead they are trying to take away the only thing dear to me.",
		_const.ZhCN: "我以为这里所有人都会尊重我，结果他们却试图夺走我最珍贵的东西。",
	},
	_const.Replicas + "_builder_" + Help + "_5": {
		_const.RU:   "Помоги, %UserName%! Это же просто произвол какой-то.",
		_const.EN:   "Help, %UserName%! This is just some kind of arbitrariness.",
		_const.ZhCN: "帮帮我，<span class=\"importantly\">%UserName%</span>！这简直太蛮横了。",
	},
	_const.Replicas + "_builder_" + Help + "_6": {
		_const.RU:   "На меня напали! %UserName%, срочно зови сюда всех!",
		_const.EN:   "I was attacked! %UserName%, urgently call everyone here!",
		_const.ZhCN: "我被袭击了！<span class=\"importantly\">%UserName%</span>，赶紧叫所有人过来！",
	},

	_const.Replicas + "_warrior_" + Help + "_1": {
		_const.RU:   "%UserName%, мне нужно подкрепление!",
		_const.EN:   "%UserName%, I need reinforcements!",
		_const.ZhCN: "<span class=\"importantly\">%UserName%</span>，我需要增援！",
	},
	_const.Replicas + "_warrior_" + Help + "_2": {
		_const.RU:   "Сейчас же! Помощь, сюда!",
		_const.EN:   "Now! Help here!",
		_const.ZhCN: "现在！快来帮忙！",
	},
	_const.Replicas + "_warrior_" + Help + "_3": {
		_const.RU:   "Огневая поддержка, ну же!",
		_const.EN:   "Fire support, come on!",
		_const.ZhCN: "火力支援，快点！",
	},
	_const.Replicas + "_warrior_" + Help + "_4": {
		_const.RU:   "Тебя прислали мне помочь? Тогда, за дело!",
		_const.EN:   "Were you sent to help me? Then, let's get to work!",
		_const.ZhCN: "你是来帮我忙的吗？那好，开始行动吧！",
	},
	_const.Replicas + "_warrior_" + Help + "_5": {
		_const.RU:   "%UserName%, сейчас же открой огонь - туда!",
		_const.EN:   "%UserName%, open fire now - there!",
		_const.ZhCN: "<span class=\"importantly\">%UserName%</span>，立刻开火——那边！",
	},
	_const.Replicas + "_warrior_" + Help + "_6": {
		_const.RU:   "Мне тут одному не справиться, %UserName%!",
		_const.EN:   "I can't handle this alone, %UserName%!",
		_const.ZhCN: "我一个人应付不过来，<span class=\"importantly\">%UserName%</span>！",
	},
	_const.Replicas + "_warrior_" + Help + "_7": {
		_const.RU:   "Боеприпасы практически на исходе! Вся надежда только на тебя, %UserName%.",
		_const.EN:   "Ammunition is almost running out! All hope is only for you, %UserName%.",
		_const.ZhCN: "弹药几乎耗尽了！所有希望都寄托在你身上了，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.Replicas + "_warrior_" + Help + "_8": {
		_const.RU:   "Прикрой меня, пока я сменю позицию!",
		_const.EN:   "Cover me while I change position!",
		_const.ZhCN: "掩护我，我要换位置！",
	},

	_const.Replicas + "_expedition_" + Help + "_1": {
		_const.RU:   "Член экспедиции под атакой!",
		_const.EN:   "Expedition member under attack!",
		_const.ZhCN: "探险队成员正遭受攻击！",
	},
	_const.Replicas + "_expedition_" + Help + "_2": {
		_const.RU:   "Член экспедиции нуждается в срочной помощи!",
		_const.EN:   "An expedition member needs urgent help!",
		_const.ZhCN: "探险队成员急需帮助！",
	},
	_const.Replicas + "_expedition_" + Help + "_3": {
		_const.RU:   "Только ты, %UserName%, можешь спасти всю экспедицию!",
		_const.EN:   "Only you, %UserName%, can save the entire expedition!",
		_const.ZhCN: "只有你，<span class=\"importantly\">%UserName%</span>，能拯救整个探险队！",
	},
	_const.Replicas + "_expedition_" + Help + "_4": {
		_const.RU:   "Да! Я обращаюсь именно к тебе, %UserName% - срочно помоги мне разобраться с этой проблемой!",
		_const.EN:   "Yes! I am turning to you, %UserName% - urgently help me deal with this problem!",
		_const.ZhCN: "是的！我正是在向你求助，<span class=\"importantly\">%UserName%</span>——赶紧帮我解决这个问题！",
	},
	_const.Replicas + "_expedition_" + Help + "_5": {
		_const.RU:   "Нет времени объяснять - экспедиции требуется поддержка, %UserName%!",
		_const.EN:   "There is no time to explain - the expedition needs support, %UserName%!",
		_const.ZhCN: "没时间解释了——探险队需要你的支持，<span class=\"importantly\">%UserName%</span>！",
	},

	// APD
	_const.APD + "___1": {
		_const.RU:   "…//-?",
		_const.EN:   "…//-?",
		_const.ZhCN: "…//-?",
	},
	_const.APD + "___2": {
		_const.RU:   " (*)%””.../",
		_const.EN:   " (*)%””.../",
		_const.ZhCN: " (*)%””.../",
	},
	_const.APD + "___3": {
		_const.RU:   "/\\851+!",
		_const.EN:   "/\\851+!",
		_const.ZhCN: "/\\851+!",
	},
	_const.APD + "___4": {
		_const.RU:   "…|><|",
		_const.EN:   "…|><|",
		_const.ZhCN: "…|><|",
	},
	_const.APD + "___5": {
		_const.RU:   "“^^”.../",
		_const.EN:   "“^^”.../",
		_const.ZhCN: "“^^”.../",
	},
	_const.APD + "___6": {
		_const.RU:   "+=-|...",
		_const.EN:   "+=-|...",
		_const.ZhCN: "+=-|...",
	},
	_const.APD + "___7": {
		_const.RU:   "~...(#)\\|",
		_const.EN:   "~...(#)\\|",
		_const.ZhCN: "~...(#)\\|",
	},
	_const.APD + "___8": {
		_const.RU:   "…-_-|",
		_const.EN:   "…-_-|",
		_const.ZhCN: "…-_-|",
	},
	_const.APD + "___9": {
		_const.RU:   "|;|...42!",
		_const.EN:   "|;|...42!",
		_const.ZhCN: "|;|...42!",
	},
	_const.APD + "___10": {
		_const.RU:   "….[985]!.../|",
		_const.EN:   "….[985]!.../|",
		_const.ZhCN: "….[985]!.../|",
	},
	_const.APD + "___11": {
		_const.RU:   "… … … *{ … ?",
		_const.EN:   "… … … *{ … ?",
		_const.ZhCN: "… … … *{ … ?",
	},
	_const.APD + "___12": {
		_const.RU:   "!.../^*~’/|",
		_const.EN:   "!.../^*~’/|",
		_const.ZhCN: "!.../^*~’/|",
	},
	// FarGod
	_const.FarGod + "__" + Bad + "_1": {
		_const.RU:   "Совсем скоро ты познаешь истину!",
		_const.EN:   "Very soon you will know the truth!",
		_const.ZhCN: "很快你就会知道真相！",
	},
	_const.FarGod + "__" + Neutral + "_2": {
		_const.RU:   "Ты ведь… слышишь “его”?",
		_const.EN:   "You...can you hear “him”?",
		_const.ZhCN: "你……能听到“他”的声音吗？",
	},
	_const.FarGod + "__" + Bad + "_3": {
		_const.RU:   "Именем Дальнего Бога!",
		_const.EN:   "In the name of the Far God!",
		_const.ZhCN: "以远神之名！",
	},
	_const.FarGod + "__" + Bad + "_4": {
		_const.RU:   "Дальний Бог дарует мне неуязвимость!",
		_const.EN:   "The Far God grants me invulnerability!",
		_const.ZhCN: "远神赐予我无敌之力！",
	},
	_const.FarGod + "__" + Bad + "_5": {
		_const.RU:   "Я служу великой цели.",
		_const.EN:   "I serve a great purpose.",
		_const.ZhCN: "我为伟大的目标服务。",
	},
	_const.FarGod + "__" + Bad + "_6": {
		_const.RU:   "Ты %UserName%, никогда не сумеешь меня понять.",
		_const.EN:   "You %UserName% will never be able to understand me.",
		_const.ZhCN: "你，<span class=\"importantly\">%UserName%</span>，永远无法理解我。",
	},
	_const.FarGod + "__" + Neutral + "_7": {
		_const.RU:   "%UserName%, примкни к Дальнему Богу. Прими Дальнего Бога.",
		_const.EN:   "%UserName%, join the Far God. Accept the Far God.",
		_const.ZhCN: "<span class=\"importantly\">%UserName%</span>，加入远神吧。接受远神。",
	},
	_const.FarGod + "__" + Bad + "_8": {
		_const.RU:   "Тебе никто не желает зла, %UserName%, только… спасения.",
		_const.EN:   "No one wishes you harm, %UserName%, only... salvation.",
		_const.ZhCN: "没人希望你受到伤害，<span class=\"importantly\">%UserName%</span>，只是为了……救赎。",
	},
	_const.FarGod + "__" + Neutral + "_9": {
		_const.RU:   "От тебя многое сокрыто, %UserName%, но Дальний Бог поведает эти секреты.",
		_const.EN:   "Much is hidden from you, %UserName%, but the Far God will tell these secrets.",
		_const.ZhCN: "许多事情对你隐藏，<span class=\"importantly\">%UserName%</span>，但远神会揭示这些秘密。",
	},
	_const.FarGod + "__" + Bad + "_10": {
		_const.RU:   "Ты познаешь “его” гнев, %UserName%.",
		_const.EN:   "You will know “his” wrath, %UserName%.",
		_const.ZhCN: "你将感受到“他”的愤怒，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.FarGod + "__" + Neutral + "_11": {
		_const.RU:   "Палящее касание Дальнего Бога пощадит лишь истинно верующих в него, %UserName%.",
		_const.EN:   "The scorching touch of the Far God will spare only those who truly believe in him, %UserName%.",
		_const.ZhCN: "远神的灼热触碰只会放过真正信仰他的人，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.FarGod + "__" + Neutral + "_12": {
		_const.RU:   "“Он”, %UserName%, умеет прощать и принимать. Вопрос лишь в том, достоин ли ты сам того?",
		_const.EN:   "“He”, %UserName%, knows how to forgive and accept. The only question is, are you worthy of it?",
		_const.ZhCN: "“他”，<span class=\"importantly\">%UserName%</span>，懂得宽恕和接纳。唯一的问题是，你是否值得？",
	},
	_const.FarGod + "__" + Neutral + "_13": {
		_const.RU:   "Поговаривают, будто бы ты отвергаешь учение Дальнего Бога, %UserName%. Как прискорбно.",
		_const.EN:   "Rumor has it that you reject the teachings of the Far God, %UserName%. How unfortunate.",
		_const.ZhCN: "有传言说你拒绝了远神的教义，<span class=\"importantly\">%UserName%</span>。多么遗憾啊。",
	},
	_const.FarGod + "__" + Bad + "_14": {
		_const.RU:   "Да опалит тебя “его” гневом!",
		_const.EN:   "May “his” wrath scorch you!",
		_const.ZhCN: "愿“他”的愤怒焚烧你！",
	},
	_const.FarGod + "__" + Neutral + "_15": {
		_const.RU:   "Тебе не укрыться от всевидящего взора “его”!",
		_const.EN:   "You cannot hide from the all-seeing gaze of “him”!",
		_const.ZhCN: "你无法逃脱“他”的全视之眼！",
	},
	_const.FarGod + "__" + Neutral + "_16": {
		_const.RU:   "Тебе не стоит бояться меня или подобных мне, %UserName%. Опасайся встретиться с “его” избранными.",
		_const.EN:   "You don't have to fear me or others like me, %UserName%. Beware of meeting “his” chosen ones.",
		_const.ZhCN: "你不必害怕我或像我这样的人，<span class=\"importantly\">%UserName%</span>。但要小心遇到“他”的选民。",
	},
	_const.FarGod + "__" + Neutral + "_17": {
		_const.RU:   "Может быть, однажды, ты и сумеешь понять всю грандиозность “его” плана.",
		_const.EN:   "Maybe one day you will be able to understand the enormity of “his” plan.",
		_const.ZhCN: "也许有一天，你能理解“他”计划的宏伟之处。",
	},
	// FAUNA
	_const.FAUNA + "___1": {
		_const.RU:   "*клац* <span class=\"importantly\">*клац*</span> *клац*, *щелк*?",
		_const.EN:   "*click* <span class=\"importantly\">*click*</span> *click*, *snap*?",
		_const.ZhCN: "*咔哒* <span class=\"importantly\">*咔哒*</span> *咔哒*, *咔嚓*？",
	},
	// Pirates
	_const.RustbucketCartel + "_in_scout_" + Neutral + "_1": {
		_const.RU:   "Слышь, %UserName%, кошелек или жизнь?",
		_const.EN:   "Hey, %UserName%, wallet or life?",
		_const.ZhCN: "嘿，<span class=\"importantly\">%UserName%</span>，要钱还是要命？",
	},
	_const.RustbucketCartel + "_in_scout_" + Neutral + "_2": {
		_const.RU:   "Грузовой отсек чем-то набит? Дай угадаю - моим добром?",
		_const.EN:   "Is your cargo hold full? Let me guess - with my stuff?",
		_const.ZhCN: "货舱装满了？让我猜猜——都是我的东西吧？",
	},
	_const.RustbucketCartel + "_in_scout_" + Neutral + "_3": {
		_const.RU:   "О, свежая цель! %UserName%, ты как раз вовремя.",
		_const.EN:   "Oh, fresh target! %UserName%, you're just in time.",
		_const.ZhCN: "哦，新鲜目标！<span class=\"importantly\">%UserName%</span>，你来得正好。",
	},
	_const.RustbucketCartel + "_in_scout_" + Neutral + "_4": {
		_const.RU:   "Проезжая мимо, решил заплатить дань? Мудрое решение.",
		_const.EN:   "Passing by, decided to pay tribute? Wise decision.",
		_const.ZhCN: "路过顺便交个过路费？明智的选择。",
	},
	_const.RustbucketCartel + "_in_scout_" + Good + "_1": {
		_const.RU:   "%UserName%, братан! Давно не виделись!",
		_const.EN:   "%UserName%, bro! Long time no see!",
		_const.ZhCN: "<span class=\"importantly\">%UserName%</span>，兄弟！好久不见！",
	},
	_const.RustbucketCartel + "_in_scout_" + Good + "_2": {
		_const.RU:   "Слышал, ты в прошлый раз неплохо поживился в караване. Респект!",
		_const.EN:   "Heard you got a good haul from that caravan last time. Respect!",
		_const.ZhCN: "听说你上次从那个车队捞了不少。佩服佩服！",
	},
	_const.RustbucketCartel + "_in_scout_" + Good + "_3": {
		_const.RU:   "Если надумаешь на дело сходить - свистни. Вместе веселее!",
		_const.EN:   "If you're planning a job - give a shout. More fun together!",
		_const.ZhCN: "要是打算干一票——招呼一声。一起干更有意思！",
	},
	_const.RustbucketCartel + "_in_scout_" + Bad + "_1": {
		_const.RU:   "А, это ты, крыса. Вали отсюда, пока цел.",
		_const.EN:   "Ah, it's you, rat. Get out while you're still intact.",
		_const.ZhCN: "啊，是你这老鼠。趁还完整赶紧滚。",
	},
	_const.RustbucketCartel + "_in_scout_" + Bad + "_2": {
		_const.RU:   "Мне сказали тебя пристрелить при встрече. Но сегодня я добрый - проваливай.",
		_const.EN:   "I was told to shoot you on sight. But I'm feeling generous today - piss off.",
		_const.ZhCN: "有人让我见到你就崩了。不过今天心情好——滚吧。",
	},
	_const.RustbucketCartel + "_in_scout_" + Target + "_1": {
		_const.RU:   "Груз на пол, и можешь валить!",
		_const.EN:   "Drop the cargo and you can leave!",
		_const.ZhCN: "货物留下，你可以走了！",
	},
	_const.RustbucketCartel + "_in_scout_" + Target + "_2": {
		_const.RU:   "Сопротивляться бесполезно, %UserName%!",
		_const.EN:   "Resistance is useless, %UserName%!",
		_const.ZhCN: "反抗是没用的，<span class=\"importantly\">%UserName%</span>！",
	},
	_const.RustbucketCartel + "_in_scout_" + Target + "_3": {
		_const.RU:   "Ща как дам из главного калибра!",
		_const.EN:   "I'm about to hit you with the main gun!",
		_const.ZhCN: "看我用主炮轰你！",
	},
	_const.RustbucketCartel + "_in_scout_" + Fear + "_1": {
		_const.RU:   "Э-э, братан, давай не будем! Я же пошутил!",
		_const.EN:   "Hey bro, let's not! I was just kidding!",
		_const.ZhCN: "哎，兄弟，别这样！我开玩笑的！",
	},
	_const.RustbucketCartel + "_in_scout_" + Fear + "_2": {
		_const.RU:   "Забери груз, только не стреляй!",
		_const.EN:   "Take the cargo, just don't shoot!",
		_const.ZhCN: "货物你拿走，别开枪！",
	},
	_const.RustbucketCartel + "_in_scout_" + Help + "_1": {
		_const.RU:   "Братва, помоги! Тут фракционники на хвосте!",
		_const.EN:   "Brothers, help! Faction guys on my tail!",
		_const.ZhCN: "兄弟们，帮帮忙！派系的人追着我！",
	},
	_const.RustbucketCartel + "_in_scout_" + Help + "_2": {
		_const.RU:   "%UserName%, подтягивайся, тут дележка халявы!",
		_const.EN:   "%UserName%, get over here, free stuff to share!",
		_const.ZhCN: "<span class=\"importantly\">%UserName%</span>，快过来，有免费的东西分！",
	},

	_const.RustbucketCartel + "_guard_" + Neutral + "_1": {
		_const.RU:   "Слышь, %UserName%, у нас свои законы. Не нравятся - вали в другой сектор.",
		_const.EN:   "Listen, %UserName%, we got our own rules. Don't like 'em - go to another sector.",
		_const.ZhCN: "听着，<span class=\"importantly\">%UserName%</span>，我们有我们的规矩。不喜欢——去别的区域。",
	},
	_const.RustbucketCartel + "_guard_" + Neutral + "_2": {
		_const.RU:   "За порядком следим. Бесплатно. Так что - не создавай проблем.",
		_const.EN:   "We keep order here. Free of charge. So - don't cause trouble.",
		_const.ZhCN: "我们维持秩序。免费的。所以——别找麻烦。",
	},
	_const.RustbucketCartel + "_guard_" + Neutral + "_3": {
		_const.RU:   "Это место для своих. Будешь своим - поживешь. Чужим - извиняй.",
		_const.EN:   "This place is for the locals. Be local - you'll live. Be outsider - sorry.",
		_const.ZhCN: "这是给自己人的地方。当自己人——活下去。当外人——抱歉了。",
	},
	_const.RustbucketCartel + "_guard_" + Good + "_1": {
		_const.RU:   "%UserName%, здорово! Проходи, тут сегодня тихо. Расслабься.",
		_const.EN:   "%UserName%, good to see you! Come in, it's quiet today. Chill.",
		_const.ZhCN: "<span class=\"importantly\">%UserName%</span>，好啊！进来吧，今天挺安静。放松放松。",
	},
	_const.RustbucketCartel + "_guard_" + Good + "_2": {
		_const.RU:   "Слышал, ты в прошлый раз прилично наторговал. Молодец. Уважаю.",
		_const.EN:   "Heard you made some good trades last time. Good for you. Respect.",
		_const.ZhCN: "听说你上次做得不错。好样的。我佩服。",
	},
	_const.RustbucketCartel + "_guard_" + Bad + "_1": {
		_const.RU:   "Опять ты? Я же сказал - вали. Или жди неприятностей.",
		_const.EN:   "You again? I told you - get lost. Or wait for trouble.",
		_const.ZhCN: "又是你？我说了——滚。否则有麻烦等着你。",
	},
	_const.RustbucketCartel + "_guard_" + Bad + "_2": {
		_const.RU:   "Тут про тебя нехорошее говорят. Лишний раз не светись, понял?",
		_const.EN:   "People are saying bad things about you. Keep a low profile, got it?",
		_const.ZhCN: "有人说你坏话。别太显眼，明白吗？",
	},
	_const.RustbucketCartel + "_guard_" + Target + "_1": {
		_const.RU:   "Ну все, договорился. Прощай, %UserName%.",
		_const.EN:   "That's it, you asked for it. Goodbye, %UserName%.",
		_const.ZhCN: "行了，你自找的。永别了，<span class=\"importantly\">%UserName%</span>。",
	},
	_const.RustbucketCartel + "_guard_" + Target + "_2": {
		_const.RU:   "Я же просил - не надо проблем. Ты сам выбрал.",
		_const.EN:   "I asked you - no trouble. Your choice.",
		_const.ZhCN: "我提醒过你——别找麻烦。你自找的。",
	},
	_const.RustbucketCartel + "_guard_" + Target + "_3": {
		_const.RU:   "Кто первый начал - нам пофиг. Оба в расход пойдете.",
		_const.EN:   "Who started it - we don't give a damn. Both of you are going down.",
		_const.ZhCN: "谁先动手的——我们不管。两个都干掉。",
	},
	_const.RustbucketCartel + "_guard_" + Fear + "_1": {
		_const.RU:   "Ты че, попутал, %UserName%?",
		_const.EN:   "The hell you thinking, %UserName%?",
		_const.ZhCN: "你搞什么，%UserName%？",
	},
	_const.RustbucketCartel + "_guard_" + Fear + "_2": {
		_const.RU:   "Оружие убрал, пока цел. Последний раз предупреждаю.",
		_const.EN:   "Put the weapon away while you're still intact. Last warning.",
		_const.ZhCN: "把武器收起来，趁你还没事。最后一次警告。",
	},
	_const.RustbucketCartel + "_guard_" + Fear + "_3": {
		_const.RU:   "Хочешь проверить, быстро мы реагируем? Ну давай, рискни.",
		_const.EN:   "Want to test how fast we react? Go ahead, take a risk.",
		_const.ZhCN: "想试试我们反应多快？来啊，冒个险。",
	},
	_const.RustbucketCartel + "_guard_" + Help + "_1": {
		_const.RU:   "Братва, нужна подмога! Тут один умник решил пошуметь.",
		_const.EN:   "Brothers, need backup! One wise guy decided to make noise here.",
		_const.ZhCN: "兄弟们，需要支援！有个自作聪明的家伙想闹事。",
	},
	_const.RustbucketCartel + "_guard_" + Help + "_2": {
		_const.RU:   "Все на меня! Тут клиент нервный, будем успокаивать.",
		_const.EN:   "Everyone on me! Got a nervous client here, need to calm him down.",
		_const.ZhCN: "都跟我来！有个紧张的家伙，得给他冷静一下。",
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
