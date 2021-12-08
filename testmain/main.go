package main

import (
	"fmt"
	"os"

	"github.com/kaspanet/kaspad/domain/dagconfig"

	"github.com/svarogg/kaspaper/kaspaperlib"
)

func main() {
	api := kaspaperlib.NewAPI(&dagconfig.MainnetParams)
	wallet, err := api.GenerateWallet()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from GenerateWallet: %s", err)
	}

	fmt.Printf("%s\n", wallet.Mnemonic())
	fmt.Printf("%s\n", wallet.KeysJSON())
}
