package dukpt

import (
	"crypto/des"
)

var (
	// bdkMaskForIPEK as defined in ANSI X9.24-1:2017 (DUKPT) section 7.2.2.2.
	bdkMaskForIPEK = []byte{0xC0, 0xC0, 0xC0, 0xC0, 0xC0, 0xC0, 0xC0, 0xC0, 0xC0, 0xC0, 0xC0, 0xC0, 0x00, 0x00, 0x00, 0x00}
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
func deriveIPEK(bdk Key, ksn KSN) (Key, error) {
	ksnMasked := maskKSNForIPEK(ksn)

	left, err := encryptIPEKHalfWith3DES(bdk, ksnMasked)
	if err != nil {
		return Key{}, err
	}

	bdkMasked := maskBDKForIPEK(bdk)
	right, err := encryptIPEKHalfWith3DES(bdkMasked, ksnMasked)
	if err != nil {
		return Key{}, err
	}

	var ipek Key
	copy(ipek[:lenKeyHalf], left[:])
	copy(ipek[lenKeyHalf:], right[:])

	return ipek, nil
}

// encryptIPEKHalfWith3DES encrypts a half of the IPEK using 3DES with the provided key.
// The key is a BDK or a masked BDK, and the data is the masked KSN.
func encryptIPEKHalfWith3DES(key Key, data KeyHalf) (KeyHalf, error) {
	normalizedKey := normalizeKey(key)

	block, err := des.NewTripleDESCipher(normalizedKey[:])
	if err != nil {
		return KeyHalf{}, err
	}

	var encrypted KeyHalf
	block.Encrypt(encrypted[:], data[:])

	return encrypted, nil
}

// normalizeKey ensures a 16-byte key is expanded to 24 bytes for 3DES (K1|K2|K1).
// The first 16 bytes are copied, and the next 8 bytes are filled with the first 8 bytes.
func normalizeKey(key Key) TripleDESKey {
	var normalized TripleDESKey

	copy(normalized[:], key[:])
	copy(normalized[lenKey:], key[:lenKeyHalf])

	return normalized
}

// maskBDKForIPEK applies the mask defined in ANSI X9.24-1:2017 (DUKPT) section 7.2.2.2 to the BDK for IPEK derivation.
// The mask 0xC0C0C0C000000000C0C0C0C000000000 is used to generate the right half of the IPEK by XORing with the BDK.
func maskBDKForIPEK(bdk Key) Key {
	var masked Key

	for i := range lenKey {
		masked[i] = bdk[i] ^ bdkMaskForIPEK[i]
	}

	return masked
}

// maskKSNForIPEK implements KSN masking as defined in ANSI X9.24-1:2017 (DUKPT) for IPEK derivation.
// The first 8 bytes of the KSN are used, and the last byte is masked to clear the last 5 bits,
// leaving 7 * 8 + 3 bits of data (59 bits).
//
// This is the same as truncating the KSI to the first 8 bytes.
func maskKSNForIPEK(ksn KSN) KeyHalf {
	ksi := extractKSI(ksn)
	return KeyHalf(ksi[:lenMaskedKSN])
}
