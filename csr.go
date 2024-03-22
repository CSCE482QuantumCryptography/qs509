package qs509

import (
    "fmt"
    "os/exec"
)

func CreateCsr(algo SignatureAlgorithm) {

    cmd := exec.Command("build/bin/openssl", "req", "-x509", "-new", "-newkey", algo.Get(), "-keyout", "dilithium3_CA.key", "-out", "dilithium3_CA.crt", "-nodes", "-subj", "/CN=test CA", "-days", "365", "-config", "openssl/apps/openssl.cnf")

    output, err := cmd.CombinedOutput()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println(string(output))

    return
}

