package healthz

type Status string

const (
	StatusUp   Status = "up"
	StatusDown Status = "down"
)
