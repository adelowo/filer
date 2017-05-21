package storage_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"

	. "github.com/adelowo/filer/storage"
	"github.com/spf13/afero"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ Store = (*LocalAdapter)(nil)

var _ = Describe("Local", func() {

	var local Store
	var fs afero.Fs

	BeforeEach(func() {
		fs = afero.NewMemMapFs()
		local = NewLocalAdapter("users", fs, nil)
	})

	JustBeforeEach(func() {
		fs.RemoveAll("users")
	})

	Context("Writing a file", func() {

		It("Should have an error if the file cannot be written to", func() {
			fs = afero.NewReadOnlyFs(fs)
			local = NewLocalAdapter("users", fs, nil)

			Expect(local.Write("names/42",
				strings.NewReader("Lanre Adelowo"))).To(HaveOccurred())
		})

		It("should not have an error", func() {
			r := bytes.NewReader([]byte("Lanre Adelowo"))
			Expect(local.Write("users/42", r)).NotTo(HaveOccurred())
		})

		//TODO (adelowo): Add tests for situations when io.Reader returns an erro
	})

	Context("Deleting a file", func() {

		It("Should not have any error", func() {
			By("Deleting the file", func() {
				//Write some data
				local.Write("users/42", strings.NewReader("Lanre Adelowo"))

				//Then delete it
				Expect(local.Delete("users/42")).NotTo(HaveOccurred())
			})
		})

		It("Cannot delete a non existent file", func() {
			err := local.Delete("unknown/path")

			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(&os.PathError{}))
		})
	})

	It("Checks if a file exists", func() {

		By("returning an error if the file doesn't exist", func() {
			_, err := local.Has("somepath")

			Expect(err).To(HaveOccurred())
			Expect(err).To(BeEquivalentTo(ErrLocalFileDoesNotExist))
		})

		By("Returning a falsy value if the file does not exist", func() {
			exists, _ := local.Has("somepath")

			Expect(exists).To(BeFalse())
		})

		By("Returning a truthy value and no error if the file exists", func() {

			local.Write("somepath", strings.NewReader("Yup! Just wrote some data"))

			exists, err := local.Has("somepath")

			Expect(err).NotTo(HaveOccurred())
			Expect(exists).To(BeTrue())
		})
	})

	It("Should return the URL for a given path", func() {
		//From the setup, we have the base directory as "users"
		expected := filepath.Join("users", "avatars", "lanre", "large", "x.jpg")

		Expect(local.URL("avatars/lanre/large/x.jpg")).Should(Equal(expected))
	})

	It("Makes use of a custom filepath generator", func() {

		pathPrefix := filepath.Join("oops", "whoops")

		local = NewLocalAdapter("users", fs, func(path string) string {
			return filepath.Join(pathPrefix, path)
		})

		expected := filepath.Join(pathPrefix, "shoops")

		Expect(local.URL("shoops")).Should(Equal(expected))
	})
})
