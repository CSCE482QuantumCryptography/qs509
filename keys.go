package qs509

import (
	"fmt"
	"os/exec"
)

func GenerateKey() string{
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