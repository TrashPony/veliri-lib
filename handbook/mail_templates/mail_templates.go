package mail_templates

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
		_const.ZhCN: {
			"sender":  `基地: `,
			"subject": `生产完成`,
			"body":    `在基地 %BASE_NAME% 的物品生产已完成，所有物品已发送到仓库。`,
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
		_const.ZhCN: {
			"sender":  `基地: `,
			"subject": `出售完成`,
			"body":    `在基地 %BASE_NAME% 的物品已出售。 <br><br> 交易金额 %CREDITS% cr. 信用点已发送到您的账户。 <br><br> 出售的物品:`,
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
		_const.ZhCN: {
			"sender":  `基地: `,
			"subject": `购买完成`,
			"body":    `在基地 %BASE_NAME% 购买了物品，花费 %CREDITS% cr. <br><br> 购买的物品正在基地仓库等待您:`,
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
		_const.ZhCN: {
			"sender":  `转账`,
			"subject": `-`,
			"body":    `玩家 %FROM_PLAYER_NAME% 向您转账 %CREDITS% cr.`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `恭喜！`,
			"body":    `恭喜，您已创建了一个集群！`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `区域占领`,
			"body":    `您的集群已占领区域 %SECTOR_NAME%，您有 24 小时完成建设，否则基地将变得脆弱。`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `新基地`,
			"body":    `您的集群在区域 %SECTOR_NAME% 建造了一个新基地。`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `基地加油`,
			"body":    `%PLAYER_NAME% 为区域 %SECTOR_NAME% 的基地加油，基地将能够运行至: %DATE_TIME%。`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `围攻开始`,
			"body":    `集群已开始围攻区域 %SECTOR_NAME%。`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `我们的区域被围攻！`,
			"body":    `%CORPORATION_NAME% 已开始围攻我们的区域 %SECTOR_NAME%`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `基地被摧毁`,
			"body":    `我们在区域 %SECTOR_NAME% 的基地已被摧毁。`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `宣战`,
			"body":    `我们已向 %CORPORATION_NAME% 宣战。`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `宣战`,
			"body":    `%CORPORATION_NAME% 已向我们宣战！`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `相互战争`,
			"body":    `%CORPORATION_NAME% 已与您设置为相互战争状态。您不再需要支付费用。`,
		},
	},
	"corporation_war_mutual_off_1": {
		_const.RU: {
			"sender":  `Кластер`,
			"subject": `Взаимная война`,
			"body":    `%CORPORATION_NAME% убрал статус взаимной войны с вами. Что бы продолжить воевать необходимо заплотить налог.`,
		},
		_const.EN: {
			"sender":  `Cluster`,
			"subject": `Mutual war`,
			"body":    `%CORPORATION_NAME% has removed the status of mutual war with you. To continue fighting, you must pay a tax.`,
		},
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `相互战争`,
			"body":    `%CORPORATION_NAME% 已移除与您的相互战争状态。要继续战斗，您必须缴纳税款。`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `和平提议`,
			"body":    `%CORPORATION_NAME% 提出和平提议。`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `和平提议`,
			"body":    `%CORPORATION_NAME% 移除了和平提议。`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `敌方总部被摧毁`,
			"body":    `敌方总部已被摧毁，与 %CORPORATION_NAME% 的战争结束。`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `总部被摧毁`,
			"body":    `我们在与 %CORPORATION_NAME% 的战争中的总部已被摧毁，战争即将结束。`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `所有敌方基地被摧毁`,
			"body":    `所有敌方基地已被摧毁，与 %CORPORATION_NAME% 的战争结束。`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `所有基地被摧毁`,
			"body":    `敌人摧毁了我们所有的基地，我们所有的战争即将结束。`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `已达成停战协议`,
			"body":    `我们已与集群 %CORPORATION_NAME% 达成和平，14天内无法对其宣战。`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `办公室租赁`,
			"body":    `%PLAYER_NAME% 在基地 %BASE_NAME% 租赁了办公室，有效期至: %DATE_TIME%。`,
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `办公室租赁`,
			"body":    `%PLAYER_NAME% 延长了在基地 %BASE_NAME% 的办公室租赁，有效期至: %DATE_TIME%。`,
		},
	},
	"corporation_no_credits_rental_office_1": {
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
		_const.ZhCN: {
			"sender":  `集群`,
			"subject": `逾期付款`,
			"body":    `基地 %BASE_NAME% 的办公室因未支付租金已停止运作，仓库内所有财产已被扣押。`,
		},
	},
	"corporation_application_created_1": {
		_const.RU: {
			"sender":  "Кластер",
			"subject": "Заявка подана",
			"body":    "Ваша заявка на вступление в кластер %CORPORATION_NAME% успешно подана и ожидает рассмотрения.",
		},
		_const.EN: {
			"sender":  "Cluster",
			"subject": "Application submitted",
			"body":    "Your application to join %CORPORATION_NAME% has been submitted and is pending review.",
		},
		_const.ZhCN: {
			"sender":  "集群",
			"subject": "申请已提交",
			"body":    "您加入 %CORPORATION_NAME% 的申请已成功提交，正在等待审核。",
		},
	},
	"corporation_application_rejected_1": {
		_const.RU: {
			"sender":  "Кластер",
			"subject": "Заявка отклонена",
			"body":    "Ваша заявка на вступление в кластер %CORPORATION_NAME% была отклонена.",
		},
		_const.EN: {
			"sender":  "Cluster",
			"subject": "Application rejected",
			"body":    "Your application to join %CORPORATION_NAME% has been rejected.",
		},
		_const.ZhCN: {
			"sender":  "集群",
			"subject": "申请被拒绝",
			"body":    "您加入 %CORPORATION_NAME% 的申请已被拒绝。",
		},
	},
	"corporation_application_accepted_1": {
		_const.RU: {
			"sender":  "Кластер",
			"subject": "Заявка одобрена",
			"body":    "Поздравляем! Вы приняты в кластер %CORPORATION_NAME%. Доступ к корпоративным ресурсам активирован.",
		},
		_const.EN: {
			"sender":  "Cluster",
			"subject": "Application approved",
			"body":    "Congratulations! You have been accepted into %CORPORATION_NAME%. Corporate resource access has been activated.",
		},
		_const.ZhCN: {
			"sender":  "集群",
			"subject": "申请已通过",
			"body":    "恭喜！您已被 %CORPORATION_NAME% 集群接受。企业资源访问权限已激活。",
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
