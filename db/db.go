package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB{
	return db
}

func init(){
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=weather port=5432 sslmode=disable "
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}

}
// implement database connection pool
