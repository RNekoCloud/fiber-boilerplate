package config

import (
	"fmt"

	"api-service/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

func NewDBConfig(source *DBConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",

		source.Host,
		source.User,
		source.Password,
		source.DBName,
		source.Port,
		source.SSLMode,
	)

	logrus.Printf("[config][func: NewDBConfig] DB DSN: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Warnln("[database] Failed to connect Postgres:", err)
		return nil
	}

	db.AutoMigrate(&model.Course{})

	return db

}
