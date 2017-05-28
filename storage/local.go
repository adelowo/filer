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

//FilePathFunc is a file path generator
type FilePathFunc func(baseDirectory, path string) string

//LocalAdapter is a storage implementation that deals with file operations
//on a physical disk
type LocalAdapter struct {
	baseDir string
	afero   *afero.Afero
	gen     FilePathFunc
}

//NewLocalAdapter returns an instance of the Local adapter
func NewLocalAdapter(baseDir string, f afero.Fs, gen FilePathFunc) *LocalAdapter {
	l := &LocalAdapter{baseDir: baseDir, afero: &afero.Afero{Fs: f}}

	if gen == nil {
		l.gen = filePathGenerator
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

	if err := l.afero.WriteFile(l.gen(l.baseDir, path), buf, defaultFilePerm); err != nil {
		return err
	}

	return nil
}

func (l *LocalAdapter) Delete(path string) error {
	return l.afero.Remove(l.gen(l.baseDir, path))
}

func (l *LocalAdapter) Exists(path string) (bool, error) {

	//underscoring the error here since
	//the way afero handles Exists is kinda weird though
	//err isn't supposed to be nil if exists is false

	if exists, _ := l.afero.Exists(l.gen(l.baseDir, path)); !exists {
		return false, ErrLocalFileDoesNotExist
	}

	return true, nil
}

func (l *LocalAdapter) URL(path string) string {
	return l.gen(l.baseDir, path)
}

func filePathGenerator(baseDirectory, path string) string {
	return filepath.Join(baseDirectory, path)
}
