package storage

import (
	"io"
)

//Store defines the interactins available for all storage drivers
type Store interface {
	Write(path string, r io.Reader) error
	Delete(path string) error
	URL(path string) string
	Has(path string) (bool, error)
}
