package qs509

import (
	"bytes"
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

func VerifyCertificateFile(caCrtPath string, certToVerify string) (bool, error) {

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

func VerifyCertificate(caCrtPath string, certBytes []byte) (bool, error) {

	checkInit()

	cmd := exec.Command(openSSLPath, "verify", "-CAfile", caCrtPath)

	var outBuffer bytes.Buffer
	cmd.Stdout = &outBuffer

	inPipe, err := cmd.StdinPipe()
	if err != nil {
		return false, err
	}

	if err := cmd.Start(); err != nil {
		return false, err
	}

	if _, err := inPipe.Write(certBytes); err != nil {
		return false, err
	}

	if err := inPipe.Close(); err != nil {
		return false, err
	}

	if err := cmd.Wait(); err != nil {
		return false, err
	}

	return true, nil

}
