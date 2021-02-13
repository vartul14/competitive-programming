package main

import (
	"testing"
)

func TestCacheBasic(t *testing.T) {
	var c Cache
	var evPolicy EvictionPolicyTypes
	evPolicy = LRU
	c.InitCache(5, evPolicy)
	c.Set("a", "b")
	value, _ := c.Get("a")
	if value != "b" {
		t.Errorf("Error getting value for key")
	}
}

func TestCacheMultipleSet(t *testing.T) {
	var c Cache
	var evPolicy EvictionPolicyTypes
	evPolicy = LRU
	c.InitCache(5, evPolicy)
	c.Set("key1", "value1")
	c.Set("key2", "value2")
	value, _ := c.Get("key2")
	if value != "value2" {
		t.Errorf("Error getting value for key")
	}
}


func TestCacheOverflow(t *testing.T) {
	var c Cache
	var evPolicy EvictionPolicyTypes
	evPolicy = LRU
	c.InitCache(2, evPolicy)
	c.Set("key1", "value1")
	c.Set("key2", "value2")
	c.Set("key3", "value3")
	_, err := c.Get("key1")
	if err == nil {
		t.Errorf("Over flow case failing")
	}
}