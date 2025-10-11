package crypto

var (
	// AddressHRP is the Human Readable Part (HRP) for address.
	AddressHRP = "ae"
	// PublicKeyHRP is the Human Readable Part (HRP) for public key.
	PublicKeyHRP = "public"
	// PrivateKeyHRP is the Human Readable Part (HRP) for private key.
	PrivateKeyHRP = "secret"
	// XPublicKeyHRP is the Human Readable Part (HRP) for extended public key.
	XPublicKeyHRP = "xpublic"
	// XPrivateKeyHRP is the Human Readable Part (HRP) for extended private key.
	XPrivateKeyHRP = "xsecret"
)

// ToTestnetHRP makes HRPs testnet specified.
func ToTestnetHRP() {
	AddressHRP = "tae"
	PublicKeyHRP = "tpublic"
	PrivateKeyHRP = "tsecret"
	XPublicKeyHRP = "txpublic"
	XPrivateKeyHRP = "txsecret"
}

// ToMainnetHRP makes HRPs mainnet specified.
func ToMainnetHRP() {
	AddressHRP = "ae"
	PublicKeyHRP = "public"
	PrivateKeyHRP = "secret"
	XPublicKeyHRP = "xpublic"
	XPrivateKeyHRP = "xsecret"
}
