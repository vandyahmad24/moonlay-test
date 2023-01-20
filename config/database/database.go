package database

import (
	"fmt"
	"gorm.io/driver/postgres"

	"github.com/vandyahmad24/moonlay-test/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var e error

func InitDbMysql() {

	fmt.Println("Trying to connect database :" + config.GetEnvVariable("DB_NAME"))
	fmt.Println("Trying to connect MYSQL_HOST :" + config.GetEnvVariable("DB_HOST"))
	fmt.Println("MYSQL")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetEnvVariable("DB_USER"),
		config.GetEnvVariable("DB_PASSWORD"),
		config.GetEnvVariable("DB_HOST"),
		config.GetEnvVariable("DB_PORT"),
		config.GetEnvVariable("DB_NAME"),
	)
	DB, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if e != nil {
		panic(e)
	}

}

func InitDbPsql() {

	fmt.Println("Trying to connect database :" + config.GetEnvVariable("DB_NAME"))
	fmt.Println("Trying to connect MYSQL_HOST :" + config.GetEnvVariable("DB_HOST"))
	fmt.Println("POSTGRESQL")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.GetEnvVariable("DB_HOST"),
		config.GetEnvVariable("DB_USER"),
		config.GetEnvVariable("DB_PASSWORD"),
		config.GetEnvVariable("DB_NAME"),
		config.GetEnvVariable("DB_PORT"))
	DB, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if e != nil {
		panic(e)
	}

}
