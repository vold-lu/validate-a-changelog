package internal

import (
	"encoding/json"
	"testing"
)

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

func TestSortedMap_MarshalJSON(t *testing.T) {
	cases := []struct {
		Name     string
		Map      interface{ MarshalJSON() ([]byte, error) }
		Expected string
	}{
		{
			Name:     "Empty",
			Map:      NewEmptyMap[string, int](),
			Expected: `{}`,
		},
		{
			Name:     "KeepsInsertionOrder",
			Map:      NewSortedMap([]string{"c", "a", "b"}, map[string]int{"a": 1, "b": 2, "c": 3}),
			Expected: `{"c":3,"a":1,"b":2}`,
		},
		{
			Name:     "EscapesKeyAndValue",
			Map:      NewSortedMap([]string{`Les "Commande"`}, map[string]string{`Les "Commande"`: `et "Tiers"`}),
			Expected: `{"Les \"Commande\"":"et \"Tiers\""}`,
		},
		{
			Name:     "NonStringKey",
			Map:      NewSortedMap([]int{2, 1}, map[int]string{1: "one", 2: "two"}),
			Expected: `{"2":"two","1":"one"}`,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			b, err := json.Marshal(c.Map)
			if err != nil {
				t.Fatalf("MarshalJSON() returned an error: %v", err)
			}

			if string(b) != c.Expected {
				t.Fatalf("MarshalJSON(). Got %s, wanted %s", string(b), c.Expected)
			}
		})
	}
}
