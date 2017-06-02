package validator_test

import (
	"os"

	. "github.com/adelowo/filer/validator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Chain", func() {


	It("Should properly validate the files based on the chain", func() {
		validator := NewChainedValidator(
			NewExtensionValidator([]string{"go", "ts", "jpg"}),
			NewSizeValidator((1024*1024), (1024*6))) //1MB and 6 KB

		file, err := os.Open("./fixtures/gopher.jpg")

		Expect(err).NotTo(HaveOccurred())
		Expect(validator.Validate(file)).To(BeTrue())
	})

	It(`
		Should return an error if validation within the chain fails`,
		func() {

			validator := NewChainedValidator(
				NewExtensionValidator([]string{"go", "ts"}),
				NewSizeValidator((1024*1024), (1024*6))) //1MB and 6 KB

			file, err := os.Open("./fixtures/gopher.jpg")

			Expect(err).NotTo(HaveOccurred())

			ok, err := validator.Validate(file)
			Expect(err).To(HaveOccurred())
			Expect(ok).To(BeFalse())
		})

	It("Should bail out quickly if only one validator is in the chain",
		func() {
			validator := NewChainedValidator(
				NewSizeValidator((1024 * 1024), (1024 * 6))) //1MB and 6 KB

			file, err := os.Open("./fixtures/gopher.jpg")

			Expect(err).NotTo(HaveOccurred())

			ok, err := validator.Validate(file)
			Expect(err).NotTo(HaveOccurred())
			Expect(ok).To(BeTrue())

		})
})
