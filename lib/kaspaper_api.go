package lib

import (
	"github.com/kaspanet/kaspad/cmd/kaspawallet/keys"
	"github.com/kaspanet/kaspad/domain/dagconfig"
	"github.com/svarogg/kaspaper/model"
)

// Make sure we implement model.KaspaperWallet
var _ model.KaspaperAPI = &KaspaperAPI{}

type KaspaperAPI struct {
}

func (k KaspaperAPI) GenerateWallet() (model.KaspaperWallet, error) {
	encryptedPrivateKeys, extendedPublicKeys, err :=
		keys.CreateMnemonics(&dagconfig.MainnetParams, 1, "", false)
	if err != nil {
		return nil, err
	}

	return &KaspaperWallet{
		encryptedPrivateKeys: encryptedPrivateKeys,
		extendedPublicKeys:   extendedPublicKeys,
	}, nil
}
