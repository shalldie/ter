package ter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var falsyList = []any{
	false,
	0,
	"",
	nil,
}

func Test_isFalsy(t *testing.T) {
	assert := assert.New(t)

	for _, item := range falsyList {
		assert.Equal(isFalsy(item), true)
	}

	assert.Equal(isFalsy("h"), false)
	assert.Equal(isFalsy(1), false)
	assert.Equal(isFalsy(true), false)
	assert.Equal(isFalsy(make([]int, 0)), false)
	assert.Equal(isFalsy(make(map[string]string)), false)
	assert.Equal(isFalsy(struct{ name string }{name: "tom"}), false)
}

func Test_True(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(True(true, 2, 1), 2)
	assert.Equal(True(false, "a", "b"), "b")
}

func Test_Truthy(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Truthy("a", 2, 1), 2)
	assert.Equal(Truthy(nil, "a", "b"), "b")
}
