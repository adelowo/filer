package storage_test

import (
	"bytes"
	"os"
	"strings"

	. "github.com/adelowo/filer/storage"
	"github.com/spf13/afero"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ Store = (*LocalAdapter)(nil)

var _ = Describe("Local", func() {

	var fs afero.Fs

	Context("Writing a file", func() {

		BeforeEach(func() {
			fs = afero.NewMemMapFs()
		})

		It("should not have an error", func() {
			local := NewLocalAdapter("uploads", fs)
			r := bytes.NewReader([]byte("Lanre Adelowo"))

			Expect(local.Write("users/42", r)).NotTo(HaveOccurred())
		})
	})

	Context("Deleting a file", func() {

		It("Should not have any error", func() {
			By("Deleting the file", func() {
				local := NewLocalAdapter("uploads", fs)
				//Write some data
				local.Write("users/42", strings.NewReader("Lanre Adelowo"))

				//Then delete it
				Expect(local.Delete("users/42")).NotTo(HaveOccurred())
			})
		})

		It("Cannot delete a non existent file", func() {
			local := NewLocalAdapter("uploads", fs)

			err := local.Delete("unknown/path")

			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(&os.PathError{}))
		})
	})
})
