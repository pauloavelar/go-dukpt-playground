package bcd

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumericToRightPaddedBCD(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string // hex representation of expected BCD
		wantErr  bool
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "",
			wantErr:  false,
		},
		{
			name:     "odd length - should append F",
			input:    "5",
			expected: "5F",
			wantErr:  false,
		},
		{
			name:     "even length",
			input:    "12",
			expected: "12",
			wantErr:  false,
		},
		{
			name:     "16 chars real-life",
			input:    "1234123412341234",
			expected: "1234123412341234",
			wantErr:  false,
		},
		{
			name:     "invalid input valid hex",
			input:    "12A",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "invalid input not even hex",
			input:    "12G",
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NumericToRightPaddedBCD(tt.input)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				resultHex := strings.ToUpper(hex.EncodeToString(result))
				require.Equal(t, tt.expected, resultHex)
			}
		})
	}
}

// Test that regular BCD encoding still works as before (prepends 0 for odd lengths)
func TestStringToBCD_StillWorksAsExpected(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string // hex representation
	}{
		{
			name:     "even length string",
			input:    "1234",
			expected: "1234",
		},
		{
			name:     "odd length string - should prepend 0",
			input:    "123",
			expected: "0123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToBCD(tt.input)
			require.NoError(t, err)

			resultHex := strings.ToUpper(hex.EncodeToString(result))
			require.Equal(t, strings.ToUpper(tt.expected), resultHex)
		})
	}
}
