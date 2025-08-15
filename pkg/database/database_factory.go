package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseType string

const (
	Postgres DatabaseType = "postgres"
	MySQL    DatabaseType = "mysql"
	MongoDB  DatabaseType = "mongodb"
)

type DatabaseFactory interface {
	Connect(config *DBConfig) (*gorm.DB, error)
}

type PostgresDatabaseImpl struct{}

func (p *PostgresDatabaseImpl) Connect(config *DBConfig) (*gorm.DB, error) {
	// Implementation for connecting to Postgres database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Database,
		func() string {
			if config.SSL {
				return "require"
			}
			return "disable"
		}())

	fmt.Println("Connecting to Postgres with DSN:", dsn)
	// Here you would use gorm.Open with the Postgres driver
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to Postgres: %w", err)
	}

	return dbConn, nil
}

type MongoDBDatabaseImpl struct{}

func (m *MongoDBDatabaseImpl) Connect(config *DBConfig) (*gorm.DB, error) {
	return nil, fmt.Errorf("MongoDB connection not implemented yet")
}

type MySQLDatabaseImpl struct{}

func (m *MySQLDatabaseImpl) Connect(config *DBConfig) (*gorm.DB, error) {
	return nil, fmt.Errorf("MySQL connection not implemented yet")
}

func GetDatabaseConnectionFromFactory(dbType DatabaseType, config *DBConfig) (*gorm.DB, error) {
	switch dbType {
	case Postgres:
		return (&PostgresDatabaseImpl{}).Connect(config)
	case MongoDB:
		return (&MongoDBDatabaseImpl{}).Connect(config)
	case MySQL:
		return (&MySQLDatabaseImpl{}).Connect(config)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}

}
