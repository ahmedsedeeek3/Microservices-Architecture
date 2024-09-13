package discovery

import "sync"

type ServiceDiscovery struct {
	services map[string]string
	mu       sync.RWMutex
}

func NewServiceDiscovery() *ServiceDiscovery {
	return &ServiceDiscovery{
		services: make(map[string]string),
	}
}

func (sd *ServiceDiscovery) Register(name, endpoint string) {
	sd.mu.Lock()
	defer sd.mu.Unlock()
	sd.services[name] = endpoint
}

func (sd *ServiceDiscovery) Discover(name string) (string, bool) {
	sd.mu.RLock()
	defer sd.mu.RUnlock()
	endpoint, ok := sd.services[name]
	return endpoint, ok
}
