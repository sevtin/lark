package utils

import (
	"reflect"
	"strings"
)

func GetFields(obj interface{}) (fields string) {
	var (
		sType     = reflect.TypeOf(obj)
		i         int
		fCount    = sType.NumField()
		lastIndex = fCount - 1
		fieldType reflect.StructField
		builder   strings.Builder
		field     string
	)
	for i = 0; i < fCount; i++ {
		fieldType = sType.Field(i)
		field = fieldType.Tag.Get("field")
		if field == "" {
			continue
		}
		builder.WriteString(field)
		if i != lastIndex {
			builder.WriteString(",")
		}
	}
	fields = builder.String()
	if strings.HasSuffix(fields, ",") {
		fields = fields[:len(fields)-1]
	}
	return
}
