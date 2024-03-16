package special_hostiles

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointsModerate(t *testing.T) {
	m1 := getModerate("fraction")
	m1.Current = 150

	m2 := getModerate("fraction")
	assert.Equal(t, m2.Current, 0)
}
