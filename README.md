# sql2csv

[![Quality](https://goreportcard.com/badge/github.com/frederikhs/sql2csv)](https://goreportcard.com/report/github.com/frederikhs/sql2csv)
[![Test](https://github.com/frederikhs/sql2csv/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/frederikhs/sql2csv/actions/workflows/test.yml)

*Go library for generating csv files based on SQL queries*

## Usage

```go
package main

import (
	"context"
	"github.com/frederikhs/sql2csv"
)

const (
	host     = "db.example.com"
	port     = "5432"
	user     = "example"
	password = "example"
	dbname   = "example"
)

func main() {
	conn, err := sql2csv.NewConnection(context.Background(), user, password, host, port, dbname)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	query := sql2csv.NewQuery(context.Background(), "SELECT * FROM public.users LIMIT $1", 10)

	// writes results to ./results.csv
	err = conn.WriteQuery(query, "results")
	if err != nil {
		panic(err)
	}
}
```
