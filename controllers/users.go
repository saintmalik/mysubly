package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saintmalik/mysubly/configs"
	"github.com/saintmalik/mysubly/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userSignup *mongo.Collection = configs.GetCollection(configs.DB, "users")

func SignUp(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	var signup models.Signup

	if err := c.BindJSON(&signup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signup.Password), 8)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	signUp := models.Signup{
		Email:    signup.Email,
		Name:     signup.Name,
		Password: string(hashedPassword),
	}

	res, err := userSignup.InsertOne(ctx, signUp)

	fmt.Println("hhd", res)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "data updated successfully!"})

}

// func SignUp(c *gin.Context) {

// 	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
// 	defer cancel()
// 	var signup models.Signup

// 	if err := c.BindJSON(&signup); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": err})
// 		log.Fatal(err)
// 		return
// 	}

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signup.Password), 8)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
// 		return
// 	}

// 	signUp := models.Signup{
// 		Email:    signup.Email,
// 		Name:     signup.Name,
// 		Password: string(hashedPassword),
// 	}

// 	res, err := userSignup.InsertOne(ctx, signUp)

// 	fmt.Println("hhd", res)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{"message": "data updated successfully!"})

// }
