package persistence

import (
	"log"
	"os"
)

// PersistenceManager - persistence manager
type PersistenceManager struct {
	logger *log.Logger
	driver Driver
}

// NewPersistenceManager - create a new persistence manager
func NewPersistenceManager(config map[string]string) *PersistenceManager {
	pm := &PersistenceManager{
		logger: log.New(os.Stdout, "persistence: ", log.LstdFlags),
	}

	switch config["driver"] {
	case "flat_file":
		pm.driver = NewFlatFileDriver()
	case "sqlite":
		pm.driver = NewSQLiteDatabaseDriver(config["db_location"])
	default:
		pm.driver = NewFlatFileDriver()
	}

	return pm
}

// Start - start the persistence manager
func (pm *PersistenceManager) Start() {
	// TODO: implement
	pm.logger.Println("Starting Persistence Manager")
}

// Stop - stop the persistence manager
func (pm *PersistenceManager) Stop() {
	// TODO: implement
	pm.logger.Println("Stopping Persistence Manager")
}

// Write - write a record to disk
func (pm *PersistenceManager) Write(record KvRecord) error {
	return pm.driver.Write(record)
}

// Read - read a record from disk
func (pm *PersistenceManager) Read(key string) (KvRecord, error) {
	return pm.driver.Read(key)
}

// Delete - delete a record from disk
func (pm *PersistenceManager) Delete(key string) error {
	return pm.driver.Delete(key)
}

// Compare - compare a record to disk
func (pm *PersistenceManager) Compare(record KvRecord) (bool, error) {
	return pm.driver.Compare(record)
}

// Load - load all records from disk
func (pm *PersistenceManager) Load() ([]KvRecord, error) {
	return pm.driver.Load()
}