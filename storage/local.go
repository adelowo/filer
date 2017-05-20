package storage

import (
	"io"
	"path/filepath"

	"github.com/spf13/afero"
)

const (
	defaultDirectoryFilePerm = 0755
	defaultFilePerm          = 0666
)

//LocalAdapter is a storage implementation that deals with file operations
//on a physical disk
type LocalAdapter struct {
	baseDir string
	afero   *afero.Afero
}

//NewLocalAdapter returns an instance of the Local adapter
func NewLocalAdapter(baseDir string, f afero.Fs) *LocalAdapter {
	return &LocalAdapter{baseDir: baseDir, afero: &afero.Afero{Fs: f}}
}

func (l *LocalAdapter) Write(path string, r io.Reader) error {

	buf, err := afero.ReadAll(r)

	if err != nil {
		return err
	}

	if err := l.afero.WriteFile(l.filePath(path), buf, defaultFilePerm); err != nil {
		return err
	}

	return nil
}

func (l *LocalAdapter) Delete(path string) error {
	return l.afero.Remove(l.filePath(path))
}

func (l *LocalAdapter) filePath(path string) string {
	return filepath.Join(l.baseDir, path)
}
