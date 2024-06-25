package handbook

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"math/rand"
	"strings"
	"time"
)

var mailTemplates = map[string]map[string]map[string]string{
	"common_craft_1": {
		_const.RU: {
			"sender":  `База: `,
			"subject": `Завершено производство`,
			"body":    `Завершено производство предметов на базе %BASE_NAME%, все предметы отправлены на склад.`,
		},
		_const.EN: {
			"sender":  `Base: `,
			"subject": `Production completed`,
			"body":    `The production of items based on %BASE_NAME% has been completed, all items have been sent to the warehouse.`,
		},
	},
	"common_sell_1": {
		_const.RU: {
			"sender":  `База: `,
			"subject": `Совершена продажа`,
			"body":    `Был продан предмет на базе %BASE_NAME%. <br><br> Сумма следки %CREDITS% cr. кредиты отправлены на ваш счет. <br><br> Проданные предметы:`,
		},
		_const.EN: {
			"sender":  `Base: `,
			"subject": `Sale completed`,
			"body":    `An item based on %BASE_NAME% was sold. <br><br> Trading amount %CREDITS% cr. credits have been sent to your account. <br><br> Items sold:`,
		},
	},
	"common_buy_1": {
		_const.RU: {
			"sender":  `База: `,
			"subject": `Совершена покупка`,
			"body":    `Был куплен предмет на базе %BASE_NAME%, за %CREDITS% cr. <br><br> Купленные предметы ожидают вас на складе базы:`,
		},
		_const.EN: {
			"sender":  `Base: `,
			"subject": `Purchase completed`,
			"body":    `An item was purchased on the base %BASE_NAME%, for %CREDITS% cr. <br><br> The purchased items are waiting for you in the base warehouse:`,
		},
	},
	"common_translation_1": {
		_const.RU: {
			"sender":  `Перевод`,
			"subject": `-`,
			"body":    `Игрок %FROM_PLAYER_NAME% перевел вам %CREDITS% cr.`,
		},
		_const.EN: {
			"sender":  `Transfer`,
			"subject": `-`,
			"body":    `Player %FROM_PLAYER_NAME% transferred you %CREDITS% cr.`,
		},
	},
	"corporation_create_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Поздравляем!`,
			"body":    `Поздравляем, вы создали кластер!`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Congratulations!`,
			"body":    `Congratulations, you have created a cluster!`,
		},
	},
	"corporation_rental_sector_start_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Захват сектора`,
			"body":    `Ваш кластер закрепил за собой сектор %SECTOR_NAME%, у вас есть 24 часа что бы закончить постройку или база станет уязвимой.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Sector capture`,
			"body":    `Your cluster has secured the sector %SECTOR_NAME%, you have 24 hours to complete construction or the base will become vulnerable.`,
		},
	},
	"corporation_rental_sector_complete_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Новая база`,
			"body":    `Ваш кластер построил новую базу в секторе %SECTOR_NAME%.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `New base`,
			"body":    `Your cluster has built a new base in sector %SECTOR_NAME%.`,
		},
	},
	"corporation_get_credits_sector_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Заправка базы`,
			"body":    `%PLAYER_NAME% заправил базу в секторе %SECTOR_NAME%, база сможет функционировать до: %DATE_TIME%.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Refueling the base`,
			"body":    `%PLAYER_NAME% has filled up the database in sector %SECTOR_NAME%, the database will be able to function until: %DATE_TIME%.`,
		},
	},
	"corporation_siege_start_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Начало осады`,
			"body":    `Кластер начал осаду сектора %SECTOR_NAME%.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Beginning of the siege`,
			"body":    `The cluster has begun a siege on sector %SECTOR_NAME%.`,
		},
	},
	"corporation_siege_defend_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Наш сектор осаждают!`,
			"body":    `%CORPORATION_NAME% начал осаду нашего сектора %SECTOR_NAME%`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Our sector is under siege!`,
			"body":    `%CORPORATION_NAME% has begun a siege on our sector %SECTOR_NAME%`,
		},
	},
	"corporation_siege_destroy_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `База уничтожена`,
			"body":    `Наша база в секторе %SECTOR_NAME% уничтожена.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Base destroyed`,
			"body":    `Our base in the sector %SECTOR_NAME% has been destroyed.`,
		},
	},
	"corporation_war_start_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Обьявление войны`,
			"body":    `Мы обьявили войну %CORPORATION_NAME%.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Declaration of war`,
			"body":    `We have declared war on %CORPORATION_NAME%.`,
		},
	},
	"corporation_war_defend_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Обьявление войны`,
			"body":    `%CORPORATION_NAME% обьявил нам войну!`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Declaration of war`,
			"body":    `%CORPORATION_NAME% has declared war on us!`,
		},
	},
	"corporation_war_mutual_on_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Взаимная война`,
			"body":    `%CORPORATION_NAME% установил статус взаимной войны с вами. Вам больше ненадо платить взнос.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Mutual war`,
			"body":    `%CORPORATION_NAME% has set the status of mutual war with you. You no longer have to pay a fee.`,
		},
	},
	"corporation_war_mutual_off_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Взаимная война`,
			"body":    `%CORPORATION_NAME% убрал статус взаимной войны с вами. Что бы продолжить воевать необходимо заплотить нолог.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Mutual war`,
			"body":    `%CORPORATION_NAME% убрал статус взаимной войны с вами. Что бы продолжить воевать необходимо заплотить налог.`,
		},
	},
	"corporation_war_surrender_on_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Предложение мира`,
			"body":    `%CORPORATION_NAME% предлогает мир.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Peace Offer`,
			"body":    `%CORPORATION_NAME% offers peace.`,
		},
	},
	"corporation_war_surrender_off_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Предложение мира`,
			"body":    `%CORPORATION_NAME% уберает предложение о мире.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Peace Offer`,
			"body":    `%CORPORATION_NAME% removes the peace proposal.`,
		},
	},
	"corporation_endwar_hostile_headquarters_destroy_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Штаб врага уничтожен`,
			"body":    `Штаб врага уничтожен, война с %CORPORATION_NAME% завершается.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Enemy headquarters destroyed`,
			"body":    `The enemy's headquarters is destroyed, the war with %CORPORATION_NAME% ends.`,
		},
	},
	"corporation_endwar_union_headquarters_destroy_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Штаб уничтожен`,
			"body":    `Наш штаб в войне против %CORPORATION_NAME% был уничтожен, война завершается.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Headquarters destroyed`,
			"body":    `Our headquarters in the war against %CORPORATION_NAME% was destroyed, the war is ending.`,
		},
	},
	"corporation_endwar_hostile_no_base_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Все базы врага уничтожены`,
			"body":    `Все базы врага уничтожены, война против %CORPORATION_NAME% завершается.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `All enemy bases have been destroyed`,
			"body":    `All enemy bases have been destroyed, the war against %CORPORATION_NAME% ends.`,
		},
	},
	"corporation_endwar_union_no_base_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Все базы уничтожены.`,
			"body":    `Враг уничтожил все наши базы, все наши войны завершаются.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `All bases have been destroyed.`,
			"body":    `The enemy has destroyed all our bases, all our wars are ending.`,
		},
	},
	"corporation_endwar_surrender_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Заключено перемирие.`,
			"body":    `Мы заключили мир с кластером %CORPORATION_NAME%, 14 дней мы не сможем им обьявить войну.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `A truce has been concluded.`,
			"body":    `We have made peace with the %CORPORATION_NAME% cluster; for 14 days we will not be able to declare war on them.`,
		},
	},
	"corporation_rental_office_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Аренда офиса`,
			"body":    `%PLAYER_NAME% арендовал офис, на базе %BASE_NAME%, до: %DATE_TIME%.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Office rental`,
			"body":    `%PLAYER_NAME% rented an office, based on %BASE_NAME%, until: %DATE_TIME%.`,
		},
	},
	"corporation_get_credits_office_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Аренда офиса`,
			"body":    `%PLAYER_NAME% продлили аренду офиса, на базе %BASE_NAME%, до: %DATE_TIME%.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Office rental`,
			"body":    `%PLAYER_NAME% extended their office lease, based on %BASE_NAME%, to: %DATE_TIME%.`,
		},
	},
	"corporation_no_credits_rental_office_1": { // TODO
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Просрочен платеж`,
			"body":    `Офис на базе %BASE_NAME% перестает функционировать до аренды, все имущество на складах арестовано.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Payment overdue`,
			"body":    `The office based on %BASE_NAME% ceases to function before rent, all property in the warehouses is seized.`,
		},
	},
}

func GetMailTemplates(typeEvent string) map[string]map[string]string {
	textVariants := make([]map[string]map[string]string, 0)
	rand.Seed(time.Now().UnixNano())

	for key, text := range mailTemplates {
		if strings.Contains(key, typeEvent) {
			textVariants = append(textVariants, text)
		}
	}

	return textVariants[rand.Intn(len(textVariants))]
}
