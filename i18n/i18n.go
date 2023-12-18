package i18n

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("./i18n/languages/en.json")
	bundle.MustLoadMessageFile("./i18n/languages/zh-CN.json")
}

func MustLocalize(lang string, msgId string, templateData ...interface{}) string {
	if lang == "" {
		lang = "en"
	}
	if msgId == "" {
		return ""
	}
	localizer := i18n.NewLocalizer(bundle, lang)
	if templateData == nil {
		return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: msgId})
	}
	return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: msgId, TemplateData: templateData[0]})
}
