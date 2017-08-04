package validator_test

import (
	"os"

	. "github.com/adelowo/filer/validator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mimetype", func() {

	var val Validator
	var file *os.File

	BeforeEach(func() {
		val = NewMimeTypeValidator([]string{"image/jpeg", "image/png"})
		file, _ = os.Open("./fixtures/gopher.jpg")
	})

	AfterEach(func() {
		file.Close()
	})

	It("should return an error if the file cannot be read", func() {
		isValid, err := val.Validate(&mock{})
		Expect(err).To(HaveOccurred())
		Expect(isValid).To(BeFalse())
	})

	It("should not have an error if the mimetype is valid", func() {
		isValid, err := val.Validate(file)

		Expect(err).NotTo(HaveOccurred())
		Expect(isValid).To(BeTrue())
	})

	It("should return an errror if the mimetype is invalid", func() {
		val = NewMimeTypeValidator([]string{"application/octet-stream"})

		isValid, err := val.Validate(file)

		Expect(err).To(HaveOccurred())
		Expect(err).Should(Equal(ErrFileInvalidMimeType))

		Expect(isValid).To(BeFalse())
	})

	It("should have a truthy value if the mimetype is valid", func() {
		Expect(val.Validate(file)).To(BeTrue())
	})

})
