package generator

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/adelowo/filer"
)

//MD5Generator genrates the md5 sum of a file name
type MD5Generator struct {
}

//NewMD5Generator returns an instance of an MD5Generator
func NewMD5Generator() *MD5Generator {
	return &MD5Generator{}
}

func (m *MD5Generator) Generate(f filer.File) string {

	sum := md5.Sum([]byte(f.Name()))

	sumAsString := hex.EncodeToString(sum[:])

	return sumAsString + "." + filer.Extension(f)
}
