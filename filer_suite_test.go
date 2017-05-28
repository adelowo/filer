package filer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFiler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Filer Suite")
}
