package obj

import "sync"

type RwMap struct {
	sync.RWMutex
	m map[any]any
}

func NewRwMap() *RwMap {
	return &RwMap{
		m: make(map[any]any),
	}
}

func (m *RwMap) Get(k any) (any, bool) {
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[k]
	return v, existed
}

func (m *RwMap) Set(k any, v any) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

func (m *RwMap) Delete(k any) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, k)
}

func (m *RwMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}

func (m *RwMap) Each(f func(k any, v any) bool) {
	m.RLock()
	defer m.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
