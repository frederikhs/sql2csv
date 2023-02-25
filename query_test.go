package sql2csv

import (
	"testing"
)

func TestBadQuery(t *testing.T) {
	tests := []struct {
		q    string
		good bool
	}{
		{"", false},
		{" ", false},
		{"INSERT", false},
		{"insert", false},
		{"UPDATE", false},
		{"DELETE", false},
		{"SELECT", true},
		{"SELECT * FROM users", true},
		{"UPDATE users SET name='test' WHERE id IN (SELECT id FROM users WHERE admin = true)", false},
	}

	for _, test := range tests {
		r := checkIsSelect(test.q)
		if r != test.good {
			t.Errorf("expected %v, got %v for query: \"%s\"\n", test.good, r, test.q)
		}
	}
}
