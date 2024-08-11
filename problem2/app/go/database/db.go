package database

import (
	"fmt"
	"problem1/configs"
	"problem1/model"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
	once.Do(func() {
		Gdb, err := gorm.Open(mysql.Open(configs.Get().DB.DataSource), &gorm.Config{PrepareStmt: true})
		if err != nil {
			fmt.Println(err)
		}
		Gdb.AutoMigrate(&model.User{}, &model.FriendLink{})
		db = Gdb

		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println("failed to get database object: ?", err)
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	})
	return db
}

func Close() {
	db, err := db.DB()
	if err != nil {
		fmt.Println(err)
	}
	db.Close()
}
