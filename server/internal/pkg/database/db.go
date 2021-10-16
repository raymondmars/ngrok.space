package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"raymond.com/ngrok-server/internal/pkg/models"
	"raymond.com/ngrok-server/internal/pkg/utils"
)

var Db *gorm.DB

//create database if not exists ngrok character set utf8mb4 collate utf8mb4_bin;

func InstallDb() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/ngrok?charset=utf8mb4&parseTime=True&loc=Local",
		utils.GetEvnWithDefaultVal("MYSQL_USER", "root"), utils.GetEvnWithDefaultVal("MYSQL_PASSWORD", "111111"), utils.GetEvnWithDefaultVal("MYSQL_ADDR", "localhost:3306"))

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Warn),
	})

	if err != nil {
		panic(err)
	}
	setDb, _ := Db.DB()

	setDb.SetMaxIdleConns(10)
	setDb.SetMaxOpenConns(100)

	// Add table suffix when create tables
	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.User{},
		&models.Domain{},
	)

}
