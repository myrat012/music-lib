package repo

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type SongsRepo struct {
	*pgxpool.Pool
}

// NewUser -.
func NewSongs(pg *pgxpool.Pool) *SongsRepo {
	return &SongsRepo{pg}
}
