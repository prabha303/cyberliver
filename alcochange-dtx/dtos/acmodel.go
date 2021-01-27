package dtos

type UserActionConfirmationReq struct {
	DeviceUUID          string `json:"deviceUUID"`
	TimeZone            string `json:"timeZone"`
	EmailID             string `json:"emailID"`
	WarningLabelRead    bool   `json:"warningLabelRead"`
	AccessCodeVerified  bool   `json:"accessCodeVerified"`
	TermsAndPrivacyRead bool   `json:"termsAndPrivacyRead"`
}

type UserActionConfirmationResponse struct {
	Message string `json:"message"`
}
