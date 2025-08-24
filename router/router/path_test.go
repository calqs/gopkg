package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_I_Can_CleanPath(t *testing.T) {
	assert.Equal(t, "/im_the_mountain", CleanPath("im_the_mountain/////"))
}
