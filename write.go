package sql2csv

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

// WriteQuery writes the results of executing the query to a csv file at the path given
func (c *Connection) WriteQuery(ctx context.Context, query *Query, path string, loggerFn func(ln string)) error {
	loggerFn("executing query")
	rows, err := c.db.Query(ctx, query.sql, query.args...)
	if err != nil {
		return err
	}

	loggerFn("creating output file: " + path)
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	w := csv.NewWriter(file)

	loggerFn("writing data")
	err = writeRows(w, rows, loggerFn)
	if err != nil {
		return fmt.Errorf("failed writing rows: %v", err)
	}

	w.Flush()

	if err = w.Error(); err != nil {
		return fmt.Errorf("failed flusing rows: %v", err)
	}

	if err = file.Close(); err != nil {
		return fmt.Errorf("failed closing file: %v", err)
	}

	rows.Close()
	if err = rows.Err(); err != nil {
		return fmt.Errorf("failed closing rows: %v", err)
	}

	loggerFn("done")

	return nil
}

// writeRows writes all the rows from the results into the writer
func writeRows(w *csv.Writer, rows pgx.Rows, loggerFn func(ln string)) error {
	columns := extractColumns(rows)

	// write header line of column name
	err := w.Write(columns)
	if err != nil {
		return err
	}

	i := 0
	for rows.Next() {
		i++
		rv, err := rows.Values()
		if err != nil {
			return err
		}

		var resultRow []string
		for _, raw := range rv {
			resultRow = append(resultRow, interfaceToString(raw))
		}

		err = w.Write(resultRow)
		if err != nil {
			return err
		}
	}

	loggerFn(fmt.Sprintf("wrote %d lines", i))

	return nil
}
