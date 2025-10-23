package gutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPage(t *testing.T) {
	tests := []struct {
		name     string
		offset   int
		length   int
		list     []int
		expected []int
	}{
		{
			name:     "easy",
			offset:   2,
			length:   2,
			list:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{2, 3},
		}, {
			name:     "length zero -> until end",
			offset:   3,
			length:   0,
			list:     []int{10, 11, 12, 13, 14},
			expected: []int{13, 14},
		}, {
			name:     "offset beyond len -> empty",
			offset:   10,
			length:   5,
			list:     []int{1, 2, 3},
			expected: []int{},
		}, {
			name:     "negative offset/length -> clamped",
			offset:   -3,
			length:   -2,
			list:     []int{7, 8, 9},
			expected: []int{},
		}, {
			name:     "length exceeds end -> clamped",
			offset:   1,
			length:   99,
			list:     []int{5, 6, 7},
			expected: []int{6, 7},
		},
	}

	for _, tt := range tests {
		tc := tt
		t.Run(tt.name, func(t *testing.T) {
			got := Page(tc.list, tc.offset, tc.length, nil)
			assert.Equal(t, tc.expected, got)
		})
	}
}
