package examples

import (
	"os"
	"path/filepath"
)

func Panic() {
	panic("a proble")

	path := filepath.Join(os.TempDir(), "file")
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}
}
