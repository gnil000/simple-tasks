package main

import (
	"testing"
)

func TestCommon(t *testing.T) {
	key1, key2, key3 := "first", "second", "third"
	value1, value2, value3 := 1, 2, 3

	data := NewStringIntMap()

	data.Add(key1, value1)
	data.Add(key2, value2)
	data.Add(key3, value3)

	if v, ok := data.Get(key1); v != value1 && !ok {
		t.Errorf("expected value %d and exist true", value1)
	}
	if v, ok := data.Get(key2); v != value2 && !ok {
		t.Errorf("expected value %d and exist true", value2)
	}
	if v, ok := data.Get(key3); v != value3 && !ok {
		t.Errorf("expected value %d and exist true", value3)
	}

	data.Remove(key3)
	if ok := data.Exists(key3); ok {
		t.Error("expected not exist, but got exists")
	}

	data2 := data.Copy()
	for key, value := range data.pairs {
		if data2[key] != value {
			t.Errorf("by key %s expect %d, but got %v", key, value, data2[key])
		}
	}
}

func TestAdd(t *testing.T) {
	key := "first"
	value := 1
	data := NewStringIntMap()

	data.Add(key, value)

	if v, exist := data.pairs[key]; !exist && v != value {
		t.Errorf("must be exist and value %d, but not found in original map or another value %d", value, v)
	}
}

func TestRemove(t *testing.T) {
	key := "first"
	value := 1
	data := NewStringIntMap()

	data.pairs[key] = value

	data.Remove(key)

	if _, exist := data.pairs[key]; exist {
		t.Errorf("must be removed, but found in original map")
	}
}

func TestCopy(t *testing.T) {
	key1, key2, key3 := "first", "second", "third"
	value1, value2, value3 := 1, 2, 3
	data := NewStringIntMap()
	data.Add(key1, value1)
	data.Add(key2, value2)
	data.Add(key3, value3)
	data2 := data.Copy()
	for key, value := range data.pairs {
		if data2[key] != value {
			t.Errorf("by key %s expect %d, but got %v", key, value, data2[key])
		}
	}
}

func TestExists(t *testing.T) {
	key := "first"
	value := 1
	data := NewStringIntMap()

	data.pairs[key] = value

	if exist := data.Exists(key); !exist {
		t.Errorf("must be exist, but not found in original map")
	}
}

func TestGet(t *testing.T) {
	key := "first"
	value := 1
	data := NewStringIntMap()

	data.pairs[key] = value

	if v, exist := data.Get(key); !exist && v != value {
		t.Errorf("must be exist and value %d, but not found in original map or another value %d", value, v)
	}
}
