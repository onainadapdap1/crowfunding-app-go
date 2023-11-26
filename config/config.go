package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load file")
	}
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func ConnectToDB() *gorm.DB {
	var dbConfig DBConfig = DBConfig{
		Username: os.Getenv("KATO_DBUSER"),
		Password: os.Getenv("KATO_DBPASSWORD"),
		Host:     os.Getenv("KATO_DBHOST"),
		Port:     os.Getenv("KATO_DBPORT"),
		Name:     os.Getenv("KATO_DBNAME"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name)

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database Connection Error")
	}

	return DB
}
