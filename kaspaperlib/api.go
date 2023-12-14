package kaspaperlib

import (
	"github.com/kaspanet/kaspad/cmd/kaspawallet/libkaspawallet"
	"github.com/kaspanet/kaspad/domain/dagconfig"
	"github.com/svarogg/kaspaper/model"
	"github.com/kaspanet/kaspad/cmd/kaspawallet/utils"
	"github.com/pkg/errors"
	"github.com/tyler-smith/go-bip39"
	"fmt"
	"os"
	"bufio"
)

// Make sure we implement model.KaspaperWallet
var _ model.KaspaperAPI = &api{}

type api struct {
	dagParams *dagconfig.Params
}

func NewAPI(dagParams *dagconfig.Params) model.KaspaperAPI {
	return &api{dagParams: dagParams}
}

func (a *api) GenerateWallet() (model.KaspaperWallet, error) {
	
	fmt.Printf("Press blank or mnemonic #%d here:\n", 1)
	reader := bufio.NewReader(os.Stdin)
	var mnemonics string
	var err error 
	
	mnemonics, err = utils.ReadLine(reader)

	if err != nil {
		return nil, err
	}
	if mnemonics == "" {

		fmt.Printf("Creating new mnemonics\n")
		mnemonics, err = libkaspawallet.CreateMnemonic()
	} else {

		if !bip39.IsMnemonicValid(string(mnemonics)) {
			return nil, errors.Errorf("mnemonic is invalid")
		}
	}
	return newWallet(a.dagParams, mnemonics)
}
