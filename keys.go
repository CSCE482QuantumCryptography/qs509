package qs509

import (
	"fmt"
	"os/exec"
)

func GenerateKey(keyAlg SignatureAlgorithm, keyOut string) (bool, error) {

	checkInit()

	cmd := exec.Command(openSSLPath, "genpkey", "-algorithm", keyAlg.Get(), "-out", keyOut)

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	fmt.Println(string(output))

	return true, nil
}
