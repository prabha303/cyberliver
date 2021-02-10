package dbscripts

var countryID int64

//CreateDefaultValues function use to call mentioned function
func CreateDefaultValues() {
	// createCountry()
}

// //createCountry function use to create default country
// func createCountry() {
// 	db := dbcon.Get()
// 	country := &models.Country{}

// 	if err := db.Model(country).Where("name=?", "India").Select(); err != nil {
// 		log.Println("Error to get country", err.Error())
// 	}

// 	if country != nil && country.ID > 0 {
// 		countryID = country.ID
// 		return
// 	}

// 	country = &models.Country{
// 		Name:     "India",
// 		Code:     "IN",
// 		DialCode: "+91",
// 	}
// 	country.IsActive = true
// 	country.CreatedAt = time.Now().UTC()
// 	country.UpdatedAt = time.Now().UTC()

// 	if _, err := db.Model(country).Insert(); err != nil {
// 		log.Println("Error to insert default Country.", err.Error())
// 		return
// 	}
// 	countryID = country.ID
// 	log.Println("Country created successfully.")
// }
