# sql2csv

[![GoDoc](https://godoc.org/github.com/frederikhs/sql2csv?status.svg)](https://godoc.org/github.com/frederikhs/sql2csv)
[![Quality](https://goreportcard.com/badge/github.com/frederikhs/sql2csv)](https://goreportcard.com/report/github.com/frederikhs/sql2csv)
[![Test](https://github.com/frederikhs/sql2csv/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/frederikhs/sql2csv/actions/workflows/test.yml)
[![Release](https://img.shields.io/github/v/release/frederikhs/sql2csv.svg)](https://github.com/frederikhs/sql2csv/releases/latest)
[![License](https://img.shields.io/github/license/frederikhs/sql2csv)](LICENSE)

*Go library for generating csv files based on SQL queries*

Only supporting PostgreSQL

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

	query, err := sql2csv.NewQuery(context.Background(), "SELECT * FROM public.users LIMIT $1", 10)
	if err != nil {
		panic(err)
	}

	// writes results to ./results.csv
	err = conn.WriteQuery(query, "results")
	if err != nil {
		panic(err)
	}
}
```
