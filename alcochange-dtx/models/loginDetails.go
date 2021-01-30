package models

import (
	"ecargoware/alcochange-dtx/utils"
	"time"
)

type LoginDeviceDetails struct {
	ID             int64     `json:"id"`
	UserID         int64     `json:"userID"`
	User           *User     `json:"user" pg:"joinFK:id"`
	Latitude       float32   `json:"latitude"`
	Longitude      float32   `json:"longitude"`
	AppID          string    `json:"appID"`
	LastLogin      time.Time `json:"lastLogin"`
	OsVersion      string    `json:"osVersion"`
	UserAppVersion int64     `json:"userAppVersion"`
	OsType         string    `json:"OsType"`
	DeviceUUID     string    `json:"deviceUUID"`
	DeviceInfo     string    `json:"deviceInfo"`
	NetworkInfo    string    `json:"networkInfo"`
	Timezone       string    `json:"timezone"`
	Version        int64     `json:"version" sql:",notnull,default:0"`
	IsActive       bool      `json:"isActive" sql:",notnull,default:false"`
	CreatedAt      time.Time `json:"createdAt" sql:",default:now()"`
	UpdatedAt      time.Time `json:"updatedAt" sql:",default:now()"`
}

func (user *LoginDeviceDetails) BeforeInsert(zone string) {
	user.IsActive = true
	user.Version++
	currentTime, _ := utils.CurrentTimeWithZone(zone)
	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime
	user.LastLogin = currentTime
}

func (loginDeviceDetails *LoginDeviceDetails) BeforeUpdate(zone string) {
	currentTime, _ := utils.CurrentTimeWithZone(zone)
	loginDeviceDetails.IsActive = true
	loginDeviceDetails.Version++
	loginDeviceDetails.UpdatedAt = currentTime
	loginDeviceDetails.LastLogin = currentTime
}
