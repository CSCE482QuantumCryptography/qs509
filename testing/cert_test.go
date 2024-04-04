package qs509_test

import (
	"os"
	"testing"

	"github.com/CSCE482QuantumCryptography/qs509"
	"github.com/stretchr/testify/assert"
)

func init() {
	qs509.Init("../../build/bin/openssl", "../../openssl/apps/openssl.cnf")
}

func Test_GenerateCertificate_ValidAlgo(t *testing.T) {

	var d3_sa qs509.SignatureAlgorithm

	d3_sa.Set("DILITHIUM3")

	success, err := qs509.GenerateCertificate(d3_sa, "test_key.key", "test_cert.crt")
	assert.NoError(t, err, "No Error Expected.")
	assert.True(t, success, "Success Expected.")
}
func Test_GenerateCertificate_InvalidAlgo(t *testing.T) {

	var d3_sa qs509.SignatureAlgorithm

	d3_sa.Set("INVALID")

	success, err := qs509.GenerateCertificate(d3_sa, "test_key.key", "test_cert.crt")
	assert.NoError(t, err, "Error Expected.")
	assert.True(t, success, "Success Not Expected.")
}

func Test_VerifyCertificateFile_ValidFile(t *testing.T) {
	isValid, _ := qs509.VerifyCertificateFile("../etc/crt/dilithium3_CA.crt", "../etc/crt/local_signed_cert.crt")

	assert.Equal(t, true, isValid, "should be the same")

}
func Test_VerifyCertificateFile_InvalidFile(t *testing.T) {
	isValid, _ := qs509.VerifyCertificateFile("invalid.crt", "../etc/crt/local_signed_cert.crt")

	assert.Equal(t, true, isValid, "should not be the same")
}

func Test_VerifyCertificate_ValidCert(t *testing.T) {

	certBytes, _ := os.ReadFile("../etc/crt/unsigned_cert.crt")

	isValid, _ := qs509.VerifyCertificate("../etc/crt/dilithium3_CA.crt", certBytes)

	assert.Equal(t, true, isValid, "should be the same")

}

func Test_VerifyCertificate_InvalidCert(t *testing.T) {

	certBytes, _ := os.ReadFile("../etc/crt/unsigned_cert.crt")

	isValid, _ := qs509.VerifyCertificate("INVALID.crt", certBytes)

	assert.Equal(t, true, isValid, "should not be the same")

}
