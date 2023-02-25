package sql2csv

import (
	"encoding/csv"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

// WriteQuery writes the results of executing the query to a csv file at the path given
func (c *Connection) WriteQuery(query *Query, path string) error {
	rows, err := c.db.Query(query.ctx, query.sql, query.args...)

	if err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("%s.csv", path))
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	w := csv.NewWriter(file)

	err = writeRows(w, rows)
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

	return nil
}

// writeRows writes all the rows from the results into the writer
func writeRows(w *csv.Writer, rows pgx.Rows) error {
	columns := extractColumns(rows)

	// write header line of column name
	err := w.Write(columns)
	if err != nil {
		return err
	}

	for rows.Next() {
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

	return nil
}
