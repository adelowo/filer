package generator

import (
	"math/rand"
	"path"

	"github.com/adelowo/filer"
)

const knownCharacters = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

//RandomGenerator generates a random string that can be used as a file name
//instead of it's original name
type RandomGenerator struct {
	length        int
	keepExtension bool
}

//NewRandomGenerator returns an instance of a RandomGenerator
//If length is less than (or equal to) 0, a panic would occur
//If keepExt is false, only the random string would be returned
//If true, the random string and the file extension is returned
func NewRandomGenerator(l int, keepExt bool) *RandomGenerator {

	if l <= 0 {
		panic("Length cannot be zero or less")
	}

	return &RandomGenerator{length: l, keepExtension: keepExt}
}

func (r *RandomGenerator) Generate(f filer.File) string {
	byt := make([]byte, r.length)

	for i := range byt {
		byt[i] = knownCharacters[rand.Intn(len(knownCharacters))]
	}

	ret := string(byt)

	if r.keepExtension {
		ret = ret + path.Ext(f.Name())
	}

	return ret
}
