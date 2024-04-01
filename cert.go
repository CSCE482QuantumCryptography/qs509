package qs509

import (
	"fmt"
	"os/exec"
)

func GenerateCertificate(keyAlg SignatureAlgorithm, keyOut string, certOut string) (bool, error) {

	checkInit()

	cmd := exec.Command(openSSLPath, "req", "-x509", "-new", "-newkey", keyAlg.Get(), "-keyout", keyOut, "-out", certOut, "-nodes", "-subj", "/CN=test CA", "-days", "365", "-config", openSSLConfigPath)

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	fmt.Println(string(output))

	return true, nil
}

func VerifyCertificate(caCrtPath string, certToVerify string) (bool, error) {

	checkInit()

	cmd := exec.Command(openSSLPath, "verify", "-CAfile", caCrtPath, certToVerify)

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	fmt.Println(string(output))

	return true, nil

}
