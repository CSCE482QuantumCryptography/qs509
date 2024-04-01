package qs509_test

import (
	"testing"

	"github.com/CSCE482QuantumCryptography/qs509"
	// "github.com/stretchr/testify/assert"
)

func init() {
	qs509.Init("../../../build/bin/openssl", "../../../openssl/apps/openssl.cnf")
}

func Test_GenerateCsr(t *testing.T) {

	var d3_sa qs509.SignatureAlgorithm

	d3_sa.Set("DILITHIUM3")

	qs509.GenerateCsr(d3_sa, "d3_keyOut.key", "d3_csrOut.csr")

}

func Test_SignCsr(t *testing.T) {

	qs509.SignCsr("./d3_csrOut.csr", "d3_crtOut.crt", "../etc/dilithium3_CA.crt", "../etc/dilithium3_CA.key")

}
