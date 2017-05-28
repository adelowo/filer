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
			slugged := gen.Generate(&mock{original})
			Expect(strings.EqualFold(slugged, slugified)).
				To(BeTrue())
		},
		Entry("Slufies name with single space",
			"some name.MD", "some-name.MD"),
		Entry("Slufies name with multiple spaces",
			"some multi  spaced name.MD", "some-multi-spaced-name.MD"),
	)
})
