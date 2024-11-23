package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/myrat012/test-work-song-lib/internal/dto"
	"github.com/myrat012/test-work-song-lib/internal/usecase"
	"github.com/rs/zerolog"
)

type songsRouter struct {
	u usecase.SongsUseCase
}

func songsRegister(r *http.ServeMux, u usecase.SongsUseCase) {
	d := &songsRouter{u}

	r.HandleFunc("POST /songs", d.add)
	r.HandleFunc("DELETE /songs/{id}", d.delete)
	r.HandleFunc("PUT /songs/{id}", d.update)

}

func (route *songsRouter) add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	zLog := zerolog.Ctx(ctx).With().
		Str("remote-addr", GetRemoteAddress(r)).
		Str("uri", r.RequestURI).
		Str("method", r.Method).
		Str("handle", "SongCreate").Logger()

	var songCreateRequest dto.SongCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&songCreateRequest); err != nil {
		zLog.Err(err).Msg("format is wrong")
		responseWithCodeAndMessage(w, http.StatusBadRequest, "invalid request")
		return
	}
	defer r.Body.Close()

	if err := route.u.Create(ctx, songCreateRequest); err != nil {
		zLog.Err(err).Msg("SongsUseCase - error processing uc.repo.Create")
		responseWithCodeAndMessage(w, http.StatusInternalServerError, "error processing uc.repo.Create")
		return
	}
	responseWithCodeAndMessage(w, http.StatusOK, "Ok")
}

func (route *songsRouter) delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	zLog := zerolog.Ctx(ctx).With().
		Str("remote-addr", GetRemoteAddress(r)).
		Str("uri", r.RequestURI).
		Str("method", r.Method).
		Str("handle", "SongDelete").Logger()

	vars := r.PathValue("id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		zLog.Err(err).Msg("format is wrong")
		responseWithCodeAndMessage(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := route.u.Delete(ctx, id); err != nil {
		zLog.Err(err).Msg("SongsUseCase - error processing uc.repo.Delete")
		responseWithCodeAndMessage(w, http.StatusInternalServerError, "error processing uc.repo.Delete")
		return
	}
	responseWithCodeAndMessage(w, http.StatusOK, "Ok")
}
