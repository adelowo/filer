package validator_test

import (
	"io/ioutil"
	"os"

	. "github.com/adelowo/filer/validator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Extension", func() {

	var val Validator
	var file *os.File

	BeforeEach(func() {
		file, _ = os.Open("./fixtures/gopher.jpg")
		val = NewExtensionValidator([]string{"go", "php", "md", "rb", "ts"})
	})

	Context("When validating a file with an invalid extension", func() {

		It("should fail with an error ", func() {

			_, err := val.Validate(file)

			Expect(err).To(HaveOccurred())
		})

		It("should have a falsy value", func() {
			isValid, _ := val.Validate(file)

			Expect(isValid).To(BeFalse())
		})
	})

	Context("When validating a file with a valid extension", func() {
		BeforeEach(func() {
			val = NewExtensionValidator([]string{"jpg", "png"})
		})

		It("should have a valid extension", func() {
			Expect(val.Validate(file)).To(BeTrue())
		})

		It("should not have an error", func() {
			_, err := val.Validate(file)

			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("Reading the name of a file saved in the temporary directory", func() {

		It("Should remove all non alphabetic values left", func() {
			val = NewExtensionValidator([]string{"jpg", "png"})

			//This would return something like gopher.png1234
			file, err := ioutil.TempFile("", "gopher.png")

			defer os.Remove(file.Name())

			defer file.Close()

			Expect(err).ToNot(HaveOccurred())

			Expect(val.Validate(file)).To(BeTrue())
		})
	})

})
