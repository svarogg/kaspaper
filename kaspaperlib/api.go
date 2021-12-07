package kaspaperlib

import (
	"github.com/kaspanet/kaspad/cmd/kaspawallet/keys"
	"github.com/kaspanet/kaspad/domain/dagconfig"
	"github.com/svarogg/kaspaper/model"
)

// Make sure we implement model.KaspaperWallet
var _ model.KaspaperAPI = &api{}

type api struct {
}

func (a *api) GenerateWallet() (model.KaspaperWallet, error) {
	encryptedMnemonics, extendedPublicKeys, err :=
		keys.CreateMnemonics(&dagconfig.MainnetParams, 1, "", false)
	if err != nil {
		return nil, err
	}
	// It's safe to use [0] because we know there's exactly 1 key, since we passed numKeys: 1 to CreateMnemonics
	return newWallet(encryptedMnemonics[0], extendedPublicKeys[0]), nil
}
