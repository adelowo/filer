
### Filer - Utils for dealing with file uploads and validation in Golang

[![Coverage Status](https://coveralls.io/repos/github/adelowo/filer/badge.svg)](https://coveralls.io/github/adelowo/filer)[![Build Status](https://img.shields.io/travis/adelowo/filer/master.svg?style=flat-square)](https://travis-ci.org/adelowo/filer.svg?branch=master)

- [Installation](#install)
- [Utilities](#utils)
- [Example usage](#usage)

<div id="install"> </div>

```bash

$ go get github.com/adelowo/filer

```

<div id="utils"> </div>

- Validators : Validate files according to a set of rules

  - `SizeValidator`

  > Validate a file by it's size

  ```go
  max, _ := filer.LengthInBytes("2MB")
  min, _ := filer.LengthInBytes("200KB")

  val = NewSizeValidator(max, min)
  file, _ = os.Open("./path/to/file")

  isValid, err := val.Validate(file)
  ```

  - `MimeTypeValidator`

  > Validate a file by it's mimetype.

  ```go
  val = NewMimeTypeValidator([]string{"image/jpeg", "image/png"})
  file, _ = os.Open("./path/to/file")

  isValid, err := val.Validate(file)
  ```

  - `ExtensionValidator`

  > Validate a file by it's extension.. Caveat, this is probably not what you need.

  ```go
  val = NewExtensionValidator([]string{"go", "php", "md", "rb", "ts"})
  file, _ = os.Open("./path/to/file")

  isValid, err := val.Validate(file)
  ```

> You can also make use of a chained validator to ease the pain of having to deal with multiple validators manually

```go

validator := NewChainedValidator(
  NewExtensionValidator([]string{"go", "ts", "jpg"}),
  NewSizeValidator((1024*1024), (1024*6))) //1MB and 6 KB

file, _ := os.Open("./path/to/file.jpg")

isValid, err := validator.Validate(file)

```

- File Name Generators:

  - `MD5Generator` :
  > Get the ms5 hash sum of the file name

  ```go
  gen := generator.NewMD5Generator()
  s := gen.Generate(file) //Assuming file is `os.Open("file")`
  ```

  - `SlugGenerator` :

  > Generates a slugified version of the file name.

  ```go
  gen := generator.NewSlugGenerator()
  s := gen.Generate(file)
  ```

  - `RandomGenerator` :

  > This generates a random name for the file.. Discards the file name itself

  ```go
  ran := generator.NewRandomGenerator(12) //random name with a length of 12
  generatedName := ran.Generate(file)
  ```

- File Storage:
 > For storage, filer utilizes [Afero][afero]. This is to allow making use of a single API while having multiple backends to choose from

 ```go

 fs = afero.NewMemMapFs() //Using an inmemory store here..
 //Could be a local store,s3 store or anything.. As long as it implements `afero.Fs`
 storeAdapter = storage.NewFilerStorage(fs, nil) //can also pass in a PathFunc instead of nil

 ```

<div id="usage"> </div>

A sample usage of this library would be :

```go
package main

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/adelowo/filer/storage"
	"github.com/adelowo/filer/validator"
	"github.com/spf13/afero"
)

func main() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			io.WriteString(w, "You came here")
			return
		}

		if r.Method == "POST" {
			file, header, err := r.FormFile("formNameInHTMLform")

			if err != nil {
				io.WriteString(w, err.Error())
				return
			}

			buf, err := ioutil.ReadAll(file)

			if err != nil {
				io.WriteString(w, err.Error())
				return
			}

			//Move to a temporary location
			//This is allow us be able to get details like it's size and others

			f, err := ioutil.TempFile("", header.Filename)

			if err != nil {
				panic("An error occurred while trying to create a temporary file")
			}

			val = validator.NewMimeTypeValidator([]string{"image/jpeg", "image/png"})
			val2 = validator.NewSizeValidator((1024 * 1024 * 2), (200 * 1024)) //2MB(maxSize) and 200KB(minSize)

			if _, err := val.Validate(f); err != nil {
				panic("Validation failed")
			}
			if _, err := val2.Validate(f); err != nil {
				panic("Validation failed")
			}


			//Upload to some place
			fs = afero.NewMemMapFs() //Using an inmemory store here..
			storeAdapter = storage.NewFilerStorage(fs, nil)

			if err := storeAdapter.Write("some/path", f); err != nil {
				panic("An error occurred while writing the file")
			}

			io.WriteString(w, "The upload was successful")
      return
		}
	})
}

```

### License
MIT

[afero]: https://github.com/spf13/afero
