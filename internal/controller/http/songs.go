package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/myrat012/test-work-song-lib/internal/dto"
	"github.com/myrat012/test-work-song-lib/internal/usecase"
	"github.com/myrat012/test-work-song-lib/pkg/util"
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
	r.HandleFunc("GET /info", d.info)
	r.HandleFunc("GET /songs/{id}/lyrics", d.lyrics)
}

func (route *songsRouter) add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	zLog := zerolog.Ctx(ctx).With().
		Str("remote-addr", GetRemoteAddress(r)).
		Str("uri", r.RequestURI).
		Str("method", r.Method).
		Str("handle", "SongCreate").Logger()

	var songCreateRequest *dto.SongCreateRequest
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

func (route *songsRouter) update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	zLog := zerolog.Ctx(ctx).With().
		Str("remote-addr", GetRemoteAddress(r)).
		Str("uri", r.RequestURI).
		Str("method", r.Method).
		Str("handle", "SongUpdate").Logger()

	var songUpdateRequest *dto.SongRequest
	if err := json.NewDecoder(r.Body).Decode(&songUpdateRequest); err != nil {
		zLog.Err(err).Msg("format is wrong")
		responseWithCodeAndMessage(w, http.StatusBadRequest, "invalid request")
		return
	}
	defer r.Body.Close()

	err := route.u.Update(ctx, songUpdateRequest)
	if err != nil {
		zLog.Err(err).Msg("SongsUseCase - error processing uc.repo.Delete")
		responseWithCodeAndMessage(w, http.StatusInternalServerError, "error processing uc.repo.Delete")
		return
	}

	responseWithCodeAndMessage(w, http.StatusOK, "Ok")
}

func (route *songsRouter) info(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	zLog := zerolog.Ctx(ctx).With().
		Str("remote-addr", GetRemoteAddress(r)).
		Str("uri", r.RequestURI).
		Str("method", r.Method).
		Str("handle", "SongInfo").Logger()

	getRequest := &dto.SongGetRequest{}

	if err := getRequest.ToStruct(r.URL.Query().Get("group"), r.URL.Query().Get("song"), r.URL.Query().Get("page"), r.URL.Query().Get("limit")); err != nil {
		zLog.Err(err).Msg("limit or page format is wrong")
		responseWithCodeAndMessage(w, http.StatusBadRequest, "invalid request")
		return
	}

	getRequest.Page, getRequest.Limit = util.PageToLimitOffset(getRequest.Page, getRequest.Limit)

	a, err := route.u.Info(ctx, getRequest)
	if err != nil {
		responseWithCodeAndMessage(w, http.StatusBadRequest, "invalid request")
		return
	}

	jsonResponseWithCode(http.StatusOK, w, a)
}

func (route *songsRouter) lyrics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	zLog := zerolog.Ctx(ctx).With().
		Str("remote-addr", GetRemoteAddress(r)).
		Str("uri", r.RequestURI).
		Str("method", r.Method).
		Str("handle", "SongLyrics").Logger()

	vars := r.PathValue("id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		zLog.Err(err).Msg("format is wrong")
		responseWithCodeAndMessage(w, http.StatusBadRequest, "invalid request")
		return
	}
	p, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		zLog.Err(err).Msg("format is wrong")
		responseWithCodeAndMessage(w, http.StatusBadRequest, "invalid request")
		return
	}
	l, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		zLog.Err(err).Msg("format is wrong")
		responseWithCodeAndMessage(w, http.StatusBadRequest, "invalid request")
		return
	}

	page, limit := util.PageToLimitOffset(p, l)

	sp, err := route.u.GetSongText(ctx, id, page, limit)
	if err != nil {
		zLog.Err(err).Msg("SongsUseCase - error processing uc.repo.GetSongText")
		responseWithCodeAndMessage(w, http.StatusInternalServerError, "error processing uc.repo.GetSongText")
		return
	}

	jsonResponseWithCode(http.StatusOK, w, sp)
}
