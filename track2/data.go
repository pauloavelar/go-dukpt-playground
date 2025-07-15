package track2

import (
	"fmt"

	"github.com/pauloavelar/go-dukpt-playground/bcd"
)

const (
	startSentinel = ';'
	separator     = '='
	separatorAlt  = 'D'
	endSentinel   = '?'

	baseLen = 5 // start sentinel + YY + MM + separator + end sentinel
)

type Data struct {
	PAN               string // Primary Account Number
	ExpYear           uint16 // Expiry year (supports YY and YYYY formats)
	ExpMonth          uint8  // Expiry month (MM)
	ServiceCode       string
	DiscretionaryData string
}

// FormatISO formats Data using '=' as the separator (ISO/IEC 7813)
func (d *Data) FormatISO() ([]byte, error) {
	return d.formatTrack2(separator)
}

// FormatAlt formats Data using 'D' as the separator (alternative)
func (d *Data) FormatAlt() ([]byte, error) {
	return d.formatTrack2(separatorAlt)
}

func (d *Data) formatTrack2(sep rune) ([]byte, error) {
	// Validate PAN is not empty
	if len(d.PAN) == 0 {
		return nil, fmt.Errorf("invalid PAN: %s", d.PAN)
	}

	pan, err := bcd.NumericToRightPaddedBCD(d.PAN)
	if err != nil {
		return nil, err
	}

	etc, err := bcd.NumericToRightPaddedBCD(d.ServiceCode + d.DiscretionaryData)
	if err != nil {
		return nil, err
	}

	expYear, err := sanitizeExpYear(d.ExpYear)
	if err != nil {
		return nil, err
	}

	t2 := make([]byte, 0, d.byteLen())
	t2 = append(t2, startSentinel)
	t2 = append(t2, pan...)
	t2 = append(t2, byte(sep))
	t2 = append(t2, bcd.IntegerToBCD(expYear)...)
	t2 = append(t2, bcd.IntegerToBCD(d.ExpMonth)...)
	t2 = append(t2, etc...)
	t2 = append(t2, endSentinel)

	return t2, nil
}

func (d *Data) byteLen() int {
	return baseLen + bcd.ByteLen(d.PAN) + bcd.ByteLen(d.ServiceCode+d.DiscretionaryData)
}
