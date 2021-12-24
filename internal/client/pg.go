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
	Exec(ctx context.Context, query string, args ...interface{}) *errorx.Error
	QueryRow(ctx context.Context, query string, args ...interface{}) (pgx.Row, *errorx.Error)
}

type PgClient struct {
	conn *pgxpool.Pool
}

var pgRunOnce sync.Once
var pgClient *PgClient

func newPgClient() *errorx.Error {
	conn, err := pgxpool.Connect(context.Background(), config.GetConfig().Postgres.Url)
	if err != nil {
		errx := errorx.Decorate(err, "failed to connect to Database")
		return errx
	}
	pgClient = &PgClient{
		conn: conn,
	}
	return nil
}

func GetPgClient() *PgClient {
	pgRunOnce.Do(func() {
		pgClient = &PgClient{}
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

func (pg *PgClient) getConnection(ctx context.Context) (*pgxpool.Pool, *errorx.Error) {
	err := pg.Ping(ctx)
	if err != nil {
		errx := errorx.InternalError.New("failed to query database")
		return nil, errx
	}
	return pg.conn, nil
}

func (pg *PgClient) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, *errorx.Error) {
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

func (pg *PgClient) Exec(ctx context.Context, query string, args ...interface{}) *errorx.Error {
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

func (pg *PgClient) QueryRow(ctx context.Context, query string, args ...interface{}) (pgx.Row, *errorx.Error) {
	connection, errx := pg.getConnection(ctx)
	if errx != nil {
		return nil, errx
	}
	rows := connection.QueryRow(ctx, query, args...)
	return rows, nil
}
