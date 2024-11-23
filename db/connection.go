package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/myrat012/test-work-song-lib/pkg/config"
)

func NewPool(dbConf *config.DbConfig) (*pgxpool.Pool, error) {
	conf, err := pgxpool.ParseConfig(dbConf.ConnString)
	if err != nil {
		return nil, err
	}

	conf.MaxConns = int32(dbConf.MaxConns)
	conf.ConnConfig.ConnectTimeout = time.Duration(dbConf.ConnTimeout) * time.Second

	pool, err := pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
