package generator_test

import (
	"github.com/adelowo/filer"
	"github.com/adelowo/filer/generator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Md5", func() {

	var gen generator.Generator
	var f filer.File

	BeforeEach(func() {
		gen = generator.NewMD5Generator()
		f = &mock{}
	})

	It("Should generate a string for a non empty string", func() {

		sum := gen.Generate(&mock{"Non-empty-file.jpg"})
		testGeneratedHashSum(sum)
	})

	It("Should generate a name and keep the extension of the original file name",
		func() {
			f = &mock{"picture.jpg"}
			gen = generator.NewMD5Generator()
			sum := gen.Generate(f)
			Expect(sum).To(HaveSuffix("jpg"))
		})

})

func testGeneratedHashSum(hash string) {
	Expect(hash).NotTo(BeEmpty())
}
