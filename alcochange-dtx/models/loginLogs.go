package models

import (
	"time"
)

type LoginLogs struct {
	ID          int64     `json:"id"`
	UsersID     int64     `json:"usersID"`
	Users       *Users    `json:"users" pg:"joinFK:id"`
	Latitude    float32   `json:"latitude"`
	Longitude   float32   `json:"longitude"`
	OsVersion   string    `json:"osVersion"`
	OsType      string    `json:"OsType"`
	DeviceUUID  string    `json:"deviceUUID"`
	DeviceInfo  string    `json:"deviceInfo"`
	NetworkInfo string    `json:"networkInfo"`
	Version     int64     `json:"version" sql:",notnull,default:0"`
	IsActive    bool      `json:"isActive" sql:",notnull,default:false"`
	CreatedAt   time.Time `json:"createdAt" sql:",notnull"`
	UpdatedAt   time.Time `json:"updatedAt" sql:",notnull"`
}
