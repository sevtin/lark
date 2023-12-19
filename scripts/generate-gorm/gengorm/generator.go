package gengorm

import (
	"fmt"
	"github.com/gertd/go-pluralize"
	"lark/pkg/common/xmysql"
	"lark/pkg/conf"
	"log/slog"
	"strings"
)

func GenGorm(cfg *conf.Mysql, table string, directory string) {
	var (
		model   = pluralize.NewClient().Singular(table)
		columns []TableColumn
		builder strings.Builder
		column  TableColumn
		field   string
		err     error
	)
	xmysql.NewMysqlClient(cfg)
	columns, err = QueryTableColumn(xmysql.GetDB(), cfg.Db, table)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	builder.WriteString("package po\n\n")
	builder.WriteString(fmt.Sprintf("type %s struct {\n", ToCamelCase(model)))
	for _, column = range columns {
		field = fmt.Sprintf("%s      %s    `gorm:\"column:%s%s%s%s\" json:\"%s\"` // %s\n",
			Capitalize(column.ColumnName),
			FieldType(column.DataType),
			column.ColumnName,
			//ColumnType(column.ColumnType),
			PrimaryKey(column.ColumnKey.String),
			DefaultValue(column.ColumnDefault.String, FieldType(column.DataType)),
			IsNull(column.IsNullable),
			//Comment(column.ColumnComment.String),
			column.ColumnName,
			column.ColumnComment.String)
		builder.WriteString(field)
	}
	builder.WriteString("}\n")
	CreateFile(builder.String(), directory, model)
}
