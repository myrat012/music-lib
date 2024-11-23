package http

import (
	"fmt"
	"net/http"

	"github.com/myrat012/test-work-song-lib/internal/usecase"
)

type songsRouter struct {
	u usecase.SongsUseCase
}

func songsRegister(r *http.ServeMux, u usecase.SongsUseCase) {
	d := &songsRouter{u}

	r.HandleFunc("POST /songs", d.add)

}

func (route *songsRouter) add(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "got path\n")

}
