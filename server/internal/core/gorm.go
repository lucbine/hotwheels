package core

import (
	"fmt"
	"hotwheels/server/internal/config"
	"time"

	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	//  todo 使用sync.Once  参考qms 框架
	DB *gorm.DB
)

//func DB(tableName string) {
//
//}

func InitGDB() (err error) {
	if DB, err = loadDB("hotwheels"); err != nil {
		fmt.Printf("db hotwheels init failed" + err.Error())
		return err
	}

	fmt.Println("db init sucessful")
	return nil
}

func loadDB(key string) (db *gorm.DB, err error) {
	keyPrefix := fmt.Sprintf("database.%s", key)
	host := config.GetString(keyPrefix + ".host")
	port := config.GetString(keyPrefix + ".port")
	dbname := config.GetString(keyPrefix + ".dbname")
	user := config.GetString(keyPrefix + ".username")
	password := config.GetString(keyPrefix + ".password")
	charset := config.GetString(keyPrefix + ".charset")
	readTimeout := config.GetString(keyPrefix + ".read_timeout")
	writeTimeout := config.GetString(keyPrefix + ".write_timeout")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&readTimeout=%s&writeTimeout=%s&parseTime=True&loc=Local", user, password, host, port,
		dbname, charset, readTimeout, writeTimeout)
	if db, err = gorm.Open("mysql", dsn); err != nil {

		fmt.Println("--------", dsn)
		fmt.Println("--------", err.Error())

		return nil, errors.Wrap(err, "db connect error")
	}
	db.DB().SetMaxOpenConns(255)
	db.DB().SetMaxIdleConns(32)
	db.DB().SetConnMaxLifetime(600 * time.Second)
	if config.Env() != "prd" {
		db.LogMode(true)
	}
	return db, nil
}
