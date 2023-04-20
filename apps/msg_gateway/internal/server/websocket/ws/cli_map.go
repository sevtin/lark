package ws

import "sync"

type CliMap struct {
	sync.RWMutex
	clients map[string]*Client
}

func NewCliMap() *CliMap {
	return &CliMap{clients: make(map[string]*Client)}
}

func (m *CliMap) Get(k string) (*Client, bool) {
	m.RLock()
	defer m.RUnlock()
	v, existed := m.clients[k]
	return v, existed
}

func (m *CliMap) Set(k string, v *Client) {
	m.Lock()
	defer m.Unlock()
	m.clients[k] = v
}

func (m *CliMap) Delete(k string) {
	m.Lock()
	defer m.Unlock()
	delete(m.clients, k)
}

func (m *CliMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.clients)
}

func (m *CliMap) Each(f func(k string, v *Client) bool) {
	m.RLock()
	defer m.RUnlock()

	for k, v := range m.clients {
		if !f(k, v) {
			return
		}
	}
}
