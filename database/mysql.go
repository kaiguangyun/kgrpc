package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kaiguangyun/kgrpc/helper"
	"strconv"
	"strings"
	"time"
)

func init() {
	Config.TablePrefix = helper.GetEnv("MysqlPrefix")
	Config.Debug = strings.ToLower(helper.GetEnv("MysqlDebug")) == "true"
	Config.Lifetime, _ = strconv.Atoi(helper.GetEnv("MysqlConnectionMaxLifetime"))
	Config.OpenConn, _ = strconv.Atoi(helper.GetEnv("MysqlMaxOpenConnections"))
}

// New Mysql Connection
func NewMysqlConn() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}
	gormDB, err := gorm.Open("mysql", getMysqlDataSource())
	if err != nil {
		return nil, err
	}
	setMysqlTablePrefix()
	setMysqlDebug(gormDB)
	setMysqlMaxOpenConnections(gormDB)
	setMysqlConnectionMaxLifetime(gormDB)

	return gormDB, err
}

// default 30 second
func setMysqlConnectionMaxLifetime(gormDB *gorm.DB) {
	if Config.Lifetime > 0 {
		gormDB.DB().SetConnMaxLifetime(time.Duration(Config.Lifetime) * time.Second)
	} else {
		gormDB.DB().SetConnMaxLifetime(30 * time.Second)
	}
}

// max open connections
func setMysqlMaxOpenConnections(gormDB *gorm.DB) {
	if Config.OpenConn > 0 {
		gormDB.DB().SetMaxOpenConns(Config.OpenConn)
	}
}

// setMysqlDebug
func setMysqlDebug(gormDB *gorm.DB) {
	gormDB.LogMode(Config.Debug)
}

// set table prefix
func setMysqlTablePrefix() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		if !strings.HasPrefix(defaultTableName, Config.TablePrefix) {
			return Config.TablePrefix + defaultTableName
		}
		return defaultTableName
	}
}

// dbConfig := mysql.Config{}, dbConfig.FormatDSN()
func getMysqlDataSource() string {
	dataSourceName := ""
	dataSourceName += helper.GetEnv("MysqlUser") + ":"
	dataSourceName += helper.GetEnv("MysqlPassword") + "@"
	dataSourceName += "tcp(" + helper.GetEnv("MysqlHost") + ":" + helper.GetEnv("MysqlPort") + ")/"
	dataSourceName += helper.GetEnv("MysqlName") + "?"
	dataSourceName += helper.GetEnv("MysqlParameters")

	return dataSourceName
}
