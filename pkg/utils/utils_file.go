// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package utils

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// CopyFile will copy a file from src path to dst path.
// Overwrites any existing files at dst.
// Permissions are copied from file at src to the new file at dst.
func CopyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	if err = os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
		return
	}
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	stat, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, stat.Mode())
	if err != nil {
		return
	}

	return
}

// CopyDir will copy a directory and all contained files and directories.
// src must exist and dst must not exist.
// Permissions are preserved when possible. Symlinks are skipped.
func CopyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	stat, err := os.Stat(src)
	if err != nil {
		return
	}
	if !stat.IsDir() {
		return fmt.Errorf("source must be a directory")
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if err == nil {
		return fmt.Errorf("destination already exists")
	}

	err = os.MkdirAll(dst, stat.Mode())
	if err != nil {
		return
	}

	items, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, item := range items {
		srcPath := filepath.Join(src, item.Name())
		dstPath := filepath.Join(dst, item.Name())

		if item.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			if item.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}

var SizeLimitExceeded = errors.New("Size limit exceeded")

type LimitedReaderWithError struct {
	limitedReader *io.LimitedReader
}

func NewLimitedReaderWithError(reader io.Reader, maxBytes int64) *LimitedReaderWithError {
	return &LimitedReaderWithError{
		limitedReader: &io.LimitedReader{R: reader, N: maxBytes + 1},
	}
}

func (l *LimitedReaderWithError) Read(p []byte) (int, error) {
	n, err := l.limitedReader.Read(p)
	if l.limitedReader.N <= 0 && err == io.EOF {
		return n, SizeLimitExceeded
	}
	return n, err
}

func Remove(name string) (err error) {
	fmt.Println(name)
	return os.Remove(name)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Mkdir(path string) (err error) {
	err = os.MkdirAll(path, 0776)
	if err != nil {
		return
	}
	err = os.Chmod(path, 0776)
	return
}

//func AppendToFile(filename string, content string) (err error) {
//	var (
//		file *os.File
//		wr   = new(bytes.Buffer)
//		buf  []byte
//	)
//	file, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
//	if err != nil {
//		return
//	}
//	defer file.Close()
//	wr.WriteString(content)
//	if buf, err = format.Source(wr.Bytes()); err != nil {
//		return
//	}
//	wr = bytes.NewBuffer(buf)
//	_, err = io.WriteString(file, content)
//	if err != nil {
//		return
//	}
//	return
//}
