package qs509

import (
	"testing"
	"time"

	"github.com/CSCE482QuantumCryptography/qs509"
)

func Test_Benchmark(t *testing.T) {
	startTime := time.Now()

	endTime := time.Now()

	qs509.Benchmark(startTime, endTime)
}
