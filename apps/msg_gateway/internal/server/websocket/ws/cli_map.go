package ws

import "sync"

type CliMap struct {
	sync.RWMutex
	m map[string]*Client
}

func NewCliMap() *CliMap {
	return &CliMap{
		m: make(map[string]*Client),
	}
}

func (m *CliMap) Get(k string) (*Client, bool) {
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[k]
	return v, existed
}

func (m *CliMap) Set(k string, v *Client) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

func (m *CliMap) Delete(k string) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, k)
}

func (m *CliMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}

func (m *CliMap) Each(f func(k string, v *Client) bool) {
	m.RLock()
	defer m.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
