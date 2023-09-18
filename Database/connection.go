package Database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"task/Config"
)

var Connection *pgx.Conn

func Connect() error {
	c, err := pgx.Connect(context.Background(), Config.Config.DbUri)
	if err == nil {
		Connection = c
		return nil
	}
	return err
}
