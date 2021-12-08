package main

import (
	"fmt"
	"os"

	"github.com/svarogg/kaspaper/kaspaperlib"
)

func main() {
	api := kaspaperlib.NewAPI()
	wallet, err := api.GenerateWallet()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from GenerateWallet: %s", err)
	}

	fmt.Printf("%s\n", wallet.Mnemonic())
}
