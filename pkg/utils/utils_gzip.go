package utils

import (
	"bytes"
	"compress/gzip"
)

func Gzip(in []byte) (out []byte, err error) {
	var (
		buffer bytes.Buffer
		writer *gzip.Writer
	)
	writer = gzip.NewWriter(&buffer)
	_, err = writer.Write(in)
	if err != nil {
		_ = writer.Close()
		return
	}
	if err = writer.Close(); err != nil {
		return
	}
	out = buffer.Bytes()
	return
}

func GzipEncode(in string) (out string, err error) {
	var (
		buf []byte
	)
	buf, err = Gzip(Str2Bytes(in))
	if err != nil {
		return
	}
	out = Bytes2Str(buf)
	return
}

func UnGzip(in []byte) (out []byte, err error) {
	var (
		gzipReader *gzip.Reader
		buffer     *bytes.Buffer
	)
	gzipReader, err = gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		return
	}
	defer func() {
		_ = gzipReader.Close()
	}()
	buffer = new(bytes.Buffer)
	if _, err = buffer.ReadFrom(gzipReader); err != nil {
		return
	}
	out = buffer.Bytes()
	return
}

func GzipDecode(in string) (out string, err error) {
	var (
		buf []byte
	)
	buf, err = UnGzip(Str2Bytes(in))
	if err != nil {
		return
	}
	out = Bytes2Str(buf)
	return
}
