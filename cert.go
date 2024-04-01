package qs509

import (
	"crypto/x509"
	"fmt"
	"os/exec"
)

func GenerateCertificate(keyAlg SignatureAlgorithm, keyOut string, certOut string) string {

	checkInit()

	cmd := exec.Command(openSSLPath, "req", "-x509", "-new", "-newkey", keyAlg.Get(), "-keyout", keyOut, "-out", certOut, "-nodes", "-subj", "/CN=test CA", "-days", "365", "-config", openSSLConfigPath)

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("error")
		fmt.Println(err.Error())
		return string(output)
	}

	fmt.Println(string(output))

	return string(output)
}

func VerifyCertificate(cert *x509.Certificate) string {

	return "not implemented yet"

}
