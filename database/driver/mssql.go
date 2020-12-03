package driver

import (
	"github.com/pangxianfei/framework/config"
)

type mssql struct {
	conn string
}

func NewMssql(connection string) *mssql {
	_db := new(mssql)
	_db.setConnection(connection)
	return _db
}

func (_mys *mssql) setConnection(connection string) {
	_mys.conn = connection
}

func (_mys *mssql) connection() string {
	return _mys.conn
}

func (_mys *mssql) username() string {
	return _mys.config("username")
}
func (_mys *mssql) password() string {
	return _mys.config("password")
}
func (_mys *mssql) host() string {
	return _mys.config("host")
}
func (_mys *mssql) port() string {
	return _mys.config("port")
}
func (_mys *mssql) database() string {
	return _mys.config("database")
}
func (_mys *mssql) charset() string {
	return _mys.config("charset")
}
func (_mys *mssql) Prefix() string {
	return _mys.config("prefix")
}

func (_mys *mssql) Driver() string {
	return _mys.config("driver")
}
func (_mys *mssql) collation() string {
	return _mys.config("collation")
}
func (_mys *mssql) config(key string) string {
	value := config.GetString("database." + _mys.connection() + "." + key)
	if value == "" {
		panic("database " + key + " parse error")
	}
	return value
}
func (_mys *mssql) ConnectionArgs() string {
	return "sqlserver://" +_mys.username() + ":" + _mys.password() + "@" + _mys.host() + ":" + _mys.port() + "?database=" + _mys.database()
}
