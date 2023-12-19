package gengorm

import (
	"fmt"
	"go/format"
	"gorm.io/gorm"
	"lark/pkg/constant"
	"lark/pkg/utils"
	"log/slog"
	"os"
	"strings"
)

func QueryTableColumn(db *gorm.DB, dbName string, tableName string) ([]TableColumn, error) {
	var columns []TableColumn
	sqlTableColumn := fmt.Sprintf("SELECT `ORDINAL_POSITION`,`COLUMN_NAME`,`COLUMN_TYPE`,`DATA_TYPE`,`COLUMN_KEY`,`IS_NULLABLE`,`EXTRA`,`COLUMN_COMMENT`,`COLUMN_DEFAULT` FROM `information_schema`.`columns` WHERE `table_schema`= '%s' AND `table_name`= '%s' ORDER BY `ORDINAL_POSITION` ASC",
		dbName, tableName)

	rows, err := db.Raw(sqlTableColumn).Rows()
	if err != nil {
		fmt.Printf("execute query table column action error, detail is [%v]\n", err.Error())
		return columns, err
	}
	defer rows.Close()

	for rows.Next() {
		var column TableColumn
		err = rows.Scan(
			&column.OrdinalPosition,
			&column.ColumnName,
			&column.ColumnType,
			&column.DataType,
			&column.ColumnKey,
			&column.IsNullable,
			&column.Extra,
			&column.ColumnComment,
			&column.ColumnDefault)
		if err != nil {
			fmt.Printf("query table column scan error, detail is [%v]\n", err.Error())
			return columns, err
		}
		columns = append(columns, column)
	}
	return columns, err
}

func Capitalize(s string) string {
	var upperStr string
	chars := strings.Split(s, "_")
	for _, val := range chars {
		vv := []rune(val)
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				if vv[i] >= 97 && vv[i] <= 122 {
					vv[i] -= 32
				}
				upperStr += string(vv[i])
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr
}

func ToCamelCase(str string) string {
	parts := strings.Split(str, "_")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}
	return strings.Join(parts, "")
}

func FieldType(s string) string {
	return constant.ColumnTypes[s]
}

func ColumnType(s string) string {
	return ";type:" + s
}

func IsNull(s string) string {
	if s == "NO" {
		return ";not null"
	}
	return ""
}

func PrimaryKey(s string) string {
	if s == "PRI" {
		return ";primary_key"
	}
	return ""
}

func DefaultValue(s string, columnType string) string {
	if s != "" {
		return ";default:" + s
	}
	switch columnType {
	case "string":
		return ";default:''"
	case "time.Time":
		return ";default:'1001-01-01'"
	case "int32", "int64", "int", "uint32", "uint64", "uint", "float32", "float64":
		return "0"
	}
	return ""
}

func Comment(s string) string {
	if s != "" {
		return ";comment:" + s
	}
	return ""
}

func CreateFile(code string, directory string, model string) (err error) {
	var (
		filePath      = directory + "po_" + model + ".go"
		exists        bool
		formattedCode []byte
	)
	if exists, _ = utils.PathExists(filePath); exists == true {
		return
	}
	if exists, _ = utils.PathExists(directory); exists == false {
		err = utils.Mkdir(directory)
		if err != nil {
			slog.Error(err.Error())
			return
		}
	}
	formattedCode, err = format.Source([]byte(code))
	if err != nil {
		slog.Error(err.Error())
		return
	}
	err = os.WriteFile(filePath, formattedCode, 0776)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	return
}
