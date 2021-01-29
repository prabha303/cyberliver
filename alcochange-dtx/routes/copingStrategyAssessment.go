package routes

import (
	"ecargoware/alcochange-dtx/internals/services/copingStrategyAssessmentService"
	"ecargoware/alcochange-dtx/sentryaccounts"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func copingStrategyAssessment(router *httprouter.Router) {
	router.GET("/v1/acdtx/copingStrategyAssessment", GetCopingStrategyAssessment)
}

// GetCopingStrategyAssessment func to send the CopingStrategy Assessment to the client
func GetCopingStrategyAssessment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	wl := copingStrategyAssessmentService.NewCopingStrategyAssessment(rd.l, rd.dbConn)
	res, errW := wl.GetCopingStrategyAssessmentMessage()
	if errW != nil {
		rd.l.Errorf("GetCopingStrategyAssessment - Error : ", errW.Error())
		sentryaccounts.SentryLogExceptions(errW)
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
