package bcd

import (
	"fmt"
	"regexp"

	"golang.org/x/exp/constraints"
)

var (
	regexpNumeric = regexp.MustCompile("^[0-9]+$")
)

// ByteLen calculates the number of BCD bytes needed to represent a numeric string.
// BCD encoding requires 2 characters per byte (1 is added for odd lengths).
func ByteLen(str string) int {
	return (len(str) + 1) / 2
}

// StringToBCD converts a numeric string to BCD bytes.
func StringToBCD(str string) ([]byte, error) {
	if !regexpNumeric.MatchString(str) {
		return nil, fmt.Errorf("invalid numeric string: %str", str)
	}

	if len(str)%2 != 0 {
		str = "0" + str
	}

	bcd := make([]byte, len(str)/2)
	for i := 0; i < len(str); i += 2 {
		h := str[i] - '0'
		l := str[i+1] - '0'

		bcd[i/2] = (h << 4) | l
	}

	return bcd, nil
}

// PANToBCD converts a PAN (Primary Account Number) string to BCD bytes.
// According to ISO/IEC 7813, if the PAN has an odd number of digits,
// a padding digit 'F' is appended to the end before BCD encoding.
func PANToBCD(pan string) ([]byte, error) {
	if !regexpNumeric.MatchString(pan) {
		return nil, fmt.Errorf("invalid PAN: %s", pan)
	}

	// For odd-length PANs, append 'F' padding per ISO/IEC 7813
	if len(pan)%2 != 0 {
		pan = pan + "F"
	}

	bcd := make([]byte, len(pan)/2)
	for i := 0; i < len(pan); i += 2 {
		var h, l byte
		
		// Handle first digit
		if pan[i] == 'F' {
			h = 0xF
		} else {
			h = pan[i] - '0'
		}
		
		// Handle second digit
		if pan[i+1] == 'F' {
			l = 0xF
		} else {
			l = pan[i+1] - '0'
		}

		bcd[i/2] = (h << 4) | l
	}

	return bcd, nil
}

// IntegerToBCD converts any int-based number to BCD.
func IntegerToBCD[T constraints.Integer](n T) []byte {
	str := fmt.Sprintf("%02d", n)
	bcd, _ := StringToBCD(str)
	return bcd
}
