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

// NumericToRightPaddedBCD converts a numeric string to BCD bytes with right F-padding.
// According to ISO/IEC 7813, if the input has an odd number of digits,
// a padding digit 'F' is appended to the end before BCD encoding.
func NumericToRightPaddedBCD(input string) ([]byte, error) {
	// Handle empty input as a special case
	if len(input) == 0 {
		return []byte{}, nil
	}

	if !regexpNumeric.MatchString(input) {
		return nil, fmt.Errorf("invalid numeric string: %s", input)
	}

	// For odd-length input, append 'F' padding per ISO/IEC 7813
	if len(input)%2 != 0 {
		input = input + "F"
	}

	bcd := make([]byte, len(input)/2)
	for i := 0; i < len(input); i += 2 {
		h := charToNibble(input[i])
		l := charToNibble(input[i+1])
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

// charToNibble converts a character to a nibble value.
// Handles both numeric characters ('0'-'9') and 'F' padding character.
func charToNibble(c byte) byte {
	if c == 'F' {
		return 0xF
	} else {
		return c - '0'
	}
}
