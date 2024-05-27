package constant

var (
	ColumnTypes = map[string]string{
		"tinyint":           "int8",
		"smallint":          "int16",
		"mediumint":         "int32",
		"int":               "int32",
		"integer":           "int64",
		"bigint":            "int64",
		"float":             "float32",
		"double":            "float64",
		"decimal":           "float64",
		"date":              "time.Time", // or "string"
		"time":              "time.Time", // or "string"
		"year":              "int16",     // or "string"
		"datetime":          "time.Time",
		"timestamp":         "time.Time",
		"char":              "string",
		"varchar":           "string",
		"tinyblob":          "[]byte",
		"tinytext":          "string",
		"blob":              "[]byte",
		"text":              "string",
		"mediumblob":        "[]byte",
		"mediumtext":        "string",
		"longblob":          "[]byte",
		"longtext":          "string",
		"binary":            "[]byte",
		"varbinary":         "[]byte",
		"unsigned bigint":   "uint64",
		"unsigned int":      "uint32",
		"unsigned smallint": "uint16",
		"unsigned tinyint":  "uint8",
		"bit":               "uint64",
		"enum":              "string",
		"set":               "string",
		"json":              "string",
	}
)
