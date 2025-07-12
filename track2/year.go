package track2

import (
	"fmt"
)

func sanitizeExpYear(year uint16) (uint8, error) {
	if year >= 1 && year <= 12 {
		return uint8(year), nil
	}

	if year >= 1900 && year < 3000 {
		return uint8(year % 100), nil
	}

	return 0, fmt.Errorf("year out of range: %d", year)
}
