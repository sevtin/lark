package utils

import (
	"bytes"
	"fmt"
	"go/format"
	"lark/scripts/gencode/config"

	"os"
	"text/template"
)

const (
	MODE_PERM_0776 os.FileMode = 0766
)

func GenCode(tpl *template.Template, conf *config.GenConfig) (err error) {
	defer func() {
		if err != nil {
			fmt.Println(err.Error())
		}
	}()
	var (
		wr            = new(bytes.Buffer)
		path          = conf.Path + "/" + conf.Prefix + conf.PackageName + conf.Suffix
		exists        bool
		formattedCode []byte
	)
	switch conf.FileType {
	case config.FILE_TYPE_GO:
		path += ".go"
	case config.FILE_TYPE_PROTO:
		path += ".proto"
	case config.FILE_TYPE_YAML:
		path += ".yaml"
	}
	err = tpl.Execute(wr, conf.Dict)
	if err != nil {
		return
	}
	if conf.Path == "" {
		return
	}
	// 避免覆盖
	if exists, err = pathExists(path); exists == true {
		return
	}
	if exists, err = pathExists(conf.Path); exists == false {
		mkdir(conf.Path)
	}
	switch conf.FileType {
	case config.FILE_TYPE_GO:
		if formattedCode, err = format.Source(wr.Bytes()); err != nil {
			return
		}
		wr = bytes.NewBuffer(formattedCode)
	}
	err = os.WriteFile(path, wr.Bytes(), MODE_PERM_0776)
	return
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func mkdir(path string) (err error) {
	err = os.MkdirAll(path, MODE_PERM_0776)
	if err != nil {
		return
	}
	err = os.Chmod(path, MODE_PERM_0776)
	return
}
