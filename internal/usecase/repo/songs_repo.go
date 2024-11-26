package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/myrat012/test-work-song-lib/internal/dto"
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

func (r *SongsRepo) DeleteById(ctx context.Context, id int) error {
	zLog := zerolog.Ctx(ctx).With().
		Str("unit", "internal.usecase.repo.SongsRepo").
		Str("method", "DeleteById").Logger()

	sqlCommand := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", model.TBL_NAME)

	_, err := r.Exec(ctx, sqlCommand, id)
	if err != nil {
		zLog.Err(err).Msgf("SongsRepo - DeleteById - sqlCommand")
		return err
	}
	return nil
}

func (r *SongsRepo) GetById(ctx context.Context, id int) (*model.Song, error) {
	zLog := zerolog.Ctx(ctx).With().
		Str("unit", "internal.usecase.repo.SongsRepo").
		Str("method", "GetById").Logger()

	sqlCommand := fmt.Sprintf("SELECT id, group_song, song, release_date, text, link FROM %s WHERE id = $1;", model.TBL_NAME)

	row := r.QueryRow(ctx, sqlCommand, id)

	song := &model.Song{}

	if err := row.Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link); err != nil {
		zLog.Err(err).Msgf("SongsRepo - GetById - sqlCommand")
		return nil, err
	}

	return song, nil
}

func (r *SongsRepo) UpdateById(ctx context.Context, m *model.Song) error {
	zLog := zerolog.Ctx(ctx).With().
		Str("unit", "internal.usecase.repo.SongsRepo").
		Str("method", "UpdateById").Logger()

	sqlCommand := fmt.Sprintf("UPDATE %s SET group_song=$1, song=$2, release_date=$3, text=$4, link=$5 WHERE id=$6;", model.TBL_NAME)

	_, err := r.Exec(ctx, sqlCommand, m.Group, m.Song, m.ReleaseDate, m.Text, m.Link, m.ID)
	if err != nil {
		zLog.Err(err).Msgf("SongsRepo - UpdateById - sqlCommand")
		return err
	}
	return nil
}

// by 'group' & 'songs'
func (r *SongsRepo) GetByFields(ctx context.Context, getRequest *dto.SongGetRequest) ([]*model.Song, error) {
	zLog := zerolog.Ctx(ctx).With().
		Str("unit", "internal.usecase.repo.SongsRepo").
		Str("method", "GetByFields").Logger()

	sqlCommand := fmt.Sprintf("SELECT id, group_song, song, release_date, text, link FROM %s WHERE ", model.TBL_NAME)

	var filters []string
	var args []interface{}
	argIx := 1
	if getRequest.Group != "" {
		filters = append(filters, fmt.Sprintf("group_song ILIKE $%d", argIx))
		args = append(args, "%"+getRequest.Group+"%")
		argIx++
	}

	if getRequest.Song != "" {
		filters = append(filters, fmt.Sprintf("song ILIKE $%d", argIx))
		args = append(args, "%"+getRequest.Song+"%")
		argIx++
	}

	if len(filters) == 1 {
		sqlCommand += filters[0]
	}
	if len(filters) == 2 {
		sqlCommand += strings.Join(filters, " AND ")
	}

	sqlCommand += fmt.Sprintf(" LIMIT $%d OFFSET $%d;", argIx, argIx+1)
	args = append(args, getRequest.Limit, getRequest.Page)

	rows, err := r.Query(ctx, sqlCommand, args...)
	if err != nil {
		zLog.Err(err).Msgf("SongsRepo - GetByFields - sqlCommand")
		return nil, err
	}

	a := make([]*model.Song, 0)
	for rows.Next() {
		song := &model.Song{}
		if err := rows.Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link); err != nil {
			zLog.Err(err).Msgf("SongsRepo - GetByFields - rows.Scan")
			return nil, err
		}
		a = append(a, song)
	}

	return a, nil
}
