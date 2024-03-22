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
	if value != "EC" && value != "RSA" && value != "DILITHIUM3" {
		return errors.New("invalid option -- must specify either 'EC' or 'RSA' or 'DILITHIUM3")
	}
	*f = SignatureAlgorithm(value)
	return nil
}

func (f *SignatureAlgorithm) EC() bool {
	return f.Get() == "EC"
}

func (f *SignatureAlgorithm) RSA() bool {
	return f.Get() == "RSA"
}

func (f *SignatureAlgorithm) D3() bool {
	return f.Get() == "DILITHIUM3"
}

func (f *SignatureAlgorithm) Get() string {
	return string(*f)
}

func (f *SignatureAlgorithm) Type() string {
	return "RSA|EC|DILITHIUM3"
}
