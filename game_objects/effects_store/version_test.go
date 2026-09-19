package effects_store

import (
	"testing"

	"github.com/TrashPony/veliri-lib/game_objects/effect"
)

func TestVersionChangesOnlyWhenSetChanges(t *testing.T) {
	e := &EffectsStore{}
	v0 := e.Version()

	if !e.AddEffect(&effect.Effect{UUID: "a", Parameter: "view", Quantity: 10}) {
		t.Fatal("эффект не добавился")
	}
	v1 := e.Version()
	if v1 == v0 {
		t.Fatal("добавление обязано менять версию")
	}

	if e.AddEffect(&effect.Effect{UUID: "a", Parameter: "view", Quantity: 10}) {
		t.Fatal("дубликат UUID не должен добавляться")
	}
	if e.Version() != v1 {
		t.Fatal("неудачное добавление не должно менять версию")
	}

	if ok, _ := e.RemoveEffect("nope"); ok || e.Version() != v1 {
		t.Fatal("удаление несуществующего не должно менять версию")
	}

	if ok, _ := e.RemoveEffect("a"); !ok || e.Version() == v1 {
		t.Fatal("удаление обязано менять версию")
	}
}
