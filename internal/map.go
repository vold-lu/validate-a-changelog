package internal

import (
	"bytes"
	"cmp"
	"encoding/json"
	"fmt"
	"sync"
)

type SortedMap[K cmp.Ordered, V any] struct {
	values map[K]V
	keys   []K

	mutex sync.RWMutex
}

func NewEmptyMap[K cmp.Ordered, V any]() *SortedMap[K, V] {
	return &SortedMap[K, V]{
		values: make(map[K]V),
		keys:   make([]K, 0),
		mutex:  sync.RWMutex{},
	}
}

func NewSortedMap[K cmp.Ordered, V any](keys []K, values map[K]V) *SortedMap[K, V] {
	return &SortedMap[K, V]{
		values: values,
		keys:   keys,
		mutex:  sync.RWMutex{},
	}
}

func (sm *SortedMap[K, V]) Get(key K) (V, bool) {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	val, exists := sm.values[key]
	return val, exists
}

func (sm *SortedMap[K, V]) Set(key K, val V) error {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	sm.values[key] = val
	sm.keys = append(sm.keys, key)

	return nil
}

func (sm *SortedMap[K, V]) Del(key K) error {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	delete(sm.values, key)

	idx := -1
	for i, k := range sm.keys {
		if k == key {
			idx = i
		}
	}

	// Should not happen
	if idx == -1 {
		return fmt.Errorf("SortedMap: unable to find key in cache: %v", key)
	}

	sm.keys = append(sm.keys[:idx], sm.keys[idx+1:]...)

	return nil
}

func (sm *SortedMap[K, V]) Keys() []K {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	return sm.keys
}

func (sm *SortedMap[K, V]) Len() int {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	return len(sm.keys)
}

func (sm *SortedMap[K, V]) Has(key K) bool {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	_, exists := sm.values[key]
	return exists
}

// MarshalJSON Re implement the JSON serialization to ensure JSON serialization respect ordering
func (sm *SortedMap[K, V]) MarshalJSON() ([]byte, error) {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	var bb bytes.Buffer
	bb.WriteRune('{')

	for i, key := range sm.keys {
		if i != 0 {
			bb.WriteRune(',')
		}

		val, exists := sm.values[key]

		// Should not happen
		if !exists {
			return nil, fmt.Errorf("SortedMap: unable to find value in cache: %v", key)
		}

		// Serialize the key
		bb.WriteString(fmt.Sprintf("\"%v\"", key))

		bb.WriteRune(':')

		// Serialize the value
		if b, err := json.Marshal(val); err != nil {
			return nil, err
		} else {
			bb.Write(b)
		}
	}

	bb.WriteRune('}')

	return bb.Bytes(), nil
}
