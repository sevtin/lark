package xmysql

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"time"
)

var (
	ERR_DB_INSTANCE_IS_EMPTY = errors.New("database instance is empty")
)

var (
	cli *MysqlClient
)

type MysqlClient struct {
	db  *gorm.DB
	cfg *conf.Mysql
}

func NewMysqlClient(cfg *conf.Mysql) *MysqlClient {
	cli = &MysqlClient{cfg: cfg}
	cli.db, _ = ConnectDB(cfg)
	return cli
}

func GetDB() *gorm.DB {
	if cli.db == nil {
		cli.db, _ = ConnectDB(cli.cfg)
	}
	return cli.db
}

func GetTX() *gorm.DB {
	return GetDB().Begin()
}

// 事务处理
func Transaction(handle func(tx *gorm.DB) (err error)) (err error) {
	var (
		db   *gorm.DB
		terr error
	)
	db = GetDB()
	if db == nil {
		err = ERR_DB_INSTANCE_IS_EMPTY
		return
	}
	tx := db.Begin(&sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	err = handle(tx)
	if err != nil {
		terr = tx.Rollback().Error
		if terr != nil {
			xlog.Error(terr.Error())
		}
		return
	}
	err = tx.Commit().Error
	return
}

func ConnectDB(cfg *conf.Mysql) (db *gorm.DB, err error) {
	var (
		opts     *gorm.Config
		sqlDB    *sql.DB
		sources  = make([]gorm.Dialector, 0)
		replicas = make([]gorm.Dialector, 0)
		dsn      string
		dsn1     string
	)

	for i, c := range cfg.Sources {
		dsn = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=true&loc=Local",
			c.Username,
			c.Password,
			c.Address,
			c.Db)
		if i == 0 {
			dsn1 = dsn
		}
		sources = append(sources, mysql.Open(dsn))
	}
	for _, c := range cfg.Replicas {
		dsn = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=true&loc=Local",
			c.Username,
			c.Password,
			c.Address,
			c.Db)
		replicas = append(replicas, mysql.Open(dsn))
	}

	opts = &gorm.Config{
		SkipDefaultTransaction: false, // 禁用默认事务(true: Error 1295: This command is not supported in the prepared statement protocol yet)
		PrepareStmt:            false, // 创建并缓存预编译语句(true: Error 1295)
	}
	db, err = gorm.Open(mysql.Open(dsn1), opts)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  sources,
		Replicas: replicas,
		// sources/replicas load balancing policy
		Policy: dbresolver.RandomPolicy{},
	}).SetMaxIdleConns(cfg.MaxIdleConns).
		SetMaxOpenConns(cfg.MaxOpenConns).
		SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Millisecond))

	db = db.Debug()

	sqlDB, err = db.DB()
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Millisecond)
	return
}

/*
SetMaxOpenConns：设置池中与数据打开的最大连接数，默认不限制连接数量。一般来说，该值设置的越大，可以并发执行的数据库查询就越多。
SetMaxIdleConns：设置池中最大空闲连接数，默认值是2. 理论上有更多的空闲连接可以减少从头建立新连接的概率，建立连接的过程比较耗时。但是过多的空闲连接会浪费内存占用。如果一个连接空闲时间过长，它也可能变得不可用。MySQL默认会自动关闭8小时未使用的连接。
SetConnMaxIdleTime：设置池中连接在关闭之前可用空闲的最长时间，默认是不限制时间。如果设置为2小时，表示池中自上次使用以后在池中空闲了2小时的连接将标为过期被清理。
SetConnMaxLifetime：设置池中连接关闭前可以保持打开的最长时间，默认是不限制时间。
*/
