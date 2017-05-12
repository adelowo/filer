package validator

import "mime/multipart"

type Validator interface {
	Validate(*multipart.FileHeader) bool
}
