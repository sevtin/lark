package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"go/format"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"
)

var (
	db        *gorm.DB
	directory = "./domain/pdo/"
)

func init() {
	var (
		dsn = "root:lark2022@tcp(lark-mysql-user-01:13306)/lark_user?charset=utf8mb4&parseTime=True&loc=Local"
		err error
	)
	dsn = "root:@tcp(127.0.0.1:3306)/canary?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//directory = "./apps/apis/canary/internal/domain/pdo/"
}

func main() {
	sql := `
SELECT wallet_id,wallet_type,uid,pay_password,balance
FROM wallets
`
	_, err := SqlToPdo(db, sql, "WithdrawVerify")
	if err != nil {
		log.Println(err)
	}
}

func SqlToPdo(db *gorm.DB, sql string, obj string) (code string, err error) {
	// 1、获取字段类型
	var (
		cts map[string]string
		sel string
	)
	cts, err = getColumnTypes(db, sql)
	if err != nil {
		return
	}

	// 2、格式化SQL
	sel = formatSQL(sql)
	if sel == "" {
		return
	}

	// 3、处理字段
	var (
		fields [][]string
	)
	fields, sel = processFields(sel)

	// 4、处理列
	var (
		columns []string
	)
	columns = processColumns(sel)

	// 5、生成代码
	code = generateCode(obj, columns, fields, cts, sql)
	err = createFile(code, camelToUnderscore(obj))
	log.Println(code)
	return
}

// 格式化SQL
func formatSQL(sql string) (s string) {
	sql = strings.ReplaceAll(sql, "\n\t", " ")
	sql = strings.ReplaceAll(sql, "\n", " ")
	sql = regexp.MustCompile(`\s+`).ReplaceAllString(sql, " ")
	matches := regexp.MustCompile(`SELECT\s+(.+?)\s+FROM`).FindStringSubmatch(sql)
	if len(matches) != 2 {
		return
	}
	s = matches[1]
	return
}

// 处理字段
func processFields(sql string) ([][]string, string) {
	var (
		regex   = regexp.MustCompile(`\s+`)
		fields  = make([][]string, 0)
		splits  []string
		matches []string
	)
	regex = regexp.MustCompile(`(SUM|IF|COUNT)(.*?) AS \w+|CASE.*? AS \w+`)
	// regex = regexp.MustCompile(`(SUM\(.*?\)) AS \w+|(IF\(.*?\)) AS \w+|CASE.*? AS \w+`)

	matches = regex.FindAllString(sql, -1)
	for _, match := range matches {
		splits = strings.Split(match, " AS ")
		if len(splits) == 2 {
			fields = append(fields, splits)
			sql = strings.ReplaceAll(sql, match+",", "")
			sql = strings.ReplaceAll(sql, match, "")
		}
	}
	return fields, sql
}

// 处理列
func processColumns(sql string) (columns []string) {
	sql = strings.TrimSpace(sql)
	columns = strings.Split(sql, ",")
	return
}

// 生成代码
func generateCode(obj string, columns []string, fields [][]string, cts map[string]string, sql string) string {
	var (
		builder  strings.Builder
		fieldTag = "field_tag_" + camelToUnderscore(obj)
		column   string
	)
	builder.WriteString("package pdo\n\nimport \"lark/pkg/utils\"\n\n")
	builder.WriteString(fmt.Sprintf("type %s struct {\n", obj))

	// 生成列
	for _, column = range columns {
		if column == "" {
			continue
		}
		column = strings.TrimSpace(column)
		jsonTag := toJsonTag(column)
		builder.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\" field:\"%s\"`\n", toPropertyField(column), getFieldType(cts, jsonTag), jsonTag, column))
	}

	// 生成字段
	for _, splits := range fields {
		jsonTag := toJsonTag(splits[1])
		builder.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\" field:\"%s\"`\n", toPropertyField(jsonTag), getFieldType(cts, jsonTag), jsonTag, strings.Join(splits, " AS ")))
	}
	builder.WriteString("}\n\n")
	// 生成SQL
	builder.WriteString(fmt.Sprintf("/*%s*/\n\n", sql))
	// 字段
	builder.WriteString(fmt.Sprintf("var (\n\t%s string\n)\n\n", fieldTag))
	// 生成 GetFields 方法
	builder.WriteString(fmt.Sprintf("func (p *%s) GetFields() string {\n\tif %s == \"\" {\n\t\t%s = utils.GetFields(*p)\n\t}\n\treturn %s\n}", obj, fieldTag, fieldTag, fieldTag))
	return builder.String()
}

