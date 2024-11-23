package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/myrat012/test-work-song-lib/internal/model"
	"github.com/rs/zerolog"
)

type SongsRepo struct {
	*pgxpool.Pool
}

// NewUser -.
func NewSongs(pg *pgxpool.Pool) *SongsRepo {
	return &SongsRepo{pg}
}

func (r *SongsRepo) Create(ctx context.Context, m *model.Song) error {
	zLog := zerolog.Ctx(ctx).With().
		Str("unit", "internal.usecase.repo.SongsRepo").
		Str("method", "Create").Logger()

	sqlCommand := fmt.Sprintf("INSERT INTO %s (group_song, song, release_date, text, link) VALUES ($1, $2, $3, $4, $5);", model.TBL_NAME)
	fmt.Println(sqlCommand, m.Group, m.Song, m.ReleaseDate, m.Text, m.Link)
	_, err := r.Exec(ctx, sqlCommand, m.Group, m.Song, m.ReleaseDate, m.Text, m.Link)
	if err != nil {
		zLog.Err(err).Msgf("SongsRepo - Create - sqlCommand")
		return err
	}
	return nil
}

func (r *SongsRepo) Delete(ctx context.Context, id int) error {
	zLog := zerolog.Ctx(ctx).With().
		Str("unit", "internal.usecase.repo.SongsRepo").
		Str("method", "Delete").Logger()

	sqlCommand := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", model.TBL_NAME)

	_, err := r.Exec(ctx, sqlCommand, id)
	if err != nil {
		zLog.Err(err).Msgf("SongsRepo - Delete - sqlCommand")
		return err
	}
	return nil
}
