package storage_test

import (
	"bytes"

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
})
