package generator

import (
	"regexp"

	"github.com/adelowo/filer"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`\s+`)
}

//SlugGenerator is a generator that slugifies the name of a given file
type SlugGenerator struct{}

//NewSlugGenerator returns an instance of a SlugGenerator
func NewSlugGenerator() *SlugGenerator {
	return &SlugGenerator{}
}

func (s *SlugGenerator) Generate(f filer.File) string {
	return re.ReplaceAllString(f.Name(), "-")
}
