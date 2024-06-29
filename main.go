package main

import (
	"ecommerce-project/controllers"
	"ecommerce-project/database"
	"ecommerce-project/middleware"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Server is running...")

	if os.Getenv("ENV") != "production" {
		// Load the .env file if not in production
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatal("Error loading .env file: ", err)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.User(gin.Logger())

	router.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuyt", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
