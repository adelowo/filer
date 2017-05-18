package validator

import (
	"errors"
	"os"
)

//ErrFileSizeTooLarge is an error type for files that signifies a file s
//too large
var ErrFileSizeTooLarge = errors.New(`
	size validator: File size is too large`)

var ErrFileSizeTooSmall = errors.New(`
	size validator: File size is too small`)

//SizeValidator is a validator that checks the size of a file
//to determine if it is valid or not based on predefined criterias
type SizeValidator struct {
	maxSize, minSize int64
}

//NewSizeValidator returns an instance of a SizeValidator
func NewSizeValidator(maxSize, minSize int64) *SizeValidator {
	s := &SizeValidator{}

	if maxSize > 0 {
		s.maxSize = maxSize * 1024
	}

	if minSize > 0 {
		s.minSize = minSize * 1024
	}

	return s
}

//Validate validates a file based on it's size
func (s *SizeValidator) Validate(f *os.File) (bool, error) {

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
