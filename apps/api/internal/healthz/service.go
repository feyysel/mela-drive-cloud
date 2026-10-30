package healthz

import (
	"fmt"

	"github.com/feyysel/mela-drive-cloud/internal/config"
	"github.com/gosimple/slug"
)

type Service struct {
	appCfg config.AppConfig
}

func NewService(appCfg config.AppConfig) *Service {
	return &Service{
		appCfg: appCfg,
	}
}

func (s *Service) Info() *InfoResponse {
	info := &InfoResponse{
		Service: fmt.Sprintf("%s-api", slug.Make(s.appCfg.Name)),
		Status:  StatusUp,
		Env:     s.appCfg.Env,
	}

	return info
}

func (s *Service) Healthz() *HealthzResponse {
	return &HealthzResponse{Status: StatusUp}
}
