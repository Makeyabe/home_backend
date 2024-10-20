package main

import (
	"fmt"
	"log"

	"github.com/Makeyabe/Home_Backend/initializers"
	"github.com/Makeyabe/Home_Backend/model"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&model.Student{})
	initializers.DB.AutoMigrate(&model.Teacher{})
	initializers.DB.AutoMigrate(&model.Booking{})
	initializers.DB.AutoMigrate(&model.Summary{})
	initializers.DB.AutoMigrate(&model.Formvisit{})
	fmt.Println("? Migration complete")
}
