package database

import "github.com/jinzhu/gorm"

// database connection
var db *gorm.DB

// connection config
type ConnectionConfig struct {
	//dsn         string
	TablePrefix string
	Debug       bool
	Lifetime    int
	OpenConn    int
}

var Config ConnectionConfig

func NewConn() (*gorm.DB, error) {
	return NewMysqlConn()
}
