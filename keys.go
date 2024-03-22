package qs509

import (
	"fmt"
	"os/exec"
)

func GenerateKey() string{
	cmd := exec.Command("../build/bin/openssl", "req", "-x509", "-new", "-newkey", "dilithium3", "-keyout", "dilithium3_CA.key", "-out", "dilithium3_CA.crt", "-nodes", "-subj", "/CN=test CA", "-days", "365", "-config", "openssl/apps/openssl.cnf")

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