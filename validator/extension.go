package validator

import (
	"mime/multipart"
	"strings"
)

const extensionSeparator = '.'

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
func (e *ExtensionValidator) Validate(m *multipart.FileHeader) bool {
	//Stub
	return true
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

func getExtensionFromFileName(fileName string) string {
	slice := strings.Split(fileName, string(extensionSeparator))
	return slice[len(slice)-1]
}
