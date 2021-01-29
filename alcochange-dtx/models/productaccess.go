package models

import (
	"ecargoware/alcochange-dtx/utils"
	"time"
)

type ProductAccess struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required,min=1,max=50" sql:",notnull,unique=idx_code_name_pa"`
	Code      string    `json:"code" sql:",notnull,unique=idx_code_name_pa"`
	Version   int64     `json:"version" sql:",notnull,default:0"`
	IsActive  bool      `json:"isActive" sql:",notnull,default:false"`
	CreatedAt time.Time `json:"createdAt" sql:",default:now()"`
	UpdatedAt time.Time `json:"updatedAt" sql:",default:now()"`
}

func (r *ProductAccess) BeforeInsert(zone string) {
	currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.IsActive = true
	r.CreatedAt = currentTime
	r.UpdatedAt = currentTime
}

func (r *ProductAccess) BeforeUpdate(zone string) {
	currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Version++
	r.UpdatedAt = currentTime
}
