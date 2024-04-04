package qs509

import (
	"testing"
	"time"
	"os"
	"fmt"

	"github.com/CSCE482QuantumCryptography/qs509"
)

func Test_Benchmark_Sunny(t *testing.T) {
	// Delete any exiting Instance Files
	_, err := os.Stat("../benchmarkLog/benchmarkInstance.xlsx");
	if err == nil {
		os.Remove("../benchmarkLog/benchmarkInstance.xlsx")
    }

    // Run Benchmark Code
	startTime := time.Now()
	endTime := time.Now()
	qs509.Benchmark(startTime, endTime)

	// Output Fail if the Instance File was not created
	_, err = os.Stat("../benchmarkLog/benchmarkInstance.xlsx");
	if os.IsNotExist(err) {
		fmt.Println("Instance File not generated.")
		t.Fail();
    }

    // Output if Existing Tracking File is in the correct location
    _, err = os.Stat("../benchmarkLog/benchmarkTime.xlsx")
    if os.IsNotExist(err) {
		fmt.Println("Tracking File missing.")
		t.Fail();
    }
}

func Test_Benchmark_Rainy(t *testing.T) {
	// Delete any exiting Instance Files
	_, err := os.Stat("../benchmarkLog/benchmarkInstance.xlsx");
	if err == nil {
		os.Remove("../benchmarkLog/benchmarkInstance.xlsx")
    }

    // Run Benchmark Code
	endTime := time.Now()
	startTime := time.Now()
	qs509.Benchmark(startTime, endTime)

	// Output Fail if the Instance File was not created
	_, err = os.Stat("../benchmarkLog/benchmarkInstance.xlsx");
	if os.IsNotExist(err) {
		fmt.Println("Instance File not generated.")
		t.Fail();
    }

    // Output if Existing Tracking File is in the correct location
    _, err = os.Stat("../benchmarkLog/benchmarkTime.xlsx")
    if os.IsNotExist(err) {
		fmt.Println("Tracking File missing.")
		t.Fail();
    }
}

func Test_CreateFile(t *testing.T) {
	qs509.CreateFile("fileName.xlsx")
}
