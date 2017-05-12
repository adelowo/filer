package validator

import (
	"testing"
)

var _ Validator = (*ExtensionValidator)(nil)

func Test_getExtensionFromFileName(t *testing.T) {

	cases := map[string]struct {
		fileName, expected string
	}{
		"filename with one dot": {"somefile.go", "go"},
		"file with two dots":    {"go.trumps.php", "php"},
	}

	for k, v := range cases {
		if ext := getExtensionFromFileName(v.fileName); ext != v.expected {
			t.Fatalf(`
        Test for a %s failed ==== File extensions do not match..\n
        Expected %s, Got %s`, k, v.expected, ext)
		}
	}
}

func Test_isValidExtension(t *testing.T) {

	cases := []struct {
		allowedExts []string
		currentExt  string
		isValid     bool
	}{
		{[]string{"go", "php", "rb"}, "js", false},
		{[]string{"go", "php", "rb"}, "go", true},
	}

	for _, v := range cases {
		if ok := isValidExtension(v.allowedExts, v.currentExt); ok != v.isValid {
			t.Fatalf(`
        Encountered a non valid extension.. \n
        Expected %v Got %v`, v.isValid, ok)
		}
	}
}
