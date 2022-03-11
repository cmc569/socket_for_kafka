package file

import (
	"os"
)

func CreateDirectoryIfNotExist(directory string) error{
	if exist := IsFileExists(directory); !exist{
		return os.MkdirAll(directory, 0777)
	}
	return nil
}

func IsFileExists(path string) bool {
	if path == ""{
		return false
	}
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}