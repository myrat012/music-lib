package usecase

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/myrat012/test-work-song-lib/internal/usecase/repo"
)

type UseCases struct {
	SongsUseCase *SongsUseCase
}

func LoadUseCases(pool *pgxpool.Pool) *UseCases {
	songsRepo := repo.NewSongs(pool)

	return &UseCases{
		SongsUseCase: NewSongsUseCase(songsRepo),
	}
}
