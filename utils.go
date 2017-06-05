package filer

import (
	"fmt"
	"path"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`(\d+)([KMGT]B$|B)`)
}

const (
	BYTES     = 1
	KILOBYTES = 1024 * BYTES
	MEGABYTES = 1024 * KILOBYTES
	GIGABYTES = 1024 * MEGABYTES
)

//LengthInBytes is a helper function that parses an human readable string and returns it's equivalent in bytes.
//Supported units are B,KB,MB and GB
func LengthInBytes(format string) (int64, error) {

	slicedParts := re.FindStringSubmatch(format)

	if x := len(slicedParts); x == 0 || x != 3 {
		return -1, fmt.Errorf(`Invalid format %s`, format)
	}

	//Underscoring the error here because if we ever get here,
	//we should be left with a string that is parseable by strconv (thanks regex)
	humanReadableSize, _ := strconv.ParseInt(slicedParts[1], 10, 64)

	unit := slicedParts[2]

	if unit == "KB" {
		return humanReadableSize * KILOBYTES, nil
	}

	if unit == "MB" {
		return humanReadableSize * MEGABYTES, nil
	}

	if unit == "GB" {
		return humanReadableSize * GIGABYTES, nil
	}

	//We would assume it is in bytes format already (e.g "1B")
	return humanReadableSize * BYTES, nil
}

//Extension returns the known extension of a given file
func Extension(f File) string {
	return NormalizeExtension(path.Ext(f.Name())[1:])
}

//This is here so as to remove all non-aphabetic characters.
//The reasoning behind this is the fact that files are saved in the temp dir of the
//system and Go suffixes them with some weird integer hence path.Ext would return
//the integer alongside the original extension
func NormalizeExtension(s string) string {

	return strings.Map(func(r rune) rune {

		if upper := unicode.ToUpper(r); upper > 'A' && upper <= unicode.MaxASCII {
			return r
		}

		return -1
	}, s)
}
