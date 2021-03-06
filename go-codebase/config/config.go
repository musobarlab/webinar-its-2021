package config

import (
	"context"
	"crypto/rsa"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gitlab.com/Wuriyanto/go-codebase/config/key"
	"gitlab.com/Wuriyanto/go-codebase/pkg/database"
	"gorm.io/gorm"
)

// Config app
type Config struct {
	ReadDB, WriteDB *gorm.DB
	PrivateKey      *rsa.PrivateKey
	PublicKey       *rsa.PublicKey
}

// Env model
type Env struct {
	RootApp string

	// Profile application
	Profile string

	// Development env checking, this env for debug purpose
	Development string

	// HTTPPort config
	HTTPPort uint16
	// GRPCPort Config
	GRPCPort uint16

	// BasicAuthUsername config
	BasicAuthUsername string
	// BasicAuthPassword config
	BasicAuthPassword string

	// GRPC auth key
	GRPCAuthKey string

	// CacheExpired config
	CacheExpired time.Duration

	ReadDB struct {
		Host     string
		User     string
		Password string
		DBName   string
		Port     uint16
		SSLMode  string
	}

	WriteDB struct {
		Host     string
		User     string
		Password string
		DBName   string
		Port     uint16
		SSLMode  string
	}

	AccessTokenExpired  time.Duration
	RefreshTokenExpired time.Duration
}

// GlobalEnv global environment
var GlobalEnv Env

// Init app config
func Init(ctx context.Context, rootApp string) *Config {
	loadEnv(rootApp)

	cfgChan := make(chan *Config)
	go func() {
		defer close(cfgChan)

		var (
			cfg Config
			err error
		)
		err = database.GormDBReadInit()
		if err != nil {
			panic(err)
		}

		err = database.GormDBWriteInit()
		if err != nil {
			panic(err)
		}

		cfg.ReadDB, err = database.GetGormReadDB()
		if err != nil {
			panic(err)
		}

		cfg.WriteDB, err = database.GetGormWriteDB()
		if err != nil {
			panic(err)
		}
		cfg.PrivateKey = key.LoadPrivateKey()
		cfg.PublicKey = key.LoadPublicKey()

		cfgChan <- &cfg
	}()

	// with timeout to init configuration
	select {
	case cfg := <-cfgChan:
		return cfg
	case <-ctx.Done():
		panic(fmt.Errorf("Failed to init configuration: %v", ctx.Err()))
	}
}

