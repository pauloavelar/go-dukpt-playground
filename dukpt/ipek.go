package dukpt

import (
	"crypto/des"
	"fmt"
)

var (
	// bdkMaskForIPEK as defined in ANSI X9.24-1:2017 (DUKPT) section 7.2.2.2.
	bdkMaskForIPEK = []byte{0xC0, 0xC0, 0xC0, 0xC0, 0xC0, 0xC0, 0xC0, 0xC0, 0xC0, 0xC0, 0xC0, 0xC0, 0x00, 0x00, 0x00, 0x00}

	// ksnMaskForIPEK as defined in ANSI X9.24-1:2017 (DUKPT) section 7.2.2.2.
	ksnMaskForIPEK = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0b11100000}
)

// deriveIPEK derives the Initial PIN Encryption Key (IPEK) from the BDK and KSN.
// The IPEK is a 16-byte key used for encrypting PINs in DUKPT transactions.
//
// There are two prerequisites:
// - The KSN must be masked according to the DUKPT spec: take first 8 bytes, clear the last 5 bits.
// - The BDK must be masked according to the DUKPT spec: XOR with the mask 0xC0C0C0C000000000C0C0C0C000000000.
//
// Then, the IPEK is derived as follows:
// - The left half is derived by encrypting the masked KSN with the *original* BDK.
// - The right half is derived by encrypting the masked KSN with a *masked* BDK.
func deriveIPEK(bdk, ksn []byte) ([]byte, error) {
	ksnMasked, err := maskKSNForIPEK(ksn)
	if err != nil {
		return nil, err
	}

	// data for the right half of the IPEK is derived from the BDK masked with the BDK mask
	bdkMasked, err := maskBDKForIPEK(bdk)
	if err != nil {
		return nil, err
	}

	left, err := encryptIPEKHalfWith3DES(bdk, ksnMasked)
	if err != nil {
		return nil, err
	}

	right, err := encryptIPEKHalfWith3DES(bdkMasked, ksnMasked)
	if err != nil {
		return nil, err
	}

	ipek := make([]byte, lenIPEK)
	copy(ipek[:lenIPEK/2], left)
	copy(ipek[lenIPEK/2:], right)

	return ipek, nil
}

func encryptIPEKHalfWith3DES(key []byte, data []byte) ([]byte, error) {
	if len(data) != 8 {
		return nil, fmt.Errorf("each IPEK half must be exactly 8 bytes, got %d bytes", len(data))
	}

	normalizedKey, err := normalizeKey(key)
	if err != nil {
		return nil, err
	}

	block, err := des.NewTripleDESCipher(normalizedKey)
	if err != nil {
		return nil, err
	}

	encrypted := make([]byte, 8)
	block.Encrypt(encrypted, data)

	return encrypted, nil
}

// maskBDKForIPEK applies the mask defined in ANSI X9.24-1:2017 (DUKPT) section 7.2.2.2 to the BDK for IPEK derivation.
// The mask 0xC0C0C0C000000000C0C0C0C000000000 is used to generate the right half of the IPEK by XORing with the BDK.
func maskBDKForIPEK(bdk []byte) ([]byte, error) {
	if len(bdk) != lenBDK {
		return nil, fmt.Errorf("invalid BDK length: expected %d bytes but got %d bytes", lenBDK, len(bdk))
	}

	masked := make([]byte, lenBDK)
	for i := 0; i < lenBDK; i++ {
		masked[i] = bdk[i] ^ bdkMaskForIPEK[i]
	}

	return masked, nil
}

// maskKSNForIPEK implements KSN masking as defined in ANSI X9.24-1:2017 (DUKPT) for IPEK derivation.
// The first 8 bytes of the KSN are used, and the last byte is masked to clear the last 5 bits.
func maskKSNForIPEK(ksn []byte) ([]byte, error) {
	if len(ksn) != lenKSN {
		return nil, fmt.Errorf("invalid KSN length: expected %d bytes but got %d bytes", lenKSN, len(ksn))
	}

	masked := make([]byte, lenMaskedKSN)
	for i := 0; i < lenMaskedKSN; i++ {
		masked[i] = ksn[i] & ksnMaskForIPEK[i]
	}

	return masked, nil
}
