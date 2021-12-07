package lib

import (
	"github.com/kaspanet/kaspad/cmd/kaspawallet/keys"
	"github.com/svarogg/kaspaper/model"
)

// Make sure we implement model.KaspaperWallet
var _ model.KaspaperWallet = &KaspaperWallet{}

type KaspaperWallet struct {
	encryptedPrivateKeys []*keys.EncryptedMnemonic
	extendedPublicKeys   []string
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
