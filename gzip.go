package iofunc

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
)

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
