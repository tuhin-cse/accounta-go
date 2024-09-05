package config

import (
	"accounta-go/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB
var err error

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBName:   os.Getenv("MYSQL_DB"),
	}
	return &dbConfig
}

func ConnectDB() error {
	dbConfig := BuildDBConfig()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

func GetDB() *gorm.DB {
	return db
}

func MigrateDB() {
	err := db.AutoMigrate(&models.User{}, &models.Category{})
	if err != nil {
		return
	}
}
