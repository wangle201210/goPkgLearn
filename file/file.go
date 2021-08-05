package file

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"runtime"
)

func T1()  {
	println("in")
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	pa := "./1111.jpg"
	fileWriter, err := bodyWriter.CreateFormFile("image", pa)
	if err != nil {
		println("000",err)
		return
	}
	runtime.GC()
	fh, e := os.Open(pa)
	if e != nil {
		println("1111",e.Error())
		return
	}
	defer fh.Close()

	if _, err = io.Copy(fileWriter, fh); err != nil {
		println("2222",err)
		return
	}
	println("out")
}
