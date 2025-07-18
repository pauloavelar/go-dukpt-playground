package track2

import (
	"fmt"
	"strings"
)

const (
	padding   = 'F'
	separator = '='
)

// Data represents the fields required to generate EMV Track 2 Equivalent Data (tag 57).
//
// Reference: EMV Book 3, Annex B1; ISO/IEC 7812-1:2017(E)
//
// Fields:
//   - PAN: Primary Account Number (10 to 19 digits)
//   - ExpYear: Expiry year (2 digits, YY)
//   - ExpMonth: Expiry month (2 digits, MM)
//   - ServiceCode: 3-digit service code
//   - DiscretionaryData: Up to 13 digits
//
// All fields must be valid and numeric as per EMV/Track 2 specifications.
type Data struct {
	PAN               string `minLength:"10" maxLength:"19"` // Primary Account Number (10 to 19 digits, numeric only)
	ExpYear           string `exactLength:"2"`               // Expiry year (YY)
	ExpMonth          string `minLength:"1" maxLength:"2"`   // Expiry month (MM)
	ServiceCode       string `exactLength:"3"`               // 3-digit, numeric only
	DiscretionaryData string `maxLength:"13"`                // Up to 13 digits, numeric only
}

// Format returns the human-readable Track 2 Equivalent Data for EMV (tag 57)
func (d *Data) Format() (string, error) {
	err := d.Validate()
	if err != nil {
		return "", fmt.Errorf("invalid data - %w", err)
	}

	var b strings.Builder
	b.WriteString(d.PAN)
	b.WriteRune(separator)
	b.WriteString(d.ExpYear)
	b.WriteString(fmt.Sprintf("%02s", d.ExpMonth))
	b.WriteString(d.ServiceCode)
	b.WriteString(d.DiscretionaryData)

	if b.Len()%2 != 0 {
		b.WriteRune(padding)
	}

	return b.String(), nil
}

// Encode returns a valid TLV-encoded Track 2 Equivalent Data.
func (d *Data) Encode() ([]byte, error) {
	formatted, err := d.Format()
	if err != nil {
		return nil, err
	}

	encoded := make([]byte, len(formatted)/2)
	for i := 0; i < len(formatted); i += 2 {
		h := charToNibble(formatted[i])
		l := charToNibble(formatted[i+1])

		encoded[i/2] = h<<4 | l
	}

	return encoded, nil
}

// Validate checks if all fields are numeric and have valid lengths.
// Returns an error describing the first problem found, if any.
func (d *Data) Validate() error {
	return validateNumericFields(d)
}
