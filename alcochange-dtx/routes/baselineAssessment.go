package routes

import (
	"ecargoware/alcochange-dtx/internals/services/baselineAssessmentService"

	"ecargoware/alcochange-dtx/sentryaccounts"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func baselineAssessment(router *httprouter.Router) {
	router.GET("/v1/acdtx/baselineAssessment", GetBaselineAssessment)
}

// GetBaselineAssessment func to send the BaselineAssessment to the client
func GetBaselineAssessment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	wl := baselineAssessmentService.NewBaselineAssessment(rd.l, rd.dbConn)
	res, errW := wl.GetBaselineAssessmentMessage()
	if errW != nil {
		rd.l.Errorf("GetBaselineAssessment - Error : ", errW.Error())
		sentryaccounts.SentryLogExceptions(errW)
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
