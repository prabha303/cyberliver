package models

import (
	"ecargoware/alcochange-dtx/utils"
	"log"
	"time"
)

type UserActionConfirmation struct {
	ID                         int64     `json:"id"`
	DeviceUUID                 string    `json:"deviceUUID" sql:",notnull"`
	EmailID                    string    `json:"emailID"`
	WarningLabelRedeemed       bool      `json:"warningLabelRedeemed" sql:",notnull,default:false"`
	AccessCodeVerified         bool      `json:"accessCodeVerified" sql:",notnull,default:false"`
	TermsAndConditionsRedeemed bool      `json:"termsAndConditionsRedeemed" sql:",notnull,default:false"`
	IsSignedUp                 bool      `json:"isSignedUp" sql:",notnull,default:false"`
	Version                    int64     `json:"version" sql:",notnull,default:0"`
	IsActive                   bool      `json:"isActive" sql:",notnull,default:false"`
	CreatedAt                  time.Time `json:"createdAt" sql:",notnull"`
	UpdatedAt                  time.Time `json:"updatedAt" sql:",notnull"`
}

func (userAction *UserActionConfirmation) BeforeUpdate(zone string) {
	currentTime, _ := utils.CurrentTimeWithZone(zone)
	userAction.IsActive = true
	userAction.Version++
	log.Println("currentTime--", zone, currentTime)
	userAction.UpdatedAt = currentTime
}

func (userAction *UserActionConfirmation) BeforeInsert(zone string) {
	currentTime, _ := utils.CurrentTimeWithZone(zone)
	userAction.Version++
	userAction.IsActive = true
	log.Println("currentTime-create-", zone, currentTime)
	userAction.CreatedAt = currentTime
	userAction.UpdatedAt = currentTime
}