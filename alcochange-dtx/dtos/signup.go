package dtos

type SignUpRequest struct {
	CountryCode         string              `json:"countryCode"`
	Timezone            string              `json:"timezone"`
	Lang                string              `json:"lang"`
	RegisterUserRequest RegisterUserRequest `json:"registerUserRequest"`
	LoginDeviceDetails  LoginDeviceDetails  `json:"loginDeviceDetails"`
}

type RegisterUserRequest struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	EmailID    string `json:"emailID"`
	AccessCode string `json:"accessCode"`
	Password   string `json:"password"`
	Gender     string `json:"gender"`
	//TermsAndConditions int     `json:"termsAndConditions"`
	SolutionType     string  `json:"solutionType"`
	DOB              string  `json:"dob"`
	Latitude         float32 `json:"latitude"`
	Longitude        float32 `json:"longitude"`
	AppID            string  `json:"appID"`
	LoggedSrc        string  `json:"loggedSrc"`
	SignUpFor        string  `json:"signUpFor"`
	ProductAccessFor string  `json:"productAccessFor"`
}
type LoginDeviceDetails struct {
	OsVersion   string `json:"osVersion"`
	OsType      string `json:"OsType"`
	DeviceUUID  string `json:"deviceUUID"`
	DeviceInfo  string `json:"deviceInfo"`
	NetworkInfo string `json:"networkInfo"`
}

type SignUpResponse struct {
	DeviceUUID  string `json:"deviceUUID"`
	JoinedDate  string `json:"joinedDate"`
	TokenID     string `json:"tokenID"`
	CitiesID    string `json:"citiesID"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DisplayName string `json:"displayName"`
	LoggedSrc   string `json:"loggedSrc"`
	AccessToken string `json:"accessToken"`
	Email       string `json:"email"`
}
