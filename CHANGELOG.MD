# Changelog

## 1.1.2 (2017-07-29)

- Improved performance of `MimeTypeValidator` by not reading the entire file into memory. Since the algorithm used at https://golang.org/pkg/net/http/#DetectContentType considers at most the first 512 bytes, it makes no sense to get the entire bytes.

## 1.1.1 (2017-06-05)

- Fixed bug in `LengthInBytes` accepting invalid units (such as "1KBB","1MBB")

## 1.1.0 (2017-06-03)

- Added `ChainedValidator` to ease dealing with multiple validators manually
- Added `LengthInBytes` to parse a human readable string into it's size n bytes.

## 1.0.0 (2017-05-30)

- Initial release
- Comes with validators, file name generators and abstraction for storage
