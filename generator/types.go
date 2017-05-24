package generator

import "github.com/adelowo/filer"

//Generator defines an interface for generating file names for files
type Generator interface {
	Generate(f filer.File) string
}
