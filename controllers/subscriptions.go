package controllers

import (
	"context"
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

var subsCollection *mongo.Collection = configs.GetCollection(configs.DB, "subs")

func CreateSub(c *gin.Context) {

	session, _ := c.Request.Cookie("Cookie_1")
	// fmt.Printf("headers: %v\n", cookies)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	var sub models.Subs
	defer cancel()

	if err := c.BindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	newSubs := models.Subs{
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
	c.JSON(http.StatusCreated, gin.H{"message": "Subs successfully", "Data": map[string]interface{}{"data": result}, session.Name: session.Value, "expiration": session.Expires})
}

func SubById(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	subId := c.Param("subId")
	defer cancel()

	var sub models.Subs

	objId, _ := primitive.ObjectIDFromHex(subId)

	err := subsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&sub)

	res := map[string]interface{}{"data": sub}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{err.Error(): err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success", "data": res})
}

func EditSub(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	subId := c.Param("subId")
	var subs models.Subs
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(subId)

	if err := c.BindJSON(&subs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	edited := bson.M{"name": subs.Name, "category": subs.Category, "custom_icons": subs.CustomIcon}

	result, err := subsCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": edited})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	// 	var updatedUser models.CreateSub
	// 	if result.MatchedCount == 1 {
	//         err := subsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)

	// 		if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data doesn't exist", "error": err})
	// 		return
	// 	}
	// }

	c.JSON(http.StatusCreated, gin.H{"message": "data updated successfully!", "Data": map[string]interface{}{"data": result}})
}

func DeleteSub(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	subId := c.Param("subId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(subId)

	result, err := subsCollection.DeleteOne(ctx, bson.M{"_id": objId})
	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No Sub to delete"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Subs deleted successfully", "Data": res})
}

func GetSubs(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var subs []models.Subs
	defer cancel()

	results, err := subsCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{err.Error(): err})
		return
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleSubs models.Subs
		if err = results.Decode(&singleSubs); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{err.Error(): err})
			return
		}

		subs = append(subs, singleSubs)
	}

	c.JSON(http.StatusOK, gin.H{"message": "data found successfully!", "Data": map[string]interface{}{"data": subs}})
}
