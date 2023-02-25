package sql2csv

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"reflect"
	"time"
)

// createTypeMap creates a slice of interface which defines the types of each column
func createTypeMap(rows pgx.Rows, length int) []interface{} {
	valuePtr := make([]interface{}, length)
	rows.Next()
	values, _ := rows.Values()
	for i, v := range values {
		if v == nil {
			valuePtr[i] = nil
			continue
		}
		valuePtr[i] = reflect.New(reflect.TypeOf(v)).Interface() // allocate pointer to type
	}

	return valuePtr
}

// extractColumns extracts each columns' name
func extractColumns(rows pgx.Rows) []string {
	fieldDescriptions := rows.FieldDescriptions()
	var columns []string
	for _, col := range fieldDescriptions {
		columns = append(columns, col.Name)
	}

	return columns
}

// rowToStrings converts each column in the row to a string and returns a slice of strings
func rowToStrings(ptrs []interface{}, columns []string) []string {
	ss := []string{}
	for i := range columns {
		if ptrs[i] == nil {
			ss = append(ss, "")
			continue
		}
		var v interface{}
		val := reflect.ValueOf(ptrs[i]).Elem().Interface() // dereference pointer
		b, ok := val.([]byte)
		if ok {
			v = string(b)
		} else {
			v = val
		}
		ss = append(ss, interfaceToString(v))
	}

	return ss
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
