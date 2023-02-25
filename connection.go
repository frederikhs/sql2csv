package sql2csv

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Connection struct {
	db *pgx.Conn
}

// NewConnection creates a new connection object
func NewConnection(ctx context.Context, user, pass, host, port, dbname string) (*Connection, error) {
	urlExample := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, dbname)
	conn, err := pgx.Connect(ctx, urlExample)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v\n", err)
	}

	return &Connection{db: conn}, nil
}

func (c *Connection) Close(ctx context.Context) error {
	return c.db.Close(ctx)
}
