package main

import (
	_ "embed"
	"encoding/base64"

	"github.com/svarogg/kaspaper/model"
)

//go:embed template.html
var template string

type walletTemplate struct {
	Mnemonic  *model.MnemonicString
	Address   string
	AddressQR string
}

func walletToWalletTempalte(wallet model.KaspaperWallet) (*walletTemplate, error) {
	const addressIndex = 0

	address, err := wallet.Address(addressIndex)
	if err != nil {
		return nil, err
	}

	addressQRbytes, err := wallet.AddressQR(addressIndex)
	if err != nil {
		return nil, err
	}
	addressQRBase64 := base64.StdEncoding.EncodeToString(addressQRbytes)

	return &walletTemplate{
		Mnemonic:  wallet.Mnemonic(),
		Address:   address,
		AddressQR: addressQRBase64,
	}, nil
}
