package database

import (
	"fmt"
	"github.com/Ahmad-mufied/online-store/config"
	"github.com/Ahmad-mufied/online-store/model"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB connect to db
func ConnectDB() {

	fmt.Println("Connecting to Database")
	time.Sleep(3 * time.Second)
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("failed to parse database port")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"),
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	err = DB.AutoMigrate(&model.Product{}, &model.User{})
	if err != nil {
		fmt.Println("Migrate Error : ", err)
		return
	}
	fmt.Println("Database Migrated")

}
