package os

import (
	// import built-in packages
	"os"
)

// 创建文件夹
func Mkdir(path string) {
	os.MkdirAll(path, 0777)
}

// 创建文件
func CreateFile(path string) {
	os.Create(path)
}
