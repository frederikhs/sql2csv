package sql2csv

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Connection struct {
	db *pgx.Conn
}

func NewConnectionString(user, pass, host, port, dbname string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, dbname)
	//	postgres://user:pass@host:port/dbname
}

// NewConnection creates a new connection object
func NewConnection(ctx context.Context, connectiongString string) (*Connection, error) {
	conn, err := pgx.Connect(ctx, connectiongString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return &Connection{db: conn}, nil
}

func (c *Connection) Close(ctx context.Context) error {
	return c.db.Close(ctx)
}
