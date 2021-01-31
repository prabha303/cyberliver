package models

import (
	"cyberliver/alcochange-dtx/utils"
	"time"
)

//Role struct
type UserType struct {
	ID          int64     `json:"id"`
	Code        string    `json:"code" sql:",notnull,unique=idx_code_name_role"`
	Name        string    `json:"name" validate:"required,min=1,max=50" sql:",notnull,unique=idx_code_name_role"`
	Description string    `json:"description" validate:"max=250"`
	Version     int64     `json:"version" sql:",notnull,default:0"`
	IsActive    bool      `json:"isActive" sql:",notnull,default:false"`
	CreatedAt   time.Time `json:"createdAt" sql:",notnull,default:now()"`
	UpdatedAt   time.Time `json:"updatedAt" sql:",notnull,default:now()"`
}

//BeforeInsert func
func (r *UserType) BeforeInsert(zone string) {
	currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Name = utils.ToCamelCase(r.Name)
	r.Version++
	r.IsActive = true
	r.CreatedAt = currentTime
	r.UpdatedAt = currentTime
}

//BeforeUpdate func
func (r *UserType) BeforeUpdate(zone string) {
	currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Name = utils.ToCamelCase(r.Name)
	r.Version++
	r.UpdatedAt = currentTime
}
