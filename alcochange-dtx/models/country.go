package models

import (
	"cyberliver/alcochange-dtx/utils"
	"time"
)

//Country provides entities for an Country structure
type Country struct {
	ID                int64     `json:"id"`
	Name              string    `json:"name" sql:",notnull,unique"`
	Code              string    `json:"code" sql:",notnull,unique"`
	DialCode          string    `json:"dialCode" sql:",notnull,unique"`
	CurrencySymbol    string    `json:"currencySymbol"`
	CurrenyName       string    `json:"currenyName"`
	CurrencyShortCode string    `json:"currencyShortCode"`
	Timezone          string    `json:"timezone"`
	IsActive          bool      `json:"isActive"`
	CreatedAt         time.Time `json:"-" sql:",default:now()"`
	UpdatedAt         time.Time `json:"-" sql:",default:now()"`
}

// type CountryMinified struct {
// 	ID   int64  `json:"id"`
// 	Name string `json:"name" sql:",notnull,unique"`
// }

//BeforeInsert func
func (r *Country) BeforeInsert(zone string) {
	currentTime, _ := utils.CurrentTimeWithZone(zone)
	r.Name = utils.ToCamelCase(r.Name)
	r.IsActive = true
	r.CreatedAt = currentTime
	r.UpdatedAt = currentTime
}
