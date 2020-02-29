package util

import (
	"io/ioutil"
	"os"
)

func FileExists(file string) (bool, error) {
	_, err := os.Stat(file)
	if err == nil {
		return true, nil //文件或者文件夹存在
	}
	if os.IsNotExist(err) {
		return false, nil //不存在
	}
	return false, err //不存在，这里的err可以查到具体的错误信息
}

//判断目录是否存在
func IsDir(dir string) bool {

	if exist,_ := FileExists(dir);exist == false{
		return false
	}

	info, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func IsFile(file string) bool {

	if exist,_ := FileExists(file);exist == false{
		return false
	}


	info, err := os.Stat(file)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func IsNonEmptyDir(dir string) bool {
	if exist := IsDir(dir); exist == false{
		return false
	}
	if listDir,_ := ioutil.ReadDir(dir);len(listDir) == 0{
		return false
	}
	return true
}

func CreateFolder(folder string)bool{
	if exist,err:=FileExists(folder); exist == false && err == nil {
		err := os.MkdirAll(folder,os.ModePerm)
		if err != nil {
			return false
		}else{
			return true
		}
	}
	return false
}
