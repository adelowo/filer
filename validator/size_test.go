package validator_test

import (
	"mime/multipart"

	. "github.com/adelowo/filer/validator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Size", func() {

	var m *multipart.FileHeader
	var val *SizeValidator
	var fileFixturePath string

	Context("Validating a file with an invalid path", func() {

		BeforeEach(func() {
			val = NewSizeValidator(int64(60), int64(70))
			fileFixturePath = "./fixtures/image.pn"
			m = &multipart.FileHeader{Filename: fileFixturePath}
		})

		It("should return an error", func() {
			_, err := val.Validate(m)

			Expect(err).To(HaveOccurred())
		})

		It("should return a falsy value", func() {
			isValid, err := val.Validate(m)
			Expect(err).To(HaveOccurred())
			Expect(isValid).Should(BeFalse())
		})

	})

	Context("Validating a file", func() {

		BeforeEach(func() {
			val = NewSizeValidator(int64(60), int64(70))
			fileFixturePath = "./fixtures/image.png"
			m = &multipart.FileHeader{Filename: fileFixturePath}
		})

		It("should return an error", func() {

			By("Inspecting the file size and deeming it too large", func() {
				_, err := val.Validate(m)

				Expect(err).Should(HaveOccurred())
				Expect(err).To(Equal(ErrFileSizeTooLarge))
			})

			By("Inspecting the file size and deeming it too small", func() {
				val := NewSizeValidator(int64(120), int64(100))

				_, err := val.Validate(m)
				Expect(err).To(Equal(ErrFileSizeTooSmall))
			})
		})

		It("passes validation", func() {
			val := NewSizeValidator(int64(250), int64(70))

			isValid, err := val.Validate(m)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(isValid).To(BeTrue())
		})
	})
})
