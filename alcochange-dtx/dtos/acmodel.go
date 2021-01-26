package dtos

type UserActionConfirmationReq struct {
	DeviceUUID                 string `json:"deviceUUID"`
	TimeZone                   string `json:"timeZone"`
	EmailID                    string `json:"emailID"`
	WarningLabelRedeemed       bool   `json:"warningLabelRedeemed"`
	AccessCodeVerified         bool   `json:"accessCodeVerified"`
	TermsAndConditionsRedeemed bool   `json:"termsAndConditionsRedeemed"`
}

type UserActionConfirmationResponse struct {
	Message string `json:"message"`
}
