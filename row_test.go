package sql2csv

import (
	"reflect"
	"testing"
	"time"
)

func TestInterfaceToString(t *testing.T) {
	tests := []struct {
		i interface{}
		s string
	}{
		{"", ""},
		{true, "true"},
		{false, "false"},
		{"string", "string"},
		{int64(1), "1"},
		{float64(1), "1"},
		{1, "1"},
		{1.2345, "1.2345"},
		{nil, ""},
		{time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local), "2009-11-10 23:00:00"},
		{time.Date(2009, time.November, 10, 23, 0, 0, 1, time.Local), "2009-11-10 23:00:00"},
		{time.Date(2009, time.November, 10, 0, 0, 0, 0, time.Local), "2009-11-10"},
		{time.Date(2009, time.November, 10, 0, 0, 0, 1, time.Local), "2009-11-10"},
	}

	for _, test := range tests {
		r := interfaceToString(test.i)
		if r != test.s {
			t.Fatalf("expected: \"%s\", got \"%s\" for type: %v", test.s, r, reflect.TypeOf(test.i))
		}
	}
}
