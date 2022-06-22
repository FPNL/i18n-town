package idatabase

import (
	"fmt"
	"github.com/FPNL/i18n-town/src/core/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func setupGorm() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("IWORD_PSQL_HOST"),
		os.Getenv("IWORD_PSQL_USERNAME"),
		os.Getenv("IWORD_PSQL_PASSWORD"),
		os.Getenv("IWORD_PSQL_PORT"),
		os.Getenv("IWORD_PSQL_DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	if err = migration(db); err != nil {
		return nil, err
	}

	return db, err
}

func migration(db *gorm.DB) error {
	return db.AutoMigrate(&entity.CommittedWord{}, &entity.StageWord{})
}
