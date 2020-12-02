package server

// HostManager manages host instances. TODO: Make it safe for concurrent access.
type HostManager struct {
	hosts []Host
}

// AddHost adds a new host instance to the LoadBalancer.
func (manager *HostManager) AddHost(h Host) error {
	// TODO: Improve runtime using better data structure.
	// Since we're using endpoint as unique identifier, we should use a map.
	for _, existingHost := range manager.hosts {
		if existingHost.EqualsHost(h) {
			return ErrHostAlreadyExists
		}
	}
	manager.hosts = append(manager.hosts, h)
	return nil
}

// RemoveHost removes a host if it exists.
func (manager *HostManager) RemoveHost(h Host) error {
	indexToRemove := -1
	for i, existingHost := range manager.hosts {
		if existingHost.EqualsHost(h) {
			indexToRemove = i
			break
		}
	}
	if indexToRemove < 0 {
		return ErrHostNotFound
	}

	manager.hosts = append(manager.hosts[:indexToRemove], manager.hosts[indexToRemove+1:]...)
	return nil
}
