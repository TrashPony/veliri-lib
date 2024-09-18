package handbook

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"math/rand"
	"strings"
	"time"
)

var configs = map[string]nameConfigs{
	_const.Replicas: {
		names:     []string{"Destroyer", "Warbringer", "Annihilator", "Ravager", "Blitzkrieg", "Crusher", "Dominator", "Marauder", "Vindicator", "Obliterator", "Conqueror", "Warlord", "Berserker", "Devastator", "Overlord", "Raptor"},
		postfixes: []string{"X1", "MK-II", "3000", "Prime", "Alpha", "Beta", "Gamma", "Omega", "X5", "MK-V", "4000", "Ultra", "Delta", "Epsilon", "Zeta", "Theta"},
	},
	_const.Explores: {
		names:     []string{"Explorer", "Pathfinder", "Seeker", "Navigator", "Surveyor", "Scout", "Pioneer", "Voyager", "Adventurer", "Discoverer", "Trailblazer", "Prospector", "Ranger", "Frontier", "Questor", "Wanderer"},
		postfixes: []string{"X2", "MK-III", "4000", "Explorer", "Delta", "Epsilon", "Zeta", "Theta", "X6", "MK-VI", "5000", "Pioneer", "Lambda", "Mu", "Nu", "Xi"},
	},
	_const.Reverses: {
		names:     []string{"RedStar", "ComradeBot", "Proletariat", "Unionizer", "Collectivist", "Revolutionary", "Soviet", "Bolshevik", "Marxist", "Leninist", "Trotsky", "Stalinist", "Maoist", "Vanguard", "WorkerBot", "PeopleBot"},
		postfixes: []string{"X3", "MK-IV", "5000", "Union", "Sigma", "Tau", "Upsilon", "Phi", "X7", "MK-VII", "6000", "Collective", "Chi", "Psi", "Omega", "Kappa"},
	},
	_const.APD: {
		names:     []string{"MadMachine", "InsaneBot", "Lunatic", "Deranged", "Maniac", "Psycho", "Frenzy", "Berserk", "Havoc", "Chaos", "Pandemonium", "Bedlam", "Rage", "Fury", "Wrath", "Mayhem"},
		postfixes: []string{""}, //[]string{"X8", "MK-VIII", "7000", "Insane", "Madness", "Frenzy", "Chaos", "Rage", "X9", "MK-IX", "8000", "Havoc", "Pandemonium", "Bedlam", "Wrath", "Mayhem"}
	},
	_const.FarGod: {
		names:     []string{"Zealot", "Cultist", "Devotee", "Worshipper", "Acolyte", "Disciple", "Believer", "Prophet", "Oracle", "Seer", "Mystic", "Shaman", "Priest", "Cleric", "Monk", "Hermit"},
		postfixes: []string{"Divine", "Holy", "Sacred", "Mystic", "Prophet", "X11", "MK-XI", "10000", "Oracle", "Seer", "Shaman", "Priest", "Cleric"},
	},
}

type nameConfigs struct {
	names     []string
	postfixes []string
}

func GetFractionName(fraction string) string {

	cfg, ok := configs[fraction]
	if !ok {
		return ""
	}

	mc := NewMarkovChain()
	for _, name := range cfg.names {
		sequence := strings.Split(name, "")
		mc.Add(sequence)
	}

	minName := 3
	maxName := 10

	return mc.generateWitchPrefix(rand.Intn(maxName-minName)+minName, cfg.postfixes)
}

// MarkovChain представляет цепь Маркова
type MarkovChain struct {
	prefixMap map[string][]string
}

// NewMarkovChain создает новую цепь Маркова
func NewMarkovChain() *MarkovChain {
	return &MarkovChain{prefixMap: make(map[string][]string)}
}

// Add добавляет последовательность в цепь Маркова
func (mc *MarkovChain) Add(sequence []string) {
	for i := 0; i < len(sequence)-1; i++ {
		prefix := sequence[i]
		suffix := sequence[i+1]
		mc.prefixMap[prefix] = append(mc.prefixMap[prefix], suffix)
	}
}

// Generate генерирует новое имя на основе цепи Маркова
func (mc *MarkovChain) generate(start string, length int) string {
	name := []string{start}
	current := start

	for i := 0; i < length-1; i++ {
		choices := mc.prefixMap[current]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		name = append(name, next)
		current = next
	}

	return strings.Title(strings.Join(name, ""))
}

// generateWitchPrefix генерирует имя
func (mc *MarkovChain) generateWitchPrefix(length int, postfixes []string) string {
	rand.Seed(time.Now().UnixNano())

	start := mc.getRandomStart()
	name := mc.generate(start, length)

	postfix := postfixes[rand.Intn(len(postfixes))]
	if postfix != "" && rand.Intn(3) == 0 {
		return name + "-" + postfix
	}

	return name
}

// getRandomStart выбирает случайную стартовую позицию
func (mc *MarkovChain) getRandomStart() string {
	keys := make([]string, 0, len(mc.prefixMap))
	for key := range mc.prefixMap {
		keys = append(keys, key)
	}
	return keys[rand.Intn(len(keys))]
}
