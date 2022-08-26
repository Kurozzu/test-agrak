package libs

import (
	"fmt"
	"log"
	"sync"

	mProduct "test-agrak/app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type prod mProduct.Product

var (
	user     string = "root"
	password string = "root1234"
	host     string = "127.0.0.1"
	port     string = "3306"
	database string = "agrak"
	db       *gorm.DB
	once     sync.Once
)

//InitMysql allows open connection to database once. If the connection is already exists then it returns it.
func InitMySql() *gorm.DB {
	once.Do(func() {
		config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
		var err error
		db, err = gorm.Open(mysql.Open(config), &gorm.Config{})
		if err != nil {
			db = nil
			log.Println(err)
		}
		if db != nil {
			db.AutoMigrate(&prod{})
		}
	})
	return db
}
