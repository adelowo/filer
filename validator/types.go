package validator

import "os"

type Validator interface {
	Validate(*os.File) (bool, error)
}
