package lib

import "github.com/svarogg/kaspaper/model"

type KaspaperWallet struct{
	
}

func (k KaspaperWallet) Mnemonic() model.MnemonicString {
	panic("implement me")
}

func (k KaspaperWallet) KeysJSON() string {
	panic("implement me")
}

func (k KaspaperWallet) Address(index int) string {
	panic("implement me")
}

