package interface_window

import _const "github.com/TrashPony/veliri-lib/const"

var Global = map[string]map[string]string{
	"text_1": {
		_const.RU:   `Помощь`,
		_const.EN:   `Help`,
		_const.ZhCN: `帮助`,
	},
	"text_2": {
		_const.RU:   `Доберитесь до указанной на <span style="color: yellow">мини-карте</span> зоне.`,
		_const.EN:   `Get to the area indicated on the <span style="color: yellow">minimap</span>.`,
		_const.ZhCN: `前往<span style="color: yellow">小地图</span>上指示的区域。`,
	},
	"text_3": {
		_const.RU:   `Примечание:`,
		_const.EN:   `Note:`,
		_const.ZhCN: `注意：`,
	},
	"text_4": {
		_const.RU:   `движение.`,
		_const.EN:   `movement.`,
		_const.ZhCN: `移动。`,
	},
	"text_5": {
		_const.RU:   `ускорение.`,
		_const.EN:   `acceleration.`,
		_const.ZhCN: `加速。`,
	},
	"text_6": {
		_const.RU:   `тормоз.`,
		_const.EN:   `brake.`,
		_const.ZhCN: `刹车。`,
	},
	"text_7": {
		_const.RU:   `Мышь`,
		_const.EN:   `Mouse`,
		_const.ZhCN: `鼠标`,
	},
	"text_8": {
		_const.RU:   `направление движения.`,
		_const.EN:   `direction of movement.`,
		_const.ZhCN: `移动方向。`,
	},
	"text_9": {
		_const.RU:   `Уничтожьте бочки и баки на пути <span style="color: grey; font-size: 12px">(рекомендуется активировать щит)</span>.`,
		_const.EN:   `Destroy the barrels and tanks on the way <span style="color: grey; font-size: 12px">(it is recommended to activate the shield)</span>.`,
		_const.ZhCN: `摧毁途中的桶和罐子 <span style="color: grey; font-size: 12px">（建议激活护盾）</span>。`,
	},
	"text_10": {
		_const.RU:   `Примечание:`,
		_const.EN:   `Note:`,
		_const.ZhCN: `注意：`,
	},
	"text_11": {
		_const.RU: `Что бы активировать оружие/снаряжение его надо выбрать на панели в <span
            style="color: yellow">центре с низу
            экрана</span> или нажать на указанной на ячейке клавишу (<span style="color: yellow">2</span> или '<span
            style="color: yellow">E</span>').`,
		_const.EN: `To activate the weapons/equipment, you need to select it on the panel in the <span style="color: yellow">center
			at the bottom of the screen</span> or press the key indicated in the cell (<span style="color: yellow">2</span>
			or '<span style="color: yellow">E</span>').`,
		_const.ZhCN: `要激活武器/装备，您需要在<span style="color: yellow">屏幕底部中心</span>的面板上选择它，或按下单元格中指示的键（<span style="color: yellow">2</span>或<span style="color: yellow">E</span>）。`,
	},
	"text_12": {
		_const.RU:   `прицеливание (камера будет следовать за мышкой).`,
		_const.EN:   `aiming (the camera will follow the mouse).`,
		_const.ZhCN: `瞄准（摄像头将跟随鼠标移动）。`,
	},
	"text_13": {
		_const.RU:   `ЛКМ`,
		_const.EN:   `LMB`,
		_const.ZhCN: `左键`,
	},
	"text_14": {
		_const.RU:   `атака.`,
		_const.EN:   `attack.`,
		_const.ZhCN: `攻击。`,
	},
	"text_15": {
		_const.RU:   `Красный круг`,
		_const.EN:   `Red circle`,
		_const.ZhCN: `红色圆圈`,
	},
	"text_16": {
		_const.RU:   `зона поражения.`,
		_const.EN:   `affected area.`,
		_const.ZhCN: `影响区域。`,
	},
	"text_17": {
		_const.RU:   `перезарядка.`,
		_const.EN:   `reloading.`,
		_const.ZhCN: `重新装填。`,
	},
	"text_18": {
		_const.RU:   `Неизвестная угроза!`,
		_const.EN:   `Unknown threat!`,
		_const.ZhCN: `未知威胁！`,
	},
	"text_19": {
		_const.RU:   `Вы являетесь врагом в этом секторе!`,
		_const.EN:   `You are the enemy in this sector!`,
		_const.ZhCN: `您是这个区域的敌人！`,
	},
	"text_20": {
		_const.RU:   `включить фары.`,
		_const.EN:   `turn off the headlights.`,
		_const.ZhCN: `关闭车灯。`,
	},
	"text_21": {
		_const.RU:   `Что бы выбрать цель, <span style="color: yellow">наведите мышку</span> на транспорт и нажмите <span style="color: yellow">TAB</span>.`,
		_const.EN:   `To select a target, <span style="color: yellow">move the mouse</span> over the vehicle and press <span style="color: yellow">TAB</span>.`,
		_const.ZhCN: `要选择目标，<span style="color: yellow">将鼠标</span>悬停在车辆上并按<span style="color: yellow">TAB</span>。`,
	},
	"text_22": {
		_const.RU:   `Что бы убрать цель, нажмите <span style="color: yellow">TAB</span> в месте где нет транспорта.`,
		_const.EN:   `To remove a target, press <span style="color: yellow">TAB</span> in a place where there is no vehicle.`,
		_const.ZhCN: `要取消目标，在没有车辆的地方按<span style="color: yellow">TAB</span>。`,
	},
	"text_23": {
		_const.RU:   `Колесо мышки`,
		_const.EN:   `Mouse wheel`,
		_const.ZhCN: `鼠标滚轮`,
	},
	"text_24": {
		_const.RU:   `зум.`,
		_const.EN:   `zoom.`,
		_const.ZhCN: `缩放。`,
	},
	"text_25": {
		_const.RU:   `Ваш запрос остался без ответа. Удачи!`,
		_const.EN:   `Your request has gone unanswered. Good luck!`,
		_const.ZhCN: `您的请求未得到回应。祝您好运！`,
	},
	"text_26": {
		_const.RU:   `Транспорт на подходе: `,
		_const.EN:   `Transport is on the way: `,
		_const.ZhCN: `运输工具正在路上：`,
	},
	"text_27": {
		_const.RU:   `Транспорт в секторе...`,
		_const.EN:   `Transport is in the sector...`,
		_const.ZhCN: `运输工具在区域中...`,
	},
}
