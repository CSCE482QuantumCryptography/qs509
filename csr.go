package qs509

import (
	"fmt"
	"os/exec"
)

func GenerateCsr(keyAlg SignatureAlgorithm, keyOut string, csrOut string) {

	checkInit()

	cmd := exec.Command(openSSLPath, "req", "-new", "-newkey", keyAlg.Get(), "-keyout", keyOut, "-out", csrOut, "-nodes", "-subj", "/CN=test server", "-config", openSSLConfigPath)

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(output))

	return
}

func SignCsr(csrPath string, crtOut string, caCrtPath string, caKeyPath string) {

	checkInit()

	cmd := exec.Command(openSSLPath, "x509", "-req", "-in", csrPath, "-out", crtOut, "-CA", caCrtPath, "-CAkey", caKeyPath, "-CAcreateserial", "-days", "365")

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(output))

	return
}
