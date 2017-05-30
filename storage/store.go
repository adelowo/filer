package storage

import (
	"errors"
	"io"

	"github.com/spf13/afero"
)

const (
	defaultDirectoryFilePerm = 0755
	defaultFilePerm          = 0666
)

var ErrFileDoesNotExists = errors.New("storage : File does not exist")

//FilerStorage is a storage types that embed afero in other to save files in multiple storage systems
type FilerStorage struct {
	afero *afero.Afero
	gen   PathFunc
}

//PathFunc is a path generator. Manipulate as wish in other to save files as per the app.
type PathFunc func(path string) string

func NewFilerStorage(fs afero.Fs, pathGenerator PathFunc) *FilerStorage {
	store := &FilerStorage{afero: &afero.Afero{Fs: fs}}

	if pathGenerator == nil {
		store.gen = defaultPathFunc
	} else {
		store.gen = pathGenerator
	}

	return store
}

func defaultPathFunc(path string) string {
	return path
}

func (l *FilerStorage) Write(path string, r io.Reader) error {

	buf, err := afero.ReadAll(r)

	if err != nil {
		return err
	}

	if err := l.afero.WriteFile(l.gen(path), buf, defaultFilePerm); err != nil {
		return err
	}

	return nil
}

func (l *FilerStorage) Delete(path string) error {
	return l.afero.Remove(l.gen(path))
}

func (l *FilerStorage) Exists(path string) (bool, error) {

	//underscoring the error here since
	//the way afero handles Exists is kinda weird though
	//err isn't supposed to be nil if exists is false

	if exists, _ := l.afero.Exists(l.gen(path)); !exists {
		return false, ErrFileDoesNotExists
	}

	return true, nil
}

func (l *FilerStorage) URL(path string) string {
	return l.gen(path)
}
