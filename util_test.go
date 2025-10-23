package gutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClamp(t *testing.T) {
	tests := [][4]int{
		{0, 0, 0, 0},
		{0, 22, 10, 10},
		{5, 0, 10, 5},
		{5, 100, 10, 10},
		{53333, 100, 10, 100},
		{53333, 9, 10, 10},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d-%d-%d", tt[0], tt[1], tt[2]), func(t *testing.T) {
			assert.Equalf(t, tt[3], Clamp(tt[0], tt[1], tt[2]), "Clamp(%v, %v, %v)", tt[0], tt[1], tt[2])
		})
	}
}
