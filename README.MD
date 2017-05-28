
### Filer - Utils for dealing with file uploads and validation in Golang

[![Coverage Status](https://coveralls.io/repos/github/adelowo/filer/badge.svg)](https://coveralls.io/github/adelowo/filer)[![Build Status](https://img.shields.io/travis/adelowo/filer/master.svg?style=flat-square)](https://travis-ci.org/adelowo/filer.svg?branch=master)

- [Installation](#install)
- [Utilities](#utils)

<div id="install"> </div>

```bash

$ go get https://github.com/adelowo/filer

```

<div id="utils"> </div>

- Validators : Validate files according to a set of rules

  - `SizeValidator`

  > Validate a file by it's size

  ```go
  val = NewSizeValidator(int64(60), int64(70))
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

### License
MIT

[afero]: https://github.com/spf13/afero
