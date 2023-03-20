package lru

import (
	"testing"
)

func TestLru(t *testing.T) {
	lru := Constructor(3)
	lru.Put("key1", 100)
	lru.Put("key2", "value1")
	t.Log(lru.Get("key1"))
	lru.Put("key1", "value11")
	t.Log(lru.Get("key1"))
	lru.Put("key3", "value3")
	lru.Put("key4", "value4")
	t.Log(lru.Get("key2"))
	t.Log(lru.cache)
}
