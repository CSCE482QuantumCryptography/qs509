package qs509_test

import (
	"testing"

	"github.com/CSCE482QuantumCryptography/qs509"
	// "github.com/stretchr/testify/assert"
)

func init() {
	qs509.Init("../../../build/bin/openssl", "../../../openssl/apps/openssl.cnf")
}

func Test_GenerateCertificate(t *testing.T) {

	var d3_sa qs509.SignatureAlgorithm

	d3_sa.Set("DILITHIUM3")

	qs509.GenerateCertificate(d3_sa, "keyOut.key", "certOut.crt")

}
