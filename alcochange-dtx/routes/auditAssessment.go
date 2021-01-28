package routes

import (
	"ecargoware/alcochange-dtx/internals/services/auditAssessmentService"
	"ecargoware/alcochange-dtx/sentryaccounts"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func auditAssessment(router *httprouter.Router) {
	router.GET("/v1/acdtx/auditAssessment", GetAuditAssessment)
}

// GetAuditAssessment func to send the Audit Assessment to the client
func GetAuditAssessment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	wl := auditAssessmentService.NewAuditAssessment(rd.l, rd.dbConn)
	res, errW := wl.GetAuditAssessmentMessage()
	if errW != nil {
		rd.l.Errorf("GetAuditAssessment - Error : ", errW.Error())
		sentryaccounts.SentryLogExceptions(errW)
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
