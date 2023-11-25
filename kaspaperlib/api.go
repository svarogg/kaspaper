package kaspaperlib

import (
	"github.com/kaspanet/kaspad/cmd/kaspawallet/libkaspawallet"
	"github.com/kaspanet/kaspad/domain/dagconfig"
	"github.com/svarogg/kaspaper/model"
	"github.com/kaspanet/kaspad/cmd/kaspawallet/utils"
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
		return newWallet(a.dagParams, mnemonics)//
	} else {
	// It's safe to use [0] because we know there's exactly 1 key, since we passed numKeys: 1 to CreateMnemonics
		return newWallet(a.dagParams, mnemonics)
	}
}
