package kaspaperlib

import (
	"github.com/kaspanet/kaspad/cmd/kaspawallet/libkaspawallet"
	"github.com/svarogg/kaspaper/model"
)

// Make sure we implement model.KaspaperWallet
var _ model.KaspaperAPI = &api{}

type api struct {
}

func NewAPI() model.KaspaperAPI {
	return &api{}
}

func (a *api) GenerateWallet() (model.KaspaperWallet, error) {
	mnemonics, err := libkaspawallet.CreateMnemonic()
	if err != nil {
		return nil, err
	}
	// It's safe to use [0] because we know there's exactly 1 key, since we passed numKeys: 1 to CreateMnemonics
	return newWallet(mnemonics)
}
