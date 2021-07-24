package iofunc

import (
	"io"
	"os"
)

func CopyFile(src string, dst string) (returnValue error) {
	srcFile, err := os.Open(src)
	if err != nil {
		returnValue = err
		return
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		returnValue = err
		return
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		returnValue = err
		return
	}
	return
}

func Exists(name string) (returnValue bool) {
	returnValue = true
	if _, err := os.Stat(name); os.IsNotExist(err) {
		returnValue = false
	}
	return
}
