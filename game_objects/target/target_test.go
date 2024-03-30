package target

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTarget(t *testing.T) {
	target := &Target{
		Type: "111",
		UUID: "222",
		ID:   4,
	}

	copyTarget := target.GetCopy()
	target.ID = 5
	assert.Equal(t, copyTarget.ID, 4)
}
