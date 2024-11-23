package http

import (
	"fmt"
	"net/http"

	"github.com/myrat012/test-work-song-lib/internal/usecase"
	"github.com/myrat012/test-work-song-lib/pkg/config"
)

func NewService(config config.AppConfig, useCases *usecase.UseCases) (*http.Server, error) {
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, config.Port),
		Handler: registerRouter(useCases),
	}
	return httpServer, nil
}
