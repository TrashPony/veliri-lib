package skill

import (
	"math"
)

type Skill struct {
	ID                  int      `json:"id"`
	TypeID              int      `json:"type_id"`
	Name                string   `json:"name"`
	Specification       string   `json:"specification"`
	ExperiencePoint     int      `json:"experience_point"`
	UserExperiencePoint int      `json:"user_experience_point"`
	ExperienceFactor    float64  `json:"experience_factor"`
	PercentFactor       float64  `json:"percent_factor"`
	Type                string   `json:"type"`
	Level               int      `json:"level"`
	Icon                string   `json:"icon"`
	NeedPointsToUp      int      `json:"need_points_to_up"`
	Parameters          []string `json:"parameters"`
}

func (s *Skill) GetPointToUp() int {
	s.NeedPointsToUp = int(math.Pow(float64(s.ExperiencePoint*s.Level), s.ExperienceFactor)) + s.ExperiencePoint
	return s.NeedPointsToUp
}

func (s *Skill) GetLevel() int {
	return s.Level
}

func (s *Skill) SetLevel(newLvl int) {
	s.Level = newLvl
}

func (s *Skill) GetUserExperiencePoint() int {
	return s.UserExperiencePoint
}

func (s *Skill) SetUserExperiencePoint(points int) {
	s.UserExperiencePoint = points
}

func (s *Skill) GetCopy() *Skill {
	copy := *s
	return &copy
}
