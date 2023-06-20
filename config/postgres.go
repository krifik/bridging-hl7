package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/k0kubun/pp"
	"github.com/krifik/bridging-hl7/exception"

	"github.com/jackc/pgx/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type postgresConfig struct {
	ConnConfig *pgx.ConnConfig
}

func NewPostgresDatabase(configuration Config) *gorm.DB {
	postgresPoolMin, err := strconv.Atoi(configuration.Get("POSTGRES_POOL_MIN"))
	exception.PanicIfNeeded(err)
	postgresPoolMax, err := strconv.Atoi(configuration.Get("POSTGRES_POOL_MAX"))
	exception.PanicIfNeeded(err)
	postgresMaxIdleTime, err := strconv.Atoi(configuration.Get("POSTGRES_MAX_IDLE_TIME_SECOND"))
	exception.PanicIfNeeded(err)
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// disable log mode
		Logger: logger.Default.LogMode(logger.Silent),

		// skip default transaction
		// SkipDefaultTransaction: true,
	})
	exception.PanicIfNeeded(err)
	sqlDB, err := db.DB()
	exception.PanicIfNeeded(err)
	var dbName string
	result := sqlDB.QueryRow("SELECT current_database()")
	err = result.Scan(&dbName)
	if err != nil {
		panic(err)
	}

	fmt.Println("Current database is:", dbName)
	sqlDB.SetMaxOpenConns(postgresPoolMax)
	sqlDB.SetConnMaxIdleTime(time.Duration(postgresMaxIdleTime) * time.Second)
	sqlDB.SetMaxIdleConns(postgresPoolMin)
	sqlDB.SetConnMaxLifetime(time.Duration(postgresMaxIdleTime) * time.Second)
	pp.Println("Connected to Database")
	return db
}

func NewPostgresContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
