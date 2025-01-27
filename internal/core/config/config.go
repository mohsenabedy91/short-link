package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	once   sync.Once
	config Config
)

type App struct {
	Name               string
	Env                string
	Debug              bool
	GracefullyShutdown time.Duration
	URL                string
	Port               string
	Locale             string
	PathLocale         string
}

type ShortLink struct {
	Host string
}

type SwaggerInfo struct {
	Title       string
	Description string
	Version     string
}

type Swagger struct {
	Host     string
	Schemes  string
	Info     SwaggerInfo
	Enable   bool
	Username string
	Password string
}

type DBPostgres struct {
	SSLMode            string
	MaxOpenConnections int
	MaxIdleConnections int
	MaxLifetime        time.Duration
	Timezone           string
}

type DB struct {
	Connection string
	Host       string
	Port       string
	Name       string
	Username   string
	Password   string
	Postgres   DBPostgres
}

type Config struct {
	App       App
	DB        DB
	ShortLink ShortLink
	Swagger   Swagger
}

// LoadConfig loads configuration from .env file and populates the Config struct.
func LoadConfig(envPath ...string) (Config, error) {
	err := godotenv.Load(envPath...)
	if err != nil {
		return Config{}, fmt.Errorf("error loading .env file: %v", err)
	}

	var app App
	app.Name = os.Getenv("APP_NAME")
	app.Env = os.Getenv("APP_ENV")
	app.Debug = getBoolEnv("APP_DEBUG", false)
	app.GracefullyShutdown = time.Duration(getIntEnv("APP_GRACEFULLY_SHUTDOWN", 5))
	app.URL = os.Getenv("APP_URL")
	app.Port = os.Getenv("APP_PORT")
	app.Locale = os.Getenv("APP_LOCALE")
	app.PathLocale = os.Getenv("APP_PATH_LOCALE")

	var db DB
	db.Connection = os.Getenv("DB_CONNECTION")
	db.Host = os.Getenv("DB_HOST")
	db.Port = os.Getenv("DB_PORT")
	db.Name = os.Getenv("DB_NAME")
	db.Username = os.Getenv("DB_USERNAME")
	db.Password = os.Getenv("DB_PASSWORD")
	db.Postgres.SSLMode = os.Getenv("DB_POSTGRES_SSL_MODE")
	db.Postgres.MaxOpenConnections = getIntEnv("DB_POSTGRES_MAX_OPEN_CONNECTIONS", 0)
	db.Postgres.MaxIdleConnections = getIntEnv("DB_POSTGRES_MAX_IDLE_CONNECTIONS", 0)
	db.Postgres.MaxLifetime = time.Duration(getIntEnv("DB_POSTGRES_MAX_LIFETIME", 0))
	db.Postgres.Timezone = os.Getenv("DB_POSTGRES_TIMEZONE")

	var shortLink ShortLink
	shortLink.Host = os.Getenv("SHORT_LINK_HOST")

	var swagger Swagger
	swagger.Host = os.Getenv("SWAGGER_HOST")
	swagger.Schemes = os.Getenv("SWAGGER_SCHEMES")
	swagger.Info.Title = os.Getenv("SWAGGER_INFO_TITLE")
	swagger.Info.Description = os.Getenv("SWAGGER_INFO_DESCRIPTION")
	swagger.Info.Version = os.Getenv("SWAGGER_INFO_VERSION")
	swagger.Enable = getBoolEnv("SWAGGER_ENABLE", false)
	swagger.Username = os.Getenv("SWAGGER_USERNAME")
	swagger.Password = os.Getenv("SWAGGER_PASSWORD")

	return Config{
		App:       app,
		DB:        db,
		ShortLink: shortLink,
		Swagger:   swagger,
	}, nil
}

// Helper function to convert string environment variable to bool
func getBoolEnv(key string, defaultValue bool) bool {
	val, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return defaultValue
	}
	return val
}

// Helper function to convert string environment variable to int
func getIntEnv(key string, defaultValue int) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return defaultValue
	}
	return val
}

// GetConfig loads the configuration once and returns it.
func GetConfig(envPath ...string) Config {
	once.Do(func() {
		var err error
		config, err = LoadConfig(envPath...)
		if err != nil {
			panic(err)
		}
	})
	return config
}
