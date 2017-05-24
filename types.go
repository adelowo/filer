package filer

import (
	"io"
	"os"
)

//File is a type that represents a file
//(can be in memory or a concrete file type)
//*os.File already implements this interface
type File interface {
	Name() string
	Stat() (os.FileInfo, error)
	io.Reader
}
