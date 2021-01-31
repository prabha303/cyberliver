package models

import (
	"cyberliver/alcochange-dtx/utils"
	"strings"
	"time"
)

type UserAccess struct {
	ID                  int64     `json:"id"`
	UserID              int64     `json:"userID" sql:",notnull"`
	User                *User     `json:"user" pg:"joinFK:id"`
	UserTypeID          int64     `json:"userTypeID" sql:",notnull"`
	FirstName           string    `json:"firstName" validate:"required,min=3,max=50" sql:",notnull"`
	LastName            string    `json:"lastName" validate:"required,max=25" sql:",notnull"`
	MobileNo            string    `json:"mobileNo" validate:"unique=mobile_no" sql:",unique"`
	EmailID             string    `json:"emailID" validate:"unique=email_id" sql:",notnull,unique"`
	PatientAccessCode   string    `json:"patientAccessCode" validate:"unique=patient_access_code" sql:",unique"`
	CountryMobileCode   string    `json:"countryMobileCode"`
	PasswordStr         string    `json:"passwordStr,omitempty" validate:"required, min=6" sql:"-"`
	Password            string    `json:"-" validate:"required,min=3,max=50" sql:",notnull"`
	Timezone            string    `json:"timezone"`
	DeviceUUID          string    `json:"deviceUUID" sql:",notnull"`
	AlcoChangeDtxAccess bool      `json:"alcoChangeDtxAccess" sql:",notnull,default:false"`
	DryDayAccess        bool      `json:"drydayAccess" sql:",notnull,default:false"`
	LastLogin           time.Time `json:"lastLogin"`
	UserType            *UserType `json:"role" pg:"joinFK:id"`
	Version             int64     `json:"version" sql:",notnull,default:0"`
	IsActive            bool      `json:"isActive" sql:",notnull,default:false"`
	CreatedAt           time.Time `json:"createdAt" sql:",default:now()"`
	UpdatedAt           time.Time `json:"updatedAt" sql:",default:now()"`
}

func (user *UserAccess) BeforeInsert(zone string) {
	user.FirstName = utils.ToCamelCase(user.FirstName)
	user.LastName = utils.ToCamelCase(user.LastName)
	user.EmailID = strings.ToLower(user.EmailID)
	user.IsActive = true
	user.Version++
	currentTime, _ := utils.CurrentTimeWithZone(zone)
	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime
	user.LastLogin = currentTime
}
func (userAccess *UserAccess) BeforeUpdate(zone string) {
	currentTime, _ := utils.CurrentTimeWithZone(zone)
	userAccess.IsActive = true
	userAccess.Version++
	userAccess.UpdatedAt = currentTime
	userAccess.LastLogin = currentTime
}
