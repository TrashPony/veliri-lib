package user_skills

import (
	_const "github.com/TrashPony/veliri-lib/const"
	"github.com/TrashPony/veliri-lib/game_objects/skill"
	"sync"
)

type UserSkills struct {
	skills           map[string]*skill.Skill
	experiencePoints int
	mx               sync.RWMutex
}

func (u *UserSkills) GetCopyMap() map[string]*skill.Skill {
	u.mx.RLock()
	defer u.mx.RUnlock()

	skills := make(map[string]*skill.Skill)
	for k, s := range u.skills {
		skills[k] = s
	}

	return skills
}

func (u *UserSkills) AddSkill(newSkill *skill.Skill) {

	if u == nil {
		return
	}

	u.mx.Lock()
	defer u.mx.Unlock()

	if u.skills == nil {
		u.skills = make(map[string]*skill.Skill)
	}

	u.skills[newSkill.Name] = newSkill
}

func (u *UserSkills) GetSkill(skillName string) *skill.Skill {

	if u == nil {
		return nil
	}

	u.mx.RLock()
	defer u.mx.RUnlock()

	if u.skills == nil {
		return nil
	}

	return u.skills[skillName]
}

func (u *UserSkills) GetChanSkills() <-chan *skill.Skill {

	if u == nil {
		return nil
	}

	u.mx.RLock()

	skills := make(chan *skill.Skill, len(u.skills))

	go func() {

		defer func() {
			u.mx.RUnlock()
			close(skills)
		}()

		for _, currentSkill := range u.skills {
			skills <- currentSkill
		}
	}()

	return skills
}

func (u *UserSkills) FindSkill(id int) *skill.Skill {

	if u == nil {
		return nil
	}

	u.mx.RLock()
	defer u.mx.RUnlock()

	for _, currentSkill := range u.skills {
		if currentSkill.TypeID == id {
			return currentSkill
		}
	}

	return nil
}

//func (u *UserSkills) AddPointsInSkill(points int, skillName string) (bool, bool, *skill.Skill) {
//
//	if u == nil {
//		return false, false, nil
//	}
//
//	lvlUp := false
//
//	upgradeSkill := u.GetSkill(skillName)
//	if upgradeSkill != nil {
//		upgradeSkill.SetUserExperiencePoint(upgradeSkill.GetUserExperiencePoint() + points)
//
//		needToLvlUp := upgradeSkill.GetPointToUp()
//
//		if needToLvlUp <= upgradeSkill.GetUserExperiencePoint() {
//
//			lvlUp = true
//
//			upgradeSkill.SetLevel(upgradeSkill.GetLevel() + 1)
//			if upgradeSkill.GetLevel() > 25 {
//				upgradeSkill.SetLevel(25)
//				lvlUp = false
//			}
//
//			upgradeSkill.SetUserExperiencePoint(upgradeSkill.GetUserExperiencePoint() - needToLvlUp)
//			u.AddPointsInSkill(0, skillName)
//		}
//
//		return true, lvlUp, upgradeSkill
//	}
//
//	return false, lvlUp, nil
//}

func (u *UserSkills) GetProcessingSkillFactor() int {
	if u == nil {
		return _const.BaseRecyclerLost
	}

	curSkill := u.GetSkill("processing")
	if curSkill != nil {
		return int(_const.BaseRecyclerLost - float64(curSkill.GetLevel())*curSkill.PercentFactor)
	} else {
		return _const.BaseRecyclerLost
	}
}

func (u *UserSkills) GetMaterialsProductionSkillFactor() int {
	curSkill := u.GetSkill("production_economy")
	if curSkill != nil {
		return int(float64(curSkill.GetLevel()) * curSkill.PercentFactor)
	} else {
		return 0
	}
}

func (u *UserSkills) GetProductionTimeSkillFactor() int {
	curSkill := u.GetSkill("production_time")
	if curSkill != nil {
		return int(float64(curSkill.GetLevel()) * curSkill.PercentFactor)
	} else {
		return 0
	}
}

func (u *UserSkills) UpSkill(skillName string) (bool, *skill.Skill) {
	if u == nil {
		return false, nil
	}

	upgradeSkill := u.GetSkill(skillName)
	if upgradeSkill == nil {
		u.AddSkill(upgradeSkill)
		upgradeSkill = u.GetSkill(skillName)
	}

	needToLvlUp := upgradeSkill.GetPointToUp()

	if u.experiencePoints >= needToLvlUp {

		u.experiencePoints -= needToLvlUp
		upgradeSkill.SetLevel(upgradeSkill.GetLevel() + 1)

		if upgradeSkill.GetLevel() > 25 {
			upgradeSkill.SetLevel(25)
		}

		return true, upgradeSkill
	}

	return false, upgradeSkill
}

func (u *UserSkills) AddPoints(points int) {
	if u == nil {
		return
	}

	u.mx.Lock()
	defer u.mx.Unlock()
	u.experiencePoints += points
}

func (u *UserSkills) SetPoints(points int) {
	u.mx.Lock()
	defer u.mx.Unlock()
	u.experiencePoints = points
}

func (u *UserSkills) GetPoints() int {
	u.mx.RLock()
	defer u.mx.RUnlock()
	return u.experiencePoints
}

func (u *UserSkills) CheckRequirements(requirements map[int]int) bool {

	if len(requirements) == 0 {
		// нет требований
		return true
	}

	if u == nil {
		return false
	}

	for id, lvl := range requirements {
		curSkill := u.FindSkill(id)
		if curSkill == nil || curSkill.GetLevel() < lvl {
			return false
		}
	}

	return true
}
