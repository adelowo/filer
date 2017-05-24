package generator

import "github.com/adelowo/filer/validator"

//Generator defines an interface for generating file names for files
type Generator interface {
	Generate(f validator.File) (string, error)
}
