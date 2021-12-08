package kaspaperlib

import (
	"strings"

	"github.com/kaspanet/kaspad/domain/dagconfig"

	"github.com/kaspanet/kaspad/cmd/kaspawallet/keys"

	"github.com/svarogg/kaspaper/model"
)

// Make sure we implement model.KaspaperWallet
var _ model.KaspaperWallet = &wallet{}

type wallet struct {
	dagParams *dagconfig.Params
	mnemonic  *model.MnemonicString
	keysFile  *keys.File
	keysJSON  string
	password  string
}

func newWallet(dagParams *dagconfig.Params, mnemonic string) (model.KaspaperWallet, error) {
	mnemonicString := &model.MnemonicString{}
	copy(mnemonicString[:], strings.Split(mnemonic, " ")) // We assume it splits to 24 words

	password := generatePassword()

	keysFile, err := keys.NewFileFromMnemonics(dagParams, mnemonicString.String(), password)
	if err != nil {
		return nil, err
	}
	keysJSON, err := keysFile.JSONString()
	if err != nil {
		return nil, err
	}

	return &wallet{
		dagParams: dagParams,
		mnemonic:  mnemonicString,
		keysFile:  keysFile,
		keysJSON:  keysJSON,
		password:  password,
	}, nil
}

func (w *wallet) Mnemonic() *model.MnemonicString {
	return w.mnemonic
}

func (w *wallet) KeysJSON() string {
	return w.keysJSON
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
