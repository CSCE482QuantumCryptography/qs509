package qs509

import (
	"testing"
	"crypto/tls"
)

func Test_GenerateCertificate(t *testing.T) {
	GenerateCertificate()
	return
}
func Test_VerifyCertificate(t *testing.T) {
	cert, err = tls.LoadX509KeyPair("../dilithium5_CA.crt", "../dilithium5_CA.key")

	

	if err != nil {
		t.Fail()
	}

	VerifyCertificate(conn)
	return
}