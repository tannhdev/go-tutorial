package controllers

import (
	"context"
	"ecommerce-project/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Hash Password
func HashPassword(password string) string {

}

// Verify Password
func VerifyPassword(userPassword string, givenPassword string) (bool, string) {

}

// Signup

func Signup() gin.HandlerFunc {
	
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout((context.Background(), 100*time.Second))

		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H("error": err.Error()))
			return
		}

		validationErr := validate.Struct(user)
		if validationErr!= nil {
            c.JSON(http.StatusBadRequest, gin.H("error": validationErr.Error()))
            return
        }

		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, gin.H("error": err.Error))
			return
		}

		if count > 0 {
            c.JSON(http.StatusBadRequest, gin.H("error": "User already exists"))
            return
        }

		count, err := UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})

		defer cancel()

		if err != nil {
			log.Panic(err)
            c.JSON(http.StatusInternalServerError, gin.H("error": err.Error))
            return
		}

		if count > 0 {
            c.JSON(http.StatusBadRequest, gin.H("error": "This phone numer already exist"))
            return
        }

		password := HashPassword(*user.Password)

		user.Password = &password
        user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.ObjectID
		user.User_ID = user.ID.Hex()

		token, refreshtoken, _ := generate.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, user.User_ID) 

		user.Token = &token
		user.Refresh_Token = &refreshtoken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)

		_, inserterr := UserCollection.InsertOne(ctx, user)
		
		if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "The user did not get created"})
			return
		}

		defer cancel()

		c.JSON(http.StatusCreated, "Sign Up successful")
	}
}

// Login
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.M{"error": err})
			return
		}

		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)

		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.M{"error": "Login or password incorrect"})
			return
		}

		IsValidPassword, msg := VerifyPassword(*user.Password, *founduser.Password)

		defer cancel()

		if !IsValidPassword {
			c.JSON(http.StatusInternalServerError, gin.M{"error": msg})
			fmt.Println((msg))
			return
		}

		token, refreshToken, _ := generate.TokenGenerator(*founduser.Email, *founduser.First_Name, *founduser.Last_Name, *founduser.User_ID)

		defer cancel()

		generate.UpdateAllTokens(token, refreshToken, founduser.User_ID)

		c.JSON(http.StatusFound, founduser)
	}
}

// Product Viewer Admin
func ProductViewerAdmin() gin.HandlerFunc {

}

// SearchProduct
func SearchProduct() gin.HandlerFunc {

}

// SearchProductByQuery
func SearchProductByQuery() gin.HandlerFunc {

}
