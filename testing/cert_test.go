package qs509_test

import (
	"os"
	"testing"

	"github.com/CSCE482QuantumCryptography/qs509"
	"github.com/stretchr/testify/assert"
)

func init() {
	qs509.Init("../../../build/bin/openssl", "../../../openssl/apps/openssl.cnf")
}

func Test_GenerateCertificate(t *testing.T) {

	var d3_sa qs509.SignatureAlgorithm

	d3_sa.Set("DILITHIUM3")

	qs509.GenerateCertificate(d3_sa, "test_key.key", "test_cert.crt")

}

func Test_VerifyCertificateFile(t *testing.T) {
	isValid, _ := qs509.VerifyCertificate("../etc/crt/dilithium3_CA.crt", "../etc/crt/local_signed_cert.crt")

	assert.Equal(t, true, isValid, "should be the same")

}

func Test_VerifyCertificate(t *testing.T) {

	certBytes, _ := os.ReadFile("../etc/crt/unsigned_cert.crt")

	isValid, _ := qs509.VerifyCertificate("../etc/crt/dilithium3_CA.crt", certBytes)

	assert.Equal(t, true, isValid, "should be the same")

}
