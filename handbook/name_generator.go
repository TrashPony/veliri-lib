package handbook

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"math/rand"
	"strings"
)

var commonSyllable = []string{
	"АЛ",
	"ЯЛ",
	"АР",
	"ПА",
	"ТА",
	"НА",
	"ЛА",
	"КИ",
	"МА",
	"КО",
	"ЛИ",
	"ЛЬ",
	"НО",
	"ОК",
	"УК",
	"СА",
	"КЛ",
	"ОС",
	"ИР",
	"ЛЕ",
	"ЛО",
	"ЕЙ",
	"УР",
	"ИС",
	"ОЛ",
	"СО",
	"ОВ",
	"ТО",
	"РО",
	"ОН",
	"ПО",
}

var replicStart = []string{
	"РА",
	"ДЕ",
	"ИГ",
	"ЮД",
}

var exploresStart = []string{
	"УЛ",
	"МЯ",
	"ЕЛ",
	"АН",
}

var reversesStart = []string{
	"ИК",
	"УТ",
	"ЯК",
	"АН",
	"КА",
}

var apdStart = []string{
	"ФА",
	"ЦЕ",
	"ЁС",
	"СО",
}

// todo это пока чисто для тестов, надо еще давать префиксы от роли и ваще нормальный генератор сделать
func GetFractionName(fraction string) string {
	if fraction == _const.Replicas {
		return strings.Title(strings.ToLower(replicStart[rand.Intn(len(replicStart))] +
			commonSyllable[rand.Intn(len(commonSyllable))] +
			commonSyllable[rand.Intn(len(commonSyllable))]))
	}

	if fraction == _const.Reverses {
		return strings.Title(strings.ToLower(reversesStart[rand.Intn(len(reversesStart))] +
			commonSyllable[rand.Intn(len(commonSyllable))] +
			commonSyllable[rand.Intn(len(commonSyllable))]))
	}

	if fraction == _const.Explores {
		return strings.Title(strings.ToLower(exploresStart[rand.Intn(len(exploresStart))] +
			commonSyllable[rand.Intn(len(commonSyllable))] +
			commonSyllable[rand.Intn(len(commonSyllable))]))
	}

	if fraction == _const.APD {
		return strings.Title(strings.ToLower(apdStart[rand.Intn(len(apdStart))] +
			commonSyllable[rand.Intn(len(commonSyllable))] +
			commonSyllable[rand.Intn(len(commonSyllable))]))
	}

	return ""
}
