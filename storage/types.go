package storage

import (
	"io"
)

type Store interface {
	Write(path string, r io.Reader) error
	Delete(path string) error
	URL(path string) (string, error)
	Has(path string) (bool, error)
}
