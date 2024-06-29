package controllers

import (
	"context"
	"ecommerce-project/database"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")

		if productQueryID == "" {
			log.Println(("Product id is empty"))

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product ID is empty"))
			return
		}

		userQueryID := c.Query("userID")

		if userQueryID == "" {
			log.Println("UserId is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("USer ID is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)

			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(200, "Successfully added to cart.")
	}
}

func (app *Application) RemoveFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")

		if productQueryID == "" {
			log.Println(("Product id is empty"))

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product ID is empty"))
			return
		}

		userQueryID := c.Query("userID")

		if userQueryID == "" {
			log.Println("UserId is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("USer ID is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)

			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.RemoveCartItem(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(200, "Successfully remove from cart.")
	}
}

func GetItomFromCart() gin.HandlerFunc {

}

func BuyFromCart() gin.HandlerFunc {

}

func InstantBuy() gin.HandlerFunc {

}
