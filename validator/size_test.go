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
			val = NewSizeValidator(155648, 145408) //152KB and 142KB
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
				val := NewSizeValidator(1048576, 204800) //1MB and 200KB

				_, err := val.Validate(file)
				Expect(err).To(Equal(ErrFileSizeTooSmall))
			})
		})

		It("passes validation", func() {
			val := NewSizeValidator((168 * 1024), (100 * 1024)) //168KB and 100KB

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
