package kaspaperlib

import (
	"github.com/kaspanet/kaspad/cmd/kaspawallet/keys"
	"github.com/svarogg/kaspaper/model"
)

// Make sure we implement model.KaspaperWallet
var _ model.KaspaperWallet = &wallet{}

type wallet struct {
	mnemonic model.MnemonicString
}

func newWallet(encryptedMnemonic *keys.EncryptedMnemonic, extendedPublicKey string) *wallet {
	return &wallet{
		mnemonic: decryptMnemonic(encryptedMnemonic),
	}
}

func (w *wallet) Mnemonic() model.MnemonicString {
	return w.mnemonic
}

func (w *wallet) KeysJSON() string {
	panic("implement me")
}

func (w *wallet) Address(index int) string {
	panic("implement me")
}
