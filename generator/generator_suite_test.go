package generator_test

import (
	"bytes"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGenerator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Generator Suite")
}

type mock struct {
	name string
}

func (m *mock) Name() string { return m.name }

func (m *mock) Stat() (os.FileInfo, error) {
	return nil, &os.PathError{Op: "stat"}
}

func (m *mock) Read(p []byte) (n int, err error) {
	return -1, bytes.ErrTooLarge
}
