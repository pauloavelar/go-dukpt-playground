package dukpt

const (
	// lenBDK is the length of the Base Derivation Key (BDK) in bytes.
	lenBDK = 16

	// lenIPEK is the length of the Initial PIN Encryption Key (IPEK) in bytes.
	lenIPEK = 16

	// lenKSN is the length of the Key Serial Number (KSN) in bytes.
	lenKSN = 10

	// lenMaskedKSN is the length of the masked KSN used for IPEK derivation.
	lenMaskedKSN = 8 // Masked KSN length for IPEK derivation

	// len3DESKey is the length of a 3DES key in bytes as required by [des.NewTripleDESCipher].
	len3DESKey = 24
)
