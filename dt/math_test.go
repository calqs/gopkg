package dt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_I_Can_Round_Float(t *testing.T) {
	assert.Equal(t, RoundFloat(1.2345, 2), 1.23)
	assert.Equal(t, RoundFloat(1.2345, 3), 1.235)
	assert.Equal(t, RoundFloat(1.2345, 4), 1.2345)
	assert.Equal(t, RoundFloat(1.2345, 5), 1.23450)
}

func TestSum(t *testing.T) {
	assert.Equal(t, Sum([]int{1, 2, 3}, func(i int) int { return i }), 6)
}

func TestMax(t *testing.T) {
	assert.Equal(t, Max(1, 2), 2)
	assert.Equal(t, Max(2, 1), 2)
}
