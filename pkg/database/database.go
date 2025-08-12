package database

import (
	"fmt"
	"sync"

	"gorm.io/gorm"
)

type DbInstance struct {
	DB *gorm.DB
	mu sync.RWMutex
}

var (
	dbInstance *DbInstance
	once       sync.Once
)

func GetDatabaseInstance() *DbInstance {
	once.Do(func() {
		dbInstance = &DbInstance{}
	})
	return dbInstance
}

func (d *DbInstance) SetDB(gormDb *gorm.DB) {
	d.mu.Lock()
	defer d.mu.Unlock()
	// if d.DB != nil {
	// 	d.DB = nil // Reset the DB if it was already set
	// }
	d.DB = gormDb
}

func (d *DbInstance) GetDB() *gorm.DB {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.DB
}

func InitializeDatabase() {
	configBuilder := NewDbConfigBuilder()
	config, err := configBuilder.SetHost("localhost").
		SetPort(5432).
		SetCredentials("postgres", "password").
		SetDatabase("postgres").
		SetSSL(false).
		Build()

	if err != nil {
		fmt.Println("Error building database config:", err)
		return
	}
	fmt.Printf("Database configuration built successfully: %v\n", config)

	dbConn, err := GetDatabaseConnectionFromFactory(Postgres, config)
	if err != nil {
		fmt.Println("Error getting database connection from factory:", err)
		return
	}

	dbInstance := GetDatabaseInstance()
	dbInstance.SetDB(dbConn)
}
