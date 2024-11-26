package http

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/myrat012/test-work-song-lib/internal/usecase"
	"github.com/myrat012/test-work-song-lib/pkg/config"
	"github.com/rs/zerolog/log"
)

func NewService(config config.AppConfig, useCases *usecase.UseCases) (*http.Server, error) {
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, config.Port),
		Handler: registerRouter(useCases),
	}
	return httpServer, nil
}
func GetRemoteAddress(r *http.Request) string {
	if val := r.Header.Get("X-Forwarded-For"); val != "" {
		return val
	} else if val := r.Header.Get("X-Real-IP"); val != "" {
		return val
	} else {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Error().Err(err).
				Str("remote-addr", r.RemoteAddr).
				Msg("error parsing remote address")
			return "0.0.0.0"
		}
		return ip
	}
}

func responseWithCodeAndMessage(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	_, _ = fmt.Fprintln(w, message)
}

func jsonResponseWithCode(httpStatusCode int, w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Err(err).Msg("error in json.Encode")
	}
}
