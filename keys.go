package qs509

import (
	"fmt"
	"os/exec"
	"time"
	"net"
	"crypto/x509"
)

func GenerateCertificate() string{
	cmd := exec.Command("../build/bin/openssl", "req", "-x509", "-new", "-newkey", "dilithium5", "-keyout", "../dilithium5_CA.key", "-out", "../dilithium5_CA.crt", "-nodes", "-subj", "/CN=test CA", "-days", "365", "-config", "../openssl/apps/openssl.cnf")

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("error")
        fmt.Println(err.Error())
        return string(output)
    }

    // Print the output
    fmt.Println(string(output))

	return string(output)
}

func VerifyCertificate(cert *x509.Certificate) string{
	// parse from the cert var
	toBeVerified, err := x509.parseCertificate(cert)

	// Check that certificate is valid (not before, not after, SignatureAlgorithm)
	fmt.Println(toBeVerified.notBefore)
	fmt.Println(toBeVerified.notAfter)
	fmt.Println(toBeVerified.SignatureAlgorithm)

	return "done"
}