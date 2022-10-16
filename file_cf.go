package cf

import (
	"errors"
	"os"
)

// 读取文件 返回string
func ReadFile(path string) (string, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

// 创建文件夹
func CreateFolder(path string) (bool, error) {
	if FileExists(path) {
		return true, nil
	}
	// 不存在时创建文件夹
	if err := os.MkdirAll(path, 0775); err != nil {
		return false, errors.New("创建目录失败")
	}
	return true, nil
}

// 判断所给路径文件/文件夹是否存在
func FileExists(path string) bool {
	//os.Stat获取文件信息
	_, err := os.Stat(path)
	if err != nil {
		if ok := os.IsExist(err); ok {
			return true
		}
		return false
	}
	return true
}

// 取文件大小
func FileSize(path string) int64 {
	//os.Stat获取文件信息
	info, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return info.Size()
}
