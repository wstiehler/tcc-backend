package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/wstiehler/tcc-backend/internal/environment"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {

	env := environment.GetInstance()

	dbConfig := DBConfig{
		Host:     env.HOST,
		Port:     env.PORT,
		User:     env.USER,
		Password: env.PASSWORD,
		DBName:   env.DB_NAME,
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {

	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	return dbUrl
}
