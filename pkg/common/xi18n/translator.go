package xi18n

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

// Translator 结构体用于存储和管理翻译
type Translator struct {
	sync.RWMutex
	translations map[string]map[any]string
	defaultLang  string // 默认语言
}

// NewTranslator 创建一个新的 Translator 实例
func NewTranslator() *Translator {
	return &Translator{
		translations: make(map[string]map[any]string),
		defaultLang:  LANG_EN,
	}
}

// AddTranslation 添加一个翻译
func (t *Translator) AddTranslation(lang string, key any, translation string) {
	t.Lock()
	defer t.Unlock()
	if t.translations[lang] == nil {
		t.translations[lang] = make(map[any]string)
	}
	t.translations[lang][key] = translation
}

// Translate 翻译一个 key
func (t *Translator) Translate(lang string, key any) (translation string) {
	t.RLock()
	defer t.RUnlock()
	var (
		langMap map[any]string
		ok      bool
	)
	if langMap, ok = t.translations[lang]; ok {
		if translation, ok = langMap[key]; ok {
			return
		}
	}
	if t.defaultLang == "" {
		translation, _ = toString(key)
		return
	}
	if langMap, ok = t.translations[t.defaultLang]; ok {
		if translation, ok = langMap[key]; ok {
			return
		}
	}
	translation, _ = toString(key)
	return
}

func toString(i any) (string, error) {
	if i == nil {
		return "", nil
	}
	switch s := i.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case int:
		return strconv.Itoa(s), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case int32:
		return strconv.Itoa(int(s)), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case uint:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(s), 10), nil
	case json.Number:
		return s.String(), nil
	case []byte:
		return string(s), nil
	case template.HTML:
		return string(s), nil
	case template.URL:
		return string(s), nil
	case template.JS:
		return string(s), nil
	case template.CSS:
		return string(s), nil
	case template.HTMLAttr:
		return string(s), nil
	case nil:
		return "", nil
	case fmt.Stringer:
		return s.String(), nil
	case error:
		return s.Error(), nil
	default:
		return "", fmt.Errorf("unable to cast %#v of type %T to string", i, i)
	}
}

// LoadTranslationsFromFiles 从指定目录加载所有 YAML 翻译文件
func (t *Translator) LoadTranslationsFromFiles(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("error reading directory: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".yaml" || filepath.Ext(file.Name()) == ".yml" {
			lang := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			err := t.loadTranslationFile(filepath.Join(dir, file.Name()), lang)
			if err != nil {
				return fmt.Errorf("error loading translation file %s: %v", file.Name(), err)
			}
		}
	}

	return nil
}

// loadTranslationFile 加载单个翻译文件
func (t *Translator) loadTranslationFile(filePath, lang string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var translations map[any]string
	err = yaml.Unmarshal(data, &translations)
	if err != nil {
		return err
	}

	for key, translation := range translations {
		t.AddTranslation(lang, key, translation)
	}

	return nil
}
