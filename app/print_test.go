package app

import (
	"os"
	"testing"
)

type testString struct {
	in  string
	out string
}

var strs []testString

func TestMain(m *testing.M) {
	strs = []testString{
		{"a", "aaa"},
		{"b", "bbb"},
		{"c", "ccc"},
	}

	code := m.Run()

	os.Exit(code)
}

func TestPrint(t *testing.T) {
	for _, str := range strs {
		result := print(str.in)
		if result != str.out {
			t.Errorf("Result of \"%v\", want %v", result, str.out)
		}
	}
}
