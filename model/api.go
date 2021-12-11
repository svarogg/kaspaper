package model

type KaspaperAPI interface {
	GenerateWallet() (KaspaperWallet, error)
}

type KaspaperWallet interface {
	Mnemonic() *MnemonicString
	KeysJSON() string
	Address(index int) (string, error)
	AddressQR(index int) ([]byte, error)
}
