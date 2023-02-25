package sql2csv

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"time"
)

// extractColumns extracts each columns' name
func extractColumns(rows pgx.Rows) []string {
	fieldDescriptions := rows.FieldDescriptions()
	var columns []string
	for _, col := range fieldDescriptions {
		columns = append(columns, col.Name)
	}

	return columns
}

// interfaceToString converts an interface to a string
func interfaceToString(i interface{}) string {
	if i == nil {
		return ""
	}

	switch v := i.(type) {
	case time.Time:
		if v.Hour() == 0 && v.Minute() == 0 && v.Second() == 0 {
			return v.Format(time.DateOnly)
		}

		return v.Format(time.DateTime)
	default:
		return fmt.Sprint(v)
	}
}