// 创建文件
func createFile(code string, filename string) (err error) {
	var (
		path          = directory
		filePath      = path + "pdo_" + filename + ".go"
		exists        bool
		formattedCode []byte
	)
	if exists, err = pathExists(filePath); exists == true {
		return
	}
	if exists, err = pathExists(path); exists == false {
		mkdir(path)
	}
	formattedCode, err = format.Source([]byte(code))
	if err != nil {
		return
	}
	err = os.WriteFile(filePath, formattedCode, 0776)
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
	err = os.MkdirAll(path, 0776)
	if err != nil {
		return
	}
	err = os.Chmod(path, 0776)
	return
}

func getColumnTypes(db *gorm.DB, s string) (cts map[string]string, err error) {
	var (
		rows     *sql.Rows
		types    []*sql.ColumnType
		ct       *sql.ColumnType
		typename string
		ok       bool
	)
	cts = map[string]string{}
	rows, err = db.Raw(s).Rows()
	if err != nil {
		return
	}
	types, _ = rows.ColumnTypes()
	for _, ct = range types {
		if typename, ok = columnTypes[strings.ToLower(ct.DatabaseTypeName())]; ok == true {
			cts[ct.Name()] = typename
		}
	}
	return
}

func getFieldType(fts map[string]string, name string) (t string) {
	var (
		ok bool
	)
	if t, ok = fts[name]; ok == true {
		return
	}
	t = "string"
	return
}

func toJsonTag(s string) string {
	s = strings.ToLower(s)
	parts := strings.Split(s, ".")
	if len(parts) > 0 {
		s = parts[len(parts)-1]
	}
	parts = strings.Split(s, " ")
	if len(parts) > 0 {
		s = parts[len(parts)-1]
	}
	return s
}

func toPropertyField(s string) string {
	var (
		parts = strings.Split(s, ".")
		c     = cases.Title(language.English, cases.NoLower)
	)
	if len(parts) > 0 {
		s = parts[len(parts)-1]
	}
	parts = strings.Split(s, " ")
	if len(parts) > 0 {
		s = parts[len(parts)-1]
	}
	parts = strings.Split(s, "_")
	for i := range parts {
		// parts[i] = strings.Title(parts[i])
		parts[i] = c.String(parts[i])
	}
	return strings.Join(parts, "")
}

func camelToUnderscore(word string) string {
	var buffer bytes.Buffer
	for i, char := range word {
		if unicode.IsUpper(char) {
			if i > 0 {
				buffer.WriteRune('_')
			}
			buffer.WriteRune(unicode.ToLower(char))
		} else {
			buffer.WriteRune(char)
		}
	}
	return buffer.String()
}

var (
	columnTypes = map[string]string{
		"tinyint":          "int32",
		"smallint":         "int32",
		"mediumint":        "int32",
		"int":              "int32",
		"integer":          "int64",
		"bigint":           "int64",
		"float":            "float64",
		"double":           "float64",
		"decimal":          "float64",
		"date":             "string",
		"time":             "string",
		"year":             "string",
		"datetime":         "time.Time",
		"timestamp":        "time.Time",
		"char":             "string",
		"varchar":          "string",
		"tinyblob":         "string",
		"tinytext":         "string",
		"blob":             "string",
		"text":             "string",
		"mediumblob":       "string",
		"mediumtext":       "string",
		"longblob":         "string",
		"longtext":         "string",
		"unsigned bigint":  "int64",
		"unsigned int":     "int",
		"unsigned tinyint": "int32",
	}
)
