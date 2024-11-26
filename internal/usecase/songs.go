package usecase

import (
	"context"

	"github.com/myrat012/test-work-song-lib/internal/dto"
	"github.com/myrat012/test-work-song-lib/internal/model"
	"github.com/myrat012/test-work-song-lib/internal/usecase/repo"
	"github.com/rs/zerolog"
)

type SongsUseCase struct {
	songsRepo *repo.SongsRepo
}

// NewUserUseCase -.
func NewSongsUseCase(sr *repo.SongsRepo) *SongsUseCase {
	return &SongsUseCase{sr}
}

func (uc *SongsUseCase) Create(ctx context.Context, createReq *dto.SongCreateRequest) error {
	zLog := zerolog.Ctx(ctx).With().
		Str("unit", "internal.usecase.SongsUseCase").
		Str("method", "Create").Logger()

	songModel := &model.Song{
		Group: createReq.Group,
		Song:  createReq.Song,
	}
	if err := uc.songsRepo.Create(ctx, songModel); err != nil {
		zLog.Err(err).Msg("SongsUseCase - error processing uc.repo.Create")
		return err
	}

	return nil
}

func (uc *SongsUseCase) Delete(ctx context.Context, id int) error {
	zLog := zerolog.Ctx(ctx).With().
		Str("unit", "internal.usecase.SongsUseCase").
		Str("method", "Delete").Logger()

	if err := uc.songsRepo.DeleteById(ctx, id); err != nil {
		zLog.Err(err).Msg("SongsUseCase - error processing uc.repo.DeleteById")
		return err
	}
	return nil
}

func (uc *SongsUseCase) Update(ctx context.Context, updateRequest *dto.SongRequest) error {
	zLog := zerolog.Ctx(ctx).With().
		Str("unit", "internal.usecase.SongsUseCase").
		Str("method", "Update").Logger()

	songOld, err := uc.songsRepo.GetById(ctx, updateRequest.ID)
	if err != nil {
		zLog.Err(err).Msg("SongsUseCase - error processing uc.repo.GetById")
		return err
	}
	if songOld == nil {
		zLog.Info().Msg("SongsUseCase - song not found")
		return nil
	}

	updatedSong := &model.Song{
		ID:          updateRequest.ID,
		Group:       updateRequest.Group,
		Song:        updateRequest.Song,
		ReleaseDate: updateRequest.ReleaseDate,
		Text:        updateRequest.Text,
		Link:        updateRequest.Link,
	}

	if err := uc.songsRepo.UpdateById(ctx, updatedSong); err != nil {
		zLog.Err(err).Msg("SongsUseCase - error processing uc.repo.UpdateById")
		return err
	}

	return nil
}

func (uc *SongsUseCase) Info(ctx context.Context, getRequest *dto.SongGetRequest) ([]*model.Song, error) {
	zLog := zerolog.Ctx(ctx).With().
		Str("unit", "internal.usecase.SongsUseCase").
		Str("method", "Info").Logger()

	if getRequest.Page <= 0 {
		getRequest.Page = 1
	}
	if getRequest.Limit <= 0 {
		getRequest.Limit = 10
	}

	a, err := uc.songsRepo.GetByFields(ctx, getRequest)
	if err != nil {
		zLog.Err(err).Msg("SongsUseCase - error processing uc.repo.GetByFields")
		return nil, err
	}
	return a, nil
}
