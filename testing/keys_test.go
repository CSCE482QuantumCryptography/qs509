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

func Test_GenerateKey_Dilithium3(t *testing.T) {
	var d3_sa qs509.SignatureAlgorithm
	d3_sa.Set("DILITHIUM3")

	success, err := qs509.GenerateKey(d3_sa, "dilithium3_test.key")
	assert.NoError(t, err, "should not have an error")
	assert.True(t, success, "should be successful")

	// Clean up generated key file
	os.Remove("dilithium3_test.key")
}

func Test_GenerateKey_RSA(t *testing.T) {
	var rsa_sa qs509.SignatureAlgorithm
	rsa_sa.Set("RSA")

	success, err := qs509.GenerateKey(rsa_sa, "rsa_test.key")
	assert.NoError(t, err, "should not have an error")
	assert.True(t, success, "should be successful")

	// Clean up generated key file
	os.Remove("rsa_test.key")
}

func Test_GenerateKey_Invalid(t *testing.T) {
	var ab_sa qs509.SignatureAlgorithm
	ab_sa.Set("INVLAID")

	success, err := qs509.GenerateKey(ab_sa, "invalid_test.key")
	assert.Error(t, err, "should have an error")
	assert.False(t, success, "should not be successful")
}
