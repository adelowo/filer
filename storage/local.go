package storage

import (
	"errors"
	"io"
	"path/filepath"

	"github.com/spf13/afero"
)

const (
	defaultDirectoryFilePerm = 0755
	defaultFilePerm          = 0666
)

//ErrLocalFileDoesNotExist is an error type that indicates that a file
//with a given path does not exist on the disk
var ErrLocalFileDoesNotExist = errors.New("Local: File does not exist")

//KeyFunc is a file path generator
type KeyFunc func(path string) string

//LocalAdapter is a storage implementation that deals with file operations
//on a physical disk
type LocalAdapter struct {
	baseDir string
	afero   *afero.Afero
	gen     KeyFunc
}

//NewLocalAdapter returns an instance of the Local adapter
func NewLocalAdapter(baseDir string, f afero.Fs, gen KeyFunc) *LocalAdapter {
	l := &LocalAdapter{baseDir: baseDir, afero: &afero.Afero{Fs: f}}

	if gen == nil {
		l.gen = l.filePath
	} else {
		l.gen = gen
	}

	return l
}

func (l *LocalAdapter) Write(path string, r io.Reader) error {

	buf, err := afero.ReadAll(r)

	if err != nil {
		return err
	}

	if err := l.afero.WriteFile(l.gen(path), buf, defaultFilePerm); err != nil {
		return err
	}

	return nil
}

func (l *LocalAdapter) Delete(path string) error {
	return l.afero.Remove(l.gen(path))
}

func (l *LocalAdapter) Has(path string) (bool, error) {
	exists, err := l.afero.Exists(l.gen(path))

	//The way afero handles Exists is kinda weird though
	//err isn't supposed to be nil if exists is false
	if !exists && (err == nil) {
		return exists, ErrLocalFileDoesNotExist
	}

	return exists, err
}

func (l *LocalAdapter) URL(path string) string {
	return l.gen(path)
}

func (l *LocalAdapter) filePath(path string) string {
	return filepath.Join(l.baseDir, path)
}
