package model

const NumThreads = 8

type KaspaperAPI interface {
	GenerateWallet() (KaspaperWallet, error)
}

type KaspaperWallet interface {
	Mnemonic() MnemonicString
	KeysJSON() string
	Address(index int) string
}
