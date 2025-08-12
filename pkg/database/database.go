package database

import (
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
