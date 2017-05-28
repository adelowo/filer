package storage_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"github.com/adelowo/filer/storage"
	"github.com/spf13/afero"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ storage.Store = (*storage.FilerStorage)(nil)

var _ = Describe("Local", func() {

	var storeAdapter storage.Store
	var fs afero.Fs

	BeforeEach(func() {
		fs = afero.NewMemMapFs()
		storeAdapter = storage.NewFilerStorage(fs, nil)
	})

	Context("Writing a file", func() {

		It("Should have an error if the file cannot be written to", func() {
			fs = afero.NewReadOnlyFs(fs)
			storeAdapter = storage.NewFilerStorage(fs, nil)

			Expect(storeAdapter.Write("names/42",
				strings.NewReader("Lanre Adelowo"))).To(HaveOccurred())
		})

		It("should not have an error", func() {
			r := bytes.NewReader([]byte("Lanre Adelowo"))
			Expect(storeAdapter.Write("users/42", r)).NotTo(HaveOccurred())
		})

		//TODO (adelowo): Add tests for situations when io.Reader returns an error
	})

	Context("Deleting a file", func() {

		It("Should not have any error", func() {
			By("Deleting the file", func() {
				//Write some data
				storeAdapter.Write("users/42", strings.NewReader("Lanre Adelowo"))

				//Then delete it
				Expect(storeAdapter.Delete("users/42")).NotTo(HaveOccurred())
			})
		})

		It("Cannot delete a non existent file", func() {
			err := storeAdapter.Delete("unknown/path")

			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(&os.PathError{}))
		})
	})

	It("Checks if a file exists", func() {

		By("returning an error if the file doesn't exist", func() {
			_, err := storeAdapter.Exists("somepath")

			Expect(err).To(HaveOccurred())
			Expect(err).To(BeEquivalentTo(storage.ErrFileDoesNotExists))
		})

		By("Returning a falsy value if the file does not exist", func() {
			exists, _ := storeAdapter.Exists("somepath")

			Expect(exists).To(BeFalse())
		})

		By("Returning a truthy value and no error if the file exists", func() {

			storeAdapter.Write("somepath", strings.NewReader("Yup! Just wrote some data"))

			exists, err := storeAdapter.Exists("somepath")

			Expect(err).NotTo(HaveOccurred())
			Expect(exists).To(BeTrue())
		})
	})

	It("Should return the URL for a given path", func() {
		expected := filepath.Join("avatars", "lanre", "large", "x.jpg")

		Expect(storeAdapter.URL("avatars/lanre/large/x.jpg")).Should(Equal(expected))
	})

	It("Makes use of a custom filepath generator", func() {

		pathPrefix := filepath.Join("oops", "whoops")

		storeAdapter = storage.NewFilerStorage(fs, func(path string) string {
			return filepath.Join(pathPrefix, path)
		})

		expected := filepath.Join(pathPrefix, "shoops")

		Expect(storeAdapter.URL("shoops")).Should(Equal(expected))
	})
})
