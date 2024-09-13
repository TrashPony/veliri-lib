package descriptions

import _const "github.com/TrashPony/veliri-lib/const"

var BoxDescription = map[string]map[string]DescriptionItem{
	_const.EN: {
		"box":                    {Name: "Standard container", Description: "<p> A container for storing and transporting a variety of things. Susceptible to external influences and prone to destruction. </p>"},
		"underground_box":        {Name: "Underground container", Description: "<p> A container that can contain useful things or artifacts of the former civilization of the planet Veliri-6. </p>"},
		"secure_underground_box": {Name: "Underground secure container", Description: "<p> Well protected and locked by an ingenious system. The best way to unlock such a container so as not to damage its contents is to enter the correct password - a combination of certain four numbers. </p> <p> But, like other containers, it is prone to destruction. </p>"},
	},
	_const.RU: {
		"box":                    {Name: "Стандартный контейнер", Description: "<p>Вместилище для хранения и перевозки разнообразных вещей. Подвержен внешнему воздействию и склонен к разрушению.</p>"},
		"underground_box":        {Name: "Подземный контейнер", Description: "<p>Ёмкость, что может содержать в себе полезные вещи или, артефакты прежней цивилизации планеты Veliri-6.</p>"},
		"secure_underground_box": {Name: "Подземный защищенный контейнер", Description: "<p>Хорошо защищённый и запертый посредством хитроумной системы ящик. Самый лучший способ отпереть подобный контейнер, чтобы не повредить его содержимое – ввести верный пароль – комбинацию из определённых четырёх чисел. </p><p>Но, как и прочие контейнеры, и он склонен к разрушению.</p>"},
	},
}
