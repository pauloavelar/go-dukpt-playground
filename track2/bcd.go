package track2

// charToNibble converts a character to a nibble value.
// Handles numeric characters ('0'-'9'), the '=' separator and 'F' padding character.
func charToNibble(c byte) byte {
	switch c {
	case '=':
		return 0xD
	case 'F':
		return 0xF
	default:
		return c - '0'
	}
}
