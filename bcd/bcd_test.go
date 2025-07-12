package bcd

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPANToBCD(t *testing.T) {
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
			wantErr:  true, // empty PAN should be invalid
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := PANToBCD(tt.pan)
			
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