package special_hostiles

import "time"

var pointsModerateTypes = map[string]PointsModerate{
	"fraction": { // базовое отношение фракции в цели
		Key:     "fraction",
		Min:     min,
		Max:     max,
		Current: 0,
		Factor:  0.5,
	},
	"battle": { // помощь в бою или наоборот
		Key:     "battle",
		Min:     min,
		Max:     max,
		Current: 0,
		Factor:  0.5,
	},
}

func getModerate(t string) *PointsModerate {
	m, ok := pointsModerateTypes[t]
	if !ok {
		return nil
	}

	return &m
}

type PointsModerate struct {
	Key        string  `json:"key"`
	Min        int     `json:"min"`
	Max        int     `json:"max"`
	Current    int     `json:"current"`
	Factor     float64 `json:"factor"`
	LastUpdate int64   `json:"last_update"`
}

func (p *PointsModerate) SetPoints(hatePoint int) {
	p.Current = hatePoint

	if p.Current > max {
		p.Current = max
	}

	if p.Current < min {
		p.Current = min
	}

	p.LastUpdate = time.Now().UnixNano()
}

func (p *PointsModerate) AddPoints(hatePoint int) {
	p.Current += hatePoint

	if p.Current > max {
		p.Current = max
	}

	if p.Current < min {
		p.Current = min
	}

	p.LastUpdate = time.Now().UnixNano()
}

func (p *PointsModerate) GetPoints() int {
	return int(float64(p.Current) * p.Factor)
}
