package validator

import (
	"errors"

	"github.com/adelowo/filer"
)

//ErrFileInvalidExtension is an error type that is rendered when
//an extension is shown
var ErrFileInvalidExtension = errors.New(
	`extension validator: File has an invalid extension type`)

//ExtensionValidator is a validator that validates a file based on it's extension
//Extremely dumb and cannot be trusted
type ExtensionValidator struct {
	validExtensions []string
}

//NewExtensionValidator returns an instance of an ExtensionValidator
func NewExtensionValidator(allowedExtensions []string) *ExtensionValidator {
	return &ExtensionValidator{validExtensions: allowedExtensions}
}

//Validate checks if a file is valid by looking at it's extension
func (e *ExtensionValidator) Validate(f filer.File) (bool, error) {
	if isValidExtension(
		e.validExtensions, filer.Extension(f)) {
		return true, nil
	}

	return false, ErrFileInvalidExtension
}

func isValidExtension(allowed []string, current string) bool {
	var valid bool

	for _, v := range allowed {
		if v == current {
			valid = true
			break
		}
	}

	return valid
}
