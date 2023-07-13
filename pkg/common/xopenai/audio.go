package xopenai

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
)

func audioMultipart(req *AudioToTextReq, w *multipart.Writer) (err error) {
	var (
		file   *os.File
		writer io.Writer
		reader *bytes.Reader
	)
	if file, err = os.Open(req.File); err != nil {
		return
	}
	defer file.Close()

	if writer, err = w.CreateFormFile("file", file.Name()); err != nil {
		return
	}
	if _, err = io.Copy(writer, file); err != nil {
		return
	}
	if writer, err = w.CreateFormField("model"); err != nil {
		return
	}
	defer w.Close()
	reader = bytes.NewReader([]byte(req.Model))
	if _, err = io.Copy(writer, reader); err != nil {
		return
	}
	return
}
