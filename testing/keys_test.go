package qs509_test

import (
	"os"
	"testing"

	"github.com/CSCE482QuantumCryptography/qs509"
	"github.com/stretchr/testify/assert"
)

func Test_GenerateKey_Dilithium3(t *testing.T) {
	var d3_sa qs509.SignatureAlgorithm
	d3_sa.Set("DILITHIUM3")

	success, err := qs509.GenerateKey(d3_sa, "dilithium3_CA.key")
	assert.NoError(t, err, "should not have an error")
	assert.True(t, success, "should be successful")

	// Clean up generated key file
	os.Remove("dilithium3_CA.key")
}

func Test_GenerateKey_LocalSignedCert(t *testing.T) {
	var rsa_sa qs509.SignatureAlgorithm
	rsa_sa.Set("RSA")

	success, err := qs509.GenerateKey(rsa_sa, "local_signed_cert.key")
	assert.NoError(t, err, "should not have an error")
	assert.True(t, success, "should be successful")

	// Clean up generated key file
	os.Remove("local_signed_cert.key")
}

func Test_GenerateKey_UnsignedCert(t *testing.T) {
	var ec_sa qs509.SignatureAlgorithm
	ec_sa.Set("EC")

	success, err := qs509.GenerateKey(ec_sa, "unsigned_cert.key")
	assert.NoError(t, err, "should not have an error")
	assert.True(t, success, "should be successful")

	// Clean up generated key file
	os.Remove("unsigned_cert.key")
}
