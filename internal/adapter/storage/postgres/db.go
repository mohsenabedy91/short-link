package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"path"
	"runtime"
	"short-link/internal/core/config"
	"time"
)

var dbClient *sql.DB

func InitClient(ctx context.Context, conf config.Config) error {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		conf.DB.Host, conf.DB.Port, conf.DB.Username, conf.DB.Password,
		conf.DB.Name, conf.DB.Postgres.SSLMode, conf.DB.Postgres.Timezone)

	if dbClient, err = sql.Open("postgres", dsn); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err = dbClient.PingContext(ctx); err != nil {
		return err
	}

	dbClient.SetMaxOpenConns(conf.DB.Postgres.MaxOpenConnections)
	dbClient.SetMaxIdleConns(conf.DB.Postgres.MaxIdleConnections)
	dbClient.SetConnMaxLifetime(conf.DB.Postgres.MaxLifetime * time.Minute)

	return nil
}

func Get() *sql.DB {
	return dbClient
}

func Close() error {
	if cErr := dbClient.Close(); cErr != nil {
		return cErr
	}

	return nil
}

func getMigrateInstance(ctx context.Context) (*migrate.Migrate, error) {
	conf := config.GetConfig()

	// Initialize the database client
	if err := InitClient(ctx, conf); err != nil {
		return nil, err
	}

	// Get the database client return *sql.DB
	db := Get()

	// Create a new postgres driver
	dbDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	_, filename, _, _ := runtime.Caller(0)
	sourceURL := "file://" + path.Join(path.Dir(filename), "migrations")

	// Create a new migrate instance
	return migrate.NewWithDatabaseInstance(sourceURL, conf.DB.Name, dbDriver)
}

func RunMigrations() error {
	// Get the migration instance
	instance, err := getMigrateInstance(context.Background())
	if err != nil {
		return err
	}

	// Run all up migrations
	if err = instance.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

func RunDownMigration(step int) error {
	// Get the migration instance
	instance, err := getMigrateInstance(context.Background())
	if err != nil {
		return err
	}
	if step > 0 {
		step = -step
	}
	// Run down migration to revert the last migration
	if err = instance.Steps(step); err != nil {
		return err
	}

	return nil
}
