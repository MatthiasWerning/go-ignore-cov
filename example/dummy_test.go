package example_test

import (
	"testing"

	"github.com/MatthiasWerning/go-ignore-cov/example"
)

func TestSayHello(t *testing.T) {
	example.MaybeSayHello()
}
