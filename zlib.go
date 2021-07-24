// iofunc project iofunc.go
package iofunc

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
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
