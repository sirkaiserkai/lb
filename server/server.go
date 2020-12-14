package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/sirkaiserkai/lb/server/background"
)

var (
	// ErrHostNotFound error returned when host does not exist.
	ErrHostNotFound = errors.New("host not found")
	// ErrHostAlreadyExists error returned when host already exists.
	ErrHostAlreadyExists = errors.New("host already exists")
)

// NewLoadBalancer returns a new LoadBalancer instance.
func NewLoadBalancer(c LoadBalancerConfig) LoadBalancer {
	hostManager := HostManager{}

	runnables := []background.Runnable{
		HealthCheck{manager: &hostManager},
	}

	return LoadBalancer{
		C:           c,
		Port:        "8081",
		hostManager: &hostManager,
		backgroundRunner: background.Runner{
			Cooldown:  time.Second * 10,
			Runnables: runnables,
		},
	}
}

// LoadBalancer is a simple LB server.
type LoadBalancer struct {
	C                LoadBalancerConfig
	Port             string
	hostManager      *HostManager
	backgroundRunner background.Runner
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
	if err := lb.hostManager.AddHost(h); err != nil {
		SetError(w, err)
	}
	SetJSONResponse(w, ModifyHostReponse{Status: "created"})
	lb.hostManager.PrintHosts()
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
	h := NewHost(rm.Endpoint)
	if err := lb.hostManager.RemoveHost(h); err != nil {
		SetError(w, err)
	}
	SetJSONResponse(w, ModifyHostReponse{Status: "removed"})
	lb.hostManager.PrintHosts()
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
	// Runs background processes.
	go lb.backgroundRunner.Run()
	fmt.Println("Running on port ", lb.Port)

	http.HandleFunc("/add", lb.Add)
	http.HandleFunc("/remove", lb.Remove)
	http.HandleFunc("/health", lb.Health)
	log.Fatal(http.ListenAndServe(":"+lb.Port, nil))
}
