package filer_test

import (
	"bytes"
	"os"
	"strings"

	"github.com/adelowo/filer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {

	var _ = DescribeTable("It should remove all numeric characters from the extension",
		func(original, expected string) {
			Expect(strings.EqualFold(expected,
				filer.NormalizeExtension(original))).To(BeTrue())
		},
		Entry("Single numeric character", "MD1", "MD"),
		Entry("Multiple numeric characters", "MD1234", "MD"),
		Entry("Multiple numeric characters", "567MD1234", "MD"),
	)

	It("Returns the extension when given a file name", func() {
		Expect(filer.Extension(&mock{name: "fileName.MD"})).
			To(Equal("MD"))
	})
})

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
