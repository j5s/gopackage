package meFile

import (
	"bufio"
	"os"
)

type FileObj struct {

}

func (f *FileObj )ReturnFileContentAsList(filename string)([]string, error)  {
	// 定义文件读取，返回切片对象
	var filelist []string
	file ,err := os.Open(filename)
	if err != nil {
		return []string{},err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		filelist = append(filelist,scanner.Text())
	}
	return filelist,nil
}
