package gosafe

import (
	"sync"
)

type Map[K comparable, V any] struct {
	sync.RWMutex
	dict map[K]V
}

func (m *Map[K, V]) SetUnsafe(key K, value V) {
	if m.dict == nil {
		m.dict = map[K]V{}
	}
	m.dict[key] = value
}

func (m *Map[K, V]) Set(key K, value V) {
	m.Lock()
	defer m.Unlock()
	m.SetUnsafe(key, value)
}

func (m *Map[K, V]) GetUnsafe(key K) (V, bool) {
	value, ok := m.dict[key]
	return value, ok
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	m.RLock()
	defer m.RUnlock()
	return m.GetUnsafe(key)
}

func (m *Map[K, V]) DeleteUnsafe(key K) {
	delete(m.dict, key)
}

func (m *Map[K, V]) Delete(key K) {
	m.Lock()
	defer m.Unlock()
	m.DeleteUnsafe(key)
}
