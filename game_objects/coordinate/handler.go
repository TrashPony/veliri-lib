package coordinate

import _const "github.com/TrashPony/veliri-lib/const"

func (coor *Coordinate) ApplyCorporationPolicy(policy string, unions, rivals, hostiles []string, owner string) {
	if policy == _const.CorporationPolicyAll {
		coor.RemovePolicy()
		return
	}

	coor.mx.Lock()
	defer coor.mx.Unlock()

	coor.Access = true
	coor.AccessInvers = policy == _const.CorporationPolicyExceptUnions || policy == _const.CorporationPolicyExceptRivals || policy == _const.CorporationPolicyExceptHostiles
	coor.access = map[string]bool{owner: !coor.AccessInvers}

	if policy == _const.CorporationPolicyOnlyUnions || policy == _const.CorporationPolicyExceptUnions {
		for _, u := range unions {
			coor.access[u] = true
		}
	}

	if policy == _const.CorporationPolicyOnlyRivals || policy == _const.CorporationPolicyExceptRivals {
		for _, r := range rivals {
			coor.access[r] = true
		}
	}

	if policy == _const.CorporationPolicyExceptHostiles {
		for _, h := range hostiles {
			coor.access[h] = true
		}
	}
}

func (coor *Coordinate) RemovePolicy() {
	coor.mx.Lock()
	defer coor.mx.Unlock()

	coor.Access = false
	coor.AccessInvers = false
	coor.access = map[string]bool{}
}
