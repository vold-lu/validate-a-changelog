package internal

import "testing"

func TestNewEmptyMap(t *testing.T) {
	m := NewEmptyMap[string, string]()

	if m == nil {
		t.Fatal("NewEmptyMap() returned nil")
	}

	if m.Len() != 0 {
		t.Fatal("NewEmptyMap() returned wrong length")
	}
}

func TestNewSortedMap(t *testing.T) {
	keys := []string{"a", "b", "c"}
	values := map[string]int{"a": 1, "b": 2, "c": 3}

	m := NewSortedMap(keys, values)

	if m == nil {
		t.Fatal("NewSortedMap() returned nil")
	}

	if m.Len() != 3 {
		t.Fatal("NewSortedMap() returned wrong length")
	}
}

func TestSortedMap_Get(t *testing.T) {
	keys := []string{"a", "b", "c"}
	values := map[string]int{"a": 1, "b": 2, "c": 3}

	m := NewSortedMap(keys, values)

	for _, key := range keys {
		if val, _ := m.Get(key); val != values[key] {
			t.Fatal("NewSortedMap() returned wrong value")
		}
	}
}

func TestSortedMap_Set(t *testing.T) {
	m := NewEmptyMap[string, string]()

	if err := m.Set("test", "alois"); err != nil {
		t.Fatal(err)
	}

	if m.Len() != 1 {
		t.Fatal("wrong length")
	}

	if val, _ := m.Get("test"); val != "alois" {
		t.Fatal("wrong value")
	}

	if err := m.Set("test", "alois2"); err != nil {
		t.Fatal(err)
	}

	if m.Len() != 1 {
		t.Fatal("wrong length")
	}

	if val, _ := m.Get("test"); val != "alois2" {
		t.Fatal("wrong value")
	}
}

func TestSortedMap_Del(t *testing.T) {
	keys := []string{"a", "b", "c"}
	values := map[string]int{"a": 1, "b": 2, "c": 3}

	m := NewSortedMap(keys, values)

	if m.Len() != 3 {
		t.Fatal("NewSortedMap() returned wrong length")
	}

	if m.Del("a") != nil {
		t.Fatal("Del() returned wrong value")
	}

	if m.Len() != 2 {
		t.Fatal("Len() returned wrong length")
	}

	if m.Del("a") == nil {
		t.Fatal("Del() returned wrong value")
	}
}

func TestSortedMap_Keys(t *testing.T) {
	keys := []string{"a", "b", "c"}
	values := map[string]int{"a": 1, "b": 2, "c": 3}

	m := NewSortedMap(keys, values)

	for i, key := range keys {
		if m.Keys()[i] != key {
			t.Fatal("NewSortedMap() returned wrong key")
		}
	}
}

func TestSortedMap_Len(t *testing.T) {
	keys := []string{"a", "b", "c"}
	values := map[string]int{"a": 1, "b": 2, "c": 3}

	m := NewSortedMap(keys, values)

	if m.Len() != 3 {
		t.Fatal("NewSortedMap() returned wrong length")
	}
}

func TestSortedMap_Has(t *testing.T) {
	keys := []string{"a", "b", "c"}
	values := map[string]int{"a": 1, "b": 2, "c": 3}

	m := NewSortedMap(keys, values)

	for _, key := range keys {
		if !m.Has(key) {
			t.Fatal("NewSortedMap() returned wrong value")
		}
	}
}
