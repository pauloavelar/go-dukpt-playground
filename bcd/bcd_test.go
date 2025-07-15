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
		// Basic cases
		{
			name:     "empty input",
			input:    "",
			expected: "",
			wantErr:  false,
		},
		{
			name:     "single digit - should append F",
			input:    "5",
			expected: "5F",
			wantErr:  false,
		},
		{
			name:     "two digits even length",
			input:    "12",
			expected: "12",
			wantErr:  false,
		},
		{
			name:     "three digits odd length - should append F",
			input:    "123",
			expected: "123F",
			wantErr:  false,
		},
		{
			name:     "four digits even length",
			input:    "1234",
			expected: "1234",
			wantErr:  false,
		},
		// Real-life scenarios with longer inputs
		{
			name:     "6 digits even length",
			input:    "601123",
			expected: "601123",
			wantErr:  false,
		},
		{
			name:     "5 digits odd length - should append F",
			input:    "60112",
			expected: "60112F",
			wantErr:  false,
		},
		{
			name:     "13 digits odd length - should append F",
			input:    "1234123412341",
			expected: "1234123412341F",
			wantErr:  false,
		},
		{
			name:     "15 digits odd length - should append F",
			input:    "123412341234123",
			expected: "123412341234123F",
			wantErr:  false,
		},
		{
			name:     "16 digits even length",
			input:    "1234123412341234",
			expected: "1234123412341234",
			wantErr:  false,
		},
		// Error cases
		{
			name:     "invalid input with letters",
			input:    "12A",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "invalid input with letters in longer string",
			input:    "123412341234A",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "invalid input with letters mixed",
			input:    "601A",
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
				require.Equal(t, strings.ToUpper(tt.expected), resultHex)
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
