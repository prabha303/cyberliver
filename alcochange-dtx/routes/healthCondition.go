package routes

import (
	"ecargoware/alcochange-dtx/internals/services/healthConditionService"
	"ecargoware/alcochange-dtx/sentryaccounts"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func healthConditionAssessment(router *httprouter.Router) {
	router.GET("/v1/acdtx/healthConditionAssessment", GetHealthConditionAssessment)
}

// GetHealthConditionAssessment func to send the Health Condition Assessment to the client
func GetHealthConditionAssessment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	wl := healthConditionService.NewHealthConditionAssessment(rd.l, rd.dbConn)
	res, errW := wl.GetHealthConditionAssessmentMessage()
	if errW != nil {
		rd.l.Errorf("GetHealthConditionAssessment - Error : ", errW.Error())
		sentryaccounts.SentryLogExceptions(errW)
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
