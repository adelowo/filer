package validator_test

import (
	"mime/multipart"

	. "github.com/adelowo/filer/validator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Extension", func() {

	var val Validator

	BeforeEach(func() {
		val = NewExtensionValidator([]string{"go", "php", "md", "rb", "ts"})
	})

	Context("When validating a file with an invalid extension", func() {

		It("should fail with an error ", func() {
			_, err := val.Validate(&multipart.FileHeader{Filename: "index.js"})

			Expect(err).To(HaveOccurred())
		})

		It("should have a falsy value", func() {
			isValid, _ := val.Validate(&multipart.FileHeader{Filename: "index.js"})

			Expect(isValid).To(BeFalse())
		})
	})

	Context("When validating a file with a valid extension", func() {
		It("should have a valid extension", func() {
			Expect(val.Validate(&multipart.FileHeader{Filename: "main.go"})).To(BeTrue())
		})

		It("should not have an error", func() {
			_, err := val.Validate(&multipart.FileHeader{Filename: "log.rb"})

			Expect(err).NotTo(HaveOccurred())
		})
	})
})
