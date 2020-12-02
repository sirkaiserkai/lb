package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	// ErrHostDoesNotExist error returned when host does not exist.
	ErrHostDoesNotExist = errors.New("host does not exist")
	// ErrHostAlreadyExists error returned when host already exists.
	ErrHostAlreadyExists = errors.New("host already exists")
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

// AddHost adds a new host instance to the LoadBalancer.
func (lb *LoadBalancer) AddHost(h Host) error {
	// TODO: return error if host already exists.
	lb.hosts = append(lb.hosts, h)
	return nil
}

// RemoveHost removes a host if it exists.
func (lb *LoadBalancer) RemoveHost(endpoint string) error {
	// i := 0
	// for i < len(lb.hosts) {
	// 	h :=
	// 	if lb.hosts[i].Endpoint() == endpoint {

	// 	}
	// 	i++
	// }
	return errors.New("Unimplemented")
}

// Add supports adding a new endpoint to send traffic to.
func (lb *LoadBalancer) Add(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SetError(w, err)
	}
	defer r.Body.Close()
	addHostRequest := AddHostRequest{}
	if err := json.Unmarshal(body, &addHostRequest); err != nil {
		SetError(w, err)
	}
	h := NewHostForAddHostRequest(addHostRequest)
	if err := lb.AddHost(h); err != nil {
		SetError(w, err)
	}
	SetJSONResponse(w, AddHostResponse{Status: "created"})
}

// Remove deletes an endpoint to send traffic to.
func (lb *LoadBalancer) Remove(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SetError(w, err)
	}
	defer r.Body.Close()
	rm := RemoveHostRequest{}
	if err := json.Unmarshal(body, &rm); err != nil {
		SetError(w, err)
	}
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
