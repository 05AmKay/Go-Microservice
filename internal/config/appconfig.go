package config

import (
	"os"
	"sync"

	"example.com/api/pkg/helpers"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSL      bool
}

type AppConfig struct {
	AppEnv   string   `json:"env"`
	Port     string   `json:"port"`
	DBConfig DBConfig `json:"db_config"`
}

type configInstance struct {
	AppConfig *AppConfig
	mu        sync.RWMutex
}

var (
	instance *configInstance
	once     sync.Once
)

func GetConfig() *configInstance {
	return instance
}

func LoadAppConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	dbPort, err := helpers.ConvertStringToInt(os.Getenv("DB_PORT"))
	if err != nil {
		return err
	}

	once.Do(func() {
		instance = &configInstance{
			AppConfig: &AppConfig{
				AppEnv: os.Getenv("APP_ENV"),
				Port:   os.Getenv("PORT"),
				DBConfig: DBConfig{
					Host:     os.Getenv("DB_HOST"),
					Port:     dbPort,
					User:     os.Getenv("DB_USER"),
					Password: os.Getenv("DB_PASSWORD"),
					Database: os.Getenv("DB_NAME"),
					SSL:      os.Getenv("DB_SSLMODE") == "true",
				},
			}}

	})

	return nil
}
