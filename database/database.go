package database

import (
	"dumpro/calculate/domain"
	"dumpro/utils"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(config utils.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.Host, config.User, config.Password, config.Db, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(domain.CalculationHistory{})
	if err != nil {
		return nil, err
	}
	return db, err
}
