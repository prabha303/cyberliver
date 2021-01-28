package models

import "time"

//Users Fields
type Users struct {
	ID                int64     `json:"id"`
	FirstName         string    `json:"firstName" validate:"required,min=3,max=50" sql:",notnull"`
	MiddleName        string    `json:"middleName" validate:"max=25"`
	LastName          string    `json:"lastName" validate:"required,max=25" sql:",notnull"`
	DisplayName       string    `json:"displayName"`
	MobileNo          string    `json:"mobileNo" validate:"unique=mobile_no" sql:",unique"`
	EmailID           string    `json:"emailID" validate:"unique=email_id" sql:","`
	DOB               string    `json:"dob"`
	Gender            string    `json:"gender" sql:",notnull"`
	JoinedDate        time.Time `json:"joinedDate" sql:",notnull"`
	SocialID          string    `json:"socialID"`
	UUID              string    `json:"uuid" sql:",notnull"`
	AccessCode        string    `json:"accessCode"`
	LoggedSrc         string    `json:"loggedSrc"`
	SolutionType      string    `json:"solutionType"`
	RoleID            int64     `json:"roleID" sql:",notnull"`
	Role              *Role     `json:"role" pg:"joinFK:id"`
	AppID             string    `json:"appID"`
	Timezone          string    `json:"timezone"`
	CountryMobileCode string    `json:"countryMobileCode"`
	Lang              string    `json:"lang"`
	PatientAccessCode string    `json:"patientAccessCode"`
	AddressLine1      string    `json:"addressLine1" validate:"max=100"`
	AddressLine2      string    `json:"addressLine2" validate:"max=100"`
	City              string    `json:"city" validate:"max=50"`
	Area              string    `json:"area" validate:"max=50"`
	StateID           int64     `json:"stateID"`
	CountryCode       string    `json:"countryCode"`
	PostCode          string    `json:"postCode"`
	Archived          string    `json:"archived"`
	ArchivedDate      string    `json:"archivedDate"`
	Version           int64     `json:"version" sql:",notnull,default:0"`
	IsActive          bool      `json:"isActive" sql:",notnull,default:false"`
	CreatedAt         time.Time `json:"createdAt" sql:",notnull"`
	UpdatedAt         time.Time `json:"updatedAt" sql:",notnull"`
}
