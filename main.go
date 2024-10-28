package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Makeyabe/Home_Backend/controllers"
	"github.com/Makeyabe/Home_Backend/initializers"
	"github.com/Makeyabe/Home_Backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server                 *gin.Engine
	AuthController         controllers.AuthController
	AuthRouteController    routes.AuthRouteController
	TeacherController      *controllers.TeacherController
	TeacherRouteController routes.TeacherRouteController
	StudentController      *controllers.StudentController
	FormController         *controllers.FormController
	FormResponse           *controllers.FormResponseController
	ImageController        *controllers.ImageController
	BookingController      *controllers.BookingController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	TeacherController = controllers.NewTeacherController(initializers.DB)
	TeacherRouteController = routes.NewTeacherRouteController(TeacherController)

	StudentController = controllers.NewStudentController(initializers.DB)
	FormController = controllers.NewFormController(initializers.DB)
	FormResponse = controllers.NewFormResponseController(initializers.DB)

	ImageController = controllers.NewImageController(initializers.DB)
	BookingController = controllers.NewBookingController(initializers.DB)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin, "http://26.132.242.117:3000", "http://26.250.208.152:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Access-Control-Allow-Origin", "*"}
	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	TeacherRouteController.TeacherRoutes(router)
	routes.StudentRoutes(router, StudentController)
	routes.FormRoutes(router, FormController)
	routes.FormResponseRoutes(router, FormResponse)
	routes.SetupImageRoutes(router, ImageController) // Setup routes for images
	routes.SetupBookingRoutes(server, initializers.DB)

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Route Not Found"})
	})
	server.Static("/images", "./uploads/images")

	if _, err := os.Stat("uploads/images"); os.IsNotExist(err) {
		err := os.MkdirAll("uploads/images", os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
	}

	log.Fatal(server.Run(":" + config.ServerPort))
}
