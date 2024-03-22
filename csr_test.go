package qs509

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_CreateCsr(t *testing.T) {
	var d3_sa SignatureAlgorithm

	d3_sa.Set("DILITHIUM3")

	CreateCsr(d3_sa);

}