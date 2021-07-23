// iofunc project iofunc.go
package iofunc

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ZlibCompress(src []byte) (returnValue []byte, returnError error) {
	defer func() {
		if err := recover(); err != nil {
			returnError = fmt.Errorf("%v", err)
		}
	}()
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	defer w.Close()
	_, err := w.Write(src)
	if err != nil {
		returnError = err
		return
	}
	w.Flush()
	returnValue = in.Bytes()
	return
}

func ZlibDecompress(src []byte) (returnValue []byte, returnError error) {
	defer func() {
		if err := recover(); err != nil {
			returnError = fmt.Errorf("%v", err)
		}
	}()
	b := bytes.NewReader(src)
	var out bytes.Buffer
	r, err := zlib.NewReader(b)
	if err != nil {
		returnError = err
		return
	}
	defer r.Close()
	io.Copy(&out, r)
	returnValue = out.Bytes()
	return
}

func GZipCompress(byteArray []byte) (returnValue []byte, returnError error) {
	defer func() {
		if err := recover(); err != nil {
			returnError = fmt.Errorf("%v", err)
		}
	}()
	var buffer bytes.Buffer
	writer, err := gzip.NewWriterLevel(&buffer, gzip.BestCompression)
	if err != nil {
		returnError = err
		return
	}
	defer writer.Close()
	_, err = writer.Write(byteArray)
	if err != nil {
		returnError = err
		return
	}
	writer.Flush()
	returnValue = buffer.Bytes()
	return
}

func GZipDecompress(byteArray []byte) (returnValue []byte, returnError error) {
	defer func() {
		if err := recover(); err != nil {
			returnError = fmt.Errorf("%v", err)
		}
	}()
	var buffer *bytes.Buffer = bytes.NewBuffer(byteArray)
	reader, err := gzip.NewReader(buffer)
	if err != nil {
		returnError = err
		return
	}
	defer reader.Close()
	contents, err := ioutil.ReadAll(reader)
	if err != nil {
		//returnError = err
		//return
	}
	returnValue = contents
	return
}

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
