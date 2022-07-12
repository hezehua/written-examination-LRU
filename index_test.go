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
	ok1 := l.put("key1","1")
	assert.Equal(t, ok1, true)
	ok2 := l.put("key2","2")
	assert.Equal(t, ok2, true)
	ok3 := l.put("key3","3")
	assert.Equal(t, ok3, true)
	ret, ok := l.get("key1")
	assert.Equal(t, ok, true)
	assert.Equal(t, ret, "1")
	l.put("key4","4")
	ret, ok = l.get("key5")
	assert.Equal(t, ok, false)
	assert.Equal(t, ret, "")
	dataList := l.showAll()
	assert.Equal(t, dataList[0].key, "key4")
	assert.Equal(t, dataList[0].val, "4")
	assert.Equal(t, dataList[1].key, "key1")
	assert.Equal(t, dataList[1].val, "1")
	assert.Equal(t, dataList[2].key, "key3")
	assert.Equal(t, dataList[2].val, "3")

}