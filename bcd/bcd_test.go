package bcd

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumericToRightPaddedBCD_PAN(t *testing.T) {
	tests := []struct {
		name     string
		pan      string
		expected string // hex representation of expected BCD
		wantErr  bool
	}{
		{
			name:     "even length PAN (16 digits)",
			pan:      "1234123412341234",
			expected: "1234123412341234",
			wantErr:  false,
		},
		{
			name:     "odd length PAN (15 digits) - should append F",
			pan:      "123412341234123",
			expected: "123412341234123F",
			wantErr:  false,
		},
		{
			name:     "odd length PAN (13 digits) - should append F",
			pan:      "1234123412341",
			expected: "1234123412341F",
			wantErr:  false,
		},
		{
			name:     "odd length PAN (1 digit) - should append F",
			pan:      "1",
			expected: "1F",
			wantErr:  false,
		},
		{
			name:     "invalid PAN with letters",
			pan:      "123412341234A",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "empty PAN",
			pan:      "",
			expected: "",
			wantErr:  false, // NumericToRightPaddedBCD handles empty input
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NumericToRightPaddedBCD(tt.pan)

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

func TestNumericToRightPaddedBCD_ServiceData(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string // hex representation of expected BCD
		wantErr  bool
	}{
		{
			name:     "even length service data",
			input:    "601123",
			expected: "601123",
			wantErr:  false,
		},
		{
			name:     "odd length service data - should append F",
			input:    "60112",
			expected: "60112F",
			wantErr:  false,
		},
		{
			name:     "single digit - should append F",
			input:    "1",
			expected: "1F",
			wantErr:  false,
		},
		{
			name:     "three digits - should append F",
			input:    "601",
			expected: "601F",
			wantErr:  false,
		},
		{
			name:     "invalid service data with letters",
			input:    "601A",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "empty service data",
			input:    "",
			expected: "",
			wantErr:  false, // empty should be valid (results in empty BCD)
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

func TestNumericToRightPaddedBCD(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string // hex representation of expected BCD
		wantErr  bool
	}{
		{
			name:     "even length input",
			input:    "1234",
			expected: "1234",
			wantErr:  false,
		},
		{
			name:     "odd length input - should append F",
			input:    "123",
			expected: "123F",
			wantErr:  false,
		},
		{
			name:     "single digit - should append F",
			input:    "5",
			expected: "5F",
			wantErr:  false,
		},
		{
			name:     "empty input",
			input:    "",
			expected: "",
			wantErr:  false,
		},
		{
			name:     "invalid input with letters",
			input:    "12A",
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
