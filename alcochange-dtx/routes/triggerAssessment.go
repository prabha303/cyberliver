package routes

import (
	"ecargoware/alcochange-dtx/internals/services/triggerAssessmentService"
	"ecargoware/alcochange-dtx/sentryaccounts"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func triggerAssessment(router *httprouter.Router) {
	router.GET("/v1/acdtx/triggerAssessment", GetTriggerAssessment)
}

// GetTriggerAssessment func to send the Trigger Assessment to the client
func GetTriggerAssessment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	wl := triggerAssessmentService.NewTriggerAssessment(rd.l, rd.dbConn)
	res, errW := wl.GetTriggerAssessmentMessage()
	if errW != nil {
		rd.l.Errorf("GetTriggerAssessment - Error : ", errW.Error())
		sentryaccounts.SentryLogExceptions(errW)
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
