package model

type KaspaperAPI interface{
	GenerateWallet() KaspaperWallet
}

type KaspaperWallet interface {
	Mnemonic() MnemonicString
	KeysJSON() string
	Address(index int) string
}