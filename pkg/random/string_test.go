package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewString(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{
			name:   "Length = 1",
			length: 1,
		},
		{
			name:   "Length = 5",
			length: 5,
		},
		{
			name:   "Length = 10",
			length: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := NewString(tt.length)
			assert.Len(t, str, tt.length)
			str2 := NewString(tt.length)
			assert.NotEqual(t, str, str2)
			dictionary := "abcdefghijklmnopqrstuvwxyz" +
				"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
				"0123456789"
			for _, c := range str {
				assert.Contains(t, dictionary, string(c))
			}
		})
	}
}
