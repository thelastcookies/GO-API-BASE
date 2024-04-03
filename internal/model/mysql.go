package model

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	mysqlG "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	mysqlConfig "thelastcookies/api-base/config"
	"thelastcookies/api-base/pkg/config"
	"time"
)

var Db *sql.DB

var GDB *gorm.DB

var MConf *mysqlConfig.MysqlConfig

func InitMySQL() {
	c := config.New("config/local")
	if err := c.Load("database", "yaml", &MConf); err != nil {
		panic(err)
	}
	cfg := mysql.Config{
		User:                 MConf.UserName,
		Passwd:               MConf.Password,
		Addr:                 MConf.Addr,
		DBName:               MConf.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
		Loc:                  time.Local,
	}
	var err error
	// 连接 MySQL
	Db, err = sql.Open("mysql", cfg.FormatDSN())

	// 用于设置最大打开的连接数，默认值为0表示不限制。设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误
	Db.SetMaxOpenConns(MConf.MaxOpenConn)
	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	Db.SetMaxIdleConns(MConf.MaxIdleConn)
	Db.SetConnMaxLifetime(MConf.ConnMaxLifeTime)
	// 使用 gorm 接管数据库操作
	GDB, _ = gorm.Open(mysqlG.New(mysqlG.Config{
		Conn: Db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	pingErr := Db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Database connected successfully.")
}

// GetDB 返回默认的数据库
func GetDB() *gorm.DB {
	return GDB
}
