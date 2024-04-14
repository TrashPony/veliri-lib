package game_math

import (
	"fmt"
	"testing"
)

func TestMath(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(GetRangeRand(7, 12, nil))
	}
}
