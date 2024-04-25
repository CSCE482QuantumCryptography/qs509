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

func Test_GenerateCsr_ValidAlgo(t *testing.T) {

	var d3_sa qs509.SignatureAlgorithm

	d3_sa.Set("DILITHIUM3")

	success, err := qs509.GenerateCsr(d3_sa, "test_d3key.key", "test_d3Csr.csr")
	assert.NoError(t, err, "No Error Expected.")
	assert.True(t, success, "Success Expected.")

	os.Remove("test_d3key.key")
	os.Remove("test_d3Csr.csr")
}

func Test_GenerateCsr_InvalidAlgo(t *testing.T) {

	var d3_sa qs509.SignatureAlgorithm

	d3_sa.Set("DILITHIUM0")

	success, err := qs509.GenerateCsr(d3_sa, "test_d3key.key", "test_d3Csr.csr")
	assert.Error(t, err, "Error Expected.")
	assert.False(t, success, "Success Not Expected.")
}

func Test_SignCsr_ValidCert(t *testing.T) {

	success, err := qs509.SignCsr("../etc/csr/test_d3Csr.csr", "local_signed_cert.crt", "../etc/crt/dilithium3_CA.crt", "../etc/keys/dilithium3_CA.key")
	assert.NoError(t, err, "No Error Expected.")
	assert.True(t, success, "Success Expected.")

	os.Remove("local_signed_cert.crt")
}

func Test_SignCsr_InvalidCert(t *testing.T) {

	success, err := qs509.SignCsr("invalid.csr", "local_signed_cert.crt", "../etc/crt/dilithium3_CA.crt", "../etc/keys/dilithium3_CA.key")
	assert.Error(t, err, "Error Expected.")
	assert.False(t, success, "Success Not Expected.")
}

func Test_SignCsr_InvalidCA(t *testing.T) {

	success, err := qs509.SignCsr("../etc/csr/test_d3Csr.csr", "local_signed_cert.crt", "invalid.crt", "../etc/keys/dilithium3_CA.key")
	assert.Error(t, err, "Error Expected.")
	assert.False(t, success, "Success Not Expected.")
}

func Test_SignCsr_InvalidCAKey(t *testing.T) {

	success, err := qs509.SignCsr("../etc/csr/test_d3Csr.csr", "local_signed_cert.crt", "../etc/crt/dilithium3_CA.crt", "invalid.key")
	assert.Error(t, err, "Error Expected.")
	assert.False(t, success, "Success Not Expected.")
}
