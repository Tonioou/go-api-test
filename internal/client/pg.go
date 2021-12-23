package client

import (
	"context"
	"sync"

	"github.com/Tonioou/go-person-crud/internal/config"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joomcode/errorx"
	"github.com/sirupsen/logrus"
)

type Postgres interface {
	Ping(ctx context.Context) *errorx.Error
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, *errorx.Error)
}

type PgClient struct {
	conn *pgxpool.Pool
}

var pgRunOnce sync.Once
var pgClient *PgClient

func newPgClient() *errorx.Error {
	conn, err := pgxpool.Connect(context.Background(), config.GetConfig().Postgres.Url)
	if err != nil {
		errx := errorx.Decorate(err, "failed to connect")
		return errx
	}
	pgClient = &PgClient{
		conn: conn,
	}
	return nil
}

func GetPgClient() *PgClient {
	pgRunOnce.Do(func() {
		err := newPgClient()
		if err != nil {
			logrus.Error(err)
		}
	})
	return pgClient
}

func (pg *PgClient) Ping(ctx context.Context) *errorx.Error {
	if pg.conn == nil {
		err := newPgClient()
		if err != nil {
			return err
		}
	}
	query := "SELECT 1;"

	_, err := pg.conn.Query(ctx, query)
	if err != nil {
		logrus.Error(errorx.Decorate(err, "failed to query database"))
		err := newPgClient()
		if err != nil {
			return err
		}
	}
	return nil
}

func (pg *PgClient) getConnection(ctx context.Context) *pgxpool.Pool {
	err := pg.Ping(ctx)
	if err != nil {
		logrus.Error(errorx.Decorate(err, "failed to query database"))
	}
	return pg.conn
}
func (pg *PgClient) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, *errorx.Error) {
	pg.Ping(ctx)
	rows, err := pg.getConnection(ctx).Query(ctx, query, args)
	if err != nil {
		return nil, errorx.Decorate(err, "failed to query database")
	}
	return rows, nil
}
