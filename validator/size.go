package validator

import "mime/multipart"

type SizeValidator struct {
	maxSize int64
}

func NewSizeValidator(maxSize int64) *SizeValidator {
	s := &SizeValidator{}

	if maxSize > 0 {
		s.maxSize = maxSize
	}

	return s
}

func (s *SizeValidator) Validate(m *multipart.FileHeader) bool {

	return true
}
