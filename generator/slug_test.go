package generator_test

import (
	"strings"

	"github.com/adelowo/filer/generator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Slug", func() {

	var gen = generator.NewSlugGenerator()

	var _ = DescribeTable("Slugifies name",
		func(original, slugified string) {
			slugged, _ := gen.Generate(&mock{original})
			Expect(strings.EqualFold(slugged, slugified)).
				To(BeTrue())
		},
		Entry("Slufies name with single space",
			"some name", "some-name"),
		Entry("Slufies name with multiple spaces",
			"some multi  spaced name", "some-multi-spaced-name"),
	)
})
