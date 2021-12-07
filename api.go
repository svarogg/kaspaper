package kaspaper

type KaspaperAPI interface{
	GenerateWallet() KaspaperWallet
}

type MnemonicString [24]string
type KaspaperWallet interface {
	Mnemonic() MnemonicString
	KeysJSON() string
	Address(index int) string
}