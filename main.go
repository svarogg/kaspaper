package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"

	"github.com/kaspanet/kaspad/domain/dagconfig"
	"github.com/svarogg/kaspaper/kaspaperlib"
)

func main() {
	if len(os.Args) == 1 {
		printUsage()
		os.Exit(1)
	}
	filename := os.Args[1]

	wallet, err := kaspaperlib.NewAPI(&dagconfig.MainnetParams).GenerateWallet()
	if err != nil {
		printErrorAndExit(err)
	}

	walletString, err := renderWallet(wallet)
	if err != nil {
		printErrorAndExit(err)
	}

	err = ioutil.WriteFile(filename, []byte(walletString), 0600)
	if err != nil {
		printErrorAndExit(errors.WithStack(err))
	}
	fmt.Printf("Paperwallet written to %s\n", filename)
}

func printErrorAndExit(err error) {
	fmt.Fprintf(os.Stderr, "A critical error occured:\n%+v\n", err)
	os.Exit(1)
}

func printUsage() {
	fmt.Println("Usage: kaspaper [filename.html]")
}
