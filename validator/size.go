package validator

import (
	"errors"

	"github.com/adelowo/filer"
)

//ErrFileSizeTooLarge is an error type for files that signifies a file s
//too large
var ErrFileSizeTooLarge = errors.New(`
	size validator: File size is too large`)

//ErrFileSizeTooSmall is an error type that denotes a file is too small
//Thus the reason behind the validation's failure
var ErrFileSizeTooSmall = errors.New(`
	size validator: File size is too small`)

//SizeValidator is a validator that checks the size of a file
//to determine if it is valid or not based on predefined criterias
type SizeValidator struct {
	maxSize, minSize int64
}

//NewSizeValidator returns an instance of a SizeValidator
//maxSize and minSize are to be given in bytes
func NewSizeValidator(maxSize, minSize int64) *SizeValidator {
	return &SizeValidator{maxSize, minSize}
}

//Validate validates a file based on it's size
func (s *SizeValidator) Validate(f filer.File) (bool, error) {

	info, err := f.Stat()

	if err != nil {
		return false, err
	}

	size := info.Size()

	if size < s.minSize {
		return false, ErrFileSizeTooSmall
	}

	if size > s.maxSize {
		return false, ErrFileSizeTooLarge
	}

	return true, nil
}
