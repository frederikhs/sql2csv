package sql2csv

import (
	"context"
	"errors"
	"strings"
)

var badKeyWords = []string{
	"insert",
	"update",
	"delete",
}

type Query struct {
	ctx  context.Context
	sql  string
	args []any
}

// NewQuery create a new query object for later use when executing a query against the database
func NewQuery(ctx context.Context, sql string, args ...any) (*Query, error) {
	if !checkIsSelect(sql) {
		return nil, errors.New("query is not a select statement")
	}

	return &Query{
		ctx:  ctx,
		sql:  sql,
		args: args,
	}, nil
}

// checkIsSelect is a naive check for queries that will alter the database.
// This tool is not for modifying the database, and this is just for stupidity’s sake and not a real safety net.
func checkIsSelect(q string) bool {
	ql := strings.ToLower(q)

	if !strings.Contains(ql, "select") {
		return false
	}

	for _, kw := range badKeyWords {
		if strings.Contains(ql, kw) {
			return false
		}
	}

	return true
}
