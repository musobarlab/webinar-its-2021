package database

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gitlab.com/Wuriyanto/go-codebase/pkg/helper"
)

var (
	readDB, writeDB *gorm.DB
)

// GormDBReadInit function will initialize READ database connection
func GormDBReadInit() error {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("READ_DB_HOST"),
		os.Getenv("READ_DB_USER"),
		os.Getenv("READ_DB_PASSWORD"),
		os.Getenv("READ_DB_NAME"),
		os.Getenv("READ_DB_PORT"),
		os.Getenv("READ_DB_SSLMODE"),
		helper.TimeZoneAsia)

	var err error
	readDB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	return nil
}

// GormDBWriteInit function will initialize WRITE database connection
func GormDBWriteInit() error {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("WRITE_DB_HOST"),
		os.Getenv("WRITE_DB_USER"),
		os.Getenv("WRITE_DB_PASSWORD"),
		os.Getenv("WRITE_DB_NAME"),
		os.Getenv("WRITE_DB_PORT"),
		os.Getenv("WRITE_DB_SSLMODE"),
		helper.TimeZoneAsia)

	var err error
	writeDB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	return nil
}

// GetGormReadDB function will return gorm.DB instance
func GetGormReadDB() (*gorm.DB, error) {
	if readDB == nil {
		return nil, errors.New("write db not initialize yet")
	}

	return readDB, nil
}

// GetGormWriteDB function will return gorm.DB instance
func GetGormWriteDB() (*gorm.DB, error) {
	if writeDB == nil {
		return nil, errors.New("write db not initialize yet")
	}

	return writeDB, nil
}
