package generator_test

import (
	"github.com/adelowo/filer/generator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Random", func() {

	It("Should panic if a random string of length zero or less is requested", func() {
		go func() {
			defer GinkgoRecover()
			generator.NewRandomGenerator(0)
		}()
	})

	It("Generates a random string with the file extension suffixed", func() {
		ran := generator.NewRandomGenerator(12)

		generatedName := ran.Generate(&mock{"filename.jpg"})

		Expect(generatedName).To(HaveSuffix(".jpg"))
	})

	It("Generates a random string", func() {

		ran := generator.NewRandomGenerator(10)

		generatedName := ran.Generate(&mock{"filename.jpg"})

		Expect(generatedName).NotTo(BeEmpty())
	})
})
