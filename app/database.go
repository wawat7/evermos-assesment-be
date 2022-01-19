package app

import (
	"evermos-assessment-be/helper"
	"evermos-assessment-be/service/order"
	"evermos-assessment-be/service/product"
	"evermos-assessment-be/service/user"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(config Config) *gorm.DB {

	DB_USERNAME := config.Get("DB_USERNAME")
	DB_PASSWORD := config.Get("DB_PASSWORD")
	DB_HOST := config.Get("DB_HOST")
	DB_PORT := config.Get("DB_PORT")
	DB_NAME := config.Get("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&order.Order{})
	db.AutoMigrate(&order.OrderHistory{})
	db.AutoMigrate(&order.OrderProduct{})

	return db
}
