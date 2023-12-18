package utils

import (
	"bytes"
	"fmt"
	"go/format"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"lark/scripts/gencode/config"
	"strings"
	"unicode"

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
		path          string
		filename      string
		exists        bool
		formattedCode []byte
	)
	if conf.Filename != "" {
		filename = conf.Filename
	} else {
		filename = conf.PackageName
	}
	path = conf.Path + "/" + conf.Prefix + filename + conf.Suffix
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

func ToCamel(s string) string {
	c := cases.Title(language.English, cases.NoLower)
	return c.String(s)
}

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func CamelToSnake(input string) string {
	var buffer bytes.Buffer
	for i, r := range input {
		if unicode.IsUpper(r) {
			if i > 0 {
				buffer.WriteRune('_')
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

func GetName(serviceName string, apiName string) (name string) {
	if strings.HasPrefix(apiName, serviceName) {
		name = apiName
		return
	}
	name = serviceName + "_" + apiName
	return
}
