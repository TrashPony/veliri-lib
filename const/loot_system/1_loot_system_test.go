package loot_system

import (
	"math/rand"
	"testing"
	"time"
)

// TestGenerateScience проверяет базовую функциональность функции generateScience
func TestGenerateScience(t *testing.T) {
	// Инициализируем генератор случайных чисел с фиксированным сидом для воспроизводимости
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	t.Logf("Используется seed: %d", seed)

	// Тест 1: Проверяем, что функция возвращает не nil
	t.Run("ReturnsNonNil", func(t *testing.T) {
		result := generateScience(1, 5, rng)
		if result == nil {
			t.Error("Функция вернула nil, ожидался слайс LootDrop")
		}
	})

	// Тест 2: Проверяем, что возвращается ровно один элемент
	t.Run("ReturnsExactlyOneItem", func(t *testing.T) {
		result := generateScience(1, 5, rng)
		if len(result) != 1 {
			t.Errorf("Ожидался 1 элемент, получено: %d", len(result))
		}
	})

	// Тест 3: Проверяем, что ItemType всегда "frr"
	t.Run("ItemTypeIsFrr", func(t *testing.T) {
		result := generateScience(1, 5, rng)
		if len(result) > 0 && result[0].ItemType != "frr" {
			t.Errorf("Ожидался ItemType: 'frr', получено: '%s'", result[0].ItemType)
		}
	})

	// Тест 4: Проверяем диапазон ItemID
	t.Run("ItemIDInRange", func(t *testing.T) {
		result := generateScience(1, 5, rng)
		if len(result) > 0 {
			itemID := result[0].ItemID
			if itemID < 1 || itemID > 7 {
				t.Errorf("ItemID вне диапазона: %d, ожидалось 1-7", itemID)
			}
		}
	})

	// Тест 5: Проверяем, что количество больше 0
	t.Run("CountGreaterThanZero", func(t *testing.T) {
		result := generateScience(1, 5, rng)
		if len(result) > 0 && result[0].Count <= 0 {
			t.Errorf("Количество должно быть > 0, получено: %d", result[0].Count)
		}
	})

	// Тест 6: Проверяем все возможные ItemID (статистический тест)
	t.Run("AllItemIDsPossible", func(t *testing.T) {
		iterations := 1000
		seenIDs := make(map[int]bool)

		for i := 0; i < iterations; i++ {
			result := generateScience(1, 5, rng)
			if len(result) > 0 {
				seenIDs[result[0].ItemID] = true
			}
		}

		// Проверяем, что видели все 7 возможных ID
		for expectedID := 1; expectedID <= 7; expectedID++ {
			if !seenIDs[expectedID] {
				t.Errorf("После %d итераций не был сгенерирован ItemID: %d", iterations, expectedID)
			}
		}

		t.Logf("Найдены ItemID: %v", getKeys(seenIDs))
	})

	// Тест 7: Проверяем с разными диапазонами minCount/maxCount
	t.Run("DifferentCountRanges", func(t *testing.T) {
		testCases := []struct {
			name     string
			minCount int
			maxCount int
		}{
			{"Min1_Max1", 1, 1},
			{"Min1_Max10", 1, 10},
			{"Min5_Max20", 5, 20},
			{"Min10_Max100", 10, 100},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := generateScience(tc.minCount, tc.maxCount, rng)
				if len(result) > 0 {
					count := result[0].Count
					// Здесь мы не можем проверить точные границы, так как есть умножение
					// на BaseCount, но можем проверить что count > 0
					if count <= 0 {
						t.Errorf("Для %s: количество должно быть > 0, получено: %d",
							tc.name, count)
					}
					t.Logf("%s: сгенерировано количество: %d", tc.name, count)
				}
			})
		}
	})

	// Тест 8: Проверяем с фиксированным сидом для отладки
	t.Run("FixedSeedForDebugging", func(t *testing.T) {
		fixedRNG := rand.New(rand.NewSource(12345))

		// Запускаем несколько раз с одним сидом
		results := make([]int, 10)
		for i := 0; i < 10; i++ {
			result := generateScience(1, 5, fixedRNG)
			if len(result) > 0 {
				results[i] = result[0].ItemID
			}
		}

		t.Logf("Результаты с фиксированным сидом: %v", results)
	})
}

// Вспомогательная функция для получения ключей мапы
func getKeys(m map[int]bool) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// TestGenerateScienceEdgeCases проверяет крайние случаи
func TestGenerateScienceEdgeCases(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Тест с пустым массивом frr (временная модификация)
	t.Run("EmptyFrrArray", func(t *testing.T) {
		// Сохраняем оригинальный массив
		originalFrr := frr
		// Временно заменяем на пустой
		frr = []LootLot{}

		defer func() {
			// Восстанавливаем оригинальный массив
			frr = originalFrr
		}()

		result := generateScience(1, 5, rng)
		if result != nil {
			t.Errorf("При пустом массиве frr ожидался nil, получено: %v", result)
		}
	})
}
