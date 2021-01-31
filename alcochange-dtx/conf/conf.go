package conf

var (
	ServerFlag = false
)

var (
	AppUrl                = "http://localhost"
	Port                  = 9010
	DatabaseUsername      = "alcochange"
	DatabasePassword      = "alco$2021$cha01ge"
	DatabaseName          = "cyberliver/_platform"
	DatabaseAddr          = "alcochange.cgdbfirlicf1.eu-west-2.rds.amazonaws.com:5432"
	ProductionMode        = false
	MaxConnectionPoolSize = 100
	ServiceAppName        = "AlcoChange"
)

// uncomment when production server is available
// func updateDBPassword() {
// 	DatabasePassword = "tranzopostgres"
// }

func InitateServerConfigurations() {
	// updateDBPassword()
	ProductionMode = true
}

// Button Text
var (
	WarningAndPrivacyButtonText  = "I agree to the Terms & Privacy"
	BaselineAssessmentButtonText = "Okay! Let's do this"
)
