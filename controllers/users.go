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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateSub(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var sub models.CreateSub
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

	result, err := configs.SubsCollection.InsertOne(ctx, newSubs)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Sub created successfully", "data": result})
}

func Sub(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var subs []models.EditSub
	cursor, err := configs.SubsCollection.Find(ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
	err = cursor.All(ctx, &subs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success", "data": subs})
}

func SubById(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	subId := c.Param("subid")
	var sub models.EditSub

	Id, err := primitive.ObjectIDFromHex(subId)
	if err != nil {
		log.Fatal(err)
	}

	err = configs.SubsCollection.FindOne(ctx, bson.M{"_id": Id}).Decode(&sub)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("sub with id %s not found", subId)})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success", "data": sub})
}

func DeleteSub(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	subId := c.Param("subid")

	Id, err := primitive.ObjectIDFromHex(subId)
	if err != nil {
		log.Fatal(err)
	}

	_, err = configs.SubsCollection.DeleteOne(ctx, bson.M{"_id": Id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
