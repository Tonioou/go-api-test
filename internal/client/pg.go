package client

import (
	"context"

	"github.com/Tonioou/go-todo-list/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joomcode/errorx"
	"github.com/sirupsen/logrus"
)

type PgClient struct {
	conn *pgxpool.Pool
	cfg  *config.Configs
}

func NewPgClient(cfg *config.Configs) *PgClient {
	pgClient := &PgClient{
		cfg: cfg,
	}

	return pgClient
}

func (pg *PgClient) Ping(ctx context.Context) error {
	var (
		errx        *errorx.Error
		query       = "SELECT 1;"
		queryResult = 0
	)

	err := pg.conn.QueryRow(ctx, query).Scan(&queryResult)
	if err != nil {
		errx = errorx.Decorate(err, "failed to query database")
		logrus.Error(errx)
	}

	logrus.Error(errorx.Decorate(err, "failed to reconnect to db"))
	return errx
}

func (pg *PgClient) getConnection(ctx context.Context) (*pgxpool.Pool, error) {
	err := pg.Ping(ctx)
	if err != nil {
		errx := errorx.InternalError.New("failed to query database")
		return nil, errx
	}
	return pg.conn, nil
}

func (pg *PgClient) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	connection, errx := pg.getConnection(ctx)
	if errx != nil {
		return nil, errx
	}
	rows, err := connection.Query(ctx, query, args)
	if err != nil {
		return nil, errorx.Decorate(err, "failed to query database")
	}
	return rows, nil
}

func (pg *PgClient) Exec(ctx context.Context, query string, args ...interface{}) error {
	connection, errx := pg.getConnection(ctx)
	if errx != nil {
		return errx
	}
	_, err := connection.Exec(ctx, query, args...)
	if err != nil {
		return errorx.Decorate(err, "failed to insert on database")
	}
	return nil
}

func (pg *PgClient) QueryRow(ctx context.Context, query string, args ...interface{}) (pgx.Row, error) {
	connection, errx := pg.getConnection(ctx)
	if errx != nil {
		return nil, errx
	}
	rows := connection.QueryRow(ctx, query, args...)
	return rows, nil
}
