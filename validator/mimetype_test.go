package validator_test

import (
	"mime/multipart"

	. "github.com/adelowo/filer/validator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mimetype", func() {

	var val Validator

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

})
