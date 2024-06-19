package handbook

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"math/rand"
	"strings"
	"time"
)

var corporationNews = map[string]map[string]string{
	// create
	"create_1": { // кто то создал новый кластер
		_const.RU: `Сформирован новый кластер %CORPORATION_NAME%! Мы будем следить за их прогрессом.`,
		_const.EN: `A new cluster %CORPORATION_NAME% has been formed! We will monitor their progress.`,
	},
	// rental
	"rental_sector_start_1": { // аренда сектора, установили площадку базу
		_const.RU: `Похоже %CORPORATION_NAME% решили взять под контроль %SECTOR_NAME%, они начали строить базу и сейчас должны туда привести ресурсы, следим за развитием событий.`,
		_const.EN: `It seems that %CORPORATION_NAME% decided to take control of %SECTOR_NAME%, they began to build a base and now they must bring resources there, we are monitoring the developments.`,
	},
	"rental_sector_complete_1": { // аренда сектора, закончили строительство главной базы
		_const.RU: `Поздравляем %CORPORATION_NAME% они смогли получить контроль над сектором %SECTOR_NAME% в пустошах.`,
		_const.EN: `Congratulations to %CORPORATION_NAME% for gaining control of the %SECTOR_NAME% sector in the wasteland.`,
	},
	// siege
	"siege_start_1": { // начало осады
		_const.RU: `Кластер %CORPORATION_NAME% начинает осаду сектора %SECTOR_NAME%, где расположена база корпорации %TO_CORPORATION_NAME%`,
		_const.EN: `Cluster %CORPORATION_NAME% begins a siege of the sector %SECTOR_NAME%, where the base of the corporation %TO_CORPORATION_NAME% is located`,
	},
	"siege_added_1": { // к осаде присоеденились
		_const.RU: `Кластер %CORPORATION_NAME% присоединяется к осаде сектора %SECTOR_NAME%, против кластера %TO_CORPORATION_NAME%`,
		_const.EN: `Cluster %CORPORATION_NAME% joins the siege of sector %SECTOR_NAME%, against cluster %TO_CORPORATION_NAME%`,
	},
	"siege_close_1": { // обороняющимся удалось перехватить портал
		_const.RU: `Обороняющийся кластер %CORPORATION_NAME% смог вернуть контроль над порталом в секторе %SECTOR_NAME%`,
		_const.EN: `The defending cluster %CORPORATION_NAME% was able to regain control of the portal in the sector %SECTOR_NAME%`,
	},
	"siege_destroy_1": { // база во время осады уничтожена
		_const.RU: `Кластер %CORPORATION_NAME% не смог удержать базу в секторе %SECTOR_NAME% и теперь сектор стал нейтральным.`,
		_const.EN: `The %CORPORATION_NAME% cluster was unable to maintain a base in the %SECTOR_NAME% sector and now the sector has become neutral.`,
	},
	"siege_end_1": { // база во время осады защищена
		_const.RU: `Кластер %CORPORATION_NAME% смог защитить %SECTOR_NAME%`,
		_const.EN: `Cluster %CORPORATION_NAME% was able to protect %SECTOR_NAME%`,
	},
	// war
	"war_start_1": { // начало войны кластера
		_const.RU: `Кластер %CORPORATION_NAME% объявляет войну %TO_CORPORATION_NAME%, штаб нападающих находится в секторе %SECTOR_NAME%.`,
		_const.EN: `Cluster %CORPORATION_NAME% declares war on %TO_CORPORATION_NAME%, the attackers' headquarters are located in sector %SECTOR_NAME%.`,
	},
	"war_mutual_1": { // обороняющая сторона конфликта обьявляет войну взаимной
		_const.RU: `Кластер %CORPORATION_NAME% объявляет войну против %TO_CORPORATION_NAME% взаимной`,
		_const.EN: `Cluster %CORPORATION_NAME% declares mutual war against %TO_CORPORATION_NAME%`,
	},
	// endwar
	"endwar_surrender_1": { // война закончилось т.к. обе стороны согласились на мир
		_const.RU: `Кластер %CORPORATION_NAME% и %TO_CORPORATION_NAME% смогли заключит мир, и война между ними заканчивается.`,
		_const.EN: `Cluster %CORPORATION_NAME% and %TO_CORPORATION_NAME% were able to make peace, and the war between them ends.`,
	},
	"endwar_headquarters_destroy_1": { // война закончилось т.к. обороняющая сторона смогла уничтожить штаб врага
		_const.RU: `Кластер %CORPORATION_NAME% смог уничтожить штаб противника в войне против них, чем завершает войну с %TO_CORPORATION_NAME%`,
		_const.EN: `Cluster %CORPORATION_NAME% was able to destroy the enemy headquarters in the war against them, ending the war with %TO_CORPORATION_NAME%`,
	},
	"endwar_no_base_1": { // война закончилось т.к. обороняющая сторона потеряла все свои сектора
		_const.RU: `Кластер %CORPORATION_NAME% потерял все свои территории и поэтому больше не может продолжать войны против других кластеров. Все войны с этим кластером завершаются.`,
		_const.EN: `Cluster %CORPORATION_NAME% has all its territories and therefore can no longer continue wars against other clusters. All wars with this cluster end.`,
	},
	"endwar_no_credits_1": { // война закончилось т.к. нападающие не смогли оплатить взнос
		_const.RU: `Кластер %CORPORATION_NAME% не смог оплатить взнос поэтому его война против %TO_CORPORATION_NAME% завершается, и последующие войны будут вне закона.`,
		_const.EN: `Cluster %CORPORATION_NAME% was unable to pay the fee, so its war against %TO_CORPORATION_NAME% ends, and subsequent wars will be illegal.`,
	},
}

func GetCorporationNews(typeEvent string) map[string]string {
	textVariants := make([]map[string]string, 0)
	rand.Seed(time.Now().UnixNano())

	for key, text := range corporationNews {
		if strings.Contains(key, typeEvent) {
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
