package qs509

import (
	"fmt"
	"os/exec"
)

func GenerateKey(spec string) {
	
	cmd := exec.Command("../build/bin/openssl", "genpkey", "-algorithm", "dilithium3", "-out", "dilithium3_srv.key")

	output, err := cmd.CombinedOutput()

	if err != nil {
        fmt.Println(err.Error())
        return
    }

    // Print the output
    fmt.Println(string(output))

	return


}