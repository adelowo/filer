package validator

import (
	"github.com/adelowo/filer"
)

type Validator interface {
	Validate(filer.File) (bool, error)
}
