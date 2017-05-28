package generator

import (
	"math/rand"
	"time"

	"github.com/adelowo/filer"
)

func init() {
	rand.Seed(time.Now().Unix())
}

const knownCharacters = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

//RandomGenerator generates a random string that can be used as a file name
//instead of it's original name
type RandomGenerator struct {
	length int
}

//NewRandomGenerator returns an instance of a RandomGenerator
//If length is less than (or equal to) 0, a panic would occur
func NewRandomGenerator(l int) *RandomGenerator {

	if l <= 0 {
		panic("Length cannot be zero or less")
	}

	return &RandomGenerator{length: l}
}

func (r *RandomGenerator) Generate(f filer.File) string {
	byt := make([]byte, r.length)

	for i := range byt {
		byt[i] = knownCharacters[rand.Intn(len(knownCharacters))]
	}

	return string(byt) + "." + filer.Extension(f)
}
