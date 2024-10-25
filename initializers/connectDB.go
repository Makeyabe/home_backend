package initializers

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // Output to stdout
        logger.Config{
            SlowThreshold: time.Second,   // Log queries slower than this
            LogLevel:      logger.Info,   // Set to Info for full SQL logging
            Colorful:      true,          // Enable colorful printing
        },
    )

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Printf("? Connected Successfully to the Database %s \n", config.DBName)
}

