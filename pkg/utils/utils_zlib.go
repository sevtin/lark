package utils

import (
	"bytes"
	"compress/zlib"
	"io"
)

func Zlib(in []byte) (out []byte) {
	var (
		buffer bytes.Buffer
		writer *zlib.Writer
	)
	writer = zlib.NewWriter(&buffer)
	writer.Write(in)
	writer.Close()
	out = buffer.Bytes()
	return
}

func UnZlib(in []byte) (out []byte) {
	var (
		r  = bytes.NewReader(in)
		rc io.ReadCloser
	)
	var buf bytes.Buffer
	rc, _ = zlib.NewReader(r)
	io.Copy(&buf, rc)
	out = buf.Bytes()
	return
}
