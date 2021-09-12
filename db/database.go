package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "host=ec2-54-74-14-109.eu-west-1.compute.amazonaws.com " +
		"user=pmfbtvnvqzyukr " +
		"password=e629f697b700dcaf54c9f1bf4c6ffe1318f2b8f5771461e46764238a160cf8c1 " +
		"dbname=dd73bum723rkng port=5432 TimeZone=Europe/Moscow"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
