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
	initializers.DB.AutoMigrate(&model.Student{}, &model.Teacher{}, &model.Booking{}, &model.Summary{}, &model.Formvisit{}, &model.Form{}, &model.FormSection{}, &model.FormField{}, &model.Image{})
	fmt.Println("? Migration complete")
}
