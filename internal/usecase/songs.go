package usecase

import "github.com/myrat012/test-work-song-lib/internal/usecase/repo"

type SongsUseCase struct {
	songsRepo *repo.SongsRepo
}

// NewUserUseCase -.
func NewSongsUseCase(sr *repo.SongsRepo) *SongsUseCase {
	return &SongsUseCase{sr}
}
