package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"gowith/config"
	"gowith/logging"
)

var DB *gorm.DB

type Model struct {
	//ID int `gorm:"primary_key" json:"id"`
	//CreatedOn int `json:"gmt_create"`
	//ModifiedOn int `json:"gmt_modified"`
}

func init() {
	var err error
	DB, err = gorm.Open(config.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.DBName))
	if err != nil{
		logging.Fatal(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}

	DB.SingularTable(true) //使用单数表名
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	fmt.Println("connect to database successly!")
}

func CloseDB() {
	//defer DB.Close()
}
















