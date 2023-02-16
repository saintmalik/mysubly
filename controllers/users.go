package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saintmalik/mysubly/configs"
	"github.com/saintmalik/mysubly/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var subsCollection *mongo.Collection = configs.GetCollection(configs.DB, "subs")

func CreateSubs(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	var sub models.CreateSub
	defer cancel()

	if err := c.BindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	newSubs := models.CreateSub{
		Name:           sub.Name,
		Category:       sub.Category,
		Website:        sub.Website,
		CustomIcon:     sub.CustomIcon,
		Color:          sub.Color,
		Cost:           sub.Cost,
		ExpenseType:    sub.ExpenseType,
		BillingPeriod:  sub.BillingPeriod,
		BillingNextPay: sub.BillingNextPay,
		ReminderOne:    sub.ReminderOne,
		PaymentMethod:  sub.PaymentMethod,
	}

	result, err := subsCollection.InsertOne(ctx, newSubs)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Subs successfully", "Data": map[string]interface{}{"data": result}})
}
