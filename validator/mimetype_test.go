package validator_test

import (
	"mime/multipart"

	. "github.com/adelowo/filer/validator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mimetype", func() {

	var val Validator
	var m *multipart.FileHeader

	BeforeEach(func() {
		val = NewMimeTypeValidator([]string{"image/jpeg", "image/png"})
	})

	Context("When validating a file with an invalid path", func() {
		It("should return an error", func() {

			isValid, err := val.Validate(&multipart.FileHeader{Filename: "unknown-file"})

			Expect(err).To(HaveOccurred())
			Expect(isValid).To(BeFalse())
		})

	})

	Context("When validating a file with a known path", func() {

		BeforeEach(func() {
			m = &multipart.FileHeader{Filename: "./fixtures/gopher.jpg"}
		})

		It("should not have an error if the mimetype is valid", func() {
			isValid, err := val.Validate(m)

			Expect(err).NotTo(HaveOccurred())
			Expect(isValid).To(BeTrue())
		})

		It("should have a truthy value if the mimetype is valid", func() {
			Expect(val.Validate(m)).To(BeTrue())
		})

		It("should return an errror if the mimetype is invalid", func() {
			val = NewMimeTypeValidator([]string{"application/octet-stream"})

			isValid, err := val.Validate(m)

			Expect(err).To(HaveOccurred())
			Expect(err).Should(Equal(ErrFileInvalidMimeType))

			Expect(isValid).To(BeFalse())
		})

	})

})
