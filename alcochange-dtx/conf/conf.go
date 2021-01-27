package conf

var (
	ServerFlag = false
)

var (
	//MSSubscriptionKey = "dbd9d1a29eb74da09e91c6e36d1f18cb"
	//MSBaseURL         = "https://facetrack.cognitiveservices.azure.com/face/v1.0"
	MSSubscriptionKey = "1e9ec381a31244498ed0f5319c631954"
	MSBaseURL         = "https://facetrackstaging.cognitiveservices.azure.com/face/v1.0"
	MSPersonsGroupURL = MSBaseURL + "/persongroups/"
)

var (
	AppUrl                = "http://localhost"
	Port                  = 9010
	DatabaseUsername      = "postgres"
	DatabasePassword      = "postgres"
	DatabaseName          = "cyberliver_platform"
	DatabaseAddr          = "localhost:5432"
	UserURL               = "http://localhost:9002"
	ProductionMode        = false
	MaxConnectionPoolSize = 100
	ServiceAppName        = "AlcoChange"
)

var (
	PaginationLimit = 5
)

func updateDBPassword() {
	DatabasePassword = "tranzopostgres"
}

func updateMSSubcrition() {
	MSSubscriptionKey = "dbd9d1a29eb74da09e91c6e36d1f18cb"
	MSBaseURL = "https://facetrack.cognitiveservices.azure.com/face/v1.0"
}

func InitateServerConfigurations() {
	updateDBPassword()
	updateMSSubcrition()
	ProductionMode = true
}

// Button Text
var (
	WarningAndPrivacyButtonText  = "I agree to the Terms & Privacy"
	BaselineAssessmentButtonText = "Okay! Let's do this"
)
