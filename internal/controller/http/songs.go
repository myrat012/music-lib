package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/myrat012/test-work-song-lib/internal/docs"
	"github.com/myrat012/test-work-song-lib/internal/dto"
	"github.com/myrat012/test-work-song-lib/internal/usecase"
	"github.com/myrat012/test-work-song-lib/pkg/util"
	"github.com/rs/zerolog"
	httpSwagger "github.com/swaggo/http-swagger"
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

	r.Handle("/swagger/", httpSwagger.WrapHandler)
}

// Add godoc
// @Summary      Добавление новой песни
// @Description  Добавляет новую песню в библиотеку
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        song  body      dto.SongCreateRequest  true  "Данные песни"
// @Success      200   {string}  string                 "Ok"
// @Router       /songs [post]
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

// Delete godoc
// @Summary      Удаление песни
// @Description  Удаляет песню по её идентификатору
// @Tags         songs
// @Param        id   path      int     true  "Идентификатор песни"
// @Success      200  {string}  string  "Ok"
// @Router       /songs/{id} [delete]
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

// GetInfo godoc
// @Summary      Получение песни
// @Description  Возвращает песни с поддержкой пагинации
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        group  query   string  false   "Название группы"
// @Param        song   query   string  false   "Название песни"
// @Param        page   query  int  true  "Номер страницы (по умолчанию 1)"
// @Param        limit  query  int  true  "Количество куплетов на странице (по умолчанию 2)"
// @Success      200    {object}  []model.Song
// @Router       /info [get]
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

// GetLyrics godoc
// @Summary      Получение текста песни
// @Description  Возвращает текст песни с поддержкой пагинации по куплетам
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        id     path   int  true  "ID песни"
// @Param        page   query  int  true  "Номер страницы (по умолчанию 1)"
// @Param        limit  query  int  true  "Количество куплетов на странице (по умолчанию 2)"
// @Success      200    {object}    []string
// @Router       /songs/{id}/lyrics [get]
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
