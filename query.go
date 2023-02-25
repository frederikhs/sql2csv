package sql2csv

import "context"

type Query struct {
	ctx  context.Context
	sql  string
	args []any
}

// NewQuery create a new query object for later use when executing a query against the database
func NewQuery(ctx context.Context, sql string, args ...any) *Query {
	return &Query{
		ctx:  ctx,
		sql:  sql,
		args: args,
	}
}
