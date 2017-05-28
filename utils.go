package filer

import (
	"path"
	"strings"
)

//Extension returns the known extension of a given file
func Extension(f File) string {
	return normalizeExtension(path.Ext(f.Name())[1:])
}

//This is here so as to remove all non-aphabetic characters.
//The reasoning behind this is the fact that files are saved in the temp dir of the
//system and Go suffixes them with some weird integer hence path.Ext would return
//the integer alongside the original extension
func normalizeExtension(s string) string {

	return strings.Map(func(r rune) rune {

		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return r
		}

		return -1
	}, s)
}