func loadEnv(rootApp string) {
	// load .env
	err := godotenv.Load(rootApp + "/.env")
	if err != nil {
		log.Println(err)
	}

	os.Setenv("APP_PATH", rootApp)
	GlobalEnv.RootApp = rootApp

	var ok bool

	GlobalEnv.Profile, ok = os.LookupEnv("PROFILE")
	if !ok {
		panic("missing PROFILE environment")
	}

	GlobalEnv.Development, ok = os.LookupEnv("DEVELOPMENT")
	if !ok {
		panic("missing DEVELOPMENT environment")
	}

	// ------------------------------------

	if port, err := strconv.Atoi(os.Getenv("HTTP_PORT")); err != nil {
		panic("missing HTTP_PORT environment")
	} else {
		GlobalEnv.HTTPPort = uint16(port)
	}

	if port, err := strconv.Atoi(os.Getenv("GRPC_PORT")); err != nil {
		panic("missing GRPC_PORT environment")
	} else {
		GlobalEnv.GRPCPort = uint16(port)
	}

	// ------------------------------------
	GlobalEnv.BasicAuthUsername, ok = os.LookupEnv("BASIC_AUTH_USER")
	if !ok {
		panic("missing BASIC_AUTH_USER environment")
	}
	GlobalEnv.BasicAuthPassword, ok = os.LookupEnv("BASIC_AUTH_PASS")
	if !ok {
		panic("missing BASIC_AUTH_PASS environment")
	}

	GlobalEnv.GRPCAuthKey, ok = os.LookupEnv("GRPC_AUTH_KEY")
	if !ok {
		panic("missing GRPC_AUTH_KEY environment")
	}

	accessTokenExpired, ok := os.LookupEnv("ACCESS_TOKEN_EXPIRED")
	if !ok {
		panic("missing ACCESS_TOKEN_EXPIRED environment")
	}

	refreshTokenExpired, ok := os.LookupEnv("REFRESH_TOKEN_EXPIRED")
	if !ok {
		panic("missing REFRESH_TOKEN_EXPIRED environment")
	}

	GlobalEnv.AccessTokenExpired, err = time.ParseDuration(accessTokenExpired)
	if err != nil {
		panic("invalid access token string")
	}

	GlobalEnv.RefreshTokenExpired, err = time.ParseDuration(refreshTokenExpired)
	if err != nil {
		panic("invalid refresh token string")
	}

	// -------------------------------------------------------
	GlobalEnv.ReadDB.Host, ok = os.LookupEnv("READ_DB_HOST")
	if !ok {
		panic("missing READ_DB_HOST environment")
	}

	GlobalEnv.ReadDB.User, ok = os.LookupEnv("READ_DB_USER")
	if !ok {
		panic("missing READ_DB_USER environment")
	}

	GlobalEnv.ReadDB.Password, ok = os.LookupEnv("READ_DB_PASSWORD")
	if !ok {
		panic("missing READ_DB_PASSWORD environment")
	}

	GlobalEnv.ReadDB.DBName, ok = os.LookupEnv("READ_DB_NAME")
	if !ok {
		panic("missing READ_DB_NAME environment")
	}

	if readDBPort, err := strconv.Atoi(os.Getenv("READ_DB_PORT")); err != nil {
		panic("missing READ_DB_PORT environment")
	} else {
		GlobalEnv.ReadDB.Port = uint16(readDBPort)
	}

	GlobalEnv.ReadDB.SSLMode, ok = os.LookupEnv("READ_DB_SSLMODE")
	if !ok {
		panic("missing READ_DB_SSLMODE environment")
	}

	// ----------------------------------------------------

	GlobalEnv.WriteDB.Host, ok = os.LookupEnv("WRITE_DB_HOST")
	if !ok {
		panic("missing WRITE_DB_HOST environment")
	}

	GlobalEnv.WriteDB.User, ok = os.LookupEnv("WRITE_DB_USER")
	if !ok {
		panic("missing WRITE_DB_USER environment")
	}

	GlobalEnv.WriteDB.Password, ok = os.LookupEnv("WRITE_DB_PASSWORD")
	if !ok {
		panic("missing WRITE_DB_PASSWORD environment")
	}

	GlobalEnv.WriteDB.DBName, ok = os.LookupEnv("WRITE_DB_NAME")
	if !ok {
		panic("missing WRITE_DB_NAME environment")
	}

	if writeDBPort, err := strconv.Atoi(os.Getenv("WRITE_DB_PORT")); err != nil {
		panic("missing WRITE_DB_PORT environment")
	} else {
		GlobalEnv.WriteDB.Port = uint16(writeDBPort)
	}

	GlobalEnv.WriteDB.SSLMode, ok = os.LookupEnv("WRITE_DB_SSLMODE")
	if !ok {
		panic("missing WRITE_DB_SSLMODE environment")
	}
}

// Exit release all connection, think as deferred function in main
func (c *Config) Exit(ctx context.Context) {
	// close mongo session
	// clean up the connection here. Eg: database connection
	writeDB, err := c.WriteDB.DB()
	if err != nil {
		panic(err)
	}

	writeDB.Close()

	readDB, err := c.ReadDB.DB()
	if err != nil {
		panic(err)
	}

	readDB.Close()

	log.Println("\x1b[33;1mConfig: Success close all connection\x1b[0m")
}
