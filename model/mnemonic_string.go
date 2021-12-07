package model

import "strings"

type MnemonicString [24]string

func (ms *MnemonicString) String() string{
	return strings.Join(ms[:], " ")
}
