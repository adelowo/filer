package validator_test

import (
	"bytes"
	"os"

	. "github.com/adelowo/filer/validator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Size", func() {

	var val *SizeValidator
	var file *os.File

	Context("Validating a file", func() {

		BeforeEach(func() {
			val = NewSizeValidator(int64(60), int64(70))
			file, _ = os.Open("./fixtures/gopher.jpg")
		})

		AfterEach(func() {
			file.Close()
		})

		It("should return an error", func() {

			By("Inspecting the file size and deeming it too large", func() {

				_, err := val.Validate(file)

				Expect(err).Should(HaveOccurred())
				Expect(err).To(Equal(ErrFileSizeTooLarge))
			})

			By("Inspecting the file size and deeming it too small", func() {
				val := NewSizeValidator(200, 183)

				_, err := val.Validate(file)
				Expect(err).To(Equal(ErrFileSizeTooSmall))
			})
		})

		It("passes validation", func() {
			val := NewSizeValidator(250, 70)

			isValid, err := val.Validate(file)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(isValid).To(BeTrue())
		})
	})

	It("returns an error if the file metadata cannot be read", func() {
		_, err := val.Validate(&mock{})

		Expect(err).Should(HaveOccurred())
		Expect(err).Should(BeAssignableToTypeOf(&os.PathError{}))
	})
})

type mock struct{}

func (m *mock) Name() string { return "Mock" }

func (m *mock) Stat() (os.FileInfo, error) {
	return nil, &os.PathError{Op: "stat"}
}

func (m *mock) Read(p []byte) (n int, err error) {
	return -1, bytes.ErrTooLarge
}
