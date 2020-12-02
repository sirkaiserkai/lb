package server

// AddHostRequest is the request structure to add a new host to the loadbalancer.
type AddHostRequest struct {
	Endpoint     string `json:"endpoint"`
	RegexPattern string `json:"regex"`
}

// RemoveHostRequest is the request struct to remove an existing host.
type RemoveHostRequest struct {
	Endpoint string `json:"endpoint"`
}
