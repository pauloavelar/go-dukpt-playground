package dukpt

import (
	"crypto/des"
	"encoding/hex"
	"fmt"
)

// normalizeKey ensures a 16-byte key is expanded to 24 bytes for 3DES (K1|K2|K1).
// The first 16 bytes are copied, and the next 8 bytes are filled with the first 8 bytes.
func normalizeKey(bdk []byte) ([]byte, error) {
	if len(bdk) != lenBDK {
		return nil, fmt.Errorf("invalid BDK length: expected %d bytes but got %d bytes", lenBDK, len(bdk))
	}

	key := make([]byte, len3DESKey)
	copy(key, bdk)
	copy(key[lenBDK:], bdk[:lenBDK/2])

	return key, nil
}

// deriveSessionKey derives the session key from the IPEK and KSN according to the DUKPT specification.
// It extracts the transaction counter from the KSN, applies the DUKPT key derivation algorithm, and returns
// the session key suitable for data encryption (with the data key variant applied).
func deriveSessionKey(ipek, ksn []byte) ([]byte, error) {
	if len(ipek) != lenBDK {
		return nil, fmt.Errorf("invalid IPEK length: expected %d bytes but got %d bytes", lenBDK, len(ipek))
	}
	if len(ksn) != lenKSN {
		return nil, fmt.Errorf("invalid KSN length: expected %d bytes but got %d bytes", lenKSN, len(ksn))
	}

	// Extract the 21-bit transaction counter from the KSN (rightmost 21 bits)
	counter := uint32(ksn[7]&0x1F)<<16 | uint32(ksn[8])<<8 | uint32(ksn[9])

	// Prepare the initial key (start with IPEK)
	key := make([]byte, len(ipek))
	copy(key, ipek)

	// Prepare the base KSN (KSN with transaction counter bits zeroed)
	ksnBase := make([]byte, lenKSN)
	copy(ksnBase, ksn)
	ksnBase[7] &= 0xE0
	ksnBase[8] = 0x00
	ksnBase[9] = 0x00

	// For each bit set in the transaction counter, apply the DUKPT key derivation step
	for i := 0; i < 21; i++ {
		if (counter & (1 << i)) != 0 {
			// Set the corresponding bit in the KSN
			ksnReg := make([]byte, len(ksnBase))
			copy(ksnReg, ksnBase)
			bitMask := uint32(1 << i)
			ksnReg[7] |= byte((bitMask >> 16) & 0x1F)
			ksnReg[8] |= byte((bitMask >> 8) & 0xFF)
			ksnReg[9] |= byte(bitMask & 0xFF)

			// Apply the DUKPT non-reversible key generation process
			var err error
			key, err = dukptNRKGP(key, ksnReg[2:10])
			if err != nil {
				return nil, err
			}
		}
	}

	// Apply the data key variant (XOR with 0x0000000000FF00000000000000FF0000)
	variant := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF}
	for i := 0; i < len(key); i++ {
		key[i] ^= variant[i]
	}

	return key, nil
}

// dukptNRKGP implements the DUKPT Non-Reversible Key Generation Process (NRKGP).
// It takes a 16-byte key and an 8-byte KSN register, and returns the derived key.
func dukptNRKGP(key, ksnReg []byte) ([]byte, error) {
	if len(key) != 16 || len(ksnReg) != 8 {
		return nil, fmt.Errorf("invalid input lengths for NRKGP")
	}

	// Split key into left and right halves
	keyL := make([]byte, 8)
	keyR := make([]byte, 8)
	copy(keyL, key[:8])
	copy(keyR, key[8:])

	// Step 1: Encrypt KSN register with keyL (with odd parity)
	keyLMasked := make([]byte, 8)
	for i := 0; i < 8; i++ {
		keyLMasked[i] = keyL[i] ^ 0xC0
	}
	blockL, err := des.NewCipher(keyLMasked)
	if err != nil {
		return nil, err
	}
	encL := make([]byte, 8)
	blockL.Encrypt(encL, ksnReg)

	// Step 2: Encrypt KSN register with keyR (with odd parity)
	keyRMasked := make([]byte, 8)
	for i := 0; i < 8; i++ {
		keyRMasked[i] = keyR[i] ^ 0xC0
	}
	blockR, err := des.NewCipher(keyRMasked)
	if err != nil {
		return nil, err
	}
	encR := make([]byte, 8)
	blockR.Encrypt(encR, ksnReg)

	// Step 3: Concatenate results
	derived := append(encL, encR...)
	return derived, nil
}

// EncryptTrack2 encrypts formatted track2 data using DUKPT and returns the encrypted data.
//
// Parameters:
//   - formatted: The formatted track2 data (numeric only)
//   - bdk: The Base Derivation Key (16 bytes)
//   - ksn: The Key Serial Number (10 bytes)
//
// Returns:
//   - encrypted data ([]byte)
//   - error if encryption fails
func EncryptTrack2(formatted string, bdk, ksn []byte) ([]byte, error) {
	// Step 1: Derive IPEK
	ipek, err := deriveIPEK(bdk, ksn)
	if err != nil {
		return nil, err
	}
	// Step 2: Derive session key (not yet implemented)
	sessionKey, err := deriveSessionKey(ipek, ksn)
	if err != nil {
		return nil, err
	}
	// Step 3: Encrypt track2 data (3DES ECB, pad to 16 bytes)
	data, err := hex.DecodeString(formatted)
	if err != nil {
		return nil, err
	}
	if len(data)%8 != 0 {
		pad := 8 - (len(data) % 8)
		for i := 0; i < pad; i++ {
			data = append(data, 0x00)
		}
	}
	block, err := des.NewTripleDESCipher(sessionKey)
	if err != nil {
		return nil, err
	}
	encrypted := make([]byte, len(data))
	for bs, be := 0, 8; bs < len(data); bs, be = bs+8, be+8 {
		block.Encrypt(encrypted[bs:be], data[bs:be])
	}
	return encrypted, nil
}
