package postgresql

import (
	"dumpro/domain"
	"dumpro/utils"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(config utils.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.Host, config.User, config.Password, config.Db, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(domain.CalculationHistory{})

	return db, err
}
