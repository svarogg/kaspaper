package kaspaperlib

import (
	"github.com/kaspanet/kaspad/cmd/kaspawallet/keys"
	"github.com/svarogg/kaspaper/model"
)

func decryptMnemonic(encryptedMnemonic *keys.EncryptedMnemonic) (*model.MnemonicString, error) {
	decryptedMnemonic, err := keys.DecryptMnemonic(keys.DefaultNumThreads, encryptedMnemonic, []byte(""))
	if err != nil {
		return nil, err
	}
	// Panice so that I can see how the hell this looks
	panic(decryptedMnemonic)
}
