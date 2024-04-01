package qs509

import (
	"fmt"
	"os/exec"
)

func GenerateKey(keyAlg SignatureAlgorithm, keyOut string) {

	checkInit()

	cmd := exec.Command(openSSLPath, "genpkey", "-algorithm", keyAlg.Get(), "-out", keyOut)

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("error")
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(output))

}
