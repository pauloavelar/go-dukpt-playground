package dukpt

import (
	"crypto/des"
	"encoding/hex"
	"fmt"
)

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
	if len(bdk) != lenKey {
		return nil, fmt.Errorf("BDK must be %d bytes long", lenKey)
	}

	if len(ksn) != lenKSN {
		return nil, fmt.Errorf("KSN must be %d bytes long", lenKSN)
	}

	ipek, err := deriveIPEK(Key(bdk), KSN(ksn))
	if err != nil {
		return nil, err
	}

	sessionKey, err := deriveSessionKey(ipek, KSN(ksn))
	if err != nil {
		return nil, err
	}

	sessionKey = maskSessionKeyForData(sessionKey)

	// encrypt track2 data (3DES ECB, pad to 16 bytes)
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

	block, err := des.NewTripleDESCipher(sessionKey[:])
	if err != nil {
		return nil, err
	}

	encrypted := make([]byte, len(data))
	for bs, be := 0, 8; bs < len(data); bs, be = bs+8, be+8 {
		block.Encrypt(encrypted[bs:be], data[bs:be])
	}

	return encrypted, nil
}

// deriveSessionKey derives the session key from the IPEK and KSN according to the DUKPT specification.
// It extracts the transaction counter from the KSN, applies the DUKPT key derivation algorithm, and returns
// the session key suitable for data encryption (with the data key variant applied).
func deriveSessionKey(ipek Key, ksn KSN) (Key, error) {
	ksi := extractKSI(ksn)
	ctr := extractKSNTransactionCounter(ksn)

	var sessionKey Key
	copy(sessionKey[:], ipek[:])

	// for each bit set in the transaction counter,
	// apply the DUKPT non-reversible key generation process
	for i := range ksnTransactionCounterBits {
		if isBitSet(ctr, i) {
			ksnReg := buildKsnRegister(ksi, i)

			var err error
			sessionKey, err = dukptNRKGP(sessionKey, KeyHalf(ksnReg[lenKSN-lenKeyHalf:]))
			if err != nil {
				return Key{}, err
			}
		}
	}

	return sessionKey, nil
}

// dukptNRKGP implements the DUKPT Non-Reversible Key Generation Process (NRKGP).
// It takes a 16-byte key and the last 8 bytes of a KSN register, and returns the derived key.
func dukptNRKGP(key Key, ksnReg KeyHalf) (Key, error) {
	keyL, keyR := splitKey(key)

	encL, err := encryptWithDES(keyL, ksnReg)
	if err != nil {
		return Key{}, err
	}

	encR, err := encryptWithDES(keyR, ksnReg)
	if err != nil {
		return Key{}, err
	}

	var derived Key
	copy(derived[:lenKeyHalf], encL[:])
	copy(derived[lenKeyHalf:], encR[:])

	return derived, nil
}

// encryptWithDES encrypts the provided data using DES with the given key.
// It applies a mask to the key before encryption, as specified in the DUKPT standard.
func encryptWithDES(desKey KeyHalf, data KeyHalf) (KeyHalf, error) {
	var maskedKey KeyHalf
	for i := range lenKeyHalf {
		maskedKey[i] = desKey[i] ^ 0xC0
	}

	blockL, err := des.NewCipher(maskedKey[:])
	if err != nil {
		return KeyHalf{}, err
	}

	var enc KeyHalf
	blockL.Encrypt(enc[:], data[:])

	return enc, nil
}

// isBitSet returns true if the bit at bitPosition in the value is set (1).
// The bitPosition is zero-based, where 0 is the least significant bit.
func isBitSet(value uint32, bitPosition int) bool {
	return (value & (1 << bitPosition)) != 0
}

// buildKsnRegister builds a KSN register from the KSI and a 'bit' position.
// The register represents the KSI with a single bit set in the transaction counter, and the bit
// variable represents the position, which is zero-based and starts from the least significant bit.
// This is used in the DUKPT session key derivation, in an iterative process for each bit set to 1.
func buildKsnRegister(ksi KSN, bit int) KSN {
	var ksnReg KSN
	copy(ksnReg[:], ksi[:])

	// apply the bit mask to each byte of the counter
	bitMask := uint32((1 << bit) & maskKSNCounter)
	ksnReg[lenKSN-3] |= byte(bitMask >> 16)
	ksnReg[lenKSN-2] |= byte(bitMask >> 8)
	ksnReg[lenKSN-1] |= byte(bitMask >> 0)

	return ksnReg
}

// splitKey splits a 16-byte key into two 8-byte halves.
func splitKey(key Key) (keyL, keyR KeyHalf) {
	copy(keyL[:], key[:lenKeyHalf])
	copy(keyR[:], key[lenKeyHalf:])

	return keyL, keyR
}
