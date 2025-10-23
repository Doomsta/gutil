package gutil

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
)

var ErrInvalidTimestamp = errors.New("no valid timestamp provided")

func ParseString(in string) (time.Time, error) {
	return parseString(in, time.Now())
}

func ParseAny(in any, now time.Time) (time.Time, error) {
	switch v := in.(type) {
	case time.Time:
		return v, nil

	case string:
		return parseString(v, now)

	case json.Number:
		if s := v.String(); s != "" {
			return parseString(s, now)
		}
		return time.Time{}, ErrInvalidTimestamp

	case int:
		return parseFromInt64(int64(v))
	case int8:
		return parseFromInt64(int64(v))
	case int16:
		return parseFromInt64(int64(v))
	case int32:
		return parseFromInt64(int64(v))
	case int64:
		return parseFromInt64(v)

	case uint:
		return parseFromInt64(int64(v))
	case uint8:
		return parseFromInt64(int64(v))
	case uint16:
		return parseFromInt64(int64(v))
	case uint32:
		return parseFromInt64(int64(v))
	case uint64:
		return parseFromInt64(int64(v))
	case float32:
		return parseFromFloat64(float64(v))
	case float64:
		return parseFromFloat64(v)
	}

	return time.Time{}, ErrInvalidTimestamp
}

func parseString(in string, now time.Time) (time.Time, error) {
	s := normalizeInput(in)

	if s == "" {
		return now, ErrInvalidTimestamp
	}

	if strings.HasPrefix(s, "now") {
		if s == "now" {
			return now, nil
		}
		if strings.HasPrefix(s, "now-") || strings.HasPrefix(s, "now+") {
			return parseNowOffset(s, now)
		}
	}

	if isAllDigits(s) {
		return parseUnixDigits(s)
	}

	if t, err := time.Parse(time.RFC3339Nano, s); err == nil {
		return t, nil
	}
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t, nil
	}

	return time.Time{}, ErrInvalidTimestamp
}

func normalizeInput(in string) string {
	return strings.Trim(strings.TrimSpace(in), "\"")
}

func parseNowOffset(s string, now time.Time) (time.Time, error) {
	var sign time.Duration = 1
	var rest string

	switch {
	case strings.HasPrefix(s, "now-"):
		sign = -1
		rest = strings.TrimPrefix(s, "now-")
	case strings.HasPrefix(s, "now+"):
		sign = 1
		rest = strings.TrimPrefix(s, "now+")
	default:
		return time.Time{}, ErrInvalidTimestamp
	}

	d, err := time.ParseDuration(rest)
	if err != nil {
		return time.Time{}, err
	}
	return now.Add(d * sign), nil
}

func isAllDigits(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func parseUnixDigits(s string) (time.Time, error) {
	s = rightPadToLen(s, 19, '0')
	switch len(s) {
	case 19:
		ns, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return time.Time{}, err
		}
		return parseFromInt64(ns)
	}
	return time.Time{}, ErrInvalidTimestamp
}

func rightPadToLen(s string, length int, pad byte) string {
	if len(s) >= length {
		return s
	}
	b := make([]byte, length)
	copy(b, s)
	for i := len(s); i < length; i++ {
		b[i] = pad
	}
	return string(b)
}

func parseFromInt64(v int64) (time.Time, error) {
	abs := v
	if abs < 0 {
		abs = -abs
	}

	return time.Unix(
		v/time.Second.Nanoseconds(),
		v%time.Second.Nanoseconds(),
	), nil
}

func parseFromFloat64(v float64) (time.Time, error) {
	av := v
	if av < 0 {
		av = -av
	}

	ns := v
	sec := int64(ns / 1_000_000_000)
	nsec := int64(ns - float64(sec)*1_000_000_000)
	return time.Unix(sec, nsec), nil
}
