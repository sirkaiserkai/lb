package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// DummyServer is a simple server which hosts a /health endpoint.
type DummyServer struct {
	LoadBalancerEndpoint string
	Port                 string
	DummyHost            Host
}

// NewDummy returns a new dummy, dummy.
func NewDummy() DummyServer {
	port := "8082"
	return DummyServer{
		LoadBalancerEndpoint: "http://localhost:8081",
		Port:                 port,
		DummyHost:            NewHost("localhost:" + port),
	}
}

func (ds DummyServer) addToLoadBalancer() error {
	req := AddHostRequest{Endpoint: ds.DummyHost.Endpoint(), RegexPattern: "123"}
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	reader := strings.NewReader(string(b))
	resp, err := http.Post(ds.LoadBalancerEndpoint+"/add", "application/json", reader)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	modifyResponse := ModifyHostReponse{}
	if err := json.Unmarshal(body, &modifyResponse); err != nil {
		return err
	}
	return nil
}

// Health returns the current health of the service.
func (ds DummyServer) Health(w http.ResponseWriter, r *http.Request) {
	health := HealthStatusResponse{
		Health: HealthStatusOK,
	}
	b, err := json.Marshal(health)
	if err != nil {
		SetError(w, err)
	}
	fmt.Fprint(w, string(b))
}

// Run executes the dummy api.
func (ds DummyServer) Run() {
	fmt.Println("Running on port ", ds.Port)
	// TODO: Reach out to LB to add itself
	if err := ds.addToLoadBalancer(); err != nil {
		panic(err)
	}
	// Note: there's a race condition between a DS adding itself.
	http.HandleFunc("/health", ds.Health)
	log.Fatal(http.ListenAndServe(":"+ds.Port, nil))
}
