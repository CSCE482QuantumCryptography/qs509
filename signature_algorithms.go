package qs509

import (
	"errors"
	"fmt"
	"strings"
)

type SignatureAlgorithm string

func (f *SignatureAlgorithm) String() string {
	return fmt.Sprint(string(*f))
}

func (f *SignatureAlgorithm) Set(value string) error {
	value = strings.ToUpper(value)
	if !strings.Contains(strings.ToUpper(SupportedSignatureAlgs), value) {
		return errors.New("invalid option -- must specify a supported algorithm! Check SignatureAlgorithm.Type() for a list of supported algorithms")
	}
	*f = SignatureAlgorithm(value)
	return nil
}

func (f *SignatureAlgorithm) Get() string {
	return string(*f)
}

func (f *SignatureAlgorithm) Type() string {
	return SupportedSignatureAlgs
}
