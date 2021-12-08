package kaspaperlib

import (
	"strings"

	"github.com/kaspanet/kaspad/cmd/kaspawallet/keys"
	"github.com/svarogg/kaspaper/model"
)

func decryptMnemonic(encryptedMnemonic *keys.EncryptedMnemonic) (*model.MnemonicString, error) {
	decryptedMnemonic, err := keys.DecryptMnemonic(keys.DefaultNumThreads, encryptedMnemonic, []byte(""))
	if err != nil {
		return nil, err
	}

	mnemonicString := &model.MnemonicString{}
	copy(mnemonicString[:], strings.Split(decryptedMnemonic, " ")) // We assume it splits to 24 words

	return mnemonicString, nil
}
