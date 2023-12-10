package filer

import (
	"github.com/bqqsrc/bqqg/errer"
	"os"
	"path"
)

func callError(caller, format string, args ...interface{}) error {
	return errer.CallerErrNew("filer", caller, format, args...)
}
func PathExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func IsFile(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !file.IsDir()
}

func JoinPathExist(args ...string) bool {
	path := path.Join(args...)
	return PathExist(path)
}

func ReadFile(file string) ([]byte, error) {
	caller := "ReadFile"
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, callError(caller, "os.ReadFile(%s) error: %s", file, err)
	}
	return content, nil
}

func WriteFile(data []byte, file string) error {
	caller := "WriteFile"
	fi, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fi.Close()
	if err != nil {
		return callError(caller, "os.OpenFile(%s, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644) error: %s", file, err)
	}
	if _, err = fi.Write(data); err != nil {
		return callError(caller, "fi.Write error: %s", file, err)
	}
	return nil
}
