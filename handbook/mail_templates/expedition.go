package mail_templates

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"math/rand"
	"strings"
	"time"
)

var expeditionMailTemplates = map[string]map[string]map[string]string{
	"exit_expedition": {
		_const.RU: {
			"body": `Покинуть объект`,
		},
		_const.EN: {
			"body": `Leave the facility`,
		},
		_const.ZhCN: {
			"body": `离开设施`,
		},
	},
	"exp_value": {
		_const.RU: {
			"body": `<p><i>Заработано %v очков опыта.</i></p>`,
		},
		_const.EN: {
			"body": `<p><i>%v experience points earned.</i></p>`,
		},
		_const.ZhCN: {
			"body": `<p><i>获得 %v 经验值。</i></p>`,
		},
	},
	"cr_value": {
		_const.RU: {
			"body": `<p><i>Входящий перевод: %v cr.</i></p>`,
		},
		_const.EN: {
			"body": `<p><i>Incoming transfer: %v cr.</i></p>`,
		},
		_const.ZhCN: {
			"body": `<p><i>收到转账：%v 信用点</i></p>`,
		},
	},
	"start_expedition": {
		_const.RU: {
			"body": `<div class="ram_message">
						  <div class="ram_icon"></div>
						  <p class="ram_text">Перехвачен зашифрованный пакет. Подключаюсь... Это же <span class="importantly">ИПЦИ</span>!</p>
					</div>
					<div class="ipci_message">
						<div class="ipci_icon"></div>
					
						<p>
							<span style="color: #a4abaf;'">Входящее сообщение...</span><br> 
							<span class="importantly">ИПЦИ</span>
							<span class="ipci_sub_head">(Институт Пост-Цивилизационных Исследований)</span>
						</p>

						<p>Приветствуем, иследователь!</p>
						<p>Наши датчики зафиксировали вашу активность в зонах аномалий. Ваши полевые данные бесценны для наших архивов.</p>
						<p>Мы готовы финансировать ваши изыскания. <span class="importantly">Первая транзакция за вашу находку уже переведена</span>!</p>
						<p>Продолжайте в том же духе и помните: чем страннее ваши находки, тем веселее нашему отделу аналитики. Не подводите их!</p>
					</div>`,
		},
		_const.EN: {
			"body": `<div class="ram_message">
							<div class="ram_icon"></div>
							<p class="ram_text">Encrypted packet intercepted. Connecting... It's the <span class="importantly">IPCR</span>!</p>
					</div>
					<div class="ipci_message">
						<div class="ipci_icon"></div>

						<p>
							<span style="color: #a4abaf;'">Incoming message...</span><br> 
							<span class="importantly">IPCR</span>
							<span class="ipci_sub_head">(Institute for Post-Civilization Research)</span>
						</p>

						<p>Greetings, explorer!</p>
						<p>Our sensors have detected your activity in anomaly zones. Your field data is invaluable to our archives.</p>
						<p>We are ready to fund your research. <span class="importantly">The first transaction for your find has been transferred</span>!</p>
						<p>Keep up the good work and remember: the stranger your findings, the more fun our analytics department has. Don't let them down!</p>
					</div>`,
		},
		_const.ZhCN: {
			"body": `<div class="ram_message">
							<div class="ram_icon"></div>
							<p class="ram_text">截获加密数据包。正在连接...这竟然是<span class="importantly">IPCR</span>！</p>
					</div>
					<div class="ipci_message">
						<div class="ipci_icon"></div>

						<p>
							<span style="color: #a4abaf;'">传入消息，来自</span><br> 
							<span class="importantly">IPCR</span>
							<span class="ipci_sub_head">(后文明研究所)</span>
						</p>

						<p>向你致敬，探索者！</p>
						<p>我们的传感器侦测到你在异常区域的活动。你提供的现场数据对我们的档案库而言无比珍贵。</p>
						<p>我们已准备好为你的研究提供资金。<span class="importantly">针对你此次发现的第一笔款项已转账</span>！</p>
						<p>请继续保持，并记住：你的发现越奇特，我们的分析部门就越开心。可别让他们失望！</p>
					</div>`,
		},
	},
	"end_expedition_high_1": {
		_const.RU: {
			"body": `
				<div class="ipci_message">
					<div class="ipci_icon"></div>

					<p>
						<span style="color: #a4abaf;'">Входящее сообщение...</span><br> 
						<span class="importantly">ИПЦИ</span>
						<span class="ipci_sub_head">(Институт Пост-Цивилизационных Исследований)</span>
					</p>

					<p>Приветствуем, иследователь!</p>
					<p>Вы провели образцовые полевые иследования! Ставим вас в пример другим исследователям а так же конечно переводим средства на ваш счет!</p>
				</div>
			`,
		},
		_const.EN: {
			"body": `<div class="ipci_message">
						<div class="ipci_icon"></div>
						
						<p>
							<span style="color: #a4abaf;'">Incoming message...</span><br> 
							<span class="importantly">IPCR</span>
							<span class="ipci_sub_head">(Institute for Post-Civilization Research)</span>
						</p>

						<p>Greetings, explorer!</p>
						<p>You have conducted exemplary field research! We are setting you up as an example for other researchers and, of course, transferring the funds to your account!</p>
					</div>`,
		},
		_const.ZhCN: {
			"body": `
					<div class="ipci_message">
						<div class="ipci_icon"></div>

						<p>
							<span style="color: #a4abaf;'">传入消息，来自</span><br> 
							<span class="importantly">IPCR</span>
							<span class="ipci_sub_head">(后文明研究所)</span>
						</p>

						<p>向你致敬，探索者！</p>
						<p>您完成了出色的实地研究工作！我们将您树立为其他研究员的榜样，当然，酬金也已汇入您的账户！</p>
					</div>`,
		},
	},
	"end_expedition_high_2": {
		_const.RU: {
			"body": `
				<div class="ipci_message">
					<div class="ipci_icon"></div>

					<p>
						<span style="color: #a4abaf;'">Входящее сообщение...</span><br> 
						<span class="importantly">ИПЦИ</span>
						<span class="ipci_sub_head">(Институт Пост-Цивилизационных Исследований)</span>
					</p>

					<p>Приветствуем, иследователь!</p>
					<p>Передаём высшую оценку от нашего отдела аналитики и переводим премию на ваш счёт!</p>
				</div>
			`,
		},
		_const.EN: {
			"body": `<div class="ipci_message">
						<div class="ipci_icon"></div>

						<p>
							<span style="color: #a4abaf;'">Incoming message...</span><br> 
							<span class="importantly">IPCR</span>
							<span class="ipci_sub_head">(Institute for Post-Civilization Research)</span>
						</p>

						<p>Greetings, explorer!</p>
						<p>We convey the highest rating from our analytics department and are transferring the bonus to your account!</p>
					</div>`,
		},
		_const.ZhCN: {
			"body": `<div class="ipci_message">
						<div class="ipci_icon"></div>

						<p>
							<span style="color: #a4abaf;'">传入消息，来自</span><br> 
							<span class="importantly">IPCR</span>
							<span class="ipci_sub_head">(后文明研究所)</span>
						</p>

						<p>向你致敬，探索者！</p>
						<p>我们的分析部门给予您最高评价，奖金已汇入您的账户！</p>
					</div>`,
		},
	},
	"end_expedition_med_1": {
		_const.RU: {
			"body": `
				<div class="ipci_message">
					<div class="ipci_icon"></div>

					<p>
						<span style="color: #a4abaf;'">Входящее сообщение...</span><br> 
						<span class="importantly">ИПЦИ</span>
						<span class="ipci_sub_head">(Институт Пост-Цивилизационных Исследований)</span>
					</p>

					<p>Приветствуем, иследователь!</p>
					<p>Собранные данные представляют несомненный интерес для архива. Стандратное вознаграждение уже в пути.</p>
				</div>
			`,
		},
		_const.EN: {
			"body": `<div class="ipci_message">
						<div class="ipci_icon"></div>

						<p>
							<span style="color: #a4abaf;'">Incoming message...</span><br> 
							<span class="importantly">IPCR</span>
							<span class="ipci_sub_head">(Institute for Post-Civilization Research)</span>
						</p>

						<p>Greetings, explorer!</p>
						<p>The collected data is of undeniable interest to the archive. Standard remuneration is already on its way.</p>
					</div>`,
		},
		_const.ZhCN: {
			"body": `<div class="ipci_message">
						<div class="ipci_icon"></div>

						<p>
							<span style="color: #a4abaf;'">传入消息，来自</span><br> 
							<span class="importantly">IPCR</span>
							<span class="ipci_sub_head">(后文明研究所)</span>
						</p>

						<p>向你致敬，探索者！</p>
						<p>您好，研究员！</p> <p>您收集的数据对档案库具有毋庸置疑的价值。标准报酬已在途中。</p>
					</div>`,
		},
	},
	"end_expedition_low_1": {
		_const.RU: {
			"body": `
				<div class="ipci_message">
					<div class="ipci_icon"></div>

					<p>
						<span style="color: #a4abaf;'">Входящее сообщение...</span><br> 
						<span class="importantly">ИПЦИ</span>
						<span class="ipci_sub_head">(Институт Пост-Цивилизационных Исследований)</span>
					</p>

					<p>Приветствуем, иследователь!</p>
					<p>Материалы получены, но их недостаточно для серьёзных выводов. Переводим утешительную компенсацию на ваши усилия.</p>
				</div>
			`,
		},
		_const.EN: {
			"body": `<div class="ipci_message">
						<div class="ipci_icon"></div>

						<p>
							<span style="color: #a4abaf;'">Incoming message...</span><br> 
							<span class="importantly">IPCR</span>
							<span class="ipci_sub_head">(Institute for Post-Civilization Research)</span>
						</p>

						<p>Greetings, explorer!</p>
						<p>The materials have been received, but they are insufficient for substantial conclusions. We are transferring a consolation compensation for your efforts.</p>
					</div>`,
		},
		_const.ZhCN: {
			"body": `<div class="ipci_message">
						<div class="ipci_icon"></div>

						<p>
							<span style="color: #a4abaf;'">传入消息，来自</span><br> 
							<span class="importantly">IPCR</span>
							<span class="ipci_sub_head">(后文明研究所)</span>
						</p>

						<p>向你致敬，探索者！</p>
						<p>资料已收到，但其内容不足以支撑重要结论。我们已为您汇出一笔慰劳金，以补偿您的付出。</p>
					</div>`,
		},
	},
	"end_expedition_low_2": {
		_const.RU: {
			"body": `
				<div class="ipci_message">
					<div class="ipci_icon"></div>

					<p>
						<span style="color: #a4abaf;'">Входящее сообщение...</span><br> 
						<span class="importantly">ИПЦИ</span>
						<span class="ipci_sub_head">(Институт Пост-Цивилизационных Исследований)</span>
					</p>

					<p>Приветствуем, иследователь!</p>
					<p>Надеемся, в следующий раз вы будете более аккуратны с бесценными артефактами. Переводим компенсацию.</p>
				</div>
			`,
		},
		_const.EN: {
			"body": `<div class="ipci_message">
						<div class="ipci_icon"></div>

						<p>
							<span style="color: #a4abaf;'">Incoming message...</span><br> 
							<span class="importantly">IPCR</span>
							<span class="ipci_sub_head">(Institute for Post-Civilization Research)</span>
						</p>

						<p>Greetings, explorer!</p>
						<p>We hope you will be more careful with priceless artifacts next time. Transferring compensation.</p>
					</div>`,
		},
		_const.ZhCN: {
			"body": `<div class="ipci_message">
						<div class="ipci_icon"></div>

						<p>
							<span style="color: #a4abaf;'">传入消息，来自</span><br> 
							<span class="importantly">IPCR</span>
							<span class="ipci_sub_head">(后文明研究所)</span>
						</p>

						<p>向你致敬，探索者！</p>
						<p>您好，研究员！</p> <p>希望您下次能更谨慎地对待这些珍贵的文物。补偿款已汇出。</p>
					</div>`,
		},
	},
}

func GetExpeditionMailTemplates(typeEvent string) map[string]map[string]string {
	textVariants := make([]map[string]map[string]string, 0)
	rand.Seed(time.Now().UnixNano())

	for key, text := range expeditionMailTemplates {
		if strings.Contains(key, typeEvent) {
			textVariants = append(textVariants, text)
		}
	}

	return textVariants[rand.Intn(len(textVariants))]
}
