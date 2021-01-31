package routes

import (
	"cyberliver/alcochange-dtx/internals/services/reasonAssessmentService"
	"cyberliver/alcochange-dtx/sentryaccounts"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func reasonAssessment(router *httprouter.Router) {
	router.GET("/v1/acdtx/reasonAssessment", GetReasonAssessment)
}

// GetReasonAssessment func to send the Reason Assessment to the client
func GetReasonAssessment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	wl := reasonAssessmentService.NewReasonAssessment(rd.l, rd.dbConn)
	res, errW := wl.GetReasonAssessmentMessage()
	if errW != nil {
		rd.l.Errorf("GetReasonAssessment - Error : ", errW.Error())
		sentryaccounts.SentryLogExceptions(errW)
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
