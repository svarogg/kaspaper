package kaspaperlib

import (
	"strings"

	"github.com/svarogg/kaspaper/model"
)

// Make sure we implement model.KaspaperWallet
var _ model.KaspaperWallet = &wallet{}

type wallet struct {
	mnemonic *model.MnemonicString
}

func newWallet(mnemonic string) (model.KaspaperWallet, error) {
	mnemonicString := &model.MnemonicString{}
	copy(mnemonicString[:], strings.Split(mnemonic, " ")) // We assume it splits to 24 words

	return &wallet{
		mnemonic: mnemonicString,
	}, nil
}

func (w *wallet) Mnemonic() *model.MnemonicString {
	return w.mnemonic
}

func (w *wallet) KeysJSON() string {
	panic("implement me")
}

func (w *wallet) Address(index int) string {
	panic("implement me")
}

func (w *wallet) QR() []byte {
	panic("implement me")
}

func (w *wallet) AddressQR() []byte {
	panic("implement me")
}
