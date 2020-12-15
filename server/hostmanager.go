package server

import "fmt"

// HostManager manages host instances. TODO: Make it safe for concurrent access.
type HostManager struct {
	hosts map[string]Host
}

// NewHostManager returns a new instance of a host manager.
func NewHostManager() HostManager {
	return HostManager{
		hosts: map[string]Host{},
	}
}

// AddHost adds a new host instance to the LoadBalancer.
func (manager *HostManager) AddHost(h Host) error {
	if _, ok := manager.hosts[h.Endpoint()]; ok {
		return ErrHostAlreadyExists
	}
	manager.hosts[h.Endpoint()] = h
	return nil
}

// RemoveHost removes a host if it exists.
func (manager *HostManager) RemoveHost(h Host) error {
	key := h.Endpoint()
	if _, ok := manager.hosts[key]; !ok {
		return ErrHostNotFound
	}
	delete(manager.hosts, key)
	return nil
}

// GetHosts returns a copy of the current hosts.
func (manager HostManager) GetHosts() []Host {
	var hosts []Host
	for _, h := range manager.hosts {
		hosts = append(hosts, h)
	}
	return hosts
}

// SetHosts overrides teh hosts for the provided input.
func (manager *HostManager) SetHosts(hosts []Host) {
	for _, h := range hosts {
		key := h.Endpoint()
		manager.hosts[key] = h
	}
}

// PrintHosts is a helper method to print the current hosts.
func (manager HostManager) PrintHosts() {
	for _, h := range manager.hosts {
		fmt.Printf("Host{endpoint: \"%s\"}\n", h.Endpoint())
	}
}
