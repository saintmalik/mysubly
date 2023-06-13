package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subs struct {
	Id             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	Category       string             `json:"category,omitempty" bson:"category,omitempty"`
	Website        string             `json:"website,omitempty" bson:"website,omitempty"`
	CustomIcon     string             `json:"custom_icons,omitempty" bson:"custom_icons,omitempty"`
	Color          string             `json:"color,omitempty" bson:"color,omitempty"`
	Cost           string             `json:"cost,omitempty" bson:"cost,omitempty"`
	ExpenseType    string             `json:"expense_type,omitempty" bson:"expense_type,omitempty"`
	BillingPeriod  string             `json:"billing_period,omitempty" bson:"billing_period,omitempty"`
	BillingNextPay string             `json:"billing_next_pay,omitempty" bson:"billing_next_pay,omitempty"`
	ReminderOne    string             `json:"reminder_one,omitempty" bson:"reminder_one,omitempty"`
	PaymentMethod  string             `json:"payment_method,omitempty" bson:"payment_method,omitempty"`
}

type EditSub struct {
	Id             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	Category       string             `json:"category,omitempty" bson:"category,omitempty"`
	Website        string             `json:"website,omitempty" bson:"website,omitempty"`
	CustomIcon     string             `json:"custom_icons,omitempty" bson:"custom_icons,omitempty"`
	Color          string             `json:"color,omitempty" bson:"color,omitempty"`
	Cost           string             `json:"cost,omitempty" bson:"cost,omitempty"`
	ExpenseType    string             `json:"expense_type,omitempty" bson:"expense_type,omitempty"`
	BillingPeriod  string             `json:"billing_period,omitempty" bson:"billing_period,omitempty"`
	BillingNextPay string             `json:"billing_next_pay,omitempty" bson:"billing_next_pay,omitempty"`
	ReminderOne    string             `json:"reminder_one,omitempty" bson:"reminder_one,omitempty"`
	PaymentMethod  string             `json:"payment_method,omitempty" bson:"payment_method,omitempty"`
}
