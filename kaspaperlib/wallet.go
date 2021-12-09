package kaspaperlib

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/kaspanet/kaspad/cmd/kaspawallet/keys"
	"github.com/kaspanet/kaspad/cmd/kaspawallet/libkaspawallet"
	"github.com/kaspanet/kaspad/domain/dagconfig"
	"github.com/skip2/go-qrcode"
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

func (w *wallet) Address(index int) (string, error) {
	path := fmt.Sprintf("m/%d/%d", libkaspawallet.ExternalKeychain, index)
	address, err := libkaspawallet.Address(w.dagParams, w.keysFile.ExtendedPublicKeys, 1, path, false)
	if err != nil {
		return "", err
	}
	return address.String(), nil
}

func (w *wallet) QR() ([]byte, error) {
	contents := fmt.Sprintf("kaspamnemonic:%s", strings.Join(w.mnemonic[:], "."))

	qr, err := qrcode.Encode(contents, qrcode.High, 256)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return qr, nil
}

func (w *wallet) AddressQR(index int) ([]byte, error) {
	address, err := w.Address(index)
	if err != nil {
		return nil, err
	}

	qr, err := qrcode.Encode(address, qrcode.High, 256)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return qr, nil
}
