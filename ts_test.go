package gutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseTS(t *testing.T) {
	now := time.Date(2022, time.April, 6, 8, 18, 47, 0, time.UTC)

	cases := []struct {
		input    string
		expected time.Time
	}{
		{"now", now},
		{"now-1h", now.Add(-1 * time.Hour)},
		{"now-25h", now.Add(-25 * time.Hour)},
		{"now+1h", now.Add(time.Hour * 1)},
		{"now+5h10m15s", now.Add(5 * time.Hour).Add(10 * time.Minute).Add(15 * time.Second)},
		{"164923312", now.Add(-time.Second * 7)},
		{"1649233127", now},
		{"16492331270", now},
		{"164923312700", now},
		{"1649233127000", now},
		{"16492331270000", now},
		{"164923312700000", now},
		{"1649233127000000", now},
		{"16492331270000000", now},
		{"164923312700000000", now},
		{"1649233127901173114", now.Add(901173114 * time.Nanosecond)},
		{"2022-04-06T10:18:47+02:00", now},
		{"2022-04-06T10:18:47.901173114+02:00", now.Add(901173114 * time.Nanosecond)},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			actual, err := parseString(c.input, now)
			require.NoError(t, err)
			assert.Equal(t, actual.UTC(), c.expected)
		})
	}
}
