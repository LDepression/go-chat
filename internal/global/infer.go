/**
 * @Author: lyc
 * @Description:
 * @File:  infer
 * @Version: 1.0.0
 * @Date: 2023/03/20 14:37
 */

package global

import (
	"errors"
	"os"
	"path/filepath"
	"sync"
)

/*
	用于推断当前项目根路径
*/

var (
	once = new(sync.Once)
)

func exist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || errors.Is(err, os.ErrExist)
}

// 计算项目路径
func inferRootDir() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var infer func(path string) string
	infer = func(path string) string {
		if exist(path + "/config") {
			return path
		}
		return infer(filepath.Dir(path))
	}
	RootDir = infer(cwd)
}

func init() {
	once.Do(func() {
		inferRootDir()
	})
}
