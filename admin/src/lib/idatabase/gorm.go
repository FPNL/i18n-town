package idatabase

import (
	"fmt"
	"os"

	"github.com/FPNL/admin/src/core/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupGorm() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("IADMIN_PSQL_HOST"),
		os.Getenv("IADMIN_PSQL_USERNAME"),
		os.Getenv("IADMIN_PSQL_PASSWORD"),
		os.Getenv("IADMIN_PSQL_PORT"),
		os.Getenv("IADMIN_PSQL_DB_NAME"),
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
	return db.AutoMigrate(&entity.User{}, &entity.Organize{})
}
