package file

import (
	"os"
)

// GetContents 把整个文件读入一个字符串中[PHP:file_get_contents]
func GetContents(filename string) (str string, err error) {
	var data []byte
	data, err = GetContentsToByte(filename)
	if err != nil {
		return
	}
	str = string(data)
	return
}

// GetContentsToByte 把整个文件读入byte中[PHP:file_get_contents]
func GetContentsToByte(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

// PutContents 把一个字符串写入文件中[覆盖原文件内容][PHP:file_put_contents]
func PutContents(filename string, data string) error {
	return PutContentsByByte(filename, []byte(data))
}

// PutContentsToAppend 把一个字符串写入文件中[追加至原文件][PHP:file_put_contents,FILE_APPEND]
func PutContentsToAppend(filename string, data string) error {
	return PutContentsByByteToAppend(filename, []byte(data))
}

// PutContentsByByte 把byte写入文件中[覆盖原文件内容][PHP:file_put_contents]
func PutContentsByByte(filename string, data []byte) error {
	return os.WriteFile(filename, data, os.ModePerm)
}

// PutContentsByByteToAppend 把byte写入文件中[追加至原文件][PHP:file_put_contents,FILE_APPEND]
func PutContentsByByteToAppend(filename string, data []byte) (err error) {
	var f *os.File
	f, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = f.Write(data)
	return
}
