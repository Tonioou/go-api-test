package client

import (
	"context"
	"fmt"
	"log"

	"github.com/Tonioou/go-todo-list/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joomcode/errorx"
	"github.com/sirupsen/logrus"
)

type PgClient struct {
	conn       *pgxpool.Pool
	connString string
}

func NewPgClient(ctx context.Context, cfg *config.Replica) *PgClient {
	pgClient := &PgClient{
		connString: createConnString(cfg),
	}
	connPool, err := pgxpool.New(ctx, pgClient.connString)
	if err != nil {
		log.Fatal(err)
	}

	pgClient.conn = connPool
	return pgClient
}

func createConnString(cfg *config.Replica) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database.Name,
	)
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

func (pg *PgClient) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	rows, err := pg.conn.Query(ctx, query, args)
	if err != nil {
		return nil, errorx.Decorate(err, "failed to query database")
	}
	return rows, nil
}

func (pg *PgClient) Exec(ctx context.Context, query string, args ...interface{}) error {
	_, err := pg.conn.Exec(ctx, query, args...)
	if err != nil {
		return errorx.Decorate(err, "failed to insert on database")
	}
	return nil
}

func (pg *PgClient) QueryRow(ctx context.Context, query string, args ...interface{}) (pgx.Row, error) {
	rows := pg.conn.QueryRow(ctx, query, args...)
	return rows, nil
}
