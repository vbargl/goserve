package os

import (
	"os"
	"path/filepath"
)

func GetWorkingDirectory() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Clean(path)
}

func GetDirectory(path string) string {
	var dir string

	switch {
	case path == "":
		dir = GetWorkingDirectory()
	case path[0] == '/':
		dir = filepath.Clean(path)
	default:
		cwd := GetWorkingDirectory()
		dir = filepath.Join(cwd, path)
	}

	if _, err := os.Stat(dir); err == nil {
		return dir
	} else {
		panic(err)
	}
}
