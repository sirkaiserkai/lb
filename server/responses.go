package server

const (
	// HealthStatusOK the state for a operational service.
	HealthStatusOK = "OK"
	// HealthStatusWarning the state when a service is operational, however, with some warning condition.
	HealthStatusWarning = "WARNING"
	// HealthStatusBad the state when something dire is occuring in the service.
	HealthStatusBad = "BAD"
)

// HealthStatusResponse the response for a health request.
type HealthStatusResponse struct {
	Health string `json:"health"`
}

// ModifyHostReponse is the response upon successfully removing an existing host.
type ModifyHostReponse struct {
	Status string `json:"status"`
}
