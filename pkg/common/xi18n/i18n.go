package xi18n

import (
	"fmt"
)

var translator = NewTranslator()

func init() {
	err := translator.LoadTranslationsFromFiles("./i18n/languages")
	if err != nil {
		fmt.Printf("Error loading translations: %v\n", err)
	}
}

func SetDefaultLang(lang string) {
	translator.defaultLang = lang
}

func Translate(lang string, key any) string {
	return translator.Translate(lang, key)
}
