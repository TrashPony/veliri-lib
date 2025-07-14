package access_manager

import (
	_const "github.com/TrashPony/veliri-lib/const"
)

func (am *AccessManager) ApplyCorporationPolicy(policy string, unions, rivals, hostiles []string, owner string) {
	if policy == _const.CorporationPolicyAll {
		am.RemovePolicy()
		return
	}

	am.mx.Lock()
	defer am.mx.Unlock()

	am.Access = true
	am.AccessInvers = policy == _const.CorporationPolicyExceptUnions || policy == _const.CorporationPolicyExceptRivals || policy == _const.CorporationPolicyExceptHostiles
	am.access = map[string]bool{owner: !am.AccessInvers}

	if policy == _const.CorporationPolicyOnlyUnions || policy == _const.CorporationPolicyExceptUnions {
		for _, u := range unions {
			am.access[u] = true
		}
	}

	if policy == _const.CorporationPolicyOnlyRivals || policy == _const.CorporationPolicyExceptRivals {
		for _, r := range rivals {
			am.access[r] = true
		}
	}

	if policy == _const.CorporationPolicyExceptHostiles {
		for _, h := range hostiles {
			am.access[h] = true
		}
	}
}

func (am *AccessManager) RemovePolicy() {
	am.mx.Lock()
	defer am.mx.Unlock()

	am.Access = false
	am.AccessInvers = false
	am.access = map[string]bool{}
}
