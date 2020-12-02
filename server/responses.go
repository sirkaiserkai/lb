package server

const (
	HealthStatusOK = "OK"
)

// HealthStatusResponse the response for a health request.
type HealthStatusResponse struct {
	Health string `json:"health"`
}

// ModifyHostReponse is the response upon successfully removing an existing host.
type ModifyHostReponse struct {
	Status string `json:"status"`
}
