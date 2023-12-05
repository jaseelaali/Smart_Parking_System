package database

import (
	"fmt"

	"github.com/jaseelaali/Smart_Parking_System/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlinfo := fmt.Sprintf("host=%s name=%s user=%s port=%s password=%s", cfg.DBHost, cfg.DBName, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, err := gorm.Open(postgres.Open(psqlinfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	db.AutoMigrate()
	if err != nil {
		return nil, err
	}
	return db, err
}
