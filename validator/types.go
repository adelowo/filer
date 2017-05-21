package validator

import (
	"io"
	"os"
)

type Validator interface {
	Validate(File) (bool, error)
}

//File is a type that represents a file meant to be validated
//os.File already implements this interface
type File interface {
	Name() string
	Stat() (os.FileInfo, error)
	io.Reader
}
