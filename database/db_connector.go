package database

import (
	gormv1 "github.com/jinzhu/gorm"
	mysql "gorm.io/driver/mysql"
	//gormv2 "gorm.io/gorm"
	"gorm.io/gorm"
	"database/sql"
	"github.com/pangxianfei/framework/config"
	"github.com/pangxianfei/framework/database/driver"
	"github.com/pangxianfei/framework/helpers/zone"
	"time"
)

//var dbv1 *gormv1.DB
var db *gorm.DB
var dber databaser
/***
 pangxianfeu by add
 ***/
func Initialize() {
	dber, db = setv2Connection("default")
	//dber, db = setConnection("default")
	//dber, db = setConnection("default")
	//dber, dbv1 = setConnection("default")
	//configOrm(dber)
}

func setConnection(conn string) (dber databaser, db *gormv1.DB) {


	if conn == "default" {
		conn = config.GetString("database." + conn)
		if conn == "" {
			panic("database connection parse error")
		}
	}

	// get driver instance
	switch conn {
	case "mysql":
		dber = driver.NewMysql(conn)
		break
	default:
		panic("incorrect database connection provided")
	}

	// connect database
	db, err := gormv1.Open(conn, dber.ConnectionArgs())
	if err != nil {
		panic("failed to connect database")
	}

	err = db.DB().Ping()
	if err != nil {
		panic("failed to connect database by ping")
	}
	if config.GetBool("app.debug") {
		db = db.Debug().LogMode(true)
	}

	db.DB().SetMaxIdleConns(config.GetInt("database.max_idle_connections"))
	db.DB().SetMaxOpenConns(config.GetInt("database.max_open_connections"))
	db.DB().SetConnMaxLifetime(zone.Duration(config.GetInt("database.max_life_seconds")) * zone.Second)

	//defer _db.Close()
	return dber, db
}



/**
  pangxianfei by add
 */
func setv2Connection(conn string) (dber databaser, sqlDb *gorm.DB) {


	if conn == "default" {
		conn = config.GetString("database." + conn)
		if conn == "" {
			panic("database connection parse error")
		}
	}

	// get driver instance
	switch conn {
	case "mysql":
		dber = driver.NewMysql(conn)
		break
	default:
		panic("incorrect database connection provided")
	}


	Db, err := sql.Open("mysql", dber.ConnectionArgs()) //sql.Open("mysql", "mydb_dsn")
	sqlDb, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dber.ConnectionArgs(),
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
		Conn:                      Db,
	}), &gorm.Config{})

	err = Db.Ping()

	if err != nil {
		panic("failed to connect database by ping")
	}

	Db.SetConnMaxLifetime(time.Hour)
	Db.SetMaxIdleConns(config.GetInt("database.max_idle_connections"))
	Db.SetMaxOpenConns(config.GetInt("database.max_open_connections"))
	Db.SetConnMaxLifetime(zone.Duration(config.GetInt("database.max_life_seconds")) * zone.Second)

/*
	Db, err := sql.Open("mysql", dber.ConnectionArgs()) //sql.Open("mysql", "mydb_dsn")
	sqlDb, err = gorm.Open(mysql.New(mysql.Config{
		Conn: Db,
	}), &gorm.Config{})

	Db.SetMaxIdleConns(10)


	Db.SetMaxOpenConns(100)

*/

	if err != nil {
		panic("failed to connect database")
	}

	return dber, sqlDb
}



func Connection(conn string) (_db *gorm.DB) {
	_, _db = setv2Connection(conn)
	return _db
}

func DB() *gorm.DB {
	return db
}

func Prefix() string {
	return dber.Prefix()
}
