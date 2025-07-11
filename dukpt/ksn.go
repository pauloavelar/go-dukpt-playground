package dukpt

const (
	// maskLastKSIByte is a mask for the last byte of the KSI, which clears the rightmost 5 bits.
	maskLastKSIByte = 0xE0 // = 0b1110_0000

	// maskKSNCounter is a mask for extracting the last 21 bits (KSN transaction counter)
	maskKSNCounter = 1<<ksnTransactionCounterBits - 1 // = 0b1_1111_1111_1111_1111_1111
)

// extractKSI returns the Key Set Identifier (KSI) as a 10-byte array,
// which is the KSN with the rightmost 21 bits (transaction counter) zeroed.
func extractKSI(ksn KSN) KSN {
	var ksi KSN
	copy(ksi[:], ksn[:])

	ksi[lenKSN-3] &= maskLastKSIByte
	ksi[lenKSN-2] = 0x00
	ksi[lenKSN-1] = 0x00

	return ksi
}

// extractKSNTransactionCounter extracts the transaction counter from the KSN.
// This corresponds to the last 21 bits of the KSN, which is incremented in every transaction.
func extractKSNTransactionCounter(ksn KSN) uint32 {
	last3Bytes := ksn[lenKSN-3:]

	var counter uint32
	for _, b := range last3Bytes {
		counter = (counter << 8) | uint32(b)
	}

	return counter & maskKSNCounter
}
