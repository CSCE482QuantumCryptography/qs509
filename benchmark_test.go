package qs509

import (
	"time"
	"testing"
)

func Test_Benchmark(t *testing.T) {
	startTime := time.Now()

	endTime := time.Now()

	Benchmark(startTime, endTime)
}
