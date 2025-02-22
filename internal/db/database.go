package database

import (
	"VariableServergRPC/config"
	"VariableServergRPC/internal/logging"
	"VariableServergRPC/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	configs, err := config.ConfigSet()
	if err != nil {
		logger.Log.Error("Fail to load configurations: %v", err)
		return nil, err
	}

	connString := configs.DataBase.StringConnection
	typeDatabase := configs.DataBase.TypeDatabase

	var db *gorm.DB
	switch typeDatabase {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(connString), &gorm.Config{})
	default:
		logger.Log.Error("Database not supported: %s", typeDatabase)
		return nil, err
	}

	if err != nil {
		logger.Log.Error("Fail to connect to database: %v", err)
		return nil, err
	}

	if err := db.AutoMigrate(&model.Variable{}); err != nil {
		logger.Log.Error("Error running migration: ", err)
	}

	logger.Log.Info("Connection with database successful!")

	return db, nil
}
