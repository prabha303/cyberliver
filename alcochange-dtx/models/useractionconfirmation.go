package models

import "time"

type UserActionConfirmation struct {
	ID                         int64     `json:"id"`
	DeviceUUID                 string    `json:"deviceUUID"`
	EmailID                    string    `json:"emailID"`
	WarningLabelRedeemed       bool      `json:"warningLabelRedeemed" sql:",notnull,default:false"`
	AccessCodeVerified         bool      `json:"accessCodeVerified" sql:",notnull,default:false"`
	TermsAndConditionsRedeemed bool      `json:"termsAndConditionsRedeemed" sql:",notnull,default:false"`
	IsSignedUp                 bool      `json:"isSignedUp" sql:",notnull,default:false"`
	Version                    int64     `json:"version"`
	IsActive                   bool      `json:"isActive"`
	CreatedAt                  time.Time `json:"createdAt" sql:",default:now()"`
	UpdatedAt                  time.Time `json:"updatedAt" sql:",default:now()"`
}
