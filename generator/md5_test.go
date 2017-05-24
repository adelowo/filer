package generator_test

import (
	"bytes"
	"os"

	"github.com/adelowo/filer"
	"github.com/adelowo/filer/generator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Md5", func() {

	var gen generator.Generator
	var f filer.File

	BeforeEach(func() {
		gen = generator.NewMD5Generator(false)
		f = &mock{}
	})

	It("Should not have an error", func() {
		_, err := gen.Generate(f)

		Expect(err).NotTo(HaveOccurred())
	})

	It("Should generate a string for a non empty string", func() {

		sum, _ := gen.Generate(&mock{"Non-empty-file"})
		testGeneratedHashSum(sum)
	})

	It("Should generate a name for an empty string", func() {
		sum, _ := gen.Generate(f)
		testGeneratedHashSum(sum)
	})

	It("Should generate a name and keep the extension of the original file name",
		func() {
			f = &mock{"picture.jpg"}
			gen = generator.NewMD5Generator(true)
			sum, _ := gen.Generate(f)
			Expect(sum).To(HaveSuffix("jpg"))
		})

})

func testGeneratedHashSum(hash string) {
	Expect(hash).NotTo(BeEmpty())
	Expect(hash).To(HaveLen(32))
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
