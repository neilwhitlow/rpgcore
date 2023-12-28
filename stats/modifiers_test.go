package stats_test

import (
	"testing"

	"github.com/neilwhitlow/rpgcore/stats"
	"github.com/stretchr/testify/assert"
)

func TestGetModifierMap(t *testing.T) {
	tests := map[string]struct {
		input    int
		expected int
	}{
		"0":  {input: 0, expected: -4},
		"7":  {input: 7, expected: -1},
		"8":  {input: 8, expected: 0},
		"12": {input: 12, expected: 0},
		"13": {input: 13, expected: 1},
		"40": {input: 40, expected: 14},
	}

	modifierMap := stats.GetModifierMap()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := modifierMap[test.input]
			assert.Equal(t, test.expected, actual)
		})
	}
}
