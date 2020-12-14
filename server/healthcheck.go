package server

import (
	"fmt"
	"log"
)

// HealthCheck performs a health check against a set of hosts.
type HealthCheck struct {
	manager *HostManager
}

// Run iterates over the available hosts and performs a health check.
func (hc HealthCheck) Run() error {
	// TODO: Update Run to return a slice of errors. This will allow the HealthCheck instance to validate all the hosts before exiting.
	log.Println("Running health checks")
	for _, h := range hc.manager.GetHosts() {
		r, err := h.Health()
		if err != nil {
			return err
		}
		if r.Health != HealthStatusOK {
			return fmt.Errorf("Health status not '%s'. Received: '%s'", HealthStatusOK, r.Health)
		}

		log.Printf("Health OK: '%s'\n", h.Endpoint())
	}
	return nil
}
