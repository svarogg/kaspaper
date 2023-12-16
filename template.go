package main

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"text/template"

	"github.com/svarogg/kaspaper/model"
)

//go:embed template.html
var templateString string

type walletTemplate struct {
	Mnemonic  *model.MnemonicString
	Address   string
	AddressQR string
	KPubKey string
	KPubKeyQR string
}

func renderWallet(wallet model.KaspaperWallet) (string, error) {
	walletTemplate, err := walletToWalletTempalte(wallet)
	if err != nil {
		return "", err
	}

	funcMap := template.FuncMap{
		"sub": func(str string, i, j int) string { return str[i:j] },
	}

	tmpl := template.Must(template.New("kaspaper").Funcs(funcMap).Parse(templateString))

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, walletTemplate)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
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


	kpubKeyQRbytes, err := wallet.KPubKeyQR()
	if err != nil {
		return nil, err
	}
	kpubKeyBase64 := base64.StdEncoding.EncodeToString(kpubKeyQRbytes)
	kpubKey := wallet.KPubKey();



	return &walletTemplate{
		Mnemonic:  wallet.Mnemonic(),
		Address:   address,
		AddressQR: addressQRBase64,
		KPubKey : kpubKey,
		KPubKeyQR: kpubKeyBase64,
	}, nil
}
