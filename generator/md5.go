package generator

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"path"

	"github.com/adelowo/filer/validator"
)

//MD5Generator genrates the md5 sum of a file name
type MD5Generator struct {
	keepExtension bool
}

//NewMD5Generator returns an instance of an MD5Generator
func NewMD5Generator(keepExt bool) *MD5Generator {
	return &MD5Generator{keepExtension: keepExt}
}

func (m *MD5Generator) Generate(f validator.File) (string, error) {

	sum := md5.Sum([]byte(f.Name()))

	sumAsString := hex.EncodeToString(sum[:])

	if m.keepExtension {
		sumAsString = fmt.Sprintf("%s%s", sumAsString, path.Ext(f.Name()))
	}

	return sumAsString, nil
}