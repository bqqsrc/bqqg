//  Copyright (C) 晓白齐齐,版权所有.

package dir

import (
	"os"
)

func Getwd() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return wd
}

func GetwdDefault(defaultDir string) string {
	wd, err := os.Getwd()
	if err != nil {
		return defaultDir
	}
	return wd
}
