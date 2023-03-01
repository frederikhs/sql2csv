package sql2csv

import (
	"context"
	"fmt"
	"github.com/tg/pgpass"
)

// TryDomain is current just set to hiper.dk since I work at hiper
var TryDomain = "hiper.dk"

func hostEquals(a, b string) bool {
	if a == b {
		return true
	}

	if fmt.Sprintf("%s.%s", a, TryDomain) == b {
		return true
	}

	return false
}

func getHosts() ([]pgpass.Entry, error) {
	f, err := pgpass.OpenDefault()
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var entries []pgpass.Entry
	er := pgpass.NewEntryReader(f)
	for er.Next() {
		e := er.Entry()
		entries = append(entries, e)
	}

	return entries, er.Err()
}

func NewConnectionFromPgPass(ctx context.Context, hostQuery string) (*Connection, error) {
	hosts, err := getHosts()
	if err != nil {
		return nil, fmt.Errorf("could not read pgpass file: %v", err)
	}

	for _, d := range hosts {
		match := hostEquals(hostQuery, d.Hostname)
		if !match {
			continue
		}

		return NewConnection(ctx, NewConnectionString(d.Username, d.Password, d.Hostname, d.Port, d.Database))
	}

	return nil, fmt.Errorf("could not find host: \"%s\" in your pgpass file", hostQuery)
}
