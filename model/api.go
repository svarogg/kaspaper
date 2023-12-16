package model

type KaspaperAPI interface {
	GenerateWallet() (KaspaperWallet, error)
}

type KaspaperWallet interface {
	Mnemonic() *MnemonicString
	Address(index int) (string, error)
	AddressQR(index int) ([]byte, error)
	KPubKey()(string)
	KPubKeyQR() ([]byte, error)
}
