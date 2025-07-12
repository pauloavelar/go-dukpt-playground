// Package dukpt implements the DUKPT (Derived Unique Key Per Transaction) key management scheme.
//
// The code in this package is written with clarity in mind, so many operations are explicit and
// not optimized for performance. The private functions are well-documented for the same reason.
package dukpt

const (
	// lenDesData is the length of data blocks used in DES encryption.
	lenDesData = 8

	// lenKey is the length of any key used in DUKPT, such as BDK, IPEK or session keys.
	lenKey = 16

	// lenKeyHalf is half of the key length, used for IPEK and session key derivation.
	lenKeyHalf = lenKey / 2

	// lenKSN is the length of the Key Serial Number (KSN) in bytes.
	lenKSN = 10

	// lenMaskedKSN is the length of the masked KSN used for IPEK derivation.
	// It corresponds to half of a key because it is used twice to derive a key.
	lenMaskedKSN = lenKeyHalf

	// len3DESKey is the length of a 3DES key in bytes as required by [des.NewTripleDESCipher].
	len3DESKey = 24

	// ksnTransactionCounterBits is the number of bits in the KSN transaction counter.
	ksnTransactionCounterBits = 21
)

type (
	// KSN is a 10-byte Key Serial Number used in DUKPT transactions.
	KSN [lenKSN]byte

	// Key is a generic 16-byte key used in DUKPT.
	// It can represent BDK, IPEK or session keys and its masks.
	Key [lenKey]byte

	// KeyHalf is an 8-byte half of a Key, used in key derivation.
	KeyHalf [lenKey / 2]byte

	// TripleDESKey is a 24-byte key used for 3DES encryption in DUKPT.
	// It is derived from the BDK or IPEK through normalization.
	TripleDESKey [len3DESKey]byte
)
