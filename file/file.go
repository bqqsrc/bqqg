// Copyright (C) 晓白齐齐,版权所有.

package file

import (
	"os"
	"path/filepath"
)

// 路径是否存在
func Exist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

// 路径是否为文件
func IsFile(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !file.IsDir()
}

// 路径是否为目录
func IsDir(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return file.IsDir()
}

// 拼接路径
func JoinPath(path ...string) string {
	return filepath.Join(path...)
}

// 拼接路径，并判断拼接后的路径是否存在
// 返回参数：拼接后的路径，是否存在
func JoinPathExist(path ...string) (string, bool) {
	newPath := JoinPath(path...)
	return newPath, Exist(newPath)
}

// 拼接路径，并判断拼接后的路径是否为文件
// 返回参数：合并后的路径，是否为文件（如果不存在或者是目录，将返回false）
func JoinPathIsFile(path ...string) (string, bool) {
	newPath := JoinPath(path...)
	return newPath, IsFile(newPath)
}
