package routes

import (
	"ecargoware/alcochange-dtx/internals/services/goalSettingAssessmentService.go"
	"ecargoware/alcochange-dtx/sentryaccounts"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func goalSettingAssessment(router *httprouter.Router) {
	router.GET("/v1/acdtx/goalSettingAssessment", GetGoalSettingAssessment)
}

// GetGoalSettingAssessment func to send the Goal Setting Assessment to the client
func GetGoalSettingAssessment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	wl := goalSettingAssessmentService.NewGoalSettingAssessment(rd.l, rd.dbConn)
	res, errW := wl.GetGoalSettingAssessmentMessage()
	if errW != nil {
		rd.l.Errorf("GetGoalSettingAssessment - Error : ", errW.Error())
		sentryaccounts.SentryLogExceptions(errW)
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
