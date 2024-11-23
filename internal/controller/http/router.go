package http

import (
	"net/http"

	"github.com/myrat012/test-work-song-lib/internal/usecase"
)

func registerRouter(u *usecase.UseCases) *http.ServeMux {
	mux := http.NewServeMux()
	songsRegister(mux, *u.SongsUseCase)
	return mux
}
