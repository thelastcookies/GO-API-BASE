package model

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	mysqlG "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
	config2 "tlc.platform/web-service/config"
	"tlc.platform/web-service/pkg/config"
)

var Db *sql.DB

var GDB *gorm.DB

var Conf *config2.MysqlConfig

func InitMySQL() {
	c := config.New("config/local")
	if err := c.Load("database", "yaml", &Conf); err != nil {
		panic(err)
	}
	cfg := mysql.Config{
		User:                 Conf.UserName,
		Passwd:               Conf.Password,
		Addr:                 Conf.Addr,
		DBName:               Conf.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
		Loc:                  time.Local,
	}
	var err error
	// 连接 MySQL
	Db, err = sql.Open("mysql", cfg.FormatDSN())
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
