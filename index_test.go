package LRU

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRU(t *testing.T) {
	l := LRU{
		cap:3,
		data:make(map[string]dataNode),
	}
	ok1 := l.set("key1","1")
	assert.Equal(t, ok1, true)
	ok2 := l.set("key2","2")
	assert.Equal(t, ok2, true)
	ok3 := l.set("key3","3")
	assert.Equal(t, ok3, true)
	ret, ok := l.get("key1")
	assert.Equal(t, ok, true)
	assert.Equal(t, ret, "1")
	l.set("key4","4")
	ret, ok = l.get("key5")
	assert.Equal(t, ok, false)
	assert.Equal(t, ret, "")
}