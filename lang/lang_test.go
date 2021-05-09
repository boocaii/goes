package lang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsHan(t *testing.T) {
	assert.Equal(t, IsHan(""), false)
	assert.Equal(t, IsHan("1"), false)
	assert.Equal(t, IsHan("你"), true)
	assert.Equal(t, IsHan("你好"), true)
	assert.Equal(t, IsHan(","), false)
}

func TestIsLatin(t *testing.T) {
	assert.Equal(t, IsLatin("ᴀ"), true)
}