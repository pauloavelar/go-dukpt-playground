package bcd

import (
	"encoding/hex"
	"testing"
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
			expected: "123412341234123f",
			wantErr:  false,
		},
		{
			name:     "odd length PAN (13 digits) - should append F",
			pan:      "1234123412341",
			expected: "1234123412341f",
			wantErr:  false,
		},
		{
			name:     "odd length PAN (1 digit) - should append F",
			pan:      "1",
			expected: "1f",
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
				if err == nil {
					t.Errorf("PANToBCD() expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("PANToBCD() unexpected error: %v", err)
				return
			}
			
			resultHex := hex.EncodeToString(result)
			if resultHex != tt.expected {
				t.Errorf("PANToBCD() = %s, expected %s", resultHex, tt.expected)
			}
		})
	}
}

func TestPANByteLen(t *testing.T) {
	tests := []struct {
		name     string
		pan      string
		expected int
	}{
		{
			name:     "even length PAN (16 digits)",
			pan:      "1234123412341234",
			expected: 8,
		},
		{
			name:     "odd length PAN (15 digits)",
			pan:      "123412341234123",
			expected: 8, // 15 + 1 (F padding) = 16 digits = 8 bytes
		},
		{
			name:     "odd length PAN (13 digits)",
			pan:      "1234123412341",
			expected: 7, // 13 + 1 (F padding) = 14 digits = 7 bytes
		},
		{
			name:     "single digit PAN",
			pan:      "1",
			expected: 1, // 1 + 1 (F padding) = 2 digits = 1 byte
		},
		{
			name:     "empty PAN",
			pan:      "",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PANByteLen(tt.pan)
			if result != tt.expected {
				t.Errorf("PANByteLen() = %d, expected %d", result, tt.expected)
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
			if err != nil {
				t.Errorf("StringToBCD() unexpected error: %v", err)
				return
			}
			
			resultHex := hex.EncodeToString(result)
			if resultHex != tt.expected {
				t.Errorf("StringToBCD() = %s, expected %s", resultHex, tt.expected)
			}
		})
	}
}