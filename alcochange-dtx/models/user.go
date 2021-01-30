package models

import (
	"ecargoware/alcochange-dtx/utils"
	"strings"
	"time"
)

type User struct {
	ID                int64     `json:"id"`
	FirstName         string    `json:"firstName" validate:"required,min=3,max=50" sql:",notnull"`
	MiddleName        string    `json:"middleName" validate:"max=25"`
	LastName          string    `json:"lastName" validate:"required,max=25" sql:",notnull"`
	DisplayName       string    `json:"displayName"`
	MobileNo          string    `json:"mobileNo" validate:"unique=mobile_no" sql:",unique"`
	EmailID           string    `json:"emailID" validate:"unique=email_id" sql:",notnull,unique"`
	PatientAccessCode string    `json:"patientAccessCode" validate:"unique=patient_access_code" sql:",unique"`
	DOB               string    `json:"dob"`
	Gender            string    `json:"gender" sql:",notnull"`
	JoinedDate        time.Time `json:"joinedDate" sql:",notnull"`
	SocialID          string    `json:"socialID"`
	DeviceUUID        string    `json:"deviceUUID" sql:",notnull"`
	EthnicityID       int64     `json:"ethnicityID" sql:",default:0"`
	LoggedSrc         string    `json:"loggedSrc"`
	SolutionType      string    `json:"solutionType"`
	UserTypeID        int64     `json:"userTypeID" sql:",notnull"`
	UserType          *UserType `json:"role" pg:"joinFK:id"`
	AppID             string    `json:"appID"`
	Timezone          string    `json:"timezone"`
	CountryMobileCode string    `json:"countryMobileCode"`
	Lang              string    `json:"lang"`
	AddressLine1      string    `json:"addressLine1" validate:"max=100"`
	AddressLine2      string    `json:"addressLine2" validate:"max=100"`
	City              string    `json:"city" validate:"max=50"`
	Area              string    `json:"area" validate:"max=50"`
	StateID           int64     `json:"stateID"`
	CountryCode       string    `json:"countryCode"`
	PostCode          string    `json:"postCode"`
	Archived          bool      `json:"archived"`
	ArchivedDate      time.Time `json:"archivedDate"`
	Version           int64     `json:"version" sql:",notnull,default:0"`
	IsActive          bool      `json:"isActive" sql:",notnull,default:false"`
	CreatedAt         time.Time `json:"createdAt" sql:",default:now()"`
	UpdatedAt         time.Time `json:"updatedAt" sql:",default:now()"`
}

func (user *User) BeforeInsert(zone string) {
	user.FirstName = utils.ToCamelCase(user.FirstName)
	// emp.MiddleName = utils.ToCamelCase(emp.MiddleName)
	user.LastName = utils.ToCamelCase(user.LastName)
	user.EmailID = strings.ToLower(user.EmailID)
	user.Gender = user.SetGender()
	user.IsActive = true
	user.Version++
	currentTime, _ := utils.CurrentTimeWithZone(zone)
	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime
}

func (user *User) SetGender() string {
	g := strings.ToUpper(user.Gender)
	if g == "MALE" {
		return "Male"
	}
	if g == "FEMALE" {
		return "Female"
	}

	if g == "TRANSGENDER" {
		return "Transgender"
	}
	return ""
}
