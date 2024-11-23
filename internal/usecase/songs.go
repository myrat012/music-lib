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

func (uc *SongsUseCase) Create(ctx context.Context, createReq dto.SongCreateRequest) error {
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

	if err := uc.songsRepo.Delete(ctx, id); err != nil {
		zLog.Err(err).Msg("SongsUseCase - error processing uc.repo.Delete")
		return err
	}
	return nil
}
