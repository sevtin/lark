package template

import (
	"strings"
	"text/template"
)

func ParseTemplate(t string) *template.Template {
	//if strings.HasPrefix(t, "\n") == true {
	//	t = t[1:]
	//}
	t = strings.TrimSpace(t)
	tpl, err := template.New("output_template").Parse(t)
	if err != nil {
		panic(err)
	}
	return tpl
}
