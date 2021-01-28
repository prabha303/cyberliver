package models

import (
	"time"
)

type UserAccess struct {
	ID                  int64     `json:"id"`
	EmailID             string    `json:"emailID" validate:"unique=email_id" sql:","`
	MobileNo            string    `json:"mobileNo" validate:"unique=mobile_no" sql:",unique"`
	CountryMobileCode   string    `json:"countryMobileCode"`
	PasswordStr         string    `json:"passwordStr,omitempty" validate:"required, min=6" sql:"-"`
	Password            string    `json:"-" sql:",notnull"`
	Timezone            string    `json:"timezone"`
	UsersID             int64     `json:"usersID"`
	Users               *Users    `json:"users" pg:"joinFK:id"`
	AlcoChangeDtxAccess bool      `json:"alcoChangeDtxAccess"`
	DryDayAccess        bool      `json:"DrydayAccess"`
	AccessCode          string    `json:"accessCode"`
	LastLogin           time.Time `json:"lastLogin"`
	RoleID              int64     `json:"roleID" sql:",notnull"`
	Role                *Role     `json:"role" pg:"joinFK:id"`
	Version             int64     `json:"version" sql:",notnull,default:0"`
	IsActive            bool      `json:"isActive" sql:",notnull,default:false"`
	CreatedAt           time.Time `json:"createdAt" sql:",notnull"`
	UpdatedAt           time.Time `json:"updatedAt" sql:",notnull"`
}
