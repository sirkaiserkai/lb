package server

// HealthStatusResponse the response for a health request.
type HealthStatusResponse struct {
	Health string `json:"health"`
}

//AddHostResponse the response for adding an ew host.
type AddHostResponse struct {
	Status string `json:"status"`
}
