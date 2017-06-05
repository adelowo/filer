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

	var _ = DescribeTable(`
		Converting human readable strings to their equivalent in Bytes`,
		func(original string, expected int64, hasError bool) {
			val, err := filer.LengthInBytes(original)

			if hasError {
				Expect(err).To(HaveOccurred())
				return
			}

			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal(expected))
		},

		Entry("Converting an unsupported unit type", "1ZB", int64(-1), true),
		Entry("An error should occur when trying to convert a word into an integer",
			"twoKB", int64(-1), true),
		Entry("Converting Single digit byte", "1B", int64(1), false),
		Entry("Converting bytes with a double digit number", "10B", int64(10), false),
		Entry("Converting Single digit Kilobyte", "1KB", int64(1024), false),
		Entry("Converting kilobytes with a double digit number", "12KB", int64(12288), false),
		Entry("Conversion should not work for 'K' but 'KB' alone", "1K", int64(-1), true),
		Entry("Converting a single digit Megabyte", "1MB", int64(1048576), false),
		Entry("Converting a double digit Megabyte", "31MB", int64(32505856), false),
		Entry("Conversion should not work for 'M' but 'MB' alone", "31M", int64(-1), true),
		Entry("Converting a single digit Gigabyte", "1GB", int64(1073741824), false),
		Entry("Converting a double digit Gigabyte", "20GB", int64(21474836480), false),
		Entry("Conversion should not work for 'G' but 'GB' alone", "1M", int64(-1), true),

		Entry("Conversion should not work for 'KBB'", "1KBB", int64(-1), true),
		Entry("Conversion should not work for 'GBB'", "1GBB", int64(-1), true),
		Entry("Conversion should not work for 'MB'", "MBB", int64(-1), true),
		Entry("Conversion should not work for 'TBB'", "TKBB", int64(-1), true),
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
