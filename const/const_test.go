package _const

import "testing"

// MapBinItemID обязан совпадать с MapBinItems - иначе визуальные хранилища начнут паниковать или путать типы.
func TestMapBinItemIDMatchesMap(t *testing.T) {
	for name, want := range MapBinItems {
		got, ok := MapBinItemID(name)
		if !ok || got != want {
			t.Fatalf("MapBinItemID(%q) = (%d,%v), в MapBinItems %d", name, got, ok, want)
		}
	}

	for _, unknown := range []string{"units", "Unit", " unit", "unit ", "nope", "objects"} {
		if id, ok := MapBinItemID(unknown); ok {
			t.Fatalf("MapBinItemID(%q) = %d, тип неизвестен и должен давать ok=false", unknown, id)
		}
	}
}
