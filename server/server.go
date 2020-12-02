package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// NewLoadBalancer returns a new LoadBalancer instance.
func NewLoadBalancer(c LoadBalancerConfig) LoadBalancer {
	return LoadBalancer{
		C:    c,
		Port: "8081",
	}
}

// LoadBalancer is a simple LB server.
type LoadBalancer struct {
	C     LoadBalancerConfig
	Port  string
	hosts []Host
}

// Add supports adding a new endpoint to send traffic to.
func (lb *LoadBalancer) Add(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SetError(w, err)
	}
	addHostRequest := AddHostRequest{}
	if err := json.Unmarshal(body, &addHostRequest); err != nil {
		SetError(w, err)
	}
	h := NewHostForAddHostRequest(addHostRequest)
	lb.hosts = append(lb.hosts, h)
	SetJSONResponse(w, AddHostResponse{Status: "created"})
}

// Remove deletes an endpoint to send traffic to.
func (lb *LoadBalancer) Remove(w http.ResponseWriter, r *http.Request) {

}

// Route is the default router.
func (lb LoadBalancer) Route(w http.ResponseWriter, r *http.Request) {

}

// Health returns the current health of the service.
func (lb LoadBalancer) Health(w http.ResponseWriter, r *http.Request) {
	health := HealthStatusResponse{
		Health: "OK",
	}
	b, err := json.Marshal(health)
	if err != nil {
		SetError(w, err)
	}
	fmt.Fprint(w, string(b))
}

// Run runs the server instance.
func (lb LoadBalancer) Run() {
	http.HandleFunc("/add", lb.Add)
	http.HandleFunc("/remove", lb.Remove)
	http.HandleFunc("/health", lb.Health)
	log.Fatal(http.ListenAndServe(":"+lb.Port, nil))
}
