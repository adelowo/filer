package generator

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/adelowo/filer/validator"
)

type MD5Generator struct {
}

func NewMD5Generator() *MD5Generator {
	return &MD5Generator{}
}

func (m *MD5Generator) Generate(f validator.File) (string, error) {

	sum := md5.Sum([]byte(f.Name()))

	return hex.EncodeToString(sum[:]), nil
}
