package models

import (
	"ecargoware/alcochange-dtx/utils"
	"time"
)

type LoginLogs struct {
	ID             int64     `json:"id"`
	UserID         int64     `json:"userID"`
	User           *User     `json:"user" pg:"joinFK:id"`
	UserAppVersion int64     `json:"userAppVersion"`
	AppID          string    `json:"appID"`
	Latitude       float32   `json:"latitude"`
	Longitude      float32   `json:"longitude"`
	OsVersion      string    `json:"osVersion"`
	OsType         string    `json:"OsType"`
	DeviceUUID     string    `json:"deviceUUID"`
	DeviceInfo     string    `json:"deviceInfo"`
	NetworkInfo    string    `json:"networkInfo"`
	Version        int64     `json:"version" sql:",notnull,default:0"`
	IsActive       bool      `json:"isActive" sql:",notnull,default:false"`
	CreatedAt      time.Time `json:"createdAt" sql:",default:now()"`
	UpdatedAt      time.Time `json:"updatedAt" sql:",default:now()"`
}

func (loginLogs *LoginLogs) BeforeInsert(zone string) {
	loginLogs.IsActive = true
	loginLogs.Version++
	currentTime, _ := utils.CurrentTimeWithZone(zone)
	loginLogs.CreatedAt = currentTime
	loginLogs.UpdatedAt = currentTime
}
