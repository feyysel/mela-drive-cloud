package healthz

type InfoResponse struct {
	Service string `json:"service"        binding:"required"`
	Status  Status `json:"status"         binding:"required"`
	Env     string `json:"environment"    binding:"required"`
}

type HealthzResponse struct {
	Status Status `json:"status" binding:"required"`
}
