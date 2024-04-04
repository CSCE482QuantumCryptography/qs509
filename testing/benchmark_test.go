package qs509

import (
	"testing"
	"time"
	"fmt"

	"github.com/CSCE482QuantumCryptography/qs509"
)

func Test_Benchmark_Sunny(t *testing.T) {
	startTime := time.Now()
	endTime := time.Now()

	qs509.Benchmark(startTime, endTime)
}

func Test_Benchmark_Rainy(t *testing.T) {
	startTime := time.Now()
	endTime := time.Now()

	qs509.Benchmark(startTime, endTime)

	fmt.Println("test")
}

func Test_CreateFile(t *testing.T) {
	qs509.CreateFile("fileName.xlsx")
}
