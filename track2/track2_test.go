package track2

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestData_FormatISO_PANPadding(t *testing.T) {
	tests := []struct {
		name        string
		pan         string
		expPANInHex string // expected PAN portion in the hex output
	}{
		{
			name:        "even length PAN (16 digits)",
			pan:         "1234123412341234",
			expPANInHex: "1234123412341234", // no padding needed
		},
		{
			name:        "odd length PAN (15 digits) - should append F",
			pan:         "123412341234123",
			expPANInHex: "123412341234123F", // F padding appended
		},
		{
			name:        "odd length PAN (13 digits) - should append F",
			pan:         "1234123412341",
			expPANInHex: "1234123412341F", // F padding appended
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &Data{
				PAN:               tt.pan,
				ExpYear:           2025,
				ExpMonth:          12,
				ServiceCode:       "601",
				DiscretionaryData: "123123",
			}

			result, err := data.FormatISO()
			require.NoError(t, err)

			resultHex := strings.ToUpper(hex.EncodeToString(result))
			
			// The result should start with 3B (start sentinel), followed by the PAN portion
			expectedStart := "3B" + strings.ToUpper(tt.expPANInHex)
			
			require.True(t, strings.HasPrefix(resultHex, expectedStart))

			// Verify the PAN portion specifically
			// Skip the start sentinel (3B) and extract the PAN portion
			actualPANHex := resultHex[2 : 2+len(tt.expPANInHex)]
			require.Equal(t, strings.ToUpper(tt.expPANInHex), actualPANHex)
		})
	}
}

func TestData_FormatAlt_PANPadding(t *testing.T) {
	// Test that the alternative format also works correctly with PAN padding
	data := &Data{
		PAN:               "123412341234123", // 15 digits (odd)
		ExpYear:           2025,
		ExpMonth:          12,
		ServiceCode:       "601",
		DiscretionaryData: "123123",
	}

	result, err := data.FormatAlt()
	require.NoError(t, err)

	resultHex := strings.ToUpper(hex.EncodeToString(result))
	
	// Should start with 3B (start sentinel) + PAN with F padding
	expectedStart := "3B123412341234123F"
	
	require.True(t, strings.HasPrefix(resultHex, expectedStart))
}