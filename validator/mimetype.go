package validator

import (
	"errors"
	"net/http"

	"github.com/adelowo/filer"
)

//ErrFileInvalidMimeType is an error type denoting a file with a mimetype
//unrecognized by the validator
var ErrFileInvalidMimeType = errors.New(`
  mimetype validator: File has an invalid mimetype`)

//MimeTypeValidator is a validator that checks the mimetype of a file
//and compares it to a list of acceptable mimetypes to determine wheter
//it is valid or not
type MimeTypeValidator struct {
	validMimeTypes []string
}

//NewMimeTypeValidator returns an instance of a MimeTypeValidator
func NewMimeTypeValidator(mimeTypes []string) *MimeTypeValidator {
	return &MimeTypeValidator{validMimeTypes: mimeTypes}
}

//Validate validates a file by looking at it's mimetype.
//Currrently, the mimetype of the file is gotten by through the DetectContentType
//function in net/http.
func (mime *MimeTypeValidator) Validate(f filer.File) (bool, error) {
	buf := make([]byte, 513)
	if _, err := f.Read(buf[0:512]); err != nil {
		return false, err
	}

	return isValidMimeType(mime.validMimeTypes, http.DetectContentType(buf))
}

func isValidMimeType(allowed []string, currentMimeType string) (bool, error) {

	var isValid bool

	for _, v := range allowed {
		if v == currentMimeType {
			isValid = true
			break
		}
	}

	if isValid {
		return true, nil
	}

	return false, ErrFileInvalidMimeType
}
