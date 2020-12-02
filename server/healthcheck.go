package server

import "fmt"

// HealthCheck performs a health check against a set of hosts.
type HealthCheck struct {
	manager *HostManager
}

// Run iterates over the available hosts and performs a health check.
func (hc HealthCheck) Run() error {
	fmt.Println("running health check")
	for _, h := range hc.manager.GetHosts() {
		r, err := h.Health()
		if err != nil {
			return err
		}
		if r.Health != HealthStatusOK {
			return fmt.Errorf("Health status not '%s'. Received: '%s'", HealthStatusOK, r.Health)
		}
	}
	return nil
}
