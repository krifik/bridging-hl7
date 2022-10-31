package config

import (
	"context"
	"log"
	"mangojek-backend/exception"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		// Logger: logger.Default.LogMode(logger.Silent),

		// skip default transaction
		// SkipDefaultTransaction: true,
	})
	exception.PanicIfNeeded(err)
	sqlDB, err := db.DB()

	sqlDB.SetMaxOpenConns(postgresPoolMax)
	sqlDB.SetConnMaxIdleTime(time.Duration(postgresMaxIdleTime) * time.Second)
	sqlDB.SetMaxIdleConns(postgresPoolMin)
	sqlDB.SetConnMaxLifetime(time.Duration(postgresMaxIdleTime) * time.Second)

	return db
}

func NewRunMigration(db *gorm.DB) {
	for _, entity := range RegisterEntities() {
		err := db.Debug().AutoMigrate(entity.Entity)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func NewRunSeed(db *gorm.DB) error {
	// product := db.Create(faker.ProductFaker(db))
	// pp.Print(product)
	for _, entity := range RegisterSeeder(db) {
		err := db.Debug().Create(entity.Seeder).Error
		if err != nil {
			return err
		}

	}
	return nil
}

func NewPostgresContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
