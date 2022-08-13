package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInt(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		s    string
		dv   int
		want int
	}{
		{"15", -1, 15},
		{"-2", -1, -2},
		{"0", -1, 0},
		{"a1", -1, -1},
	}
	for _, test := range tests {
		got, _ := ParseInt(test.s, test.dv)
		assert.Equal(test.want, got)
	}
}
