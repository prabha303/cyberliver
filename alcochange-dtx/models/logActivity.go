package models

import (
	"cyberliver/alcochange-dtx/dbcon"
	"time"
)

// Activity logger
type APILogActivity struct {
	ID          int64       `json:"id"`
	UserID      int64       `json:"userid"`
	User        string      `json:"user"`
	Resource    string      `json:"resource"`
	Action      string      `json:"action"`
	RequestBody interface{} `json:"requestBody"`
	Response    interface{} `json:"response"`
	CreatedAt   time.Time   `json:"createdAt"`
}

func SaveAPILogActivity(api *APILogActivity) {

	db := dbcon.Get()
	db.Model(api).Insert()
}
