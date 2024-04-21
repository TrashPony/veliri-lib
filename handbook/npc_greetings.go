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
	_const.Replicas + "_miner_" + Neutral + "_1": {
		_const.RU: "Привет, %UserName%. Я иду копать руду, не мешай мне.",
		_const.EN: "Hello, %UserName%. I'm going to dig for ore, don't bother me.",
	},
	_const.Replicas + "_guard_" + Neutral + "_1": {
		_const.RU: "Проезжай, не задерживайся.",
		_const.EN: "Pass through, don't stop",
	},
	_const.Replicas + "_out_scout_" + Neutral + "_1": {
		_const.RU: "Привет, я занимаюсь выполнением правительственных миссий, не мешай.",
		_const.EN: "Hi, I'm on a government mission, don't interfere.",
	},
	_const.Replicas + "_in_scout_" + Neutral + "_1": {
		_const.RU: "Чего тебе?",
		_const.EN: "What do you want?",
	},
	_const.Replicas + "_transport_" + Neutral + "_1": {
		_const.RU: "Не мешай, я занимаюсь торговлей",
		_const.EN: "Do not disturb, I'm doing business",
	},
	_const.Replicas + "_builder_" + Neutral + "_1": {
		_const.RU: "Опять кто-то что-то сломал, а мне чинить…",
		_const.EN: "Someone has broken something again, and I have to fix it…",
	},
	_const.Replicas + "_warrior_" + Neutral + "_1": {
		_const.RU: "Не мешай выполнению военных обязанностей",
		_const.EN: "Do not interfere with my military duties",
	},
	_const.Replicas + "_expedition_" + Neutral + "_1": {
		_const.RU: "Наша задача — осваивать пустоши, не мешай нам и будешь цел",
		_const.EN: "Our task is to explore the wasteland, do not disturb us and you will be safe",
	},

	_const.Replicas + "_miner_" + Good + "_1": {
		_const.RU: "Кажется, я нашёл руду! Хочешь помочь мне её разработать?",
		_const.EN: "I think I've found some ore! Want to help me mine it?",
	},
	_const.Replicas + "_miner_" + Good + "_2": {
		_const.RU: "Похоже, в этом секторе уже выкопали всю руду, лучше поискать её где-то ещё",
		_const.EN: "It seems like all the ore has been mined in this sector, better look somewhere else",
	},
	_const.Replicas + "_guard_" + Good + "_1": {
		_const.RU: "Вас что-то беспокоит?",
		_const.EN: "Is something bothering you?",
	},
	_const.Replicas + "_out_scout_" + Good + "_1": {
		_const.RU: "Кажется, в пустошах стало слишком много враждебных ботов, будь осторожен!",
		_const.EN: "There seem to be a lot of hostile bots in the wasteland, be careful!",
	},
	_const.Replicas + "_in_scout_" + Good + "_1": {
		_const.RU: "Может, храбанем кого-нибудь?",
		_const.EN: "Maybe we'll rob someone?",
	},
	_const.Replicas + "_transport_" + Good + "_1": {
		_const.RU: "Я везу %ProductName% на базу %BaseName%, это будет выгодная сделка!",
		_const.EN: "I'm taking %ProductName% to the %BaseName% base, it'll be a profitable deal!",
	},
	_const.Replicas + "_transport_" + Good + "_2": {
		_const.RU: "Я следую на %BaseName%, база очень выгодно продаёт %ProductName%!",
		_const.EN: "I'm heading to %BaseName%, the base sells %ProductName% very profitably!",
	},
	_const.Replicas + "_transport_" + Good + "_3": {
		_const.RU: "Привет, %UserName%, как дела?",
		_const.EN: "Hello, %UserName%, how are you?",
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

	// Explores
	_const.Explores + "_miner_" + Neutral + "_1": {
		_const.RU: "Привет %UserName%. Я иду копать руду, не мешай мне.",
		_const.EN: "Hello %UserName%. I'm going to mine ore, don't disturb me.",
	},
	_const.Explores + "_guard_" + Neutral + "_1": {
		_const.RU: "Проезжай, не задерживайся.",
		_const.EN: "Move on, don't stop.",
	},
	_const.Explores + "_out_scout_" + Neutral + "_1": {
		_const.RU: "Привет, я занимаюсь какой-то важной ерундой, не мешай.",
		_const.EN: "Hi, I'm doing something important, don't bother me.",
	},
	_const.Explores + "_in_scout_" + Neutral + "_1": {
		_const.RU: "Чего тебе?",
		_const.EN: "What do you want?",
	},
	_const.Explores + "_transport_" + Neutral + "_1": {
		_const.RU: "Не мешай, я занимаюсь торговлей",
		_const.EN: "Do not disturb me, I'm doing some trading",
	},
	_const.Explores + "_builder_" + Neutral + "_1": {
		_const.RU: "Опять кто-то что-то сломал, а мне чинить...",
		_const.EN: "Once again, someone broke something, and I have to fix it...",
	},
	_const.Explores + "_warrior_" + Neutral + "_1": {
		_const.RU: "Не мешай выполнению военных обязанностей.",
		_const.EN: "Do not interfere with my military duties.",
	},
	_const.Explores + "_expedition_" + Neutral + "_1": {
		_const.RU: "Наша задача — осваивать пустоши, не мешай нам и будешь цел.",
		_const.EN: "Our task is to explore the wasteland, do not interfere with us and you will be safe.",
	},

	_const.Explores + "_miner_" + Good + "_1": {
		_const.RU: "Кажется, я нашел руду! Хочешь помочь мне её разработать?",
		_const.EN: "I think I found some ore! Do you want to help me mine it?",
	},
	_const.Explores + "_miner_" + Good + "_2": {
		_const.RU: "Похоже, в этом секторе уже выкопали всю руду, лучше поискать её где-то ещё.",
		_const.EN: "It seems like all the ore has already been mined in this sector, it's better to look for it somewhere else.",
	},
	_const.Explores + "_guard_" + Good + "_1": {
		_const.RU: "Вас что-то беспокоит?",
		_const.EN: "Is there anything bothering you?",
	},
	_const.Explores + "_out_scout_" + Good + "_1": {
		_const.RU: "Кажется, в пустошах стало слишком много враждебных ботов, будь осторожен!",
		_const.EN: "It seems like there are too many hostile bots in the wasteland, be careful!",
	},
	_const.Explores + "_in_scout_" + Good + "_1": {
		_const.RU: "Может, храбанем кого-нибудь?",
		_const.EN: "Maybe we'll rob someone?",
	},
	_const.Explores + "_transport_" + Good + "_1": {
		_const.RU: "Я везу %ProductName% на базу %BaseName%, это будет выгодная сделка!",
		_const.EN: "I'm taking %ProductName% to the %BaseName% base, it will be a profitable deal!",
	},
	_const.Explores + "_transport_" + Good + "_2": {
		_const.RU: "Я следую на %BaseName%, база очень выгодно продаёт %ProductName%!",
		_const.EN: "I'm heading to %BaseName%, the base sells %ProductName% very profitably!",
	},
	_const.Explores + "_transport_" + Good + "_3": {
		_const.RU: "Привет %UserName%, как дела?",
		_const.EN: "Hi %UserName%, how are you?",
	},
	_const.Explores + "_builder_" + Good + "_1": {
		_const.RU: "Привет, как дела?",
		_const.EN: "Hello, how are you?",
	},
	_const.Explores + "_warrior_" + Good + "_1": {
		_const.RU: "Мы отправляемся в сектор %SectorName%, можешь присоединиться!",
		_const.EN: "We're going to sector %SectorName%, you can join us!",
	},
	_const.Explores + "_expedition_" + Good + "_1": {
		_const.RU: "Мы пробуем освоить секторе %SectorName%, будем рады содействию!",
		_const.EN: "We are trying to explore the sector %SectorName%, we will be glad to cooperate!",
	},
	_const.Explores + "_expedition_" + Good + "_2": {
		_const.RU: "Наша задача выбирать противника с территории пустошь в секторе %SectorName%! Будем рады содействию!",
		_const.EN: "Our task is to choose an enemy from the wasteland territory in the sector %SectorName%! We will be glad to cooperate!",
	},

	_const.Explores + "_miner_" + Bad + "_1": {
		_const.RU: "Ну, иди сюда, попробуй меня трахнуть — я тебя сам трахну!",
		_const.EN: "Well, come here, try to fuck me - I’ll fuck you myself!",
	},
	_const.Explores + "_guard_" + Bad + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "Don't make things worse!",
	},
	_const.Explores + "_out_scout_" + Bad + "_1": {
		_const.RU: "Лучше уходи пока можешь!",
		_const.EN: "Better leave while you can!",
	},
	_const.Explores + "_in_scout_" + Bad + "_1": {
		_const.RU: "Приветствую. Хочешь внести плату за проезд?",
		_const.EN: "Greetings. Do you want to pay the fare?",
	},
	_const.Explores + "_transport_" + Bad + "_1": {
		_const.RU: "Если ты помешаешь мне торговать то я за себя не ручаюсь!",
		_const.EN: "If you interfere with my trade, I won't be responsible for my actions!",
	},
	_const.Explores + "_builder_" + Bad + "_1": {
		_const.RU: "Отвали, ты мне не интересен.",
		_const.EN: "Get lost, you're not interesting to me.",
	},
	_const.Explores + "_warrior_" + Bad + "_1": {
		_const.RU: "Советую тебе избегай встречи со мной!",
		_const.EN: "I advise you to avoid meeting me!",
	},
	_const.Explores + "_expedition_" + Bad + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "Stay away!",
	},

	_const.Explores + "_miner_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your vehicle!",
	},
	_const.Explores + "_guard_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your vehicle!",
	},
	_const.Explores + "_out_scout_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your vehicle!",
	},
	_const.Explores + "_in_scout_" + Target + "_1": {
		_const.RU: "Уже сдаешься? Это не бесплатная услуга!",
		_const.EN: "Are you giving up already? This is not a free service!",
	},
	_const.Explores + "_transport_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your vehicle!",
	},
	_const.Explores + "_builder_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your vehicle!",
	},
	_const.Explores + "_warrior_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your vehicle!",
	},
	_const.Explores + "_expedition_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your vehicle!",
	},

	_const.Explores + "_miner_" + Fear + "_1": {
		_const.RU: "Прекрати! Чего тебе надо?!",
		_const.EN: "Stop it! What do you want?!",
	},
	_const.Explores + "_guard_" + Fear + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "Don't make things worse!",
	},
	_const.Explores + "_out_scout_" + Fear + "_1": {
		_const.RU: "А знаешь что? У меня есть дела поважнее!",
		_const.EN: "You know what? I have more important things to do!",
	},
	_const.Explores + "_in_scout_" + Fear + "_1": {
		_const.RU: "Я всё понял! Пожалуйста, не разрушай мой транспорт!",
		_const.EN: "I've got it! Please, don't destroy my vehicle!",
	},
	_const.Explores + "_transport_" + Fear + "_1": {
		_const.RU: "Может, мы как-то договоримся?!",
		_const.EN: "Maybe we can come to some agreement?!",
	},
	_const.Explores + "_builder_" + Fear + "_1": {
		_const.RU: "Может, мы как-то договоримся?!",
		_const.EN: "Maybe we can come to some agreement?!",
	},
	_const.Explores + "_warrior_" + Fear + "_1": {
		_const.RU: "Советую тебе избегать встречи со мной!",
		_const.EN: "I advise you to avoid meeting me!",
	},
	_const.Explores + "_expedition_" + Fear + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "Stay away!",
	},

	_const.Explores + "_miner_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "Help, the hooligans are depriving me of my sight!",
	},
	_const.Explores + "_guard_" + Help + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "Don't make things worse!",
	},
	_const.Explores + "_out_scout_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "Help, the hooligans are depriving me of my sight!",
	},
	_const.Explores + "_in_scout_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "Help, the hooligans are depriving me of my sight!",
	},
	_const.Explores + "_transport_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "Help, the hooligans are depriving me of my sight!",
	},
	_const.Explores + "_builder_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "Help, the hooligans are depriving me of my sight!",
	},
	_const.Explores + "_warrior_" + Help + "_1": {
		_const.RU: "Советую тебе избегать встречи со мной!",
		_const.EN: "I advise you to avoid meeting me!",
	},
	_const.Explores + "_expedition_" + Help + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "Stay away!",
	},

	// Reverses Neutral
	_const.Reverses + "_miner_" + Neutral + "_1": {
		_const.RU: "Привет %UserName%. Я иду копать руду, не мешай мне.",
		_const.EN: "Hello, %UserName%. I'm going to mine ore, don't disturb me.",
	},
	_const.Reverses + "_guard_" + Neutral + "_1": {
		_const.RU: "Проезжай, не задерживайся.",
		_const.EN: "Move on, don't stop.",
	},
	_const.Reverses + "_out_scout_" + Neutral + "_1": {
		_const.RU: "Привет, я занимаюсь какой-то важной херней, не мешай.",
		_const.EN: "Hi, I'm doing some important shit, don't bother me.",
	},
	_const.Reverses + "_in_scout_" + Neutral + "_1": {
		_const.RU: "Чего тебе?",
		_const.EN: "What do you want?",
	},
	_const.Reverses + "_transport_" + Neutral + "_1": {
		_const.RU: "Не мешай, я занимаюсь торговлей",
		_const.EN: "Do not disturb me, I'm doing business.",
	},
	_const.Reverses + "_builder_" + Neutral + "_1": {
		_const.RU: "Опять кто-то что-то сломал, а мне чинить...",
		_const.EN: "Someone has broken something again, and I have to fix it...",
	},
	_const.Reverses + "_warrior_" + Neutral + "_1": {
		_const.RU: "Не мешай выполнению военных обязанностей",
		_const.EN: "Do not interfere with my military duties.",
	},
	_const.Reverses + "_expedition_" + Neutral + "_1": {
		_const.RU: "Наша задача осваивать пустоши, не мешай нам и будешь цел",
		_const.EN: "Our task is to explore the wasteland, do not interfere with us and you will be safe.",
	},

	_const.Reverses + "_miner_" + Good + "_1": {
		_const.RU: "Кажется, я нашёл руду! Хочешь помочь мне её разработать?",
		_const.EN: "I think I've found some ore! Do you want to help me mine it?",
	},
	_const.Reverses + "_miner_" + Good + "_2": {
		_const.RU: "Похоже в этом секторе уже выкопали всю руду, лучше поискать её где-то ещё",
		_const.EN: "It seems that all the ore has been mined in this sector, it is better to look for it somewhere else.",
	},
	_const.Reverses + "_guard_" + Good + "_1": {
		_const.RU: "Вас что-то беспокоит?",
		_const.EN: "Is something bothering you?",
	},
	_const.Reverses + "_out_scout_" + Good + "_1": {
		_const.RU: "Кажется, в пустошах стало слишком много враждебных ботов, будь осторожен!",
		_const.EN: "It seems there are too many hostile bots in the wasteland, be careful!",
	},
	_const.Reverses + "_in_scout_" + Good + "_1": {
		_const.RU: "Может, храбанем кого-нибудь?",
		_const.EN: "Maybe we'll rob someone?",
	},
	_const.Reverses + "_transport_" + Good + "_1": {
		_const.RU: "Я везу %ProductName% на базу %BaseName%, это будет выгодная сделка!",
		_const.EN: "I'm carrying %ProductName% to the %BaseName% base, it's going to be a profitable deal!",
	},
	_const.Reverses + "_transport_" + Good + "_2": {
		_const.RU: "Я следую на %BaseName%, база очень выгодно продаёт %ProductName%! ",
		_const.EN: "I'm heading to %BaseName%, the base sells %ProductName% very profitably!",
	},
	_const.Reverses + "_transport_" + Good + "_3": {
		_const.RU: "Привет %UserName%, как дела?",
		_const.EN: "Hi, %UserName%, how are you doing?",
	},
	_const.Reverses + "_builder_" + Good + "_1": {
		_const.RU: "Привет, как дела?",
		_const.EN: "Hello, how are you?",
	},
	_const.Reverses + "_warrior_" + Good + "_1": {
		_const.RU: "Мы отправляемся в сектор %SectorName%, можешь присоединиться!",
		_const.EN: "We're heading to the %SectorName% sector, you can join us!",
	},
	_const.Reverses + "_expedition_" + Good + "_1": {
		_const.RU: "Мы пробуем освоить сектор %SectorName%, будем рады содействию!",
		_const.EN: "We are trying to explore the %SectorName% sector, we will appreciate your help!",
	},
	_const.Reverses + "_expedition_" + Good + "_2": {
		_const.RU: "Наша задача — вытеснить противника с территории пустоши в секторе %SectorName%! Будем рады содействию!",
		_const.EN: "Our task is to oust the enemy from the wasteland territory in the %SectorName% sector! We will appreciate your assistance!",
	},

	_const.Reverses + "_miner_" + Bad + "_1": {
		_const.RU: "Ну, иди сюда, попробуй меня трахнуть — я тебя сам трахну!",
		_const.EN: "Well, come here, try to fuck me - I’ll fuck you myself!\n",
	},
	_const.Reverses + "_guard_" + Bad + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "Do not make things worse!",
	},
	_const.Reverses + "_out_scout_" + Bad + "_1": {
		_const.RU: "Лучше уходи, пока можешь!",
		_const.EN: "Better leave while you can!",
	},
	_const.Reverses + "_in_scout_" + Bad + "_1": {
		_const.RU: "Приветствую. Хочешь внести плату за проезд?",
		_const.EN: "Greetings. Do you want to pay the fare?",
	},
	_const.Reverses + "_transport_" + Bad + "_1": {
		_const.RU: "Если ты помешаешь мне торговать, то я за себя не ручаюсь!",
		_const.EN: "If you interfere with my trade, I won't be responsible for my actions!",
	},
	_const.Reverses + "_builder_" + Bad + "_1": {
		_const.RU: "Отвали, ты мне не интересен.",
		_const.EN: "Get lost, you're not interesting to me.",
	},
	_const.Reverses + "_warrior_" + Bad + "_1": {
		_const.RU: "Советую тебе избегать встречи со мной!",
		_const.EN: "I advise you to avoid meeting me!",
	},
	_const.Reverses + "_expedition_" + Bad + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "Do not come closer!",
	},

	_const.Reverses + "_miner_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your transport!",
	},
	_const.Reverses + "_guard_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your transport!",
	},
	_const.Reverses + "_out_scout_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your transport!",
	},
	_const.Reverses + "_in_scout_" + Target + "_1": {
		_const.RU: "Уже сдаешься? Это не бесплатная услуга!",
		_const.EN: "Are you giving up already? This is not a free service!",
	},
	_const.Reverses + "_transport_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your transport!",
	},
	_const.Reverses + "_builder_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your transport!",
	},
	_const.Reverses + "_warrior_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your transport!",
	},
	_const.Reverses + "_expedition_" + Target + "_1": {
		_const.RU: "Сейчас я уничтожу твой транспорт!",
		_const.EN: "I'm going to destroy your transport!",
	},

	_const.Reverses + "_miner_" + Fear + "_1": {
		_const.RU: "Прекрати! Чего тебе надо?!",
		_const.EN: "Stop it! What do you want?!",
	},
	_const.Reverses + "_guard_" + Fear + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "Don't make things worse!",
	},
	_const.Reverses + "_out_scout_" + Fear + "_1": {
		_const.RU: "А знаешь что? У меня есть дела поважнее!",
		_const.EN: "You know what? I have more important things to do!",
	},
	_const.Reverses + "_in_scout_" + Fear + "_1": {
		_const.RU: "Я всё понял! Пожалуйста, не разрушай мой транспорт!",
		_const.EN: "I get it! Please don't destroy my transport!",
	},
	_const.Reverses + "_transport_" + Fear + "_1": {
		_const.RU: "Может, мы как-то договоримся?!",
		_const.EN: "Maybe we can come to some agreement?!",
	},
	_const.Reverses + "_builder_" + Fear + "_1": {
		_const.RU: "Может, мы как-то договоримся?!",
		_const.EN: "Maybe we can come to some agreement?!",
	},
	_const.Reverses + "_warrior_" + Fear + "_1": {
		_const.RU: "Советую тебе избегать встречи со мной!",
		_const.EN: "I advise you to avoid meeting me!",
	},
	_const.Reverses + "_expedition_" + Fear + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "Do not come closer!",
	},

	_const.Reverses + "_miner_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "Help, hooligans are depriving me of my sight!",
	},
	_const.Reverses + "_guard_" + Help + "_1": {
		_const.RU: "Не усугубляй!",
		_const.EN: "Don't make things worse!",
	},
	_const.Reverses + "_out_scout_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "Help, hooligans are depriving me of my sight!",
	},
	_const.Reverses + "_in_scout_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "Help, hooligans are depriving me of my sight!",
	},
	_const.Reverses + "_transport_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "Help, hooligans are depriving me of my sight!",
	},
	_const.Reverses + "_builder_" + Help + "_1": {
		_const.RU: "Помогите, хулиганы зрения лишают!",
		_const.EN: "Help, hooligans are depriving me of my sight!",
	},
	_const.Reverses + "_warrior_" + Help + "_1": {
		_const.RU: "Советую тебе избегать встречи со мной!",
		_const.EN: "I advise you to avoid meeting me!",
	},
	_const.Reverses + "_expedition_" + Help + "_1": {
		_const.RU: "Не приближайся!",
		_const.EN: "Do not come closer!",
	},

	// APD
	_const.APD + "___1": {
		_const.RU: "ПУ ПУ ПУ...",
		_const.EN: "PU PU PU...",
	},

	// FarGod
	_const.FarGod + "___1": {
		_const.RU: "Скоро всё изменится!",
		_const.EN: "Soon everything will change!",
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
