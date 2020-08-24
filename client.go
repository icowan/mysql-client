/**
 * @Time : 2020/8/24 10:45 AM
 * @Author : solacowa@gmail.com
 * @File : client
 * @Software: GoLand
 */

package mysqlclient

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewMysql(dbUrl string, debug bool) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
	db.LogMode(debug)

	if err = db.DB().Ping(); err != nil {
		log.Println("db", "ping", "err", err)
	}

	log.Println("mysql", "connect", "success", true)

	return db, nil
}
