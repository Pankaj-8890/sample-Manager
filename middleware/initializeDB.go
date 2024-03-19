package middleware

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"sampleManager/model"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


func DatabaseConnection() *gorm.DB {
	host := "localhost"
	port := "5432"
	dbName := "sampleManager"
	dbUser := "postgres"
	password := "postgres"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, dbUser, dbName, password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info),})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")

	db.Exec("DROP TABLE IF EXISTS orders;")
	db.Exec("DROP TABLE IF EXISTS delivery_people;")
	err = db.AutoMigrate(&model.Mapping{}, &model.Segment{})
	if err != nil {
		log.Fatalf("Error migrating Mapping table: %v", err)
	}

	return db
}